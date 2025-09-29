# CharactersCharacterIdFwStatsGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentRank** | Pointer to **int64** | The given character&#39;s current faction rank | [optional] 
**EnlistedOn** | Pointer to **time.Time** | The enlistment date of the given character into faction warfare. Will not be included if character is not enlisted in faction warfare | [optional] 
**FactionId** | Pointer to **int64** | The faction the given character is enlisted to fight for. Will not be included if character is not enlisted in faction warfare | [optional] 
**HighestRank** | Pointer to **int64** | The given character&#39;s highest faction rank achieved | [optional] 
**Kills** | [**CharactersCharacterIdFwStatsGetKills**](CharactersCharacterIdFwStatsGetKills.md) |  | 
**VictoryPoints** | [**CharactersCharacterIdFwStatsGetVictoryPoints**](CharactersCharacterIdFwStatsGetVictoryPoints.md) |  | 

## Methods

### NewCharactersCharacterIdFwStatsGet

`func NewCharactersCharacterIdFwStatsGet(kills CharactersCharacterIdFwStatsGetKills, victoryPoints CharactersCharacterIdFwStatsGetVictoryPoints, ) *CharactersCharacterIdFwStatsGet`

NewCharactersCharacterIdFwStatsGet instantiates a new CharactersCharacterIdFwStatsGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdFwStatsGetWithDefaults

`func NewCharactersCharacterIdFwStatsGetWithDefaults() *CharactersCharacterIdFwStatsGet`

NewCharactersCharacterIdFwStatsGetWithDefaults instantiates a new CharactersCharacterIdFwStatsGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentRank

`func (o *CharactersCharacterIdFwStatsGet) GetCurrentRank() int64`

GetCurrentRank returns the CurrentRank field if non-nil, zero value otherwise.

### GetCurrentRankOk

`func (o *CharactersCharacterIdFwStatsGet) GetCurrentRankOk() (*int64, bool)`

GetCurrentRankOk returns a tuple with the CurrentRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRank

`func (o *CharactersCharacterIdFwStatsGet) SetCurrentRank(v int64)`

SetCurrentRank sets CurrentRank field to given value.

### HasCurrentRank

`func (o *CharactersCharacterIdFwStatsGet) HasCurrentRank() bool`

HasCurrentRank returns a boolean if a field has been set.

### GetEnlistedOn

`func (o *CharactersCharacterIdFwStatsGet) GetEnlistedOn() time.Time`

GetEnlistedOn returns the EnlistedOn field if non-nil, zero value otherwise.

### GetEnlistedOnOk

`func (o *CharactersCharacterIdFwStatsGet) GetEnlistedOnOk() (*time.Time, bool)`

GetEnlistedOnOk returns a tuple with the EnlistedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnlistedOn

`func (o *CharactersCharacterIdFwStatsGet) SetEnlistedOn(v time.Time)`

SetEnlistedOn sets EnlistedOn field to given value.

### HasEnlistedOn

`func (o *CharactersCharacterIdFwStatsGet) HasEnlistedOn() bool`

HasEnlistedOn returns a boolean if a field has been set.

### GetFactionId

`func (o *CharactersCharacterIdFwStatsGet) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *CharactersCharacterIdFwStatsGet) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *CharactersCharacterIdFwStatsGet) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *CharactersCharacterIdFwStatsGet) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetHighestRank

`func (o *CharactersCharacterIdFwStatsGet) GetHighestRank() int64`

GetHighestRank returns the HighestRank field if non-nil, zero value otherwise.

### GetHighestRankOk

`func (o *CharactersCharacterIdFwStatsGet) GetHighestRankOk() (*int64, bool)`

GetHighestRankOk returns a tuple with the HighestRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestRank

`func (o *CharactersCharacterIdFwStatsGet) SetHighestRank(v int64)`

SetHighestRank sets HighestRank field to given value.

### HasHighestRank

`func (o *CharactersCharacterIdFwStatsGet) HasHighestRank() bool`

HasHighestRank returns a boolean if a field has been set.

### GetKills

`func (o *CharactersCharacterIdFwStatsGet) GetKills() CharactersCharacterIdFwStatsGetKills`

GetKills returns the Kills field if non-nil, zero value otherwise.

### GetKillsOk

`func (o *CharactersCharacterIdFwStatsGet) GetKillsOk() (*CharactersCharacterIdFwStatsGetKills, bool)`

GetKillsOk returns a tuple with the Kills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKills

`func (o *CharactersCharacterIdFwStatsGet) SetKills(v CharactersCharacterIdFwStatsGetKills)`

SetKills sets Kills field to given value.


### GetVictoryPoints

`func (o *CharactersCharacterIdFwStatsGet) GetVictoryPoints() CharactersCharacterIdFwStatsGetVictoryPoints`

GetVictoryPoints returns the VictoryPoints field if non-nil, zero value otherwise.

### GetVictoryPointsOk

`func (o *CharactersCharacterIdFwStatsGet) GetVictoryPointsOk() (*CharactersCharacterIdFwStatsGetVictoryPoints, bool)`

GetVictoryPointsOk returns a tuple with the VictoryPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictoryPoints

`func (o *CharactersCharacterIdFwStatsGet) SetVictoryPoints(v CharactersCharacterIdFwStatsGetVictoryPoints)`

SetVictoryPoints sets VictoryPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


