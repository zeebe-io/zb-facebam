package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/zeebe-io/zb-facebam/board"
	"github.com/zeebe-io/zb-facebam/thumbnail"
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
					Name:    "board",
					Aliases: []string{"b"},
					Usage:   "",
					Action: func(c *cli.Context) error {
						board.Run()
						return nil
					},
				},
				{
					Name:    "thumbnail",
					Aliases: []string{"t"},
					Usage:   "",
					Action: func(c *cli.Context) error {
                        thumbnail.Run()
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
