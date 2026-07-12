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
// disabled as sub-resources of an existing bundle id, keyed by the deterministic
// "<bundleId>_<TYPE>" id, without ever touching the bundle id itself.
func TestUpdateBundleIdCapabilitiesInPlace(t *testing.T) {
	created := createBundleId(t)
	bundleId := created.Data.GetId()
	defer deleteBundleId(bundleId)

	appleAuthID := addCapability(t, bundleId, "APPLE_ID_AUTH")
	assert.Equal(t, bundleId+"_APPLE_ID_AUTH", appleAuthID)

	pushID := addCapability(t, bundleId, "PUSH_NOTIFICATIONS")
	assert.Equal(t, bundleId+"_PUSH_NOTIFICATIONS", pushID)

	// Disable one capability by its deterministic id, leaving the other in place.
	_, err := apiClient.BundleIdCapabilitiesAPI.
		BundleIdCapabilitiesDeleteInstance(context.Background(), bundleId+"_PUSH_NOTIFICATIONS").Execute()
	assert.NoError(t, err)
}

// addCapability enables a single capability on an existing bundle id, matching
// the request the provider builds (APPLE_ID_AUTH carries its consent setting),
// and returns the created capability's id.
func addCapability(t *testing.T, bundleId, capabilityType string) string {
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
	resp, _, err := apiClient.BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance(context.Background()).
		BundleIdCapabilityCreateRequest(input).Execute()
	assert.NoError(t, err)
	return resp.Data.GetId()
}
