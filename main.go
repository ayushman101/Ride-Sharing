package main

import (
	"bufio"
	"fmt"
	"math"
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
	Id             string
	RiderIndex     int32
	DriverIndex    int32 //driver index is 0 by default but is changed when ride is started
	Bill           float32
	Status         string  // Matched |  Started |   Stopped.
	matchedDrivers []int32 //index of top 5 the drivers that have been matched
}

type Drivers []Driver
type Riders []Rider
type Rides []Ride

//TODO: Need to String method to all the above types. Floats must be printed to 2 decimal places.

// returns a slice of Drivers Id under 5km radius in ascending order.
// If distance is same then sort in lexicographically
func (d *Drivers) NearestFive(sourceX float32, sourceY float32) []int32 {
	var Ids []int32

	it := 0

	dri := *d

	for i := 0; i < len(*d); i++ {
		if it == 5 {
			break
		}

		dist := math.Sqrt(math.Pow(float64(dri[i].x)-float64(sourceX), 2) + math.Pow(float64(dri[i].y)-float64(sourceY), 2))

		if dist <= 5 {
			Ids = append(Ids, int32(i))
			it++
		}
	}
	return Ids
}

func (d *Drivers) ADD_DRIVER(id string, xCor float32, yCor float32) {
	newDriver := Driver{
		Id: id,
		x:  xCor,
		y:  yCor,
	}

	*d = append(*d, newDriver)
}

func (r *Riders) ADD_RIDER(id string, xCor float32, yCor float32) {
	newRider := Rider{
		Id: id,
		x:  xCor,
		y:  yCor,
	}

	*r = append(*r, newRider)
}

func (r *Riders) find(id string) (Rider, bool) {

	for _, rider := range *r {
		if rider.Id == id {
			return rider, true
		}
	}

	return Rider{}, false
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
				continue
			} else {
				xf, _ := strconv.ParseFloat(argList[2], 32)
				yf, _ := strconv.ParseFloat(argList[3], 32)
				drivers.ADD_DRIVER(argList[1], float32(xf), float32(yf))
			}
		case "ADD_RIDER":

			if len(argList) < 4 {
				fmt.Printf("arguments given %v expected 4", len(argList))
				continue
			} else {
				xf, _ := strconv.ParseFloat(argList[2], 32)
				yf, _ := strconv.ParseFloat(argList[3], 32)
				riders.ADD_RIDER(argList[1], float32(xf), float32(yf))
			}

			fmt.Printf("Riders: %+v\n", riders)
		case "MATCH":
			if len(argList) < 2 {
				fmt.Printf("arguments given %v expected 2", len(argList))
				continue
			}
			//rider must exist
			rider, ok := riders.find(argList[1])

			if !ok {
				fmt.Printf("rider not found.\n use ADD_RIDER <RiderID> <x_Coordinate> <y_Coordinate>\n")
				continue
			}

			//get list of all the drivers within 5km range
			matchedDrivers := drivers.NearestFive(rider.x, rider.y)

			if len(matchedDrivers) == 0 {
				fmt.Printf("NO_DRIVERS_AVAILABLE\n")
				continue
			}

			fmt.Printf("DRIVERS_MATCHED ")

			for _, driverIndex := range matchedDrivers {
				fmt.Printf("%s ", drivers[driverIndex].Id)
			}
			fmt.Printf("\n")

			//create a ride with status : Matched

		case "START_RIDE":
		case "STOP_RIDE":
		case "BILL":
		default:
			fmt.Printf("invalid input command\n")
			//give help instructions.
		}
	}
}
