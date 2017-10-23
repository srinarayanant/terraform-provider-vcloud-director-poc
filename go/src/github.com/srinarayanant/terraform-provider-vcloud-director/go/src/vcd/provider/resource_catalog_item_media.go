package provider

import (
	"fmt"
	//"github.com/hashicorp/terraform/helper/logging"
	"log"
	"os"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/hashicorp/logutils"
	//"github.com/golang/glog"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/proto"
)

func resourceCatalogItemMedia() *schema.Resource {
	return &schema.Resource{
		Create: resourceCatalogItemMediaCreate,
		Read:   resourceCatalogItemMediaRead,
		Update: resourceCatalogItemMediaUpdate,
		Delete: resourceCatalogItemMediaDelete,

		Schema: map[string]*schema.Schema{
			"catalog_name": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: false,
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			
		},
	}
}

func resourceCatalogItemMediaCreate(d *schema.ResourceData, m interface{}) error {
	log.SetOutput(os.Stdout)
	fmt.Printf("\n\n  === resourceCatalogItemMediaCreate ===========================================\n\n")
	itemName := d.Get("item_name").(string)
	filePath := d.Get("file_path").(string)
	catalogName := d.Get("catalog_name").(string)

	vcdClient := m.(*VCDClient)

	provider := vcdClient.getProvider()

	resp, errp := provider.CatalogUploadMedia(proto.CatalogUploadMediaInfo{CatalogName:catalogName, 
														FilePath : filePath ,
														ItemName : itemName, })

	if errp != nil {
		return fmt.Errorf("Error Creating catalog Item:%v %#v", itemName, errp)
	}

	if resp.Created {
		fmt.Printf("catalog %v is Created  ================ ", catalogName)
		d.SetId(catalogName)
		return nil
	}

	
	return nil
}

func resourceCatalogItemMediaRead(d *schema.ResourceData, m interface{}) error {
	/*
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
	*/
	return nil

}

func resourceCatalogItemMediaUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCatalogItemMediaDelete(d *schema.ResourceData, m interface{}) error {
	/*
	log.SetOutput(os.Stdout)
	fmt.Printf("\n\n  === resourceCatalogDelete ===========================================\n\n")
	cname := d.Get("name").(string)

	vcdClient := m.(*VCDClient)

	provider := vcdClient.getProvider()

	_, err := provider.DeleteCatalog(cname)

	if err != nil {
		return fmt.Errorf("Error Creating catalog :%v %#v", cname, err)
	}
	*/

	return nil
}
