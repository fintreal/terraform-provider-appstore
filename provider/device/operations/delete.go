package operations

import (
	"context"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*openapi.APIClient)

	// Apple's App Store Connect API has no DELETE for devices — a registered
	// device can only be disabled. So "deleting" the resource disables the device
	// (it still counts against the account's yearly device limit until the annual
	// reset) and drops it from Terraform state.
	if err := setDeviceStatus(ctx, client, d.Id(), "DISABLED"); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
