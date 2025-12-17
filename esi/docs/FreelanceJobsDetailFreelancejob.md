# FreelanceJobsDetailFreelancejob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LastModified** | **time.Time** | Job&#39;s last modified | 
**Name** | **string** | Job&#39;s name | 
**Progress** | [**FreelanceJobsDetailProgress**](FreelanceJobsDetailProgress.md) |  | 
**Reward** | Pointer to [**FreelanceJobsDetailReward**](FreelanceJobsDetailReward.md) |  | [optional] 
**State** | **string** | Job&#39;s state | 

## Methods

### NewFreelanceJobsDetailFreelancejob

`func NewFreelanceJobsDetailFreelancejob(id string, lastModified time.Time, name string, progress FreelanceJobsDetailProgress, state string, ) *FreelanceJobsDetailFreelancejob`

NewFreelanceJobsDetailFreelancejob instantiates a new FreelanceJobsDetailFreelancejob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreelanceJobsDetailFreelancejobWithDefaults

`func NewFreelanceJobsDetailFreelancejobWithDefaults() *FreelanceJobsDetailFreelancejob`

NewFreelanceJobsDetailFreelancejobWithDefaults instantiates a new FreelanceJobsDetailFreelancejob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FreelanceJobsDetailFreelancejob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FreelanceJobsDetailFreelancejob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FreelanceJobsDetailFreelancejob) SetId(v string)`

SetId sets Id field to given value.


### GetLastModified

`func (o *FreelanceJobsDetailFreelancejob) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *FreelanceJobsDetailFreelancejob) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *FreelanceJobsDetailFreelancejob) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetName

`func (o *FreelanceJobsDetailFreelancejob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FreelanceJobsDetailFreelancejob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FreelanceJobsDetailFreelancejob) SetName(v string)`

SetName sets Name field to given value.


### GetProgress

`func (o *FreelanceJobsDetailFreelancejob) GetProgress() FreelanceJobsDetailProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FreelanceJobsDetailFreelancejob) GetProgressOk() (*FreelanceJobsDetailProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FreelanceJobsDetailFreelancejob) SetProgress(v FreelanceJobsDetailProgress)`

SetProgress sets Progress field to given value.


### GetReward

`func (o *FreelanceJobsDetailFreelancejob) GetReward() FreelanceJobsDetailReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *FreelanceJobsDetailFreelancejob) GetRewardOk() (*FreelanceJobsDetailReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *FreelanceJobsDetailFreelancejob) SetReward(v FreelanceJobsDetailReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *FreelanceJobsDetailFreelancejob) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetState

`func (o *FreelanceJobsDetailFreelancejob) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FreelanceJobsDetailFreelancejob) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FreelanceJobsDetailFreelancejob) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


