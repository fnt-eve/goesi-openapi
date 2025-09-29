# CharactersCharacterIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** | The character&#39;s alliance ID | [optional] 
**Birthday** | **time.Time** | Creation date of the character | 
**BloodlineId** | **int64** |  | 
**CorporationId** | **int64** | The character&#39;s corporation ID | 
**Description** | Pointer to **string** |  | [optional] 
**FactionId** | Pointer to **int64** | ID of the faction the character is fighting for, if the character is enlisted in Factional Warfare | [optional] 
**Gender** | **string** |  | 
**Name** | **string** |  | 
**RaceId** | **int64** |  | 
**SecurityStatus** | Pointer to **float64** |  | [optional] 
**Title** | Pointer to **string** | The individual title of the character | [optional] 

## Methods

### NewCharactersCharacterIdGet

`func NewCharactersCharacterIdGet(birthday time.Time, bloodlineId int64, corporationId int64, gender string, name string, raceId int64, ) *CharactersCharacterIdGet`

NewCharactersCharacterIdGet instantiates a new CharactersCharacterIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdGetWithDefaults

`func NewCharactersCharacterIdGetWithDefaults() *CharactersCharacterIdGet`

NewCharactersCharacterIdGetWithDefaults instantiates a new CharactersCharacterIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *CharactersCharacterIdGet) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *CharactersCharacterIdGet) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *CharactersCharacterIdGet) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *CharactersCharacterIdGet) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetBirthday

`func (o *CharactersCharacterIdGet) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *CharactersCharacterIdGet) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *CharactersCharacterIdGet) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.


### GetBloodlineId

`func (o *CharactersCharacterIdGet) GetBloodlineId() int64`

GetBloodlineId returns the BloodlineId field if non-nil, zero value otherwise.

### GetBloodlineIdOk

`func (o *CharactersCharacterIdGet) GetBloodlineIdOk() (*int64, bool)`

GetBloodlineIdOk returns a tuple with the BloodlineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloodlineId

`func (o *CharactersCharacterIdGet) SetBloodlineId(v int64)`

SetBloodlineId sets BloodlineId field to given value.


### GetCorporationId

`func (o *CharactersCharacterIdGet) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *CharactersCharacterIdGet) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *CharactersCharacterIdGet) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.


### GetDescription

`func (o *CharactersCharacterIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CharactersCharacterIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CharactersCharacterIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CharactersCharacterIdGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFactionId

`func (o *CharactersCharacterIdGet) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *CharactersCharacterIdGet) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *CharactersCharacterIdGet) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *CharactersCharacterIdGet) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetGender

`func (o *CharactersCharacterIdGet) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CharactersCharacterIdGet) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CharactersCharacterIdGet) SetGender(v string)`

SetGender sets Gender field to given value.


### GetName

`func (o *CharactersCharacterIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharactersCharacterIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharactersCharacterIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetRaceId

`func (o *CharactersCharacterIdGet) GetRaceId() int64`

GetRaceId returns the RaceId field if non-nil, zero value otherwise.

### GetRaceIdOk

`func (o *CharactersCharacterIdGet) GetRaceIdOk() (*int64, bool)`

GetRaceIdOk returns a tuple with the RaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaceId

`func (o *CharactersCharacterIdGet) SetRaceId(v int64)`

SetRaceId sets RaceId field to given value.


### GetSecurityStatus

`func (o *CharactersCharacterIdGet) GetSecurityStatus() float64`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *CharactersCharacterIdGet) GetSecurityStatusOk() (*float64, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *CharactersCharacterIdGet) SetSecurityStatus(v float64)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *CharactersCharacterIdGet) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CharactersCharacterIdGet) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CharactersCharacterIdGet) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CharactersCharacterIdGet) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CharactersCharacterIdGet) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


