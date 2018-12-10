package main

import (
	"fmt"
	"log"
	"math/rand"
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

	//Create database
	_, err = QueryDB(c, fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DATABASE")))
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	regions := []string{"distracted", "attentive", "neutral"}
	for {
		//Create random data to be written
		tags := map[string]string{
			"state": regions[rand.Intn(len(regions))],
		}
		facialFeature := rand.Float64() * 100.0
		imageColour := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"facialFeature": facialFeature,
			"imageColour":   imageColour,
		}

		writeTime := time.Now()
		WritePoints(c, "AverageImagePixel", tags, fields, writeTime)
		time.Sleep(1000 * time.Millisecond)
	}

	// q := fmt.Sprintf("SELECT count(%s) FROM %s", "busy", "cpu_usage")
	// res, err := QueryDB(c, q)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

	fmt.Println("End of execution")

}

//WritePoints function writes points to InfluxDB
func WritePoints(c client.Client, measurement string, tags map[string]string, fields map[string]interface{}, writeTime time.Time) {

	//Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  os.Getenv("DATABASE"),
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	//Create newpoint and to batch
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
