package car

type StateOn struct {
	car *Car
}

func NewStateOn(car *Car) *StateOn {
	return &StateOn{
		car: car,
	}
}

func (s *StateOn) TurnOn() {
	println("the car is already on")
}

func (s *StateOn) TurnOff() {
	println("Turning off the car")

	s.car.setNewState(NewStateOff(s.car))
}

func (s *StateOn) Accelerate() {
	s.car.gear.Accelerate()
}

func (s *StateOn) Brake() {
	s.car.gear.Brake()
}

func (s *StateOn) SwitchToNeutral() {
	s.car.gear.SwitchToNeutral()
}

func (s *StateOn) SwitchToDrive() {
	s.car.gear.SwitchToDrive()
}

func (s *StateOn) SwitchToReverse() {
	s.car.gear.SwitchToReverse()
}
