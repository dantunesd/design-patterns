package notifier

import "fmt"

type BasicNotifier struct{}

func NewBasicNotifier() *BasicNotifier {
	return &BasicNotifier{}
}

func (n BasicNotifier) Notify(message string) {
	fmt.Println(message)
}
