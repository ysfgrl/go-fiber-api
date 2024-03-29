package main

import (
	"github.com/urfave/cli"
	"github.com/ysfgrl/go-fiber-api/src/cmd"
	"github.com/ysfgrl/go-fiber-api/src/worker"
	"os"
)

func main() {

	cliApp := cli.NewApp()
	cliApp.Name = "ysfgrl"
	cliApp.Usage = "go fiber api"
	cliApp.Version = "0.0.0"

	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "run workers",
			Subcommands: []cli.Command{
				{
					Name:  "dev",
					Usage: "one of [dev, prod, test]",
					Action: func(c *cli.Context) error {
						if err := worker.Launch(); err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
			},
		},
		{
			Name:  "send",
			Usage: "test workers ",
			Subcommands: []cli.Command{
				{
					Name:  "dev",
					Usage: "one of [dev, prod, test]",
					Action: func(c *cli.Context) error {
						if err := worker.Test(); err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
			},
		},
		{
			Name:  "app",
			Usage: "run fiber app",
			Subcommands: []cli.Command{
				{
					Name:  "dev",
					Usage: "one of [dev, prod, test]",
					Action: func(c *cli.Context) error {
						if err := cmd.Listen(); err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
			},
		},
	}
	_ = cliApp.Run(os.Args)
}
