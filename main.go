package main

import (
	"github.com/urfave/cli/v2"
)

const usage = `cdtftContainer is a simple container runtime implementation.
				The purpose of this project is to learn how docker works and how to wirte a docker by ourselves`

func main() {
	app := cli.NewApp()
	app.Name = "cdtftContainer"
	app.Usage = usage
	app.Commands = []cli.Commands{

	}
}

