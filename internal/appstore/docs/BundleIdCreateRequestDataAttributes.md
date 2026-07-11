# BundleIdCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Platform** | [**BundleIdPlatform**](BundleIdPlatform.md) |  | 
**Identifier** | **string** |  | 
**SeedId** | Pointer to **string** |  | [optional] 

## Methods

### NewBundleIdCreateRequestDataAttributes

`func NewBundleIdCreateRequestDataAttributes(name string, platform BundleIdPlatform, identifier string, ) *BundleIdCreateRequestDataAttributes`

NewBundleIdCreateRequestDataAttributes instantiates a new BundleIdCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdCreateRequestDataAttributesWithDefaults

`func NewBundleIdCreateRequestDataAttributesWithDefaults() *BundleIdCreateRequestDataAttributes`

NewBundleIdCreateRequestDataAttributesWithDefaults instantiates a new BundleIdCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BundleIdCreateRequestDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleIdCreateRequestDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleIdCreateRequestDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *BundleIdCreateRequestDataAttributes) GetPlatform() BundleIdPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BundleIdCreateRequestDataAttributes) GetPlatformOk() (*BundleIdPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BundleIdCreateRequestDataAttributes) SetPlatform(v BundleIdPlatform)`

SetPlatform sets Platform field to given value.


### GetIdentifier

`func (o *BundleIdCreateRequestDataAttributes) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BundleIdCreateRequestDataAttributes) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BundleIdCreateRequestDataAttributes) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSeedId

`func (o *BundleIdCreateRequestDataAttributes) GetSeedId() string`

GetSeedId returns the SeedId field if non-nil, zero value otherwise.

### GetSeedIdOk

`func (o *BundleIdCreateRequestDataAttributes) GetSeedIdOk() (*string, bool)`

GetSeedIdOk returns a tuple with the SeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedId

`func (o *BundleIdCreateRequestDataAttributes) SetSeedId(v string)`

SetSeedId sets SeedId field to given value.

### HasSeedId

`func (o *BundleIdCreateRequestDataAttributes) HasSeedId() bool`

HasSeedId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


