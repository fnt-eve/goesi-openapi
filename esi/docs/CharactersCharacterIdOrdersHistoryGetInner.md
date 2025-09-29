# CharactersCharacterIdOrdersHistoryGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int64** | Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration | 
**Escrow** | Pointer to **float64** | For buy orders, the amount of ISK in escrow | [optional] 
**IsBuyOrder** | Pointer to **bool** | True if the order is a bid (buy) order | [optional] 
**IsCorporation** | **bool** | Signifies whether the buy/sell order was placed on behalf of a corporation. | 
**Issued** | **time.Time** | Date and time when this order was issued | 
**LocationId** | **int64** | ID of the location where order was placed | 
**MinVolume** | Pointer to **int64** | For buy orders, the minimum quantity that will be accepted in a matching sell order | [optional] 
**OrderId** | **int64** | Unique order ID | 
**Price** | **float64** | Cost per unit for this order | 
**Range** | **string** | Valid order range, numbers are ranges in jumps | 
**RegionId** | **int64** | ID of the region where order was placed | 
**State** | **string** | Current order state | 
**TypeId** | **int64** | The type ID of the item transacted in this order | 
**VolumeRemain** | **int64** | Quantity of items still required or offered | 
**VolumeTotal** | **int64** | Quantity of items required or offered at time order was placed | 

## Methods

### NewCharactersCharacterIdOrdersHistoryGetInner

`func NewCharactersCharacterIdOrdersHistoryGetInner(duration int64, isCorporation bool, issued time.Time, locationId int64, orderId int64, price float64, range_ string, regionId int64, state string, typeId int64, volumeRemain int64, volumeTotal int64, ) *CharactersCharacterIdOrdersHistoryGetInner`

NewCharactersCharacterIdOrdersHistoryGetInner instantiates a new CharactersCharacterIdOrdersHistoryGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdOrdersHistoryGetInnerWithDefaults

`func NewCharactersCharacterIdOrdersHistoryGetInnerWithDefaults() *CharactersCharacterIdOrdersHistoryGetInner`

NewCharactersCharacterIdOrdersHistoryGetInnerWithDefaults instantiates a new CharactersCharacterIdOrdersHistoryGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetEscrow

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetEscrow() float64`

GetEscrow returns the Escrow field if non-nil, zero value otherwise.

### GetEscrowOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetEscrowOk() (*float64, bool)`

GetEscrowOk returns a tuple with the Escrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrow

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetEscrow(v float64)`

SetEscrow sets Escrow field to given value.

### HasEscrow

`func (o *CharactersCharacterIdOrdersHistoryGetInner) HasEscrow() bool`

HasEscrow returns a boolean if a field has been set.

### GetIsBuyOrder

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetIsBuyOrder() bool`

GetIsBuyOrder returns the IsBuyOrder field if non-nil, zero value otherwise.

### GetIsBuyOrderOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetIsBuyOrderOk() (*bool, bool)`

GetIsBuyOrderOk returns a tuple with the IsBuyOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyOrder

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetIsBuyOrder(v bool)`

SetIsBuyOrder sets IsBuyOrder field to given value.

### HasIsBuyOrder

`func (o *CharactersCharacterIdOrdersHistoryGetInner) HasIsBuyOrder() bool`

HasIsBuyOrder returns a boolean if a field has been set.

### GetIsCorporation

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetIsCorporation() bool`

GetIsCorporation returns the IsCorporation field if non-nil, zero value otherwise.

### GetIsCorporationOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetIsCorporationOk() (*bool, bool)`

GetIsCorporationOk returns a tuple with the IsCorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCorporation

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetIsCorporation(v bool)`

SetIsCorporation sets IsCorporation field to given value.


### GetIssued

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetLocationId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetMinVolume

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetMinVolume() int64`

GetMinVolume returns the MinVolume field if non-nil, zero value otherwise.

### GetMinVolumeOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetMinVolumeOk() (*int64, bool)`

GetMinVolumeOk returns a tuple with the MinVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVolume

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetMinVolume(v int64)`

SetMinVolume sets MinVolume field to given value.

### HasMinVolume

`func (o *CharactersCharacterIdOrdersHistoryGetInner) HasMinVolume() bool`

HasMinVolume returns a boolean if a field has been set.

### GetOrderId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetPrice

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetRange

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetRange(v string)`

SetRange sets Range field to given value.


### GetRegionId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetRegionId() int64`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetRegionIdOk() (*int64, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetRegionId(v int64)`

SetRegionId sets RegionId field to given value.


### GetState

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetState(v string)`

SetState sets State field to given value.


### GetTypeId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetVolumeRemain

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetVolumeRemain() int64`

GetVolumeRemain returns the VolumeRemain field if non-nil, zero value otherwise.

### GetVolumeRemainOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetVolumeRemainOk() (*int64, bool)`

GetVolumeRemainOk returns a tuple with the VolumeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRemain

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetVolumeRemain(v int64)`

SetVolumeRemain sets VolumeRemain field to given value.


### GetVolumeTotal

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetVolumeTotal() int64`

GetVolumeTotal returns the VolumeTotal field if non-nil, zero value otherwise.

### GetVolumeTotalOk

`func (o *CharactersCharacterIdOrdersHistoryGetInner) GetVolumeTotalOk() (*int64, bool)`

GetVolumeTotalOk returns a tuple with the VolumeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTotal

`func (o *CharactersCharacterIdOrdersHistoryGetInner) SetVolumeTotal(v int64)`

SetVolumeTotal sets VolumeTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


