package operations

import (
	"context"

	openapi "github.com/fintreal/terraform-provider-appstore/internal/appstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Update reconciles capabilities in place. Apple models each capability as its
// own sub-resource, so toggling them never replaces the bundle identifier.
func Update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*openapi.APIClient)
	id := d.Id()

	oldRaw, newRaw := d.GetChange("capabilities")
	toAdd, toRemove := diffCapabilities(toStringList(oldRaw.([]any)), toStringList(newRaw.([]any)))

	for _, capability := range toAdd {
		request := NewBundleIdCapabilityRequest(capability, id)
		if _, _, err := client.BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance(ctx).
			BundleIdCapabilityCreateRequest(request).Execute(); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	// Apple keys each capability as "<bundleId>_<TYPE>" (see the id returned on
	// create), so the delete target can be built without a lookup.
	for _, capability := range toRemove {
		capabilityID := id + "_" + capability
		if _, err := client.BundleIdCapabilitiesAPI.BundleIdCapabilitiesDeleteInstance(ctx, capabilityID).Execute(); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}

func toStringList(in []any) []string {
	out := make([]string, 0, len(in))
	for _, v := range in {
		out = append(out, v.(string))
	}
	return out
}

// diffCapabilities returns the capability types to add and to remove to move the
// enabled set from old to next.
func diffCapabilities(old, next []string) (add, remove []string) {
	oldSet := make(map[string]bool, len(old))
	for _, capability := range old {
		oldSet[capability] = true
	}
	nextSet := make(map[string]bool, len(next))
	for _, capability := range next {
		nextSet[capability] = true
	}
	for _, capability := range next {
		if !oldSet[capability] {
			add = append(add, capability)
		}
	}
	for _, capability := range old {
		if !nextSet[capability] {
			remove = append(remove, capability)
		}
	}
	return add, remove
}
