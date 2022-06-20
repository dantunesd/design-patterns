package car

type ICarState interface {
	TurnOn()
	TurnOff()
	SwitchTo(gear GearState)
	Accelerate()
	Brake()
}

type IGear interface {
	SwitchTo(gear GearState)
	IsNeutral() bool
	GetDirection() string
}

type Car struct {
	state ICarState
	gear  IGear
}

func NewCar() *Car {
	car := &Car{}

	car.gear = NewGear()
	car.state = NewStateOff(car)

	return car
}

func (c *Car) TurnOn() {
	c.state.TurnOn()
}

func (c *Car) TurnOff() {
	c.state.TurnOff()
}

func (c *Car) SwitchTo(gear GearState) {
	c.state.SwitchTo(gear)
}

func (c *Car) Accelerate() {
	c.state.Accelerate()
}

func (c *Car) Brake() {
	c.state.Brake()
}

func (c *Car) setNewState(state ICarState) {
	c.state = state
}
