package operations

import (
	"context"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*openapi.APIClient)

	name := d.Get("name").(string)
	_type := d.Get("type").(string)
	bundleIdentifierId := d.Get("bundle_identifier_id").(string)
	certificateIdsAny := d.Get("certificate_ids").([]any)
	certificateIds := make([]string, len(certificateIdsAny))
	for i, v := range certificateIdsAny {
		certificateIds[i] = v.(string)
	}
	request := NewProfileCreateRequest(name, _type, bundleIdentifierId, certificateIds)

	resp, _, err := client.ProfilesAPI.ProfilesCreateInstance(context.Background()).ProfileCreateRequest(request).Execute()

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
	var newCertificateIds []string
	for _, certificate := range relationships.GetCertificates().Data {
		certificateIds = append(certificateIds, certificate.GetId())
	}
	if err := d.Set("certificate_ids", newCertificateIds); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func NewProfileCreateRequest(name string, _type string, bundleIdentifierId string, certificateIds []string) openapi.ProfileCreateRequest {
	var certificateIdData []openapi.ProfileRelationshipsCertificatesDataInner
	for _, certificateId := range certificateIds {
		certificateIdData = append(certificateIdData, openapi.ProfileRelationshipsCertificatesDataInner{Type: "certificates", Id: certificateId})
	}
	return *openapi.NewProfileCreateRequest(
		*openapi.NewProfileCreateRequestData(
			"profiles",
			*openapi.NewProfileCreateRequestDataAttributes(name, _type),
			*openapi.NewProfileCreateRequestDataRelationships(
				*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(
					*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("bundleIds", bundleIdentifierId),
				),
				*openapi.NewProfileCreateRequestDataRelationshipsCertificates(certificateIdData),
			),
		),
	)
}
