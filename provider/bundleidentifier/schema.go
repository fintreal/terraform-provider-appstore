package bundleidentifier

import (
	"github.com/fintreal/terraform-provider-appstore/provider/bundleidentifier/operations"
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
			"identifier": {
				Description: "identifier like `com.example.app`",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"capabilities": {
				Description: "Capabilities of the app",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice(capabilities, true),
				},
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

var capabilities = []string{
	"ICLOUD",
	"IN_APP_PURCHASE",
	"GAME_CENTER",
	"PUSH_NOTIFICATIONS",
	"WALLET",
	"INTER_APP_AUDIO",
	"MAPS",
	"ASSOCIATED_DOMAINS",
	"PERSONAL_VPN",
	"APP_GROUPS",
	"HEALTHKIT",
	"HOMEKIT",
	"WIRELESS_ACCESSORY_CONFIGURATION",
	"APPLE_PAY",
	"DATA_PROTECTION",
	"SIRIKIT",
	"NETWORK_EXTENSIONS",
	"MULTIPATH",
	"HOT_SPOT",
	"NFC_TAG_READING",
	"CLASSKIT",
	"AUTOFILL_CREDENTIAL_PROVIDER",
	"ACCESS_WIFI_INFORMATION",
	"NETWORK_CUSTOM_PROTOCOL",
	"COREMEDIA_HLS_LOW_LATENCY",
	"SYSTEM_EXTENSION_INSTALL",
	"USER_MANAGEMENT",
	"APPLE_ID_AUTH",
}
