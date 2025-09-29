# CharactersCharacterIdMedalsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorporationId** | **int64** |  | 
**Date** | **time.Time** |  | 
**Description** | **string** |  | 
**Graphics** | [**[]CharactersCharacterIdMedalsGetInnerGraphicsInner**](CharactersCharacterIdMedalsGetInnerGraphicsInner.md) |  | 
**IssuerId** | **int64** |  | 
**MedalId** | **int64** |  | 
**Reason** | **string** |  | 
**Status** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewCharactersCharacterIdMedalsGetInner

`func NewCharactersCharacterIdMedalsGetInner(corporationId int64, date time.Time, description string, graphics []CharactersCharacterIdMedalsGetInnerGraphicsInner, issuerId int64, medalId int64, reason string, status string, title string, ) *CharactersCharacterIdMedalsGetInner`

NewCharactersCharacterIdMedalsGetInner instantiates a new CharactersCharacterIdMedalsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdMedalsGetInnerWithDefaults

`func NewCharactersCharacterIdMedalsGetInnerWithDefaults() *CharactersCharacterIdMedalsGetInner`

NewCharactersCharacterIdMedalsGetInnerWithDefaults instantiates a new CharactersCharacterIdMedalsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorporationId

`func (o *CharactersCharacterIdMedalsGetInner) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *CharactersCharacterIdMedalsGetInner) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *CharactersCharacterIdMedalsGetInner) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.


### GetDate

`func (o *CharactersCharacterIdMedalsGetInner) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CharactersCharacterIdMedalsGetInner) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CharactersCharacterIdMedalsGetInner) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetDescription

`func (o *CharactersCharacterIdMedalsGetInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CharactersCharacterIdMedalsGetInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CharactersCharacterIdMedalsGetInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGraphics

`func (o *CharactersCharacterIdMedalsGetInner) GetGraphics() []CharactersCharacterIdMedalsGetInnerGraphicsInner`

GetGraphics returns the Graphics field if non-nil, zero value otherwise.

### GetGraphicsOk

`func (o *CharactersCharacterIdMedalsGetInner) GetGraphicsOk() (*[]CharactersCharacterIdMedalsGetInnerGraphicsInner, bool)`

GetGraphicsOk returns a tuple with the Graphics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphics

`func (o *CharactersCharacterIdMedalsGetInner) SetGraphics(v []CharactersCharacterIdMedalsGetInnerGraphicsInner)`

SetGraphics sets Graphics field to given value.


### GetIssuerId

`func (o *CharactersCharacterIdMedalsGetInner) GetIssuerId() int64`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *CharactersCharacterIdMedalsGetInner) GetIssuerIdOk() (*int64, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *CharactersCharacterIdMedalsGetInner) SetIssuerId(v int64)`

SetIssuerId sets IssuerId field to given value.


### GetMedalId

`func (o *CharactersCharacterIdMedalsGetInner) GetMedalId() int64`

GetMedalId returns the MedalId field if non-nil, zero value otherwise.

### GetMedalIdOk

`func (o *CharactersCharacterIdMedalsGetInner) GetMedalIdOk() (*int64, bool)`

GetMedalIdOk returns a tuple with the MedalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedalId

`func (o *CharactersCharacterIdMedalsGetInner) SetMedalId(v int64)`

SetMedalId sets MedalId field to given value.


### GetReason

`func (o *CharactersCharacterIdMedalsGetInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CharactersCharacterIdMedalsGetInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CharactersCharacterIdMedalsGetInner) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *CharactersCharacterIdMedalsGetInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CharactersCharacterIdMedalsGetInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CharactersCharacterIdMedalsGetInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *CharactersCharacterIdMedalsGetInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CharactersCharacterIdMedalsGetInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CharactersCharacterIdMedalsGetInner) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


