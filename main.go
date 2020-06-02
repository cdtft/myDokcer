package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

const usage = `cdtftContainer is a simple container runtime implementation.
				The purpose of this project is to learn how docker works and how to wirte a docker by ourselves`

func main() {
	app := cli.NewApp()
	app.Name = "cdtftContainer"
	app.Usage = usage
	app.Commands = []*cli.Command{
		initCommand,
		runCommand,
	}

	//设置日志打印格式
	app.Before = func(context *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
