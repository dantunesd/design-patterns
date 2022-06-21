package gear

type State interface {
	Accelerate()
	SwitchToNeutral()
	SwitchToDrive()
	SwitchToReverse()
}
