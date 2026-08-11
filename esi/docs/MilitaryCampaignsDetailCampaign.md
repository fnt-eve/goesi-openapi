# MilitaryCampaignsDetailCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finished** | Pointer to **time.Time** | Moment the campaign transitioned to a non-active state | [optional] 
**Id** | **string** |  | 
**Progress** | **int64** | Campaign&#39;s progress | 
**Started** | Pointer to **time.Time** | Moment the campaign started | [optional] 
**State** | **string** | Campaign&#39;s state | 

## Methods

### NewMilitaryCampaignsDetailCampaign

`func NewMilitaryCampaignsDetailCampaign(id string, progress int64, state string, ) *MilitaryCampaignsDetailCampaign`

NewMilitaryCampaignsDetailCampaign instantiates a new MilitaryCampaignsDetailCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilitaryCampaignsDetailCampaignWithDefaults

`func NewMilitaryCampaignsDetailCampaignWithDefaults() *MilitaryCampaignsDetailCampaign`

NewMilitaryCampaignsDetailCampaignWithDefaults instantiates a new MilitaryCampaignsDetailCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinished

`func (o *MilitaryCampaignsDetailCampaign) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *MilitaryCampaignsDetailCampaign) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *MilitaryCampaignsDetailCampaign) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *MilitaryCampaignsDetailCampaign) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *MilitaryCampaignsDetailCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MilitaryCampaignsDetailCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MilitaryCampaignsDetailCampaign) SetId(v string)`

SetId sets Id field to given value.


### GetProgress

`func (o *MilitaryCampaignsDetailCampaign) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *MilitaryCampaignsDetailCampaign) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *MilitaryCampaignsDetailCampaign) SetProgress(v int64)`

SetProgress sets Progress field to given value.


### GetStarted

`func (o *MilitaryCampaignsDetailCampaign) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *MilitaryCampaignsDetailCampaign) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *MilitaryCampaignsDetailCampaign) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *MilitaryCampaignsDetailCampaign) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetState

`func (o *MilitaryCampaignsDetailCampaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MilitaryCampaignsDetailCampaign) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MilitaryCampaignsDetailCampaign) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


