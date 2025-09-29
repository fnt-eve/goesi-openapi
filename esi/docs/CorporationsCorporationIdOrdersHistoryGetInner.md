# CorporationsCorporationIdOrdersHistoryGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int64** | Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration | 
**Escrow** | Pointer to **float64** | For buy orders, the amount of ISK in escrow | [optional] 
**IsBuyOrder** | Pointer to **bool** | True if the order is a bid (buy) order | [optional] 
**Issued** | **time.Time** | Date and time when this order was issued | 
**IssuedBy** | Pointer to **int64** | The character who issued this order | [optional] 
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
**WalletDivision** | **int64** | The corporation wallet division used for this order | 

## Methods

### NewCorporationsCorporationIdOrdersHistoryGetInner

`func NewCorporationsCorporationIdOrdersHistoryGetInner(duration int64, issued time.Time, locationId int64, orderId int64, price float64, range_ string, regionId int64, state string, typeId int64, volumeRemain int64, volumeTotal int64, walletDivision int64, ) *CorporationsCorporationIdOrdersHistoryGetInner`

NewCorporationsCorporationIdOrdersHistoryGetInner instantiates a new CorporationsCorporationIdOrdersHistoryGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdOrdersHistoryGetInnerWithDefaults

`func NewCorporationsCorporationIdOrdersHistoryGetInnerWithDefaults() *CorporationsCorporationIdOrdersHistoryGetInner`

NewCorporationsCorporationIdOrdersHistoryGetInnerWithDefaults instantiates a new CorporationsCorporationIdOrdersHistoryGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetEscrow

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetEscrow() float64`

GetEscrow returns the Escrow field if non-nil, zero value otherwise.

### GetEscrowOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetEscrowOk() (*float64, bool)`

GetEscrowOk returns a tuple with the Escrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrow

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetEscrow(v float64)`

SetEscrow sets Escrow field to given value.

### HasEscrow

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) HasEscrow() bool`

HasEscrow returns a boolean if a field has been set.

### GetIsBuyOrder

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetIsBuyOrder() bool`

GetIsBuyOrder returns the IsBuyOrder field if non-nil, zero value otherwise.

### GetIsBuyOrderOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetIsBuyOrderOk() (*bool, bool)`

GetIsBuyOrderOk returns a tuple with the IsBuyOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyOrder

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetIsBuyOrder(v bool)`

SetIsBuyOrder sets IsBuyOrder field to given value.

### HasIsBuyOrder

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) HasIsBuyOrder() bool`

HasIsBuyOrder returns a boolean if a field has been set.

### GetIssued

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetIssuedBy

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetIssuedBy() int64`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetIssuedByOk() (*int64, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetIssuedBy(v int64)`

SetIssuedBy sets IssuedBy field to given value.

### HasIssuedBy

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) HasIssuedBy() bool`

HasIssuedBy returns a boolean if a field has been set.

### GetLocationId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetMinVolume

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetMinVolume() int64`

GetMinVolume returns the MinVolume field if non-nil, zero value otherwise.

### GetMinVolumeOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetMinVolumeOk() (*int64, bool)`

GetMinVolumeOk returns a tuple with the MinVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVolume

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetMinVolume(v int64)`

SetMinVolume sets MinVolume field to given value.

### HasMinVolume

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) HasMinVolume() bool`

HasMinVolume returns a boolean if a field has been set.

### GetOrderId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetPrice

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetRange

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetRange(v string)`

SetRange sets Range field to given value.


### GetRegionId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetRegionId() int64`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetRegionIdOk() (*int64, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetRegionId(v int64)`

SetRegionId sets RegionId field to given value.


### GetState

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetState(v string)`

SetState sets State field to given value.


### GetTypeId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetVolumeRemain

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetVolumeRemain() int64`

GetVolumeRemain returns the VolumeRemain field if non-nil, zero value otherwise.

### GetVolumeRemainOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetVolumeRemainOk() (*int64, bool)`

GetVolumeRemainOk returns a tuple with the VolumeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRemain

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetVolumeRemain(v int64)`

SetVolumeRemain sets VolumeRemain field to given value.


### GetVolumeTotal

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetVolumeTotal() int64`

GetVolumeTotal returns the VolumeTotal field if non-nil, zero value otherwise.

### GetVolumeTotalOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetVolumeTotalOk() (*int64, bool)`

GetVolumeTotalOk returns a tuple with the VolumeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTotal

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetVolumeTotal(v int64)`

SetVolumeTotal sets VolumeTotal field to given value.


### GetWalletDivision

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetWalletDivision() int64`

GetWalletDivision returns the WalletDivision field if non-nil, zero value otherwise.

### GetWalletDivisionOk

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) GetWalletDivisionOk() (*int64, bool)`

GetWalletDivisionOk returns a tuple with the WalletDivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletDivision

`func (o *CorporationsCorporationIdOrdersHistoryGetInner) SetWalletDivision(v int64)`

SetWalletDivision sets WalletDivision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


