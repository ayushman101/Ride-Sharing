package main

import (
	"bufio"
	"fmt"
	"os"
)

type Driver struct {
	Id string
	x  float32
	y  float32
}

type Rider struct {
	Id string
	x  float32
	y  float32
}

type Ride struct {
	Id          string
	RiderIndex  int32
	DriverIndex int32
	Bill        float32
}

type Drivers []Driver
type Riders []Rider
type Rides []Ride

// returns a slice of Drivers Id under 5km radius in ascending order.
// If distance is same then sort in lexicographically
func (d *Drivers) NearestFive(sourceX float32, sourceY float32) ([]string, error) {
	var Ids []string

	return Ids, nil
}

func main() {

	var drivers Drivers
	var riders Riders
	var rides Rides

	//Just to remove erros whild dev
	_ = drivers
	_ = riders
	_ = rides

	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		/*
			args := scanner.Text()
			argList := strings.Fields(args)

			Add your code here to process the input commands
		*/

	}
}
