package gear

type StateDrive struct {
	gear *Gear
}

func NewStateDrive(gear *Gear) *StateDrive {
	return &StateDrive{
		gear: gear,
	}
}

func (s *StateDrive) Accelerate() {
	println("Accelerating foward")
}

func (s *StateDrive) Brake() {
	println("Braking the car")
}
