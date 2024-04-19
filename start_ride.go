package main

import (
	"errors"
	"fmt"
	"strconv"
)

func StartRideHandler(argList []string) error {

	//check if the rider exist or not
	RiderIndex, ok := riders.find(argList[3])

	if !ok {
		fmt.Printf("rider not found.\n use ADD_RIDER <RiderID> <x_Coordinate> <y_Coordinate>\n")
		return errors.New("rider not found")

	}

	//check if the ride is matched or not
	matchedRide, ok := matchedRides.findRide(RiderIndex)

	if !ok {
		fmt.Println("INVALID_RIDE")
		return errors.New("ride not matched")

	}

	//check if the value of N is valid
	n, _ := strconv.Atoi(argList[2])
	if n > len(matchedRide.MatchedRiders) {
		fmt.Println("INVALID_RIDE")
		return errors.New("driver index out of range")

	}

	//check if ride already exists
	_, ok = rides.findRidewithStatus(argList[1], STARTED)
	if ok {
		fmt.Println("INVALID_RIDE")
		return errors.New("ride already exists")
	}

	//create a ride with status: started.
	ride := Ride{
		Id:          argList[1],
		RiderIndex:  RiderIndex,
		DriverIndex: matchedRide.MatchedRiders[(n)-1],
		Bill:        float64(BASE_FARE),
		Status:      STARTED,
	}

	//Make rider busy
	drivers[matchedRide.MatchedRiders[int32(n)-1]].status = BUSY

	rides = append(rides, ride)
	fmt.Println("RIDE_STARTED ", ride.Id)
	// fmt.Println(rides)

	return nil
}
