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
	q := client.NewQuery("CREATE DATABASE "+os.Getenv("DATABASE"), "", "")
	fmt.Println(q)
	response, err := c.Query(q)
	if err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	} else {
		panic("Response Error:" + response.Error().Error() + "; Error:" + err.Error())
	}

	//Create database
	_, err = QueryDB(c, fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DATABASE")))
	if err != nil{
		log.Fatal(err)
	}

	//Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  os.Getenv("DATABASE"),
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"Colour": "Hue"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	//Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	WritePoints(c)
	// QueryDB(c,)

	fmt.Println("End of execution")

}

//WritePoints writes points to influxDB
func WritePoints(c client.Client) {

	//Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  os.Getenv("DATABASE"),
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	sampleSize := 1000
	for i := 0; i < sampleSize; i++ {
		regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
		tags := map[string]string{
			"cpu":    "cpu-total",
			"host":   fmt.Sprintf("host %d", rand.Intn(1000)),
			"region": regions[rand.Intn(len(regions))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle": idle,
			"busy": 100.0 - idle,
		}

		pt, err := client.NewPoint(
			"cpu_usage",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)

	}
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

}

//QueryDB function to query the database
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
