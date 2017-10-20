package provider

import (
	"fmt"
	//"github.com/hashicorp/terraform/helper/logging"
	"log"
	"os"
	//"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/logutils"
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
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceCatalogCreate(d *schema.ResourceData, m interface{}) error {
	log.SetOutput(os.Stdout)
	fmt.Printf("\n\n  === resourceCatalogCreate ===========================================\n\n")
	cname := d.Get("name").(string)
	desc := d.Get("description").(string)
	shared := d.Get("shared").(bool)

	vcdClient := m.(*VCDClient)

	provider := vcdClient.getProvider()

	resp, errp := provider.IsPresentCatalog(cname)

	if errp != nil {
		return fmt.Errorf("Error Creating catalog :%v %#v", cname, errp)
	}

	if resp.Present {
		fmt.Printf("catalog %v is present  ================ ", cname)
		d.SetId(cname)
		return nil
	}

	res, err := provider.CreateCatalog(cname, desc, shared)

	if err != nil {
		return fmt.Errorf("Error Creating catalog :%v %#v", cname, err)
	}
	if res.Created {
		fmt.Printf("catalog %v  created  ================ ", cname)
		d.SetId(cname)
	}
	return nil
}

func resourceCatalogRead(d *schema.ResourceData, m interface{}) error {

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	log.Print("[INFO] ========================readingres resource======= \n\n\n\n\n\n")

	cname := d.Get("name").(string)

	vcdClient := m.(*VCDClient)

	provider := vcdClient.getProvider()

	res, err := provider.IsPresentCatalog(cname)
	if err != nil {
		return fmt.Errorf("Error checking resourceCatalogRead  %#v", err)
	}
	if res.Present {
		fmt.Printf("catalog %v is present / setting state", cname)
		d.SetId(cname)
	} else {
		// resource catalog not present / clear id for recreate
		d.SetId("")
	}
	d.Set("name", cname)
	return nil

}

func resourceCatalogUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCatalogDelete(d *schema.ResourceData, m interface{}) error {

	log.SetOutput(os.Stdout)
	fmt.Printf("\n\n  === resourceCatalogDelete ===========================================\n\n")
	cname := d.Get("name").(string)

	vcdClient := m.(*VCDClient)

	provider := vcdClient.getProvider()

	_, err := provider.DeleteCatalog(cname)

	if err != nil {
		return fmt.Errorf("Error Creating catalog :%v %#v", cname, err)
	}

	return nil
}
