package car

import "fmt"

type StateOff struct {
	car *Car
}

func NewStateOff(car *Car) *StateOff {
	return &StateOff{
		car: car,
	}
}

func (s *StateOff) TurnOn() {
	fmt.Println("Turning on the car")

	s.car.setNewState(NewStateOn(s.car))
}

func (s *StateOff) TurnOff() {
	fmt.Println("trying to turn off but the car is off")
}

func (s *StateOff) SwitchTo(gear GearState) {
	fmt.Println("trying to switch gear but the car is off")
}

func (s *StateOff) Accelerate() {
	fmt.Println("trying to accelerate but the car is off")
}

func (s *StateOff) Brake() {
	fmt.Println("trying to brake but the car is off")
}
