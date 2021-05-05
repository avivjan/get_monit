package getmonitfiles

type Driverman struct {
	name string
}

func (driverman Driverman) GetName() string {
	return driverman.name
}
