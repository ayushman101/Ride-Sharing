package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	Status      string // Matched |  Started |   Stopped.
}

type Drivers []Driver
type Riders []Rider
type Rides []Ride

//TODO: Need to String method to all the above types. Floats must be printed to 2 decimal places.

// returns a slice of Drivers Id under 5km radius in ascending order.
// If distance is same then sort in lexicographically
func (d *Drivers) NearestFive(sourceX float32, sourceY float32) ([]string, error) {
	var Ids []string

	return Ids, nil
}

func (d *Drivers) ADD_DRIVER(id string, xCor float32, yCor float32) {
	newDriver := Driver{
		Id: id,
		x:  xCor,
		y:  yCor,
	}

	*d = append(*d, newDriver)
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

		args := scanner.Text()
		argList := strings.Fields(args)

		switch argList[0] {
		case "ADD_DRIVER":

			if len(argList) < 4 {
				fmt.Printf("arguments given %v expected 4", len(argList))
			} else {
				xf, _ := strconv.ParseFloat(argList[2], 32)
				yf, _ := strconv.ParseFloat(argList[3], 32)
				drivers.ADD_DRIVER(argList[1], float32(xf), float32(yf))
			}
		case "ADD_RIDER":
		case "MATCH":
		case "START_RIDE":
		case "STOP_RIDE":
		case "BILL":
		default:
			fmt.Printf("invalid input command\n")
			//give help instructions.
		}
	}
}
