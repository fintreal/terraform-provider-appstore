package test

import (
	"context"
	"os"
	"testing"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/stretchr/testify/assert"
)

// TestGetDevices exercises the devices collection endpoint. Read-only, so it is
// safe to run on every CI run.
func TestGetDevices(t *testing.T) {
	_, _, err := apiClient.DevicesAPI.DevicesGetCollection(context.Background()).Execute()
	assert.NoError(t, err)
}

// TestCreateGetUpdateDevice registers a device, reads it back, renames it, then
// disables it.
//
// Apple cannot delete devices, and every new UDID permanently counts against the
// account's yearly device limit. So this test is gated: it uses a FIXED UDID from
// APPSTORE_TEST_DEVICE_UDID (which is re-adopted on later runs, consuming at most
// one slot) and skips when that variable is unset — never burning quota on CI by
// accident.
func TestCreateGetUpdateDevice(t *testing.T) {
	udid := os.Getenv("APPSTORE_TEST_DEVICE_UDID")
	if udid == "" {
		t.Skip("set APPSTORE_TEST_DEVICE_UDID (a dedicated test UDID) to run the device mutation test")
	}
	ctx := context.Background()

	// Create, or adopt the existing device if this UDID was registered before.
	createInput := *openapi.NewDeviceCreateRequest(
		*openapi.NewDeviceCreateRequestData(
			"devices",
			*openapi.NewDeviceCreateRequestDataAttributes("Integration Test Device", "IOS", udid),
		),
	)
	var deviceId string
	created, _, err := apiClient.DevicesAPI.DevicesCreateInstance(ctx).DeviceCreateRequest(createInput).Execute()
	if err != nil {
		list, _, lerr := apiClient.DevicesAPI.DevicesGetCollection(ctx).FilterUdid([]string{udid}).Execute()
		assert.NoError(t, lerr)
		assert.NotEmpty(t, list.GetData())
		deviceId = list.GetData()[0].GetId()
	} else {
		deviceId = created.Data.GetId()
		assert.Equal(t, udid, created.Data.Attributes.GetUdid())
	}

	// Read it back.
	got, _, err := apiClient.DevicesAPI.DevicesGetInstance(ctx, deviceId).Execute()
	assert.NoError(t, err)
	assert.Equal(t, udid, got.Data.Attributes.GetUdid())

	// Rename it.
	newName := "Integration Test Device " + GenerateRandomString()
	updAttrs := openapi.NewDeviceUpdateRequestDataAttributes()
	updAttrs.SetName(newName)
	updData := openapi.NewDeviceUpdateRequestData("devices", deviceId)
	updData.SetAttributes(*updAttrs)
	updated, _, err := apiClient.DevicesAPI.DevicesUpdateInstance(ctx, deviceId).DeviceUpdateRequest(*openapi.NewDeviceUpdateRequest(*updData)).Execute()
	assert.NoError(t, err)
	assert.Equal(t, newName, updated.Data.Attributes.GetName())

	// Disable it (the closest thing to deletion Apple allows).
	disAttrs := openapi.NewDeviceUpdateRequestDataAttributes()
	disAttrs.SetStatus("DISABLED")
	disData := openapi.NewDeviceUpdateRequestData("devices", deviceId)
	disData.SetAttributes(*disAttrs)
	_, _, err = apiClient.DevicesAPI.DevicesUpdateInstance(ctx, deviceId).DeviceUpdateRequest(*openapi.NewDeviceUpdateRequest(*disData)).Execute()
	assert.NoError(t, err)
}
