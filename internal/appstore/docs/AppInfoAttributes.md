# AppInfoAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppStoreState** | Pointer to [**AppStoreVersionState**](AppStoreVersionState.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewAppInfoAttributes

`func NewAppInfoAttributes() *AppInfoAttributes`

NewAppInfoAttributes instantiates a new AppInfoAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoAttributesWithDefaults

`func NewAppInfoAttributesWithDefaults() *AppInfoAttributes`

NewAppInfoAttributesWithDefaults instantiates a new AppInfoAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppStoreState

`func (o *AppInfoAttributes) GetAppStoreState() AppStoreVersionState`

GetAppStoreState returns the AppStoreState field if non-nil, zero value otherwise.

### GetAppStoreStateOk

`func (o *AppInfoAttributes) GetAppStoreStateOk() (*AppStoreVersionState, bool)`

GetAppStoreStateOk returns a tuple with the AppStoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreState

`func (o *AppInfoAttributes) SetAppStoreState(v AppStoreVersionState)`

SetAppStoreState sets AppStoreState field to given value.

### HasAppStoreState

`func (o *AppInfoAttributes) HasAppStoreState() bool`

HasAppStoreState returns a boolean if a field has been set.

### GetState

`func (o *AppInfoAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppInfoAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppInfoAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AppInfoAttributes) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


