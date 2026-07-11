# AppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]App**](App.md) |  | 
**Included** | Pointer to [**[]AppsResponseIncludedInner**](AppsResponseIncludedInner.md) |  | [optional] 

## Methods

### NewAppsResponse

`func NewAppsResponse(data []App, ) *AppsResponse`

NewAppsResponse instantiates a new AppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsResponseWithDefaults

`func NewAppsResponseWithDefaults() *AppsResponse`

NewAppsResponseWithDefaults instantiates a new AppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppsResponse) GetData() []App`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppsResponse) GetDataOk() (*[]App, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppsResponse) SetData(v []App)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppsResponse) GetIncluded() []AppsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppsResponse) GetIncludedOk() (*[]AppsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppsResponse) SetIncluded(v []AppsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


