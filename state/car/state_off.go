package car

var carOffMessage = "The car is off"

type StateOff struct {
	car *Car
}

func NewStateOff(car *Car) *StateOff {
	return &StateOff{
		car: car,
	}
}

func (s *StateOff) TurnOn() {
	println("Turning on the car")

	s.car.setNewState(NewStateOn(s.car))
}

func (s *StateOff) TurnOff() {
	println(carOffMessage)
}

func (s *StateOff) Accelerate() {
	println(carOffMessage)
}

func (s *StateOff) Brake() {
	println(carOffMessage)
}

func (s *StateOff) SwitchToNeutral() {
	println(carOffMessage)
}

func (s *StateOff) SwitchToDrive() {
	println(carOffMessage)
}

func (s *StateOff) SwitchToReverse() {
	println(carOffMessage)
}
