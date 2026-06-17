# DeviceCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Platform** | [**BundleIdPlatform**](BundleIdPlatform.md) |  | 
**Udid** | **string** |  | 

## Methods

### NewDeviceCreateRequestDataAttributes

`func NewDeviceCreateRequestDataAttributes(name string, platform BundleIdPlatform, udid string, ) *DeviceCreateRequestDataAttributes`

NewDeviceCreateRequestDataAttributes instantiates a new DeviceCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateRequestDataAttributesWithDefaults

`func NewDeviceCreateRequestDataAttributesWithDefaults() *DeviceCreateRequestDataAttributes`

NewDeviceCreateRequestDataAttributesWithDefaults instantiates a new DeviceCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceCreateRequestDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceCreateRequestDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceCreateRequestDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *DeviceCreateRequestDataAttributes) GetPlatform() BundleIdPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceCreateRequestDataAttributes) GetPlatformOk() (*BundleIdPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceCreateRequestDataAttributes) SetPlatform(v BundleIdPlatform)`

SetPlatform sets Platform field to given value.


### GetUdid

`func (o *DeviceCreateRequestDataAttributes) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *DeviceCreateRequestDataAttributes) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *DeviceCreateRequestDataAttributes) SetUdid(v string)`

SetUdid sets Udid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


