package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "rcost"
	app.Usage = "Just try"
	app.Author = "moritz31"
	app.Version = "0.0.1"
}

func main() {
	info()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
