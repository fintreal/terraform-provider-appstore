# CertificateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CertificateType** | [**CertificateType**](CertificateType.md) |  | 
**DisplayName** | **string** |  | 
**SerialNumber** | **string** |  | 
**Platform** | Pointer to [**BundleIdPlatform**](BundleIdPlatform.md) |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**CertificateContent** | Pointer to **string** |  | [optional] 
**Activated** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificateAttributes

`func NewCertificateAttributes(name string, certificateType CertificateType, displayName string, serialNumber string, ) *CertificateAttributes`

NewCertificateAttributes instantiates a new CertificateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAttributesWithDefaults

`func NewCertificateAttributesWithDefaults() *CertificateAttributes`

NewCertificateAttributesWithDefaults instantiates a new CertificateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCertificateType

`func (o *CertificateAttributes) GetCertificateType() CertificateType`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *CertificateAttributes) GetCertificateTypeOk() (*CertificateType, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *CertificateAttributes) SetCertificateType(v CertificateType)`

SetCertificateType sets CertificateType field to given value.


### GetDisplayName

`func (o *CertificateAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSerialNumber

`func (o *CertificateAttributes) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateAttributes) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateAttributes) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetPlatform

`func (o *CertificateAttributes) GetPlatform() BundleIdPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CertificateAttributes) GetPlatformOk() (*BundleIdPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CertificateAttributes) SetPlatform(v BundleIdPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CertificateAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CertificateAttributes) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CertificateAttributes) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CertificateAttributes) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CertificateAttributes) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetCertificateContent

`func (o *CertificateAttributes) GetCertificateContent() string`

GetCertificateContent returns the CertificateContent field if non-nil, zero value otherwise.

### GetCertificateContentOk

`func (o *CertificateAttributes) GetCertificateContentOk() (*string, bool)`

GetCertificateContentOk returns a tuple with the CertificateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateContent

`func (o *CertificateAttributes) SetCertificateContent(v string)`

SetCertificateContent sets CertificateContent field to given value.

### HasCertificateContent

`func (o *CertificateAttributes) HasCertificateContent() bool`

HasCertificateContent returns a boolean if a field has been set.

### GetActivated

`func (o *CertificateAttributes) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *CertificateAttributes) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *CertificateAttributes) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *CertificateAttributes) HasActivated() bool`

HasActivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


