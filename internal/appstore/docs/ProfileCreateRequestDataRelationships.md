# ProfileCreateRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | [**BundleIdCapabilityCreateRequestDataRelationshipsBundleId**](BundleIdCapabilityCreateRequestDataRelationshipsBundleId.md) |  | 
**Devices** | Pointer to [**ProfileRelationshipsDevices**](ProfileRelationshipsDevices.md) |  | [optional] 
**Certificates** | [**ProfileCreateRequestDataRelationshipsCertificates**](ProfileCreateRequestDataRelationshipsCertificates.md) |  | 

## Methods

### NewProfileCreateRequestDataRelationships

`func NewProfileCreateRequestDataRelationships(bundleId BundleIdCapabilityCreateRequestDataRelationshipsBundleId, certificates ProfileCreateRequestDataRelationshipsCertificates, ) *ProfileCreateRequestDataRelationships`

NewProfileCreateRequestDataRelationships instantiates a new ProfileCreateRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileCreateRequestDataRelationshipsWithDefaults

`func NewProfileCreateRequestDataRelationshipsWithDefaults() *ProfileCreateRequestDataRelationships`

NewProfileCreateRequestDataRelationshipsWithDefaults instantiates a new ProfileCreateRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *ProfileCreateRequestDataRelationships) GetBundleId() BundleIdCapabilityCreateRequestDataRelationshipsBundleId`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ProfileCreateRequestDataRelationships) GetBundleIdOk() (*BundleIdCapabilityCreateRequestDataRelationshipsBundleId, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ProfileCreateRequestDataRelationships) SetBundleId(v BundleIdCapabilityCreateRequestDataRelationshipsBundleId)`

SetBundleId sets BundleId field to given value.


### GetDevices

`func (o *ProfileCreateRequestDataRelationships) GetDevices() ProfileRelationshipsDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ProfileCreateRequestDataRelationships) GetDevicesOk() (*ProfileRelationshipsDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ProfileCreateRequestDataRelationships) SetDevices(v ProfileRelationshipsDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *ProfileCreateRequestDataRelationships) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetCertificates

`func (o *ProfileCreateRequestDataRelationships) GetCertificates() ProfileCreateRequestDataRelationshipsCertificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ProfileCreateRequestDataRelationships) GetCertificatesOk() (*ProfileCreateRequestDataRelationshipsCertificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ProfileCreateRequestDataRelationships) SetCertificates(v ProfileCreateRequestDataRelationshipsCertificates)`

SetCertificates sets Certificates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


