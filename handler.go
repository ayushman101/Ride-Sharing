package main

import (
	"math"
	"strings"
)

type Driver struct {
	Id     string
	x      int32
	y      int32
	status string // FREE | BUSY
}

type Rider struct {
	Id string
	x  int32
	y  int32
}

type Ride struct {
	Id          string
	RiderIndex  int32
	DriverIndex int32 //driver index is 0 by default but is changed when ride is started
	Bill        float64
	Status      string // |  Started |   Stopped.
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
		if ride.Id == Id && ride.Status == "STARTED" {
			return ride, true
		}
	}

	return Ride{}, false
}

func (rd *Rides) findCompletedRide(Id string) (Ride, bool) {

	for _, ride := range *rd {
		if ride.Id == Id && ride.Status == "STOPPED" {
			return ride, true
		}
	}

	return Ride{}, false
}

func (rd *Rides) Remove(Id string) {

	oldrd := *rd
	var newrd Rides

	for _, ride := range oldrd {
		if ride.Id == Id && ride.Status == "STOPPED" {
			continue
		}
		newrd = append(newrd, ride)
	}

	*rd = newrd

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
func (d *Drivers) NearestFive(sourceX int32, sourceY int32) []int32 {
	var Ids []int32
	var distances []float32
	it := 0

	dri := *d

	for i := 0; i < len(dri); i++ {

		if dri[i].status == "BUSY" {
			continue
		}

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

			} else if distances[j] == dis && strings.Compare(dri[Ids[j]].Id, dri[Ids[i]].Id) < 0 {
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

func (d *Drivers) ADD_DRIVER(id string, xCor int32, yCor int32) error {
	newDriver := Driver{
		Id:     id,
		x:      xCor,
		y:      yCor,
		status: "FREE",
	}

	*d = append(*d, newDriver)

	return nil
}

func (r *Riders) ADD_RIDER(id string, xCor int32, yCor int32) {
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
