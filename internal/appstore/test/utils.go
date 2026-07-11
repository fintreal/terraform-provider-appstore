package test

import (
	"context"
	"math/rand"
	"os"
	"testing"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/stretchr/testify/require"
)

// distributionCertID is a real iOS Distribution certificate in the account.
// Certificates cannot be created via the API (they require a CSR), so this is
// the one fixed reference the integration tests rely on. Everything else the
// tests need (bundle IDs, capabilities, profiles) is created and deleted per run.
const distributionCertID = "64A2UF5HD4"

var apiClient = GetClient()

func GetClient() *openapi.APIClient {
	keyData := os.Getenv("KEY_DATA")
	keyId := os.Getenv("KEY_ID")
	issuerId := os.Getenv("ISSUER_ID")
	var token, err = openapi.GenerateToken(keyData, keyId, issuerId)
	if err != nil {
		panic(err)
	}
	var configuration = openapi.NewConfiguration(*token)
	var apiClient = openapi.NewAPIClient(configuration)

	return apiClient
}

const charset = "abcdefghijklmnopqrstuvwxyz"

func GenerateRandomString() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// createBundleId registers a throwaway explicit App ID for a single test and
// fails the test immediately if creation does not succeed.
func createBundleId(t *testing.T) *openapi.BundleIdResponse {
	t.Helper()
	identifier := "com.fintreal.test." + GenerateRandomString()
	input := *openapi.NewBundleIdCreateRequest(
		*openapi.NewBundleIdCreateRequestData(
			"bundleIds",
			*openapi.NewBundleIdCreateRequestDataAttributes("Integration Test Bundle ID", "IOS", identifier),
		),
	)
	resp, _, err := apiClient.BundleIdsAPI.BundleIdsCreateInstance(context.Background()).BundleIdCreateRequest(input).Execute()
	require.NoError(t, err)
	return resp
}

func deleteBundleId(id string) {
	apiClient.BundleIdsAPI.BundleIdsDeleteInstance(context.Background(), id).Execute()
}

// createProfile provisions an App Store profile for the given bundle ID using
// the distribution certificate, failing the test immediately on error.
func createProfile(t *testing.T, bundleId, name string) *openapi.ProfileResponse {
	t.Helper()
	input := *openapi.NewProfileCreateRequest(
		*openapi.NewProfileCreateRequestData(
			"profiles",
			*openapi.NewProfileCreateRequestDataAttributes(name, "IOS_APP_STORE"),
			*openapi.NewProfileCreateRequestDataRelationships(
				*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(
					*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("bundleIds", bundleId),
				),
				*openapi.NewProfileCreateRequestDataRelationshipsCertificates(
					[]openapi.ProfileRelationshipsCertificatesDataInner{{Type: "certificates", Id: distributionCertID}},
				),
			),
		),
	)
	resp, _, err := apiClient.ProfilesAPI.ProfilesCreateInstance(context.Background()).ProfileCreateRequest(input).Execute()
	require.NoError(t, err)
	return resp
}

func deleteProfile(id string) {
	apiClient.ProfilesAPI.ProfilesDeleteInstance(context.Background(), id).Execute()
}
