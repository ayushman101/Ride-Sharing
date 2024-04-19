package main

import (
	"strconv"
)

func AddDriverHandler(argList []string) error {

	xf, _ := strconv.Atoi(argList[2])
	yf, _ := strconv.Atoi(argList[3])
	drivers.ADD_DRIVER(argList[1], int32(xf), int32(yf))

	return nil
}
