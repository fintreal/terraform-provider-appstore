# CapabilitySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnabledByDefault** | Pointer to **bool** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**AllowedInstances** | Pointer to **string** |  | [optional] 
**MinInstances** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**[]CapabilityOption**](CapabilityOption.md) |  | [optional] 

## Methods

### NewCapabilitySetting

`func NewCapabilitySetting() *CapabilitySetting`

NewCapabilitySetting instantiates a new CapabilitySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySettingWithDefaults

`func NewCapabilitySettingWithDefaults() *CapabilitySetting`

NewCapabilitySettingWithDefaults instantiates a new CapabilitySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CapabilitySetting) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CapabilitySetting) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CapabilitySetting) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CapabilitySetting) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CapabilitySetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilitySetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilitySetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilitySetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilitySetting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilitySetting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilitySetting) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilitySetting) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabledByDefault

`func (o *CapabilitySetting) GetEnabledByDefault() bool`

GetEnabledByDefault returns the EnabledByDefault field if non-nil, zero value otherwise.

### GetEnabledByDefaultOk

`func (o *CapabilitySetting) GetEnabledByDefaultOk() (*bool, bool)`

GetEnabledByDefaultOk returns a tuple with the EnabledByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledByDefault

`func (o *CapabilitySetting) SetEnabledByDefault(v bool)`

SetEnabledByDefault sets EnabledByDefault field to given value.

### HasEnabledByDefault

`func (o *CapabilitySetting) HasEnabledByDefault() bool`

HasEnabledByDefault returns a boolean if a field has been set.

### GetVisible

`func (o *CapabilitySetting) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *CapabilitySetting) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *CapabilitySetting) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *CapabilitySetting) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetAllowedInstances

`func (o *CapabilitySetting) GetAllowedInstances() string`

GetAllowedInstances returns the AllowedInstances field if non-nil, zero value otherwise.

### GetAllowedInstancesOk

`func (o *CapabilitySetting) GetAllowedInstancesOk() (*string, bool)`

GetAllowedInstancesOk returns a tuple with the AllowedInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInstances

`func (o *CapabilitySetting) SetAllowedInstances(v string)`

SetAllowedInstances sets AllowedInstances field to given value.

### HasAllowedInstances

`func (o *CapabilitySetting) HasAllowedInstances() bool`

HasAllowedInstances returns a boolean if a field has been set.

### GetMinInstances

`func (o *CapabilitySetting) GetMinInstances() int32`

GetMinInstances returns the MinInstances field if non-nil, zero value otherwise.

### GetMinInstancesOk

`func (o *CapabilitySetting) GetMinInstancesOk() (*int32, bool)`

GetMinInstancesOk returns a tuple with the MinInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstances

`func (o *CapabilitySetting) SetMinInstances(v int32)`

SetMinInstances sets MinInstances field to given value.

### HasMinInstances

`func (o *CapabilitySetting) HasMinInstances() bool`

HasMinInstances returns a boolean if a field has been set.

### GetOptions

`func (o *CapabilitySetting) GetOptions() []CapabilityOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CapabilitySetting) GetOptionsOk() (*[]CapabilityOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CapabilitySetting) SetOptions(v []CapabilityOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CapabilitySetting) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


