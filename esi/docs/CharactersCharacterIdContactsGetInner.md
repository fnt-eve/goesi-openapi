# CharactersCharacterIdContactsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | **int64** |  | 
**ContactType** | **string** |  | 
**IsBlocked** | Pointer to **bool** | Whether this contact is in the blocked list. Note a missing value denotes unknown, not true or false | [optional] 
**IsWatched** | Pointer to **bool** | Whether this contact is being watched | [optional] 
**LabelIds** | Pointer to **[]int64** |  | [optional] 
**Standing** | **float64** | Standing of the contact | 

## Methods

### NewCharactersCharacterIdContactsGetInner

`func NewCharactersCharacterIdContactsGetInner(contactId int64, contactType string, standing float64, ) *CharactersCharacterIdContactsGetInner`

NewCharactersCharacterIdContactsGetInner instantiates a new CharactersCharacterIdContactsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdContactsGetInnerWithDefaults

`func NewCharactersCharacterIdContactsGetInnerWithDefaults() *CharactersCharacterIdContactsGetInner`

NewCharactersCharacterIdContactsGetInnerWithDefaults instantiates a new CharactersCharacterIdContactsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *CharactersCharacterIdContactsGetInner) GetContactId() int64`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *CharactersCharacterIdContactsGetInner) GetContactIdOk() (*int64, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *CharactersCharacterIdContactsGetInner) SetContactId(v int64)`

SetContactId sets ContactId field to given value.


### GetContactType

`func (o *CharactersCharacterIdContactsGetInner) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *CharactersCharacterIdContactsGetInner) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *CharactersCharacterIdContactsGetInner) SetContactType(v string)`

SetContactType sets ContactType field to given value.


### GetIsBlocked

`func (o *CharactersCharacterIdContactsGetInner) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *CharactersCharacterIdContactsGetInner) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *CharactersCharacterIdContactsGetInner) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *CharactersCharacterIdContactsGetInner) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetIsWatched

`func (o *CharactersCharacterIdContactsGetInner) GetIsWatched() bool`

GetIsWatched returns the IsWatched field if non-nil, zero value otherwise.

### GetIsWatchedOk

`func (o *CharactersCharacterIdContactsGetInner) GetIsWatchedOk() (*bool, bool)`

GetIsWatchedOk returns a tuple with the IsWatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWatched

`func (o *CharactersCharacterIdContactsGetInner) SetIsWatched(v bool)`

SetIsWatched sets IsWatched field to given value.

### HasIsWatched

`func (o *CharactersCharacterIdContactsGetInner) HasIsWatched() bool`

HasIsWatched returns a boolean if a field has been set.

### GetLabelIds

`func (o *CharactersCharacterIdContactsGetInner) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CharactersCharacterIdContactsGetInner) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CharactersCharacterIdContactsGetInner) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CharactersCharacterIdContactsGetInner) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetStanding

`func (o *CharactersCharacterIdContactsGetInner) GetStanding() float64`

GetStanding returns the Standing field if non-nil, zero value otherwise.

### GetStandingOk

`func (o *CharactersCharacterIdContactsGetInner) GetStandingOk() (*float64, bool)`

GetStandingOk returns a tuple with the Standing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanding

`func (o *CharactersCharacterIdContactsGetInner) SetStanding(v float64)`

SetStanding sets Standing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


