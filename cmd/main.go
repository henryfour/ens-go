package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"syscall"
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

func TrapSignal(cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Printf("signal captured: %v, exiting...\n", sig)
			if cb != nil {
				cb()
			}
			os.Exit(0)
		}
	}()
}
