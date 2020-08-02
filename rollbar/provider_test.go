package rollbar

import (
	helper "github.com/davidji99/terraform-provider-rollbar/helper/test"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testAccConfig *helper.TestConfig

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"rollbar": testAccProvider,
	}
	testAccConfig = helper.NewTestConfig()
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	testAccConfig.GetOrAbort(t, helper.TestConfigAccountAccessToken)
}
