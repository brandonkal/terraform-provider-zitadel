package default_oidc_settings_test

import (
	"fmt"
	"testing"

	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/admin"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
)

func TestAccDefaultOIDCSettings(t *testing.T) {
	resourceName := "zitadel_default_oidc_settings"
	initialProperty := "123h0m0s"
	updatedProperty := "456h0m0s"
	frame, err := test_utils.NewInstanceTestFrame(resourceName)
	if err != nil {
		t.Fatalf("setting up test context failed: %v", err)
	}
	test_utils.RunLifecyleTest(
		t,
		frame.BaseTestFrame,
		func(accessTokenLifetime, _ interface{}) string {
			return fmt.Sprintf(`
resource "%s" "%s" {
	access_token_lifetime = "%s"
  	id_token_lifetime = "777h0m0s"
  	refresh_token_idle_expiration = "888h0m0s"
  	refresh_token_expiration = "999h0m0s"
}`, resourceName, frame.UniqueResourcesID, accessTokenLifetime)
		},
		initialProperty, updatedProperty,
		"", "",
		checkRemoteProperty(*frame),
		test_utils.ZITADEL_GENERATED_ID_REGEX,
		test_utils.CheckNothing,
		nil, nil, "", "",
	)
}

func checkRemoteProperty(frame test_utils.InstanceTestFrame) func(interface{}) resource.TestCheckFunc {
	return func(expect interface{}) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			resp, err := frame.GetOIDCSettings(frame, &admin.GetOIDCSettingsRequest{})
			if err != nil {
				return fmt.Errorf("getting oidc settings failed: %w", err)
			}
			actual := resp.GetSettings().GetAccessTokenLifetime().AsDuration().String()
			if actual != expect {
				return fmt.Errorf("expected %s, but got %s", expect, actual)
			}
			return nil
		}
	}
}