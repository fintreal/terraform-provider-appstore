# DeviceCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**DeviceCreateRequestDataAttributes**](DeviceCreateRequestDataAttributes.md) |  | 

## Methods

### NewDeviceCreateRequestData

`func NewDeviceCreateRequestData(type_ string, attributes DeviceCreateRequestDataAttributes, ) *DeviceCreateRequestData`

NewDeviceCreateRequestData instantiates a new DeviceCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateRequestDataWithDefaults

`func NewDeviceCreateRequestDataWithDefaults() *DeviceCreateRequestData`

NewDeviceCreateRequestDataWithDefaults instantiates a new DeviceCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeviceCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *DeviceCreateRequestData) GetAttributes() DeviceCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeviceCreateRequestData) GetAttributesOk() (*DeviceCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeviceCreateRequestData) SetAttributes(v DeviceCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


