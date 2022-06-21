package gear

type StateNeutral struct {
	gear *Gear
}

func NewStateNeutral(gear *Gear) *StateNeutral {
	return &StateNeutral{
		gear: gear,
	}
}

func (s *StateNeutral) Accelerate() {
	println("The gear is in neutral. Change it.")
}

func (s *StateNeutral) SwitchToNeutral() {
	println("The gear is already neutral")
}

func (s *StateNeutral) SwitchToDrive() {
	s.gear.setNewState(NewStateDrive(s.gear))
}

func (s *StateNeutral) SwitchToReverse() {
	s.gear.setNewState(NewStateReverse(s.gear))
}
