package main

import "fmt"

//import "os"

import (
	"./vcd"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"pyvcloudvcd_server": resourceServer(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	fmt.Printf(" in Provider go lang")
	config := vcd.Config{
		User:     d.Get("username").(string),
		Password: d.Get("password").(string),
		Org:      d.Get("org").(string),
		//Href:            d.Get("url").(string),
		//VDC:             d.Get("vdc").(string),
		//MaxRetryTimeout: maxRetryTimeout,
		//InsecureFlag:    d.Get("allow_unverified_ssl").(bool),
	}
	return config.Client()
}
