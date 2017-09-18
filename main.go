package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/zeebe-io/zb-facebam/board"
)

const version = "0.1.0"

func main() {

	app := cli.NewApp()
	app.Usage = "facebam control application"
	app.Version = version
	app.Flags = []cli.Flag{}
	app.Before = cli.BeforeFunc(func(c *cli.Context) error {
		return nil
	})

	app.Authors = []cli.Author{
		{Name: "Zeebe Team", Email: "info@zeebe.io"},
	}

	app.Commands = []cli.Command{
		{
			Name:    "service",
			Aliases: []string{"s"},
			Usage:   "start a service",
			Subcommands: []cli.Command{
				{
					Name:    "analysis",
					Aliases: []string{"a"},
					Usage:   "",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:    "board",
					Aliases: []string{"b"},
					Usage:   "",
					Action: func(c *cli.Context) error {
						board.Run()
						return nil
					},
				},
				{
					Name:    "cropper",
					Aliases: []string{"c"},
					Usage:   "",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:    "processing",
					Aliases: []string{"c"},
					Usage:   "",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:    "storage",
					Aliases: []string{"c"},
					Usage:   "",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
