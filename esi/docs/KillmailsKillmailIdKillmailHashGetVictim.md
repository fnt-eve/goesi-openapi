# KillmailsKillmailIdKillmailHashGetVictim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** |  | [optional] 
**CharacterId** | Pointer to **int64** |  | [optional] 
**CorporationId** | Pointer to **int64** |  | [optional] 
**DamageTaken** | **int64** | How much total damage was taken by the victim  | 
**FactionId** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]KillmailsKillmailIdKillmailHashGetVictimItemsInner**](KillmailsKillmailIdKillmailHashGetVictimItemsInner.md) |  | [optional] 
**Position** | Pointer to [**KillmailsKillmailIdKillmailHashGetVictimPosition**](KillmailsKillmailIdKillmailHashGetVictimPosition.md) |  | [optional] 
**ShipTypeId** | **int64** | The ship that the victim was piloting and was destroyed  | 

## Methods

### NewKillmailsKillmailIdKillmailHashGetVictim

`func NewKillmailsKillmailIdKillmailHashGetVictim(damageTaken int64, shipTypeId int64, ) *KillmailsKillmailIdKillmailHashGetVictim`

NewKillmailsKillmailIdKillmailHashGetVictim instantiates a new KillmailsKillmailIdKillmailHashGetVictim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKillmailsKillmailIdKillmailHashGetVictimWithDefaults

`func NewKillmailsKillmailIdKillmailHashGetVictimWithDefaults() *KillmailsKillmailIdKillmailHashGetVictim`

NewKillmailsKillmailIdKillmailHashGetVictimWithDefaults instantiates a new KillmailsKillmailIdKillmailHashGetVictim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetCharacterId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetCharacterId() int64`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetCharacterIdOk() (*int64, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetCharacterId(v int64)`

SetCharacterId sets CharacterId field to given value.

### HasCharacterId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) HasCharacterId() bool`

HasCharacterId returns a boolean if a field has been set.

### GetCorporationId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.

### HasCorporationId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) HasCorporationId() bool`

HasCorporationId returns a boolean if a field has been set.

### GetDamageTaken

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetDamageTaken() int64`

GetDamageTaken returns the DamageTaken field if non-nil, zero value otherwise.

### GetDamageTakenOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetDamageTakenOk() (*int64, bool)`

GetDamageTakenOk returns a tuple with the DamageTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageTaken

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetDamageTaken(v int64)`

SetDamageTaken sets DamageTaken field to given value.


### GetFactionId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetItems

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetItems() []KillmailsKillmailIdKillmailHashGetVictimItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetItemsOk() (*[]KillmailsKillmailIdKillmailHashGetVictimItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetItems(v []KillmailsKillmailIdKillmailHashGetVictimItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *KillmailsKillmailIdKillmailHashGetVictim) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPosition

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetPosition() KillmailsKillmailIdKillmailHashGetVictimPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetPositionOk() (*KillmailsKillmailIdKillmailHashGetVictimPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetPosition(v KillmailsKillmailIdKillmailHashGetVictimPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *KillmailsKillmailIdKillmailHashGetVictim) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetShipTypeId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetShipTypeId() int64`

GetShipTypeId returns the ShipTypeId field if non-nil, zero value otherwise.

### GetShipTypeIdOk

`func (o *KillmailsKillmailIdKillmailHashGetVictim) GetShipTypeIdOk() (*int64, bool)`

GetShipTypeIdOk returns a tuple with the ShipTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTypeId

`func (o *KillmailsKillmailIdKillmailHashGetVictim) SetShipTypeId(v int64)`

SetShipTypeId sets ShipTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


