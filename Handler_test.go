package main

import (
	"fmt"
	"testing"
)

// func TestAddRiderHandler(t *testing.T) {

// 	testList := [][]string{
// 		{
// 			"ADD_RIDER",
// 			"R1",
// 			"0",
// 			"0",
// 		},
// 		{
// 			"ADD_RIDER",
// 		},
// 	}

// 	if err := AddRiderHandler(testList[0]); err != nil {
// 		t.Errorf("add rider error")
// 	}

// 	if err := AddRiderHandler(testList[1]); err == nil {
// 		t.Errorf("add rider should have given error but didn't")
// 	}

// 	fmt.Printf("add rider success")
// }

// func TestAddDriverHandler(t *testing.T) {

// 	testList := [][]string{
// 		{
// 			"ADD_DRIDER",
// 			"D1",
// 			"2",
// 			"2",
// 		},
// 		{
// 			"ADD_DRIVER",
// 		},
// 	}

// 	if err := AddDriverHandler(testList[0]); err != nil {
// 		t.Errorf("add driver error")
// 	}

// 	if err := AddDriverHandler(testList[1]); err == nil {
// 		t.Errorf("add driver should have given error but didn't ")
// 	}

// 	fmt.Printf("add driver success")

// }

func TestDriverMethods(t *testing.T) {
	var drivers Drivers

	if err := drivers.ADD_DRIVER("D1", 2, 2); err != nil {
		t.Errorf("driver add error")
	}

	if err := drivers.ADD_DRIVER("D2", 1, 1); err != nil {
		t.Errorf("driver add error")
	}

	if err := drivers.ADD_DRIVER("D3", 2, 2); err != nil {
		t.Errorf("driver add error")
	}

	if err := drivers.ADD_DRIVER("D4", 4, 2); err != nil {
		t.Errorf("driver add error")
	}

	if err := drivers.ADD_DRIVER("D5", 3, 2); err != nil {
		t.Errorf("driver add error")
	}

	if err := drivers.ADD_DRIVER("D6", 1, 2); err != nil {
		t.Errorf("driver add error")
	}

	matchedDrivers := drivers.NearestFive(0, 0)

	if len(matchedDrivers) == 0 {
		t.Errorf("matched drivers length should not be zero")
	}

	fmt.Println("test drivers methods success")
}

func TestRide(t *testing.T) {

	var rides Rides

	rides = append(rides, Ride{
		"RIDE-001",
		1,
		0,
		50,
		STARTED,
	})

	rides = append(rides, Ride{
		"RIDE-002",
		1,
		0,
		50,
		STOPPED,
	})

	_, ok := rides.findRidewithStatus("RIDE-001", STARTED)

	if !ok {
		t.Errorf("ride should exist")
	}

	_, ok = rides.findRidewithStatus("RIDE-002", STOPPED)

	if !ok {
		t.Errorf("stopped ride not found")
	}

	_, ok = rides.findRidewithStatus("RIDE-003", STARTED)

	if ok {
		t.Errorf("ride should not exist")
	}

	rides.Remove("RIDE-002")

	fmt.Println(rides)

	_, ok = rides.findRidewithStatus("RIDE-002", STARTED)

	if ok {
		t.Errorf("ride wasn't removed")
	}

}

func TestMatchedRides(t *testing.T) {
	var matchedRides MatchedRides

	matchedRides = append(matchedRides, MatchedRide{
		1,
		[]int32{
			2,
			3,
		},
	})

	_, ok := matchedRides.findRide(1)

	if !ok {
		t.Errorf("matched ride should exist")
	}

	_, ok = matchedRides.findRide(2)

	if ok {
		t.Errorf("matched ride should not exist")
	}
}

func TestRiders(t *testing.T) {
	var riders Riders

	riders = append(riders, Rider{
		"R1",
		2,
		2,
	})

	_, ok := riders.find("R1")

	if !ok {
		t.Errorf("rider should exist")
	}

	_, ok = riders.find("R2")

	if ok {
		t.Errorf("rider should not exist")
	}

}
func TestEntryPoint(t *testing.T) {
	cliArgs := []string{
		"input1.txt",
	}

	entrypoint(cliArgs)

	cliArgs = []string{}

	entrypoint(cliArgs)

}
