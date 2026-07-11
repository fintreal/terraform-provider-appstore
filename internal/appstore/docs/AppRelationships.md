# AppRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEncryptionDeclarations** | Pointer to [**AppRelationshipsAppEncryptionDeclarations**](AppRelationshipsAppEncryptionDeclarations.md) |  | [optional] 
**CiProduct** | Pointer to [**AppRelationshipsCiProduct**](AppRelationshipsCiProduct.md) |  | [optional] 
**BetaTesters** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**BetaGroups** | Pointer to [**AppRelationshipsBetaGroups**](AppRelationshipsBetaGroups.md) |  | [optional] 
**AppStoreVersions** | Pointer to [**AppRelationshipsAppStoreVersions**](AppRelationshipsAppStoreVersions.md) |  | [optional] 
**PreReleaseVersions** | Pointer to [**AppRelationshipsPreReleaseVersions**](AppRelationshipsPreReleaseVersions.md) |  | [optional] 
**BetaAppLocalizations** | Pointer to [**AppRelationshipsBetaAppLocalizations**](AppRelationshipsBetaAppLocalizations.md) |  | [optional] 
**Builds** | Pointer to [**AppRelationshipsBuilds**](AppRelationshipsBuilds.md) |  | [optional] 
**BetaLicenseAgreement** | Pointer to [**AppRelationshipsBetaLicenseAgreement**](AppRelationshipsBetaLicenseAgreement.md) |  | [optional] 
**BetaAppReviewDetail** | Pointer to [**AppRelationshipsBetaAppReviewDetail**](AppRelationshipsBetaAppReviewDetail.md) |  | [optional] 
**AppInfos** | Pointer to [**AppRelationshipsAppInfos**](AppRelationshipsAppInfos.md) |  | [optional] 
**AppClips** | Pointer to [**AppRelationshipsAppClips**](AppRelationshipsAppClips.md) |  | [optional] 
**AppPricePoints** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**EndUserLicenseAgreement** | Pointer to [**AppRelationshipsEndUserLicenseAgreement**](AppRelationshipsEndUserLicenseAgreement.md) |  | [optional] 
**AppPriceSchedule** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**AppAvailabilityV2** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**InAppPurchases** | Pointer to [**AppRelationshipsInAppPurchases**](AppRelationshipsInAppPurchases.md) |  | [optional] 
**SubscriptionGroups** | Pointer to [**AppRelationshipsSubscriptionGroups**](AppRelationshipsSubscriptionGroups.md) |  | [optional] 
**GameCenterEnabledVersions** | Pointer to [**AppRelationshipsGameCenterEnabledVersions**](AppRelationshipsGameCenterEnabledVersions.md) |  | [optional] 
**PerfPowerMetrics** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**AppCustomProductPages** | Pointer to [**AppRelationshipsAppCustomProductPages**](AppRelationshipsAppCustomProductPages.md) |  | [optional] 
**InAppPurchasesV2** | Pointer to [**AppRelationshipsInAppPurchasesV2**](AppRelationshipsInAppPurchasesV2.md) |  | [optional] 
**PromotedPurchases** | Pointer to [**AppRelationshipsPromotedPurchases**](AppRelationshipsPromotedPurchases.md) |  | [optional] 
**AppEvents** | Pointer to [**AppRelationshipsAppEvents**](AppRelationshipsAppEvents.md) |  | [optional] 
**ReviewSubmissions** | Pointer to [**AppRelationshipsReviewSubmissions**](AppRelationshipsReviewSubmissions.md) |  | [optional] 
**SubscriptionGracePeriod** | Pointer to [**AppRelationshipsSubscriptionGracePeriod**](AppRelationshipsSubscriptionGracePeriod.md) |  | [optional] 
**CustomerReviews** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**GameCenterDetail** | Pointer to [**AppRelationshipsGameCenterDetail**](AppRelationshipsGameCenterDetail.md) |  | [optional] 
**AppStoreVersionExperimentsV2** | Pointer to [**AppRelationshipsAppStoreVersionExperimentsV2**](AppRelationshipsAppStoreVersionExperimentsV2.md) |  | [optional] 
**AlternativeDistributionKey** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**AnalyticsReportRequests** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 
**MarketplaceSearchDetail** | Pointer to [**AppRelationshipsBetaTesters**](AppRelationshipsBetaTesters.md) |  | [optional] 

