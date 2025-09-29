# CorporationsCorporationIdOrdersGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int64** | Number of days for which order is valid (starting from the issued date). An order expires at time issued + duration | 
**Escrow** | Pointer to **float64** | For buy orders, the amount of ISK in escrow | [optional] 
**IsBuyOrder** | Pointer to **bool** | True if the order is a bid (buy) order | [optional] 
**Issued** | **time.Time** | Date and time when this order was issued | 
**IssuedBy** | **int64** | The character who issued this order | 
**LocationId** | **int64** | ID of the location where order was placed | 
**MinVolume** | Pointer to **int64** | For buy orders, the minimum quantity that will be accepted in a matching sell order | [optional] 
**OrderId** | **int64** | Unique order ID | 
**Price** | **float64** | Cost per unit for this order | 
**Range** | **string** | Valid order range, numbers are ranges in jumps | 
**RegionId** | **int64** | ID of the region where order was placed | 
**TypeId** | **int64** | The type ID of the item transacted in this order | 
**VolumeRemain** | **int64** | Quantity of items still required or offered | 
**VolumeTotal** | **int64** | Quantity of items required or offered at time order was placed | 
**WalletDivision** | **int64** | The corporation wallet division used for this order. | 

## Methods

### NewCorporationsCorporationIdOrdersGetInner

`func NewCorporationsCorporationIdOrdersGetInner(duration int64, issued time.Time, issuedBy int64, locationId int64, orderId int64, price float64, range_ string, regionId int64, typeId int64, volumeRemain int64, volumeTotal int64, walletDivision int64, ) *CorporationsCorporationIdOrdersGetInner`

NewCorporationsCorporationIdOrdersGetInner instantiates a new CorporationsCorporationIdOrdersGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdOrdersGetInnerWithDefaults

`func NewCorporationsCorporationIdOrdersGetInnerWithDefaults() *CorporationsCorporationIdOrdersGetInner`

NewCorporationsCorporationIdOrdersGetInnerWithDefaults instantiates a new CorporationsCorporationIdOrdersGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CorporationsCorporationIdOrdersGetInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CorporationsCorporationIdOrdersGetInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetEscrow

`func (o *CorporationsCorporationIdOrdersGetInner) GetEscrow() float64`

GetEscrow returns the Escrow field if non-nil, zero value otherwise.

### GetEscrowOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetEscrowOk() (*float64, bool)`

GetEscrowOk returns a tuple with the Escrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrow

`func (o *CorporationsCorporationIdOrdersGetInner) SetEscrow(v float64)`

SetEscrow sets Escrow field to given value.

### HasEscrow

`func (o *CorporationsCorporationIdOrdersGetInner) HasEscrow() bool`

HasEscrow returns a boolean if a field has been set.

### GetIsBuyOrder

`func (o *CorporationsCorporationIdOrdersGetInner) GetIsBuyOrder() bool`

GetIsBuyOrder returns the IsBuyOrder field if non-nil, zero value otherwise.

### GetIsBuyOrderOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetIsBuyOrderOk() (*bool, bool)`

GetIsBuyOrderOk returns a tuple with the IsBuyOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyOrder

`func (o *CorporationsCorporationIdOrdersGetInner) SetIsBuyOrder(v bool)`

SetIsBuyOrder sets IsBuyOrder field to given value.

### HasIsBuyOrder

`func (o *CorporationsCorporationIdOrdersGetInner) HasIsBuyOrder() bool`

HasIsBuyOrder returns a boolean if a field has been set.

### GetIssued

`func (o *CorporationsCorporationIdOrdersGetInner) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *CorporationsCorporationIdOrdersGetInner) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetIssuedBy

`func (o *CorporationsCorporationIdOrdersGetInner) GetIssuedBy() int64`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetIssuedByOk() (*int64, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *CorporationsCorporationIdOrdersGetInner) SetIssuedBy(v int64)`

SetIssuedBy sets IssuedBy field to given value.


### GetLocationId

`func (o *CorporationsCorporationIdOrdersGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CorporationsCorporationIdOrdersGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetMinVolume

`func (o *CorporationsCorporationIdOrdersGetInner) GetMinVolume() int64`

GetMinVolume returns the MinVolume field if non-nil, zero value otherwise.

### GetMinVolumeOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetMinVolumeOk() (*int64, bool)`

GetMinVolumeOk returns a tuple with the MinVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVolume

`func (o *CorporationsCorporationIdOrdersGetInner) SetMinVolume(v int64)`

SetMinVolume sets MinVolume field to given value.

### HasMinVolume

`func (o *CorporationsCorporationIdOrdersGetInner) HasMinVolume() bool`

HasMinVolume returns a boolean if a field has been set.

### GetOrderId

`func (o *CorporationsCorporationIdOrdersGetInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CorporationsCorporationIdOrdersGetInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetPrice

`func (o *CorporationsCorporationIdOrdersGetInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CorporationsCorporationIdOrdersGetInner) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetRange

`func (o *CorporationsCorporationIdOrdersGetInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CorporationsCorporationIdOrdersGetInner) SetRange(v string)`

SetRange sets Range field to given value.


### GetRegionId

`func (o *CorporationsCorporationIdOrdersGetInner) GetRegionId() int64`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetRegionIdOk() (*int64, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CorporationsCorporationIdOrdersGetInner) SetRegionId(v int64)`

SetRegionId sets RegionId field to given value.


### GetTypeId

`func (o *CorporationsCorporationIdOrdersGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsCorporationIdOrdersGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetVolumeRemain

`func (o *CorporationsCorporationIdOrdersGetInner) GetVolumeRemain() int64`

GetVolumeRemain returns the VolumeRemain field if non-nil, zero value otherwise.

### GetVolumeRemainOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetVolumeRemainOk() (*int64, bool)`

GetVolumeRemainOk returns a tuple with the VolumeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRemain

`func (o *CorporationsCorporationIdOrdersGetInner) SetVolumeRemain(v int64)`

SetVolumeRemain sets VolumeRemain field to given value.


### GetVolumeTotal

`func (o *CorporationsCorporationIdOrdersGetInner) GetVolumeTotal() int64`

GetVolumeTotal returns the VolumeTotal field if non-nil, zero value otherwise.

### GetVolumeTotalOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetVolumeTotalOk() (*int64, bool)`

GetVolumeTotalOk returns a tuple with the VolumeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTotal

`func (o *CorporationsCorporationIdOrdersGetInner) SetVolumeTotal(v int64)`

SetVolumeTotal sets VolumeTotal field to given value.


### GetWalletDivision

`func (o *CorporationsCorporationIdOrdersGetInner) GetWalletDivision() int64`

GetWalletDivision returns the WalletDivision field if non-nil, zero value otherwise.

### GetWalletDivisionOk

`func (o *CorporationsCorporationIdOrdersGetInner) GetWalletDivisionOk() (*int64, bool)`

GetWalletDivisionOk returns a tuple with the WalletDivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletDivision

`func (o *CorporationsCorporationIdOrdersGetInner) SetWalletDivision(v int64)`

SetWalletDivision sets WalletDivision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


