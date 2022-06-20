package car

type ICarState interface {
	TurnOn()
	TurnOff()
	SwitchTo(gear GearState)
	Accelerate()
	Brake()
}

type Car struct {
	state  ICarState
	engine IEngine
	gear   IGear
}

func NewCar(engine IEngine, gear IGear) *Car {
	car := &Car{}

	car.engine = engine
	car.gear = gear
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

func (c *Car) canAccelerate() bool {
	return !c.gear.IsNeutral()
}

func (c *Car) getDirection() string {
	return c.gear.GetDirection()
}
