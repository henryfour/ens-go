package main

import (
	"errors"
	"github.com/urfave/cli"
	"os"
)

var (
	ApiFlag = cli.StringFlag{
		Name:     "api",
		Usage:    "api url",
		Required: true,
	}
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		ApiFlag,
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 0 {
			return errors.New("please input at least on name")
		}
		api := c.String(ApiFlag.Name)
		if api == "" {
			return errors.New("api url is required")
		}
		return CheckNames(api, c.Args())
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
