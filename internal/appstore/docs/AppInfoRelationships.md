# AppInfoRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AppInfoRelationshipsApp**](AppInfoRelationshipsApp.md) |  | [optional] 
**AgeRatingDeclaration** | Pointer to [**AppInfoRelationshipsAgeRatingDeclaration**](AppInfoRelationshipsAgeRatingDeclaration.md) |  | [optional] 
**AppInfoLocalizations** | Pointer to [**AppInfoRelationshipsAppInfoLocalizations**](AppInfoRelationshipsAppInfoLocalizations.md) |  | [optional] 
**PrimaryCategory** | Pointer to [**AppInfoRelationshipsPrimaryCategory**](AppInfoRelationshipsPrimaryCategory.md) |  | [optional] 
**PrimarySubcategoryOne** | Pointer to [**AppInfoRelationshipsPrimaryCategory**](AppInfoRelationshipsPrimaryCategory.md) |  | [optional] 
**PrimarySubcategoryTwo** | Pointer to [**AppInfoRelationshipsPrimaryCategory**](AppInfoRelationshipsPrimaryCategory.md) |  | [optional] 
**SecondaryCategory** | Pointer to [**AppInfoRelationshipsPrimaryCategory**](AppInfoRelationshipsPrimaryCategory.md) |  | [optional] 
**SecondarySubcategoryOne** | Pointer to [**AppInfoRelationshipsPrimaryCategory**](AppInfoRelationshipsPrimaryCategory.md) |  | [optional] 
**SecondarySubcategoryTwo** | Pointer to [**AppInfoRelationshipsPrimaryCategory**](AppInfoRelationshipsPrimaryCategory.md) |  | [optional] 

## Methods

### NewAppInfoRelationships

`func NewAppInfoRelationships() *AppInfoRelationships`

NewAppInfoRelationships instantiates a new AppInfoRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoRelationshipsWithDefaults

`func NewAppInfoRelationshipsWithDefaults() *AppInfoRelationships`

