package notifier

type Duplicator struct {
	notifier Notifier
}

func NewDuplicator(notifier Notifier) *Duplicator {
	return &Duplicator{notifier: notifier}
}

func (d *Duplicator) Notify(message string) {
	d.notifier.Notify(message + " " + message)
}
