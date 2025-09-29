# FwStatsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactionId** | **int64** |  | 
**Kills** | [**FwStatsGetInnerKills**](FwStatsGetInnerKills.md) |  | 
**Pilots** | **int64** | How many pilots fight for the given faction | 
**SystemsControlled** | **int64** | The number of solar systems controlled by the given faction | 
**VictoryPoints** | [**FwStatsGetInnerVictoryPoints**](FwStatsGetInnerVictoryPoints.md) |  | 

## Methods

### NewFwStatsGetInner

`func NewFwStatsGetInner(factionId int64, kills FwStatsGetInnerKills, pilots int64, systemsControlled int64, victoryPoints FwStatsGetInnerVictoryPoints, ) *FwStatsGetInner`

NewFwStatsGetInner instantiates a new FwStatsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwStatsGetInnerWithDefaults

`func NewFwStatsGetInnerWithDefaults() *FwStatsGetInner`

NewFwStatsGetInnerWithDefaults instantiates a new FwStatsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactionId

`func (o *FwStatsGetInner) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *FwStatsGetInner) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *FwStatsGetInner) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.


### GetKills

`func (o *FwStatsGetInner) GetKills() FwStatsGetInnerKills`

GetKills returns the Kills field if non-nil, zero value otherwise.

### GetKillsOk

`func (o *FwStatsGetInner) GetKillsOk() (*FwStatsGetInnerKills, bool)`

GetKillsOk returns a tuple with the Kills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKills

`func (o *FwStatsGetInner) SetKills(v FwStatsGetInnerKills)`

SetKills sets Kills field to given value.


### GetPilots

`func (o *FwStatsGetInner) GetPilots() int64`

GetPilots returns the Pilots field if non-nil, zero value otherwise.

### GetPilotsOk

`func (o *FwStatsGetInner) GetPilotsOk() (*int64, bool)`

GetPilotsOk returns a tuple with the Pilots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPilots

`func (o *FwStatsGetInner) SetPilots(v int64)`

SetPilots sets Pilots field to given value.


### GetSystemsControlled

`func (o *FwStatsGetInner) GetSystemsControlled() int64`

GetSystemsControlled returns the SystemsControlled field if non-nil, zero value otherwise.

### GetSystemsControlledOk

`func (o *FwStatsGetInner) GetSystemsControlledOk() (*int64, bool)`

GetSystemsControlledOk returns a tuple with the SystemsControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsControlled

`func (o *FwStatsGetInner) SetSystemsControlled(v int64)`

SetSystemsControlled sets SystemsControlled field to given value.


### GetVictoryPoints

`func (o *FwStatsGetInner) GetVictoryPoints() FwStatsGetInnerVictoryPoints`

GetVictoryPoints returns the VictoryPoints field if non-nil, zero value otherwise.

### GetVictoryPointsOk

`func (o *FwStatsGetInner) GetVictoryPointsOk() (*FwStatsGetInnerVictoryPoints, bool)`

GetVictoryPointsOk returns a tuple with the VictoryPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictoryPoints

`func (o *FwStatsGetInner) SetVictoryPoints(v FwStatsGetInnerVictoryPoints)`

SetVictoryPoints sets VictoryPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


