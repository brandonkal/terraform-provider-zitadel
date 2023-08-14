package project_grant_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/management"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
)

func TestAccProjectGrant(t *testing.T) {
	resourceName := "zitadel_project_grant"
	initialProperty := "initialProperty"
	updatedProperty := "updatedProperty"
	frame, err := test_utils.NewOrgTestFrame(resourceName)
	if err != nil {
		t.Fatalf("setting up test context failed: %v", err)
	}
	project, err := frame.AddProject(frame, &management.AddProjectRequest{
		Name: frame.UniqueResourcesID,
	})
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	projectID := project.GetId()
	for _, role := range []string{initialProperty, updatedProperty} {
		_, err = frame.AddProjectRole(frame, &management.AddProjectRoleRequest{
			ProjectId:   projectID,
			RoleKey:     role,
			DisplayName: role,
		})
		if err != nil {
			t.Fatalf("failed to create project role %s: %v", role, err)
		}
	}
	org, err := frame.AddOrg(frame, &management.AddOrgRequest{
		Name: frame.UniqueResourcesID,
	})
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}
	grantedOrgID := org.GetId()
	test_utils.RunLifecyleTest[string](
		t,
		frame.BaseTestFrame,
		func(configProperty, _ string) string {
			return fmt.Sprintf(`
resource "%s" "%s" {
  org_id         = "%s"
  project_id     = "%s"
  granted_org_id = "%s"
  role_keys      = ["%s"]
}`, resourceName, frame.UniqueResourcesID, frame.OrgID, projectID, grantedOrgID, configProperty)
		},
		initialProperty, updatedProperty,
		"", "", "",
		false,
		checkRemoteProperty(*frame, projectID),
		helper.ZitadelGeneratedIdOnlyRegex,
		test_utils.CheckIsNotFoundFromPropertyCheck(checkRemoteProperty(*frame, projectID), ""),
		nil,
	)
}

func checkRemoteProperty(frame test_utils.OrgTestFrame, projectID string) func(string) resource.TestCheckFunc {
	return func(expect string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			resp, err := frame.GetProjectGrantByID(frame, &management.GetProjectGrantByIDRequest{
				ProjectId: projectID,
				GrantId:   frame.State(state).ID,
			})
			if err != nil {
				return err
			}
			actualRoleKeys := resp.GetProjectGrant().GetGrantedRoleKeys()
			if len(actualRoleKeys) != 1 {
				return fmt.Errorf("expected 1 role, but got %d", len(actualRoleKeys))
			}
			if expect != actualRoleKeys[0] {
				return fmt.Errorf("expected role key %s, but got %s", expect, actualRoleKeys[0])
			}
			return nil
		}
	}
}
