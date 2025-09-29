# CharactersCharacterIdFleetGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetBossId** | **int64** | Character ID of the current fleet boss | 
**FleetId** | **int64** | The character&#39;s current fleet ID | 
**Role** | **string** | Member’s role in fleet | 
**SquadId** | **int64** | ID of the squad the member is in. If not applicable, will be set to -1 | 
**WingId** | **int64** | ID of the wing the member is in. If not applicable, will be set to -1 | 

## Methods

### NewCharactersCharacterIdFleetGet

`func NewCharactersCharacterIdFleetGet(fleetBossId int64, fleetId int64, role string, squadId int64, wingId int64, ) *CharactersCharacterIdFleetGet`

NewCharactersCharacterIdFleetGet instantiates a new CharactersCharacterIdFleetGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdFleetGetWithDefaults

`func NewCharactersCharacterIdFleetGetWithDefaults() *CharactersCharacterIdFleetGet`

NewCharactersCharacterIdFleetGetWithDefaults instantiates a new CharactersCharacterIdFleetGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetBossId

`func (o *CharactersCharacterIdFleetGet) GetFleetBossId() int64`

GetFleetBossId returns the FleetBossId field if non-nil, zero value otherwise.

### GetFleetBossIdOk

`func (o *CharactersCharacterIdFleetGet) GetFleetBossIdOk() (*int64, bool)`

GetFleetBossIdOk returns a tuple with the FleetBossId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetBossId

`func (o *CharactersCharacterIdFleetGet) SetFleetBossId(v int64)`

SetFleetBossId sets FleetBossId field to given value.


### GetFleetId

`func (o *CharactersCharacterIdFleetGet) GetFleetId() int64`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *CharactersCharacterIdFleetGet) GetFleetIdOk() (*int64, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *CharactersCharacterIdFleetGet) SetFleetId(v int64)`

SetFleetId sets FleetId field to given value.


### GetRole

`func (o *CharactersCharacterIdFleetGet) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CharactersCharacterIdFleetGet) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CharactersCharacterIdFleetGet) SetRole(v string)`

SetRole sets Role field to given value.


### GetSquadId

`func (o *CharactersCharacterIdFleetGet) GetSquadId() int64`

GetSquadId returns the SquadId field if non-nil, zero value otherwise.

### GetSquadIdOk

`func (o *CharactersCharacterIdFleetGet) GetSquadIdOk() (*int64, bool)`

GetSquadIdOk returns a tuple with the SquadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquadId

`func (o *CharactersCharacterIdFleetGet) SetSquadId(v int64)`

SetSquadId sets SquadId field to given value.


### GetWingId

`func (o *CharactersCharacterIdFleetGet) GetWingId() int64`

GetWingId returns the WingId field if non-nil, zero value otherwise.

### GetWingIdOk

`func (o *CharactersCharacterIdFleetGet) GetWingIdOk() (*int64, bool)`

GetWingIdOk returns a tuple with the WingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWingId

`func (o *CharactersCharacterIdFleetGet) SetWingId(v int64)`

SetWingId sets WingId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


