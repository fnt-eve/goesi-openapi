# CorporationsCorporationIdContractsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptorId** | **int64** | Who will accept the contract | 
**AssigneeId** | **int64** | ID to whom the contract is assigned, can be corporation or character ID | 
**Availability** | **string** | To whom the contract is available | 
**Buyout** | Pointer to **float64** | Buyout price (for Auctions only) | [optional] 
**Collateral** | Pointer to **float64** | Collateral price (for Couriers only) | [optional] 
**ContractId** | **int64** |  | 
**DateAccepted** | Pointer to **time.Time** | Date of confirmation of contract | [optional] 
**DateCompleted** | Pointer to **time.Time** | Date of completed of contract | [optional] 
**DateExpired** | **time.Time** | Expiration date of the contract | 
**DateIssued** | **time.Time** | Сreation date of the contract | 
**DaysToComplete** | Pointer to **int64** | Number of days to perform the contract | [optional] 
**EndLocationId** | Pointer to **int64** | End location ID (for Couriers contract) | [optional] 
**ForCorporation** | **bool** | true if the contract was issued on behalf of the issuer&#39;s corporation | 
**IssuerCorporationId** | **int64** | Character&#39;s corporation ID for the issuer | 
**IssuerId** | **int64** | Character ID for the issuer | 
**Price** | Pointer to **float64** | Price of contract (for ItemsExchange and Auctions) | [optional] 
**Reward** | Pointer to **float64** | Remuneration for contract (for Couriers only) | [optional] 
**StartLocationId** | Pointer to **int64** | Start location ID (for Couriers contract) | [optional] 
**Status** | **string** | Status of the the contract | 
**Title** | Pointer to **string** | Title of the contract | [optional] 
**Type** | **string** | Type of the contract | 
**Volume** | Pointer to **float64** | Volume of items in the contract | [optional] 

## Methods

### NewCorporationsCorporationIdContractsGetInner

`func NewCorporationsCorporationIdContractsGetInner(acceptorId int64, assigneeId int64, availability string, contractId int64, dateExpired time.Time, dateIssued time.Time, forCorporation bool, issuerCorporationId int64, issuerId int64, status string, type_ string, ) *CorporationsCorporationIdContractsGetInner`

NewCorporationsCorporationIdContractsGetInner instantiates a new CorporationsCorporationIdContractsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdContractsGetInnerWithDefaults

`func NewCorporationsCorporationIdContractsGetInnerWithDefaults() *CorporationsCorporationIdContractsGetInner`

NewCorporationsCorporationIdContractsGetInnerWithDefaults instantiates a new CorporationsCorporationIdContractsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptorId

`func (o *CorporationsCorporationIdContractsGetInner) GetAcceptorId() int64`

GetAcceptorId returns the AcceptorId field if non-nil, zero value otherwise.

### GetAcceptorIdOk

`func (o *CorporationsCorporationIdContractsGetInner) GetAcceptorIdOk() (*int64, bool)`

GetAcceptorIdOk returns a tuple with the AcceptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptorId

`func (o *CorporationsCorporationIdContractsGetInner) SetAcceptorId(v int64)`

SetAcceptorId sets AcceptorId field to given value.


### GetAssigneeId

`func (o *CorporationsCorporationIdContractsGetInner) GetAssigneeId() int64`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *CorporationsCorporationIdContractsGetInner) GetAssigneeIdOk() (*int64, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *CorporationsCorporationIdContractsGetInner) SetAssigneeId(v int64)`

SetAssigneeId sets AssigneeId field to given value.


### GetAvailability

`func (o *CorporationsCorporationIdContractsGetInner) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CorporationsCorporationIdContractsGetInner) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CorporationsCorporationIdContractsGetInner) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetBuyout

`func (o *CorporationsCorporationIdContractsGetInner) GetBuyout() float64`

GetBuyout returns the Buyout field if non-nil, zero value otherwise.

### GetBuyoutOk

`func (o *CorporationsCorporationIdContractsGetInner) GetBuyoutOk() (*float64, bool)`

GetBuyoutOk returns a tuple with the Buyout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyout

