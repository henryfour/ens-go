package ens

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/common/log"
	"github.com/wealdtech/go-ens/v3"
	"strings"
	"time"
)

var zeroAddr = common.Address{}

type Ens struct {
	client   *ethclient.Client
	registry *ens.Registry
}

type DomainInfo struct {
	Name      string
	Owner     string
	Available bool
}

func (d DomainInfo) String() string {
	avs := "✔"
	if !d.Available {
		avs = "✘"
	}
	return fmt.Sprintf("%s %s %s", avs, d.Name, d.Owner)
}

func NewEns(api string) (*Ens, error) {
	client, err := ethclient.Dial(api)
	if err != nil {
		return nil, err
	}
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
func (x *Ens) GetDomainInfo(name string) (*DomainInfo, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	if len(name) > 4 && name[len(name)-4:] == ".eth" {
		name = name[:len(name)-4]
	}
	if name == "" {
		return nil, nil
	}
	// todo: 这个得到的只是 CONTROLLER, 并不是 REGISTRANT; 比如: https://app.ens.domains/name/bigger.eth
	owner, err := x.registry.Owner(name + ".eth")
	if err != nil {
		return nil, err
	}
	d := &DomainInfo{
		Name:      name,
		Owner:     owner.Hex(),
		Available: owner == zeroAddr,
	}
	return d, nil
}

func (x *Ens) GetDomainInfos(api string, names []string) ([]*DomainInfo, error) {
	dis := make([]*DomainInfo, 0, len(names))
	for _, name := range names {
		di, err := x.GetDomainInfo(name)
		if err != nil {
			log.Info("failed to check %s: err=%s\n", name, err)
		} else {
			dis = append(dis, di)
		}
		time.Sleep(20 * time.Millisecond)
	}
	return dis, nil
}
