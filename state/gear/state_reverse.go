package gear

type StateReverse struct {
	gear *Gear
}

func NewStateReverse(gear *Gear) *StateReverse {
	return &StateReverse{
		gear: gear,
	}
}

func (s *StateReverse) Accelerate() {
	println("Accelerating backwards")
}

func (s *StateReverse) Brake() {
	println("Braking the car")
}
