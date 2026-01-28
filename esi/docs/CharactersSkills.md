# CharactersSkills

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skills** | [**[]CharactersSkillsSkill**](CharactersSkillsSkill.md) | The trained skills | 
**TotalSp** | **int64** | The total Skill Points spent on skills | 
**UnallocatedSp** | Pointer to **int64** | The amount of unallocated Skill Points | [optional] 

## Methods

### NewCharactersSkills

`func NewCharactersSkills(skills []CharactersSkillsSkill, totalSp int64, ) *CharactersSkills`

NewCharactersSkills instantiates a new CharactersSkills object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersSkillsWithDefaults

`func NewCharactersSkillsWithDefaults() *CharactersSkills`

NewCharactersSkillsWithDefaults instantiates a new CharactersSkills object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkills

`func (o *CharactersSkills) GetSkills() []CharactersSkillsSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *CharactersSkills) GetSkillsOk() (*[]CharactersSkillsSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *CharactersSkills) SetSkills(v []CharactersSkillsSkill)`

SetSkills sets Skills field to given value.


### GetTotalSp

`func (o *CharactersSkills) GetTotalSp() int64`

GetTotalSp returns the TotalSp field if non-nil, zero value otherwise.

### GetTotalSpOk

`func (o *CharactersSkills) GetTotalSpOk() (*int64, bool)`

GetTotalSpOk returns a tuple with the TotalSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSp

`func (o *CharactersSkills) SetTotalSp(v int64)`

SetTotalSp sets TotalSp field to given value.


### GetUnallocatedSp

`func (o *CharactersSkills) GetUnallocatedSp() int64`

GetUnallocatedSp returns the UnallocatedSp field if non-nil, zero value otherwise.

### GetUnallocatedSpOk

`func (o *CharactersSkills) GetUnallocatedSpOk() (*int64, bool)`

GetUnallocatedSpOk returns a tuple with the UnallocatedSp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnallocatedSp

`func (o *CharactersSkills) SetUnallocatedSp(v int64)`

SetUnallocatedSp sets UnallocatedSp field to given value.

### HasUnallocatedSp

`func (o *CharactersSkills) HasUnallocatedSp() bool`

HasUnallocatedSp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


