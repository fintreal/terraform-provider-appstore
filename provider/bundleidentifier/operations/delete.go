package operations

import (
	"context"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*openapi.APIClient)

	id := d.Get("id").(string)

	_, err := client.BundleIdsAPI.BundleIdsDeleteInstance(context.Background(), id).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics
	return diags
}
