package main

import (
	maphub "game-server/internal"
	"game-server/pkg/instrumentation"
	"log"
)

func main() {
	app, err := maphub.New()

	if err != nil {
		log.Fatal(err)
	}

	errs := make(chan error, 1)
	go func() { errs <- app.Run() }()

	select {
	case err := <-errs:
		log.Fatal(err)
	case <-instrumentation.Done():
	}
}
