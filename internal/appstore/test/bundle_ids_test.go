package test

import (
	"context"
	"testing"

	openapi "github.com/fintreal/terraform-provider-appstore/internal/appstore"
	"github.com/stretchr/testify/assert"
)

func TestGetBundleIds(t *testing.T) {
	apiClient.BundleIdsAPI.BundleIdsGetCollection(context.Background()).Execute()
}

func TestGetBundleId(t *testing.T) {
	created := createBundleId(t)
	id := created.Data.GetId()
	defer deleteBundleId(id)

	bundleId, _, err := apiClient.BundleIdsAPI.BundleIdsGetInstance(context.Background(), id).Execute()
	assert.NoError(t, err)
	assert.Equal(t, id, bundleId.Data.GetId())
	assert.Equal(t, created.Data.Attributes.GetIdentifier(), bundleId.Data.Attributes.GetIdentifier())
	assert.Equal(t, "Integration Test Bundle ID", bundleId.Data.Attributes.GetName())
}

func TestCreateAndDeleteBundleId(t *testing.T) {
	identifier := "com.fintreal.test." + GenerateRandomString()
	name := "Integration Test Bundle ID"
	input := *openapi.NewBundleIdCreateRequest(
		*openapi.NewBundleIdCreateRequestData(
			"bundleIds",
			*openapi.NewBundleIdCreateRequestDataAttributes(name, "IOS", identifier),
		),
	)

	data, _, err := apiClient.BundleIdsAPI.BundleIdsCreateInstance(context.Background()).BundleIdCreateRequest(input).Execute()

	assert.Equal(t, identifier, data.Data.Attributes.GetIdentifier())
	assert.Equal(t, name, data.Data.Attributes.GetName())
	assert.NoError(t, err)

	_, err = apiClient.BundleIdsAPI.BundleIdsDeleteInstance(context.Background(), data.Data.GetId()).Execute()
	assert.NoError(t, err)
}
