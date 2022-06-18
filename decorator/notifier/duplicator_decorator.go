package notifier

type Duplicator struct {
	notifier Notifier
}

func NewDuplicator(notifier Notifier) *Duplicator {
	return &Duplicator{notifier: notifier}
}

func (N Duplicator) Notify(message string) {
	N.notifier.Notify(message + " " + message)
}
