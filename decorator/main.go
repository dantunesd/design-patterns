package main

import "design-patterns/decorator/notifier"

func main() {
	basicNotifier := notifier.NewBasicNotifier()
	basicNotifier.Notify("basic message")

	duplicator := notifier.NewDuplicator(basicNotifier)
	duplicator.Notify("duplicated message")

	encrypter := notifier.NewEncrypter(basicNotifier)
	encrypter.Notify("encrypted message")

	encrypterAndDuplicator := notifier.NewEncrypter(duplicator)
	encrypterAndDuplicator.Notify("encrypted and duplicated message")
}
