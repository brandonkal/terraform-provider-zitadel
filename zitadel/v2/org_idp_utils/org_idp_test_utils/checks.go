package org_idp_test_utils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/management"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
)

func CheckProviderName(frame test_utils.OrgTestFrame) func(string) resource.TestCheckFunc {
	return func(expectName string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			remoteProvider, err := frame.GetProviderByID(frame, &management.GetProviderByIDRequest{Id: frame.State(state).ID})
			if err != nil {
				return err
			}
			actual := remoteProvider.GetIdp().GetName()
			if actual != expectName {
				return fmt.Errorf("expected name %s, but got name %s", expectName, actual)
			}
			return nil
		}
	}
}

func CheckDestroy(frame test_utils.OrgTestFrame) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		err := CheckProviderName(frame)("")(state)
		if status.Code(err) != codes.NotFound {
			return fmt.Errorf("expected not found error but got: %w", err)
		}
		return nil
	}
}
