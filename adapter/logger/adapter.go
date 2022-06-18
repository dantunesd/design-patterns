package logger

import "log"

type Adapter struct {
	logger *log.Logger
}

func NewAdapter(logger *log.Logger) *Adapter {
	return &Adapter{
		logger: logger,
	}
}

func (l *Adapter) Debug(message string) {
	l.logger.Printf("DEBUG: %s", message)
}

func (l *Adapter) Info(message string) {
	l.logger.Printf("INFO: %s", message)
}

func (l *Adapter) Error(message string) {
	l.logger.Printf("ERROR: %s", message)
}
