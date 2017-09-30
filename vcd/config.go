package vcd

//import (
//	"fmt"
//	"net/url"
//)

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
}

func (c *Config) Client() (*VCDClient, error) {

	vcdclient := &VCDClient{}
	Login()
	return vcdclient, nil
}
