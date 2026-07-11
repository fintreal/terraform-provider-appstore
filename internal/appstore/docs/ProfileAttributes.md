# ProfileAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to [**BundleIdPlatform**](BundleIdPlatform.md) |  | [optional] 
**ProfileType** | Pointer to **string** |  | [optional] 
**ProfileState** | Pointer to **string** |  | [optional] 
**ProfileContent** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProfileAttributes

`func NewProfileAttributes() *ProfileAttributes`

NewProfileAttributes instantiates a new ProfileAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileAttributesWithDefaults

`func NewProfileAttributesWithDefaults() *ProfileAttributes`

NewProfileAttributesWithDefaults instantiates a new ProfileAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProfileAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProfileAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *ProfileAttributes) GetPlatform() BundleIdPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ProfileAttributes) GetPlatformOk() (*BundleIdPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ProfileAttributes) SetPlatform(v BundleIdPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ProfileAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProfileType

`func (o *ProfileAttributes) GetProfileType() string`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *ProfileAttributes) GetProfileTypeOk() (*string, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *ProfileAttributes) SetProfileType(v string)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *ProfileAttributes) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### GetProfileState

`func (o *ProfileAttributes) GetProfileState() string`

GetProfileState returns the ProfileState field if non-nil, zero value otherwise.

### GetProfileStateOk

`func (o *ProfileAttributes) GetProfileStateOk() (*string, bool)`

GetProfileStateOk returns a tuple with the ProfileState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileState

`func (o *ProfileAttributes) SetProfileState(v string)`

SetProfileState sets ProfileState field to given value.

### HasProfileState

`func (o *ProfileAttributes) HasProfileState() bool`

HasProfileState returns a boolean if a field has been set.

### GetProfileContent

`func (o *ProfileAttributes) GetProfileContent() string`

GetProfileContent returns the ProfileContent field if non-nil, zero value otherwise.

### GetProfileContentOk

`func (o *ProfileAttributes) GetProfileContentOk() (*string, bool)`

GetProfileContentOk returns a tuple with the ProfileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileContent

`func (o *ProfileAttributes) SetProfileContent(v string)`

SetProfileContent sets ProfileContent field to given value.

### HasProfileContent

`func (o *ProfileAttributes) HasProfileContent() bool`

HasProfileContent returns a boolean if a field has been set.

### GetUuid

`func (o *ProfileAttributes) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProfileAttributes) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProfileAttributes) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProfileAttributes) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ProfileAttributes) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProfileAttributes) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProfileAttributes) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ProfileAttributes) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ProfileAttributes) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ProfileAttributes) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ProfileAttributes) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ProfileAttributes) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


