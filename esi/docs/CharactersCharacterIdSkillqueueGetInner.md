# CharactersCharacterIdSkillqueueGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishDate** | Pointer to **time.Time** | Date on which training of the skill will complete. Omitted if the skill queue is paused. | [optional] 
**FinishedLevel** | **int64** |  | 
**LevelEndSp** | Pointer to **int64** |  | [optional] 
**LevelStartSp** | Pointer to **int64** | Amount of SP that was in the skill when it started training it&#39;s current level. Used to calculate % of current level complete. | [optional] 
**QueuePosition** | **int64** |  | 
**SkillId** | **int64** |  | 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**TrainingStartSp** | Pointer to **int64** |  | [optional] 

## Methods

### NewCharactersCharacterIdSkillqueueGetInner

`func NewCharactersCharacterIdSkillqueueGetInner(finishedLevel int64, queuePosition int64, skillId int64, ) *CharactersCharacterIdSkillqueueGetInner`

NewCharactersCharacterIdSkillqueueGetInner instantiates a new CharactersCharacterIdSkillqueueGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdSkillqueueGetInnerWithDefaults

`func NewCharactersCharacterIdSkillqueueGetInnerWithDefaults() *CharactersCharacterIdSkillqueueGetInner`

NewCharactersCharacterIdSkillqueueGetInnerWithDefaults instantiates a new CharactersCharacterIdSkillqueueGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishDate

`func (o *CharactersCharacterIdSkillqueueGetInner) GetFinishDate() time.Time`

GetFinishDate returns the FinishDate field if non-nil, zero value otherwise.

### GetFinishDateOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetFinishDateOk() (*time.Time, bool)`

GetFinishDateOk returns a tuple with the FinishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishDate

`func (o *CharactersCharacterIdSkillqueueGetInner) SetFinishDate(v time.Time)`

SetFinishDate sets FinishDate field to given value.

### HasFinishDate

`func (o *CharactersCharacterIdSkillqueueGetInner) HasFinishDate() bool`

HasFinishDate returns a boolean if a field has been set.

### GetFinishedLevel

`func (o *CharactersCharacterIdSkillqueueGetInner) GetFinishedLevel() int64`

GetFinishedLevel returns the FinishedLevel field if non-nil, zero value otherwise.

### GetFinishedLevelOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetFinishedLevelOk() (*int64, bool)`

GetFinishedLevelOk returns a tuple with the FinishedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedLevel

`func (o *CharactersCharacterIdSkillqueueGetInner) SetFinishedLevel(v int64)`

SetFinishedLevel sets FinishedLevel field to given value.


### GetLevelEndSp

`func (o *CharactersCharacterIdSkillqueueGetInner) GetLevelEndSp() int64`

GetLevelEndSp returns the LevelEndSp field if non-nil, zero value otherwise.

### GetLevelEndSpOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetLevelEndSpOk() (*int64, bool)`

GetLevelEndSpOk returns a tuple with the LevelEndSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelEndSp

`func (o *CharactersCharacterIdSkillqueueGetInner) SetLevelEndSp(v int64)`

SetLevelEndSp sets LevelEndSp field to given value.

### HasLevelEndSp

`func (o *CharactersCharacterIdSkillqueueGetInner) HasLevelEndSp() bool`

HasLevelEndSp returns a boolean if a field has been set.

### GetLevelStartSp

`func (o *CharactersCharacterIdSkillqueueGetInner) GetLevelStartSp() int64`

GetLevelStartSp returns the LevelStartSp field if non-nil, zero value otherwise.

### GetLevelStartSpOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetLevelStartSpOk() (*int64, bool)`

GetLevelStartSpOk returns a tuple with the LevelStartSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelStartSp

`func (o *CharactersCharacterIdSkillqueueGetInner) SetLevelStartSp(v int64)`

SetLevelStartSp sets LevelStartSp field to given value.

### HasLevelStartSp

`func (o *CharactersCharacterIdSkillqueueGetInner) HasLevelStartSp() bool`

HasLevelStartSp returns a boolean if a field has been set.

### GetQueuePosition

`func (o *CharactersCharacterIdSkillqueueGetInner) GetQueuePosition() int64`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetQueuePositionOk() (*int64, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *CharactersCharacterIdSkillqueueGetInner) SetQueuePosition(v int64)`

SetQueuePosition sets QueuePosition field to given value.


### GetSkillId

`func (o *CharactersCharacterIdSkillqueueGetInner) GetSkillId() int64`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetSkillIdOk() (*int64, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *CharactersCharacterIdSkillqueueGetInner) SetSkillId(v int64)`

SetSkillId sets SkillId field to given value.


### GetStartDate

`func (o *CharactersCharacterIdSkillqueueGetInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CharactersCharacterIdSkillqueueGetInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CharactersCharacterIdSkillqueueGetInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTrainingStartSp

`func (o *CharactersCharacterIdSkillqueueGetInner) GetTrainingStartSp() int64`

GetTrainingStartSp returns the TrainingStartSp field if non-nil, zero value otherwise.

### GetTrainingStartSpOk

`func (o *CharactersCharacterIdSkillqueueGetInner) GetTrainingStartSpOk() (*int64, bool)`

GetTrainingStartSpOk returns a tuple with the TrainingStartSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingStartSp

`func (o *CharactersCharacterIdSkillqueueGetInner) SetTrainingStartSp(v int64)`

SetTrainingStartSp sets TrainingStartSp field to given value.

### HasTrainingStartSp

`func (o *CharactersCharacterIdSkillqueueGetInner) HasTrainingStartSp() bool`

HasTrainingStartSp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


