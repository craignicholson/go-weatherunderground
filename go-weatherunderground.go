package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	ApiKey string
	Uri    string
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
	// hourlyEndPoint := uri+apikey+"/hourly/q/CA/San_Francisco.json"

	// Returns an hourly forecast for the next 10 days
	// hourlyEndPointForcast := uri++apikey+"/hourly10day/q/CA/San_Francisco.json"

	// Example Query using US ZipCode
	hourlyEndPointZipCode := uri + apikey + "/hourly10day/q/75219.json"

	// Example Query using airport code
	// hourlyEndPointAirportCode :=uri+apikey+"/hourly10day/q/KJFK.json"

	// Example using County/City {Japan/Tokyo}
	// hourlyEndPointHistoryInter := uri+apikey+"/geolookup/conditions/forecast/q/Japan/Tokyo.json"

	// History_YYYYMMDD returns a summary of the observed weather for the specified date.
	// hourlyEndPointHistory :=uri+apikey+"/history_19701127/q/CA/San_Francisco.json"

	//Run the Api query to fetch the weather data
	response, err := http.Get(hourlyEndPointZipCode)
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

	// If statement because the data is huge and I need to see
	// the ouputs above
	if len(weatherdata) > 0 {
		fmt.Println("yes")
	}
}
