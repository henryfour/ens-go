package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "ens-go"
	app.Description = "A simple gadget of ens"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		cmdQuery,
		cmdRobot,
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
