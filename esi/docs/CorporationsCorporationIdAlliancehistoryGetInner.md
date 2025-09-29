# CorporationsCorporationIdAlliancehistoryGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** |  | [optional] 
**IsDeleted** | Pointer to **bool** | True if the alliance has been closed | [optional] 
**RecordId** | **int64** | An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous | 
**StartDate** | **time.Time** |  | 

## Methods

### NewCorporationsCorporationIdAlliancehistoryGetInner

`func NewCorporationsCorporationIdAlliancehistoryGetInner(recordId int64, startDate time.Time, ) *CorporationsCorporationIdAlliancehistoryGetInner`

NewCorporationsCorporationIdAlliancehistoryGetInner instantiates a new CorporationsCorporationIdAlliancehistoryGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdAlliancehistoryGetInnerWithDefaults

`func NewCorporationsCorporationIdAlliancehistoryGetInnerWithDefaults() *CorporationsCorporationIdAlliancehistoryGetInner`

NewCorporationsCorporationIdAlliancehistoryGetInnerWithDefaults instantiates a new CorporationsCorporationIdAlliancehistoryGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetRecordId

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetRecordId() int64`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetRecordIdOk() (*int64, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) SetRecordId(v int64)`

SetRecordId sets RecordId field to given value.


### GetStartDate

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CorporationsCorporationIdAlliancehistoryGetInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


