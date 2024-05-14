package commands

import (
	"errors"

	"github.com/urfave/cli/v2"
)

var Create = &cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "Create project from template",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "template",
			Value:   "default",
			Usage:   "template to use to create project",
			Aliases: []string{"t"},
		},
	},
	Action: func(ctx *cli.Context) error {
		if !ctx.Args().Present() {
			return errors.New("Put the target directory")
		}

		if ctx.Args().Len() > 1 {
			return errors.New("Expected only 1 argument")
		}

		targetDirectory := ctx.Args().Get(0)
		templateName := ctx.String("template")

		// Execute action

		return nil
	},
	Args: true,
}
