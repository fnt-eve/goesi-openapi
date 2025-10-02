# CorporationsProjectsDetailDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Career** | **string** | Assigned career path | 
**Created** | **time.Time** | Moment this project was created | 
**Description** | **string** | Description | 
**Expires** | Pointer to **time.Time** | Moment this project expires | [optional] 
**Finished** | Pointer to **time.Time** | Moment this project transitioned to a non-active state. | [optional] 

## Methods

### NewCorporationsProjectsDetailDetails

`func NewCorporationsProjectsDetailDetails(career string, created time.Time, description string, ) *CorporationsProjectsDetailDetails`

NewCorporationsProjectsDetailDetails instantiates a new CorporationsProjectsDetailDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsProjectsDetailDetailsWithDefaults

`func NewCorporationsProjectsDetailDetailsWithDefaults() *CorporationsProjectsDetailDetails`

NewCorporationsProjectsDetailDetailsWithDefaults instantiates a new CorporationsProjectsDetailDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCareer

`func (o *CorporationsProjectsDetailDetails) GetCareer() string`

GetCareer returns the Career field if non-nil, zero value otherwise.

### GetCareerOk

`func (o *CorporationsProjectsDetailDetails) GetCareerOk() (*string, bool)`

GetCareerOk returns a tuple with the Career field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareer

`func (o *CorporationsProjectsDetailDetails) SetCareer(v string)`

SetCareer sets Career field to given value.


### GetCreated

`func (o *CorporationsProjectsDetailDetails) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CorporationsProjectsDetailDetails) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CorporationsProjectsDetailDetails) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDescription

`func (o *CorporationsProjectsDetailDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CorporationsProjectsDetailDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CorporationsProjectsDetailDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpires

`func (o *CorporationsProjectsDetailDetails) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CorporationsProjectsDetailDetails) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CorporationsProjectsDetailDetails) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CorporationsProjectsDetailDetails) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFinished

`func (o *CorporationsProjectsDetailDetails) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *CorporationsProjectsDetailDetails) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *CorporationsProjectsDetailDetails) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *CorporationsProjectsDetailDetails) HasFinished() bool`

HasFinished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


