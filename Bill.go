package main

import "fmt"

func calcBill(argList []string) {
	if len(argList) < 2 {
		fmt.Printf("arguments given %v expected 2", len(argList))
		return
	}

	//find the ride
	ride, ok := rides.findCompletedRide(argList[1])
	if !ok {
		fmt.Println("INVALID_RIDE")
		return
	}

	//check if its completed or not
	if ride.Status != "STOPPED" {
		fmt.Printf("RIDE_NOT_COMPLETED\n")
		return
	}

	fmt.Printf("BILL %s %s %0.2f\n", ride.Id, drivers[ride.DriverIndex].Id, ride.Bill)

	rides.Remove(ride.Id)

}
