# MarketsRegionIdOrdersGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int64** |  | 
**IsBuyOrder** | **bool** |  | 
**Issued** | **time.Time** |  | 
**LocationId** | **int64** |  | 
**MinVolume** | **int64** |  | 
**OrderId** | **int64** |  | 
**Price** | **float64** |  | 
**Range** | **string** |  | 
**SystemId** | **int64** | The solar system this order was placed | 
**TypeId** | **int64** |  | 
**VolumeRemain** | **int64** |  | 
**VolumeTotal** | **int64** |  | 

## Methods

### NewMarketsRegionIdOrdersGetInner

`func NewMarketsRegionIdOrdersGetInner(duration int64, isBuyOrder bool, issued time.Time, locationId int64, minVolume int64, orderId int64, price float64, range_ string, systemId int64, typeId int64, volumeRemain int64, volumeTotal int64, ) *MarketsRegionIdOrdersGetInner`

NewMarketsRegionIdOrdersGetInner instantiates a new MarketsRegionIdOrdersGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketsRegionIdOrdersGetInnerWithDefaults

`func NewMarketsRegionIdOrdersGetInnerWithDefaults() *MarketsRegionIdOrdersGetInner`

NewMarketsRegionIdOrdersGetInnerWithDefaults instantiates a new MarketsRegionIdOrdersGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *MarketsRegionIdOrdersGetInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MarketsRegionIdOrdersGetInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MarketsRegionIdOrdersGetInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetIsBuyOrder

`func (o *MarketsRegionIdOrdersGetInner) GetIsBuyOrder() bool`

GetIsBuyOrder returns the IsBuyOrder field if non-nil, zero value otherwise.

### GetIsBuyOrderOk

`func (o *MarketsRegionIdOrdersGetInner) GetIsBuyOrderOk() (*bool, bool)`

GetIsBuyOrderOk returns a tuple with the IsBuyOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyOrder

`func (o *MarketsRegionIdOrdersGetInner) SetIsBuyOrder(v bool)`

SetIsBuyOrder sets IsBuyOrder field to given value.


### GetIssued

`func (o *MarketsRegionIdOrdersGetInner) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *MarketsRegionIdOrdersGetInner) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *MarketsRegionIdOrdersGetInner) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetLocationId

`func (o *MarketsRegionIdOrdersGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *MarketsRegionIdOrdersGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *MarketsRegionIdOrdersGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetMinVolume

`func (o *MarketsRegionIdOrdersGetInner) GetMinVolume() int64`

GetMinVolume returns the MinVolume field if non-nil, zero value otherwise.

### GetMinVolumeOk

`func (o *MarketsRegionIdOrdersGetInner) GetMinVolumeOk() (*int64, bool)`

GetMinVolumeOk returns a tuple with the MinVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVolume

`func (o *MarketsRegionIdOrdersGetInner) SetMinVolume(v int64)`

SetMinVolume sets MinVolume field to given value.


### GetOrderId

`func (o *MarketsRegionIdOrdersGetInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarketsRegionIdOrdersGetInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarketsRegionIdOrdersGetInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetPrice

`func (o *MarketsRegionIdOrdersGetInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketsRegionIdOrdersGetInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketsRegionIdOrdersGetInner) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetRange

`func (o *MarketsRegionIdOrdersGetInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *MarketsRegionIdOrdersGetInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *MarketsRegionIdOrdersGetInner) SetRange(v string)`

SetRange sets Range field to given value.


### GetSystemId

`func (o *MarketsRegionIdOrdersGetInner) GetSystemId() int64`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *MarketsRegionIdOrdersGetInner) GetSystemIdOk() (*int64, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *MarketsRegionIdOrdersGetInner) SetSystemId(v int64)`

SetSystemId sets SystemId field to given value.


### GetTypeId

`func (o *MarketsRegionIdOrdersGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *MarketsRegionIdOrdersGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *MarketsRegionIdOrdersGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetVolumeRemain

`func (o *MarketsRegionIdOrdersGetInner) GetVolumeRemain() int64`

GetVolumeRemain returns the VolumeRemain field if non-nil, zero value otherwise.

### GetVolumeRemainOk

`func (o *MarketsRegionIdOrdersGetInner) GetVolumeRemainOk() (*int64, bool)`

GetVolumeRemainOk returns a tuple with the VolumeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRemain

`func (o *MarketsRegionIdOrdersGetInner) SetVolumeRemain(v int64)`

SetVolumeRemain sets VolumeRemain field to given value.


### GetVolumeTotal

`func (o *MarketsRegionIdOrdersGetInner) GetVolumeTotal() int64`

GetVolumeTotal returns the VolumeTotal field if non-nil, zero value otherwise.

### GetVolumeTotalOk

`func (o *MarketsRegionIdOrdersGetInner) GetVolumeTotalOk() (*int64, bool)`

GetVolumeTotalOk returns a tuple with the VolumeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTotal

`func (o *MarketsRegionIdOrdersGetInner) SetVolumeTotal(v int64)`

SetVolumeTotal sets VolumeTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


