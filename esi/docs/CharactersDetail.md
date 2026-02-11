# CharactersDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** |  | [optional] 
**Birthday** | **time.Time** | Character&#39;s creation date | 
**BloodlineId** | **int64** |  | 
**CorporationId** | **int64** |  | 
**Description** | Pointer to **string** | Character&#39;s description (biography) | [optional] 
**FactionId** | Pointer to **int64** |  | [optional] 
**Gender** | **string** | Character&#39;s gender | 
**Name** | **string** | Character&#39;s name | 
**RaceId** | **int64** |  | 
**SecurityStatus** | Pointer to **float64** | Character&#39;s security status | [optional] 
**Title** | Pointer to **string** | Character&#39;s corporation title | [optional] 

## Methods

### NewCharactersDetail

`func NewCharactersDetail(birthday time.Time, bloodlineId int64, corporationId int64, gender string, name string, raceId int64, ) *CharactersDetail`

NewCharactersDetail instantiates a new CharactersDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersDetailWithDefaults

`func NewCharactersDetailWithDefaults() *CharactersDetail`

NewCharactersDetailWithDefaults instantiates a new CharactersDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *CharactersDetail) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *CharactersDetail) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *CharactersDetail) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *CharactersDetail) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetBirthday

`func (o *CharactersDetail) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *CharactersDetail) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *CharactersDetail) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.


### GetBloodlineId

`func (o *CharactersDetail) GetBloodlineId() int64`

GetBloodlineId returns the BloodlineId field if non-nil, zero value otherwise.

### GetBloodlineIdOk

`func (o *CharactersDetail) GetBloodlineIdOk() (*int64, bool)`

GetBloodlineIdOk returns a tuple with the BloodlineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloodlineId

`func (o *CharactersDetail) SetBloodlineId(v int64)`

SetBloodlineId sets BloodlineId field to given value.


### GetCorporationId

`func (o *CharactersDetail) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *CharactersDetail) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *CharactersDetail) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.


### GetDescription

`func (o *CharactersDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CharactersDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CharactersDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CharactersDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFactionId

`func (o *CharactersDetail) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *CharactersDetail) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *CharactersDetail) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *CharactersDetail) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetGender

`func (o *CharactersDetail) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CharactersDetail) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CharactersDetail) SetGender(v string)`

SetGender sets Gender field to given value.


### GetName

`func (o *CharactersDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharactersDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharactersDetail) SetName(v string)`

SetName sets Name field to given value.


### GetRaceId

`func (o *CharactersDetail) GetRaceId() int64`

GetRaceId returns the RaceId field if non-nil, zero value otherwise.

### GetRaceIdOk

`func (o *CharactersDetail) GetRaceIdOk() (*int64, bool)`

GetRaceIdOk returns a tuple with the RaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaceId

`func (o *CharactersDetail) SetRaceId(v int64)`

SetRaceId sets RaceId field to given value.


### GetSecurityStatus

`func (o *CharactersDetail) GetSecurityStatus() float64`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *CharactersDetail) GetSecurityStatusOk() (*float64, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *CharactersDetail) SetSecurityStatus(v float64)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *CharactersDetail) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CharactersDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CharactersDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CharactersDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CharactersDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


