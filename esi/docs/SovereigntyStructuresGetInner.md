# SovereigntyStructuresGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | **int64** | The alliance that owns the structure.  | 
**SolarSystemId** | **int64** | Solar system in which the structure is located.  | 
**StructureId** | **int64** | Unique item ID for this structure. | 
**StructureTypeId** | **int64** | A reference to the type of structure this is.  | 
**VulnerabilityOccupancyLevel** | Pointer to **float64** | The occupancy level for the next or current vulnerability window. This takes into account all development indexes and capital system bonuses. Also known as Activity Defense Multiplier from in the client. It increases the time that attackers must spend using their entosis links on the structure.  | [optional] 
**VulnerableEndTime** | Pointer to **time.Time** | The time at which the next or current vulnerability window ends. At the end of a vulnerability window the next window is recalculated and locked in along with the vulnerabilityOccupancyLevel. If the structure is not in 100% entosis control of the defender, it will go in to &#39;overtime&#39; and stay vulnerable for as long as that situation persists. Only once the defenders have 100% entosis control and has the vulnerableEndTime passed does the vulnerability interval expire and a new one is calculated.  | [optional] 
**VulnerableStartTime** | Pointer to **time.Time** | The next time at which the structure will become vulnerable. Or the start time of the current window if current time is between this and vulnerableEndTime.  | [optional] 

## Methods

### NewSovereigntyStructuresGetInner

`func NewSovereigntyStructuresGetInner(allianceId int64, solarSystemId int64, structureId int64, structureTypeId int64, ) *SovereigntyStructuresGetInner`

NewSovereigntyStructuresGetInner instantiates a new SovereigntyStructuresGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSovereigntyStructuresGetInnerWithDefaults

`func NewSovereigntyStructuresGetInnerWithDefaults() *SovereigntyStructuresGetInner`

NewSovereigntyStructuresGetInnerWithDefaults instantiates a new SovereigntyStructuresGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *SovereigntyStructuresGetInner) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *SovereigntyStructuresGetInner) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *SovereigntyStructuresGetInner) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.


### GetSolarSystemId

`func (o *SovereigntyStructuresGetInner) GetSolarSystemId() int64`

GetSolarSystemId returns the SolarSystemId field if non-nil, zero value otherwise.

### GetSolarSystemIdOk

`func (o *SovereigntyStructuresGetInner) GetSolarSystemIdOk() (*int64, bool)`

GetSolarSystemIdOk returns a tuple with the SolarSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarSystemId

`func (o *SovereigntyStructuresGetInner) SetSolarSystemId(v int64)`

SetSolarSystemId sets SolarSystemId field to given value.


### GetStructureId

`func (o *SovereigntyStructuresGetInner) GetStructureId() int64`

GetStructureId returns the StructureId field if non-nil, zero value otherwise.

### GetStructureIdOk

`func (o *SovereigntyStructuresGetInner) GetStructureIdOk() (*int64, bool)`

GetStructureIdOk returns a tuple with the StructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureId

`func (o *SovereigntyStructuresGetInner) SetStructureId(v int64)`

SetStructureId sets StructureId field to given value.


### GetStructureTypeId

`func (o *SovereigntyStructuresGetInner) GetStructureTypeId() int64`

GetStructureTypeId returns the StructureTypeId field if non-nil, zero value otherwise.

### GetStructureTypeIdOk

`func (o *SovereigntyStructuresGetInner) GetStructureTypeIdOk() (*int64, bool)`

GetStructureTypeIdOk returns a tuple with the StructureTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureTypeId

`func (o *SovereigntyStructuresGetInner) SetStructureTypeId(v int64)`

SetStructureTypeId sets StructureTypeId field to given value.


### GetVulnerabilityOccupancyLevel

`func (o *SovereigntyStructuresGetInner) GetVulnerabilityOccupancyLevel() float64`

GetVulnerabilityOccupancyLevel returns the VulnerabilityOccupancyLevel field if non-nil, zero value otherwise.

### GetVulnerabilityOccupancyLevelOk

`func (o *SovereigntyStructuresGetInner) GetVulnerabilityOccupancyLevelOk() (*float64, bool)`

GetVulnerabilityOccupancyLevelOk returns a tuple with the VulnerabilityOccupancyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityOccupancyLevel

`func (o *SovereigntyStructuresGetInner) SetVulnerabilityOccupancyLevel(v float64)`

SetVulnerabilityOccupancyLevel sets VulnerabilityOccupancyLevel field to given value.

### HasVulnerabilityOccupancyLevel

`func (o *SovereigntyStructuresGetInner) HasVulnerabilityOccupancyLevel() bool`

HasVulnerabilityOccupancyLevel returns a boolean if a field has been set.

### GetVulnerableEndTime

`func (o *SovereigntyStructuresGetInner) GetVulnerableEndTime() time.Time`

GetVulnerableEndTime returns the VulnerableEndTime field if non-nil, zero value otherwise.

### GetVulnerableEndTimeOk

`func (o *SovereigntyStructuresGetInner) GetVulnerableEndTimeOk() (*time.Time, bool)`

GetVulnerableEndTimeOk returns a tuple with the VulnerableEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableEndTime

`func (o *SovereigntyStructuresGetInner) SetVulnerableEndTime(v time.Time)`

SetVulnerableEndTime sets VulnerableEndTime field to given value.

### HasVulnerableEndTime

`func (o *SovereigntyStructuresGetInner) HasVulnerableEndTime() bool`

HasVulnerableEndTime returns a boolean if a field has been set.

### GetVulnerableStartTime

`func (o *SovereigntyStructuresGetInner) GetVulnerableStartTime() time.Time`

GetVulnerableStartTime returns the VulnerableStartTime field if non-nil, zero value otherwise.

### GetVulnerableStartTimeOk

`func (o *SovereigntyStructuresGetInner) GetVulnerableStartTimeOk() (*time.Time, bool)`

GetVulnerableStartTimeOk returns a tuple with the VulnerableStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableStartTime

`func (o *SovereigntyStructuresGetInner) SetVulnerableStartTime(v time.Time)`

SetVulnerableStartTime sets VulnerableStartTime field to given value.

### HasVulnerableStartTime

`func (o *SovereigntyStructuresGetInner) HasVulnerableStartTime() bool`

HasVulnerableStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


