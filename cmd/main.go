package main

import (
	"github.com/0x716/Turnipin/bootstrap"
	"log"
)

func main() {
	app := bootstrap.NewApplication()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
