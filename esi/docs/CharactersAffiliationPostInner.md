# CharactersAffiliationPostInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** | The character&#39;s alliance ID, if their corporation is in an alliance | [optional] 
**CharacterId** | **int64** | The character&#39;s ID | 
**CorporationId** | **int64** | The character&#39;s corporation ID | 
**FactionId** | Pointer to **int64** | The character&#39;s faction ID, if their corporation is in a faction | [optional] 

## Methods

### NewCharactersAffiliationPostInner

`func NewCharactersAffiliationPostInner(characterId int64, corporationId int64, ) *CharactersAffiliationPostInner`

NewCharactersAffiliationPostInner instantiates a new CharactersAffiliationPostInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersAffiliationPostInnerWithDefaults

`func NewCharactersAffiliationPostInnerWithDefaults() *CharactersAffiliationPostInner`

NewCharactersAffiliationPostInnerWithDefaults instantiates a new CharactersAffiliationPostInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *CharactersAffiliationPostInner) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *CharactersAffiliationPostInner) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *CharactersAffiliationPostInner) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *CharactersAffiliationPostInner) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetCharacterId

`func (o *CharactersAffiliationPostInner) GetCharacterId() int64`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *CharactersAffiliationPostInner) GetCharacterIdOk() (*int64, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *CharactersAffiliationPostInner) SetCharacterId(v int64)`

SetCharacterId sets CharacterId field to given value.


### GetCorporationId

`func (o *CharactersAffiliationPostInner) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *CharactersAffiliationPostInner) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *CharactersAffiliationPostInner) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.


### GetFactionId

`func (o *CharactersAffiliationPostInner) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *CharactersAffiliationPostInner) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *CharactersAffiliationPostInner) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *CharactersAffiliationPostInner) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


