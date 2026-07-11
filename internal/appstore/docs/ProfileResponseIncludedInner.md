# ProfileResponseIncludedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**CertificateAttributes**](CertificateAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**BundleIdRelationships**](BundleIdRelationships.md) |  | [optional] 

## Methods

### NewProfileResponseIncludedInner

`func NewProfileResponseIncludedInner(type_ string, id string, ) *ProfileResponseIncludedInner`

NewProfileResponseIncludedInner instantiates a new ProfileResponseIncludedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileResponseIncludedInnerWithDefaults

`func NewProfileResponseIncludedInnerWithDefaults() *ProfileResponseIncludedInner`

NewProfileResponseIncludedInnerWithDefaults instantiates a new ProfileResponseIncludedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProfileResponseIncludedInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileResponseIncludedInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileResponseIncludedInner) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ProfileResponseIncludedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileResponseIncludedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileResponseIncludedInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ProfileResponseIncludedInner) GetAttributes() CertificateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProfileResponseIncludedInner) GetAttributesOk() (*CertificateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProfileResponseIncludedInner) SetAttributes(v CertificateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProfileResponseIncludedInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ProfileResponseIncludedInner) GetRelationships() BundleIdRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ProfileResponseIncludedInner) GetRelationshipsOk() (*BundleIdRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ProfileResponseIncludedInner) SetRelationships(v BundleIdRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ProfileResponseIncludedInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


