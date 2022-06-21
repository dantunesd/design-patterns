package gear

import "reflect"

type Gear struct {
	state State
}

func NewGear() *Gear {
	gear := &Gear{}
	gear.state = NewStateNeutral(gear)

	return gear
}

func (g *Gear) Accelerate() {
	g.state.Accelerate()
}

func (g *Gear) SwitchToNeutral() {
	g.state.SwitchToNeutral()
}

func (g *Gear) SwitchToDrive() {
	g.state.SwitchToDrive()
}

func (g *Gear) SwitchToReverse() {
	g.state.SwitchToReverse()
}

func (g *Gear) setNewState(state State) {
	println("Switching gear: " + reflect.TypeOf(state).Elem().Name())

	g.state = state
}