## Methods

### NewAppRelationships

`func NewAppRelationships() *AppRelationships`

NewAppRelationships instantiates a new AppRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRelationshipsWithDefaults

`func NewAppRelationshipsWithDefaults() *AppRelationships`

NewAppRelationshipsWithDefaults instantiates a new AppRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEncryptionDeclarations

`func (o *AppRelationships) GetAppEncryptionDeclarations() AppRelationshipsAppEncryptionDeclarations`

GetAppEncryptionDeclarations returns the AppEncryptionDeclarations field if non-nil, zero value otherwise.

### GetAppEncryptionDeclarationsOk

`func (o *AppRelationships) GetAppEncryptionDeclarationsOk() (*AppRelationshipsAppEncryptionDeclarations, bool)`

GetAppEncryptionDeclarationsOk returns a tuple with the AppEncryptionDeclarations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEncryptionDeclarations

`func (o *AppRelationships) SetAppEncryptionDeclarations(v AppRelationshipsAppEncryptionDeclarations)`

SetAppEncryptionDeclarations sets AppEncryptionDeclarations field to given value.

### HasAppEncryptionDeclarations

`func (o *AppRelationships) HasAppEncryptionDeclarations() bool`

HasAppEncryptionDeclarations returns a boolean if a field has been set.

### GetCiProduct

`func (o *AppRelationships) GetCiProduct() AppRelationshipsCiProduct`

GetCiProduct returns the CiProduct field if non-nil, zero value otherwise.

### GetCiProductOk

`func (o *AppRelationships) GetCiProductOk() (*AppRelationshipsCiProduct, bool)`

GetCiProductOk returns a tuple with the CiProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiProduct

`func (o *AppRelationships) SetCiProduct(v AppRelationshipsCiProduct)`

SetCiProduct sets CiProduct field to given value.

### HasCiProduct

`func (o *AppRelationships) HasCiProduct() bool`

HasCiProduct returns a boolean if a field has been set.

### GetBetaTesters

`func (o *AppRelationships) GetBetaTesters() AppRelationshipsBetaTesters`

GetBetaTesters returns the BetaTesters field if non-nil, zero value otherwise.

### GetBetaTestersOk

`func (o *AppRelationships) GetBetaTestersOk() (*AppRelationshipsBetaTesters, bool)`

GetBetaTestersOk returns a tuple with the BetaTesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaTesters

`func (o *AppRelationships) SetBetaTesters(v AppRelationshipsBetaTesters)`

SetBetaTesters sets BetaTesters field to given value.

### HasBetaTesters

`func (o *AppRelationships) HasBetaTesters() bool`

HasBetaTesters returns a boolean if a field has been set.

### GetBetaGroups

`func (o *AppRelationships) GetBetaGroups() AppRelationshipsBetaGroups`

GetBetaGroups returns the BetaGroups field if non-nil, zero value otherwise.

### GetBetaGroupsOk

`func (o *AppRelationships) GetBetaGroupsOk() (*AppRelationshipsBetaGroups, bool)`

GetBetaGroupsOk returns a tuple with the BetaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaGroups

`func (o *AppRelationships) SetBetaGroups(v AppRelationshipsBetaGroups)`

SetBetaGroups sets BetaGroups field to given value.

### HasBetaGroups

`func (o *AppRelationships) HasBetaGroups() bool`

HasBetaGroups returns a boolean if a field has been set.

