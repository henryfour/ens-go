package main

import (
	"github.com/urfave/cli"
)

var cmdRobot = cli.Command{
	Name:      "robot",
	ShortName: "r",
	Usage:     "run ens-go robot",
	Flags:     []cli.Flag{flagApi, flagToken},
	Action:    actionRobot,
}

func actionRobot(c *cli.Context) error {
	return nil
}
