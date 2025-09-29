# CharactersCharacterIdClonesGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HomeLocation** | Pointer to [**CharactersCharacterIdClonesGetHomeLocation**](CharactersCharacterIdClonesGetHomeLocation.md) |  | [optional] 
**JumpClones** | [**[]CharactersCharacterIdClonesGetJumpClonesInner**](CharactersCharacterIdClonesGetJumpClonesInner.md) |  | 
**LastCloneJumpDate** | Pointer to **time.Time** |  | [optional] 
**LastStationChangeDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCharactersCharacterIdClonesGet

`func NewCharactersCharacterIdClonesGet(jumpClones []CharactersCharacterIdClonesGetJumpClonesInner, ) *CharactersCharacterIdClonesGet`

NewCharactersCharacterIdClonesGet instantiates a new CharactersCharacterIdClonesGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdClonesGetWithDefaults

`func NewCharactersCharacterIdClonesGetWithDefaults() *CharactersCharacterIdClonesGet`

NewCharactersCharacterIdClonesGetWithDefaults instantiates a new CharactersCharacterIdClonesGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHomeLocation

`func (o *CharactersCharacterIdClonesGet) GetHomeLocation() CharactersCharacterIdClonesGetHomeLocation`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *CharactersCharacterIdClonesGet) GetHomeLocationOk() (*CharactersCharacterIdClonesGetHomeLocation, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *CharactersCharacterIdClonesGet) SetHomeLocation(v CharactersCharacterIdClonesGetHomeLocation)`

SetHomeLocation sets HomeLocation field to given value.

### HasHomeLocation

`func (o *CharactersCharacterIdClonesGet) HasHomeLocation() bool`

HasHomeLocation returns a boolean if a field has been set.

### GetJumpClones

`func (o *CharactersCharacterIdClonesGet) GetJumpClones() []CharactersCharacterIdClonesGetJumpClonesInner`

GetJumpClones returns the JumpClones field if non-nil, zero value otherwise.

### GetJumpClonesOk

`func (o *CharactersCharacterIdClonesGet) GetJumpClonesOk() (*[]CharactersCharacterIdClonesGetJumpClonesInner, bool)`

GetJumpClonesOk returns a tuple with the JumpClones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpClones

`func (o *CharactersCharacterIdClonesGet) SetJumpClones(v []CharactersCharacterIdClonesGetJumpClonesInner)`

SetJumpClones sets JumpClones field to given value.


### GetLastCloneJumpDate

`func (o *CharactersCharacterIdClonesGet) GetLastCloneJumpDate() time.Time`

GetLastCloneJumpDate returns the LastCloneJumpDate field if non-nil, zero value otherwise.

### GetLastCloneJumpDateOk

`func (o *CharactersCharacterIdClonesGet) GetLastCloneJumpDateOk() (*time.Time, bool)`

GetLastCloneJumpDateOk returns a tuple with the LastCloneJumpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloneJumpDate

`func (o *CharactersCharacterIdClonesGet) SetLastCloneJumpDate(v time.Time)`

SetLastCloneJumpDate sets LastCloneJumpDate field to given value.

### HasLastCloneJumpDate

`func (o *CharactersCharacterIdClonesGet) HasLastCloneJumpDate() bool`

HasLastCloneJumpDate returns a boolean if a field has been set.

### GetLastStationChangeDate

`func (o *CharactersCharacterIdClonesGet) GetLastStationChangeDate() time.Time`

GetLastStationChangeDate returns the LastStationChangeDate field if non-nil, zero value otherwise.

### GetLastStationChangeDateOk

`func (o *CharactersCharacterIdClonesGet) GetLastStationChangeDateOk() (*time.Time, bool)`

GetLastStationChangeDateOk returns a tuple with the LastStationChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStationChangeDate

`func (o *CharactersCharacterIdClonesGet) SetLastStationChangeDate(v time.Time)`

SetLastStationChangeDate sets LastStationChangeDate field to given value.

### HasLastStationChangeDate

`func (o *CharactersCharacterIdClonesGet) HasLastStationChangeDate() bool`

HasLastStationChangeDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


