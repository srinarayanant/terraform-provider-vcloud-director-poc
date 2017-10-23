package provider;
import (
	//"fmt"
	"testing"

	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/helper/schema"

)
var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func TestAccResourceCatalogItemMedia(t *testing.T) {
	

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCatalogItemDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCatalogItem_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCatalogItelUpload("c1"),
					
				),
			},
		},
	})
}


func testAccCheckCatalogItelUpload(catalogName string) resource.TestCheckFunc {
	return testAccCheckCatalogItemDestroy
}
func testAccCheckCatalogItemDestroy(s *terraform.State) error {
	

	return nil


}

const testAccCatalogItem_basic = `
resource "vcloud-director_catalog_item_media" "item1" {
	item_name = "item1"
	catalog_name= "c3"
	file_path="/home/file.iso"
}
`

func testAccPreCheck(t *testing.T) {

}