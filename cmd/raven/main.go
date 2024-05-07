package main

import (
	"fmt"
	"log"
	"os"

	"github.com/4strodev/raven/cmd/raven/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "raven",
		Usage: "Create projects from templates and apply schematics",
		Action: func(*cli.Context) error {
			fmt.Println("Root command")
			return nil
		},
		Commands: []*cli.Command{
			commands.Create,
		},
	}

	// Attach commands
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
