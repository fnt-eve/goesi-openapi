# CharactersCharacterIdCorporationhistoryGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorporationId** | **int64** |  | 
**IsDeleted** | Pointer to **bool** | True if the corporation has been deleted | [optional] 
**RecordId** | **int64** | An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous | 
**StartDate** | **time.Time** |  | 

## Methods

### NewCharactersCharacterIdCorporationhistoryGetInner

`func NewCharactersCharacterIdCorporationhistoryGetInner(corporationId int64, recordId int64, startDate time.Time, ) *CharactersCharacterIdCorporationhistoryGetInner`

NewCharactersCharacterIdCorporationhistoryGetInner instantiates a new CharactersCharacterIdCorporationhistoryGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdCorporationhistoryGetInnerWithDefaults

`func NewCharactersCharacterIdCorporationhistoryGetInnerWithDefaults() *CharactersCharacterIdCorporationhistoryGetInner`

NewCharactersCharacterIdCorporationhistoryGetInnerWithDefaults instantiates a new CharactersCharacterIdCorporationhistoryGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorporationId

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *CharactersCharacterIdCorporationhistoryGetInner) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.


### GetIsDeleted

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CharactersCharacterIdCorporationhistoryGetInner) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CharactersCharacterIdCorporationhistoryGetInner) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetRecordId

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetRecordId() int64`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetRecordIdOk() (*int64, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *CharactersCharacterIdCorporationhistoryGetInner) SetRecordId(v int64)`

SetRecordId sets RecordId field to given value.


### GetStartDate

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CharactersCharacterIdCorporationhistoryGetInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CharactersCharacterIdCorporationhistoryGetInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


