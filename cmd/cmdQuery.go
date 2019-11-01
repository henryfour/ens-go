package main

import (
	"ens-go/core"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"strings"
	"time"
)

var cmdQuery = cli.Command{
	Name:      "query",
	ShortName: "q",
	Usage:     "query infomation of ens domains",
	Flags:     []cli.Flag{flagApi, flagFile},
	Action:    actionQuery,
}

func actionQuery(c *cli.Context) error {
	names := c.Args()
	filename := c.String("file")
	if filename != "" {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		names = strings.Fields(string(content))
	}
	return queryDomainInfos(fvApi, names)
}

func queryDomainInfos(api string, names []string) error {
	ens, err := core.NewEns(api)
	if err != nil {
		return err
	}
	for _, name := range names {
		di, err := ens.GetDomainInfo(name)
		if err != nil {
			fmt.Printf("failed to check %s: err=%s\n", name, err)
		} else {
			println(di.String())
		}
		time.Sleep(20 * time.Millisecond)
	}
	return nil
}
