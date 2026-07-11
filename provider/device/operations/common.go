package operations

import (
	"context"
	"net/http"

	openapi "github.com/fintreal/terraform-provider-appstore/internal/appstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// setDeviceState writes a Device's fields into the resource data.
func setDeviceState(d *schema.ResourceData, device openapi.Device) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId(device.GetId())
	attrs := device.GetAttributes()
	values := map[string]any{
		"id":           device.GetId(),
		"udid":         attrs.GetUdid(),
		"name":         attrs.GetName(),
		"platform":     string(attrs.GetPlatform()),
		"status":       attrs.GetStatus(),
		"device_class": attrs.GetDeviceClass(),
		"model":        attrs.GetModel(),
	}
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

// findDeviceByUdid returns the device with the given UDID, or nil if none exists.
func findDeviceByUdid(ctx context.Context, client *openapi.APIClient, udid string) (*openapi.Device, error) {
	resp, _, err := client.DevicesAPI.DevicesGetCollection(ctx).FilterUdid([]string{udid}).Execute()
	if err != nil {
		return nil, err
	}
	for _, device := range resp.GetData() {
		attrs := device.GetAttributes()
		if attrs.GetUdid() == udid {
			found := device
			return &found, nil
		}
	}
	return nil, nil
}

// setDeviceStatus patches a device to ENABLED or DISABLED.
func setDeviceStatus(ctx context.Context, client *openapi.APIClient, id string, status string) error {
	attrs := openapi.NewDeviceUpdateRequestDataAttributes()
	attrs.SetStatus(status)
	data := openapi.NewDeviceUpdateRequestData("devices", id)
	data.SetAttributes(*attrs)
	req := openapi.NewDeviceUpdateRequest(*data)
	_, _, err := client.DevicesAPI.DevicesUpdateInstance(ctx, id).DeviceUpdateRequest(*req).Execute()
	return err
}

func isNotFound(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusNotFound
}
