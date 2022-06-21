package car

type State interface {
	TurnOn()
	TurnOff()
	Brake()
	Accelerate()
	SwitchToNeutral()
	SwitchToDrive()
	SwitchToReverse()
}
