# ContractsPublicRegionIdGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buyout** | Pointer to **float64** | Buyout price (for Auctions only) | [optional] 
**Collateral** | Pointer to **float64** | Collateral price (for Couriers only) | [optional] 
**ContractId** | **int64** |  | 
**DateExpired** | **time.Time** | Expiration date of the contract | 
**DateIssued** | **time.Time** | Сreation date of the contract | 
**DaysToComplete** | Pointer to **int64** | Number of days to perform the contract | [optional] 
**EndLocationId** | Pointer to **int64** | End location ID (for Couriers contract) | [optional] 
**ForCorporation** | Pointer to **bool** | true if the contract was issued on behalf of the issuer&#39;s corporation | [optional] 
**IssuerCorporationId** | **int64** | Character&#39;s corporation ID for the issuer | 
**IssuerId** | **int64** | Character ID for the issuer | 
**Price** | Pointer to **float64** | Price of contract (for ItemsExchange and Auctions) | [optional] 
**Reward** | Pointer to **float64** | Remuneration for contract (for Couriers only) | [optional] 
**StartLocationId** | Pointer to **int64** | Start location ID (for Couriers contract) | [optional] 
**Title** | Pointer to **string** | Title of the contract | [optional] 
**Type** | **string** | Type of the contract | 
**Volume** | Pointer to **float64** | Volume of items in the contract | [optional] 

## Methods

### NewContractsPublicRegionIdGetInner

`func NewContractsPublicRegionIdGetInner(contractId int64, dateExpired time.Time, dateIssued time.Time, issuerCorporationId int64, issuerId int64, type_ string, ) *ContractsPublicRegionIdGetInner`

NewContractsPublicRegionIdGetInner instantiates a new ContractsPublicRegionIdGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsPublicRegionIdGetInnerWithDefaults

`func NewContractsPublicRegionIdGetInnerWithDefaults() *ContractsPublicRegionIdGetInner`

NewContractsPublicRegionIdGetInnerWithDefaults instantiates a new ContractsPublicRegionIdGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyout

`func (o *ContractsPublicRegionIdGetInner) GetBuyout() float64`

GetBuyout returns the Buyout field if non-nil, zero value otherwise.

### GetBuyoutOk

`func (o *ContractsPublicRegionIdGetInner) GetBuyoutOk() (*float64, bool)`

GetBuyoutOk returns a tuple with the Buyout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyout

`func (o *ContractsPublicRegionIdGetInner) SetBuyout(v float64)`

SetBuyout sets Buyout field to given value.

### HasBuyout

`func (o *ContractsPublicRegionIdGetInner) HasBuyout() bool`

HasBuyout returns a boolean if a field has been set.

### GetCollateral

`func (o *ContractsPublicRegionIdGetInner) GetCollateral() float64`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *ContractsPublicRegionIdGetInner) GetCollateralOk() (*float64, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *ContractsPublicRegionIdGetInner) SetCollateral(v float64)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *ContractsPublicRegionIdGetInner) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetContractId

`func (o *ContractsPublicRegionIdGetInner) GetContractId() int64`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractsPublicRegionIdGetInner) GetContractIdOk() (*int64, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractsPublicRegionIdGetInner) SetContractId(v int64)`

SetContractId sets ContractId field to given value.


### GetDateExpired

`func (o *ContractsPublicRegionIdGetInner) GetDateExpired() time.Time`

GetDateExpired returns the DateExpired field if non-nil, zero value otherwise.

### GetDateExpiredOk

`func (o *ContractsPublicRegionIdGetInner) GetDateExpiredOk() (*time.Time, bool)`

GetDateExpiredOk returns a tuple with the DateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpired

`func (o *ContractsPublicRegionIdGetInner) SetDateExpired(v time.Time)`

SetDateExpired sets DateExpired field to given value.


### GetDateIssued

`func (o *ContractsPublicRegionIdGetInner) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *ContractsPublicRegionIdGetInner) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *ContractsPublicRegionIdGetInner) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.


### GetDaysToComplete

`func (o *ContractsPublicRegionIdGetInner) GetDaysToComplete() int64`

GetDaysToComplete returns the DaysToComplete field if non-nil, zero value otherwise.

### GetDaysToCompleteOk

`func (o *ContractsPublicRegionIdGetInner) GetDaysToCompleteOk() (*int64, bool)`

GetDaysToCompleteOk returns a tuple with the DaysToComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToComplete

`func (o *ContractsPublicRegionIdGetInner) SetDaysToComplete(v int64)`

