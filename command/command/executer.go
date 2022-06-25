package command

type Executer struct {
}

func NewExecuter() *Executer {
	return &Executer{}
}

func (e *Executer) Execute(command Command) {
	command.Execute()
}
