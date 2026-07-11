package operations

import (
	"context"

	openapi "github.com/fintreal/terraform-provider-appstore/internal/appstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*openapi.APIClient)

	udid := d.Get("udid").(string)
	name := d.Get("name").(string)
	platform := d.Get("platform").(string)

	// Apple never deletes devices, so this UDID may already be registered
	// (possibly DISABLED). Adopt it and ensure it is enabled instead of failing.
	existing, err := findDeviceByUdid(ctx, client, udid)
	if err != nil {
		return diag.FromErr(err)
	}
	if existing != nil {
		d.SetId(existing.GetId())
		existingAttrs := existing.GetAttributes()
		if existingAttrs.GetStatus() != "ENABLED" {
			if err := setDeviceStatus(ctx, client, existing.GetId(), "ENABLED"); err != nil {
				return diag.FromErr(err)
			}
		}
		return Read(ctx, d, m)
	}

	attributes := openapi.NewDeviceCreateRequestDataAttributes(name, openapi.BundleIdPlatform(platform), udid)
	data := openapi.NewDeviceCreateRequestData("devices", *attributes)
	request := openapi.NewDeviceCreateRequest(*data)

	resp, _, err := client.DevicesAPI.DevicesCreateInstance(ctx).DeviceCreateRequest(*request).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return setDeviceState(d, resp.GetData())
}
