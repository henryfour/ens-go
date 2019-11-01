package main

import "github.com/urfave/cli"

var (
	// global
	fvApi   string
	flagApi = cli.StringFlag{
		Name:        "api",
		Usage:       "api url (required)",
		Required:    true,
		Destination: &fvApi,
	}

	// query
	fvFile   string
	flagFile = cli.StringFlag{
		Name:        "file,f",
		Usage:       "input file contains ens names, one name per line",
		Destination: &fvFile,
	}
	// robot
	fvToken   string
	flagToken = cli.StringFlag{
		Name:        "token",
		Usage:       "robot token",
		Required:    true,
		Destination: &fvToken,
	}
)
