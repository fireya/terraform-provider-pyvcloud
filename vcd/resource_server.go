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

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"catalog": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
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
	provider.Login("user1", "Admin!23", "O1", "10.112.83.27")
	res, err := vcdClient.getProvider().IsPresentCatalog("c1")
	fmt.Println(res)
	fmt.Println(err)
	return fmt.Errorf("Unable to find edge gateway: %#v", res)
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
