package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"myDocker/container"
	"myDocker/subsystems"
	"strings"
)

var runCommand = &cli.Command{
	Name:  "run",
	Usage: `create a container with namespace and cgroups limit cdtftContainer run -ti [command]`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		&cli.StringFlag{
			Name:  "m",
			Usage: "limit memory",
		},
	},
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Slice()
		tty := context.Bool("ti")
		memoryLimit := context.String("m")
		resourceConfig := &subsystems.ResourceConfig{
			MemoryLimit: memoryLimit,
		}
		logrus.Infof("resourceConfig:[%v]", resourceConfig)
		Run(tty, parseArgs(cmd), resourceConfig)
		return nil
	},
}

func parseArgs(args []string) string {
	return strings.Join(args, " ")
}

var initCommand = &cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container",
	Action: func(context *cli.Context) error {
		logrus.Info("init come on")
		//cmd := context.Args().Get(0)
		//命令不再通过参数传递，而是通过管道通信。
		err := container.RunContainerInitProcess()
		return err
	},
}
