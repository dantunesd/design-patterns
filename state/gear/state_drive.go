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

func (s *StateDrive) SwitchToNeutral() {
	s.gear.setNewState(NewStateNeutral(s.gear))
}

func (s *StateDrive) SwitchToDrive() {
	println("The gear is already Drive")
}

func (s *StateDrive) SwitchToReverse() {
	s.gear.setNewState(NewStateReverse(s.gear))
}
