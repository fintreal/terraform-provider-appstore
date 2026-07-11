package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCertificate(t *testing.T) {
	data, _, err := apiClient.CertificatesAPI.CertificatesGetInstance(context.Background(), "64A2UF5HD4").Execute()

	assert.NoError(t, err)
	assert.Equal(t, "64A2UF5HD4", data.Data.Id)
	assert.Equal(t, "certificates", data.Data.Type)
	assert.Equal(t, "6AB7A99542D0D58517914E6EA8119579", data.Data.Attributes.SerialNumber)
	assert.NotEmpty(t, data.Data.Attributes.Name)
	assert.NotEmpty(t, data.Data.Attributes.CertificateType)
	assert.NotEmpty(t, data.Data.Attributes.DisplayName)
}

func TestGetBySerialNumber(t *testing.T) {
	serialNumber := "6AB7A99542D0D58517914E6EA8119579"
	data, _, err := apiClient.CertificatesAPI.CertificatesGetCollection(context.Background()).FilterSerialNumber([]string{serialNumber}).Execute()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(data.Data))
	assert.Equal(t, serialNumber, data.Data[0].Attributes.SerialNumber)
}
