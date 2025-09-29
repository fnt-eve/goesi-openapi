# KillmailsKillmailIdKillmailHashGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attackers** | [**[]KillmailsKillmailIdKillmailHashGetAttackersInner**](KillmailsKillmailIdKillmailHashGetAttackersInner.md) |  | 
**KillmailId** | **int64** | ID of the killmail | 
**KillmailTime** | **time.Time** | Time that the victim was killed and the killmail generated  | 
**MoonId** | Pointer to **int64** | Moon if the kill took place at one | [optional] 
**SolarSystemId** | **int64** | Solar system that the kill took place in  | 
**Victim** | [**KillmailsKillmailIdKillmailHashGetVictim**](KillmailsKillmailIdKillmailHashGetVictim.md) |  | 
**WarId** | Pointer to **int64** | War if the killmail is generated in relation to an official war  | [optional] 

## Methods

### NewKillmailsKillmailIdKillmailHashGet

`func NewKillmailsKillmailIdKillmailHashGet(attackers []KillmailsKillmailIdKillmailHashGetAttackersInner, killmailId int64, killmailTime time.Time, solarSystemId int64, victim KillmailsKillmailIdKillmailHashGetVictim, ) *KillmailsKillmailIdKillmailHashGet`

NewKillmailsKillmailIdKillmailHashGet instantiates a new KillmailsKillmailIdKillmailHashGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKillmailsKillmailIdKillmailHashGetWithDefaults

`func NewKillmailsKillmailIdKillmailHashGetWithDefaults() *KillmailsKillmailIdKillmailHashGet`

NewKillmailsKillmailIdKillmailHashGetWithDefaults instantiates a new KillmailsKillmailIdKillmailHashGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackers

`func (o *KillmailsKillmailIdKillmailHashGet) GetAttackers() []KillmailsKillmailIdKillmailHashGetAttackersInner`

GetAttackers returns the Attackers field if non-nil, zero value otherwise.

### GetAttackersOk

`func (o *KillmailsKillmailIdKillmailHashGet) GetAttackersOk() (*[]KillmailsKillmailIdKillmailHashGetAttackersInner, bool)`

GetAttackersOk returns a tuple with the Attackers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackers

`func (o *KillmailsKillmailIdKillmailHashGet) SetAttackers(v []KillmailsKillmailIdKillmailHashGetAttackersInner)`

SetAttackers sets Attackers field to given value.


### GetKillmailId

`func (o *KillmailsKillmailIdKillmailHashGet) GetKillmailId() int64`

GetKillmailId returns the KillmailId field if non-nil, zero value otherwise.

### GetKillmailIdOk

`func (o *KillmailsKillmailIdKillmailHashGet) GetKillmailIdOk() (*int64, bool)`

GetKillmailIdOk returns a tuple with the KillmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillmailId

`func (o *KillmailsKillmailIdKillmailHashGet) SetKillmailId(v int64)`

SetKillmailId sets KillmailId field to given value.


### GetKillmailTime

`func (o *KillmailsKillmailIdKillmailHashGet) GetKillmailTime() time.Time`

GetKillmailTime returns the KillmailTime field if non-nil, zero value otherwise.

### GetKillmailTimeOk

`func (o *KillmailsKillmailIdKillmailHashGet) GetKillmailTimeOk() (*time.Time, bool)`

GetKillmailTimeOk returns a tuple with the KillmailTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillmailTime

`func (o *KillmailsKillmailIdKillmailHashGet) SetKillmailTime(v time.Time)`

SetKillmailTime sets KillmailTime field to given value.


### GetMoonId

`func (o *KillmailsKillmailIdKillmailHashGet) GetMoonId() int64`

GetMoonId returns the MoonId field if non-nil, zero value otherwise.

### GetMoonIdOk

`func (o *KillmailsKillmailIdKillmailHashGet) GetMoonIdOk() (*int64, bool)`

GetMoonIdOk returns a tuple with the MoonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonId

`func (o *KillmailsKillmailIdKillmailHashGet) SetMoonId(v int64)`

SetMoonId sets MoonId field to given value.

### HasMoonId

`func (o *KillmailsKillmailIdKillmailHashGet) HasMoonId() bool`

HasMoonId returns a boolean if a field has been set.

### GetSolarSystemId

`func (o *KillmailsKillmailIdKillmailHashGet) GetSolarSystemId() int64`

GetSolarSystemId returns the SolarSystemId field if non-nil, zero value otherwise.

### GetSolarSystemIdOk

`func (o *KillmailsKillmailIdKillmailHashGet) GetSolarSystemIdOk() (*int64, bool)`

GetSolarSystemIdOk returns a tuple with the SolarSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarSystemId

`func (o *KillmailsKillmailIdKillmailHashGet) SetSolarSystemId(v int64)`

SetSolarSystemId sets SolarSystemId field to given value.


### GetVictim

`func (o *KillmailsKillmailIdKillmailHashGet) GetVictim() KillmailsKillmailIdKillmailHashGetVictim`

GetVictim returns the Victim field if non-nil, zero value otherwise.

### GetVictimOk

`func (o *KillmailsKillmailIdKillmailHashGet) GetVictimOk() (*KillmailsKillmailIdKillmailHashGetVictim, bool)`

GetVictimOk returns a tuple with the Victim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictim

`func (o *KillmailsKillmailIdKillmailHashGet) SetVictim(v KillmailsKillmailIdKillmailHashGetVictim)`

SetVictim sets Victim field to given value.


### GetWarId

`func (o *KillmailsKillmailIdKillmailHashGet) GetWarId() int64`

GetWarId returns the WarId field if non-nil, zero value otherwise.

### GetWarIdOk

`func (o *KillmailsKillmailIdKillmailHashGet) GetWarIdOk() (*int64, bool)`

GetWarIdOk returns a tuple with the WarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarId

`func (o *KillmailsKillmailIdKillmailHashGet) SetWarId(v int64)`

SetWarId sets WarId field to given value.

### HasWarId

`func (o *KillmailsKillmailIdKillmailHashGet) HasWarId() bool`

HasWarId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


