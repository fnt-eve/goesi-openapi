# CorporationsCorporationIdRolesHistoryGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedAt** | **time.Time** |  | 
**CharacterId** | **int64** | The character whose roles are changed | 
**IssuerId** | **int64** | ID of the character who issued this change | 
**NewRoles** | **[]string** |  | 
**OldRoles** | **[]string** |  | 
**RoleType** | **string** |  | 

## Methods

### NewCorporationsCorporationIdRolesHistoryGetInner

`func NewCorporationsCorporationIdRolesHistoryGetInner(changedAt time.Time, characterId int64, issuerId int64, newRoles []string, oldRoles []string, roleType string, ) *CorporationsCorporationIdRolesHistoryGetInner`

NewCorporationsCorporationIdRolesHistoryGetInner instantiates a new CorporationsCorporationIdRolesHistoryGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdRolesHistoryGetInnerWithDefaults

`func NewCorporationsCorporationIdRolesHistoryGetInnerWithDefaults() *CorporationsCorporationIdRolesHistoryGetInner`

NewCorporationsCorporationIdRolesHistoryGetInnerWithDefaults instantiates a new CorporationsCorporationIdRolesHistoryGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangedAt

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetChangedAt() time.Time`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetChangedAtOk() (*time.Time, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *CorporationsCorporationIdRolesHistoryGetInner) SetChangedAt(v time.Time)`

SetChangedAt sets ChangedAt field to given value.


### GetCharacterId

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetCharacterId() int64`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetCharacterIdOk() (*int64, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *CorporationsCorporationIdRolesHistoryGetInner) SetCharacterId(v int64)`

SetCharacterId sets CharacterId field to given value.


### GetIssuerId

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetIssuerId() int64`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetIssuerIdOk() (*int64, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *CorporationsCorporationIdRolesHistoryGetInner) SetIssuerId(v int64)`

SetIssuerId sets IssuerId field to given value.


### GetNewRoles

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetNewRoles() []string`

GetNewRoles returns the NewRoles field if non-nil, zero value otherwise.

### GetNewRolesOk

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetNewRolesOk() (*[]string, bool)`

GetNewRolesOk returns a tuple with the NewRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRoles

`func (o *CorporationsCorporationIdRolesHistoryGetInner) SetNewRoles(v []string)`

SetNewRoles sets NewRoles field to given value.


### GetOldRoles

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetOldRoles() []string`

GetOldRoles returns the OldRoles field if non-nil, zero value otherwise.

### GetOldRolesOk

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetOldRolesOk() (*[]string, bool)`

GetOldRolesOk returns a tuple with the OldRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldRoles

`func (o *CorporationsCorporationIdRolesHistoryGetInner) SetOldRoles(v []string)`

SetOldRoles sets OldRoles field to given value.


### GetRoleType

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *CorporationsCorporationIdRolesHistoryGetInner) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *CorporationsCorporationIdRolesHistoryGetInner) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


