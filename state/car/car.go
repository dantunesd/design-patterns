package car

import "design-patterns/state/gear"

type ICarState interface {
	TurnOn()
	TurnOff()
	Accelerate()
	Brake()
	SwitchToNeutral()
	SwitchToDrive()
	SwitchToReverse()
}

type IGear interface {
	Accelerate()
	Brake()
	SwitchToNeutral()
	SwitchToDrive()
	SwitchToReverse()
}

type Car struct {
	state ICarState
	gear  IGear
}

func NewCar() *Car {
	car := &Car{}

	car.gear = gear.NewGear()
	car.state = NewStateOff(car)

	return car
}

func (c *Car) TurnOn() {
	c.state.TurnOn()
}

func (c *Car) TurnOff() {
	c.state.TurnOff()
}

func (c *Car) Accelerate() {
	c.state.Accelerate()
}

func (c *Car) Brake() {
	c.state.Brake()
}

func (c *Car) SwitchToNeutral() {
	c.state.SwitchToNeutral()
}

func (c *Car) SwitchToDrive() {
	c.state.SwitchToDrive()
}

func (c *Car) SwitchToReverse() {
	c.state.SwitchToReverse()
}

func (c *Car) setNewState(state ICarState) {
	c.state = state
}
