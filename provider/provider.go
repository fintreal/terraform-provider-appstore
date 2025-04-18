package provider

import (
	"context"

	appstore "github.com/fintreal/app-store-sdk-go"
	"github.com/fintreal/terraform-provider-appstore/provider/bundleidentifier"
	"github.com/fintreal/terraform-provider-appstore/provider/provisioningprofile"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"appstore_key": {
				Description: "App Store API Key. You can set this via `APPSTORE_KEY` environment variable.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPSTORE_KEY", ""),
			},
			"appstore_key_id": {
				Description: "App Store API Key ID. You can set this via `APPSTORE_KEY_ID` environment variable.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPSTORE_KEY_ID", ""),
			},
			"appstore_key_issuer_id": {
				Description: "App Store Issuer ID. You can set this via `APPSTORE_KEY_ISSUER_ID` environment variable.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPSTORE_KEY_ISSUER_ID", ""),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"appstore_bundle_identifier":    bundleidentifier.Resource(),
			"appstore_provisioning_profile": provisioningprofile.Resource(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			appstoreKey := d.Get("appstore_key").(string)
			appstoreKeyId := d.Get("appstore_key_id").(string)
			appstoreKeyIssuerId := d.Get("appstore_key_issuer_id").(string)

			var diags diag.Diagnostics

			if appstoreKey == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "`appstore_key` value cannot be an empty string",
					Detail:   "set the token value in the provider configuration or via the APPSTORE_KEY environment variable",
				})
			}

			if appstoreKeyId == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "`appstore_key_id` value cannot be an empty string",
					Detail:   "set the token value in the provider configuration or via the APPSTORE_KEY_ID environment variable",
				})
			}

			if appstoreKeyIssuerId == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "`appstore_issuer_id` value cannot be an empty string",
					Detail:   "set the token value in the provider configuration or via the APPSTORE_KEY_ISSUER_ID environment variable",
				})
			}

			if len(diags) > 0 {
				return nil, diags
			}

			token, err := appstore.GenerateToken(appstoreKey, appstoreKeyId, appstoreKeyIssuerId)

			if err != nil || *token == "" {
				return nil, diag.FromErr(err)
			}

			configuration := appstore.NewConfiguration(*token)

			client := appstore.NewAPIClient(configuration)

			return client, diags
		},
	}
}
