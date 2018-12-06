package main

import (
	"fmt"
	"os"

	client "github.com/influxdata/influxdb/client/v2"
)

func main() {

	// Create client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     os.Getenv("TSDBADDR"),
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	})
	if err != nil {
		panic("Error creating InfluxDB Client: " + err.Error())
	}
	defer c.Close()

	// Create database
	q := client.NewQuery("CREATE DATABASE telegraf","","")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}

	fmt.Println("End of execution")

}
