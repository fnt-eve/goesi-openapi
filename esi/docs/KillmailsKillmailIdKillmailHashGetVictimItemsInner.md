# KillmailsKillmailIdKillmailHashGetVictimItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flag** | **int64** | Flag for the location of the item  | 
**ItemTypeId** | **int64** |  | 
**Items** | Pointer to [**[]KillmailsKillmailIdKillmailHashGetVictimItemsInnerItemsInner**](KillmailsKillmailIdKillmailHashGetVictimItemsInnerItemsInner.md) |  | [optional] 
**QuantityDestroyed** | Pointer to **int64** | How many of the item were destroyed if any  | [optional] 
**QuantityDropped** | Pointer to **int64** | How many of the item were dropped if any  | [optional] 
**Singleton** | **int64** |  | 

## Methods

### NewKillmailsKillmailIdKillmailHashGetVictimItemsInner

`func NewKillmailsKillmailIdKillmailHashGetVictimItemsInner(flag int64, itemTypeId int64, singleton int64, ) *KillmailsKillmailIdKillmailHashGetVictimItemsInner`

NewKillmailsKillmailIdKillmailHashGetVictimItemsInner instantiates a new KillmailsKillmailIdKillmailHashGetVictimItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKillmailsKillmailIdKillmailHashGetVictimItemsInnerWithDefaults

`func NewKillmailsKillmailIdKillmailHashGetVictimItemsInnerWithDefaults() *KillmailsKillmailIdKillmailHashGetVictimItemsInner`

NewKillmailsKillmailIdKillmailHashGetVictimItemsInnerWithDefaults instantiates a new KillmailsKillmailIdKillmailHashGetVictimItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlag

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetFlag() int64`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetFlagOk() (*int64, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) SetFlag(v int64)`

SetFlag sets Flag field to given value.


### GetItemTypeId

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetItemTypeId() int64`

GetItemTypeId returns the ItemTypeId field if non-nil, zero value otherwise.

### GetItemTypeIdOk

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetItemTypeIdOk() (*int64, bool)`

GetItemTypeIdOk returns a tuple with the ItemTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypeId

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) SetItemTypeId(v int64)`

SetItemTypeId sets ItemTypeId field to given value.


### GetItems

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetItems() []KillmailsKillmailIdKillmailHashGetVictimItemsInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetItemsOk() (*[]KillmailsKillmailIdKillmailHashGetVictimItemsInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) SetItems(v []KillmailsKillmailIdKillmailHashGetVictimItemsInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetQuantityDestroyed

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetQuantityDestroyed() int64`

GetQuantityDestroyed returns the QuantityDestroyed field if non-nil, zero value otherwise.

### GetQuantityDestroyedOk

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetQuantityDestroyedOk() (*int64, bool)`

GetQuantityDestroyedOk returns a tuple with the QuantityDestroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDestroyed

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) SetQuantityDestroyed(v int64)`

SetQuantityDestroyed sets QuantityDestroyed field to given value.

### HasQuantityDestroyed

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) HasQuantityDestroyed() bool`

HasQuantityDestroyed returns a boolean if a field has been set.

### GetQuantityDropped

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetQuantityDropped() int64`

GetQuantityDropped returns the QuantityDropped field if non-nil, zero value otherwise.

### GetQuantityDroppedOk

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetQuantityDroppedOk() (*int64, bool)`

GetQuantityDroppedOk returns a tuple with the QuantityDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDropped

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) SetQuantityDropped(v int64)`

SetQuantityDropped sets QuantityDropped field to given value.

### HasQuantityDropped

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) HasQuantityDropped() bool`

HasQuantityDropped returns a boolean if a field has been set.

### GetSingleton

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetSingleton() int64`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) GetSingletonOk() (*int64, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *KillmailsKillmailIdKillmailHashGetVictimItemsInner) SetSingleton(v int64)`

SetSingleton sets Singleton field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


