package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Create *cli.Command = &cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage: "Create project from template",
	Action: func(ctx *cli.Context) error {
		fmt.Println("Create not implemented")
		return nil
	},
}
