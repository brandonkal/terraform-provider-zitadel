package user_grant

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing the authorization given to a user directly, including the given roles.",
		Schema: map[string]*schema.Schema{
			userGrantProjectIDVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the project",
				ForceNew:    true,
			},
			userGrantProjectGrantIDVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the granted project",
				ForceNew:    true,
			},
			userGrantUserIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the user",
				ForceNew:    true,
			},
			userGrantRoleKeysVar: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "List of roles granted",
			},
			userGrantOrgIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the organization which owns the resource",
				ForceNew:    true,
			},
		},
		DeleteContext: delete,
		CreateContext: create,
		UpdateContext: update,
		ReadContext:   read,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}
