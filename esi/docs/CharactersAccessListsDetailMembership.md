# CharactersAccessListsDetailMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alliances** | [**[]CharactersAccessListsDetailAllianceentry**](CharactersAccessListsDetailAllianceentry.md) | Alliances in the Access List | 
**AllowEveryone** | **bool** | Whether everyone is allowed unless blocked | 
**Characters** | [**[]CharactersAccessListsDetailCharacterentry**](CharactersAccessListsDetailCharacterentry.md) | Characters in the Access List | 
**Corporations** | [**[]CharactersAccessListsDetailCorporationentry**](CharactersAccessListsDetailCorporationentry.md) | Corporations in the Access List | 

## Methods

### NewCharactersAccessListsDetailMembership

`func NewCharactersAccessListsDetailMembership(alliances []CharactersAccessListsDetailAllianceentry, allowEveryone bool, characters []CharactersAccessListsDetailCharacterentry, corporations []CharactersAccessListsDetailCorporationentry, ) *CharactersAccessListsDetailMembership`

NewCharactersAccessListsDetailMembership instantiates a new CharactersAccessListsDetailMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersAccessListsDetailMembershipWithDefaults

`func NewCharactersAccessListsDetailMembershipWithDefaults() *CharactersAccessListsDetailMembership`

NewCharactersAccessListsDetailMembershipWithDefaults instantiates a new CharactersAccessListsDetailMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlliances

`func (o *CharactersAccessListsDetailMembership) GetAlliances() []CharactersAccessListsDetailAllianceentry`

GetAlliances returns the Alliances field if non-nil, zero value otherwise.

### GetAlliancesOk

`func (o *CharactersAccessListsDetailMembership) GetAlliancesOk() (*[]CharactersAccessListsDetailAllianceentry, bool)`

GetAlliancesOk returns a tuple with the Alliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlliances

`func (o *CharactersAccessListsDetailMembership) SetAlliances(v []CharactersAccessListsDetailAllianceentry)`

SetAlliances sets Alliances field to given value.


### GetAllowEveryone

`func (o *CharactersAccessListsDetailMembership) GetAllowEveryone() bool`

GetAllowEveryone returns the AllowEveryone field if non-nil, zero value otherwise.

### GetAllowEveryoneOk

`func (o *CharactersAccessListsDetailMembership) GetAllowEveryoneOk() (*bool, bool)`

GetAllowEveryoneOk returns a tuple with the AllowEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEveryone

`func (o *CharactersAccessListsDetailMembership) SetAllowEveryone(v bool)`

SetAllowEveryone sets AllowEveryone field to given value.


### GetCharacters

`func (o *CharactersAccessListsDetailMembership) GetCharacters() []CharactersAccessListsDetailCharacterentry`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *CharactersAccessListsDetailMembership) GetCharactersOk() (*[]CharactersAccessListsDetailCharacterentry, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *CharactersAccessListsDetailMembership) SetCharacters(v []CharactersAccessListsDetailCharacterentry)`

SetCharacters sets Characters field to given value.


### GetCorporations

`func (o *CharactersAccessListsDetailMembership) GetCorporations() []CharactersAccessListsDetailCorporationentry`

GetCorporations returns the Corporations field if non-nil, zero value otherwise.

### GetCorporationsOk

`func (o *CharactersAccessListsDetailMembership) GetCorporationsOk() (*[]CharactersAccessListsDetailCorporationentry, bool)`

GetCorporationsOk returns a tuple with the Corporations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporations

`func (o *CharactersAccessListsDetailMembership) SetCorporations(v []CharactersAccessListsDetailCorporationentry)`

SetCorporations sets Corporations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


