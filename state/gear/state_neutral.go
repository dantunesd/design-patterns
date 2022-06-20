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

func (s *StateNeutral) SwitchToNeutral() {
	println("The gear is already Neutral")
}

func (s *StateNeutral) SwitchToDrive() {
	s.gear.setNewState(NewStateDrive(s.gear))
}

func (s *StateNeutral) SwitchToReverse() {
	s.gear.setNewState(NewStateReverse(s.gear))
}
