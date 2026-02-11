# AllianceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorCorporationId** | **int64** |  | 
**CreatorId** | **int64** |  | 
**DateFounded** | **time.Time** | Alliance&#39;s founding date | 
**ExecutorCorporationId** | Pointer to **int64** |  | [optional] 
**FactionId** | Pointer to **int64** |  | [optional] 
**Name** | **string** | Alliance&#39;s name | 
**Ticker** | **string** | Alliance&#39;s ticker | 

## Methods

### NewAllianceDetail

`func NewAllianceDetail(creatorCorporationId int64, creatorId int64, dateFounded time.Time, name string, ticker string, ) *AllianceDetail`

NewAllianceDetail instantiates a new AllianceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllianceDetailWithDefaults

`func NewAllianceDetailWithDefaults() *AllianceDetail`

NewAllianceDetailWithDefaults instantiates a new AllianceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatorCorporationId

`func (o *AllianceDetail) GetCreatorCorporationId() int64`

GetCreatorCorporationId returns the CreatorCorporationId field if non-nil, zero value otherwise.

### GetCreatorCorporationIdOk

`func (o *AllianceDetail) GetCreatorCorporationIdOk() (*int64, bool)`

GetCreatorCorporationIdOk returns a tuple with the CreatorCorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCorporationId

`func (o *AllianceDetail) SetCreatorCorporationId(v int64)`

SetCreatorCorporationId sets CreatorCorporationId field to given value.


### GetCreatorId

`func (o *AllianceDetail) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AllianceDetail) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AllianceDetail) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.


### GetDateFounded

`func (o *AllianceDetail) GetDateFounded() time.Time`

GetDateFounded returns the DateFounded field if non-nil, zero value otherwise.

### GetDateFoundedOk

`func (o *AllianceDetail) GetDateFoundedOk() (*time.Time, bool)`

GetDateFoundedOk returns a tuple with the DateFounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFounded

`func (o *AllianceDetail) SetDateFounded(v time.Time)`

SetDateFounded sets DateFounded field to given value.


### GetExecutorCorporationId

`func (o *AllianceDetail) GetExecutorCorporationId() int64`

GetExecutorCorporationId returns the ExecutorCorporationId field if non-nil, zero value otherwise.

### GetExecutorCorporationIdOk

`func (o *AllianceDetail) GetExecutorCorporationIdOk() (*int64, bool)`

GetExecutorCorporationIdOk returns a tuple with the ExecutorCorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutorCorporationId

`func (o *AllianceDetail) SetExecutorCorporationId(v int64)`

SetExecutorCorporationId sets ExecutorCorporationId field to given value.

### HasExecutorCorporationId

`func (o *AllianceDetail) HasExecutorCorporationId() bool`

HasExecutorCorporationId returns a boolean if a field has been set.

### GetFactionId

`func (o *AllianceDetail) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *AllianceDetail) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *AllianceDetail) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *AllianceDetail) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetName

`func (o *AllianceDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllianceDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllianceDetail) SetName(v string)`

SetName sets Name field to given value.


### GetTicker

`func (o *AllianceDetail) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *AllianceDetail) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *AllianceDetail) SetTicker(v string)`

SetTicker sets Ticker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


