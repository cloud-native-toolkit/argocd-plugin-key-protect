package main

import (
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "generate-secret"
	app.Usage = "Generates secrets populated with values from Key Protect"
	app.Version = "1.0.0"
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
    }
}
