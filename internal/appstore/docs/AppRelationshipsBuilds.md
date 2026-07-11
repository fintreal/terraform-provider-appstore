# AppRelationshipsBuilds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**RelationshipLinks**](RelationshipLinks.md) |  | [optional] 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Data** | Pointer to [**[]AppRelationshipsBuildsDataInner**](AppRelationshipsBuildsDataInner.md) |  | [optional] 

## Methods

### NewAppRelationshipsBuilds

`func NewAppRelationshipsBuilds() *AppRelationshipsBuilds`

NewAppRelationshipsBuilds instantiates a new AppRelationshipsBuilds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRelationshipsBuildsWithDefaults

`func NewAppRelationshipsBuildsWithDefaults() *AppRelationshipsBuilds`

NewAppRelationshipsBuildsWithDefaults instantiates a new AppRelationshipsBuilds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AppRelationshipsBuilds) GetLinks() RelationshipLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppRelationshipsBuilds) GetLinksOk() (*RelationshipLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppRelationshipsBuilds) SetLinks(v RelationshipLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppRelationshipsBuilds) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AppRelationshipsBuilds) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppRelationshipsBuilds) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppRelationshipsBuilds) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppRelationshipsBuilds) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *AppRelationshipsBuilds) GetData() []AppRelationshipsBuildsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppRelationshipsBuilds) GetDataOk() (*[]AppRelationshipsBuildsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppRelationshipsBuilds) SetData(v []AppRelationshipsBuildsDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AppRelationshipsBuilds) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


