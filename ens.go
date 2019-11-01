package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
	"strings"
	"time"
)

var zeroAddr = common.Address{}

type Ens struct {
	client   *ethclient.Client
	registry *ens.Registry
}

func NewEns(api string) (*Ens, error) {
	client, err := ethclient.Dial(api)
	if err != nil {
		return nil, err
	}
	// netId, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("netId=%v\n", netId.String())
	registry, err := ens.NewRegistry(client)
	if err != nil {
		return nil, err
	}
	return &Ens{
		client:   client,
		registry: registry,
	}, nil
}

// 查询域名是否被注册
func (x *Ens) checkName(name string) error {
	name = strings.ToLower(strings.TrimSpace(name))
	if len(name) > 4 && name[len(name)-4:] == ".eth" {
		name = name[:len(name)-4]
	}
	if name == "" {
		return nil
	}
	// todo: 这个得到的只是 CONTROLLER, 并不是 REGISTRANT; 比如: https://app.ens.domains/name/bigger.eth
	owner, err := x.registry.Owner(name + ".eth")
	if err != nil {
		return err
	}
	if owner == zeroAddr {
		fmt.Printf("owner: %s ✔ %s\n", owner.Hex(), name)
	} else {
		fmt.Printf("owner: %s ✘ %s\n", owner.Hex(), name)
	}
	return nil
}

func CheckNames(api string, names []string) error {
	x, err := NewEns(api)
	if err != nil {
		return err
	}
	for _, name := range names {
		if err = x.checkName(name); err != nil {
			fmt.Printf("failed to check %s: err=%s\n", name, err.Error())
		}
		time.Sleep(20 * time.Millisecond)
	}
	return nil
}
