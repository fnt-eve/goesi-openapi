# CharactersCharacterIdContractsContractIdItemsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsIncluded** | **bool** | true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract | 
**IsSingleton** | **bool** |  | 
**Quantity** | **int64** | Number of items in the stack | 
**RawQuantity** | Pointer to **int64** | -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy | [optional] 
**RecordId** | **int64** | Unique ID for the item | 
**TypeId** | **int64** | Type ID for item | 

## Methods

### NewCharactersCharacterIdContractsContractIdItemsGetInner

`func NewCharactersCharacterIdContractsContractIdItemsGetInner(isIncluded bool, isSingleton bool, quantity int64, recordId int64, typeId int64, ) *CharactersCharacterIdContractsContractIdItemsGetInner`

NewCharactersCharacterIdContractsContractIdItemsGetInner instantiates a new CharactersCharacterIdContractsContractIdItemsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdContractsContractIdItemsGetInnerWithDefaults

`func NewCharactersCharacterIdContractsContractIdItemsGetInnerWithDefaults() *CharactersCharacterIdContractsContractIdItemsGetInner`

NewCharactersCharacterIdContractsContractIdItemsGetInnerWithDefaults instantiates a new CharactersCharacterIdContractsContractIdItemsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsIncluded

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetIsIncluded() bool`

GetIsIncluded returns the IsIncluded field if non-nil, zero value otherwise.

### GetIsIncludedOk

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetIsIncludedOk() (*bool, bool)`

GetIsIncludedOk returns a tuple with the IsIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncluded

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) SetIsIncluded(v bool)`

SetIsIncluded sets IsIncluded field to given value.


### GetIsSingleton

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetIsSingleton() bool`

GetIsSingleton returns the IsSingleton field if non-nil, zero value otherwise.

### GetIsSingletonOk

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetIsSingletonOk() (*bool, bool)`

GetIsSingletonOk returns a tuple with the IsSingleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleton

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) SetIsSingleton(v bool)`

SetIsSingleton sets IsSingleton field to given value.


### GetQuantity

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetRawQuantity

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetRawQuantity() int64`

GetRawQuantity returns the RawQuantity field if non-nil, zero value otherwise.

### GetRawQuantityOk

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetRawQuantityOk() (*int64, bool)`

GetRawQuantityOk returns a tuple with the RawQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuantity

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) SetRawQuantity(v int64)`

SetRawQuantity sets RawQuantity field to given value.

### HasRawQuantity

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) HasRawQuantity() bool`

HasRawQuantity returns a boolean if a field has been set.

### GetRecordId

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetRecordId() int64`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetRecordIdOk() (*int64, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) SetRecordId(v int64)`

SetRecordId sets RecordId field to given value.


### GetTypeId

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CharactersCharacterIdContractsContractIdItemsGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


