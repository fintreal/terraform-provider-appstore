package provisioningprofile

import (
	"github.com/fintreal/terraform-provider-appstore/provider/provisioningprofile/operations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: operations.Create,
		ReadContext:   operations.Read,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "identifier id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "name",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"type": {
				Description:  "provisioning profile type",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(profileTypes, true),
			},
			"bundle_identifier_id": {
				Description: "id of the bundle identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"certificate_ids": {
				Description: "certificate ids for the provisioning profile",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				ForceNew: true,
			},
			"content_base64": {
				Description: "base64 encoded content of the provisioning profile",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

var profileTypes = []string{
	"IOS_APP_DEVELOPMENT",
	"IOS_APP_STORE",
	"IOS_APP_ADHOC",
	"IOS_APP_INHOUSE",
	"MAC_APP_DEVELOPMENT",
	"MAC_APP_STORE",
	"MAC_APP_DIRECT",
	"TVOS_APP_DEVELOPMENT",
	"TVOS_APP_STORE",
	"TVOS_APP_ADHOC",
	"TVOS_APP_INHOUSE",
	"MAC_CATALYST_APP_DEVELOPMENT",
	"MAC_CATALYST_APP_STORE",
	"MAC_CATALYST_APP_DIRECT",
}
