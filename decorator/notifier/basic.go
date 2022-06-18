package notifier

import "fmt"

type Default struct{}

func NewDefault() *Default {
	return &Default{}
}

func (n Default) Notify(message string) {
	fmt.Println(message)
}
