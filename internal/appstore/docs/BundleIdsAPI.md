# \BundleIdsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BundleIdsCreateInstance**](BundleIdsAPI.md#BundleIdsCreateInstance) | **Post** /v1/bundleIds | 
[**BundleIdsDeleteInstance**](BundleIdsAPI.md#BundleIdsDeleteInstance) | **Delete** /v1/bundleIds/{id} | 
[**BundleIdsGetCollection**](BundleIdsAPI.md#BundleIdsGetCollection) | **Get** /v1/bundleIds | 
[**BundleIdsGetInstance**](BundleIdsAPI.md#BundleIdsGetInstance) | **Get** /v1/bundleIds/{id} | 
[**BundleIdsUpdateInstance**](BundleIdsAPI.md#BundleIdsUpdateInstance) | **Patch** /v1/bundleIds/{id} | 



## BundleIdsCreateInstance

> BundleIdResponse BundleIdsCreateInstance(ctx).BundleIdCreateRequest(bundleIdCreateRequest).Execute()



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
	bundleIdCreateRequest := *openapiclient.NewBundleIdCreateRequest(*openapiclient.NewBundleIdCreateRequestData("Type_example", *openapiclient.NewBundleIdCreateRequestDataAttributes("Name_example", openapiclient.BundleIdPlatform("IOS"), "Identifier_example"))) // BundleIdCreateRequest | BundleId representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleIdsAPI.BundleIdsCreateInstance(context.Background()).BundleIdCreateRequest(bundleIdCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BundleIdsCreateInstance`: BundleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleIdCreateRequest** | [**BundleIdCreateRequest**](BundleIdCreateRequest.md) | BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsDeleteInstance

> BundleIdsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.BundleIdsAPI.BundleIdsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBundleIdsDeleteInstanceRequest struct via the builder pattern


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


## BundleIdsGetCollection

> BundleIdsResponse BundleIdsGetCollection(ctx).Execute()



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
	resp, r, err := apiClient.BundleIdsAPI.BundleIdsGetCollection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BundleIdsGetCollection`: BundleIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsGetCollection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsGetCollectionRequest struct via the builder pattern


### Return type

[**BundleIdsResponse**](BundleIdsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsGetInstance

> BundleIdResponse BundleIdsGetInstance(ctx, id).Include(include).Execute()



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
	resp, r, err := apiClient.BundleIdsAPI.BundleIdsGetInstance(context.Background(), id).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BundleIdsGetInstance`: BundleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsUpdateInstance

> BundleIdResponse BundleIdsUpdateInstance(ctx, id).BundleIdUpdateRequest(bundleIdUpdateRequest).Execute()



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
	bundleIdUpdateRequest := *openapiclient.NewBundleIdUpdateRequest(*openapiclient.NewBundleIdUpdateRequestData("Type_example", "Id_example")) // BundleIdUpdateRequest | BundleId representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleIdsAPI.BundleIdsUpdateInstance(context.Background(), id).BundleIdUpdateRequest(bundleIdUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BundleIdsUpdateInstance`: BundleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleIdUpdateRequest** | [**BundleIdUpdateRequest**](BundleIdUpdateRequest.md) | BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

