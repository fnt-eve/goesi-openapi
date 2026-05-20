# CorporationsStructuresSkyhooksDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveWorkforce** | Pointer to **int64** | Skyhook&#39;s effective workforce; this can differ from the Skyhook&#39;s normal workforce due to the influence of an attached Mercenary Den | [optional] 
**Id** | **int64** |  | 
**IsActive** | **bool** | Whether the Skyhook is active and producing workforce/power/reagents | 
**PlanetId** | **int64** |  | 
**Reagents** | Pointer to [**[]CorporationsStructuresSkyhooksDetailReagent**](CorporationsStructuresSkyhooksDetailReagent.md) | Skyhook&#39;s reagents | [optional] 
**ReinforcementTimer** | Pointer to [**CorporationsStructuresSkyhooksDetailReinforcementtimer**](CorporationsStructuresSkyhooksDetailReinforcementtimer.md) |  | [optional] 
**State** | **string** | Skyhook&#39;s state | 
**TheftVulnerability** | Pointer to [**CorporationsStructuresSkyhooksDetailTheftvulnerability**](CorporationsStructuresSkyhooksDetailTheftvulnerability.md) |  | [optional] 

## Methods

### NewCorporationsStructuresSkyhooksDetail

`func NewCorporationsStructuresSkyhooksDetail(id int64, isActive bool, planetId int64, state string, ) *CorporationsStructuresSkyhooksDetail`

NewCorporationsStructuresSkyhooksDetail instantiates a new CorporationsStructuresSkyhooksDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsStructuresSkyhooksDetailWithDefaults

`func NewCorporationsStructuresSkyhooksDetailWithDefaults() *CorporationsStructuresSkyhooksDetail`

NewCorporationsStructuresSkyhooksDetailWithDefaults instantiates a new CorporationsStructuresSkyhooksDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveWorkforce

`func (o *CorporationsStructuresSkyhooksDetail) GetEffectiveWorkforce() int64`

GetEffectiveWorkforce returns the EffectiveWorkforce field if non-nil, zero value otherwise.

### GetEffectiveWorkforceOk

`func (o *CorporationsStructuresSkyhooksDetail) GetEffectiveWorkforceOk() (*int64, bool)`

GetEffectiveWorkforceOk returns a tuple with the EffectiveWorkforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveWorkforce

`func (o *CorporationsStructuresSkyhooksDetail) SetEffectiveWorkforce(v int64)`

SetEffectiveWorkforce sets EffectiveWorkforce field to given value.

### HasEffectiveWorkforce

`func (o *CorporationsStructuresSkyhooksDetail) HasEffectiveWorkforce() bool`

HasEffectiveWorkforce returns a boolean if a field has been set.

### GetId

`func (o *CorporationsStructuresSkyhooksDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporationsStructuresSkyhooksDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporationsStructuresSkyhooksDetail) SetId(v int64)`

SetId sets Id field to given value.


### GetIsActive

`func (o *CorporationsStructuresSkyhooksDetail) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CorporationsStructuresSkyhooksDetail) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CorporationsStructuresSkyhooksDetail) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPlanetId

`func (o *CorporationsStructuresSkyhooksDetail) GetPlanetId() int64`

GetPlanetId returns the PlanetId field if non-nil, zero value otherwise.

### GetPlanetIdOk

`func (o *CorporationsStructuresSkyhooksDetail) GetPlanetIdOk() (*int64, bool)`

GetPlanetIdOk returns a tuple with the PlanetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanetId

`func (o *CorporationsStructuresSkyhooksDetail) SetPlanetId(v int64)`

SetPlanetId sets PlanetId field to given value.


### GetReagents

`func (o *CorporationsStructuresSkyhooksDetail) GetReagents() []CorporationsStructuresSkyhooksDetailReagent`

GetReagents returns the Reagents field if non-nil, zero value otherwise.

### GetReagentsOk

`func (o *CorporationsStructuresSkyhooksDetail) GetReagentsOk() (*[]CorporationsStructuresSkyhooksDetailReagent, bool)`

GetReagentsOk returns a tuple with the Reagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReagents

`func (o *CorporationsStructuresSkyhooksDetail) SetReagents(v []CorporationsStructuresSkyhooksDetailReagent)`

SetReagents sets Reagents field to given value.

### HasReagents

`func (o *CorporationsStructuresSkyhooksDetail) HasReagents() bool`

HasReagents returns a boolean if a field has been set.

### GetReinforcementTimer

`func (o *CorporationsStructuresSkyhooksDetail) GetReinforcementTimer() CorporationsStructuresSkyhooksDetailReinforcementtimer`

GetReinforcementTimer returns the ReinforcementTimer field if non-nil, zero value otherwise.

### GetReinforcementTimerOk

`func (o *CorporationsStructuresSkyhooksDetail) GetReinforcementTimerOk() (*CorporationsStructuresSkyhooksDetailReinforcementtimer, bool)`

GetReinforcementTimerOk returns a tuple with the ReinforcementTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinforcementTimer

`func (o *CorporationsStructuresSkyhooksDetail) SetReinforcementTimer(v CorporationsStructuresSkyhooksDetailReinforcementtimer)`

SetReinforcementTimer sets ReinforcementTimer field to given value.

### HasReinforcementTimer

`func (o *CorporationsStructuresSkyhooksDetail) HasReinforcementTimer() bool`

HasReinforcementTimer returns a boolean if a field has been set.

### GetState

`func (o *CorporationsStructuresSkyhooksDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CorporationsStructuresSkyhooksDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CorporationsStructuresSkyhooksDetail) SetState(v string)`

SetState sets State field to given value.


### GetTheftVulnerability

`func (o *CorporationsStructuresSkyhooksDetail) GetTheftVulnerability() CorporationsStructuresSkyhooksDetailTheftvulnerability`

GetTheftVulnerability returns the TheftVulnerability field if non-nil, zero value otherwise.

### GetTheftVulnerabilityOk

`func (o *CorporationsStructuresSkyhooksDetail) GetTheftVulnerabilityOk() (*CorporationsStructuresSkyhooksDetailTheftvulnerability, bool)`

GetTheftVulnerabilityOk returns a tuple with the TheftVulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheftVulnerability

`func (o *CorporationsStructuresSkyhooksDetail) SetTheftVulnerability(v CorporationsStructuresSkyhooksDetailTheftvulnerability)`

SetTheftVulnerability sets TheftVulnerability field to given value.

### HasTheftVulnerability

`func (o *CorporationsStructuresSkyhooksDetail) HasTheftVulnerability() bool`

HasTheftVulnerability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


