# KillmailsKillmailIdKillmailHashGetAttackersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** |  | [optional] 
**CharacterId** | Pointer to **int64** |  | [optional] 
**CorporationId** | Pointer to **int64** |  | [optional] 
**DamageDone** | **int64** |  | 
**FactionId** | Pointer to **int64** |  | [optional] 
**FinalBlow** | **bool** | Was the attacker the one to achieve the final blow  | 
**SecurityStatus** | **float64** | Security status for the attacker  | 
**ShipTypeId** | Pointer to **int64** | What ship was the attacker flying  | [optional] 
**WeaponTypeId** | Pointer to **int64** | What weapon was used by the attacker for the kill  | [optional] 

## Methods

### NewKillmailsKillmailIdKillmailHashGetAttackersInner

`func NewKillmailsKillmailIdKillmailHashGetAttackersInner(damageDone int64, finalBlow bool, securityStatus float64, ) *KillmailsKillmailIdKillmailHashGetAttackersInner`

NewKillmailsKillmailIdKillmailHashGetAttackersInner instantiates a new KillmailsKillmailIdKillmailHashGetAttackersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKillmailsKillmailIdKillmailHashGetAttackersInnerWithDefaults

`func NewKillmailsKillmailIdKillmailHashGetAttackersInnerWithDefaults() *KillmailsKillmailIdKillmailHashGetAttackersInner`

NewKillmailsKillmailIdKillmailHashGetAttackersInnerWithDefaults instantiates a new KillmailsKillmailIdKillmailHashGetAttackersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetCharacterId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetCharacterId() int64`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetCharacterIdOk() (*int64, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetCharacterId(v int64)`

SetCharacterId sets CharacterId field to given value.

### HasCharacterId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) HasCharacterId() bool`

HasCharacterId returns a boolean if a field has been set.

### GetCorporationId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.

### HasCorporationId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) HasCorporationId() bool`

HasCorporationId returns a boolean if a field has been set.

### GetDamageDone

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetDamageDone() int64`

GetDamageDone returns the DamageDone field if non-nil, zero value otherwise.

### GetDamageDoneOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetDamageDoneOk() (*int64, bool)`

GetDamageDoneOk returns a tuple with the DamageDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDone

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetDamageDone(v int64)`

SetDamageDone sets DamageDone field to given value.


### GetFactionId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetFinalBlow

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetFinalBlow() bool`

GetFinalBlow returns the FinalBlow field if non-nil, zero value otherwise.

### GetFinalBlowOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetFinalBlowOk() (*bool, bool)`

GetFinalBlowOk returns a tuple with the FinalBlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalBlow

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetFinalBlow(v bool)`

SetFinalBlow sets FinalBlow field to given value.


### GetSecurityStatus

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetSecurityStatus() float64`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetSecurityStatusOk() (*float64, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetSecurityStatus(v float64)`

SetSecurityStatus sets SecurityStatus field to given value.


### GetShipTypeId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetShipTypeId() int64`

GetShipTypeId returns the ShipTypeId field if non-nil, zero value otherwise.

### GetShipTypeIdOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetShipTypeIdOk() (*int64, bool)`

GetShipTypeIdOk returns a tuple with the ShipTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTypeId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetShipTypeId(v int64)`

SetShipTypeId sets ShipTypeId field to given value.

### HasShipTypeId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) HasShipTypeId() bool`

HasShipTypeId returns a boolean if a field has been set.

### GetWeaponTypeId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetWeaponTypeId() int64`

GetWeaponTypeId returns the WeaponTypeId field if non-nil, zero value otherwise.

### GetWeaponTypeIdOk

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) GetWeaponTypeIdOk() (*int64, bool)`

GetWeaponTypeIdOk returns a tuple with the WeaponTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponTypeId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) SetWeaponTypeId(v int64)`

SetWeaponTypeId sets WeaponTypeId field to given value.

### HasWeaponTypeId

`func (o *KillmailsKillmailIdKillmailHashGetAttackersInner) HasWeaponTypeId() bool`

HasWeaponTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


