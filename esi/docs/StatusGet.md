# StatusGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Players** | **int32** | Current online player count | 
**ServerVersion** | **string** | Running version as string | 
**StartTime** | **time.Time** | Server start timestamp | 
**Vip** | Pointer to **bool** | If the server is in VIP mode | [optional] 

## Methods

### NewStatusGet

`func NewStatusGet(players int32, serverVersion string, startTime time.Time, ) *StatusGet`

NewStatusGet instantiates a new StatusGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusGetWithDefaults

`func NewStatusGetWithDefaults() *StatusGet`

NewStatusGetWithDefaults instantiates a new StatusGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayers

`func (o *StatusGet) GetPlayers() int32`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *StatusGet) GetPlayersOk() (*int32, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *StatusGet) SetPlayers(v int32)`

SetPlayers sets Players field to given value.


### GetServerVersion

`func (o *StatusGet) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *StatusGet) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *StatusGet) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetStartTime

`func (o *StatusGet) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *StatusGet) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *StatusGet) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetVip

`func (o *StatusGet) GetVip() bool`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *StatusGet) GetVipOk() (*bool, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *StatusGet) SetVip(v bool)`

SetVip sets Vip field to given value.

### HasVip

`func (o *StatusGet) HasVip() bool`

HasVip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


