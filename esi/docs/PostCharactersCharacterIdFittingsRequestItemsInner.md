# PostCharactersCharacterIdFittingsRequestItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flag** | **string** | Fitting location for the item. Entries placed in &#39;Invalid&#39; will be discarded. If this leaves the fitting with nothing, it will cause an error. | 
**Quantity** | **int64** |  | 
**TypeId** | **int64** |  | 

## Methods

### NewPostCharactersCharacterIdFittingsRequestItemsInner

`func NewPostCharactersCharacterIdFittingsRequestItemsInner(flag string, quantity int64, typeId int64, ) *PostCharactersCharacterIdFittingsRequestItemsInner`

NewPostCharactersCharacterIdFittingsRequestItemsInner instantiates a new PostCharactersCharacterIdFittingsRequestItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCharactersCharacterIdFittingsRequestItemsInnerWithDefaults

`func NewPostCharactersCharacterIdFittingsRequestItemsInnerWithDefaults() *PostCharactersCharacterIdFittingsRequestItemsInner`

NewPostCharactersCharacterIdFittingsRequestItemsInnerWithDefaults instantiates a new PostCharactersCharacterIdFittingsRequestItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlag

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) SetFlag(v string)`

SetFlag sets Flag field to given value.


### GetQuantity

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetTypeId

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *PostCharactersCharacterIdFittingsRequestItemsInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


