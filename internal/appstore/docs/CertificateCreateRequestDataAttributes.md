# CertificateCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrContent** | **string** |  | 
**CertificateType** | [**CertificateType**](CertificateType.md) |  | 

## Methods

### NewCertificateCreateRequestDataAttributes

`func NewCertificateCreateRequestDataAttributes(csrContent string, certificateType CertificateType, ) *CertificateCreateRequestDataAttributes`

NewCertificateCreateRequestDataAttributes instantiates a new CertificateCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCreateRequestDataAttributesWithDefaults

`func NewCertificateCreateRequestDataAttributesWithDefaults() *CertificateCreateRequestDataAttributes`

NewCertificateCreateRequestDataAttributesWithDefaults instantiates a new CertificateCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrContent

`func (o *CertificateCreateRequestDataAttributes) GetCsrContent() string`

GetCsrContent returns the CsrContent field if non-nil, zero value otherwise.

### GetCsrContentOk

`func (o *CertificateCreateRequestDataAttributes) GetCsrContentOk() (*string, bool)`

GetCsrContentOk returns a tuple with the CsrContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrContent

`func (o *CertificateCreateRequestDataAttributes) SetCsrContent(v string)`

SetCsrContent sets CsrContent field to given value.


### GetCertificateType

`func (o *CertificateCreateRequestDataAttributes) GetCertificateType() CertificateType`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *CertificateCreateRequestDataAttributes) GetCertificateTypeOk() (*CertificateType, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *CertificateCreateRequestDataAttributes) SetCertificateType(v CertificateType)`

SetCertificateType sets CertificateType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


