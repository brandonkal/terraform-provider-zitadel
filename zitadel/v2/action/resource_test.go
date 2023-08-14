package action_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/management"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
)

func TestAccAction(t *testing.T) {
	resourceName := "zitadel_action"
	initialProperty := "initialproperty"
	updatedProperty := "updatedproperty"
	frame, err := test_utils.NewOrgTestFrame(resourceName)
	if err != nil {
		t.Fatalf("setting up test context failed: %v", err)
	}
	test_utils.RunLifecyleTest[string](
		t,
		frame.BaseTestFrame,
		func(configProperty, _ string) string {
			return fmt.Sprintf(`
resource "%s" "%s" {
  org_id          = "%s"
  name            = "testaction"
  script          = "%s"
  timeout         = "10s"
  allowed_to_fail = true
}`, resourceName, frame.UniqueResourcesID, frame.OrgID, configProperty)
		},
		initialProperty, updatedProperty,
		"", "", "",
		true,
		checkRemoteProperty(frame),
		helper.ZitadelGeneratedIdOnlyRegex,
		test_utils.CheckIsNotFoundFromPropertyCheck(checkRemoteProperty(frame), updatedProperty),
		test_utils.ImportStateIdWithOrg(frame),
	)
}

func checkRemoteProperty(frame *test_utils.OrgTestFrame) func(string) resource.TestCheckFunc {
	return func(expect string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			remoteResource, err := frame.GetAction(frame, &management.GetActionRequest{Id: frame.State(state).ID})
			if err != nil {
				return err
			}
			actual := remoteResource.GetAction().GetScript()
			if actual != expect {
				return fmt.Errorf("expected %s, but got %s", expect, actual)
			}
			return nil
		}
	}
}
