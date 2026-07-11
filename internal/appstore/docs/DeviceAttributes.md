# DeviceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to [**BundleIdPlatform**](BundleIdPlatform.md) |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 
**DeviceClass** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**AddedDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeviceAttributes

`func NewDeviceAttributes() *DeviceAttributes`

NewDeviceAttributes instantiates a new DeviceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAttributesWithDefaults

`func NewDeviceAttributesWithDefaults() *DeviceAttributes`

NewDeviceAttributesWithDefaults instantiates a new DeviceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceAttributes) GetPlatform() BundleIdPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceAttributes) GetPlatformOk() (*BundleIdPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceAttributes) SetPlatform(v BundleIdPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUdid

`func (o *DeviceAttributes) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *DeviceAttributes) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *DeviceAttributes) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *DeviceAttributes) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetDeviceClass

`func (o *DeviceAttributes) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *DeviceAttributes) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *DeviceAttributes) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.

### HasDeviceClass

`func (o *DeviceAttributes) HasDeviceClass() bool`

HasDeviceClass returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetModel

`func (o *DeviceAttributes) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceAttributes) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceAttributes) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceAttributes) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAddedDate

`func (o *DeviceAttributes) GetAddedDate() time.Time`

GetAddedDate returns the AddedDate field if non-nil, zero value otherwise.

### GetAddedDateOk

`func (o *DeviceAttributes) GetAddedDateOk() (*time.Time, bool)`

GetAddedDateOk returns a tuple with the AddedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedDate

`func (o *DeviceAttributes) SetAddedDate(v time.Time)`

SetAddedDate sets AddedDate field to given value.

### HasAddedDate

`func (o *DeviceAttributes) HasAddedDate() bool`

HasAddedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


