# CorporationsCorporationIdCustomsOfficesGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceTaxRate** | Pointer to **float64** | Only present if alliance access is allowed | [optional] 
**AllowAccessWithStandings** | **bool** | standing_level and any standing related tax rate only present when this is true | 
**AllowAllianceAccess** | **bool** |  | 
**BadStandingTaxRate** | Pointer to **float64** |  | [optional] 
**CorporationTaxRate** | Pointer to **float64** |  | [optional] 
**ExcellentStandingTaxRate** | Pointer to **float64** | Tax rate for entities with excellent level of standing, only present if this level is allowed, same for all other standing related tax rates | [optional] 
**GoodStandingTaxRate** | Pointer to **float64** |  | [optional] 
**NeutralStandingTaxRate** | Pointer to **float64** |  | [optional] 
**OfficeId** | **int64** | unique ID of this customs office | 
**ReinforceExitEnd** | **int64** |  | 
**ReinforceExitStart** | **int64** | Together with reinforce_exit_end, marks a 2-hour period where this customs office could exit reinforcement mode during the day after initial attack | 
**StandingLevel** | Pointer to **string** | Access is allowed only for entities with this level of standing or better | [optional] 
**SystemId** | **int64** | ID of the solar system this customs office is located in | 
**TerribleStandingTaxRate** | Pointer to **float64** |  | [optional] 

## Methods

### NewCorporationsCorporationIdCustomsOfficesGetInner

`func NewCorporationsCorporationIdCustomsOfficesGetInner(allowAccessWithStandings bool, allowAllianceAccess bool, officeId int64, reinforceExitEnd int64, reinforceExitStart int64, systemId int64, ) *CorporationsCorporationIdCustomsOfficesGetInner`

NewCorporationsCorporationIdCustomsOfficesGetInner instantiates a new CorporationsCorporationIdCustomsOfficesGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdCustomsOfficesGetInnerWithDefaults

`func NewCorporationsCorporationIdCustomsOfficesGetInnerWithDefaults() *CorporationsCorporationIdCustomsOfficesGetInner`

NewCorporationsCorporationIdCustomsOfficesGetInnerWithDefaults instantiates a new CorporationsCorporationIdCustomsOfficesGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetAllianceTaxRate() float64`

GetAllianceTaxRate returns the AllianceTaxRate field if non-nil, zero value otherwise.

### GetAllianceTaxRateOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetAllianceTaxRateOk() (*float64, bool)`

GetAllianceTaxRateOk returns a tuple with the AllianceTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetAllianceTaxRate(v float64)`

SetAllianceTaxRate sets AllianceTaxRate field to given value.

### HasAllianceTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasAllianceTaxRate() bool`

HasAllianceTaxRate returns a boolean if a field has been set.

### GetAllowAccessWithStandings

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetAllowAccessWithStandings() bool`

GetAllowAccessWithStandings returns the AllowAccessWithStandings field if non-nil, zero value otherwise.

### GetAllowAccessWithStandingsOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetAllowAccessWithStandingsOk() (*bool, bool)`

GetAllowAccessWithStandingsOk returns a tuple with the AllowAccessWithStandings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAccessWithStandings

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetAllowAccessWithStandings(v bool)`

SetAllowAccessWithStandings sets AllowAccessWithStandings field to given value.


### GetAllowAllianceAccess

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetAllowAllianceAccess() bool`

GetAllowAllianceAccess returns the AllowAllianceAccess field if non-nil, zero value otherwise.

### GetAllowAllianceAccessOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetAllowAllianceAccessOk() (*bool, bool)`

GetAllowAllianceAccessOk returns a tuple with the AllowAllianceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllianceAccess

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetAllowAllianceAccess(v bool)`

SetAllowAllianceAccess sets AllowAllianceAccess field to given value.


### GetBadStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetBadStandingTaxRate() float64`

GetBadStandingTaxRate returns the BadStandingTaxRate field if non-nil, zero value otherwise.

### GetBadStandingTaxRateOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetBadStandingTaxRateOk() (*float64, bool)`

GetBadStandingTaxRateOk returns a tuple with the BadStandingTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetBadStandingTaxRate(v float64)`

SetBadStandingTaxRate sets BadStandingTaxRate field to given value.

### HasBadStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasBadStandingTaxRate() bool`

HasBadStandingTaxRate returns a boolean if a field has been set.

### GetCorporationTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetCorporationTaxRate() float64`

GetCorporationTaxRate returns the CorporationTaxRate field if non-nil, zero value otherwise.

### GetCorporationTaxRateOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetCorporationTaxRateOk() (*float64, bool)`

