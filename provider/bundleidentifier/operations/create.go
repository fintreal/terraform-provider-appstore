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

	identifier := d.Get("identifier").(string)
	name := d.Get("name").(string)
	platform := "IOS"
	capabilities := d.Get("capabilities").([]any)

	request := NewBundleIdCreateRequest(name, platform, identifier)

	resp, _, err := client.BundleIdsAPI.BundleIdsCreateInstance(context.Background()).BundleIdCreateRequest(request).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	data := resp.GetData()
	attributes := data.GetAttributes()
	d.SetId(data.GetId())

	var newCapabilities []any
	for _, capability := range capabilities {
		request := NewBundleIdCapabilityRequest(capability.(string), data.GetId())
		resp, _, err := client.BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance(context.Background()).BundleIdCapabilityCreateRequest(request).Execute()
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
		newCapabilities = append(newCapabilities, resp.GetData().Attributes.GetCapabilityType())
	}
	if len(diags) > 0 {
		return diags
	}

	if err := d.Set("id", data.GetId()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", attributes.GetName()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("identifier", attributes.GetIdentifier()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if len(newCapabilities) > 0 {
		if err := d.Set("capabilities", newCapabilities); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func NewBundleIdCreateRequest(name string, platform string, identifier string) openapi.BundleIdCreateRequest {
	return *openapi.NewBundleIdCreateRequest(
		*openapi.NewBundleIdCreateRequestData(
			"bundleIds",
			*openapi.NewBundleIdCreateRequestDataAttributes(name, openapi.BundleIdPlatform(platform), identifier),
		),
	)
}

func NewBundleIdCapabilityRequest(capability string, bundleIdentifierId string) openapi.BundleIdCapabilityCreateRequest {
	settings := []openapi.CapabilitySetting{}
	if capability == "APPLE_ID_AUTH" {
		key := "APPLE_ID_AUTH_APP_CONSENT"
		optionKey := "PRIMARY_APP_CONSENT"
		setting := openapi.CapabilitySetting{
			Key: &key,
			Options: []openapi.CapabilityOption{{
				Key: &optionKey,
			}},
		}
		settings = append(settings, setting)
	}
	return *openapi.NewBundleIdCapabilityCreateRequest(
		*openapi.NewBundleIdCapabilityCreateRequestData(
			"bundleIdCapabilities",
			*openapi.NewBundleIdCapabilityCreateRequestDataAttributes(openapi.CapabilityType(capability), settings),
			*openapi.NewBundleIdCapabilityCreateRequestDataRelationships(
				*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(
					*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("bundleIds", bundleIdentifierId),
				),
			),
		),
	)
}