NewAppInfoRelationshipsWithDefaults instantiates a new AppInfoRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AppInfoRelationships) GetApp() AppInfoRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AppInfoRelationships) GetAppOk() (*AppInfoRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AppInfoRelationships) SetApp(v AppInfoRelationshipsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AppInfoRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAgeRatingDeclaration

`func (o *AppInfoRelationships) GetAgeRatingDeclaration() AppInfoRelationshipsAgeRatingDeclaration`

GetAgeRatingDeclaration returns the AgeRatingDeclaration field if non-nil, zero value otherwise.

### GetAgeRatingDeclarationOk

`func (o *AppInfoRelationships) GetAgeRatingDeclarationOk() (*AppInfoRelationshipsAgeRatingDeclaration, bool)`

GetAgeRatingDeclarationOk returns a tuple with the AgeRatingDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeRatingDeclaration

`func (o *AppInfoRelationships) SetAgeRatingDeclaration(v AppInfoRelationshipsAgeRatingDeclaration)`

SetAgeRatingDeclaration sets AgeRatingDeclaration field to given value.

### HasAgeRatingDeclaration

`func (o *AppInfoRelationships) HasAgeRatingDeclaration() bool`

HasAgeRatingDeclaration returns a boolean if a field has been set.

### GetAppInfoLocalizations

`func (o *AppInfoRelationships) GetAppInfoLocalizations() AppInfoRelationshipsAppInfoLocalizations`

GetAppInfoLocalizations returns the AppInfoLocalizations field if non-nil, zero value otherwise.

### GetAppInfoLocalizationsOk

`func (o *AppInfoRelationships) GetAppInfoLocalizationsOk() (*AppInfoRelationshipsAppInfoLocalizations, bool)`

GetAppInfoLocalizationsOk returns a tuple with the AppInfoLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfoLocalizations

`func (o *AppInfoRelationships) SetAppInfoLocalizations(v AppInfoRelationshipsAppInfoLocalizations)`

SetAppInfoLocalizations sets AppInfoLocalizations field to given value.

### HasAppInfoLocalizations

`func (o *AppInfoRelationships) HasAppInfoLocalizations() bool`

HasAppInfoLocalizations returns a boolean if a field has been set.

### GetPrimaryCategory

`func (o *AppInfoRelationships) GetPrimaryCategory() AppInfoRelationshipsPrimaryCategory`

GetPrimaryCategory returns the PrimaryCategory field if non-nil, zero value otherwise.

### GetPrimaryCategoryOk

`func (o *AppInfoRelationships) GetPrimaryCategoryOk() (*AppInfoRelationshipsPrimaryCategory, bool)`

GetPrimaryCategoryOk returns a tuple with the PrimaryCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCategory

`func (o *AppInfoRelationships) SetPrimaryCategory(v AppInfoRelationshipsPrimaryCategory)`

SetPrimaryCategory sets PrimaryCategory field to given value.

### HasPrimaryCategory

`func (o *AppInfoRelationships) HasPrimaryCategory() bool`

HasPrimaryCategory returns a boolean if a field has been set.

### GetPrimarySubcategoryOne

`func (o *AppInfoRelationships) GetPrimarySubcategoryOne() AppInfoRelationshipsPrimaryCategory`

GetPrimarySubcategoryOne returns the PrimarySubcategoryOne field if non-nil, zero value otherwise.

### GetPrimarySubcategoryOneOk

`func (o *AppInfoRelationships) GetPrimarySubcategoryOneOk() (*AppInfoRelationshipsPrimaryCategory, bool)`

GetPrimarySubcategoryOneOk returns a tuple with the PrimarySubcategoryOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubcategoryOne

`func (o *AppInfoRelationships) SetPrimarySubcategoryOne(v AppInfoRelationshipsPrimaryCategory)`

SetPrimarySubcategoryOne sets PrimarySubcategoryOne field to given value.

### HasPrimarySubcategoryOne

`func (o *AppInfoRelationships) HasPrimarySubcategoryOne() bool`

HasPrimarySubcategoryOne returns a boolean if a field has been set.

### GetPrimarySubcategoryTwo

`func (o *AppInfoRelationships) GetPrimarySubcategoryTwo() AppInfoRelationshipsPrimaryCategory`

GetPrimarySubcategoryTwo returns the PrimarySubcategoryTwo field if non-nil, zero value otherwise.

### GetPrimarySubcategoryTwoOk

`func (o *AppInfoRelationships) GetPrimarySubcategoryTwoOk() (*AppInfoRelationshipsPrimaryCategory, bool)`

GetPrimarySubcategoryTwoOk returns a tuple with the PrimarySubcategoryTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubcategoryTwo

`func (o *AppInfoRelationships) SetPrimarySubcategoryTwo(v AppInfoRelationshipsPrimaryCategory)`

SetPrimarySubcategoryTwo sets PrimarySubcategoryTwo field to given value.

### HasPrimarySubcategoryTwo

`func (o *AppInfoRelationships) HasPrimarySubcategoryTwo() bool`

HasPrimarySubcategoryTwo returns a boolean if a field has been set.

### GetSecondaryCategory

`func (o *AppInfoRelationships) GetSecondaryCategory() AppInfoRelationshipsPrimaryCategory`

GetSecondaryCategory returns the SecondaryCategory field if non-nil, zero value otherwise.

### GetSecondaryCategoryOk

`func (o *AppInfoRelationships) GetSecondaryCategoryOk() (*AppInfoRelationshipsPrimaryCategory, bool)`

GetSecondaryCategoryOk returns a tuple with the SecondaryCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCategory

`func (o *AppInfoRelationships) SetSecondaryCategory(v AppInfoRelationshipsPrimaryCategory)`

SetSecondaryCategory sets SecondaryCategory field to given value.

### HasSecondaryCategory

`func (o *AppInfoRelationships) HasSecondaryCategory() bool`

HasSecondaryCategory returns a boolean if a field has been set.

### GetSecondarySubcategoryOne

`func (o *AppInfoRelationships) GetSecondarySubcategoryOne() AppInfoRelationshipsPrimaryCategory`

GetSecondarySubcategoryOne returns the SecondarySubcategoryOne field if non-nil, zero value otherwise.

### GetSecondarySubcategoryOneOk

`func (o *AppInfoRelationships) GetSecondarySubcategoryOneOk() (*AppInfoRelationshipsPrimaryCategory, bool)`

GetSecondarySubcategoryOneOk returns a tuple with the SecondarySubcategoryOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubcategoryOne

`func (o *AppInfoRelationships) SetSecondarySubcategoryOne(v AppInfoRelationshipsPrimaryCategory)`

SetSecondarySubcategoryOne sets SecondarySubcategoryOne field to given value.

### HasSecondarySubcategoryOne

`func (o *AppInfoRelationships) HasSecondarySubcategoryOne() bool`

HasSecondarySubcategoryOne returns a boolean if a field has been set.

### GetSecondarySubcategoryTwo

`func (o *AppInfoRelationships) GetSecondarySubcategoryTwo() AppInfoRelationshipsPrimaryCategory`

GetSecondarySubcategoryTwo returns the SecondarySubcategoryTwo field if non-nil, zero value otherwise.

### GetSecondarySubcategoryTwoOk

`func (o *AppInfoRelationships) GetSecondarySubcategoryTwoOk() (*AppInfoRelationshipsPrimaryCategory, bool)`

GetSecondarySubcategoryTwoOk returns a tuple with the SecondarySubcategoryTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubcategoryTwo

`func (o *AppInfoRelationships) SetSecondarySubcategoryTwo(v AppInfoRelationshipsPrimaryCategory)`

SetSecondarySubcategoryTwo sets SecondarySubcategoryTwo field to given value.

### HasSecondarySubcategoryTwo

`func (o *AppInfoRelationships) HasSecondarySubcategoryTwo() bool`

HasSecondarySubcategoryTwo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


