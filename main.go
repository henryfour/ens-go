package main

import (
	"errors"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ApiFlag = cli.StringFlag{
		Name:     "api",
		Usage:    "api url",
		Required: true,
	}
	InputFile = cli.StringFlag{
		Name:     "file,f",
		Usage:    "input file",
	}
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		ApiFlag,
		InputFile,
	}
	app.Action = func(c *cli.Context) error {
		api := c.String(ApiFlag.Name)
		if api == "" {
			return errors.New("api url is required")
		}
		names := c.Args()
		filename := c.String("file")
		if filename != "" {
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			names = strings.Split(string(content), "\n")
		}
		return CheckNames(api, names)
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