### GetAppStoreVersions

`func (o *AppRelationships) GetAppStoreVersions() AppRelationshipsAppStoreVersions`

GetAppStoreVersions returns the AppStoreVersions field if non-nil, zero value otherwise.

### GetAppStoreVersionsOk

`func (o *AppRelationships) GetAppStoreVersionsOk() (*AppRelationshipsAppStoreVersions, bool)`

GetAppStoreVersionsOk returns a tuple with the AppStoreVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersions

`func (o *AppRelationships) SetAppStoreVersions(v AppRelationshipsAppStoreVersions)`

SetAppStoreVersions sets AppStoreVersions field to given value.

### HasAppStoreVersions

`func (o *AppRelationships) HasAppStoreVersions() bool`

HasAppStoreVersions returns a boolean if a field has been set.

### GetPreReleaseVersions

`func (o *AppRelationships) GetPreReleaseVersions() AppRelationshipsPreReleaseVersions`

GetPreReleaseVersions returns the PreReleaseVersions field if non-nil, zero value otherwise.

### GetPreReleaseVersionsOk

`func (o *AppRelationships) GetPreReleaseVersionsOk() (*AppRelationshipsPreReleaseVersions, bool)`

GetPreReleaseVersionsOk returns a tuple with the PreReleaseVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreReleaseVersions

`func (o *AppRelationships) SetPreReleaseVersions(v AppRelationshipsPreReleaseVersions)`

SetPreReleaseVersions sets PreReleaseVersions field to given value.

### HasPreReleaseVersions

`func (o *AppRelationships) HasPreReleaseVersions() bool`

HasPreReleaseVersions returns a boolean if a field has been set.

### GetBetaAppLocalizations

`func (o *AppRelationships) GetBetaAppLocalizations() AppRelationshipsBetaAppLocalizations`

GetBetaAppLocalizations returns the BetaAppLocalizations field if non-nil, zero value otherwise.

### GetBetaAppLocalizationsOk

`func (o *AppRelationships) GetBetaAppLocalizationsOk() (*AppRelationshipsBetaAppLocalizations, bool)`

GetBetaAppLocalizationsOk returns a tuple with the BetaAppLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaAppLocalizations

`func (o *AppRelationships) SetBetaAppLocalizations(v AppRelationshipsBetaAppLocalizations)`

SetBetaAppLocalizations sets BetaAppLocalizations field to given value.

### HasBetaAppLocalizations

`func (o *AppRelationships) HasBetaAppLocalizations() bool`

HasBetaAppLocalizations returns a boolean if a field has been set.

### GetBuilds

`func (o *AppRelationships) GetBuilds() AppRelationshipsBuilds`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *AppRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *AppRelationships) SetBuilds(v AppRelationshipsBuilds)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *AppRelationships) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetBetaLicenseAgreement

`func (o *AppRelationships) GetBetaLicenseAgreement() AppRelationshipsBetaLicenseAgreement`

GetBetaLicenseAgreement returns the BetaLicenseAgreement field if non-nil, zero value otherwise.

### GetBetaLicenseAgreementOk

`func (o *AppRelationships) GetBetaLicenseAgreementOk() (*AppRelationshipsBetaLicenseAgreement, bool)`

GetBetaLicenseAgreementOk returns a tuple with the BetaLicenseAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaLicenseAgreement

`func (o *AppRelationships) SetBetaLicenseAgreement(v AppRelationshipsBetaLicenseAgreement)`

SetBetaLicenseAgreement sets BetaLicenseAgreement field to given value.

### HasBetaLicenseAgreement

`func (o *AppRelationships) HasBetaLicenseAgreement() bool`

HasBetaLicenseAgreement returns a boolean if a field has been set.

### GetBetaAppReviewDetail

`func (o *AppRelationships) GetBetaAppReviewDetail() AppRelationshipsBetaAppReviewDetail`

