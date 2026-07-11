# ErrorLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**ErrorLinksAssociated**](ErrorLinksAssociated.md) |  | [optional] 

## Methods

### NewErrorLinks

`func NewErrorLinks() *ErrorLinks`

NewErrorLinks instantiates a new ErrorLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLinksWithDefaults

`func NewErrorLinksWithDefaults() *ErrorLinks`

NewErrorLinksWithDefaults instantiates a new ErrorLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *ErrorLinks) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ErrorLinks) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ErrorLinks) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ErrorLinks) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetAssociated

`func (o *ErrorLinks) GetAssociated() ErrorLinksAssociated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *ErrorLinks) GetAssociatedOk() (*ErrorLinksAssociated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *ErrorLinks) SetAssociated(v ErrorLinksAssociated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *ErrorLinks) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


