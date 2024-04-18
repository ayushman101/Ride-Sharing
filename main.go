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
	Id          string
	RiderIndex  int32
	DriverIndex int32 //driver index is 0 by default but is changed when ride is started
	Bill        float32
	Status      string // Matched |  Started |   Stopped.
}

type MatchedRide struct {
	RiderIndex    int32
	MatchedRiders []int32
}

type Drivers []Driver
type Riders []Rider
type Rides []Ride
type MatchedRides []MatchedRide

//TODO: Need to String method to all the above types. Floats must be printed to 2 decimal places.

//check if ride exist or not

func (rd *Rides) findRide(Id string) (Ride, bool) {

	for _, ride := range *rd {
		if ride.Id == Id {
			return ride, true
		}
	}

	return Ride{}, false
}

// returns the true if there is a matched ride
func (mr *MatchedRides) findRide(riderIndex int32) (MatchedRide, bool) {

	for _, matchedRide := range *mr {
		if matchedRide.RiderIndex == riderIndex {
			return matchedRide, true
		}
	}
	return MatchedRide{}, false
}

// returns a slice of Drivers Id under 5km radius in ascending order.
// If distance is same then sort in lexicographically
func (d *Drivers) NearestFive(sourceX float32, sourceY float32) []int32 {
	var Ids []int32
	var distances []float32
	it := 0

	dri := *d

	for i := 0; i < len(dri); i++ {

		dist := math.Sqrt(math.Pow(float64(dri[i].x)-float64(sourceX), 2) + math.Pow(float64(dri[i].y)-float64(sourceY), 2))

		if dist <= 5 {
			if it < 5 {
				Ids = append(Ids, int32(i))
				distances = append(distances, float32(dist))
				it++
			} else {
				//If we already have 5 drivers and there is another under 5km
				//Then we check if it is in top 5 by comparing it with farthest driver among the five
				maxDist := distances[0]
				k := 0
				for j := 1; j < len(Ids); j++ {
					if distances[j] > maxDist {
						maxDist = distances[j]
						_ = k
						k = j
					}
				}

				if dist < float64(maxDist) {
					distances[k] = float32(dist)
					Ids[k] = int32(i)
				}
			}
		}

	}

	for i := 1; i < len(Ids); i++ {
		pos := i
		id := Ids[i]
		dis := distances[i]
		for j := i - 1; j >= 0; j-- {
			if distances[j] > dis {
				distances[j+1] = distances[j]
				Ids[j+1] = Ids[j]
				pos--

			} else if distances[j] == dis && strings.Compare(dri[Ids[j]].Id, dri[Ids[i]].Id) > 0 {
				distances[j+1] = distances[j]
				Ids[j+1] = Ids[j]
				pos--
			} else {
				break
			}
		}

		distances[pos] = dis
		Ids[pos] = id
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

func (r *Riders) find(id string) (int32, bool) {

	riders := *r

	for i := 0; i < len(riders); i++ {
		if riders[i].Id == id {
			return int32(i), true
		}
	}

	return -1, false
}

func main() {

	var drivers Drivers
	var riders Riders
	var rides Rides
	var matchedRides MatchedRides

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
			RiderIndex, ok := riders.find(argList[1])

			if !ok {
				fmt.Printf("rider not found.\n use ADD_RIDER <RiderID> <x_Coordinate> <y_Coordinate>\n")
				continue
			}

			//get list of all the drivers within 5km range
			matchedDrivers := drivers.NearestFive(riders[RiderIndex].x, riders[RiderIndex].y)

			if len(matchedDrivers) == 0 {
				fmt.Printf("NO_DRIVERS_AVAILABLE\n")
				continue
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

			matchedRides = append(matchedRides, matchedRide)
			_ = matchedRides
		case "START_RIDE":
			//check if we have enough args
			if len(argList) < 4 {
				fmt.Printf("arguments given %v expected 4", len(argList))
				continue
			}
			//check if the rider exist or not
			RiderIndex, ok := riders.find(argList[3])

			if !ok {
				fmt.Printf("rider not found.\n use ADD_RIDER <RiderID> <x_Coordinate> <y_Coordinate>\n")
				continue
			}

			//check if the ride is matched or not
			matchedRide, ok := matchedRides.findRide(RiderIndex)

			if !ok {
				fmt.Println("INVALID RIDE")
				continue
			}

			//check if the value of N is valid
			n, _ := strconv.Atoi(argList[2])
			if n > len(matchedRide.MatchedRiders) {
				fmt.Println("INVALID RIDE")
				continue
			}

			//check if ride already exists
			_, ok = rides.findRide(argList[1])
			if ok {
				fmt.Println("INVALID RIDE")
				continue
			}

			//create a ride with status: started.
			ride := Ride{
				Id:          argList[1],
				RiderIndex:  RiderIndex,
				DriverIndex: int32(n),
				Bill:        50,
				Status:      "Started",
			}

			rides = append(rides, ride)
			fmt.Println("RIDE_STARTED ", ride.Id)

		case "STOP_RIDE":
		case "BILL":
		default:
			fmt.Printf("invalid input command\n")
			//give help instructions.
		}
	}
}
