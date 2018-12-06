package main

import (
	"os"

	client "github.com/influxdata/influxdb/client/v2"
)

func main() {

	// Create a client	
	_, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: os.Getenv("TSDBADDR"),
        Username: os.Getenv("INFLUX_USER"),
        Password: os.Getenv("INFLUX_PWD"),
    })
    if err != nil {		
        panic("Error creating InfluxDB Client: "+err.Error())
	}

}
