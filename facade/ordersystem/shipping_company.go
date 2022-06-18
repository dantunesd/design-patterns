package ordersystem

import "fmt"

type ShippingCompany struct{}

func NewShippingCompany() ShippingCompany {
	return ShippingCompany{}
}

func (s *ShippingCompany) Transport(postage *Postage) {
	fmt.Printf("postageID: %s - transported\n", postage.ID)
}