GetBetaAppReviewDetail returns the BetaAppReviewDetail field if non-nil, zero value otherwise.

### GetBetaAppReviewDetailOk

`func (o *AppRelationships) GetBetaAppReviewDetailOk() (*AppRelationshipsBetaAppReviewDetail, bool)`

GetBetaAppReviewDetailOk returns a tuple with the BetaAppReviewDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaAppReviewDetail

`func (o *AppRelationships) SetBetaAppReviewDetail(v AppRelationshipsBetaAppReviewDetail)`

SetBetaAppReviewDetail sets BetaAppReviewDetail field to given value.

### HasBetaAppReviewDetail

`func (o *AppRelationships) HasBetaAppReviewDetail() bool`

HasBetaAppReviewDetail returns a boolean if a field has been set.

### GetAppInfos

`func (o *AppRelationships) GetAppInfos() AppRelationshipsAppInfos`

GetAppInfos returns the AppInfos field if non-nil, zero value otherwise.

### GetAppInfosOk

`func (o *AppRelationships) GetAppInfosOk() (*AppRelationshipsAppInfos, bool)`

GetAppInfosOk returns a tuple with the AppInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfos

`func (o *AppRelationships) SetAppInfos(v AppRelationshipsAppInfos)`

SetAppInfos sets AppInfos field to given value.

### HasAppInfos

`func (o *AppRelationships) HasAppInfos() bool`

HasAppInfos returns a boolean if a field has been set.

### GetAppClips

`func (o *AppRelationships) GetAppClips() AppRelationshipsAppClips`

GetAppClips returns the AppClips field if non-nil, zero value otherwise.

### GetAppClipsOk

`func (o *AppRelationships) GetAppClipsOk() (*AppRelationshipsAppClips, bool)`

GetAppClipsOk returns a tuple with the AppClips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClips

`func (o *AppRelationships) SetAppClips(v AppRelationshipsAppClips)`

SetAppClips sets AppClips field to given value.

### HasAppClips

`func (o *AppRelationships) HasAppClips() bool`

HasAppClips returns a boolean if a field has been set.

### GetAppPricePoints

`func (o *AppRelationships) GetAppPricePoints() AppRelationshipsBetaTesters`

GetAppPricePoints returns the AppPricePoints field if non-nil, zero value otherwise.

### GetAppPricePointsOk

`func (o *AppRelationships) GetAppPricePointsOk() (*AppRelationshipsBetaTesters, bool)`

GetAppPricePointsOk returns a tuple with the AppPricePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPricePoints

`func (o *AppRelationships) SetAppPricePoints(v AppRelationshipsBetaTesters)`

SetAppPricePoints sets AppPricePoints field to given value.

### HasAppPricePoints

`func (o *AppRelationships) HasAppPricePoints() bool`

HasAppPricePoints returns a boolean if a field has been set.

### GetEndUserLicenseAgreement

`func (o *AppRelationships) GetEndUserLicenseAgreement() AppRelationshipsEndUserLicenseAgreement`

GetEndUserLicenseAgreement returns the EndUserLicenseAgreement field if non-nil, zero value otherwise.

### GetEndUserLicenseAgreementOk

`func (o *AppRelationships) GetEndUserLicenseAgreementOk() (*AppRelationshipsEndUserLicenseAgreement, bool)`

GetEndUserLicenseAgreementOk returns a tuple with the EndUserLicenseAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserLicenseAgreement

`func (o *AppRelationships) SetEndUserLicenseAgreement(v AppRelationshipsEndUserLicenseAgreement)`

SetEndUserLicenseAgreement sets EndUserLicenseAgreement field to given value.

### HasEndUserLicenseAgreement

`func (o *AppRelationships) HasEndUserLicenseAgreement() bool`

HasEndUserLicenseAgreement returns a boolean if a field has been set.

### GetAppPriceSchedule

