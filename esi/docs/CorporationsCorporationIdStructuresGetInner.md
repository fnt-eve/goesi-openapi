# CorporationsCorporationIdStructuresGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorporationId** | **int64** | ID of the corporation that owns the structure | 
**FuelExpires** | Pointer to **time.Time** | Date on which the structure will run out of fuel | [optional] 
**Name** | Pointer to **string** | The structure name | [optional] 
**NextReinforceApply** | Pointer to **time.Time** | The date and time when the structure&#39;s newly requested reinforcement times (e.g. next_reinforce_hour and next_reinforce_day) will take effect | [optional] 
**NextReinforceHour** | Pointer to **int64** | The requested change to reinforce_hour that will take effect at the time shown by next_reinforce_apply | [optional] 
**ProfileId** | **int64** | The id of the ACL profile for this citadel | 
**ReinforceHour** | Pointer to **int64** | The hour of day that determines the four hour window when the structure will randomly exit its reinforcement periods and become vulnerable to attack against its armor and/or hull. The structure will become vulnerable at a random time that is +/- 2 hours centered on the value of this property | [optional] 
**Services** | Pointer to [**[]CorporationsCorporationIdStructuresGetInnerServicesInner**](CorporationsCorporationIdStructuresGetInnerServicesInner.md) | Contains a list of service upgrades, and their state | [optional] 
**State** | **string** |  | 
**StateTimerEnd** | Pointer to **time.Time** | Date at which the structure will move to it&#39;s next state | [optional] 
**StateTimerStart** | Pointer to **time.Time** | Date at which the structure entered it&#39;s current state | [optional] 
**StructureId** | **int64** | The Item ID of the structure | 
**SystemId** | **int64** | The solar system the structure is in | 
**TypeId** | **int64** | The type id of the structure | 
**UnanchorsAt** | Pointer to **time.Time** | Date at which the structure will unanchor | [optional] 

## Methods

### NewCorporationsCorporationIdStructuresGetInner

`func NewCorporationsCorporationIdStructuresGetInner(corporationId int64, profileId int64, state string, structureId int64, systemId int64, typeId int64, ) *CorporationsCorporationIdStructuresGetInner`

NewCorporationsCorporationIdStructuresGetInner instantiates a new CorporationsCorporationIdStructuresGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdStructuresGetInnerWithDefaults

`func NewCorporationsCorporationIdStructuresGetInnerWithDefaults() *CorporationsCorporationIdStructuresGetInner`

NewCorporationsCorporationIdStructuresGetInnerWithDefaults instantiates a new CorporationsCorporationIdStructuresGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorporationId

`func (o *CorporationsCorporationIdStructuresGetInner) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *CorporationsCorporationIdStructuresGetInner) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.


### GetFuelExpires

`func (o *CorporationsCorporationIdStructuresGetInner) GetFuelExpires() time.Time`

GetFuelExpires returns the FuelExpires field if non-nil, zero value otherwise.

### GetFuelExpiresOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetFuelExpiresOk() (*time.Time, bool)`

GetFuelExpiresOk returns a tuple with the FuelExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelExpires

`func (o *CorporationsCorporationIdStructuresGetInner) SetFuelExpires(v time.Time)`

SetFuelExpires sets FuelExpires field to given value.

### HasFuelExpires

`func (o *CorporationsCorporationIdStructuresGetInner) HasFuelExpires() bool`

HasFuelExpires returns a boolean if a field has been set.

### GetName

`func (o *CorporationsCorporationIdStructuresGetInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporationsCorporationIdStructuresGetInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CorporationsCorporationIdStructuresGetInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextReinforceApply

`func (o *CorporationsCorporationIdStructuresGetInner) GetNextReinforceApply() time.Time`

GetNextReinforceApply returns the NextReinforceApply field if non-nil, zero value otherwise.

### GetNextReinforceApplyOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetNextReinforceApplyOk() (*time.Time, bool)`

GetNextReinforceApplyOk returns a tuple with the NextReinforceApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReinforceApply

`func (o *CorporationsCorporationIdStructuresGetInner) SetNextReinforceApply(v time.Time)`

SetNextReinforceApply sets NextReinforceApply field to given value.

### HasNextReinforceApply

`func (o *CorporationsCorporationIdStructuresGetInner) HasNextReinforceApply() bool`

HasNextReinforceApply returns a boolean if a field has been set.

### GetNextReinforceHour

`func (o *CorporationsCorporationIdStructuresGetInner) GetNextReinforceHour() int64`

GetNextReinforceHour returns the NextReinforceHour field if non-nil, zero value otherwise.

### GetNextReinforceHourOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetNextReinforceHourOk() (*int64, bool)`

GetNextReinforceHourOk returns a tuple with the NextReinforceHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReinforceHour

`func (o *CorporationsCorporationIdStructuresGetInner) SetNextReinforceHour(v int64)`

