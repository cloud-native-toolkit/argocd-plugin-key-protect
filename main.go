package main

import (
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/generate_secrets_from_files"
	"fmt"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version   string
)

var app = cli.NewApp()

func main() {
	app := &cli.App{
		Name: "generate-secret",
		Usage: "Generates secrets populated with values from Key Protect",
		Version: version,
		ArgsUsage: "[path]",
		Action: func(c *cli.Context) error {
			path, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			if c.Args().Len() > 0 {
				path = c.Args().Get(0)
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