SetDaysToComplete sets DaysToComplete field to given value.

### HasDaysToComplete

`func (o *ContractsPublicRegionIdGetInner) HasDaysToComplete() bool`

HasDaysToComplete returns a boolean if a field has been set.

### GetEndLocationId

`func (o *ContractsPublicRegionIdGetInner) GetEndLocationId() int64`

GetEndLocationId returns the EndLocationId field if non-nil, zero value otherwise.

### GetEndLocationIdOk

`func (o *ContractsPublicRegionIdGetInner) GetEndLocationIdOk() (*int64, bool)`

GetEndLocationIdOk returns a tuple with the EndLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocationId

`func (o *ContractsPublicRegionIdGetInner) SetEndLocationId(v int64)`

SetEndLocationId sets EndLocationId field to given value.

### HasEndLocationId

`func (o *ContractsPublicRegionIdGetInner) HasEndLocationId() bool`

HasEndLocationId returns a boolean if a field has been set.

### GetForCorporation

`func (o *ContractsPublicRegionIdGetInner) GetForCorporation() bool`

GetForCorporation returns the ForCorporation field if non-nil, zero value otherwise.

### GetForCorporationOk

`func (o *ContractsPublicRegionIdGetInner) GetForCorporationOk() (*bool, bool)`

GetForCorporationOk returns a tuple with the ForCorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForCorporation

`func (o *ContractsPublicRegionIdGetInner) SetForCorporation(v bool)`

SetForCorporation sets ForCorporation field to given value.

### HasForCorporation

`func (o *ContractsPublicRegionIdGetInner) HasForCorporation() bool`

HasForCorporation returns a boolean if a field has been set.

### GetIssuerCorporationId

`func (o *ContractsPublicRegionIdGetInner) GetIssuerCorporationId() int64`

GetIssuerCorporationId returns the IssuerCorporationId field if non-nil, zero value otherwise.

### GetIssuerCorporationIdOk

`func (o *ContractsPublicRegionIdGetInner) GetIssuerCorporationIdOk() (*int64, bool)`

GetIssuerCorporationIdOk returns a tuple with the IssuerCorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCorporationId

`func (o *ContractsPublicRegionIdGetInner) SetIssuerCorporationId(v int64)`

SetIssuerCorporationId sets IssuerCorporationId field to given value.


### GetIssuerId

`func (o *ContractsPublicRegionIdGetInner) GetIssuerId() int64`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *ContractsPublicRegionIdGetInner) GetIssuerIdOk() (*int64, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *ContractsPublicRegionIdGetInner) SetIssuerId(v int64)`

SetIssuerId sets IssuerId field to given value.


### GetPrice

`func (o *ContractsPublicRegionIdGetInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ContractsPublicRegionIdGetInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ContractsPublicRegionIdGetInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ContractsPublicRegionIdGetInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReward

`func (o *ContractsPublicRegionIdGetInner) GetReward() float64`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *ContractsPublicRegionIdGetInner) GetRewardOk() (*float64, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *ContractsPublicRegionIdGetInner) SetReward(v float64)`

SetReward sets Reward field to given value.

### HasReward

`func (o *ContractsPublicRegionIdGetInner) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetStartLocationId

`func (o *ContractsPublicRegionIdGetInner) GetStartLocationId() int64`

GetStartLocationId returns the StartLocationId field if non-nil, zero value otherwise.

### GetStartLocationIdOk

`func (o *ContractsPublicRegionIdGetInner) GetStartLocationIdOk() (*int64, bool)`

GetStartLocationIdOk returns a tuple with the StartLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocationId

`func (o *ContractsPublicRegionIdGetInner) SetStartLocationId(v int64)`

SetStartLocationId sets StartLocationId field to given value.

### HasStartLocationId

`func (o *ContractsPublicRegionIdGetInner) HasStartLocationId() bool`

HasStartLocationId returns a boolean if a field has been set.

### GetTitle

`func (o *ContractsPublicRegionIdGetInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContractsPublicRegionIdGetInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContractsPublicRegionIdGetInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContractsPublicRegionIdGetInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ContractsPublicRegionIdGetInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractsPublicRegionIdGetInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractsPublicRegionIdGetInner) SetType(v string)`

SetType sets Type field to given value.


### GetVolume

`func (o *ContractsPublicRegionIdGetInner) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ContractsPublicRegionIdGetInner) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ContractsPublicRegionIdGetInner) SetVolume(v float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ContractsPublicRegionIdGetInner) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


