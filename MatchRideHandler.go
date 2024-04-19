package main

import (
	"errors"
	"fmt"
)

func MatchRideHandler(argList []string) error {

	//rider must exist
	RiderIndex, ok := riders.find(argList[1])

	if !ok {
		fmt.Printf("rider not found.\n use ADD_RIDER <RiderID> <x_Coordinate> <y_Coordinate>\n")
		return errors.New("rider not found")
	}

	//get list of all the drivers within 5km range
	matchedDrivers := drivers.NearestFive(riders[RiderIndex].x, riders[RiderIndex].y)

	if len(matchedDrivers) == 0 {
		fmt.Printf("NO_DRIVERS_AVAILABLE\n")
		return errors.New("no driver is available")
	}

	fmt.Printf("DRIVERS_MATCHED ")

	for _, driverIndex := range matchedDrivers {
		fmt.Printf("%s ", drivers[driverIndex].Id)
	}
	fmt.Printf("\n")

	//create a MatchedRide
	matchedRide := MatchedRide{
		RiderIndex:    RiderIndex,
		MatchedRiders: matchedDrivers,
	}

	//TODO:check if matched ride already exists
	_, ok = matchedRides.findRide(RiderIndex)

	//if rider hasn't been matched already then append new match
	//else modify the existing matched ride.
	if !ok {

		matchedRides = append(matchedRides, matchedRide)
		_ = matchedRides

		return nil

	}

	for i := 0; i < len(matchedRides); i++ {
		if matchedRides[i].RiderIndex == RiderIndex {
			matchedRides[i] = matchedRide
		}
	}

	return nil

}
