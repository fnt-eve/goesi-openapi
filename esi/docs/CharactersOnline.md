# CharactersOnline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogin** | Pointer to **time.Time** | When the character last logged in | [optional] 
**LastLogout** | Pointer to **time.Time** | When the character last logged out | [optional] 
**Logins** | Pointer to **int64** | Total logins (all-time) | [optional] 
**Online** | **bool** | Whether the character is online | 

## Methods

### NewCharactersOnline

`func NewCharactersOnline(online bool, ) *CharactersOnline`

NewCharactersOnline instantiates a new CharactersOnline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersOnlineWithDefaults

`func NewCharactersOnlineWithDefaults() *CharactersOnline`

NewCharactersOnlineWithDefaults instantiates a new CharactersOnline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogin

`func (o *CharactersOnline) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *CharactersOnline) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *CharactersOnline) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *CharactersOnline) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastLogout

`func (o *CharactersOnline) GetLastLogout() time.Time`

GetLastLogout returns the LastLogout field if non-nil, zero value otherwise.

### GetLastLogoutOk

`func (o *CharactersOnline) GetLastLogoutOk() (*time.Time, bool)`

GetLastLogoutOk returns a tuple with the LastLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogout

`func (o *CharactersOnline) SetLastLogout(v time.Time)`

SetLastLogout sets LastLogout field to given value.

### HasLastLogout

`func (o *CharactersOnline) HasLastLogout() bool`

HasLastLogout returns a boolean if a field has been set.

### GetLogins

`func (o *CharactersOnline) GetLogins() int64`

GetLogins returns the Logins field if non-nil, zero value otherwise.

### GetLoginsOk

`func (o *CharactersOnline) GetLoginsOk() (*int64, bool)`

GetLoginsOk returns a tuple with the Logins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogins

`func (o *CharactersOnline) SetLogins(v int64)`

SetLogins sets Logins field to given value.

### HasLogins

`func (o *CharactersOnline) HasLogins() bool`

HasLogins returns a boolean if a field has been set.

### GetOnline

`func (o *CharactersOnline) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *CharactersOnline) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *CharactersOnline) SetOnline(v bool)`

SetOnline sets Online field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


