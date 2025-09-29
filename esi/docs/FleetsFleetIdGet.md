# FleetsFleetIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFreeMove** | **bool** | Is free-move enabled | 
**IsRegistered** | **bool** | Does the fleet have an active fleet advertisement | 
**IsVoiceEnabled** | **bool** | Is EVE Voice enabled | 
**Motd** | **string** | Fleet MOTD in CCP flavoured HTML | 

## Methods

### NewFleetsFleetIdGet

`func NewFleetsFleetIdGet(isFreeMove bool, isRegistered bool, isVoiceEnabled bool, motd string, ) *FleetsFleetIdGet`

NewFleetsFleetIdGet instantiates a new FleetsFleetIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetsFleetIdGetWithDefaults

`func NewFleetsFleetIdGetWithDefaults() *FleetsFleetIdGet`

NewFleetsFleetIdGetWithDefaults instantiates a new FleetsFleetIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFreeMove

`func (o *FleetsFleetIdGet) GetIsFreeMove() bool`

GetIsFreeMove returns the IsFreeMove field if non-nil, zero value otherwise.

### GetIsFreeMoveOk

`func (o *FleetsFleetIdGet) GetIsFreeMoveOk() (*bool, bool)`

GetIsFreeMoveOk returns a tuple with the IsFreeMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeMove

`func (o *FleetsFleetIdGet) SetIsFreeMove(v bool)`

SetIsFreeMove sets IsFreeMove field to given value.


### GetIsRegistered

`func (o *FleetsFleetIdGet) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *FleetsFleetIdGet) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *FleetsFleetIdGet) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.


### GetIsVoiceEnabled

`func (o *FleetsFleetIdGet) GetIsVoiceEnabled() bool`

GetIsVoiceEnabled returns the IsVoiceEnabled field if non-nil, zero value otherwise.

### GetIsVoiceEnabledOk

`func (o *FleetsFleetIdGet) GetIsVoiceEnabledOk() (*bool, bool)`

GetIsVoiceEnabledOk returns a tuple with the IsVoiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoiceEnabled

`func (o *FleetsFleetIdGet) SetIsVoiceEnabled(v bool)`

SetIsVoiceEnabled sets IsVoiceEnabled field to given value.


### GetMotd

`func (o *FleetsFleetIdGet) GetMotd() string`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *FleetsFleetIdGet) GetMotdOk() (*string, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *FleetsFleetIdGet) SetMotd(v string)`

SetMotd sets Motd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


