package main

import (
	"ens-go/core"
	"ens-go/robot"
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
	ens, err := core.NewEns(fvApi)
	if err != nil {
		return err
	}
	r := robot.NewRobot(ens, fvToken)
	if err = r.Start(); err != nil {
		return err
	}
	TrapSignal(func() {
		r.Stop()
	})
	select {}
}
