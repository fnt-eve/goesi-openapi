# CorporationsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** |  | [optional] 
**CeoId** | Pointer to **int64** |  | [optional] 
**CreatorId** | Pointer to **int64** |  | [optional] 
**DateFounded** | Pointer to **time.Time** | Corporation&#39;s founding date | [optional] 
**Description** | **string** | Corporation&#39;s description | 
**EnlistedFactionId** | Pointer to **int64** |  | [optional] 
**FriendlyFire** | **string** | Corporation&#39;s friendly fire status | 
**HomeStationId** | **int64** |  | 
**MemberCount** | **int64** | Corporation&#39;s member count | 
**Name** | **string** | Corporation&#39;s name | 
**Palette** | Pointer to [**CorporationsDetailPalette**](CorporationsDetailPalette.md) |  | [optional] 
**Shares** | **int64** | Corporation&#39;s shares | 
**State** | **string** | Corporation&#39;s state | 
**TaxRates** | [**CorporationsDetailTaxrates**](CorporationsDetailTaxrates.md) |  | 
**Ticker** | **string** | Corporation&#39;s short name | 
**Type** | **string** | Corporation&#39;s type | 
**Url** | Pointer to **string** | Corporation&#39;s URL | [optional] 
**WarEligible** | **bool** | Corporation&#39;s war eligible | 

## Methods

### NewCorporationsDetail

`func NewCorporationsDetail(description string, friendlyFire string, homeStationId int64, memberCount int64, name string, shares int64, state string, taxRates CorporationsDetailTaxrates, ticker string, type_ string, warEligible bool, ) *CorporationsDetail`

NewCorporationsDetail instantiates a new CorporationsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsDetailWithDefaults

`func NewCorporationsDetailWithDefaults() *CorporationsDetail`

NewCorporationsDetailWithDefaults instantiates a new CorporationsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *CorporationsDetail) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *CorporationsDetail) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *CorporationsDetail) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *CorporationsDetail) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetCeoId

`func (o *CorporationsDetail) GetCeoId() int64`

GetCeoId returns the CeoId field if non-nil, zero value otherwise.

### GetCeoIdOk

`func (o *CorporationsDetail) GetCeoIdOk() (*int64, bool)`

GetCeoIdOk returns a tuple with the CeoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeoId

`func (o *CorporationsDetail) SetCeoId(v int64)`

SetCeoId sets CeoId field to given value.

### HasCeoId

`func (o *CorporationsDetail) HasCeoId() bool`

HasCeoId returns a boolean if a field has been set.

### GetCreatorId

`func (o *CorporationsDetail) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CorporationsDetail) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CorporationsDetail) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *CorporationsDetail) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDateFounded

`func (o *CorporationsDetail) GetDateFounded() time.Time`

GetDateFounded returns the DateFounded field if non-nil, zero value otherwise.

### GetDateFoundedOk

`func (o *CorporationsDetail) GetDateFoundedOk() (*time.Time, bool)`

GetDateFoundedOk returns a tuple with the DateFounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFounded

`func (o *CorporationsDetail) SetDateFounded(v time.Time)`

SetDateFounded sets DateFounded field to given value.

### HasDateFounded

`func (o *CorporationsDetail) HasDateFounded() bool`

HasDateFounded returns a boolean if a field has been set.

### GetDescription

`func (o *CorporationsDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CorporationsDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CorporationsDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnlistedFactionId

`func (o *CorporationsDetail) GetEnlistedFactionId() int64`

GetEnlistedFactionId returns the EnlistedFactionId field if non-nil, zero value otherwise.

### GetEnlistedFactionIdOk

`func (o *CorporationsDetail) GetEnlistedFactionIdOk() (*int64, bool)`

GetEnlistedFactionIdOk returns a tuple with the EnlistedFactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnlistedFactionId

`func (o *CorporationsDetail) SetEnlistedFactionId(v int64)`

SetEnlistedFactionId sets EnlistedFactionId field to given value.

### HasEnlistedFactionId

`func (o *CorporationsDetail) HasEnlistedFactionId() bool`

HasEnlistedFactionId returns a boolean if a field has been set.

### GetFriendlyFire

`func (o *CorporationsDetail) GetFriendlyFire() string`

GetFriendlyFire returns the FriendlyFire field if non-nil, zero value otherwise.

### GetFriendlyFireOk

`func (o *CorporationsDetail) GetFriendlyFireOk() (*string, bool)`

GetFriendlyFireOk returns a tuple with the FriendlyFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyFire

`func (o *CorporationsDetail) SetFriendlyFire(v string)`

SetFriendlyFire sets FriendlyFire field to given value.


### GetHomeStationId

`func (o *CorporationsDetail) GetHomeStationId() int64`

GetHomeStationId returns the HomeStationId field if non-nil, zero value otherwise.

### GetHomeStationIdOk

`func (o *CorporationsDetail) GetHomeStationIdOk() (*int64, bool)`

GetHomeStationIdOk returns a tuple with the HomeStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeStationId

`func (o *CorporationsDetail) SetHomeStationId(v int64)`

SetHomeStationId sets HomeStationId field to given value.


### GetMemberCount

`func (o *CorporationsDetail) GetMemberCount() int64`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *CorporationsDetail) GetMemberCountOk() (*int64, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *CorporationsDetail) SetMemberCount(v int64)`

SetMemberCount sets MemberCount field to given value.


### GetName

`func (o *CorporationsDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporationsDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporationsDetail) SetName(v string)`

SetName sets Name field to given value.


### GetPalette

`func (o *CorporationsDetail) GetPalette() CorporationsDetailPalette`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *CorporationsDetail) GetPaletteOk() (*CorporationsDetailPalette, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *CorporationsDetail) SetPalette(v CorporationsDetailPalette)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *CorporationsDetail) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

### GetShares

`func (o *CorporationsDetail) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CorporationsDetail) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CorporationsDetail) SetShares(v int64)`

SetShares sets Shares field to given value.


### GetState

`func (o *CorporationsDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CorporationsDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CorporationsDetail) SetState(v string)`

SetState sets State field to given value.


### GetTaxRates

`func (o *CorporationsDetail) GetTaxRates() CorporationsDetailTaxrates`

GetTaxRates returns the TaxRates field if non-nil, zero value otherwise.

### GetTaxRatesOk

`func (o *CorporationsDetail) GetTaxRatesOk() (*CorporationsDetailTaxrates, bool)`

GetTaxRatesOk returns a tuple with the TaxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRates

`func (o *CorporationsDetail) SetTaxRates(v CorporationsDetailTaxrates)`

SetTaxRates sets TaxRates field to given value.


### GetTicker

`func (o *CorporationsDetail) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *CorporationsDetail) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *CorporationsDetail) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### GetType

`func (o *CorporationsDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CorporationsDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CorporationsDetail) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CorporationsDetail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CorporationsDetail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CorporationsDetail) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CorporationsDetail) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWarEligible

`func (o *CorporationsDetail) GetWarEligible() bool`

GetWarEligible returns the WarEligible field if non-nil, zero value otherwise.

### GetWarEligibleOk

`func (o *CorporationsDetail) GetWarEligibleOk() (*bool, bool)`

GetWarEligibleOk returns a tuple with the WarEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarEligible

`func (o *CorporationsDetail) SetWarEligible(v bool)`

SetWarEligible sets WarEligible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


