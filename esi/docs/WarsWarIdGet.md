# WarsWarIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggressor** | [**WarsWarIdGetAggressor**](WarsWarIdGetAggressor.md) |  | 
**Allies** | Pointer to [**[]WarsWarIdGetAlliesInner**](WarsWarIdGetAlliesInner.md) | allied corporations or alliances, each object contains either corporation_id or alliance_id | [optional] 
**Declared** | **time.Time** | Time that the war was declared | 
**Defender** | [**WarsWarIdGetDefender**](WarsWarIdGetDefender.md) |  | 
**Finished** | Pointer to **time.Time** | Time the war ended and shooting was no longer allowed | [optional] 
**Id** | **int64** | ID of the specified war | 
**Mutual** | **bool** | Was the war declared mutual by both parties | 
**OpenForAllies** | **bool** | Is the war currently open for allies or not | 
**Retracted** | Pointer to **time.Time** | Time the war was retracted but both sides could still shoot each other | [optional] 
**Started** | Pointer to **time.Time** | Time when the war started and both sides could shoot each other | [optional] 

## Methods

### NewWarsWarIdGet

`func NewWarsWarIdGet(aggressor WarsWarIdGetAggressor, declared time.Time, defender WarsWarIdGetDefender, id int64, mutual bool, openForAllies bool, ) *WarsWarIdGet`

NewWarsWarIdGet instantiates a new WarsWarIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarsWarIdGetWithDefaults

`func NewWarsWarIdGetWithDefaults() *WarsWarIdGet`

NewWarsWarIdGetWithDefaults instantiates a new WarsWarIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggressor

`func (o *WarsWarIdGet) GetAggressor() WarsWarIdGetAggressor`

GetAggressor returns the Aggressor field if non-nil, zero value otherwise.

### GetAggressorOk

`func (o *WarsWarIdGet) GetAggressorOk() (*WarsWarIdGetAggressor, bool)`

GetAggressorOk returns a tuple with the Aggressor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggressor

`func (o *WarsWarIdGet) SetAggressor(v WarsWarIdGetAggressor)`

SetAggressor sets Aggressor field to given value.


### GetAllies

`func (o *WarsWarIdGet) GetAllies() []WarsWarIdGetAlliesInner`

GetAllies returns the Allies field if non-nil, zero value otherwise.

### GetAlliesOk

`func (o *WarsWarIdGet) GetAlliesOk() (*[]WarsWarIdGetAlliesInner, bool)`

GetAlliesOk returns a tuple with the Allies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllies

`func (o *WarsWarIdGet) SetAllies(v []WarsWarIdGetAlliesInner)`

SetAllies sets Allies field to given value.

### HasAllies

`func (o *WarsWarIdGet) HasAllies() bool`

HasAllies returns a boolean if a field has been set.

### GetDeclared

`func (o *WarsWarIdGet) GetDeclared() time.Time`

GetDeclared returns the Declared field if non-nil, zero value otherwise.

### GetDeclaredOk

`func (o *WarsWarIdGet) GetDeclaredOk() (*time.Time, bool)`

GetDeclaredOk returns a tuple with the Declared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclared

`func (o *WarsWarIdGet) SetDeclared(v time.Time)`

SetDeclared sets Declared field to given value.


### GetDefender

`func (o *WarsWarIdGet) GetDefender() WarsWarIdGetDefender`

GetDefender returns the Defender field if non-nil, zero value otherwise.

### GetDefenderOk

`func (o *WarsWarIdGet) GetDefenderOk() (*WarsWarIdGetDefender, bool)`

GetDefenderOk returns a tuple with the Defender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefender

`func (o *WarsWarIdGet) SetDefender(v WarsWarIdGetDefender)`

SetDefender sets Defender field to given value.


### GetFinished

`func (o *WarsWarIdGet) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *WarsWarIdGet) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *WarsWarIdGet) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *WarsWarIdGet) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *WarsWarIdGet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarsWarIdGet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarsWarIdGet) SetId(v int64)`

SetId sets Id field to given value.


### GetMutual

`func (o *WarsWarIdGet) GetMutual() bool`

GetMutual returns the Mutual field if non-nil, zero value otherwise.

### GetMutualOk

`func (o *WarsWarIdGet) GetMutualOk() (*bool, bool)`

GetMutualOk returns a tuple with the Mutual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutual

`func (o *WarsWarIdGet) SetMutual(v bool)`

SetMutual sets Mutual field to given value.


### GetOpenForAllies

`func (o *WarsWarIdGet) GetOpenForAllies() bool`

GetOpenForAllies returns the OpenForAllies field if non-nil, zero value otherwise.

### GetOpenForAlliesOk

`func (o *WarsWarIdGet) GetOpenForAlliesOk() (*bool, bool)`

GetOpenForAlliesOk returns a tuple with the OpenForAllies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenForAllies

`func (o *WarsWarIdGet) SetOpenForAllies(v bool)`

SetOpenForAllies sets OpenForAllies field to given value.


### GetRetracted

`func (o *WarsWarIdGet) GetRetracted() time.Time`

GetRetracted returns the Retracted field if non-nil, zero value otherwise.

### GetRetractedOk

`func (o *WarsWarIdGet) GetRetractedOk() (*time.Time, bool)`

GetRetractedOk returns a tuple with the Retracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetracted

`func (o *WarsWarIdGet) SetRetracted(v time.Time)`

SetRetracted sets Retracted field to given value.

### HasRetracted

`func (o *WarsWarIdGet) HasRetracted() bool`

HasRetracted returns a boolean if a field has been set.

### GetStarted

`func (o *WarsWarIdGet) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *WarsWarIdGet) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *WarsWarIdGet) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *WarsWarIdGet) HasStarted() bool`

HasStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


