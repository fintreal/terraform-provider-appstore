# BundleIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BundleId**](BundleId.md) |  | 
**Included** | Pointer to [**[]BundleIdResponseIncludedInner**](BundleIdResponseIncludedInner.md) |  | [optional] 

## Methods

### NewBundleIdResponse

`func NewBundleIdResponse(data BundleId, ) *BundleIdResponse`

NewBundleIdResponse instantiates a new BundleIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdResponseWithDefaults

`func NewBundleIdResponseWithDefaults() *BundleIdResponse`

NewBundleIdResponseWithDefaults instantiates a new BundleIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BundleIdResponse) GetData() BundleId`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BundleIdResponse) GetDataOk() (*BundleId, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BundleIdResponse) SetData(v BundleId)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BundleIdResponse) GetIncluded() []BundleIdResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BundleIdResponse) GetIncludedOk() (*[]BundleIdResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BundleIdResponse) SetIncluded(v []BundleIdResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BundleIdResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


