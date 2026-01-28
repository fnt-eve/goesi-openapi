# CharactersSkillqueueSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishDate** | Pointer to **time.Time** | The date the skill training will finish | [optional] 
**FinishedLevel** | **int64** | The level the skill is training for | 
**LevelEndSp** | Pointer to **int64** | The Skill Points at the end of the level | [optional] 
**LevelStartSp** | Pointer to **int64** | The Skill Points at the start of the level | [optional] 
**QueuePosition** | **int64** | The position of the skill in the queue | 
**SkillId** | **int64** |  | 
**StartDate** | Pointer to **time.Time** | The date the skill training will start/continue | [optional] 
**TrainingStartSp** | Pointer to **int64** | The Skill Points at the start of training | [optional] 

## Methods

### NewCharactersSkillqueueSkill

`func NewCharactersSkillqueueSkill(finishedLevel int64, queuePosition int64, skillId int64, ) *CharactersSkillqueueSkill`

NewCharactersSkillqueueSkill instantiates a new CharactersSkillqueueSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersSkillqueueSkillWithDefaults

`func NewCharactersSkillqueueSkillWithDefaults() *CharactersSkillqueueSkill`

NewCharactersSkillqueueSkillWithDefaults instantiates a new CharactersSkillqueueSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishDate

`func (o *CharactersSkillqueueSkill) GetFinishDate() time.Time`

GetFinishDate returns the FinishDate field if non-nil, zero value otherwise.

### GetFinishDateOk

`func (o *CharactersSkillqueueSkill) GetFinishDateOk() (*time.Time, bool)`

GetFinishDateOk returns a tuple with the FinishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishDate

`func (o *CharactersSkillqueueSkill) SetFinishDate(v time.Time)`

SetFinishDate sets FinishDate field to given value.

### HasFinishDate

`func (o *CharactersSkillqueueSkill) HasFinishDate() bool`

HasFinishDate returns a boolean if a field has been set.

### GetFinishedLevel

`func (o *CharactersSkillqueueSkill) GetFinishedLevel() int64`

GetFinishedLevel returns the FinishedLevel field if non-nil, zero value otherwise.

### GetFinishedLevelOk

`func (o *CharactersSkillqueueSkill) GetFinishedLevelOk() (*int64, bool)`

GetFinishedLevelOk returns a tuple with the FinishedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedLevel

`func (o *CharactersSkillqueueSkill) SetFinishedLevel(v int64)`

SetFinishedLevel sets FinishedLevel field to given value.


### GetLevelEndSp

`func (o *CharactersSkillqueueSkill) GetLevelEndSp() int64`

GetLevelEndSp returns the LevelEndSp field if non-nil, zero value otherwise.

### GetLevelEndSpOk

`func (o *CharactersSkillqueueSkill) GetLevelEndSpOk() (*int64, bool)`

GetLevelEndSpOk returns a tuple with the LevelEndSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelEndSp

`func (o *CharactersSkillqueueSkill) SetLevelEndSp(v int64)`

SetLevelEndSp sets LevelEndSp field to given value.

### HasLevelEndSp

`func (o *CharactersSkillqueueSkill) HasLevelEndSp() bool`

HasLevelEndSp returns a boolean if a field has been set.

### GetLevelStartSp

`func (o *CharactersSkillqueueSkill) GetLevelStartSp() int64`

GetLevelStartSp returns the LevelStartSp field if non-nil, zero value otherwise.

### GetLevelStartSpOk

`func (o *CharactersSkillqueueSkill) GetLevelStartSpOk() (*int64, bool)`

GetLevelStartSpOk returns a tuple with the LevelStartSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelStartSp

`func (o *CharactersSkillqueueSkill) SetLevelStartSp(v int64)`

SetLevelStartSp sets LevelStartSp field to given value.

### HasLevelStartSp

`func (o *CharactersSkillqueueSkill) HasLevelStartSp() bool`

HasLevelStartSp returns a boolean if a field has been set.

### GetQueuePosition

`func (o *CharactersSkillqueueSkill) GetQueuePosition() int64`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *CharactersSkillqueueSkill) GetQueuePositionOk() (*int64, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *CharactersSkillqueueSkill) SetQueuePosition(v int64)`

SetQueuePosition sets QueuePosition field to given value.


### GetSkillId

`func (o *CharactersSkillqueueSkill) GetSkillId() int64`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *CharactersSkillqueueSkill) GetSkillIdOk() (*int64, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *CharactersSkillqueueSkill) SetSkillId(v int64)`

SetSkillId sets SkillId field to given value.


### GetStartDate

`func (o *CharactersSkillqueueSkill) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CharactersSkillqueueSkill) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CharactersSkillqueueSkill) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CharactersSkillqueueSkill) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTrainingStartSp

`func (o *CharactersSkillqueueSkill) GetTrainingStartSp() int64`

GetTrainingStartSp returns the TrainingStartSp field if non-nil, zero value otherwise.

### GetTrainingStartSpOk

`func (o *CharactersSkillqueueSkill) GetTrainingStartSpOk() (*int64, bool)`

GetTrainingStartSpOk returns a tuple with the TrainingStartSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingStartSp

`func (o *CharactersSkillqueueSkill) SetTrainingStartSp(v int64)`

SetTrainingStartSp sets TrainingStartSp field to given value.

### HasTrainingStartSp

`func (o *CharactersSkillqueueSkill) HasTrainingStartSp() bool`

HasTrainingStartSp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


