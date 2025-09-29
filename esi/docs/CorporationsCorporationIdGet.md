# CorporationsCorporationIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** | ID of the alliance that corporation is a member of, if any | [optional] 
**CeoId** | **int64** |  | 
**CreatorId** | **int64** |  | 
**DateFounded** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FactionId** | Pointer to **int64** |  | [optional] 
**HomeStationId** | Pointer to **int64** |  | [optional] 
**MemberCount** | **int64** |  | 
**Name** | **string** | the full name of the corporation | 
**Shares** | Pointer to **int64** |  | [optional] 
**TaxRate** | **float64** |  | 
**Ticker** | **string** | the short name of the corporation | 
**Url** | Pointer to **string** |  | [optional] 
**WarEligible** | Pointer to **bool** |  | [optional] 

## Methods

### NewCorporationsCorporationIdGet

`func NewCorporationsCorporationIdGet(ceoId int64, creatorId int64, memberCount int64, name string, taxRate float64, ticker string, ) *CorporationsCorporationIdGet`

NewCorporationsCorporationIdGet instantiates a new CorporationsCorporationIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdGetWithDefaults

`func NewCorporationsCorporationIdGetWithDefaults() *CorporationsCorporationIdGet`

NewCorporationsCorporationIdGetWithDefaults instantiates a new CorporationsCorporationIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *CorporationsCorporationIdGet) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *CorporationsCorporationIdGet) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *CorporationsCorporationIdGet) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *CorporationsCorporationIdGet) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetCeoId

`func (o *CorporationsCorporationIdGet) GetCeoId() int64`

GetCeoId returns the CeoId field if non-nil, zero value otherwise.

### GetCeoIdOk

`func (o *CorporationsCorporationIdGet) GetCeoIdOk() (*int64, bool)`

GetCeoIdOk returns a tuple with the CeoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeoId

`func (o *CorporationsCorporationIdGet) SetCeoId(v int64)`

SetCeoId sets CeoId field to given value.


### GetCreatorId

`func (o *CorporationsCorporationIdGet) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CorporationsCorporationIdGet) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CorporationsCorporationIdGet) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.


### GetDateFounded

`func (o *CorporationsCorporationIdGet) GetDateFounded() time.Time`

GetDateFounded returns the DateFounded field if non-nil, zero value otherwise.

### GetDateFoundedOk

`func (o *CorporationsCorporationIdGet) GetDateFoundedOk() (*time.Time, bool)`

GetDateFoundedOk returns a tuple with the DateFounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFounded

`func (o *CorporationsCorporationIdGet) SetDateFounded(v time.Time)`

SetDateFounded sets DateFounded field to given value.

### HasDateFounded

`func (o *CorporationsCorporationIdGet) HasDateFounded() bool`

HasDateFounded returns a boolean if a field has been set.

### GetDescription

`func (o *CorporationsCorporationIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CorporationsCorporationIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CorporationsCorporationIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CorporationsCorporationIdGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFactionId

`func (o *CorporationsCorporationIdGet) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *CorporationsCorporationIdGet) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *CorporationsCorporationIdGet) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *CorporationsCorporationIdGet) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetHomeStationId

`func (o *CorporationsCorporationIdGet) GetHomeStationId() int64`

GetHomeStationId returns the HomeStationId field if non-nil, zero value otherwise.

### GetHomeStationIdOk

`func (o *CorporationsCorporationIdGet) GetHomeStationIdOk() (*int64, bool)`

GetHomeStationIdOk returns a tuple with the HomeStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeStationId

`func (o *CorporationsCorporationIdGet) SetHomeStationId(v int64)`

SetHomeStationId sets HomeStationId field to given value.

### HasHomeStationId

`func (o *CorporationsCorporationIdGet) HasHomeStationId() bool`

HasHomeStationId returns a boolean if a field has been set.

### GetMemberCount

`func (o *CorporationsCorporationIdGet) GetMemberCount() int64`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *CorporationsCorporationIdGet) GetMemberCountOk() (*int64, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *CorporationsCorporationIdGet) SetMemberCount(v int64)`

SetMemberCount sets MemberCount field to given value.


### GetName

`func (o *CorporationsCorporationIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporationsCorporationIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporationsCorporationIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetShares

`func (o *CorporationsCorporationIdGet) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CorporationsCorporationIdGet) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CorporationsCorporationIdGet) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CorporationsCorporationIdGet) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTaxRate

`func (o *CorporationsCorporationIdGet) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *CorporationsCorporationIdGet) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *CorporationsCorporationIdGet) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.


### GetTicker

`func (o *CorporationsCorporationIdGet) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *CorporationsCorporationIdGet) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *CorporationsCorporationIdGet) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### GetUrl

`func (o *CorporationsCorporationIdGet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CorporationsCorporationIdGet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CorporationsCorporationIdGet) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CorporationsCorporationIdGet) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWarEligible

`func (o *CorporationsCorporationIdGet) GetWarEligible() bool`

GetWarEligible returns the WarEligible field if non-nil, zero value otherwise.

### GetWarEligibleOk

`func (o *CorporationsCorporationIdGet) GetWarEligibleOk() (*bool, bool)`

GetWarEligibleOk returns a tuple with the WarEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarEligible

`func (o *CorporationsCorporationIdGet) SetWarEligible(v bool)`

SetWarEligible sets WarEligible field to given value.

### HasWarEligible

`func (o *CorporationsCorporationIdGet) HasWarEligible() bool`

HasWarEligible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


