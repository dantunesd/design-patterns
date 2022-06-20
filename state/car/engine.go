package car

type IEngine interface {
	TurnOn()
	TurnOff()
	IsOn() bool
}

type Engine struct {
	state bool
}

func NewEngine() *Engine {
	return &Engine{
		state: false,
	}
}

func (e *Engine) TurnOn() {
	e.state = true
}

func (e *Engine) TurnOff() {
	e.state = false
}

func (e *Engine) IsOn() bool {
	return e.state
}