`func (o *AppRelationships) GetAppPriceSchedule() AppRelationshipsBetaTesters`

GetAppPriceSchedule returns the AppPriceSchedule field if non-nil, zero value otherwise.

### GetAppPriceScheduleOk

`func (o *AppRelationships) GetAppPriceScheduleOk() (*AppRelationshipsBetaTesters, bool)`

GetAppPriceScheduleOk returns a tuple with the AppPriceSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPriceSchedule

`func (o *AppRelationships) SetAppPriceSchedule(v AppRelationshipsBetaTesters)`

SetAppPriceSchedule sets AppPriceSchedule field to given value.

### HasAppPriceSchedule

`func (o *AppRelationships) HasAppPriceSchedule() bool`

HasAppPriceSchedule returns a boolean if a field has been set.

### GetAppAvailabilityV2

`func (o *AppRelationships) GetAppAvailabilityV2() AppRelationshipsBetaTesters`

GetAppAvailabilityV2 returns the AppAvailabilityV2 field if non-nil, zero value otherwise.

### GetAppAvailabilityV2Ok

`func (o *AppRelationships) GetAppAvailabilityV2Ok() (*AppRelationshipsBetaTesters, bool)`

GetAppAvailabilityV2Ok returns a tuple with the AppAvailabilityV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAvailabilityV2

`func (o *AppRelationships) SetAppAvailabilityV2(v AppRelationshipsBetaTesters)`

SetAppAvailabilityV2 sets AppAvailabilityV2 field to given value.

### HasAppAvailabilityV2

`func (o *AppRelationships) HasAppAvailabilityV2() bool`

HasAppAvailabilityV2 returns a boolean if a field has been set.

### GetInAppPurchases

`func (o *AppRelationships) GetInAppPurchases() AppRelationshipsInAppPurchases`

GetInAppPurchases returns the InAppPurchases field if non-nil, zero value otherwise.

### GetInAppPurchasesOk

`func (o *AppRelationships) GetInAppPurchasesOk() (*AppRelationshipsInAppPurchases, bool)`

GetInAppPurchasesOk returns a tuple with the InAppPurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAppPurchases

`func (o *AppRelationships) SetInAppPurchases(v AppRelationshipsInAppPurchases)`

SetInAppPurchases sets InAppPurchases field to given value.

### HasInAppPurchases

`func (o *AppRelationships) HasInAppPurchases() bool`

HasInAppPurchases returns a boolean if a field has been set.

### GetSubscriptionGroups

`func (o *AppRelationships) GetSubscriptionGroups() AppRelationshipsSubscriptionGroups`

GetSubscriptionGroups returns the SubscriptionGroups field if non-nil, zero value otherwise.

### GetSubscriptionGroupsOk

`func (o *AppRelationships) GetSubscriptionGroupsOk() (*AppRelationshipsSubscriptionGroups, bool)`

GetSubscriptionGroupsOk returns a tuple with the SubscriptionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionGroups

`func (o *AppRelationships) SetSubscriptionGroups(v AppRelationshipsSubscriptionGroups)`

SetSubscriptionGroups sets SubscriptionGroups field to given value.

### HasSubscriptionGroups

`func (o *AppRelationships) HasSubscriptionGroups() bool`

HasSubscriptionGroups returns a boolean if a field has been set.

### GetGameCenterEnabledVersions

`func (o *AppRelationships) GetGameCenterEnabledVersions() AppRelationshipsGameCenterEnabledVersions`

GetGameCenterEnabledVersions returns the GameCenterEnabledVersions field if non-nil, zero value otherwise.

### GetGameCenterEnabledVersionsOk

`func (o *AppRelationships) GetGameCenterEnabledVersionsOk() (*AppRelationshipsGameCenterEnabledVersions, bool)`

GetGameCenterEnabledVersionsOk returns a tuple with the GameCenterEnabledVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterEnabledVersions

