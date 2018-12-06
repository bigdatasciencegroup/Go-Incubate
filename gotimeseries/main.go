package main

import (
	"fmt"
	"log"
	"os"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
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

	//Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  os.Getenv("DATABASE"),
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.Addpoint(pt)

	//Write the batch
	if err:= c.Write(bp); err != nil{
		log.Fatal(err)
	}

	fmt.Println("End of execution")

}
