# CorporationsCorporationIdFwStatsGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnlistedOn** | Pointer to **time.Time** | The enlistment date of the given corporation into faction warfare. Will not be included if corporation is not enlisted in faction warfare | [optional] 
**FactionId** | Pointer to **int64** | The faction the given corporation is enlisted to fight for. Will not be included if corporation is not enlisted in faction warfare | [optional] 
**Kills** | [**CorporationsCorporationIdFwStatsGetKills**](CorporationsCorporationIdFwStatsGetKills.md) |  | 
**Pilots** | Pointer to **int64** | How many pilots the enlisted corporation has. Will not be included if corporation is not enlisted in faction warfare | [optional] 
**VictoryPoints** | [**CorporationsCorporationIdFwStatsGetVictoryPoints**](CorporationsCorporationIdFwStatsGetVictoryPoints.md) |  | 

## Methods

### NewCorporationsCorporationIdFwStatsGet

`func NewCorporationsCorporationIdFwStatsGet(kills CorporationsCorporationIdFwStatsGetKills, victoryPoints CorporationsCorporationIdFwStatsGetVictoryPoints, ) *CorporationsCorporationIdFwStatsGet`

NewCorporationsCorporationIdFwStatsGet instantiates a new CorporationsCorporationIdFwStatsGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdFwStatsGetWithDefaults

`func NewCorporationsCorporationIdFwStatsGetWithDefaults() *CorporationsCorporationIdFwStatsGet`

NewCorporationsCorporationIdFwStatsGetWithDefaults instantiates a new CorporationsCorporationIdFwStatsGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnlistedOn

`func (o *CorporationsCorporationIdFwStatsGet) GetEnlistedOn() time.Time`

GetEnlistedOn returns the EnlistedOn field if non-nil, zero value otherwise.

### GetEnlistedOnOk

`func (o *CorporationsCorporationIdFwStatsGet) GetEnlistedOnOk() (*time.Time, bool)`

GetEnlistedOnOk returns a tuple with the EnlistedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnlistedOn

`func (o *CorporationsCorporationIdFwStatsGet) SetEnlistedOn(v time.Time)`

SetEnlistedOn sets EnlistedOn field to given value.

### HasEnlistedOn

`func (o *CorporationsCorporationIdFwStatsGet) HasEnlistedOn() bool`

HasEnlistedOn returns a boolean if a field has been set.

### GetFactionId

`func (o *CorporationsCorporationIdFwStatsGet) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *CorporationsCorporationIdFwStatsGet) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *CorporationsCorporationIdFwStatsGet) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *CorporationsCorporationIdFwStatsGet) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetKills

`func (o *CorporationsCorporationIdFwStatsGet) GetKills() CorporationsCorporationIdFwStatsGetKills`

GetKills returns the Kills field if non-nil, zero value otherwise.

### GetKillsOk

`func (o *CorporationsCorporationIdFwStatsGet) GetKillsOk() (*CorporationsCorporationIdFwStatsGetKills, bool)`

GetKillsOk returns a tuple with the Kills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKills

`func (o *CorporationsCorporationIdFwStatsGet) SetKills(v CorporationsCorporationIdFwStatsGetKills)`

SetKills sets Kills field to given value.


### GetPilots

`func (o *CorporationsCorporationIdFwStatsGet) GetPilots() int64`

GetPilots returns the Pilots field if non-nil, zero value otherwise.

### GetPilotsOk

`func (o *CorporationsCorporationIdFwStatsGet) GetPilotsOk() (*int64, bool)`

GetPilotsOk returns a tuple with the Pilots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPilots

`func (o *CorporationsCorporationIdFwStatsGet) SetPilots(v int64)`

SetPilots sets Pilots field to given value.

### HasPilots

`func (o *CorporationsCorporationIdFwStatsGet) HasPilots() bool`

HasPilots returns a boolean if a field has been set.

### GetVictoryPoints

`func (o *CorporationsCorporationIdFwStatsGet) GetVictoryPoints() CorporationsCorporationIdFwStatsGetVictoryPoints`

GetVictoryPoints returns the VictoryPoints field if non-nil, zero value otherwise.

### GetVictoryPointsOk

`func (o *CorporationsCorporationIdFwStatsGet) GetVictoryPointsOk() (*CorporationsCorporationIdFwStatsGetVictoryPoints, bool)`

GetVictoryPointsOk returns a tuple with the VictoryPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictoryPoints

`func (o *CorporationsCorporationIdFwStatsGet) SetVictoryPoints(v CorporationsCorporationIdFwStatsGetVictoryPoints)`

SetVictoryPoints sets VictoryPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


