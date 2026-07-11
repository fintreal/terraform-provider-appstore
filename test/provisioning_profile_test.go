package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProvisioningProfile(t *testing.T) {
	bundle := createBundleId(t)
	defer deleteBundleId(bundle.Data.GetId())

	created := createProfile(t, bundle.Data.GetId(), "Test Provisioning Profile")
	defer deleteProfile(created.Data.Id)

	data, _, err := apiClient.ProfilesAPI.ProfilesGetInstance(context.Background(), created.Data.Id).Include([]string{"bundleId", "certificates"}).Execute()
	assert.NoError(t, err)
	assert.Equal(t, created.Data.Id, data.Data.Id)
	assert.Equal(t, "Test Provisioning Profile", data.Data.Attributes.GetName())
	assert.Equal(t, "IOS_APP_STORE", data.Data.Attributes.GetProfileType())
	assert.Equal(t, bundle.Data.GetId(), data.Data.Relationships.GetBundleId().Data.GetId())
	assert.Equal(t, distributionCertID, data.Data.Relationships.GetCertificates().Data[0].GetId())
}

func TestCreateAndDeleteProvisioningProfile(t *testing.T) {
	bundle := createBundleId(t)
	defer deleteBundleId(bundle.Data.GetId())

	created := createProfile(t, bundle.Data.GetId(), "Test Provisioning Profile X")

	assert.Equal(t, "Test Provisioning Profile X", created.Data.Attributes.GetName())
	assert.Equal(t, "IOS_APP_STORE", created.Data.Attributes.GetProfileType())

	_, err := apiClient.ProfilesAPI.ProfilesDeleteInstance(context.Background(), created.Data.Id).Execute()
	assert.NoError(t, err)
}
