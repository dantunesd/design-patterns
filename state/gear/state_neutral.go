package gear

var onNeutralGearMessage = "The car is on neutral gear. Change it."

type StateNeutral struct {
	gear *Gear
}

func NewStateNeutral(gear *Gear) *StateNeutral {
	return &StateNeutral{
		gear: gear,
	}
}

func (s *StateNeutral) Accelerate() {
	println(onNeutralGearMessage)
}

func (s *StateNeutral) Brake() {
	println(onNeutralGearMessage)
}