SetNextReinforceHour sets NextReinforceHour field to given value.

### HasNextReinforceHour

`func (o *CorporationsCorporationIdStructuresGetInner) HasNextReinforceHour() bool`

HasNextReinforceHour returns a boolean if a field has been set.

### GetProfileId

`func (o *CorporationsCorporationIdStructuresGetInner) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CorporationsCorporationIdStructuresGetInner) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.


### GetReinforceHour

`func (o *CorporationsCorporationIdStructuresGetInner) GetReinforceHour() int64`

GetReinforceHour returns the ReinforceHour field if non-nil, zero value otherwise.

### GetReinforceHourOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetReinforceHourOk() (*int64, bool)`

GetReinforceHourOk returns a tuple with the ReinforceHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinforceHour

`func (o *CorporationsCorporationIdStructuresGetInner) SetReinforceHour(v int64)`

SetReinforceHour sets ReinforceHour field to given value.

### HasReinforceHour

`func (o *CorporationsCorporationIdStructuresGetInner) HasReinforceHour() bool`

HasReinforceHour returns a boolean if a field has been set.

### GetServices

`func (o *CorporationsCorporationIdStructuresGetInner) GetServices() []CorporationsCorporationIdStructuresGetInnerServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetServicesOk() (*[]CorporationsCorporationIdStructuresGetInnerServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CorporationsCorporationIdStructuresGetInner) SetServices(v []CorporationsCorporationIdStructuresGetInnerServicesInner)`

SetServices sets Services field to given value.

### HasServices

`func (o *CorporationsCorporationIdStructuresGetInner) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetState

`func (o *CorporationsCorporationIdStructuresGetInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CorporationsCorporationIdStructuresGetInner) SetState(v string)`

SetState sets State field to given value.


### GetStateTimerEnd

`func (o *CorporationsCorporationIdStructuresGetInner) GetStateTimerEnd() time.Time`

GetStateTimerEnd returns the StateTimerEnd field if non-nil, zero value otherwise.

### GetStateTimerEndOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetStateTimerEndOk() (*time.Time, bool)`

GetStateTimerEndOk returns a tuple with the StateTimerEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTimerEnd

`func (o *CorporationsCorporationIdStructuresGetInner) SetStateTimerEnd(v time.Time)`

SetStateTimerEnd sets StateTimerEnd field to given value.

### HasStateTimerEnd

`func (o *CorporationsCorporationIdStructuresGetInner) HasStateTimerEnd() bool`

HasStateTimerEnd returns a boolean if a field has been set.

### GetStateTimerStart

`func (o *CorporationsCorporationIdStructuresGetInner) GetStateTimerStart() time.Time`

GetStateTimerStart returns the StateTimerStart field if non-nil, zero value otherwise.

### GetStateTimerStartOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetStateTimerStartOk() (*time.Time, bool)`

GetStateTimerStartOk returns a tuple with the StateTimerStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTimerStart

`func (o *CorporationsCorporationIdStructuresGetInner) SetStateTimerStart(v time.Time)`

SetStateTimerStart sets StateTimerStart field to given value.

### HasStateTimerStart

`func (o *CorporationsCorporationIdStructuresGetInner) HasStateTimerStart() bool`

HasStateTimerStart returns a boolean if a field has been set.

### GetStructureId

`func (o *CorporationsCorporationIdStructuresGetInner) GetStructureId() int64`

GetStructureId returns the StructureId field if non-nil, zero value otherwise.

### GetStructureIdOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetStructureIdOk() (*int64, bool)`

GetStructureIdOk returns a tuple with the StructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureId

`func (o *CorporationsCorporationIdStructuresGetInner) SetStructureId(v int64)`

SetStructureId sets StructureId field to given value.


### GetSystemId

`func (o *CorporationsCorporationIdStructuresGetInner) GetSystemId() int64`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetSystemIdOk() (*int64, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *CorporationsCorporationIdStructuresGetInner) SetSystemId(v int64)`

SetSystemId sets SystemId field to given value.


### GetTypeId

`func (o *CorporationsCorporationIdStructuresGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsCorporationIdStructuresGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetUnanchorsAt

`func (o *CorporationsCorporationIdStructuresGetInner) GetUnanchorsAt() time.Time`

GetUnanchorsAt returns the UnanchorsAt field if non-nil, zero value otherwise.

### GetUnanchorsAtOk

`func (o *CorporationsCorporationIdStructuresGetInner) GetUnanchorsAtOk() (*time.Time, bool)`

GetUnanchorsAtOk returns a tuple with the UnanchorsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnanchorsAt

`func (o *CorporationsCorporationIdStructuresGetInner) SetUnanchorsAt(v time.Time)`

SetUnanchorsAt sets UnanchorsAt field to given value.

### HasUnanchorsAt

`func (o *CorporationsCorporationIdStructuresGetInner) HasUnanchorsAt() bool`

HasUnanchorsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


