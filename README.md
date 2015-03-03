# go-weatherunderground

========================

A set of examples use go to fetch data from weather underground.


### Examples
- hourly forecast for the next 36 hours immediately following the API request.
- Returns an hourly forecast for the next 10 days
- Example Query using US ZipCode
- Example using County/City {Japan/Tokyo}
- History_YYYYMMDD returns a summary of the observed weather for the specified date.

go-weatherunderground is also a client for the http server application server.go

Once data has been collected from go-weatherunderground.go the data is posted
to the http server (:8080), and the server.go app writes this data to mongodb.

Copyright (c) 2015 Craig Nicholson
