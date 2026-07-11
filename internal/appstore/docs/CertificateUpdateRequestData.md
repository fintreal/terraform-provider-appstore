# CertificateUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**CertificateUpdateRequestDataAttributes**](CertificateUpdateRequestDataAttributes.md) |  | [optional] 

## Methods

### NewCertificateUpdateRequestData

`func NewCertificateUpdateRequestData(type_ string, id string, ) *CertificateUpdateRequestData`

NewCertificateUpdateRequestData instantiates a new CertificateUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateUpdateRequestDataWithDefaults

`func NewCertificateUpdateRequestDataWithDefaults() *CertificateUpdateRequestData`

NewCertificateUpdateRequestDataWithDefaults instantiates a new CertificateUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CertificateUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CertificateUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CertificateUpdateRequestData) GetAttributes() CertificateUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CertificateUpdateRequestData) GetAttributesOk() (*CertificateUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CertificateUpdateRequestData) SetAttributes(v CertificateUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CertificateUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