`func (o *CorporationsCorporationIdContractsGetInner) SetBuyout(v float64)`

SetBuyout sets Buyout field to given value.

### HasBuyout

`func (o *CorporationsCorporationIdContractsGetInner) HasBuyout() bool`

HasBuyout returns a boolean if a field has been set.

### GetCollateral

`func (o *CorporationsCorporationIdContractsGetInner) GetCollateral() float64`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *CorporationsCorporationIdContractsGetInner) GetCollateralOk() (*float64, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *CorporationsCorporationIdContractsGetInner) SetCollateral(v float64)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *CorporationsCorporationIdContractsGetInner) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetContractId

`func (o *CorporationsCorporationIdContractsGetInner) GetContractId() int64`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *CorporationsCorporationIdContractsGetInner) GetContractIdOk() (*int64, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *CorporationsCorporationIdContractsGetInner) SetContractId(v int64)`

SetContractId sets ContractId field to given value.


### GetDateAccepted

`func (o *CorporationsCorporationIdContractsGetInner) GetDateAccepted() time.Time`

GetDateAccepted returns the DateAccepted field if non-nil, zero value otherwise.

### GetDateAcceptedOk

`func (o *CorporationsCorporationIdContractsGetInner) GetDateAcceptedOk() (*time.Time, bool)`

GetDateAcceptedOk returns a tuple with the DateAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAccepted

`func (o *CorporationsCorporationIdContractsGetInner) SetDateAccepted(v time.Time)`

SetDateAccepted sets DateAccepted field to given value.

### HasDateAccepted

`func (o *CorporationsCorporationIdContractsGetInner) HasDateAccepted() bool`

HasDateAccepted returns a boolean if a field has been set.

### GetDateCompleted

`func (o *CorporationsCorporationIdContractsGetInner) GetDateCompleted() time.Time`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *CorporationsCorporationIdContractsGetInner) GetDateCompletedOk() (*time.Time, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *CorporationsCorporationIdContractsGetInner) SetDateCompleted(v time.Time)`

SetDateCompleted sets DateCompleted field to given value.

### HasDateCompleted

`func (o *CorporationsCorporationIdContractsGetInner) HasDateCompleted() bool`

HasDateCompleted returns a boolean if a field has been set.

### GetDateExpired

`func (o *CorporationsCorporationIdContractsGetInner) GetDateExpired() time.Time`

GetDateExpired returns the DateExpired field if non-nil, zero value otherwise.

### GetDateExpiredOk

`func (o *CorporationsCorporationIdContractsGetInner) GetDateExpiredOk() (*time.Time, bool)`

GetDateExpiredOk returns a tuple with the DateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpired

`func (o *CorporationsCorporationIdContractsGetInner) SetDateExpired(v time.Time)`

SetDateExpired sets DateExpired field to given value.


### GetDateIssued

