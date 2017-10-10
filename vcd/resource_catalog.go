package vcd

import (
	"fmt"
	"log"
	"os"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/golang/glog"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCatalog() *schema.Resource {
	return &schema.Resource{
		Create: resourceCatalogCreate,
		Read:   resourceCatalogRead,
		Update: resourceCatalogUpdate,
		Delete: resourceCatalogDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"shared": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceCatalogCreate(d *schema.ResourceData, m interface{}) error {
	log.SetOutput(os.Stdout)
	/*	config := Config{
		User:     d.Get("username").(string),
		Password: d.Get("password").(string),
		Org:      d.Get("org").(string),
		//Href:            d.Get("url").(string),
		//VDC:             d.Get("vdc").(string),
		//MaxRetryTimeout: maxRetryTimeout,
		//InsecureFlag:    d.Get("allow_unverified_ssl").(bool),
	}*/
	//config.Client()
	vcdClient := m.(*VCDClient)

	//spew.Dump(vcdClient)
	//glog.Info("============ glocg ")
	provider := vcdClient.getProvider()
	//provider.Login("user1", "Admin!23", "O1", "10.112.83.27")
	res, err := provider.IsPresentCatalog("c1")
	fmt.Println(res)
	fmt.Println(err)
	//return fmt.Errorf("Unable to find edge gateway: %#v", res)
	return nil
}

func resourceCatalogRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCatalogUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCatalogDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
