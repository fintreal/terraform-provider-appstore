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
	serialNumber := d.Get("serial_number").(string)

	var cert *openapi.Certificate
	var err error
	if serialNumber != "" {
		cert, err = getBySerialNumber(ctx, client, serialNumber)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		cert, err = getById(ctx, client, id)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	attributes := cert.GetAttributes()

	d.SetId(cert.GetId())

	if err := d.Set("id", cert.GetId()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", attributes.GetName()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("type", attributes.GetCertificateType()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("serial_number", attributes.GetSerialNumber()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("content", attributes.GetCertificateContent()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func getById(ctx context.Context, client *openapi.APIClient, id string) (*openapi.Certificate, error) {
	resp, _, err := client.CertificatesAPI.CertificatesGetInstance(ctx, id).Execute()
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func getBySerialNumber(ctx context.Context, client *openapi.APIClient, serialNumber string) (*openapi.Certificate, error) {
	resp, _, err := client.CertificatesAPI.CertificatesGetCollection(ctx).FilterSerialNumber([]string{serialNumber}).Execute()
	if err != nil {
		return nil, err
	}
	return &resp.Data[0], nil
}
