# CorporationsCorporationIdContainersLogsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**CharacterId** | **int64** | ID of the character who performed the action. | 
**ContainerId** | **int64** | ID of the container | 
**ContainerTypeId** | **int64** | Type ID of the container | 
**LocationFlag** | **string** |  | 
**LocationId** | **int64** |  | 
**LoggedAt** | **time.Time** | Timestamp when this log was created | 
**NewConfigBitmask** | Pointer to **int64** |  | [optional] 
**OldConfigBitmask** | Pointer to **int64** |  | [optional] 
**PasswordType** | Pointer to **string** | Type of password set if action is of type SetPassword or EnterPassword | [optional] 
**Quantity** | Pointer to **int64** | Quantity of the item being acted upon | [optional] 
**TypeId** | Pointer to **int64** | Type ID of the item being acted upon | [optional] 

## Methods

### NewCorporationsCorporationIdContainersLogsGetInner

`func NewCorporationsCorporationIdContainersLogsGetInner(action string, characterId int64, containerId int64, containerTypeId int64, locationFlag string, locationId int64, loggedAt time.Time, ) *CorporationsCorporationIdContainersLogsGetInner`

NewCorporationsCorporationIdContainersLogsGetInner instantiates a new CorporationsCorporationIdContainersLogsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdContainersLogsGetInnerWithDefaults

`func NewCorporationsCorporationIdContainersLogsGetInnerWithDefaults() *CorporationsCorporationIdContainersLogsGetInner`

NewCorporationsCorporationIdContainersLogsGetInnerWithDefaults instantiates a new CorporationsCorporationIdContainersLogsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetAction(v string)`

SetAction sets Action field to given value.


### GetCharacterId

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetCharacterId() int64`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetCharacterIdOk() (*int64, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetCharacterId(v int64)`

SetCharacterId sets CharacterId field to given value.


### GetContainerId

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.


### GetContainerTypeId

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetContainerTypeId() int64`

GetContainerTypeId returns the ContainerTypeId field if non-nil, zero value otherwise.

### GetContainerTypeIdOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetContainerTypeIdOk() (*int64, bool)`

GetContainerTypeIdOk returns a tuple with the ContainerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeId

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetContainerTypeId(v int64)`

SetContainerTypeId sets ContainerTypeId field to given value.


### GetLocationFlag

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetLocationFlag() string`

GetLocationFlag returns the LocationFlag field if non-nil, zero value otherwise.

### GetLocationFlagOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetLocationFlagOk() (*string, bool)`

GetLocationFlagOk returns a tuple with the LocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationFlag

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetLocationFlag(v string)`

SetLocationFlag sets LocationFlag field to given value.


### GetLocationId

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetLoggedAt

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetNewConfigBitmask

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetNewConfigBitmask() int64`

GetNewConfigBitmask returns the NewConfigBitmask field if non-nil, zero value otherwise.

### GetNewConfigBitmaskOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetNewConfigBitmaskOk() (*int64, bool)`

GetNewConfigBitmaskOk returns a tuple with the NewConfigBitmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewConfigBitmask

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetNewConfigBitmask(v int64)`

SetNewConfigBitmask sets NewConfigBitmask field to given value.

### HasNewConfigBitmask

`func (o *CorporationsCorporationIdContainersLogsGetInner) HasNewConfigBitmask() bool`

HasNewConfigBitmask returns a boolean if a field has been set.

### GetOldConfigBitmask

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetOldConfigBitmask() int64`

GetOldConfigBitmask returns the OldConfigBitmask field if non-nil, zero value otherwise.

### GetOldConfigBitmaskOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetOldConfigBitmaskOk() (*int64, bool)`

GetOldConfigBitmaskOk returns a tuple with the OldConfigBitmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldConfigBitmask

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetOldConfigBitmask(v int64)`

SetOldConfigBitmask sets OldConfigBitmask field to given value.

### HasOldConfigBitmask

`func (o *CorporationsCorporationIdContainersLogsGetInner) HasOldConfigBitmask() bool`

HasOldConfigBitmask returns a boolean if a field has been set.

### GetPasswordType

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *CorporationsCorporationIdContainersLogsGetInner) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetQuantity

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CorporationsCorporationIdContainersLogsGetInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTypeId

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsCorporationIdContainersLogsGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsCorporationIdContainersLogsGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *CorporationsCorporationIdContainersLogsGetInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


