package device

import (
	"github.com/fintreal/terraform-provider-appstore/provider/device/operations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Resource defines the appstore_device resource.
//
// Apple's App Store Connect API has no DELETE for devices — they can only be
// disabled. So `terraform destroy` (or removing a device from config) sets the
// device's status to DISABLED rather than deleting it. A device that is created
// for a UDID that already exists is adopted and re-enabled.
func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: operations.Create,
		ReadContext:   operations.Read,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "device id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"udid": {
				Description: "device UDID",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "device name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"platform": {
				Description:  "device platform (IOS or MAC_OS)",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "IOS",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"IOS", "MAC_OS"}, false),
			},
			"status": {
				Description: "device status (ENABLED or DISABLED) — managed by the resource lifecycle",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"device_class": {
				Description: "device class as reported by Apple (e.g. IPHONE, IPAD)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"model": {
				Description: "device model as reported by Apple",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}
