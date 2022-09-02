package idp_oidc

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/idp"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing a OIDC IDP of the organization.",
		Schema: map[string]*schema.Schema{
			orgIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the organization",
				ForceNew:    true,
			},
			nameVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the IDP",
			},
			stylingTypeVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Some identity providers specify the styling of the button to their login",
				ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
					return helper.EnumValueValidation(stylingTypeVar, value, idp.IDPStylingType_value)
				},
			},
			clientIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "client id generated by the identity provider",
				Sensitive:   true,
			},
			clientSecretVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "client secret generated by the identity provider",
				Sensitive:   true,
			},
			issuerVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "the oidc issuer of the identity provider",
			},
			scopesVar: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "the scopes requested by ZITADEL during the request on the identity provider",
			},
			displayNameMappingVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "definition which field is mapped to the display name of the user",
			},
			usernameMappingVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "definition which field is mapped to the email of the user",
			},
			autoRegisterVar: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "auto register for users from this idp",
			},
		},
		ReadContext:   read,
		UpdateContext: update,
		CreateContext: create,
		DeleteContext: delete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}
