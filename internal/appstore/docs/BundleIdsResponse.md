# BundleIdsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BundleId**](BundleId.md) |  | 
**Included** | Pointer to [**[]BundleIdsResponseIncludedInner**](BundleIdsResponseIncludedInner.md) |  | [optional] 

## Methods

### NewBundleIdsResponse

`func NewBundleIdsResponse(data []BundleId, ) *BundleIdsResponse`

NewBundleIdsResponse instantiates a new BundleIdsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdsResponseWithDefaults

`func NewBundleIdsResponseWithDefaults() *BundleIdsResponse`

NewBundleIdsResponseWithDefaults instantiates a new BundleIdsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BundleIdsResponse) GetData() []BundleId`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BundleIdsResponse) GetDataOk() (*[]BundleId, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BundleIdsResponse) SetData(v []BundleId)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BundleIdsResponse) GetIncluded() []BundleIdsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BundleIdsResponse) GetIncludedOk() (*[]BundleIdsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BundleIdsResponse) SetIncluded(v []BundleIdsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BundleIdsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


