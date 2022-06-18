package main

import (
	"design-patterns/adapter/logger"
	"log"
)

func main() {
	standardLogger := log.Default()
	CoupledWithStandardLogger(standardLogger)

	loggerAdapter := logger.NewAdapter(standardLogger)
	UncoupledWithStandarLogger(loggerAdapter)
}

func CoupledWithStandardLogger(logger *log.Logger) {
	logger.Println("debug message")
	logger.Println("info message")
	logger.Println("error message")
}

func UncoupledWithStandarLogger(logger logger.NewLogger) {
	logger.Debug("message")
	logger.Info("message")
	logger.Error("message")
}
