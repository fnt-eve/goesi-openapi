# CorporationsCorporationIdStarbasesGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoonId** | Pointer to **int64** | The moon this starbase (POS) is anchored on, unanchored POSes do not have this information | [optional] 
**OnlinedSince** | Pointer to **time.Time** | When the POS onlined, for starbases (POSes) in online state | [optional] 
**ReinforcedUntil** | Pointer to **time.Time** | When the POS will be out of reinforcement, for starbases (POSes) in reinforced state | [optional] 
**StarbaseId** | **int64** | Unique ID for this starbase (POS) | 
**State** | Pointer to **string** |  | [optional] 
**SystemId** | **int64** | The solar system this starbase (POS) is in, unanchored POSes have this information | 
**TypeId** | **int64** | Starbase (POS) type | 
**UnanchorAt** | Pointer to **time.Time** | When the POS started unanchoring, for starbases (POSes) in unanchoring state | [optional] 

## Methods

### NewCorporationsCorporationIdStarbasesGetInner

`func NewCorporationsCorporationIdStarbasesGetInner(starbaseId int64, systemId int64, typeId int64, ) *CorporationsCorporationIdStarbasesGetInner`

NewCorporationsCorporationIdStarbasesGetInner instantiates a new CorporationsCorporationIdStarbasesGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdStarbasesGetInnerWithDefaults

`func NewCorporationsCorporationIdStarbasesGetInnerWithDefaults() *CorporationsCorporationIdStarbasesGetInner`

NewCorporationsCorporationIdStarbasesGetInnerWithDefaults instantiates a new CorporationsCorporationIdStarbasesGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoonId

`func (o *CorporationsCorporationIdStarbasesGetInner) GetMoonId() int64`

GetMoonId returns the MoonId field if non-nil, zero value otherwise.

### GetMoonIdOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetMoonIdOk() (*int64, bool)`

GetMoonIdOk returns a tuple with the MoonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonId

`func (o *CorporationsCorporationIdStarbasesGetInner) SetMoonId(v int64)`

SetMoonId sets MoonId field to given value.

### HasMoonId

`func (o *CorporationsCorporationIdStarbasesGetInner) HasMoonId() bool`

HasMoonId returns a boolean if a field has been set.

### GetOnlinedSince

`func (o *CorporationsCorporationIdStarbasesGetInner) GetOnlinedSince() time.Time`

GetOnlinedSince returns the OnlinedSince field if non-nil, zero value otherwise.

### GetOnlinedSinceOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetOnlinedSinceOk() (*time.Time, bool)`

GetOnlinedSinceOk returns a tuple with the OnlinedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlinedSince

`func (o *CorporationsCorporationIdStarbasesGetInner) SetOnlinedSince(v time.Time)`

SetOnlinedSince sets OnlinedSince field to given value.

### HasOnlinedSince

`func (o *CorporationsCorporationIdStarbasesGetInner) HasOnlinedSince() bool`

HasOnlinedSince returns a boolean if a field has been set.

### GetReinforcedUntil

`func (o *CorporationsCorporationIdStarbasesGetInner) GetReinforcedUntil() time.Time`

GetReinforcedUntil returns the ReinforcedUntil field if non-nil, zero value otherwise.

### GetReinforcedUntilOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetReinforcedUntilOk() (*time.Time, bool)`

GetReinforcedUntilOk returns a tuple with the ReinforcedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinforcedUntil

`func (o *CorporationsCorporationIdStarbasesGetInner) SetReinforcedUntil(v time.Time)`

SetReinforcedUntil sets ReinforcedUntil field to given value.

### HasReinforcedUntil

`func (o *CorporationsCorporationIdStarbasesGetInner) HasReinforcedUntil() bool`

HasReinforcedUntil returns a boolean if a field has been set.

### GetStarbaseId

`func (o *CorporationsCorporationIdStarbasesGetInner) GetStarbaseId() int64`

GetStarbaseId returns the StarbaseId field if non-nil, zero value otherwise.

### GetStarbaseIdOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetStarbaseIdOk() (*int64, bool)`

GetStarbaseIdOk returns a tuple with the StarbaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarbaseId

`func (o *CorporationsCorporationIdStarbasesGetInner) SetStarbaseId(v int64)`

SetStarbaseId sets StarbaseId field to given value.


### GetState

`func (o *CorporationsCorporationIdStarbasesGetInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CorporationsCorporationIdStarbasesGetInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CorporationsCorporationIdStarbasesGetInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSystemId

`func (o *CorporationsCorporationIdStarbasesGetInner) GetSystemId() int64`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetSystemIdOk() (*int64, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *CorporationsCorporationIdStarbasesGetInner) SetSystemId(v int64)`

SetSystemId sets SystemId field to given value.


### GetTypeId

`func (o *CorporationsCorporationIdStarbasesGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsCorporationIdStarbasesGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetUnanchorAt

`func (o *CorporationsCorporationIdStarbasesGetInner) GetUnanchorAt() time.Time`

GetUnanchorAt returns the UnanchorAt field if non-nil, zero value otherwise.

### GetUnanchorAtOk

`func (o *CorporationsCorporationIdStarbasesGetInner) GetUnanchorAtOk() (*time.Time, bool)`

GetUnanchorAtOk returns a tuple with the UnanchorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnanchorAt

`func (o *CorporationsCorporationIdStarbasesGetInner) SetUnanchorAt(v time.Time)`

SetUnanchorAt sets UnanchorAt field to given value.

### HasUnanchorAt

`func (o *CorporationsCorporationIdStarbasesGetInner) HasUnanchorAt() bool`

HasUnanchorAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


