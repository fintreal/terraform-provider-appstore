# AppRelationshipsSubscriptionGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**RelationshipLinks**](RelationshipLinks.md) |  | [optional] 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Data** | Pointer to [**[]AppRelationshipsSubscriptionGroupsDataInner**](AppRelationshipsSubscriptionGroupsDataInner.md) |  | [optional] 

## Methods

### NewAppRelationshipsSubscriptionGroups

`func NewAppRelationshipsSubscriptionGroups() *AppRelationshipsSubscriptionGroups`

NewAppRelationshipsSubscriptionGroups instantiates a new AppRelationshipsSubscriptionGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRelationshipsSubscriptionGroupsWithDefaults

`func NewAppRelationshipsSubscriptionGroupsWithDefaults() *AppRelationshipsSubscriptionGroups`

NewAppRelationshipsSubscriptionGroupsWithDefaults instantiates a new AppRelationshipsSubscriptionGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AppRelationshipsSubscriptionGroups) GetLinks() RelationshipLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppRelationshipsSubscriptionGroups) GetLinksOk() (*RelationshipLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppRelationshipsSubscriptionGroups) SetLinks(v RelationshipLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppRelationshipsSubscriptionGroups) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AppRelationshipsSubscriptionGroups) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppRelationshipsSubscriptionGroups) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppRelationshipsSubscriptionGroups) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppRelationshipsSubscriptionGroups) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *AppRelationshipsSubscriptionGroups) GetData() []AppRelationshipsSubscriptionGroupsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppRelationshipsSubscriptionGroups) GetDataOk() (*[]AppRelationshipsSubscriptionGroupsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppRelationshipsSubscriptionGroups) SetData(v []AppRelationshipsSubscriptionGroupsDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AppRelationshipsSubscriptionGroups) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


