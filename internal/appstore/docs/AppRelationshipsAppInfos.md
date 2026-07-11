# AppRelationshipsAppInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**RelationshipLinks**](RelationshipLinks.md) |  | [optional] 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Data** | Pointer to [**[]AppRelationshipsAppInfosDataInner**](AppRelationshipsAppInfosDataInner.md) |  | [optional] 

## Methods

### NewAppRelationshipsAppInfos

`func NewAppRelationshipsAppInfos() *AppRelationshipsAppInfos`

NewAppRelationshipsAppInfos instantiates a new AppRelationshipsAppInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRelationshipsAppInfosWithDefaults

`func NewAppRelationshipsAppInfosWithDefaults() *AppRelationshipsAppInfos`

NewAppRelationshipsAppInfosWithDefaults instantiates a new AppRelationshipsAppInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AppRelationshipsAppInfos) GetLinks() RelationshipLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppRelationshipsAppInfos) GetLinksOk() (*RelationshipLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppRelationshipsAppInfos) SetLinks(v RelationshipLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppRelationshipsAppInfos) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AppRelationshipsAppInfos) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppRelationshipsAppInfos) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppRelationshipsAppInfos) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppRelationshipsAppInfos) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *AppRelationshipsAppInfos) GetData() []AppRelationshipsAppInfosDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppRelationshipsAppInfos) GetDataOk() (*[]AppRelationshipsAppInfosDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppRelationshipsAppInfos) SetData(v []AppRelationshipsAppInfosDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AppRelationshipsAppInfos) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


