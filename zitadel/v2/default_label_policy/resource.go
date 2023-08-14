package default_label_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing the default label policy.",
		Schema: map[string]*schema.Schema{
			primaryColorVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for primary color",
			},
			hideLoginNameSuffixVar: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "hides the org suffix on the login form if the scope \"urn:zitadel:iam:org:domain:primary:{domainname}\" is set. Details about this scope in https://docs.zitadel.ch/concepts#Reserved_Scopes",
			},
			warnColorVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for warn color",
			},
			backgroundColorVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for background color",
			},
			fontColorVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for font color",
			},
			primaryColorDarkVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for primary color dark theme",
			},
			backgroundColorDarkVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for background color dark theme",
			},
			warnColorDarkVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for warn color dark theme",
			},
			fontColorDarkVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "hex value for font color dark theme",
			},
			disableWatermarkVar: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "disable watermark",
			},
			logoPathVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			logoHashVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			logoURLVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			iconPathVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			iconHashVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			iconURLVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			logoDarkPathVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			logoDarkHashVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			logoURLDarkVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			iconDarkPathVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			iconDarkHashVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			iconURLDarkVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			fontPathVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			fontHashVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			fontURLVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			setActiveVar: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "set the label policy active after creating/updating",
			},
		},
		ReadContext:   read,
		CreateContext: update,
		DeleteContext: delete,
		UpdateContext: update,
		Importer:      &schema.ResourceImporter{StateContext: helper.ImportWithOptionalIDV5("instance_id")},
	}
}
