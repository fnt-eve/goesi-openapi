# CharactersCharacterIdAttributesGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedRemapCooldownDate** | Pointer to **time.Time** | Neural remapping cooldown after a character uses remap accrued over time | [optional] 
**BonusRemaps** | Pointer to **int64** | Number of available bonus character neural remaps | [optional] 
**Charisma** | **int64** |  | 
**Intelligence** | **int64** |  | 
**LastRemapDate** | Pointer to **time.Time** | Datetime of last neural remap, including usage of bonus remaps | [optional] 
**Memory** | **int64** |  | 
**Perception** | **int64** |  | 
**Willpower** | **int64** |  | 

## Methods

### NewCharactersCharacterIdAttributesGet

`func NewCharactersCharacterIdAttributesGet(charisma int64, intelligence int64, memory int64, perception int64, willpower int64, ) *CharactersCharacterIdAttributesGet`

NewCharactersCharacterIdAttributesGet instantiates a new CharactersCharacterIdAttributesGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdAttributesGetWithDefaults

`func NewCharactersCharacterIdAttributesGetWithDefaults() *CharactersCharacterIdAttributesGet`

NewCharactersCharacterIdAttributesGetWithDefaults instantiates a new CharactersCharacterIdAttributesGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedRemapCooldownDate

`func (o *CharactersCharacterIdAttributesGet) GetAccruedRemapCooldownDate() time.Time`

GetAccruedRemapCooldownDate returns the AccruedRemapCooldownDate field if non-nil, zero value otherwise.

### GetAccruedRemapCooldownDateOk

`func (o *CharactersCharacterIdAttributesGet) GetAccruedRemapCooldownDateOk() (*time.Time, bool)`

GetAccruedRemapCooldownDateOk returns a tuple with the AccruedRemapCooldownDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedRemapCooldownDate

`func (o *CharactersCharacterIdAttributesGet) SetAccruedRemapCooldownDate(v time.Time)`

SetAccruedRemapCooldownDate sets AccruedRemapCooldownDate field to given value.

### HasAccruedRemapCooldownDate

`func (o *CharactersCharacterIdAttributesGet) HasAccruedRemapCooldownDate() bool`

HasAccruedRemapCooldownDate returns a boolean if a field has been set.

### GetBonusRemaps

`func (o *CharactersCharacterIdAttributesGet) GetBonusRemaps() int64`

GetBonusRemaps returns the BonusRemaps field if non-nil, zero value otherwise.

### GetBonusRemapsOk

`func (o *CharactersCharacterIdAttributesGet) GetBonusRemapsOk() (*int64, bool)`

GetBonusRemapsOk returns a tuple with the BonusRemaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusRemaps

`func (o *CharactersCharacterIdAttributesGet) SetBonusRemaps(v int64)`

SetBonusRemaps sets BonusRemaps field to given value.

### HasBonusRemaps

`func (o *CharactersCharacterIdAttributesGet) HasBonusRemaps() bool`

HasBonusRemaps returns a boolean if a field has been set.

### GetCharisma

`func (o *CharactersCharacterIdAttributesGet) GetCharisma() int64`

GetCharisma returns the Charisma field if non-nil, zero value otherwise.

### GetCharismaOk

`func (o *CharactersCharacterIdAttributesGet) GetCharismaOk() (*int64, bool)`

GetCharismaOk returns a tuple with the Charisma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharisma

`func (o *CharactersCharacterIdAttributesGet) SetCharisma(v int64)`

SetCharisma sets Charisma field to given value.


### GetIntelligence

`func (o *CharactersCharacterIdAttributesGet) GetIntelligence() int64`

GetIntelligence returns the Intelligence field if non-nil, zero value otherwise.

### GetIntelligenceOk

`func (o *CharactersCharacterIdAttributesGet) GetIntelligenceOk() (*int64, bool)`

GetIntelligenceOk returns a tuple with the Intelligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelligence

`func (o *CharactersCharacterIdAttributesGet) SetIntelligence(v int64)`

SetIntelligence sets Intelligence field to given value.


### GetLastRemapDate

`func (o *CharactersCharacterIdAttributesGet) GetLastRemapDate() time.Time`

GetLastRemapDate returns the LastRemapDate field if non-nil, zero value otherwise.

### GetLastRemapDateOk

`func (o *CharactersCharacterIdAttributesGet) GetLastRemapDateOk() (*time.Time, bool)`

GetLastRemapDateOk returns a tuple with the LastRemapDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRemapDate

`func (o *CharactersCharacterIdAttributesGet) SetLastRemapDate(v time.Time)`

SetLastRemapDate sets LastRemapDate field to given value.

### HasLastRemapDate

`func (o *CharactersCharacterIdAttributesGet) HasLastRemapDate() bool`

HasLastRemapDate returns a boolean if a field has been set.

### GetMemory

`func (o *CharactersCharacterIdAttributesGet) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CharactersCharacterIdAttributesGet) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CharactersCharacterIdAttributesGet) SetMemory(v int64)`

SetMemory sets Memory field to given value.


### GetPerception

`func (o *CharactersCharacterIdAttributesGet) GetPerception() int64`

GetPerception returns the Perception field if non-nil, zero value otherwise.

### GetPerceptionOk

`func (o *CharactersCharacterIdAttributesGet) GetPerceptionOk() (*int64, bool)`

GetPerceptionOk returns a tuple with the Perception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerception

`func (o *CharactersCharacterIdAttributesGet) SetPerception(v int64)`

SetPerception sets Perception field to given value.


### GetWillpower

`func (o *CharactersCharacterIdAttributesGet) GetWillpower() int64`

GetWillpower returns the Willpower field if non-nil, zero value otherwise.

### GetWillpowerOk

`func (o *CharactersCharacterIdAttributesGet) GetWillpowerOk() (*int64, bool)`

GetWillpowerOk returns a tuple with the Willpower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillpower

`func (o *CharactersCharacterIdAttributesGet) SetWillpower(v int64)`

SetWillpower sets Willpower field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


