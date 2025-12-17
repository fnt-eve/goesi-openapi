# FreelanceJobsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessAndVisibility** | [**FreelanceJobsDetailAccessandvisibility**](FreelanceJobsDetailAccessandvisibility.md) |  | 
**Configuration** | [**FreelanceJobsDetailConfiguration**](FreelanceJobsDetailConfiguration.md) |  | 
**Contribution** | Pointer to [**FreelanceJobsDetailContribution**](FreelanceJobsDetailContribution.md) |  | [optional] 
**Details** | [**FreelanceJobsDetailDetails**](FreelanceJobsDetailDetails.md) |  | 
**Id** | **string** |  | 
**LastModified** | **time.Time** | Job&#39;s last modified | 
**Name** | **string** | Job&#39;s name | 
**Progress** | [**FreelanceJobsDetailProgress**](FreelanceJobsDetailProgress.md) |  | 
**Reward** | Pointer to [**FreelanceJobsDetailReward**](FreelanceJobsDetailReward.md) |  | [optional] 
**State** | **string** | Job&#39;s state | 

## Methods

### NewFreelanceJobsDetail

`func NewFreelanceJobsDetail(accessAndVisibility FreelanceJobsDetailAccessandvisibility, configuration FreelanceJobsDetailConfiguration, details FreelanceJobsDetailDetails, id string, lastModified time.Time, name string, progress FreelanceJobsDetailProgress, state string, ) *FreelanceJobsDetail`

NewFreelanceJobsDetail instantiates a new FreelanceJobsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreelanceJobsDetailWithDefaults

`func NewFreelanceJobsDetailWithDefaults() *FreelanceJobsDetail`

NewFreelanceJobsDetailWithDefaults instantiates a new FreelanceJobsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessAndVisibility

`func (o *FreelanceJobsDetail) GetAccessAndVisibility() FreelanceJobsDetailAccessandvisibility`

GetAccessAndVisibility returns the AccessAndVisibility field if non-nil, zero value otherwise.

### GetAccessAndVisibilityOk

`func (o *FreelanceJobsDetail) GetAccessAndVisibilityOk() (*FreelanceJobsDetailAccessandvisibility, bool)`

GetAccessAndVisibilityOk returns a tuple with the AccessAndVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndVisibility

`func (o *FreelanceJobsDetail) SetAccessAndVisibility(v FreelanceJobsDetailAccessandvisibility)`

SetAccessAndVisibility sets AccessAndVisibility field to given value.


### GetConfiguration

`func (o *FreelanceJobsDetail) GetConfiguration() FreelanceJobsDetailConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FreelanceJobsDetail) GetConfigurationOk() (*FreelanceJobsDetailConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FreelanceJobsDetail) SetConfiguration(v FreelanceJobsDetailConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetContribution

`func (o *FreelanceJobsDetail) GetContribution() FreelanceJobsDetailContribution`

GetContribution returns the Contribution field if non-nil, zero value otherwise.

### GetContributionOk

`func (o *FreelanceJobsDetail) GetContributionOk() (*FreelanceJobsDetailContribution, bool)`

GetContributionOk returns a tuple with the Contribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContribution

`func (o *FreelanceJobsDetail) SetContribution(v FreelanceJobsDetailContribution)`

SetContribution sets Contribution field to given value.

### HasContribution

`func (o *FreelanceJobsDetail) HasContribution() bool`

HasContribution returns a boolean if a field has been set.

### GetDetails

`func (o *FreelanceJobsDetail) GetDetails() FreelanceJobsDetailDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FreelanceJobsDetail) GetDetailsOk() (*FreelanceJobsDetailDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FreelanceJobsDetail) SetDetails(v FreelanceJobsDetailDetails)`

SetDetails sets Details field to given value.


### GetId

`func (o *FreelanceJobsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FreelanceJobsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FreelanceJobsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetLastModified

`func (o *FreelanceJobsDetail) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *FreelanceJobsDetail) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *FreelanceJobsDetail) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetName

`func (o *FreelanceJobsDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FreelanceJobsDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FreelanceJobsDetail) SetName(v string)`

SetName sets Name field to given value.


### GetProgress

`func (o *FreelanceJobsDetail) GetProgress() FreelanceJobsDetailProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FreelanceJobsDetail) GetProgressOk() (*FreelanceJobsDetailProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FreelanceJobsDetail) SetProgress(v FreelanceJobsDetailProgress)`

SetProgress sets Progress field to given value.


### GetReward

`func (o *FreelanceJobsDetail) GetReward() FreelanceJobsDetailReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *FreelanceJobsDetail) GetRewardOk() (*FreelanceJobsDetailReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *FreelanceJobsDetail) SetReward(v FreelanceJobsDetailReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *FreelanceJobsDetail) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetState

`func (o *FreelanceJobsDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FreelanceJobsDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FreelanceJobsDetail) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


