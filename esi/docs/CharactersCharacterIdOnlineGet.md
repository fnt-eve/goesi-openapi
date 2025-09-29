# CharactersCharacterIdOnlineGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogin** | Pointer to **time.Time** | Timestamp of the last login | [optional] 
**LastLogout** | Pointer to **time.Time** | Timestamp of the last logout | [optional] 
**Logins** | Pointer to **int64** | Total number of times the character has logged in | [optional] 
**Online** | **bool** | If the character is online | 

## Methods

### NewCharactersCharacterIdOnlineGet

`func NewCharactersCharacterIdOnlineGet(online bool, ) *CharactersCharacterIdOnlineGet`

NewCharactersCharacterIdOnlineGet instantiates a new CharactersCharacterIdOnlineGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdOnlineGetWithDefaults

`func NewCharactersCharacterIdOnlineGetWithDefaults() *CharactersCharacterIdOnlineGet`

NewCharactersCharacterIdOnlineGetWithDefaults instantiates a new CharactersCharacterIdOnlineGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogin

`func (o *CharactersCharacterIdOnlineGet) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *CharactersCharacterIdOnlineGet) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *CharactersCharacterIdOnlineGet) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *CharactersCharacterIdOnlineGet) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastLogout

`func (o *CharactersCharacterIdOnlineGet) GetLastLogout() time.Time`

GetLastLogout returns the LastLogout field if non-nil, zero value otherwise.

### GetLastLogoutOk

`func (o *CharactersCharacterIdOnlineGet) GetLastLogoutOk() (*time.Time, bool)`

GetLastLogoutOk returns a tuple with the LastLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogout

`func (o *CharactersCharacterIdOnlineGet) SetLastLogout(v time.Time)`

SetLastLogout sets LastLogout field to given value.

### HasLastLogout

`func (o *CharactersCharacterIdOnlineGet) HasLastLogout() bool`

HasLastLogout returns a boolean if a field has been set.

### GetLogins

`func (o *CharactersCharacterIdOnlineGet) GetLogins() int64`

GetLogins returns the Logins field if non-nil, zero value otherwise.

### GetLoginsOk

`func (o *CharactersCharacterIdOnlineGet) GetLoginsOk() (*int64, bool)`

GetLoginsOk returns a tuple with the Logins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogins

`func (o *CharactersCharacterIdOnlineGet) SetLogins(v int64)`

SetLogins sets Logins field to given value.

### HasLogins

`func (o *CharactersCharacterIdOnlineGet) HasLogins() bool`

HasLogins returns a boolean if a field has been set.

### GetOnline

`func (o *CharactersCharacterIdOnlineGet) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *CharactersCharacterIdOnlineGet) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *CharactersCharacterIdOnlineGet) SetOnline(v bool)`

SetOnline sets Online field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


