package rollbar

import (
	helper "github.com/davidji99/terraform-provider-rollbar/helper/test"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccConfig *helper.TestConfig

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"rollbar": testAccProvider,
	}
	testAccConfig = helper.NewTestConfig()
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	testAccConfig.GetOrAbort(t, helper.TestConfigAccountAccessToken)
}
