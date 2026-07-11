# \ProfilesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfilesCreateInstance**](ProfilesAPI.md#ProfilesCreateInstance) | **Post** /v1/profiles | 
[**ProfilesDeleteInstance**](ProfilesAPI.md#ProfilesDeleteInstance) | **Delete** /v1/profiles/{id} | 
[**ProfilesGetCollection**](ProfilesAPI.md#ProfilesGetCollection) | **Get** /v1/profiles | 
[**ProfilesGetInstance**](ProfilesAPI.md#ProfilesGetInstance) | **Get** /v1/profiles/{id} | 



## ProfilesCreateInstance

> ProfileResponse ProfilesCreateInstance(ctx).ProfileCreateRequest(profileCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fintreal/app-store-sdk-go"
)

func main() {
	profileCreateRequest := *openapiclient.NewProfileCreateRequest(*openapiclient.NewProfileCreateRequestData("Type_example", *openapiclient.NewProfileCreateRequestDataAttributes("Name_example", "ProfileType_example"), *openapiclient.NewProfileCreateRequestDataRelationships(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("Type_example", "Id_example")), *openapiclient.NewProfileCreateRequestDataRelationshipsCertificates([]openapiclient.ProfileRelationshipsCertificatesDataInner{*openapiclient.NewProfileRelationshipsCertificatesDataInner("Type_example", "Id_example")})))) // ProfileCreateRequest | Profile representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ProfilesCreateInstance(context.Background()).ProfileCreateRequest(profileCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProfilesCreateInstance`: ProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProfilesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileCreateRequest** | [**ProfileCreateRequest**](ProfileCreateRequest.md) | Profile representation | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesDeleteInstance

> ProfilesDeleteInstance(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fintreal/app-store-sdk-go"
)

func main() {
	id := "id_example" // string | the id of the requested resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfilesAPI.ProfilesDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesDeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesGetCollection

> ProfilesResponse ProfilesGetCollection(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fintreal/app-store-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ProfilesGetCollection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProfilesGetCollection`: ProfilesResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesGetCollection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesGetCollectionRequest struct via the builder pattern


### Return type

[**ProfilesResponse**](ProfilesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesGetInstance

> ProfileResponse ProfilesGetInstance(ctx, id).Include(include).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fintreal/app-store-sdk-go"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ProfilesGetInstance(context.Background(), id).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProfilesGetInstance`: ProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