`func (o *CorporationsCorporationIdContractsGetInner) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *CorporationsCorporationIdContractsGetInner) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *CorporationsCorporationIdContractsGetInner) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.


### GetDaysToComplete

`func (o *CorporationsCorporationIdContractsGetInner) GetDaysToComplete() int64`

GetDaysToComplete returns the DaysToComplete field if non-nil, zero value otherwise.

### GetDaysToCompleteOk

`func (o *CorporationsCorporationIdContractsGetInner) GetDaysToCompleteOk() (*int64, bool)`

GetDaysToCompleteOk returns a tuple with the DaysToComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToComplete

`func (o *CorporationsCorporationIdContractsGetInner) SetDaysToComplete(v int64)`

SetDaysToComplete sets DaysToComplete field to given value.

### HasDaysToComplete

`func (o *CorporationsCorporationIdContractsGetInner) HasDaysToComplete() bool`

HasDaysToComplete returns a boolean if a field has been set.

### GetEndLocationId

`func (o *CorporationsCorporationIdContractsGetInner) GetEndLocationId() int64`

GetEndLocationId returns the EndLocationId field if non-nil, zero value otherwise.

### GetEndLocationIdOk

`func (o *CorporationsCorporationIdContractsGetInner) GetEndLocationIdOk() (*int64, bool)`

GetEndLocationIdOk returns a tuple with the EndLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocationId

`func (o *CorporationsCorporationIdContractsGetInner) SetEndLocationId(v int64)`

SetEndLocationId sets EndLocationId field to given value.

### HasEndLocationId

`func (o *CorporationsCorporationIdContractsGetInner) HasEndLocationId() bool`

HasEndLocationId returns a boolean if a field has been set.

### GetForCorporation

`func (o *CorporationsCorporationIdContractsGetInner) GetForCorporation() bool`

GetForCorporation returns the ForCorporation field if non-nil, zero value otherwise.

### GetForCorporationOk

`func (o *CorporationsCorporationIdContractsGetInner) GetForCorporationOk() (*bool, bool)`

GetForCorporationOk returns a tuple with the ForCorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForCorporation

`func (o *CorporationsCorporationIdContractsGetInner) SetForCorporation(v bool)`

SetForCorporation sets ForCorporation field to given value.


### GetIssuerCorporationId

`func (o *CorporationsCorporationIdContractsGetInner) GetIssuerCorporationId() int64`

GetIssuerCorporationId returns the IssuerCorporationId field if non-nil, zero value otherwise.

### GetIssuerCorporationIdOk

`func (o *CorporationsCorporationIdContractsGetInner) GetIssuerCorporationIdOk() (*int64, bool)`

GetIssuerCorporationIdOk returns a tuple with the IssuerCorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCorporationId

`func (o *CorporationsCorporationIdContractsGetInner) SetIssuerCorporationId(v int64)`

SetIssuerCorporationId sets IssuerCorporationId field to given value.


### GetIssuerId

`func (o *CorporationsCorporationIdContractsGetInner) GetIssuerId() int64`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *CorporationsCorporationIdContractsGetInner) GetIssuerIdOk() (*int64, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *CorporationsCorporationIdContractsGetInner) SetIssuerId(v int64)`

SetIssuerId sets IssuerId field to given value.


### GetPrice

`func (o *CorporationsCorporationIdContractsGetInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CorporationsCorporationIdContractsGetInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CorporationsCorporationIdContractsGetInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CorporationsCorporationIdContractsGetInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReward

`func (o *CorporationsCorporationIdContractsGetInner) GetReward() float64`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *CorporationsCorporationIdContractsGetInner) GetRewardOk() (*float64, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *CorporationsCorporationIdContractsGetInner) SetReward(v float64)`

SetReward sets Reward field to given value.

### HasReward

`func (o *CorporationsCorporationIdContractsGetInner) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetStartLocationId

`func (o *CorporationsCorporationIdContractsGetInner) GetStartLocationId() int64`

GetStartLocationId returns the StartLocationId field if non-nil, zero value otherwise.

### GetStartLocationIdOk

`func (o *CorporationsCorporationIdContractsGetInner) GetStartLocationIdOk() (*int64, bool)`

GetStartLocationIdOk returns a tuple with the StartLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocationId

`func (o *CorporationsCorporationIdContractsGetInner) SetStartLocationId(v int64)`

SetStartLocationId sets StartLocationId field to given value.

### HasStartLocationId

`func (o *CorporationsCorporationIdContractsGetInner) HasStartLocationId() bool`

HasStartLocationId returns a boolean if a field has been set.

### GetStatus

`func (o *CorporationsCorporationIdContractsGetInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CorporationsCorporationIdContractsGetInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CorporationsCorporationIdContractsGetInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *CorporationsCorporationIdContractsGetInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CorporationsCorporationIdContractsGetInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CorporationsCorporationIdContractsGetInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CorporationsCorporationIdContractsGetInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CorporationsCorporationIdContractsGetInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CorporationsCorporationIdContractsGetInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CorporationsCorporationIdContractsGetInner) SetType(v string)`

SetType sets Type field to given value.


### GetVolume

`func (o *CorporationsCorporationIdContractsGetInner) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CorporationsCorporationIdContractsGetInner) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CorporationsCorporationIdContractsGetInner) SetVolume(v float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CorporationsCorporationIdContractsGetInner) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


