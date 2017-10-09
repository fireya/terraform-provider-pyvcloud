package vcd

import (
	"github.com/hashicorp/go-plugin"
	//	"fmt"
	//	"net/url"
)

type Config struct {
	User            string
	Password        string
	Org             string
	Href            string
	VDC             string
	MaxRetryTimeout int
	InsecureFlag    bool
}

type VCDClient struct {
	*plugin.Client
}

func (c *Config) Client() (*VCDClient, error) {

	vcdclient, err := CreatClient(c)
	return vcdclient, err
}
