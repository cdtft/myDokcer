package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `create a container with namespace and cgroups limit cdtftContainer run -ti [command]`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:        "ti",
			Usage:       "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run()
		return nil
	},
}