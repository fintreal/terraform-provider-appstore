package test

import (
	"context"
	"testing"

	openapi "github.com/fintreal/terraform-provider-appstore/internal/appstore"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteBundleIdCapability(t *testing.T) {
	created := createBundleId(t)
	bundleId := created.Data.GetId()
	defer deleteBundleId(bundleId)

	key := "APPLE_ID_AUTH_APP_CONSENT"
	optionKey := "PRIMARY_APP_CONSENT"
	capabilityType := "APPLE_ID_AUTH"
	input := *openapi.NewBundleIdCapabilityCreateRequest(
		*openapi.NewBundleIdCapabilityCreateRequestData(
			"bundleIdCapabilities",
			*openapi.NewBundleIdCapabilityCreateRequestDataAttributes(
				openapi.CapabilityType(capabilityType),
				[]openapi.CapabilitySetting{{
					Key: &key,
					Options: []openapi.CapabilityOption{{
						Key: &optionKey,
					}},
				}},
			),
			*openapi.NewBundleIdCapabilityCreateRequestDataRelationships(
				*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(
					*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("bundleIds", bundleId),
				),
			),
		),
	)
	data, _, err := apiClient.BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance(context.Background()).BundleIdCapabilityCreateRequest(input).Execute()
	assert.NoError(t, err)
	assert.Equal(t, capabilityType, string(data.Data.Attributes.GetCapabilityType()))
	assert.Equal(t, bundleId+"_"+capabilityType, data.Data.GetId())

	_, err = apiClient.BundleIdCapabilitiesAPI.BundleIdCapabilitiesDeleteInstance(context.Background(), data.Data.GetId()).Execute()
	assert.NoError(t, err)
}

// TestUpdateBundleIdCapabilitiesInPlace exercises the mechanism the
// appstore_bundle_identifier Update relies on: capabilities are enabled and
// disabled as sub-resources of an existing bundle id, and the currently enabled
// set is discoverable via the bundle id's included bundleIdCapabilities.
func TestUpdateBundleIdCapabilitiesInPlace(t *testing.T) {
	created := createBundleId(t)
	bundleId := created.Data.GetId()
	defer deleteBundleId(bundleId)

	addCapability(t, bundleId, "APPLE_ID_AUTH")
	addCapability(t, bundleId, "PUSH_NOTIFICATIONS")

	enabled := capabilityIDsByType(t, bundleId)
	assert.Contains(t, enabled, "APPLE_ID_AUTH")
	assert.Contains(t, enabled, "PUSH_NOTIFICATIONS")

	_, err := apiClient.BundleIdCapabilitiesAPI.
		BundleIdCapabilitiesDeleteInstance(context.Background(), enabled["PUSH_NOTIFICATIONS"]).Execute()
	assert.NoError(t, err)

	remaining := capabilityIDsByType(t, bundleId)
	assert.Contains(t, remaining, "APPLE_ID_AUTH")
	assert.NotContains(t, remaining, "PUSH_NOTIFICATIONS")
}

// capabilityIDsByType lists the bundle id's enabled capabilities keyed by type,
// mirroring the provider's Include-based lookup used to delete capabilities.
func capabilityIDsByType(t *testing.T, bundleId string) map[string]string {
	t.Helper()
	resp, _, err := apiClient.BundleIdsAPI.BundleIdsGetInstance(context.Background(), bundleId).
		Include([]string{"bundleIdCapabilities"}).Execute()
	assert.NoError(t, err)
	out := map[string]string{}
	for _, included := range resp.GetIncluded() {
		if capability := included.BundleIdCapability; capability != nil {
			attributes := capability.GetAttributes()
			out[string(attributes.GetCapabilityType())] = capability.GetId()
		}
	}
	return out
}

// addCapability enables a single capability on an existing bundle id, matching
// the request the provider builds (APPLE_ID_AUTH carries its consent setting).
func addCapability(t *testing.T, bundleId, capabilityType string) {
	t.Helper()
	settings := []openapi.CapabilitySetting{}
	if capabilityType == "APPLE_ID_AUTH" {
		key := "APPLE_ID_AUTH_APP_CONSENT"
		optionKey := "PRIMARY_APP_CONSENT"
		settings = append(settings, openapi.CapabilitySetting{
			Key:     &key,
			Options: []openapi.CapabilityOption{{Key: &optionKey}},
		})
	}
	input := *openapi.NewBundleIdCapabilityCreateRequest(
		*openapi.NewBundleIdCapabilityCreateRequestData(
			"bundleIdCapabilities",
			*openapi.NewBundleIdCapabilityCreateRequestDataAttributes(openapi.CapabilityType(capabilityType), settings),
			*openapi.NewBundleIdCapabilityCreateRequestDataRelationships(
				*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(
					*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("bundleIds", bundleId),
				),
			),
		),
	)
	_, _, err := apiClient.BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance(context.Background()).
		BundleIdCapabilityCreateRequest(input).Execute()
	assert.NoError(t, err)
}
