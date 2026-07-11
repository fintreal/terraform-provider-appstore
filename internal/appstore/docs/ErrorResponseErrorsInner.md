# ErrorResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Code** | **string** |  | 
**Title** | **string** |  | 
**Detail** | **string** |  | 
**Source** | Pointer to [**ErrorResponseErrorsInnerSource**](ErrorResponseErrorsInnerSource.md) |  | [optional] 
**Links** | Pointer to [**ErrorLinks**](ErrorLinks.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewErrorResponseErrorsInner

`func NewErrorResponseErrorsInner(status string, code string, title string, detail string, ) *ErrorResponseErrorsInner`

NewErrorResponseErrorsInner instantiates a new ErrorResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseErrorsInnerWithDefaults

`func NewErrorResponseErrorsInnerWithDefaults() *ErrorResponseErrorsInner`

NewErrorResponseErrorsInnerWithDefaults instantiates a new ErrorResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorResponseErrorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorResponseErrorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorResponseErrorsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorResponseErrorsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorResponseErrorsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponseErrorsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponseErrorsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *ErrorResponseErrorsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseErrorsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseErrorsInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetTitle

`func (o *ErrorResponseErrorsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResponseErrorsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResponseErrorsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ErrorResponseErrorsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResponseErrorsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResponseErrorsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSource

`func (o *ErrorResponseErrorsInner) GetSource() ErrorResponseErrorsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ErrorResponseErrorsInner) GetSourceOk() (*ErrorResponseErrorsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ErrorResponseErrorsInner) SetSource(v ErrorResponseErrorsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ErrorResponseErrorsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetLinks

`func (o *ErrorResponseErrorsInner) GetLinks() ErrorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ErrorResponseErrorsInner) GetLinksOk() (*ErrorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ErrorResponseErrorsInner) SetLinks(v ErrorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ErrorResponseErrorsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorResponseErrorsInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorResponseErrorsInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorResponseErrorsInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorResponseErrorsInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


