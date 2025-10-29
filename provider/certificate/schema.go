package certificate

import (
	"github.com/fintreal/terraform-provider-appstore/provider/certificate/operations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"id": {
				Description:  "id",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"id", "serial_number"},
			},
			"serial_number": {
				Description:  "serial number",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"id", "serial_number"},
			},
			"name": {
				Description: "name",
				Type:        schema.TypeString,
				Required:    false,
			},
			"type": {
				Description: "type",
				Type:        schema.TypeString,
				Required:    false,
			},
			"content": {
				Description: "content",
				Type:        schema.TypeString,
				Required:    false,
				Sensitive:   true,
			},
		},
	}
}
