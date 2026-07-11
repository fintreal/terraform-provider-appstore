# ProfileCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**ProfileCreateRequestDataAttributes**](ProfileCreateRequestDataAttributes.md) |  | 
**Relationships** | [**ProfileCreateRequestDataRelationships**](ProfileCreateRequestDataRelationships.md) |  | 

## Methods

### NewProfileCreateRequestData

`func NewProfileCreateRequestData(type_ string, attributes ProfileCreateRequestDataAttributes, relationships ProfileCreateRequestDataRelationships, ) *ProfileCreateRequestData`

NewProfileCreateRequestData instantiates a new ProfileCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileCreateRequestDataWithDefaults

`func NewProfileCreateRequestDataWithDefaults() *ProfileCreateRequestData`

NewProfileCreateRequestDataWithDefaults instantiates a new ProfileCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProfileCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ProfileCreateRequestData) GetAttributes() ProfileCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProfileCreateRequestData) GetAttributesOk() (*ProfileCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProfileCreateRequestData) SetAttributes(v ProfileCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ProfileCreateRequestData) GetRelationships() ProfileCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ProfileCreateRequestData) GetRelationshipsOk() (*ProfileCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ProfileCreateRequestData) SetRelationships(v ProfileCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


