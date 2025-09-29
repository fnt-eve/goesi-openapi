# WarsWarIdGetAggressor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | Pointer to **int64** | Alliance ID if and only if the aggressor is an alliance | [optional] 
**CorporationId** | Pointer to **int64** | Corporation ID if and only if the aggressor is a corporation | [optional] 
**IskDestroyed** | **float64** | ISK value of ships the aggressor has destroyed | 
**ShipsKilled** | **int64** | The number of ships the aggressor has killed | 

## Methods

### NewWarsWarIdGetAggressor

`func NewWarsWarIdGetAggressor(iskDestroyed float64, shipsKilled int64, ) *WarsWarIdGetAggressor`

NewWarsWarIdGetAggressor instantiates a new WarsWarIdGetAggressor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarsWarIdGetAggressorWithDefaults

`func NewWarsWarIdGetAggressorWithDefaults() *WarsWarIdGetAggressor`

NewWarsWarIdGetAggressorWithDefaults instantiates a new WarsWarIdGetAggressor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *WarsWarIdGetAggressor) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *WarsWarIdGetAggressor) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *WarsWarIdGetAggressor) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.

### HasAllianceId

`func (o *WarsWarIdGetAggressor) HasAllianceId() bool`

HasAllianceId returns a boolean if a field has been set.

### GetCorporationId

`func (o *WarsWarIdGetAggressor) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *WarsWarIdGetAggressor) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *WarsWarIdGetAggressor) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.

### HasCorporationId

`func (o *WarsWarIdGetAggressor) HasCorporationId() bool`

HasCorporationId returns a boolean if a field has been set.

### GetIskDestroyed

`func (o *WarsWarIdGetAggressor) GetIskDestroyed() float64`

GetIskDestroyed returns the IskDestroyed field if non-nil, zero value otherwise.

### GetIskDestroyedOk

`func (o *WarsWarIdGetAggressor) GetIskDestroyedOk() (*float64, bool)`

GetIskDestroyedOk returns a tuple with the IskDestroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIskDestroyed

`func (o *WarsWarIdGetAggressor) SetIskDestroyed(v float64)`

SetIskDestroyed sets IskDestroyed field to given value.


### GetShipsKilled

`func (o *WarsWarIdGetAggressor) GetShipsKilled() int64`

GetShipsKilled returns the ShipsKilled field if non-nil, zero value otherwise.

### GetShipsKilledOk

`func (o *WarsWarIdGetAggressor) GetShipsKilledOk() (*int64, bool)`

GetShipsKilledOk returns a tuple with the ShipsKilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipsKilled

`func (o *WarsWarIdGetAggressor) SetShipsKilled(v int64)`

SetShipsKilled sets ShipsKilled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


