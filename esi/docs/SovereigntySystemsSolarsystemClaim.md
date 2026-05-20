# SovereigntySystemsSolarsystemClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Faction** | Pointer to [**SovereigntySystemsFaction**](SovereigntySystemsFaction.md) |  | [optional] 
**Alliance** | Pointer to [**SovereigntySystemsAlliance**](SovereigntySystemsAlliance.md) |  | [optional] 
**Unclaimed** | Pointer to **bool** | Solar system is unclaimed | [optional] 

## Methods

### NewSovereigntySystemsSolarsystemClaim

`func NewSovereigntySystemsSolarsystemClaim() *SovereigntySystemsSolarsystemClaim`

NewSovereigntySystemsSolarsystemClaim instantiates a new SovereigntySystemsSolarsystemClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSovereigntySystemsSolarsystemClaimWithDefaults

`func NewSovereigntySystemsSolarsystemClaimWithDefaults() *SovereigntySystemsSolarsystemClaim`

NewSovereigntySystemsSolarsystemClaimWithDefaults instantiates a new SovereigntySystemsSolarsystemClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaction

`func (o *SovereigntySystemsSolarsystemClaim) GetFaction() SovereigntySystemsFaction`

GetFaction returns the Faction field if non-nil, zero value otherwise.

### GetFactionOk

`func (o *SovereigntySystemsSolarsystemClaim) GetFactionOk() (*SovereigntySystemsFaction, bool)`

GetFactionOk returns a tuple with the Faction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaction

`func (o *SovereigntySystemsSolarsystemClaim) SetFaction(v SovereigntySystemsFaction)`

SetFaction sets Faction field to given value.

### HasFaction

`func (o *SovereigntySystemsSolarsystemClaim) HasFaction() bool`

HasFaction returns a boolean if a field has been set.

### GetAlliance

`func (o *SovereigntySystemsSolarsystemClaim) GetAlliance() SovereigntySystemsAlliance`

GetAlliance returns the Alliance field if non-nil, zero value otherwise.

### GetAllianceOk

`func (o *SovereigntySystemsSolarsystemClaim) GetAllianceOk() (*SovereigntySystemsAlliance, bool)`

GetAllianceOk returns a tuple with the Alliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlliance

`func (o *SovereigntySystemsSolarsystemClaim) SetAlliance(v SovereigntySystemsAlliance)`

SetAlliance sets Alliance field to given value.

### HasAlliance

`func (o *SovereigntySystemsSolarsystemClaim) HasAlliance() bool`

HasAlliance returns a boolean if a field has been set.

### GetUnclaimed

`func (o *SovereigntySystemsSolarsystemClaim) GetUnclaimed() bool`

GetUnclaimed returns the Unclaimed field if non-nil, zero value otherwise.

### GetUnclaimedOk

`func (o *SovereigntySystemsSolarsystemClaim) GetUnclaimedOk() (*bool, bool)`

GetUnclaimedOk returns a tuple with the Unclaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclaimed

`func (o *SovereigntySystemsSolarsystemClaim) SetUnclaimed(v bool)`

SetUnclaimed sets Unclaimed field to given value.

### HasUnclaimed

`func (o *SovereigntySystemsSolarsystemClaim) HasUnclaimed() bool`

HasUnclaimed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


