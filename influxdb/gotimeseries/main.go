package main

import (
	// "encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"os"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
	"gocv.io/x/gocv"
)

func main() {

	// Create client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     os.Getenv("TSDBADDR"),
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PASS"),
	})
	if err != nil {
		panic("Error creating InfluxDB Client: " + err.Error())
	}
	defer c.Close()

	//Create database
	_, err = QueryDB(c, fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DATABASE")))
	if err != nil {
		log.Fatal(err)
	}

	//Set retention policy
	_, err = QueryDB(c, fmt.Sprintf("CREATE RETENTION POLICY short ON %s DURATION 1h REPLICATION 1 DEFAULT", os.Getenv("DATABASE")))
	if err != nil {
		log.Fatal(err)
	}

	// Capture video
	webcam, err := gocv.OpenVideoCapture(os.Getenv("RTSPLINK"))
	if err != nil {
		panic("Error in opening webcam: " + err.Error())
	}

	// ii := 0

	// Stream images from RTSP to Kafka message queue
	frame := gocv.NewMat()
	for {
		if !webcam.Read(&frame) {
			continue
		}

		// Type assert frame into RGBA image
		imgInterface, err := frame.ToImage()
		if err != nil {
			panic(err.Error())
		}
		img, ok := imgInterface.(*image.RGBA)
		if !ok {
			panic("Type assertion of pic (type image.Image interface) to type image.RGBA failed")
		}

		//Form the struct to be sent to Kafka message queue
		doc := Result{
			Pix:      img.Pix,
			Channels: frame.Channels(),
			Rows:     frame.Rows(),
			Cols:     frame.Cols(),
		}

		//Prepare message to be sent to Kafka
		docBytes, err := json.Marshal(doc)
		if err != nil {
			log.Fatal("Json marshalling error. Error:", err.Error())
		}

		//Write data to InfluxDB
		tags := map[string]string{}
		fields := map[string]interface{}{
			"frame": string(docBytes),
		}
		writeTime := time.Now()
		WritePoints(c, "VideoFrames", tags, fields, writeTime)

		log.Printf("---->>>> %v\n", time.Now())

		// if ii >= 10 {
		// 	break
		// }
		// ii = ii + 1
	}

	log.Println("End of execution")

}

//WritePoints function writes points to InfluxDB
func WritePoints(c client.Client, measurement string, tags map[string]string, fields map[string]interface{}, writeTime time.Time) {

	//Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: os.Getenv("DATABASE"),
	})
	if err != nil {
		log.Fatal(err)
	}

	//Create newpoint and add to batch
	pt, err := client.NewPoint(
		measurement,
		tags,
		fields,
		writeTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	//Write batch to database
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

}

//QueryDB function queries the database
func QueryDB(c client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: os.Getenv("DATABASE"),
	}
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}

	return res, nil
}

//Result represents the Kafka queue message format
type Result struct {
	Pix      []byte `json:"pix"`
	Channels int    `json:"channels"`
	Rows     int    `json:"rows"`
	Cols     int    `json:"cols"`
}
