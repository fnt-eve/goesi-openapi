# SovereigntyCampaignsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackersScore** | Pointer to **float64** | Score for all attacking parties, only present in Defense Events.  | [optional] 
**CampaignId** | **int64** | Unique ID for this campaign. | 
**ConstellationId** | **int64** | The constellation in which the campaign will take place.  | 
**DefenderId** | Pointer to **int64** | Defending alliance, only present in Defense Events  | [optional] 
**DefenderScore** | Pointer to **float64** | Score for the defending alliance, only present in Defense Events.  | [optional] 
**EventType** | **string** | Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as \&quot;Defense Events\&quot;, station_freeport as \&quot;Freeport Events\&quot;.  | 
**Participants** | Pointer to [**[]SovereigntyCampaignsGetInnerParticipantsInner**](SovereigntyCampaignsGetInnerParticipantsInner.md) | Alliance participating and their respective scores, only present in Freeport Events.  | [optional] 
**SolarSystemId** | **int64** | The solar system the structure is located in.  | 
**StartTime** | **time.Time** | Time the event is scheduled to start.  | 
**StructureId** | **int64** | The structure item ID that is related to this campaign.  | 

## Methods

### NewSovereigntyCampaignsGetInner

`func NewSovereigntyCampaignsGetInner(campaignId int64, constellationId int64, eventType string, solarSystemId int64, startTime time.Time, structureId int64, ) *SovereigntyCampaignsGetInner`

NewSovereigntyCampaignsGetInner instantiates a new SovereigntyCampaignsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSovereigntyCampaignsGetInnerWithDefaults

`func NewSovereigntyCampaignsGetInnerWithDefaults() *SovereigntyCampaignsGetInner`

NewSovereigntyCampaignsGetInnerWithDefaults instantiates a new SovereigntyCampaignsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackersScore

`func (o *SovereigntyCampaignsGetInner) GetAttackersScore() float64`

GetAttackersScore returns the AttackersScore field if non-nil, zero value otherwise.

### GetAttackersScoreOk

`func (o *SovereigntyCampaignsGetInner) GetAttackersScoreOk() (*float64, bool)`

GetAttackersScoreOk returns a tuple with the AttackersScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackersScore

`func (o *SovereigntyCampaignsGetInner) SetAttackersScore(v float64)`

SetAttackersScore sets AttackersScore field to given value.

### HasAttackersScore

`func (o *SovereigntyCampaignsGetInner) HasAttackersScore() bool`

HasAttackersScore returns a boolean if a field has been set.

### GetCampaignId

`func (o *SovereigntyCampaignsGetInner) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *SovereigntyCampaignsGetInner) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *SovereigntyCampaignsGetInner) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetConstellationId

`func (o *SovereigntyCampaignsGetInner) GetConstellationId() int64`

GetConstellationId returns the ConstellationId field if non-nil, zero value otherwise.

### GetConstellationIdOk

`func (o *SovereigntyCampaignsGetInner) GetConstellationIdOk() (*int64, bool)`

GetConstellationIdOk returns a tuple with the ConstellationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstellationId

`func (o *SovereigntyCampaignsGetInner) SetConstellationId(v int64)`

SetConstellationId sets ConstellationId field to given value.


### GetDefenderId

`func (o *SovereigntyCampaignsGetInner) GetDefenderId() int64`

GetDefenderId returns the DefenderId field if non-nil, zero value otherwise.

### GetDefenderIdOk

`func (o *SovereigntyCampaignsGetInner) GetDefenderIdOk() (*int64, bool)`

GetDefenderIdOk returns a tuple with the DefenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenderId

`func (o *SovereigntyCampaignsGetInner) SetDefenderId(v int64)`

SetDefenderId sets DefenderId field to given value.

### HasDefenderId

`func (o *SovereigntyCampaignsGetInner) HasDefenderId() bool`

HasDefenderId returns a boolean if a field has been set.

### GetDefenderScore

`func (o *SovereigntyCampaignsGetInner) GetDefenderScore() float64`

GetDefenderScore returns the DefenderScore field if non-nil, zero value otherwise.

### GetDefenderScoreOk

`func (o *SovereigntyCampaignsGetInner) GetDefenderScoreOk() (*float64, bool)`

GetDefenderScoreOk returns a tuple with the DefenderScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenderScore

`func (o *SovereigntyCampaignsGetInner) SetDefenderScore(v float64)`

SetDefenderScore sets DefenderScore field to given value.

### HasDefenderScore

`func (o *SovereigntyCampaignsGetInner) HasDefenderScore() bool`

HasDefenderScore returns a boolean if a field has been set.

### GetEventType

`func (o *SovereigntyCampaignsGetInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SovereigntyCampaignsGetInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SovereigntyCampaignsGetInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetParticipants

`func (o *SovereigntyCampaignsGetInner) GetParticipants() []SovereigntyCampaignsGetInnerParticipantsInner`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *SovereigntyCampaignsGetInner) GetParticipantsOk() (*[]SovereigntyCampaignsGetInnerParticipantsInner, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *SovereigntyCampaignsGetInner) SetParticipants(v []SovereigntyCampaignsGetInnerParticipantsInner)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *SovereigntyCampaignsGetInner) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetSolarSystemId

`func (o *SovereigntyCampaignsGetInner) GetSolarSystemId() int64`

GetSolarSystemId returns the SolarSystemId field if non-nil, zero value otherwise.

### GetSolarSystemIdOk

`func (o *SovereigntyCampaignsGetInner) GetSolarSystemIdOk() (*int64, bool)`

GetSolarSystemIdOk returns a tuple with the SolarSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarSystemId

`func (o *SovereigntyCampaignsGetInner) SetSolarSystemId(v int64)`

SetSolarSystemId sets SolarSystemId field to given value.


### GetStartTime

`func (o *SovereigntyCampaignsGetInner) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SovereigntyCampaignsGetInner) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SovereigntyCampaignsGetInner) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetStructureId

`func (o *SovereigntyCampaignsGetInner) GetStructureId() int64`

GetStructureId returns the StructureId field if non-nil, zero value otherwise.

### GetStructureIdOk

`func (o *SovereigntyCampaignsGetInner) GetStructureIdOk() (*int64, bool)`

GetStructureIdOk returns a tuple with the StructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureId

`func (o *SovereigntyCampaignsGetInner) SetStructureId(v int64)`

SetStructureId sets StructureId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


