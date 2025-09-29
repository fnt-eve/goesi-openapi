# FleetsFleetIdMembersGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterId** | **int64** |  | 
**JoinTime** | **time.Time** |  | 
**Role** | **string** | Member’s role in fleet | 
**RoleName** | **string** | Localized role names | 
**ShipTypeId** | **int64** |  | 
**SolarSystemId** | **int64** | Solar system the member is located in | 
**SquadId** | **int64** | ID of the squad the member is in. If not applicable, will be set to -1 | 
**StationId** | Pointer to **int64** | Station in which the member is docked in, if applicable | [optional] 
**TakesFleetWarp** | **bool** | Whether the member take fleet warps | 
**WingId** | **int64** | ID of the wing the member is in. If not applicable, will be set to -1 | 

## Methods

### NewFleetsFleetIdMembersGetInner

`func NewFleetsFleetIdMembersGetInner(characterId int64, joinTime time.Time, role string, roleName string, shipTypeId int64, solarSystemId int64, squadId int64, takesFleetWarp bool, wingId int64, ) *FleetsFleetIdMembersGetInner`

NewFleetsFleetIdMembersGetInner instantiates a new FleetsFleetIdMembersGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetsFleetIdMembersGetInnerWithDefaults

`func NewFleetsFleetIdMembersGetInnerWithDefaults() *FleetsFleetIdMembersGetInner`

NewFleetsFleetIdMembersGetInnerWithDefaults instantiates a new FleetsFleetIdMembersGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacterId

`func (o *FleetsFleetIdMembersGetInner) GetCharacterId() int64`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *FleetsFleetIdMembersGetInner) GetCharacterIdOk() (*int64, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *FleetsFleetIdMembersGetInner) SetCharacterId(v int64)`

SetCharacterId sets CharacterId field to given value.


### GetJoinTime

`func (o *FleetsFleetIdMembersGetInner) GetJoinTime() time.Time`

GetJoinTime returns the JoinTime field if non-nil, zero value otherwise.

### GetJoinTimeOk

`func (o *FleetsFleetIdMembersGetInner) GetJoinTimeOk() (*time.Time, bool)`

GetJoinTimeOk returns a tuple with the JoinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTime

`func (o *FleetsFleetIdMembersGetInner) SetJoinTime(v time.Time)`

SetJoinTime sets JoinTime field to given value.


### GetRole

`func (o *FleetsFleetIdMembersGetInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FleetsFleetIdMembersGetInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FleetsFleetIdMembersGetInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetRoleName

`func (o *FleetsFleetIdMembersGetInner) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *FleetsFleetIdMembersGetInner) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *FleetsFleetIdMembersGetInner) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetShipTypeId

`func (o *FleetsFleetIdMembersGetInner) GetShipTypeId() int64`

GetShipTypeId returns the ShipTypeId field if non-nil, zero value otherwise.

### GetShipTypeIdOk

`func (o *FleetsFleetIdMembersGetInner) GetShipTypeIdOk() (*int64, bool)`

GetShipTypeIdOk returns a tuple with the ShipTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTypeId

`func (o *FleetsFleetIdMembersGetInner) SetShipTypeId(v int64)`

SetShipTypeId sets ShipTypeId field to given value.


### GetSolarSystemId

`func (o *FleetsFleetIdMembersGetInner) GetSolarSystemId() int64`

GetSolarSystemId returns the SolarSystemId field if non-nil, zero value otherwise.

### GetSolarSystemIdOk

`func (o *FleetsFleetIdMembersGetInner) GetSolarSystemIdOk() (*int64, bool)`

GetSolarSystemIdOk returns a tuple with the SolarSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarSystemId

`func (o *FleetsFleetIdMembersGetInner) SetSolarSystemId(v int64)`

SetSolarSystemId sets SolarSystemId field to given value.


### GetSquadId

`func (o *FleetsFleetIdMembersGetInner) GetSquadId() int64`

GetSquadId returns the SquadId field if non-nil, zero value otherwise.

### GetSquadIdOk

`func (o *FleetsFleetIdMembersGetInner) GetSquadIdOk() (*int64, bool)`

GetSquadIdOk returns a tuple with the SquadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquadId

`func (o *FleetsFleetIdMembersGetInner) SetSquadId(v int64)`

SetSquadId sets SquadId field to given value.


### GetStationId

`func (o *FleetsFleetIdMembersGetInner) GetStationId() int64`

GetStationId returns the StationId field if non-nil, zero value otherwise.

### GetStationIdOk

`func (o *FleetsFleetIdMembersGetInner) GetStationIdOk() (*int64, bool)`

GetStationIdOk returns a tuple with the StationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationId

`func (o *FleetsFleetIdMembersGetInner) SetStationId(v int64)`

SetStationId sets StationId field to given value.

### HasStationId

`func (o *FleetsFleetIdMembersGetInner) HasStationId() bool`

HasStationId returns a boolean if a field has been set.

### GetTakesFleetWarp

`func (o *FleetsFleetIdMembersGetInner) GetTakesFleetWarp() bool`

GetTakesFleetWarp returns the TakesFleetWarp field if non-nil, zero value otherwise.

### GetTakesFleetWarpOk

`func (o *FleetsFleetIdMembersGetInner) GetTakesFleetWarpOk() (*bool, bool)`

GetTakesFleetWarpOk returns a tuple with the TakesFleetWarp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakesFleetWarp

`func (o *FleetsFleetIdMembersGetInner) SetTakesFleetWarp(v bool)`

SetTakesFleetWarp sets TakesFleetWarp field to given value.


### GetWingId

`func (o *FleetsFleetIdMembersGetInner) GetWingId() int64`

GetWingId returns the WingId field if non-nil, zero value otherwise.

### GetWingIdOk

`func (o *FleetsFleetIdMembersGetInner) GetWingIdOk() (*int64, bool)`

GetWingIdOk returns a tuple with the WingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWingId

`func (o *FleetsFleetIdMembersGetInner) SetWingId(v int64)`

SetWingId sets WingId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


