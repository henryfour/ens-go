package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ApiFlag = cli.StringFlag{
		Name:     "api",
		Usage:    "api url (required)",
		Required: true,
	}
	InputFile = cli.StringFlag{
		Name:     "file,f",
		Usage:    "input file contains ens names, one name per line",
	}
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "ens-go"
	app.Description = "A simple gadget of ens"
	app.Flags = []cli.Flag{
		ApiFlag,
		InputFile,
	}
	app.Action = func(c *cli.Context) error {
		api := c.String(ApiFlag.Name)
		names := c.Args()
		filename := c.String("file")
		if filename != "" {
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			names = strings.Fields(string(content))
		}
		return CheckNames(api, names)
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
