package provider
import (
  
  "testing"

  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
)

var testProvider *schema.Provider

func init() {
  testProvider = Provider().(*schema.Provider)
}

func TestProvider(t *testing.T) {
  if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
    t.Fatalf("err: %s", err)
  }
}

func TestProvider_impl(t *testing.T) {
  var _ terraform.ResourceProvider = Provider()
}

func testPreCheck(t *testing.T) {
  // We will use this function later on to make sure our test environment is valid.
  // For example, you can make sure here that some environment variables are set.
}