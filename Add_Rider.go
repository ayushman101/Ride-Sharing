package main

import (
	"errors"
	"fmt"
	"strconv"
)

func AddRiderHandler(argList []string) error {

	if len(argList) < 4 {
		fmt.Printf("arguments given %v expected 4", len(argList))
		return errors.New("wrong number of arguments")
	} else {
		xf, _ := strconv.Atoi(argList[2])
		yf, _ := strconv.Atoi(argList[3])
		riders.ADD_RIDER(argList[1], int32(xf), int32(yf))
	}

	return nil
}
