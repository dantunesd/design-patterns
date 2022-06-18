package notifier

import (
	"encoding/base64"
)

type Encrypter struct {
	notifier Notifier
}

func NewEncrypter(notifier Notifier) *Encrypter {
	return &Encrypter{notifier: notifier}
}

func (e *Encrypter) Notify(message string) {
	e.notifier.Notify(
		base64.StdEncoding.EncodeToString([]byte(message)),
	)
}
