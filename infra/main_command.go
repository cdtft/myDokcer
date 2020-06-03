package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"myDocker/infra/container"
)

var runCommand = &cli.Command{
	Name:  "run",
	Usage: `create a container with namespace and cgroups limit cdtftContainer run -ti [command]`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = &cli.Command{
	Name: "init",
	Usage: "Init container process run user's process in container",
	Action: func(context *cli.Context) error {
		logrus.Info("init come on")
		cmd := context.Args().Get(0)
		logrus.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
