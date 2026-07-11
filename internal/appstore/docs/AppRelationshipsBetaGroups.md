# AppRelationshipsBetaGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**RelationshipLinks**](RelationshipLinks.md) |  | [optional] 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Data** | Pointer to [**[]AppRelationshipsBetaGroupsDataInner**](AppRelationshipsBetaGroupsDataInner.md) |  | [optional] 

## Methods

### NewAppRelationshipsBetaGroups

`func NewAppRelationshipsBetaGroups() *AppRelationshipsBetaGroups`

NewAppRelationshipsBetaGroups instantiates a new AppRelationshipsBetaGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRelationshipsBetaGroupsWithDefaults

`func NewAppRelationshipsBetaGroupsWithDefaults() *AppRelationshipsBetaGroups`

NewAppRelationshipsBetaGroupsWithDefaults instantiates a new AppRelationshipsBetaGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AppRelationshipsBetaGroups) GetLinks() RelationshipLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppRelationshipsBetaGroups) GetLinksOk() (*RelationshipLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppRelationshipsBetaGroups) SetLinks(v RelationshipLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppRelationshipsBetaGroups) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AppRelationshipsBetaGroups) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppRelationshipsBetaGroups) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppRelationshipsBetaGroups) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppRelationshipsBetaGroups) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *AppRelationshipsBetaGroups) GetData() []AppRelationshipsBetaGroupsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppRelationshipsBetaGroups) GetDataOk() (*[]AppRelationshipsBetaGroupsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppRelationshipsBetaGroups) SetData(v []AppRelationshipsBetaGroupsDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AppRelationshipsBetaGroups) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