`func (o *AppRelationships) SetGameCenterEnabledVersions(v AppRelationshipsGameCenterEnabledVersions)`

SetGameCenterEnabledVersions sets GameCenterEnabledVersions field to given value.

### HasGameCenterEnabledVersions

`func (o *AppRelationships) HasGameCenterEnabledVersions() bool`

HasGameCenterEnabledVersions returns a boolean if a field has been set.

### GetPerfPowerMetrics

`func (o *AppRelationships) GetPerfPowerMetrics() AppRelationshipsBetaTesters`

GetPerfPowerMetrics returns the PerfPowerMetrics field if non-nil, zero value otherwise.

### GetPerfPowerMetricsOk

`func (o *AppRelationships) GetPerfPowerMetricsOk() (*AppRelationshipsBetaTesters, bool)`

GetPerfPowerMetricsOk returns a tuple with the PerfPowerMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfPowerMetrics

`func (o *AppRelationships) SetPerfPowerMetrics(v AppRelationshipsBetaTesters)`

SetPerfPowerMetrics sets PerfPowerMetrics field to given value.

### HasPerfPowerMetrics

`func (o *AppRelationships) HasPerfPowerMetrics() bool`

HasPerfPowerMetrics returns a boolean if a field has been set.

### GetAppCustomProductPages

`func (o *AppRelationships) GetAppCustomProductPages() AppRelationshipsAppCustomProductPages`

GetAppCustomProductPages returns the AppCustomProductPages field if non-nil, zero value otherwise.

### GetAppCustomProductPagesOk

`func (o *AppRelationships) GetAppCustomProductPagesOk() (*AppRelationshipsAppCustomProductPages, bool)`

GetAppCustomProductPagesOk returns a tuple with the AppCustomProductPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCustomProductPages

`func (o *AppRelationships) SetAppCustomProductPages(v AppRelationshipsAppCustomProductPages)`

SetAppCustomProductPages sets AppCustomProductPages field to given value.

### HasAppCustomProductPages

`func (o *AppRelationships) HasAppCustomProductPages() bool`

HasAppCustomProductPages returns a boolean if a field has been set.

### GetInAppPurchasesV2

`func (o *AppRelationships) GetInAppPurchasesV2() AppRelationshipsInAppPurchasesV2`

GetInAppPurchasesV2 returns the InAppPurchasesV2 field if non-nil, zero value otherwise.

### GetInAppPurchasesV2Ok

`func (o *AppRelationships) GetInAppPurchasesV2Ok() (*AppRelationshipsInAppPurchasesV2, bool)`

GetInAppPurchasesV2Ok returns a tuple with the InAppPurchasesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAppPurchasesV2

`func (o *AppRelationships) SetInAppPurchasesV2(v AppRelationshipsInAppPurchasesV2)`

SetInAppPurchasesV2 sets InAppPurchasesV2 field to given value.

### HasInAppPurchasesV2

`func (o *AppRelationships) HasInAppPurchasesV2() bool`

HasInAppPurchasesV2 returns a boolean if a field has been set.

### GetPromotedPurchases

`func (o *AppRelationships) GetPromotedPurchases() AppRelationshipsPromotedPurchases`

GetPromotedPurchases returns the PromotedPurchases field if non-nil, zero value otherwise.

### GetPromotedPurchasesOk

`func (o *AppRelationships) GetPromotedPurchasesOk() (*AppRelationshipsPromotedPurchases, bool)`

GetPromotedPurchasesOk returns a tuple with the PromotedPurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedPurchases

`func (o *AppRelationships) SetPromotedPurchases(v AppRelationshipsPromotedPurchases)`

SetPromotedPurchases sets PromotedPurchases field to given value.

### HasPromotedPurchases

`func (o *AppRelationships) HasPromotedPurchases() bool`

HasPromotedPurchases returns a boolean if a field has been set.

### GetAppEvents

`func (o *AppRelationships) GetAppEvents() AppRelationshipsAppEvents`

