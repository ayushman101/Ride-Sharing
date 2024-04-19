package main

import (
	"fmt"
	"math"
	"strconv"
)

func StopRideHandler(argList []string) {

	//check if we have enough args
	if len(argList) < 5 {
		fmt.Printf("arguments given %v expected 5", len(argList))
		return
	}

	//check if the ride exists
	ride, ok := rides.findRide(argList[1])

	if !ok {
		fmt.Println("INVALID_RIDE")
		return
	}

	//calculate the distance travalled and bill
	desX, _ := strconv.Atoi(argList[2])
	desY, _ := strconv.Atoi(argList[3])

	srcX := riders[ride.RiderIndex].x
	srcY := riders[ride.RiderIndex].y

	// dist := math.Sqrt(math.Pow(float64(desX)-float64(srcX), 2) + math.Pow(float64(desY)-float64(srcY), 2))

	delX := int32(desX) - srcX

	delY := int32(desY) - srcY

	delX *= delX
	delY *= delY

	dist := math.Sqrt(float64(delX) + float64(delY))

	// fmt.Println("Distance: ", dist)

	ride.Bill += float64(dist * 6.5)

	//get the minutes and add to bill
	min, _ := strconv.Atoi(argList[4])
	ride.Bill += (float64(min) * 2)

	//adding tax to the bill :20%
	ride.Bill *= 1.2

	// ride.Bill = math.Round(ride.Bill*100) / 100
	//replace the ride in actual slice of rides
	for i := 0; i < len(rides); i++ {
		if rides[i].Id == ride.Id {
			rides[i].Bill = ride.Bill
			rides[i].Status = "STOPPED"
			break
		}
	}

	fmt.Println("RIDE_STOPPED ", ride.Id)
	// fmt.Println(rides)

}
