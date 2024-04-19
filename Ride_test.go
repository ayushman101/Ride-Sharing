package main

import (
	"testing"
)

func TestMatchedRidesHandler(t *testing.T) {

	argList := []string{
		"ADD_RIDER",
		"R1",
		"0",
		"0",
	}

	AddRiderHandler(argList)

	argList = []string{
		"ADD_DRIVER",
		"D1",
		"1",
		"1",
	}

	AddDriverHandler(argList)

	argList = []string{
		"MATCH",
		"R1",
	}

	if err := MatchRideHandler(argList); err != nil {
		t.Errorf("match rider handler error")
	}

}

func TestStartingRide(t *testing.T) {
	argList := []string{
		"ADD_RIDER",
		"R1",
		"0",
		"0",
	}

	AddRiderHandler(argList)

	argList = []string{
		"ADD_DRIVER",
		"D1",
		"1",
		"1",
	}

	AddDriverHandler(argList)

	argList = []string{
		"MATCH",
		"R1",
	}

	if err := MatchRideHandler(argList); err != nil {
		t.Errorf("match rider handler error")
	}

	argList = []string{
		"START_RIDE",
		"R1DE-001",
		"1",
		"R1",
	}

	StartRideHandler(argList)

}

func TestStoppingRide(t *testing.T) {
	argList := []string{
		"ADD_RIDER",
		"R1",
		"0",
		"0",
	}

	AddRiderHandler(argList)

	argList = []string{
		"ADD_DRIVER",
		"D1",
		"1",
		"1",
	}

	AddDriverHandler(argList)

	argList = []string{
		"MATCH",
		"R1",
	}

	if err := MatchRideHandler(argList); err != nil {
		t.Errorf("match rider handler error")
	}

	argList = []string{
		"START_RIDE",
		"R1DE-001",
		"1",
		"R1",
	}

	StartRideHandler(argList)

	argList = []string{
		"STOP_RIDE",
		"R1DE-001",
		"4",
		"5",
		"32",
	}

	StopRideHandler(argList)

}

func TestBillCalculation(t *testing.T) {
	argList := []string{
		"ADD_RIDER",
		"R1",
		"0",
		"0",
	}

	AddRiderHandler(argList)

	argList = []string{
		"ADD_DRIVER",
		"D1",
		"1",
		"1",
	}

	AddDriverHandler(argList)

	argList = []string{
		"MATCH",
		"R1",
	}

	if err := MatchRideHandler(argList); err != nil {
		t.Errorf("match rider handler error")
	}

	argList = []string{
		"START_RIDE",
		"R1DE-001",
		"1",
		"R1",
	}

	StartRideHandler(argList)

	argList = []string{
		"STOP_RIDE",
		"R1DE-001",
		"4",
		"5",
		"32",
	}

	StopRideHandler(argList)

	argList = []string{
		"BILL",
		"RIDE-001",
	}

	calcBill(argList)

}
