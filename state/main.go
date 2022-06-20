package main

import (
	"design-patterns/state/car"
)

func main() {

	myCar := car.NewCar()

	println("----- Instanciating a car off and at neutral gear -----")

	myCar.Accelerate()    // It's off, nothing will happen
	myCar.Brake()         // It's off, nothing will happen
	myCar.SwitchToDrive() // It's off, nothing will happen
	myCar.TurnOff()       // It's off, nothing will happen

	println("----- Let's turn on the car -----")

	myCar.TurnOn()        // Now it's on.
	myCar.Accelerate()    // But it still on neutral.
	myCar.SwitchToDrive() // Switch to drive mode.
	myCar.Accelerate()    // Let's go.
	myCar.TurnOn()        // Ops, wrong command.
	myCar.Accelerate()    // Let's accelerate a little bit more.
	myCar.Brake()         // It's fast, let's brake it or we'll break.

	println("----- Let's turn off the car -----")

	myCar.SwitchToReverse() // Let's keep it in the garage
	myCar.SwitchToReverse() // Ops, I have already switched, I didn't see it
	myCar.Accelerate()      // Backing up the car
	myCar.Brake()           // Great.
	myCar.TurnOff()         // It's enough for today. See you tomorrow.
}
