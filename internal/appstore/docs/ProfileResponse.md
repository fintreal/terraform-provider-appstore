# ProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Profile**](Profile.md) |  | 
**Included** | Pointer to [**[]ProfileResponseIncludedInner**](ProfileResponseIncludedInner.md) |  | [optional] 

## Methods

### NewProfileResponse

`func NewProfileResponse(data Profile, ) *ProfileResponse`

NewProfileResponse instantiates a new ProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileResponseWithDefaults

`func NewProfileResponseWithDefaults() *ProfileResponse`

NewProfileResponseWithDefaults instantiates a new ProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProfileResponse) GetData() Profile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProfileResponse) GetDataOk() (*Profile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProfileResponse) SetData(v Profile)`

SetData sets Data field to given value.


### GetIncluded

`func (o *ProfileResponse) GetIncluded() []ProfileResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProfileResponse) GetIncludedOk() (*[]ProfileResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProfileResponse) SetIncluded(v []ProfileResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProfileResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


