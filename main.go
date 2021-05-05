package main

import (
	"fmt"
	"test/getmonitfiles"
)

func main() {
	manager := getmonitfiles.Manager{}
	fmt.Println("Hello, welcome to Get Monit")
	ScanAndInsertDrivermanData(manager)
	ScanAndInsertPassengerData(manager)

	for {
		fmt.Println("Want to add passenger? Enter 1 for true or 0 for false")
		var morePassengerToAdd int
		fmt.Scanln(&morePassengerToAdd)
		if morePassengerToAdd == 1 {
			ScanAndInsertPassengerData(manager)
		}
		if morePassengerToAdd == 0 {
			break
		}
	}
	fmt.Println(manager.GetWorkOrder())
}

func ScanAndInsertDrivermanData(manager getmonitfiles.Manager) {
	fmt.Println("Please enter the driver name: ")
	var inputDriverManName string
	fmt.Scanln(&inputDriverManName)
	manager.InsertDriverman(inputDriverManName)
}

func ScanAndInsertPassengerData(manager getmonitfiles.Manager) {
	fmt.Println("Please enter a passenger name: ")
	var inputPassengerName string
	fmt.Scanln(&inputPassengerName)
	fmt.Println("Please enter a passenger current location: ")
	var inputPassengerCurrentLocation string
	fmt.Scanln(&inputPassengerCurrentLocation)
	fmt.Println("Please enter a passenger destination: ")
	var inputPassengerDestination string
	fmt.Scanln(&inputPassengerDestination)

	manager.AddPassenger(inputPassengerName, inputPassengerCurrentLocation, inputPassengerDestination)
}
