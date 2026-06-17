package operations

import (
	"context"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*openapi.APIClient)

	// Only `name` is mutable in place; `udid` and `platform` are ForceNew, and
	// `status` is driven by the resource lifecycle (create/delete), not config.
	attributes := openapi.NewDeviceUpdateRequestDataAttributes()
	attributes.SetName(d.Get("name").(string))

	data := openapi.NewDeviceUpdateRequestData("devices", d.Id())
	data.SetAttributes(*attributes)
	request := openapi.NewDeviceUpdateRequest(*data)

	resp, _, err := client.DevicesAPI.DevicesUpdateInstance(ctx, d.Id()).DeviceUpdateRequest(*request).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return setDeviceState(d, resp.GetData())
}
