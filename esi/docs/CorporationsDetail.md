# CorporationsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** |  | [optional] 
**CeoId** | **int64** |  | 
**CreatorId** | **int64** |  | 
**DateFounded** | Pointer to **time.Time** | Corporation&#39;s founding date | [optional] 
**Description** | Pointer to **string** | Corporation&#39;s description | [optional] 
**FactionId** | Pointer to **int64** |  | [optional] 
**HomeStationId** | Pointer to **int64** |  | [optional] 
**MemberCount** | **int64** | Corporation&#39;s member count | 
**Name** | **string** | Corporation&#39;s name | 
**Shares** | Pointer to **int64** | Corporation&#39;s shares | [optional] 
**TaxRate** | **float64** | Corporation&#39;s tax rate | 
**Ticker** | **string** | Corporation&#39;s short name | 
**Url** | Pointer to **string** | Corporation&#39;s URL | [optional] 
**WarEligible** | Pointer to **bool** | Corporation&#39;s war eligible | [optional] 

## Methods

### NewCorporationsDetail

`func NewCorporationsDetail(ceoId int64, creatorId int64, memberCount int64, name string, taxRate float64, ticker string, ) *CorporationsDetail`

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

### HasDescription

`func (o *CorporationsDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFactionId

`func (o *CorporationsDetail) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *CorporationsDetail) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *CorporationsDetail) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *CorporationsDetail) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

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

### HasHomeStationId

`func (o *CorporationsDetail) HasHomeStationId() bool`

HasHomeStationId returns a boolean if a field has been set.

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

### HasShares

`func (o *CorporationsDetail) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTaxRate

`func (o *CorporationsDetail) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *CorporationsDetail) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *CorporationsDetail) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.


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

### HasWarEligible

`func (o *CorporationsDetail) HasWarEligible() bool`

HasWarEligible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


