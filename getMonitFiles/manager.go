package getmonitfiles

import "fmt"

type Manager struct{}

func (manager Manager) InsertDriverman(drivermanName string) {
	newDriverman := Driverman{
		name: drivermanName,
	}
	driverman = newDriverman
}

func (manager Manager) AddPassenger(passengerName string, passengerCurrentLocation string, passengerDestination string) {
	newPassenger := Passenger{
		name:            passengerName,
		currentLocation: passengerCurrentLocation,
		destination:     passengerDestination,
	}
	Passengers = append(Passengers, newPassenger) // Enqueue
}

func (manager Manager) GetWorkOrder() string {
	workOrder := ""
	for len(Passengers) > 0 {
		nextPassnger := (Passengers[0]) // First element
		workOrder += fmt.Sprint(driverman.GetName(), " is picking up ", nextPassnger.GetName(), " from: ", nextPassnger.GetCurrentLocation(), " to ", nextPassnger.GetDestination(), "\n")
		Passengers = Passengers[1:] // Dequeue
	}
	return workOrder
}
