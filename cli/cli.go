package cli

import (
	"sort"

	"github.com/urfave/cli"

	"github.com/Place1/swagger-version-manager/commands"
)

func Run(args []string) error {
	app := cli.NewApp()
	app.Name = "swagger-version-manager"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "current",
			Usage: "show the current swagger codegen version",
			Action: func(context *cli.Context) error {
				return commands.Current()
			},
		},
		{
			Name:  "list",
			Usage: "list available swagger codegen versions",
			Action: func(context *cli.Context) error {
				return commands.List()
			},
		},
		{
			Name:      "use",
			Usage:     "use the specified swagger-codegen-version",
			ArgsUsage: "<version>",
			Action: func(context *cli.Context) error {
				version := context.Args().First()
				if version == "" {
					return cli.NewExitError("please provide version string", 1)
				}
				return commands.Use(context.Args().First())
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	return app.Run(args)
}
