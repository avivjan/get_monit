package getmonitfiles

type Passenger struct {
	name            string
	currentLocation string
	destination     string
}

func (passenger Passenger) GetName() string {
	return passenger.name
}
func (passenger Passenger) GetCurrentLocation() string {
	return passenger.currentLocation
}
func (passenger Passenger) GetDestination() string {
	return passenger.destination
}
