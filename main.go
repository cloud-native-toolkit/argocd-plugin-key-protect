package main

import (
	"argocd-plugin-key-protect/pkg/generate_secrets_from_files"
	"fmt"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	app := &cli.App{
		Name: "generate-secret",
		Usage: "Generates secrets populated with values from Key Protect",
		Action: func(c *cli.Context) error {
			path, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			result := generate_secrets_from_files.GenerateSecretsFromFiles(path)
			fmt.Println(result)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