GetCorporationTaxRateOk returns a tuple with the CorporationTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetCorporationTaxRate(v float64)`

SetCorporationTaxRate sets CorporationTaxRate field to given value.

### HasCorporationTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasCorporationTaxRate() bool`

HasCorporationTaxRate returns a boolean if a field has been set.

### GetExcellentStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetExcellentStandingTaxRate() float64`

GetExcellentStandingTaxRate returns the ExcellentStandingTaxRate field if non-nil, zero value otherwise.

### GetExcellentStandingTaxRateOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetExcellentStandingTaxRateOk() (*float64, bool)`

GetExcellentStandingTaxRateOk returns a tuple with the ExcellentStandingTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcellentStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetExcellentStandingTaxRate(v float64)`

SetExcellentStandingTaxRate sets ExcellentStandingTaxRate field to given value.

### HasExcellentStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasExcellentStandingTaxRate() bool`

HasExcellentStandingTaxRate returns a boolean if a field has been set.

### GetGoodStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetGoodStandingTaxRate() float64`

GetGoodStandingTaxRate returns the GoodStandingTaxRate field if non-nil, zero value otherwise.

### GetGoodStandingTaxRateOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetGoodStandingTaxRateOk() (*float64, bool)`

GetGoodStandingTaxRateOk returns a tuple with the GoodStandingTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetGoodStandingTaxRate(v float64)`

SetGoodStandingTaxRate sets GoodStandingTaxRate field to given value.

### HasGoodStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasGoodStandingTaxRate() bool`

HasGoodStandingTaxRate returns a boolean if a field has been set.

### GetNeutralStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetNeutralStandingTaxRate() float64`

GetNeutralStandingTaxRate returns the NeutralStandingTaxRate field if non-nil, zero value otherwise.

### GetNeutralStandingTaxRateOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetNeutralStandingTaxRateOk() (*float64, bool)`

GetNeutralStandingTaxRateOk returns a tuple with the NeutralStandingTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetNeutralStandingTaxRate(v float64)`

SetNeutralStandingTaxRate sets NeutralStandingTaxRate field to given value.

### HasNeutralStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasNeutralStandingTaxRate() bool`

HasNeutralStandingTaxRate returns a boolean if a field has been set.

### GetOfficeId

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetOfficeId() int64`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetOfficeIdOk() (*int64, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetOfficeId(v int64)`

SetOfficeId sets OfficeId field to given value.


### GetReinforceExitEnd

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetReinforceExitEnd() int64`

GetReinforceExitEnd returns the ReinforceExitEnd field if non-nil, zero value otherwise.

### GetReinforceExitEndOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetReinforceExitEndOk() (*int64, bool)`

GetReinforceExitEndOk returns a tuple with the ReinforceExitEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinforceExitEnd

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetReinforceExitEnd(v int64)`

SetReinforceExitEnd sets ReinforceExitEnd field to given value.


### GetReinforceExitStart

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetReinforceExitStart() int64`

GetReinforceExitStart returns the ReinforceExitStart field if non-nil, zero value otherwise.

### GetReinforceExitStartOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetReinforceExitStartOk() (*int64, bool)`

GetReinforceExitStartOk returns a tuple with the ReinforceExitStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinforceExitStart

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetReinforceExitStart(v int64)`

SetReinforceExitStart sets ReinforceExitStart field to given value.


### GetStandingLevel

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetStandingLevel() string`

GetStandingLevel returns the StandingLevel field if non-nil, zero value otherwise.

### GetStandingLevelOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetStandingLevelOk() (*string, bool)`

GetStandingLevelOk returns a tuple with the StandingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingLevel

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetStandingLevel(v string)`

SetStandingLevel sets StandingLevel field to given value.

### HasStandingLevel

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasStandingLevel() bool`

HasStandingLevel returns a boolean if a field has been set.

### GetSystemId

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetSystemId() int64`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetSystemIdOk() (*int64, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetSystemId(v int64)`

SetSystemId sets SystemId field to given value.


### GetTerribleStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetTerribleStandingTaxRate() float64`

GetTerribleStandingTaxRate returns the TerribleStandingTaxRate field if non-nil, zero value otherwise.

### GetTerribleStandingTaxRateOk

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) GetTerribleStandingTaxRateOk() (*float64, bool)`

GetTerribleStandingTaxRateOk returns a tuple with the TerribleStandingTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerribleStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) SetTerribleStandingTaxRate(v float64)`

SetTerribleStandingTaxRate sets TerribleStandingTaxRate field to given value.

### HasTerribleStandingTaxRate

`func (o *CorporationsCorporationIdCustomsOfficesGetInner) HasTerribleStandingTaxRate() bool`

HasTerribleStandingTaxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


