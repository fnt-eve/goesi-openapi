# CorporationsProjectsDetailProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LastModified** | **time.Time** | Moment this project was last modified. Project contributions also count as a modification | 
**Name** | **string** | Project&#39;s name | 
**Progress** | [**CorporationsProjectsDetailProgress**](CorporationsProjectsDetailProgress.md) |  | 
**Reward** | Pointer to [**CorporationsProjectsDetailReward**](CorporationsProjectsDetailReward.md) |  | [optional] 
**State** | **string** | Project&#39;s current state | 

## Methods

### NewCorporationsProjectsDetailProject

`func NewCorporationsProjectsDetailProject(id string, lastModified time.Time, name string, progress CorporationsProjectsDetailProgress, state string, ) *CorporationsProjectsDetailProject`

NewCorporationsProjectsDetailProject instantiates a new CorporationsProjectsDetailProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsProjectsDetailProjectWithDefaults

`func NewCorporationsProjectsDetailProjectWithDefaults() *CorporationsProjectsDetailProject`

NewCorporationsProjectsDetailProjectWithDefaults instantiates a new CorporationsProjectsDetailProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CorporationsProjectsDetailProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporationsProjectsDetailProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporationsProjectsDetailProject) SetId(v string)`

SetId sets Id field to given value.


### GetLastModified

`func (o *CorporationsProjectsDetailProject) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CorporationsProjectsDetailProject) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CorporationsProjectsDetailProject) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetName

`func (o *CorporationsProjectsDetailProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporationsProjectsDetailProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporationsProjectsDetailProject) SetName(v string)`

SetName sets Name field to given value.


### GetProgress

`func (o *CorporationsProjectsDetailProject) GetProgress() CorporationsProjectsDetailProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *CorporationsProjectsDetailProject) GetProgressOk() (*CorporationsProjectsDetailProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *CorporationsProjectsDetailProject) SetProgress(v CorporationsProjectsDetailProgress)`

SetProgress sets Progress field to given value.


### GetReward

`func (o *CorporationsProjectsDetailProject) GetReward() CorporationsProjectsDetailReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *CorporationsProjectsDetailProject) GetRewardOk() (*CorporationsProjectsDetailReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *CorporationsProjectsDetailProject) SetReward(v CorporationsProjectsDetailReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *CorporationsProjectsDetailProject) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetState

`func (o *CorporationsProjectsDetailProject) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CorporationsProjectsDetailProject) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CorporationsProjectsDetailProject) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


