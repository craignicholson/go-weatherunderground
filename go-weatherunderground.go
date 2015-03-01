// This application will query WeatherUnderground API for weather data
// and post to a http service.  WeatherUnderground's data
// is more of a daily pull for most of the queries instead of
// an hourly pull.  The one cavet to daily vs. hourly is we might
// need updated hourly profiles for scada data to replace the forcast
// weather data through out the day.


package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Configuration struct {
	ApiKey string `json:"apikey"`
	Uri    string `json:"uri"`
}


func main() {

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	apikey := config.ApiKey
	uri := config.Uri
	fmt.Println(apikey)

	// Returns an hourly forecast for the next 36 hours immediately following the API request.
	const hourlyEndPoint = "/hourly/q/TX/Dallas.json"

	// Returns an hourly forecast for the next 10 days
	const hourlyEndPointForcast = "/hourly10day/q/CA/San_Francisco.json"

	// Example Query using US ZipCode
	const hourlyEndPointZipCode = "/hourly10day/q/75219.json"

	// Example Query using airport code
	const hourlyEndPointAirportCode = "/hourly10day/q/KJFK.json"

	// Example using County/City {Japan/Tokyo}
	const hourlyEndPointHistoryInter = "/geolookup/conditions/forecast/q/Japan/Tokyo.json"

	// History_YYYYMMDD returns a summary of the observed weather for the specified date.
	const hourlyEndPointHistory = "/history_20150227/q/TX/Dallas.json"

	query := hourlyEndPointHistory

	fmt.Println(uri+apikey+query, "\n")

	//Run the Api query to fetch the weather data
	response, err := http.Get(uri + apikey + query)
	if err != nil {
		log.Fatal(err)
	}

	// Read all the weather data for processing
	// to be added to our own datastore for applications
	weatherdata, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//If we do not use the %s, the bytes are not decoded for the Printf
	//fmt.Printf("%s", weatherdata)

	// If we have weatherdata, post to the server (server.go)
	// TODO: add tests and a config value to not post the value to the http server
	if len(weatherdata) > 1 {
		resp, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer(weatherdata))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.Status)
	}
}
