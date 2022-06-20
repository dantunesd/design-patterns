package car

type GearState string

const (
	Neutral GearState = "neutral"
	Reverse GearState = "reverse"
	Drive   GearState = "drive"
)

type Gear struct {
	state GearState
}

func NewGear() *Gear {
	return &Gear{
		state: Neutral,
	}
}

func (g *Gear) SwitchTo(gear GearState) {
	g.state = gear
}

func (g *Gear) IsNeutral() bool {
	return g.state == Neutral
}

func (g *Gear) GetDirection() string {
	if g.state == Drive {
		return "foward"
	}

	if g.state == Reverse {
		return "backwards"
	}

	return ""
}
