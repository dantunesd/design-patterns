package main

import (
	"design-patterns/state/car"
	"fmt"
)

func main() {

	myCar := car.NewCar(car.NewEngine(), car.NewGear())

	fmt.Println("----- Instanciating a car off and at neutral gear -----")

	myCar.Accelerate()        // It's off, nothing will happen
	myCar.Brake()             // It's off, nothing will happen
	myCar.SwitchTo(car.Drive) // It's off, nothing will happen
	myCar.TurnOff()           // It's off, nothing will happen

	fmt.Println("----- Let's turn on the car -----")

	myCar.TurnOn()            // Now it's on.
	myCar.Accelerate()        // But it still on neutral.
	myCar.SwitchTo(car.Drive) // Switch to drive mode.
	myCar.Accelerate()        // Let's go.
	myCar.TurnOn()            // Ops, wrong command.
	myCar.Accelerate()        // And accelerate a little bit more.
	myCar.Brake()             // It's fast, let's break.

	fmt.Println("----- Let's turn off the car -----")

	myCar.SwitchTo(car.Reverse) // Let's keep it in the garage
	myCar.Accelerate()          // Backing the car
	myCar.Brake()               // Great.
	myCar.TurnOff()             // It's enough for today. See you tomorrow.
}
