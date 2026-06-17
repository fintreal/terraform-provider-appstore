# DevicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Device**](Device.md) |  | 
**Links** | Pointer to [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | [optional] 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewDevicesResponse

`func NewDevicesResponse(data []Device, ) *DevicesResponse`

NewDevicesResponse instantiates a new DevicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesResponseWithDefaults

`func NewDevicesResponseWithDefaults() *DevicesResponse`

NewDevicesResponseWithDefaults instantiates a new DevicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DevicesResponse) GetData() []Device`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DevicesResponse) GetDataOk() (*[]Device, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DevicesResponse) SetData(v []Device)`

SetData sets Data field to given value.


### GetLinks

`func (o *DevicesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevicesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevicesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevicesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DevicesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DevicesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DevicesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DevicesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


