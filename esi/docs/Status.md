# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Players** | **int64** | Number of characters currently logged in | 
**ServerVersion** | **string** | Build number of the cluster | 
**StartTime** | **time.Time** | Moment the cluster started accepting connections | 
**Vip** | **bool** | Whether the cluster only accepts VIP logins | 

## Methods

### NewStatus

`func NewStatus(players int64, serverVersion string, startTime time.Time, vip bool, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayers

`func (o *Status) GetPlayers() int64`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *Status) GetPlayersOk() (*int64, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *Status) SetPlayers(v int64)`

SetPlayers sets Players field to given value.


### GetServerVersion

`func (o *Status) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *Status) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *Status) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetStartTime

`func (o *Status) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Status) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Status) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetVip

`func (o *Status) GetVip() bool`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *Status) GetVipOk() (*bool, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *Status) SetVip(v bool)`

SetVip sets Vip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


