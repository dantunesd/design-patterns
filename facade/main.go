package main

func main() {
	orderSender := NewOrderSenderFacade()
	orderSender.Send("my package example")
}
