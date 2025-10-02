# CorporationsProjectsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [**CorporationsProjectsDetailConfiguration**](CorporationsProjectsDetailConfiguration.md) |  | 
**Contribution** | Pointer to [**CorporationsProjectsDetailContribution**](CorporationsProjectsDetailContribution.md) |  | [optional] 
**Creator** | [**CorporationsProjectsDetailCreator**](CorporationsProjectsDetailCreator.md) |  | 
**Details** | [**CorporationsProjectsDetailDetails**](CorporationsProjectsDetailDetails.md) |  | 
**Id** | **string** |  | 
**LastModified** | **time.Time** | Moment this project was last modified. Project contributions also count as a modification | 
**Name** | **string** | Project&#39;s name | 
**Progress** | [**CorporationsProjectsDetailProgress**](CorporationsProjectsDetailProgress.md) |  | 
**Reward** | Pointer to [**CorporationsProjectsDetailReward**](CorporationsProjectsDetailReward.md) |  | [optional] 
**State** | **string** | Project&#39;s current state | 

## Methods

### NewCorporationsProjectsDetail

`func NewCorporationsProjectsDetail(configuration CorporationsProjectsDetailConfiguration, creator CorporationsProjectsDetailCreator, details CorporationsProjectsDetailDetails, id string, lastModified time.Time, name string, progress CorporationsProjectsDetailProgress, state string, ) *CorporationsProjectsDetail`

NewCorporationsProjectsDetail instantiates a new CorporationsProjectsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsProjectsDetailWithDefaults

`func NewCorporationsProjectsDetailWithDefaults() *CorporationsProjectsDetail`

NewCorporationsProjectsDetailWithDefaults instantiates a new CorporationsProjectsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *CorporationsProjectsDetail) GetConfiguration() CorporationsProjectsDetailConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CorporationsProjectsDetail) GetConfigurationOk() (*CorporationsProjectsDetailConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CorporationsProjectsDetail) SetConfiguration(v CorporationsProjectsDetailConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetContribution

`func (o *CorporationsProjectsDetail) GetContribution() CorporationsProjectsDetailContribution`

GetContribution returns the Contribution field if non-nil, zero value otherwise.

### GetContributionOk

`func (o *CorporationsProjectsDetail) GetContributionOk() (*CorporationsProjectsDetailContribution, bool)`

GetContributionOk returns a tuple with the Contribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContribution

`func (o *CorporationsProjectsDetail) SetContribution(v CorporationsProjectsDetailContribution)`

SetContribution sets Contribution field to given value.

### HasContribution

`func (o *CorporationsProjectsDetail) HasContribution() bool`

HasContribution returns a boolean if a field has been set.

### GetCreator

`func (o *CorporationsProjectsDetail) GetCreator() CorporationsProjectsDetailCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CorporationsProjectsDetail) GetCreatorOk() (*CorporationsProjectsDetailCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CorporationsProjectsDetail) SetCreator(v CorporationsProjectsDetailCreator)`

SetCreator sets Creator field to given value.


### GetDetails

`func (o *CorporationsProjectsDetail) GetDetails() CorporationsProjectsDetailDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CorporationsProjectsDetail) GetDetailsOk() (*CorporationsProjectsDetailDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CorporationsProjectsDetail) SetDetails(v CorporationsProjectsDetailDetails)`

SetDetails sets Details field to given value.


### GetId

`func (o *CorporationsProjectsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporationsProjectsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporationsProjectsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetLastModified

`func (o *CorporationsProjectsDetail) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CorporationsProjectsDetail) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CorporationsProjectsDetail) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetName

`func (o *CorporationsProjectsDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporationsProjectsDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporationsProjectsDetail) SetName(v string)`

SetName sets Name field to given value.


### GetProgress

`func (o *CorporationsProjectsDetail) GetProgress() CorporationsProjectsDetailProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *CorporationsProjectsDetail) GetProgressOk() (*CorporationsProjectsDetailProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *CorporationsProjectsDetail) SetProgress(v CorporationsProjectsDetailProgress)`

SetProgress sets Progress field to given value.


### GetReward

`func (o *CorporationsProjectsDetail) GetReward() CorporationsProjectsDetailReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *CorporationsProjectsDetail) GetRewardOk() (*CorporationsProjectsDetailReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *CorporationsProjectsDetail) SetReward(v CorporationsProjectsDetailReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *CorporationsProjectsDetail) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetState

`func (o *CorporationsProjectsDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CorporationsProjectsDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CorporationsProjectsDetail) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


