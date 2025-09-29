# PutFleetsFleetIdMembersMemberIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | If a character is moved to the &#x60;fleet_commander&#x60; role, neither &#x60;wing_id&#x60; or &#x60;squad_id&#x60; should be specified. If a character is moved to the &#x60;wing_commander&#x60; role, only &#x60;wing_id&#x60; should be specified. If a character is moved to the &#x60;squad_commander&#x60; role, both &#x60;wing_id&#x60; and &#x60;squad_id&#x60; should be specified. If a character is moved to the &#x60;squad_member&#x60; role, both &#x60;wing_id&#x60; and &#x60;squad_id&#x60; should be specified. | 
**SquadId** | Pointer to **int64** |  | [optional] 
**WingId** | Pointer to **int64** |  | [optional] 

## Methods

### NewPutFleetsFleetIdMembersMemberIdRequest

`func NewPutFleetsFleetIdMembersMemberIdRequest(role string, ) *PutFleetsFleetIdMembersMemberIdRequest`

NewPutFleetsFleetIdMembersMemberIdRequest instantiates a new PutFleetsFleetIdMembersMemberIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFleetsFleetIdMembersMemberIdRequestWithDefaults

`func NewPutFleetsFleetIdMembersMemberIdRequestWithDefaults() *PutFleetsFleetIdMembersMemberIdRequest`

NewPutFleetsFleetIdMembersMemberIdRequestWithDefaults instantiates a new PutFleetsFleetIdMembersMemberIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *PutFleetsFleetIdMembersMemberIdRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PutFleetsFleetIdMembersMemberIdRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PutFleetsFleetIdMembersMemberIdRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetSquadId

`func (o *PutFleetsFleetIdMembersMemberIdRequest) GetSquadId() int64`

GetSquadId returns the SquadId field if non-nil, zero value otherwise.

### GetSquadIdOk

`func (o *PutFleetsFleetIdMembersMemberIdRequest) GetSquadIdOk() (*int64, bool)`

GetSquadIdOk returns a tuple with the SquadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquadId

`func (o *PutFleetsFleetIdMembersMemberIdRequest) SetSquadId(v int64)`

SetSquadId sets SquadId field to given value.

### HasSquadId

`func (o *PutFleetsFleetIdMembersMemberIdRequest) HasSquadId() bool`

HasSquadId returns a boolean if a field has been set.

### GetWingId

`func (o *PutFleetsFleetIdMembersMemberIdRequest) GetWingId() int64`

GetWingId returns the WingId field if non-nil, zero value otherwise.

### GetWingIdOk

`func (o *PutFleetsFleetIdMembersMemberIdRequest) GetWingIdOk() (*int64, bool)`

GetWingIdOk returns a tuple with the WingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWingId

`func (o *PutFleetsFleetIdMembersMemberIdRequest) SetWingId(v int64)`

SetWingId sets WingId field to given value.

### HasWingId

`func (o *PutFleetsFleetIdMembersMemberIdRequest) HasWingId() bool`

HasWingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


