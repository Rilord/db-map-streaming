package instrumentation

import (
	"os"
	"os/signal"
	"syscall"
)

var enabled = false

func PreInit() {
	enabled = true
}

func Done() <-chan struct{} {
	if !enabled {
		return nil
	}

	done := make(chan struct{}, 1)

	sig_handler := make(chan os.Signal, 1)
	signal.Notify(sig_handler, syscall.SIGUSR2)

	go func() {
		<-sig_handler
		done <- struct{}{}
	}()

	return done
}
