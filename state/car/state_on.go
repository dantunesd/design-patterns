package car

import "fmt"

type StateOn struct {
	car *Car
}

func NewStateOn(car *Car) *StateOn {
	return &StateOn{
		car: car,
	}
}

func (s *StateOn) TurnOn() {
	fmt.Println("the car is already on")
}

func (s *StateOn) TurnOff() {
	fmt.Println("Turning off the car")

	s.car.setNewState(NewStateOff(s.car))
}

func (s *StateOn) SwitchTo(gear GearState) {
	fmt.Println("Switching gear to: " + gear)

	s.car.gear.SwitchTo(gear)
}

func (s *StateOn) Accelerate() {
	if s.car.gear.IsNeutral() {
		fmt.Println("The car is on neutral gear. Change it.")
		return
	}

	fmt.Println("Accelerating the car " + s.car.gear.GetDirection())
}

func (s *StateOn) Brake() {
	fmt.Println("braking the car")
}
