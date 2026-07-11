package operations

import (
	"context"

	openapi "github.com/fintreal/terraform-provider-appstore/internal/appstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*openapi.APIClient)

	resp, httpResp, err := client.DevicesAPI.DevicesGetInstance(ctx, d.Id()).Execute()
	if err != nil {
		if isNotFound(httpResp) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	return setDeviceState(d, resp.GetData())
}
