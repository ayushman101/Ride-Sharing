package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func entrypoint(cliArgs []string) {

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

		switchCommand(argList)
	}
}

func switchCommand(argList []string) {

	switch argList[0] {

	//HANDLING ADD_DRIVER COMMAND
	case "ADD_DRIVER":

		AddDriverHandler(argList)

		//HANDLING ADD_RIDER COMMAND
	case "ADD_RIDER":

		AddRiderHandler(argList)

		//MATCHING DRIVERS WITH RIDER
	case "MATCH":

		MatchRideHandler(argList)

		//STARTING A RIDE
	case "START_RIDE":
		StartRideHandler(argList)

		//STOPPING THE RIDE
	case "STOP_RIDE":
		StopRideHandler(argList)

		//BILL
	case "BILL":
		calcBill(argList)
	default:
		fmt.Printf("invalid input command\n")
	}

}
