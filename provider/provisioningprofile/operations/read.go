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

	resp, _, err := client.ProfilesAPI.ProfilesGetInstance(context.Background(), id).Include([]string{"bundleId", "certificates"}).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	data := resp.GetData()
	attributes := data.GetAttributes()
	relationships := data.GetRelationships()

	d.SetId(data.GetId())

	if err := d.Set("id", data.GetId()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", attributes.GetName()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("type", attributes.GetProfileType()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("bundle_identifier_id", relationships.GetBundleId().Data.GetId()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("content_base64", attributes.GetProfileContent()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	var cretificateIds []string
	for _, certificate := range relationships.GetCertificates().Data {
		cretificateIds = append(cretificateIds, certificate.GetId())
	}
	if err := d.Set("certificate_ids", cretificateIds); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
