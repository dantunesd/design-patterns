package transport

import "fmt"

const PLANE transportType = "plane"

type PlaneStrategy struct {
	transportType transportType
}

func NewPlaneStrategy() *PlaneStrategy {
	return &PlaneStrategy{
		transportType: PLANE,
	}
}

func (p *PlaneStrategy) Transport(person string) {
	fmt.Printf("transporting %s by plane\n", person)
}

func (p *PlaneStrategy) GetType() transportType {
	return p.transportType
}
