# PostCharactersCharacterIdFittingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Items** | [**[]PostCharactersCharacterIdFittingsRequestItemsInner**](PostCharactersCharacterIdFittingsRequestItemsInner.md) |  | 
**Name** | **string** |  | 
**ShipTypeId** | **int64** |  | 

## Methods

### NewPostCharactersCharacterIdFittingsRequest

`func NewPostCharactersCharacterIdFittingsRequest(description string, items []PostCharactersCharacterIdFittingsRequestItemsInner, name string, shipTypeId int64, ) *PostCharactersCharacterIdFittingsRequest`

NewPostCharactersCharacterIdFittingsRequest instantiates a new PostCharactersCharacterIdFittingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCharactersCharacterIdFittingsRequestWithDefaults

`func NewPostCharactersCharacterIdFittingsRequestWithDefaults() *PostCharactersCharacterIdFittingsRequest`

NewPostCharactersCharacterIdFittingsRequestWithDefaults instantiates a new PostCharactersCharacterIdFittingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PostCharactersCharacterIdFittingsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostCharactersCharacterIdFittingsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostCharactersCharacterIdFittingsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetItems

`func (o *PostCharactersCharacterIdFittingsRequest) GetItems() []PostCharactersCharacterIdFittingsRequestItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PostCharactersCharacterIdFittingsRequest) GetItemsOk() (*[]PostCharactersCharacterIdFittingsRequestItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PostCharactersCharacterIdFittingsRequest) SetItems(v []PostCharactersCharacterIdFittingsRequestItemsInner)`

SetItems sets Items field to given value.


### GetName

`func (o *PostCharactersCharacterIdFittingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCharactersCharacterIdFittingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCharactersCharacterIdFittingsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShipTypeId

`func (o *PostCharactersCharacterIdFittingsRequest) GetShipTypeId() int64`

GetShipTypeId returns the ShipTypeId field if non-nil, zero value otherwise.

### GetShipTypeIdOk

`func (o *PostCharactersCharacterIdFittingsRequest) GetShipTypeIdOk() (*int64, bool)`

GetShipTypeIdOk returns a tuple with the ShipTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTypeId

`func (o *PostCharactersCharacterIdFittingsRequest) SetShipTypeId(v int64)`

SetShipTypeId sets ShipTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


