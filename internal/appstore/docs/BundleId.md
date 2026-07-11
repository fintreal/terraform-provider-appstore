# BundleId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BundleIdAttributes**](BundleIdAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**BundleIdRelationships**](BundleIdRelationships.md) |  | [optional] 

## Methods

### NewBundleId

`func NewBundleId(type_ string, id string, ) *BundleId`

NewBundleId instantiates a new BundleId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdWithDefaults

`func NewBundleIdWithDefaults() *BundleId`

NewBundleIdWithDefaults instantiates a new BundleId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BundleId) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleId) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleId) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BundleId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleId) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BundleId) GetAttributes() BundleIdAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleId) GetAttributesOk() (*BundleIdAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleId) SetAttributes(v BundleIdAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BundleId) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BundleId) GetRelationships() BundleIdRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BundleId) GetRelationshipsOk() (*BundleIdRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BundleId) SetRelationships(v BundleIdRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BundleId) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


