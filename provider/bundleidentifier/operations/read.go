package operations

import (
	"context"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*openapi.APIClient)

	id := d.Get("id").(string)

	resp, _, err := client.BundleIdsAPI.BundleIdsGetInstance(context.Background(), id).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	data := resp.GetData()
	attributes := data.GetAttributes()

	d.SetId(data.GetId())

	if err := d.Set("id", data.GetId()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", attributes.GetName()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("identifier", attributes.GetIdentifier()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("platform", attributes.GetPlatform()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
