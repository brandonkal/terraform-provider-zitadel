package machine_user

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/user"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.",
		Schema: map[string]*schema.Schema{
			helper.OrgIDVar: helper.OrgIDResourceField,
			userStateVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the user",
				/* Not necessary as long as only active users are created
				ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
					return EnumValueValidation(userStateVar, value.(string), user.UserState_value)
				},*/
			},
			userNameVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username",
			},
			loginNamesVar: {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Loginnames",
			},
			preferredLoginNameVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Preferred login name",
			},
			nameVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the machine user",
			},
			descriptionVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the user",
			},
			accessTokenTypeVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Access token type" + helper.DescriptionEnumValuesList(user.AccessTokenType_name),
				ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
					return helper.EnumValueValidation(accessTokenTypeVar, value, user.AccessTokenType_value)
				},
				Default: defaultAccessTokenType,
			},
		},
		ReadContext:   read,
		CreateContext: create,
		DeleteContext: delete,
		UpdateContext: update,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}