GetAppEvents returns the AppEvents field if non-nil, zero value otherwise.

### GetAppEventsOk

`func (o *AppRelationships) GetAppEventsOk() (*AppRelationshipsAppEvents, bool)`

GetAppEventsOk returns a tuple with the AppEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEvents

`func (o *AppRelationships) SetAppEvents(v AppRelationshipsAppEvents)`

SetAppEvents sets AppEvents field to given value.

### HasAppEvents

`func (o *AppRelationships) HasAppEvents() bool`

HasAppEvents returns a boolean if a field has been set.

### GetReviewSubmissions

`func (o *AppRelationships) GetReviewSubmissions() AppRelationshipsReviewSubmissions`

GetReviewSubmissions returns the ReviewSubmissions field if non-nil, zero value otherwise.

### GetReviewSubmissionsOk

`func (o *AppRelationships) GetReviewSubmissionsOk() (*AppRelationshipsReviewSubmissions, bool)`

GetReviewSubmissionsOk returns a tuple with the ReviewSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewSubmissions

`func (o *AppRelationships) SetReviewSubmissions(v AppRelationshipsReviewSubmissions)`

SetReviewSubmissions sets ReviewSubmissions field to given value.

### HasReviewSubmissions

`func (o *AppRelationships) HasReviewSubmissions() bool`

HasReviewSubmissions returns a boolean if a field has been set.

### GetSubscriptionGracePeriod

`func (o *AppRelationships) GetSubscriptionGracePeriod() AppRelationshipsSubscriptionGracePeriod`

GetSubscriptionGracePeriod returns the SubscriptionGracePeriod field if non-nil, zero value otherwise.

### GetSubscriptionGracePeriodOk

`func (o *AppRelationships) GetSubscriptionGracePeriodOk() (*AppRelationshipsSubscriptionGracePeriod, bool)`

GetSubscriptionGracePeriodOk returns a tuple with the SubscriptionGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionGracePeriod

`func (o *AppRelationships) SetSubscriptionGracePeriod(v AppRelationshipsSubscriptionGracePeriod)`

SetSubscriptionGracePeriod sets SubscriptionGracePeriod field to given value.

### HasSubscriptionGracePeriod

`func (o *AppRelationships) HasSubscriptionGracePeriod() bool`

HasSubscriptionGracePeriod returns a boolean if a field has been set.

### GetCustomerReviews

`func (o *AppRelationships) GetCustomerReviews() AppRelationshipsBetaTesters`

GetCustomerReviews returns the CustomerReviews field if non-nil, zero value otherwise.

### GetCustomerReviewsOk

`func (o *AppRelationships) GetCustomerReviewsOk() (*AppRelationshipsBetaTesters, bool)`

GetCustomerReviewsOk returns a tuple with the CustomerReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReviews

`func (o *AppRelationships) SetCustomerReviews(v AppRelationshipsBetaTesters)`

SetCustomerReviews sets CustomerReviews field to given value.

### HasCustomerReviews

`func (o *AppRelationships) HasCustomerReviews() bool`

HasCustomerReviews returns a boolean if a field has been set.

### GetGameCenterDetail

`func (o *AppRelationships) GetGameCenterDetail() AppRelationshipsGameCenterDetail`

GetGameCenterDetail returns the GameCenterDetail field if non-nil, zero value otherwise.

### GetGameCenterDetailOk

`func (o *AppRelationships) GetGameCenterDetailOk() (*AppRelationshipsGameCenterDetail, bool)`

GetGameCenterDetailOk returns a tuple with the GameCenterDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterDetail

`func (o *AppRelationships) SetGameCenterDetail(v AppRelationshipsGameCenterDetail)`

SetGameCenterDetail sets GameCenterDetail field to given value.

### HasGameCenterDetail

`func (o *AppRelationships) HasGameCenterDetail() bool`

