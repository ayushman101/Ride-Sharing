package main

import "os"

var drivers Drivers
var riders Riders
var rides Rides
var matchedRides MatchedRides

func main() {
	cliArgs := os.Args[1:]

	entrypoint(cliArgs)
}
