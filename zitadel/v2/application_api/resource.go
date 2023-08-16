package application_api

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/app"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing an API application belonging to a project, with all configuration possibilities.",
		Schema: map[string]*schema.Schema{
			helper.OrgIDVar: helper.OrgIDResourceField,
			ProjectIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the project",
				ForceNew:    true,
			},
			nameVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the application",
			},
			authMethodTypeVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Auth method type" + helper.DescriptionEnumValuesList(app.APIAuthMethodType_name),
				ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
					return helper.EnumValueValidation(authMethodTypeVar, value, app.APIAuthMethodType_value)
				},
				Default: app.APIAuthMethodType_name[0],
			},
			ClientIDVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "generated ID for this config",
				Sensitive:   true,
			},
			ClientSecretVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "generated secret for this config",
				Sensitive:   true,
			},
		},
		DeleteContext: delete,
		CreateContext: create,
		UpdateContext: update,
		ReadContext:   read,
		Importer: &schema.ResourceImporter{
			StateContext: helper.ImportWithIDAndOptionalOrg(
				helper.ResourceIDVar,
				helper.ImportAttribute{Key: ProjectIDVar, ValueFromString: helper.ConvertID},
				helper.ImportAttribute{Key: ClientIDVar, ValueFromString: helper.ConvertNonEmpty, Optional: true},
				helper.ImportAttribute{Key: ClientSecretVar, ValueFromString: helper.ConvertNonEmpty, Optional: true},
			),
		},
	}
}
