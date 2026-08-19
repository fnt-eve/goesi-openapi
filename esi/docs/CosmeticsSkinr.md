# CosmeticsSkinr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorId** | **int64** |  | 
**Id** | **string** | SKINR ID | 
**Layout** | [**CosmeticsSkinrLayout**](CosmeticsSkinrLayout.md) |  | 
**Line** | Pointer to **string** | SKINR line name | [optional] 
**Name** | **string** | SKINR name | 
**ShipTypeId** | **int64** |  | 
**Tier** | [**CosmeticsSkinrTier**](CosmeticsSkinrTier.md) |  | 

## Methods

### NewCosmeticsSkinr

`func NewCosmeticsSkinr(creatorId int64, id string, layout CosmeticsSkinrLayout, name string, shipTypeId int64, tier CosmeticsSkinrTier, ) *CosmeticsSkinr`

NewCosmeticsSkinr instantiates a new CosmeticsSkinr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmeticsSkinrWithDefaults

`func NewCosmeticsSkinrWithDefaults() *CosmeticsSkinr`

NewCosmeticsSkinrWithDefaults instantiates a new CosmeticsSkinr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatorId

`func (o *CosmeticsSkinr) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CosmeticsSkinr) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CosmeticsSkinr) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.


### GetId

`func (o *CosmeticsSkinr) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CosmeticsSkinr) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CosmeticsSkinr) SetId(v string)`

SetId sets Id field to given value.


### GetLayout

`func (o *CosmeticsSkinr) GetLayout() CosmeticsSkinrLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *CosmeticsSkinr) GetLayoutOk() (*CosmeticsSkinrLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *CosmeticsSkinr) SetLayout(v CosmeticsSkinrLayout)`

SetLayout sets Layout field to given value.


### GetLine

`func (o *CosmeticsSkinr) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CosmeticsSkinr) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CosmeticsSkinr) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *CosmeticsSkinr) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetName

`func (o *CosmeticsSkinr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CosmeticsSkinr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CosmeticsSkinr) SetName(v string)`

SetName sets Name field to given value.


### GetShipTypeId

`func (o *CosmeticsSkinr) GetShipTypeId() int64`

GetShipTypeId returns the ShipTypeId field if non-nil, zero value otherwise.

### GetShipTypeIdOk

`func (o *CosmeticsSkinr) GetShipTypeIdOk() (*int64, bool)`

GetShipTypeIdOk returns a tuple with the ShipTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTypeId

`func (o *CosmeticsSkinr) SetShipTypeId(v int64)`

SetShipTypeId sets ShipTypeId field to given value.


### GetTier

`func (o *CosmeticsSkinr) GetTier() CosmeticsSkinrTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CosmeticsSkinr) GetTierOk() (*CosmeticsSkinrTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CosmeticsSkinr) SetTier(v CosmeticsSkinrTier)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


