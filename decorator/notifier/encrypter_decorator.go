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

func (N Encrypter) Notify(message string) {
	N.notifier.Notify(
		base64.StdEncoding.EncodeToString([]byte(message)),
	)
}