HasGameCenterDetail returns a boolean if a field has been set.

### GetAppStoreVersionExperimentsV2

`func (o *AppRelationships) GetAppStoreVersionExperimentsV2() AppRelationshipsAppStoreVersionExperimentsV2`

GetAppStoreVersionExperimentsV2 returns the AppStoreVersionExperimentsV2 field if non-nil, zero value otherwise.

### GetAppStoreVersionExperimentsV2Ok

`func (o *AppRelationships) GetAppStoreVersionExperimentsV2Ok() (*AppRelationshipsAppStoreVersionExperimentsV2, bool)`

GetAppStoreVersionExperimentsV2Ok returns a tuple with the AppStoreVersionExperimentsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersionExperimentsV2

`func (o *AppRelationships) SetAppStoreVersionExperimentsV2(v AppRelationshipsAppStoreVersionExperimentsV2)`

SetAppStoreVersionExperimentsV2 sets AppStoreVersionExperimentsV2 field to given value.

### HasAppStoreVersionExperimentsV2

`func (o *AppRelationships) HasAppStoreVersionExperimentsV2() bool`

HasAppStoreVersionExperimentsV2 returns a boolean if a field has been set.

### GetAlternativeDistributionKey

`func (o *AppRelationships) GetAlternativeDistributionKey() AppRelationshipsBetaTesters`

GetAlternativeDistributionKey returns the AlternativeDistributionKey field if non-nil, zero value otherwise.

### GetAlternativeDistributionKeyOk

`func (o *AppRelationships) GetAlternativeDistributionKeyOk() (*AppRelationshipsBetaTesters, bool)`

GetAlternativeDistributionKeyOk returns a tuple with the AlternativeDistributionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDistributionKey

`func (o *AppRelationships) SetAlternativeDistributionKey(v AppRelationshipsBetaTesters)`

SetAlternativeDistributionKey sets AlternativeDistributionKey field to given value.

### HasAlternativeDistributionKey

`func (o *AppRelationships) HasAlternativeDistributionKey() bool`

HasAlternativeDistributionKey returns a boolean if a field has been set.

### GetAnalyticsReportRequests

`func (o *AppRelationships) GetAnalyticsReportRequests() AppRelationshipsBetaTesters`

GetAnalyticsReportRequests returns the AnalyticsReportRequests field if non-nil, zero value otherwise.

### GetAnalyticsReportRequestsOk

`func (o *AppRelationships) GetAnalyticsReportRequestsOk() (*AppRelationshipsBetaTesters, bool)`

GetAnalyticsReportRequestsOk returns a tuple with the AnalyticsReportRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsReportRequests

`func (o *AppRelationships) SetAnalyticsReportRequests(v AppRelationshipsBetaTesters)`

SetAnalyticsReportRequests sets AnalyticsReportRequests field to given value.

### HasAnalyticsReportRequests

`func (o *AppRelationships) HasAnalyticsReportRequests() bool`

HasAnalyticsReportRequests returns a boolean if a field has been set.

### GetMarketplaceSearchDetail

`func (o *AppRelationships) GetMarketplaceSearchDetail() AppRelationshipsBetaTesters`

GetMarketplaceSearchDetail returns the MarketplaceSearchDetail field if non-nil, zero value otherwise.

### GetMarketplaceSearchDetailOk

`func (o *AppRelationships) GetMarketplaceSearchDetailOk() (*AppRelationshipsBetaTesters, bool)`

GetMarketplaceSearchDetailOk returns a tuple with the MarketplaceSearchDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSearchDetail

`func (o *AppRelationships) SetMarketplaceSearchDetail(v AppRelationshipsBetaTesters)`

SetMarketplaceSearchDetail sets MarketplaceSearchDetail field to given value.

### HasMarketplaceSearchDetail

`func (o *AppRelationships) HasMarketplaceSearchDetail() bool`

HasMarketplaceSearchDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


