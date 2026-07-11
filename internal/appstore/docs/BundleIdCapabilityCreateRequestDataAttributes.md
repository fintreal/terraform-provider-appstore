# BundleIdCapabilityCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapabilityType** | [**CapabilityType**](CapabilityType.md) |  | 
**Settings** | [**[]CapabilitySetting**](CapabilitySetting.md) |  | 

## Methods

### NewBundleIdCapabilityCreateRequestDataAttributes

`func NewBundleIdCapabilityCreateRequestDataAttributes(capabilityType CapabilityType, settings []CapabilitySetting, ) *BundleIdCapabilityCreateRequestDataAttributes`

NewBundleIdCapabilityCreateRequestDataAttributes instantiates a new BundleIdCapabilityCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdCapabilityCreateRequestDataAttributesWithDefaults

`func NewBundleIdCapabilityCreateRequestDataAttributesWithDefaults() *BundleIdCapabilityCreateRequestDataAttributes`

NewBundleIdCapabilityCreateRequestDataAttributesWithDefaults instantiates a new BundleIdCapabilityCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilityType

`func (o *BundleIdCapabilityCreateRequestDataAttributes) GetCapabilityType() CapabilityType`

GetCapabilityType returns the CapabilityType field if non-nil, zero value otherwise.

### GetCapabilityTypeOk

`func (o *BundleIdCapabilityCreateRequestDataAttributes) GetCapabilityTypeOk() (*CapabilityType, bool)`

GetCapabilityTypeOk returns a tuple with the CapabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityType

`func (o *BundleIdCapabilityCreateRequestDataAttributes) SetCapabilityType(v CapabilityType)`

SetCapabilityType sets CapabilityType field to given value.


### GetSettings

`func (o *BundleIdCapabilityCreateRequestDataAttributes) GetSettings() []CapabilitySetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BundleIdCapabilityCreateRequestDataAttributes) GetSettingsOk() (*[]CapabilitySetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BundleIdCapabilityCreateRequestDataAttributes) SetSettings(v []CapabilitySetting)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


