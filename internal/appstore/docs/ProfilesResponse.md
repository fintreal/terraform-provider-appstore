# ProfilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Profile**](Profile.md) |  | 
**Included** | Pointer to [**[]ProfilesResponseIncludedInner**](ProfilesResponseIncludedInner.md) |  | [optional] 

## Methods

### NewProfilesResponse

`func NewProfilesResponse(data []Profile, ) *ProfilesResponse`

NewProfilesResponse instantiates a new ProfilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilesResponseWithDefaults

`func NewProfilesResponseWithDefaults() *ProfilesResponse`

NewProfilesResponseWithDefaults instantiates a new ProfilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProfilesResponse) GetData() []Profile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProfilesResponse) GetDataOk() (*[]Profile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProfilesResponse) SetData(v []Profile)`

SetData sets Data field to given value.


### GetIncluded

`func (o *ProfilesResponse) GetIncluded() []ProfilesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProfilesResponse) GetIncludedOk() (*[]ProfilesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProfilesResponse) SetIncluded(v []ProfilesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProfilesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


