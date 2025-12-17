# FreelanceJobsDetailDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Career** | **string** | Assigned career path | 
**Created** | **time.Time** | Moment this freelance job was created | 
**Creator** | [**FreelanceJobsDetailCreator**](FreelanceJobsDetailCreator.md) |  | 
**Description** | **string** | Description | 
**Expires** | Pointer to **time.Time** | Moment this freelance job expires | [optional] 
**Finished** | Pointer to **time.Time** | Moment this freelance job transitioned to a non-active state | [optional] 

## Methods

### NewFreelanceJobsDetailDetails

`func NewFreelanceJobsDetailDetails(career string, created time.Time, creator FreelanceJobsDetailCreator, description string, ) *FreelanceJobsDetailDetails`

NewFreelanceJobsDetailDetails instantiates a new FreelanceJobsDetailDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreelanceJobsDetailDetailsWithDefaults

`func NewFreelanceJobsDetailDetailsWithDefaults() *FreelanceJobsDetailDetails`

NewFreelanceJobsDetailDetailsWithDefaults instantiates a new FreelanceJobsDetailDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCareer

`func (o *FreelanceJobsDetailDetails) GetCareer() string`

GetCareer returns the Career field if non-nil, zero value otherwise.

### GetCareerOk

`func (o *FreelanceJobsDetailDetails) GetCareerOk() (*string, bool)`

GetCareerOk returns a tuple with the Career field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareer

`func (o *FreelanceJobsDetailDetails) SetCareer(v string)`

SetCareer sets Career field to given value.


### GetCreated

`func (o *FreelanceJobsDetailDetails) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FreelanceJobsDetailDetails) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FreelanceJobsDetailDetails) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreator

`func (o *FreelanceJobsDetailDetails) GetCreator() FreelanceJobsDetailCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *FreelanceJobsDetailDetails) GetCreatorOk() (*FreelanceJobsDetailCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *FreelanceJobsDetailDetails) SetCreator(v FreelanceJobsDetailCreator)`

SetCreator sets Creator field to given value.


### GetDescription

`func (o *FreelanceJobsDetailDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FreelanceJobsDetailDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FreelanceJobsDetailDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpires

`func (o *FreelanceJobsDetailDetails) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *FreelanceJobsDetailDetails) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *FreelanceJobsDetailDetails) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *FreelanceJobsDetailDetails) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFinished

`func (o *FreelanceJobsDetailDetails) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *FreelanceJobsDetailDetails) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *FreelanceJobsDetailDetails) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *FreelanceJobsDetailDetails) HasFinished() bool`

HasFinished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


