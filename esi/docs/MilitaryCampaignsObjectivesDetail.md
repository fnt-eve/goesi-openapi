# MilitaryCampaignsObjectivesDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finished** | Pointer to **time.Time** | Moment the objective transitioned to a non-active state | [optional] 
**Id** | **string** |  | 
**LastModified** | **time.Time** | Objective&#39;s last modified | 
**Participants** | [**MilitaryCampaignsObjectivesDetailParticipants**](MilitaryCampaignsObjectivesDetailParticipants.md) |  | 
**Progress** | **int64** | Objective&#39;s progress | 
**Started** | Pointer to **time.Time** | Moment the objective started | [optional] 
**State** | **string** | Objective&#39;s state | 

## Methods

### NewMilitaryCampaignsObjectivesDetail

`func NewMilitaryCampaignsObjectivesDetail(id string, lastModified time.Time, participants MilitaryCampaignsObjectivesDetailParticipants, progress int64, state string, ) *MilitaryCampaignsObjectivesDetail`

NewMilitaryCampaignsObjectivesDetail instantiates a new MilitaryCampaignsObjectivesDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilitaryCampaignsObjectivesDetailWithDefaults

`func NewMilitaryCampaignsObjectivesDetailWithDefaults() *MilitaryCampaignsObjectivesDetail`

NewMilitaryCampaignsObjectivesDetailWithDefaults instantiates a new MilitaryCampaignsObjectivesDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinished

`func (o *MilitaryCampaignsObjectivesDetail) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *MilitaryCampaignsObjectivesDetail) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *MilitaryCampaignsObjectivesDetail) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *MilitaryCampaignsObjectivesDetail) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *MilitaryCampaignsObjectivesDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MilitaryCampaignsObjectivesDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MilitaryCampaignsObjectivesDetail) SetId(v string)`

SetId sets Id field to given value.


### GetLastModified

`func (o *MilitaryCampaignsObjectivesDetail) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *MilitaryCampaignsObjectivesDetail) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *MilitaryCampaignsObjectivesDetail) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetParticipants

`func (o *MilitaryCampaignsObjectivesDetail) GetParticipants() MilitaryCampaignsObjectivesDetailParticipants`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *MilitaryCampaignsObjectivesDetail) GetParticipantsOk() (*MilitaryCampaignsObjectivesDetailParticipants, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *MilitaryCampaignsObjectivesDetail) SetParticipants(v MilitaryCampaignsObjectivesDetailParticipants)`

SetParticipants sets Participants field to given value.


### GetProgress

`func (o *MilitaryCampaignsObjectivesDetail) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *MilitaryCampaignsObjectivesDetail) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *MilitaryCampaignsObjectivesDetail) SetProgress(v int64)`

SetProgress sets Progress field to given value.


### GetStarted

`func (o *MilitaryCampaignsObjectivesDetail) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *MilitaryCampaignsObjectivesDetail) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *MilitaryCampaignsObjectivesDetail) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *MilitaryCampaignsObjectivesDetail) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetState

`func (o *MilitaryCampaignsObjectivesDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MilitaryCampaignsObjectivesDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MilitaryCampaignsObjectivesDetail) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


