# CharactersCharacterIdShipGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipItemId** | **int64** | Item id&#39;s are unique to a ship and persist until it is repackaged. This value can be used to track repeated uses of a ship, or detect when a pilot changes into a different instance of the same ship type. | 
**ShipName** | **string** |  | 
**ShipTypeId** | **int64** |  | 

## Methods

### NewCharactersCharacterIdShipGet

`func NewCharactersCharacterIdShipGet(shipItemId int64, shipName string, shipTypeId int64, ) *CharactersCharacterIdShipGet`

NewCharactersCharacterIdShipGet instantiates a new CharactersCharacterIdShipGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdShipGetWithDefaults

`func NewCharactersCharacterIdShipGetWithDefaults() *CharactersCharacterIdShipGet`

NewCharactersCharacterIdShipGetWithDefaults instantiates a new CharactersCharacterIdShipGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipItemId

`func (o *CharactersCharacterIdShipGet) GetShipItemId() int64`

GetShipItemId returns the ShipItemId field if non-nil, zero value otherwise.

### GetShipItemIdOk

`func (o *CharactersCharacterIdShipGet) GetShipItemIdOk() (*int64, bool)`

GetShipItemIdOk returns a tuple with the ShipItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipItemId

`func (o *CharactersCharacterIdShipGet) SetShipItemId(v int64)`

SetShipItemId sets ShipItemId field to given value.


### GetShipName

`func (o *CharactersCharacterIdShipGet) GetShipName() string`

GetShipName returns the ShipName field if non-nil, zero value otherwise.

### GetShipNameOk

`func (o *CharactersCharacterIdShipGet) GetShipNameOk() (*string, bool)`

GetShipNameOk returns a tuple with the ShipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipName

`func (o *CharactersCharacterIdShipGet) SetShipName(v string)`

SetShipName sets ShipName field to given value.


### GetShipTypeId

`func (o *CharactersCharacterIdShipGet) GetShipTypeId() int64`

GetShipTypeId returns the ShipTypeId field if non-nil, zero value otherwise.

### GetShipTypeIdOk

`func (o *CharactersCharacterIdShipGet) GetShipTypeIdOk() (*int64, bool)`

GetShipTypeIdOk returns a tuple with the ShipTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTypeId

`func (o *CharactersCharacterIdShipGet) SetShipTypeId(v int64)`

SetShipTypeId sets ShipTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


