package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"net/http"

	"gocv.io/x/gocv"
)

type emonet struct {
	baseHandler
	faceCascade gocv.CascadeClassifier
}

//NewEmonet returns a new handle to specified machine learning model
func NewEmonet(modelurl string, labelurl string) (Handler, error) {

	labels := make(map[int]string)

	// Read-in labels
	dat, err := ioutil.ReadFile(labelurl)
	if err != nil {
		return &emonet{}, errors.New("Failed to read in labelurl. " + err.Error())
	}
	err = json.Unmarshal(dat, &labels)
	if err != nil {
		return &emonet{}, errors.New("Failure in unmarshalling labels. " + err.Error())
	}

	// load classifier to recognize faces
	faceCascade := gocv.NewCascadeClassifier()
	xmlFile := "/go/src/app/vendor/models/haarcascade_frontalface_default.xml"
	if !faceCascade.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
	}

	return &emonet{
		baseHandler: baseHandler{
			labels: labels,
			url:    modelurl,
			chIn:   make(chan Input),
			chOut:  make(chan Output),
		},
		faceCascade: faceCascade,
	}, nil
}

//Predict classifies input images
func (emn *emonet) Predict() {
	var resBody responseBodyTensor
	defer func() {
		if r := recover(); r != nil {
			log.Println("models.*emonet.Predict():PANICKED AND RESTARTING")
			log.Println("Panic:", r)
			go emn.Predict()
		}
	}()

	//Write initial prediction into shared output channel
	emn.chOut <- Output{Class: "Nothing"}

	// prepare image matrix
	minSize := image.Point{X: 48, Y: 48}
	maxSize := image.Point{X: 50000, Y: 50000}
	sz := image.Point{X: 48, Y: 48}

	for elem := range emn.chIn {
		img := elem.Img

		// Convert img to gray scale
		gocv.CvtColor(img, &img, gocv.ColorConversionCode(6))

		// Detect faces
		rects := emn.faceCascade.DetectMultiScaleWithParams(img, 1.3, 5, 0, minSize, maxSize)
		fmt.Printf("found %d faces\n", len(rects))

		classArr := []string{}
		for _, r := range rects {
			//Crop and resize image
			imgFace := img.Region(r)
			imgFaceDet := imgFace.Clone()
			gocv.Resize(imgFaceDet, &imgFaceDet, sz, 0, 0, 1)
			// imgFaceDet.DivideFloat(float32(255))
			imgFaceDet.ConvertTo(&imgFaceDet, gocv.MatType(5))

			// window1 := gocv.NewWindow("OrigFace")
			// window2 := gocv.NewWindow("CropFace")
			// // defer window.Close()
			// window2.IMShow(imgFaceDet)
			// window2.WaitKey(1)
			// window1.IMShow(imgGray)
			// window1.WaitKey(1)

			imgTensor, err := (&imgFaceDet).DataPtrFloat32()
			if err != nil {
				log.Println("cv mat error ", err)
			}

			// sz.X is columns
			// sz.Y is rows
			var imageTensorArr [48][48][1]float32
			for row := 0; row < sz.Y; row++ {
				for col := 0; col < sz.X; col++ {
					imageTensorArr[row][col][0] = imgTensor[row*48+col]
				}
			}

			//Prepare request message
			inference := inferTensor{
				InstancesTensor: []instanceTensor{
					instanceTensor{ImgTensor: imageTensorArr},
				},
			}

			//Query the machine learning model
			reqBody, err := json.Marshal(inference)
			if err != nil {
				log.Println("Error in Marshal: ", err)
				continue
			}
			req, err := http.NewRequest("POST", emn.url, bytes.NewBuffer(reqBody))
			if err != nil {
				log.Println("Error in NewRequest: ", err)
				continue
			}
			req.Header.Add("Content-Type", "application/json")
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("Error in DefaultClient: ", err)
				continue
			}
			defer res.Body.Close()

			// body, err := ioutil.ReadAll(res.Body)
			// fmt.Println(string(body))

			//Process response from machine learning model
			decoder := json.NewDecoder(res.Body)
			if err := decoder.Decode(&resBody); err != nil {
				log.Println("Error in Decode: ", err)
				continue
			}
			predClass := resBody.PredictionsTensor[0]
			valMax := float32(0)
			indexMax := 6
			for index, val := range predClass {
				if val > valMax {
					valMax = val
					indexMax = index
				}
			}
			pred, _ := emn.labels[indexMax]
			classArr = append(classArr, pred)
		}

		//Write prediction into shared output channel
		emn.chOut <- Output{Rects: rects, ClassArr: classArr}
	}
}

type inferTensor struct {
	InstancesTensor []instanceTensor `json:"instances"`
}

type instanceTensor struct {
	ImgTensor [48][48][1]float32 `json:"conv2d_1_input"` //json name must end with `_bytes` for json to know that it is binary data
}

type responseBodyTensor struct {
	PredictionsTensor [][]float32 `json:"predictions"`
}
