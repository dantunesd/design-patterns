package gear

type State string

type IGearState interface {
	Accelerate()
	Brake()
}

type Gear struct {
	state IGearState
}

func NewGear() *Gear {
	gear := &Gear{}
	gear.state = NewStateNeutral(gear)

	return gear
}

func (g *Gear) Accelerate() {
	g.state.Accelerate()

}

func (g *Gear) Brake() {
	g.state.Brake()
}

func (g *Gear) SwitchToNeutral() {
	println("Switching to neutral")

	g.state = NewStateNeutral(g)
}

func (g *Gear) SwitchToDrive() {
	println("Switching to drive")

	g.state = NewStateDrive(g)
}

func (g *Gear) SwitchToReverse() {
	println("Switching to reverse")

	g.state = NewStateReverse(g)
}
