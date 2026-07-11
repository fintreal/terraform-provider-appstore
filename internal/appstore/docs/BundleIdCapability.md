# BundleIdCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BundleIdCapabilityAttributes**](BundleIdCapabilityAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewBundleIdCapability

`func NewBundleIdCapability(type_ string, id string, ) *BundleIdCapability`

NewBundleIdCapability instantiates a new BundleIdCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdCapabilityWithDefaults

`func NewBundleIdCapabilityWithDefaults() *BundleIdCapability`

NewBundleIdCapabilityWithDefaults instantiates a new BundleIdCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BundleIdCapability) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleIdCapability) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleIdCapability) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BundleIdCapability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleIdCapability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleIdCapability) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BundleIdCapability) GetAttributes() BundleIdCapabilityAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleIdCapability) GetAttributesOk() (*BundleIdCapabilityAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleIdCapability) SetAttributes(v BundleIdCapabilityAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BundleIdCapability) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *BundleIdCapability) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BundleIdCapability) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BundleIdCapability) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BundleIdCapability) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


