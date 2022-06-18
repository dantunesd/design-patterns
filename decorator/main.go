package main

import "design-patterns/decorator/notifier"

func main() {
	defaultNotifier := notifier.NewDefault()
	defaultNotifier.Notify("basic message")

	duplicator := notifier.NewDuplicator(defaultNotifier)
	duplicator.Notify("duplicated message")

	encrypter := notifier.NewEncrypter(defaultNotifier)
	encrypter.Notify("encrypted message")

	encrypterAndDuplicator := notifier.NewEncrypter(duplicator)
	encrypterAndDuplicator.Notify("encrypted and duplicated message")
}
