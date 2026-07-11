# \DevicesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesCreateInstance**](DevicesAPI.md#DevicesCreateInstance) | **Post** /v1/devices | 
[**DevicesGetCollection**](DevicesAPI.md#DevicesGetCollection) | **Get** /v1/devices | 
[**DevicesGetInstance**](DevicesAPI.md#DevicesGetInstance) | **Get** /v1/devices/{id} | 
[**DevicesUpdateInstance**](DevicesAPI.md#DevicesUpdateInstance) | **Patch** /v1/devices/{id} | 



## DevicesCreateInstance

> DeviceResponse DevicesCreateInstance(ctx).DeviceCreateRequest(deviceCreateRequest).Execute()



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
	deviceCreateRequest := *openapiclient.NewDeviceCreateRequest(*openapiclient.NewDeviceCreateRequestData("Type_example", *openapiclient.NewDeviceCreateRequestDataAttributes("Name_example", openapiclient.BundleIdPlatform("IOS"), "Udid_example"))) // DeviceCreateRequest | Device representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.DevicesCreateInstance(context.Background()).DeviceCreateRequest(deviceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesCreateInstance`: DeviceResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevicesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCreateRequest** | [**DeviceCreateRequest**](DeviceCreateRequest.md) | Device representation | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesGetCollection

> DevicesResponse DevicesGetCollection(ctx).FilterUdid(filterUdid).Limit(limit).Cursor(cursor).Execute()



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
	filterUdid := []string{"Inner_example"} // []string | filter by device UDID (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	cursor := "cursor_example" // string | pagination cursor from links.next (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.DevicesGetCollection(context.Background()).FilterUdid(filterUdid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesGetCollection`: DevicesResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevicesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterUdid** | **[]string** | filter by device UDID | 
 **limit** | **int32** | maximum resources per page | 
 **cursor** | **string** | pagination cursor from links.next | 

### Return type

[**DevicesResponse**](DevicesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesGetInstance

> DeviceResponse DevicesGetInstance(ctx, id).Execute()



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
	resp, r, err := apiClient.DevicesAPI.DevicesGetInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesGetInstance`: DeviceResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevicesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUpdateInstance

> DeviceResponse DevicesUpdateInstance(ctx, id).DeviceUpdateRequest(deviceUpdateRequest).Execute()



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
	deviceUpdateRequest := *openapiclient.NewDeviceUpdateRequest(*openapiclient.NewDeviceUpdateRequestData("Type_example", "Id_example")) // DeviceUpdateRequest | Device representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.DevicesUpdateInstance(context.Background(), id).DeviceUpdateRequest(deviceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesUpdateInstance`: DeviceResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevicesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceUpdateRequest** | [**DeviceUpdateRequest**](DeviceUpdateRequest.md) | Device representation | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

