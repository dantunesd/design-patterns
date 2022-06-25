package main

import (
	"design-patterns/command/command"
)

func main() {
	request := command.NewCustomer("whatever")

	deactivateCommand := command.NewDeactivateCustomer(request)
	activateCommand := command.NewActivateCustomer(request)

	commandExecuter := command.NewExecuter()
	commandExecuter.Execute(deactivateCommand)
	commandExecuter.Execute(activateCommand)
}
