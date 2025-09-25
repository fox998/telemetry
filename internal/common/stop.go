package common

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func GetStop() chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Application is running. Press Ctrl+C to stop.")

	return sigChan
}
