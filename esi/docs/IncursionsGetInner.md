# IncursionsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstellationId** | **int64** | The constellation id in which this incursion takes place | 
**FactionId** | **int64** | The attacking faction&#39;s id | 
**HasBoss** | **bool** | Whether the final encounter has boss or not | 
**InfestedSolarSystems** | **[]int64** | A list of infested solar system ids that are a part of this incursion | 
**Influence** | **float64** | Influence of this incursion as a float from 0 to 1 | 
**StagingSolarSystemId** | **int64** | Staging solar system for this incursion | 
**State** | **string** | The state of this incursion | 
**Type** | **string** | The type of this incursion | 

## Methods

### NewIncursionsGetInner

`func NewIncursionsGetInner(constellationId int64, factionId int64, hasBoss bool, infestedSolarSystems []int64, influence float64, stagingSolarSystemId int64, state string, type_ string, ) *IncursionsGetInner`

NewIncursionsGetInner instantiates a new IncursionsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncursionsGetInnerWithDefaults

`func NewIncursionsGetInnerWithDefaults() *IncursionsGetInner`

NewIncursionsGetInnerWithDefaults instantiates a new IncursionsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstellationId

`func (o *IncursionsGetInner) GetConstellationId() int64`

GetConstellationId returns the ConstellationId field if non-nil, zero value otherwise.

### GetConstellationIdOk

`func (o *IncursionsGetInner) GetConstellationIdOk() (*int64, bool)`

GetConstellationIdOk returns a tuple with the ConstellationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstellationId

`func (o *IncursionsGetInner) SetConstellationId(v int64)`

SetConstellationId sets ConstellationId field to given value.


### GetFactionId

`func (o *IncursionsGetInner) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *IncursionsGetInner) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *IncursionsGetInner) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.


### GetHasBoss

`func (o *IncursionsGetInner) GetHasBoss() bool`

GetHasBoss returns the HasBoss field if non-nil, zero value otherwise.

### GetHasBossOk

`func (o *IncursionsGetInner) GetHasBossOk() (*bool, bool)`

GetHasBossOk returns a tuple with the HasBoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBoss

`func (o *IncursionsGetInner) SetHasBoss(v bool)`

SetHasBoss sets HasBoss field to given value.


### GetInfestedSolarSystems

`func (o *IncursionsGetInner) GetInfestedSolarSystems() []int64`

GetInfestedSolarSystems returns the InfestedSolarSystems field if non-nil, zero value otherwise.

### GetInfestedSolarSystemsOk

`func (o *IncursionsGetInner) GetInfestedSolarSystemsOk() (*[]int64, bool)`

GetInfestedSolarSystemsOk returns a tuple with the InfestedSolarSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfestedSolarSystems

`func (o *IncursionsGetInner) SetInfestedSolarSystems(v []int64)`

SetInfestedSolarSystems sets InfestedSolarSystems field to given value.


### GetInfluence

`func (o *IncursionsGetInner) GetInfluence() float64`

GetInfluence returns the Influence field if non-nil, zero value otherwise.

### GetInfluenceOk

`func (o *IncursionsGetInner) GetInfluenceOk() (*float64, bool)`

GetInfluenceOk returns a tuple with the Influence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluence

`func (o *IncursionsGetInner) SetInfluence(v float64)`

SetInfluence sets Influence field to given value.


### GetStagingSolarSystemId

`func (o *IncursionsGetInner) GetStagingSolarSystemId() int64`

GetStagingSolarSystemId returns the StagingSolarSystemId field if non-nil, zero value otherwise.

### GetStagingSolarSystemIdOk

`func (o *IncursionsGetInner) GetStagingSolarSystemIdOk() (*int64, bool)`

GetStagingSolarSystemIdOk returns a tuple with the StagingSolarSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingSolarSystemId

`func (o *IncursionsGetInner) SetStagingSolarSystemId(v int64)`

SetStagingSolarSystemId sets StagingSolarSystemId field to given value.


### GetState

`func (o *IncursionsGetInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IncursionsGetInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IncursionsGetInner) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *IncursionsGetInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncursionsGetInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncursionsGetInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


