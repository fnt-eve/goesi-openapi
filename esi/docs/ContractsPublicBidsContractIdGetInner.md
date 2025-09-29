# ContractsPublicBidsContractIdGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount bid, in ISK | 
**BidId** | **int64** | Unique ID for the bid | 
**DateBid** | **time.Time** | Datetime when the bid was placed | 

## Methods

### NewContractsPublicBidsContractIdGetInner

`func NewContractsPublicBidsContractIdGetInner(amount float64, bidId int64, dateBid time.Time, ) *ContractsPublicBidsContractIdGetInner`

NewContractsPublicBidsContractIdGetInner instantiates a new ContractsPublicBidsContractIdGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsPublicBidsContractIdGetInnerWithDefaults

`func NewContractsPublicBidsContractIdGetInnerWithDefaults() *ContractsPublicBidsContractIdGetInner`

NewContractsPublicBidsContractIdGetInnerWithDefaults instantiates a new ContractsPublicBidsContractIdGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ContractsPublicBidsContractIdGetInner) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ContractsPublicBidsContractIdGetInner) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ContractsPublicBidsContractIdGetInner) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetBidId

`func (o *ContractsPublicBidsContractIdGetInner) GetBidId() int64`

GetBidId returns the BidId field if non-nil, zero value otherwise.

### GetBidIdOk

`func (o *ContractsPublicBidsContractIdGetInner) GetBidIdOk() (*int64, bool)`

GetBidIdOk returns a tuple with the BidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidId

`func (o *ContractsPublicBidsContractIdGetInner) SetBidId(v int64)`

SetBidId sets BidId field to given value.


### GetDateBid

`func (o *ContractsPublicBidsContractIdGetInner) GetDateBid() time.Time`

GetDateBid returns the DateBid field if non-nil, zero value otherwise.

### GetDateBidOk

`func (o *ContractsPublicBidsContractIdGetInner) GetDateBidOk() (*time.Time, bool)`

GetDateBidOk returns a tuple with the DateBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBid

`func (o *ContractsPublicBidsContractIdGetInner) SetDateBid(v time.Time)`

SetDateBid sets DateBid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


