package gear

type State interface {
	Accelerate() bool
	SwitchToNeutral()
	SwitchToDrive()
	SwitchToReverse()
}
