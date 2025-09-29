# AlliancesAllianceIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorCorporationId** | **int64** | ID of the corporation that created the alliance | 
**CreatorId** | **int64** | ID of the character that created the alliance | 
**DateFounded** | **time.Time** |  | 
**ExecutorCorporationId** | Pointer to **int64** | the executor corporation ID, if this alliance is not closed | [optional] 
**FactionId** | Pointer to **int64** | Faction ID this alliance is fighting for, if this alliance is enlisted in factional warfare | [optional] 
**Name** | **string** | the full name of the alliance | 
**Ticker** | **string** | the short name of the alliance | 

## Methods

### NewAlliancesAllianceIdGet

`func NewAlliancesAllianceIdGet(creatorCorporationId int64, creatorId int64, dateFounded time.Time, name string, ticker string, ) *AlliancesAllianceIdGet`

NewAlliancesAllianceIdGet instantiates a new AlliancesAllianceIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlliancesAllianceIdGetWithDefaults

`func NewAlliancesAllianceIdGetWithDefaults() *AlliancesAllianceIdGet`

NewAlliancesAllianceIdGetWithDefaults instantiates a new AlliancesAllianceIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatorCorporationId

`func (o *AlliancesAllianceIdGet) GetCreatorCorporationId() int64`

GetCreatorCorporationId returns the CreatorCorporationId field if non-nil, zero value otherwise.

### GetCreatorCorporationIdOk

`func (o *AlliancesAllianceIdGet) GetCreatorCorporationIdOk() (*int64, bool)`

GetCreatorCorporationIdOk returns a tuple with the CreatorCorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCorporationId

`func (o *AlliancesAllianceIdGet) SetCreatorCorporationId(v int64)`

SetCreatorCorporationId sets CreatorCorporationId field to given value.


### GetCreatorId

`func (o *AlliancesAllianceIdGet) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AlliancesAllianceIdGet) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AlliancesAllianceIdGet) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.


### GetDateFounded

`func (o *AlliancesAllianceIdGet) GetDateFounded() time.Time`

GetDateFounded returns the DateFounded field if non-nil, zero value otherwise.

### GetDateFoundedOk

`func (o *AlliancesAllianceIdGet) GetDateFoundedOk() (*time.Time, bool)`

GetDateFoundedOk returns a tuple with the DateFounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFounded

`func (o *AlliancesAllianceIdGet) SetDateFounded(v time.Time)`

SetDateFounded sets DateFounded field to given value.


### GetExecutorCorporationId

`func (o *AlliancesAllianceIdGet) GetExecutorCorporationId() int64`

GetExecutorCorporationId returns the ExecutorCorporationId field if non-nil, zero value otherwise.

### GetExecutorCorporationIdOk

`func (o *AlliancesAllianceIdGet) GetExecutorCorporationIdOk() (*int64, bool)`

GetExecutorCorporationIdOk returns a tuple with the ExecutorCorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutorCorporationId

`func (o *AlliancesAllianceIdGet) SetExecutorCorporationId(v int64)`

SetExecutorCorporationId sets ExecutorCorporationId field to given value.

### HasExecutorCorporationId

`func (o *AlliancesAllianceIdGet) HasExecutorCorporationId() bool`

HasExecutorCorporationId returns a boolean if a field has been set.

### GetFactionId

`func (o *AlliancesAllianceIdGet) GetFactionId() int64`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *AlliancesAllianceIdGet) GetFactionIdOk() (*int64, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *AlliancesAllianceIdGet) SetFactionId(v int64)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *AlliancesAllianceIdGet) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetName

`func (o *AlliancesAllianceIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlliancesAllianceIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlliancesAllianceIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetTicker

`func (o *AlliancesAllianceIdGet) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *AlliancesAllianceIdGet) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *AlliancesAllianceIdGet) SetTicker(v string)`

SetTicker sets Ticker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


