package main

import "design-patterns/strategy/transport"

func main() {
	transportManager := transport.NewManager()

	transportManager.Add(transport.NewPlaneStrategy())
	transportManager.Add(transport.NewBusStrategy())
	transportManager.Add(transport.NewTaxiStrategy())

	transportManager.Transport("myself", transport.PLANE)
	transportManager.Transport("myself", transport.BUS)
	transportManager.Transport("myself", transport.TAXI)
}
