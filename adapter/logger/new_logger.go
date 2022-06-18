package logger

type NewLogger interface {
	Debug(message string)
	Info(message string)
	Error(message string)
}
