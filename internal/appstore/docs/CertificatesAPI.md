# \CertificatesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificatesCreateInstance**](CertificatesAPI.md#CertificatesCreateInstance) | **Post** /v1/certificates | 
[**CertificatesDeleteInstance**](CertificatesAPI.md#CertificatesDeleteInstance) | **Delete** /v1/certificates/{id} | 
[**CertificatesGetCollection**](CertificatesAPI.md#CertificatesGetCollection) | **Get** /v1/certificates | 
[**CertificatesGetInstance**](CertificatesAPI.md#CertificatesGetInstance) | **Get** /v1/certificates/{id} | 
[**CertificatesUpdateInstance**](CertificatesAPI.md#CertificatesUpdateInstance) | **Patch** /v1/certificates/{id} | 



## CertificatesCreateInstance

> CertificateResponse CertificatesCreateInstance(ctx).CertificateCreateRequest(certificateCreateRequest).Execute()



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
	certificateCreateRequest := *openapiclient.NewCertificateCreateRequest(*openapiclient.NewCertificateCreateRequestData("Type_example", *openapiclient.NewCertificateCreateRequestDataAttributes("CsrContent_example", openapiclient.CertificateType("APPLE_PAY")))) // CertificateCreateRequest | Certificate representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CertificatesCreateInstance(context.Background()).CertificateCreateRequest(certificateCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificatesCreateInstance`: CertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CertificatesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateCreateRequest** | [**CertificateCreateRequest**](CertificateCreateRequest.md) | Certificate representation | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesDeleteInstance

> CertificatesDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.CertificatesAPI.CertificatesDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificatesDeleteInstanceRequest struct via the builder pattern


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


## CertificatesGetCollection

> CertificatesResponse CertificatesGetCollection(ctx).FilterSerialNumber(filterSerialNumber).Execute()



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
	filterSerialNumber := []string{"Inner_example"} // []string | filter by attribute 'serialNumber' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CertificatesGetCollection(context.Background()).FilterSerialNumber(filterSerialNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificatesGetCollection`: CertificatesResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CertificatesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSerialNumber** | **[]string** | filter by attribute &#39;serialNumber&#39; | 

### Return type

[**CertificatesResponse**](CertificatesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesGetInstance

> CertificateResponse CertificatesGetInstance(ctx, id).Execute()



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
	resp, r, err := apiClient.CertificatesAPI.CertificatesGetInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificatesGetInstance`: CertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CertificatesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesUpdateInstance

> CertificateResponse CertificatesUpdateInstance(ctx, id).CertificateUpdateRequest(certificateUpdateRequest).Execute()



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
	certificateUpdateRequest := *openapiclient.NewCertificateUpdateRequest(*openapiclient.NewCertificateUpdateRequestData("Type_example", "Id_example")) // CertificateUpdateRequest | Certificate representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CertificatesUpdateInstance(context.Background(), id).CertificateUpdateRequest(certificateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificatesUpdateInstance`: CertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CertificatesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateUpdateRequest** | [**CertificateUpdateRequest**](CertificateUpdateRequest.md) | Certificate representation | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

