# BundleIdUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BundleIdUpdateRequestDataAttributes**](BundleIdUpdateRequestDataAttributes.md) |  | [optional] 

## Methods

### NewBundleIdUpdateRequestData

`func NewBundleIdUpdateRequestData(type_ string, id string, ) *BundleIdUpdateRequestData`

NewBundleIdUpdateRequestData instantiates a new BundleIdUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdUpdateRequestDataWithDefaults

`func NewBundleIdUpdateRequestDataWithDefaults() *BundleIdUpdateRequestData`

NewBundleIdUpdateRequestDataWithDefaults instantiates a new BundleIdUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BundleIdUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleIdUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleIdUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BundleIdUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleIdUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleIdUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BundleIdUpdateRequestData) GetAttributes() BundleIdUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleIdUpdateRequestData) GetAttributesOk() (*BundleIdUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleIdUpdateRequestData) SetAttributes(v BundleIdUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BundleIdUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


