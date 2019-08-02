package main

import (
	"encoding/json"
	"image"
	"image/color"
	"log"
	"models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gocv.io/x/gocv"
)

var statusColor = color.RGBA{200, 150, 50, 0}
var bkgColor = color.RGBA{255, 255, 255, 0}

func message(ev *kafka.Message) error {

	//Read message into `topicMsg` struct
	doc := &topicMsg{}
	err := json.Unmarshal(ev.Value, doc)
	if err != nil {
		log.Println(err)
		return err
	}

	// Retrieve frame
	log.Printf("%% Message sent %v on %s\n", ev.Timestamp, ev.TopicPartition)
	frame, err := gocv.NewMatFromBytes(doc.Rows, doc.Cols, doc.Type, doc.Mat)
	if err != nil {
		log.Println("Frame:", err)
		return err
	}

	frameOut := frame.Clone()

	// Form output image
	for ind := 0; ind < len(modelParams); ind++ {
		mp := modelParams[ind]

		// Get prediction
		res, err := mp.modelHandler.Get()
		if err == nil {
			mp.pred = res.Class
		}

		// Write prediction to frame
		gocv.PutText(
			&frameOut,
			mp.pred,
			image.Pt(10, ind*20+20),
			gocv.FontHersheyPlain, 1.2,
			bkgColor, 6,
		)
		// Write prediction to frame
		gocv.PutText(
			&frameOut,
			mp.pred,
			image.Pt(10, ind*20+20),
			gocv.FontHersheyPlain, 1.2,
			statusColor, 2,
		)

		// Post next frame
		mp.modelHandler.Post(models.Input{Img: frame})
	}

	// Write image to output Kafka queue
	select {
	case videoDisplay <- frameOut:
	default:
	}
	// videoDisplay <- frameOut

	return nil
}

type topicMsg struct {
	Mat      []byte       `json:"mat"`
	Channels int          `json:"channels"`
	Rows     int          `json:"rows"`
	Cols     int          `json:"cols"`
	Type     gocv.MatType `json:"type"`
}
