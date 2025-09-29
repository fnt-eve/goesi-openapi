# CharactersCharacterIdContractsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptorId** | **int64** | Who will accept the contract | 
**AssigneeId** | **int64** | ID to whom the contract is assigned, can be alliance, corporation or character ID | 
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

### NewCharactersCharacterIdContractsGetInner

`func NewCharactersCharacterIdContractsGetInner(acceptorId int64, assigneeId int64, availability string, contractId int64, dateExpired time.Time, dateIssued time.Time, forCorporation bool, issuerCorporationId int64, issuerId int64, status string, type_ string, ) *CharactersCharacterIdContractsGetInner`

NewCharactersCharacterIdContractsGetInner instantiates a new CharactersCharacterIdContractsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdContractsGetInnerWithDefaults

`func NewCharactersCharacterIdContractsGetInnerWithDefaults() *CharactersCharacterIdContractsGetInner`

NewCharactersCharacterIdContractsGetInnerWithDefaults instantiates a new CharactersCharacterIdContractsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptorId

`func (o *CharactersCharacterIdContractsGetInner) GetAcceptorId() int64`

GetAcceptorId returns the AcceptorId field if non-nil, zero value otherwise.

### GetAcceptorIdOk

`func (o *CharactersCharacterIdContractsGetInner) GetAcceptorIdOk() (*int64, bool)`

GetAcceptorIdOk returns a tuple with the AcceptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptorId

`func (o *CharactersCharacterIdContractsGetInner) SetAcceptorId(v int64)`

SetAcceptorId sets AcceptorId field to given value.


### GetAssigneeId

`func (o *CharactersCharacterIdContractsGetInner) GetAssigneeId() int64`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *CharactersCharacterIdContractsGetInner) GetAssigneeIdOk() (*int64, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *CharactersCharacterIdContractsGetInner) SetAssigneeId(v int64)`

SetAssigneeId sets AssigneeId field to given value.


### GetAvailability

`func (o *CharactersCharacterIdContractsGetInner) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CharactersCharacterIdContractsGetInner) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CharactersCharacterIdContractsGetInner) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetBuyout

`func (o *CharactersCharacterIdContractsGetInner) GetBuyout() float64`

GetBuyout returns the Buyout field if non-nil, zero value otherwise.

### GetBuyoutOk

`func (o *CharactersCharacterIdContractsGetInner) GetBuyoutOk() (*float64, bool)`

GetBuyoutOk returns a tuple with the Buyout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyout

`func (o *CharactersCharacterIdContractsGetInner) SetBuyout(v float64)`

SetBuyout sets Buyout field to given value.

### HasBuyout

`func (o *CharactersCharacterIdContractsGetInner) HasBuyout() bool`

HasBuyout returns a boolean if a field has been set.

### GetCollateral

`func (o *CharactersCharacterIdContractsGetInner) GetCollateral() float64`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *CharactersCharacterIdContractsGetInner) GetCollateralOk() (*float64, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *CharactersCharacterIdContractsGetInner) SetCollateral(v float64)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *CharactersCharacterIdContractsGetInner) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetContractId

`func (o *CharactersCharacterIdContractsGetInner) GetContractId() int64`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *CharactersCharacterIdContractsGetInner) GetContractIdOk() (*int64, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *CharactersCharacterIdContractsGetInner) SetContractId(v int64)`

SetContractId sets ContractId field to given value.


### GetDateAccepted

`func (o *CharactersCharacterIdContractsGetInner) GetDateAccepted() time.Time`

GetDateAccepted returns the DateAccepted field if non-nil, zero value otherwise.

### GetDateAcceptedOk

`func (o *CharactersCharacterIdContractsGetInner) GetDateAcceptedOk() (*time.Time, bool)`

GetDateAcceptedOk returns a tuple with the DateAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAccepted

`func (o *CharactersCharacterIdContractsGetInner) SetDateAccepted(v time.Time)`

SetDateAccepted sets DateAccepted field to given value.

### HasDateAccepted

`func (o *CharactersCharacterIdContractsGetInner) HasDateAccepted() bool`

HasDateAccepted returns a boolean if a field has been set.

### GetDateCompleted

`func (o *CharactersCharacterIdContractsGetInner) GetDateCompleted() time.Time`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *CharactersCharacterIdContractsGetInner) GetDateCompletedOk() (*time.Time, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *CharactersCharacterIdContractsGetInner) SetDateCompleted(v time.Time)`

SetDateCompleted sets DateCompleted field to given value.

### HasDateCompleted

`func (o *CharactersCharacterIdContractsGetInner) HasDateCompleted() bool`

HasDateCompleted returns a boolean if a field has been set.

### GetDateExpired

`func (o *CharactersCharacterIdContractsGetInner) GetDateExpired() time.Time`

GetDateExpired returns the DateExpired field if non-nil, zero value otherwise.

### GetDateExpiredOk

`func (o *CharactersCharacterIdContractsGetInner) GetDateExpiredOk() (*time.Time, bool)`

GetDateExpiredOk returns a tuple with the DateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpired

`func (o *CharactersCharacterIdContractsGetInner) SetDateExpired(v time.Time)`

SetDateExpired sets DateExpired field to given value.


### GetDateIssued

`func (o *CharactersCharacterIdContractsGetInner) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *CharactersCharacterIdContractsGetInner) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *CharactersCharacterIdContractsGetInner) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.


### GetDaysToComplete

`func (o *CharactersCharacterIdContractsGetInner) GetDaysToComplete() int64`

GetDaysToComplete returns the DaysToComplete field if non-nil, zero value otherwise.

### GetDaysToCompleteOk

`func (o *CharactersCharacterIdContractsGetInner) GetDaysToCompleteOk() (*int64, bool)`

GetDaysToCompleteOk returns a tuple with the DaysToComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToComplete

`func (o *CharactersCharacterIdContractsGetInner) SetDaysToComplete(v int64)`

SetDaysToComplete sets DaysToComplete field to given value.

### HasDaysToComplete

`func (o *CharactersCharacterIdContractsGetInner) HasDaysToComplete() bool`

HasDaysToComplete returns a boolean if a field has been set.

### GetEndLocationId

`func (o *CharactersCharacterIdContractsGetInner) GetEndLocationId() int64`

GetEndLocationId returns the EndLocationId field if non-nil, zero value otherwise.

### GetEndLocationIdOk

`func (o *CharactersCharacterIdContractsGetInner) GetEndLocationIdOk() (*int64, bool)`

GetEndLocationIdOk returns a tuple with the EndLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocationId

`func (o *CharactersCharacterIdContractsGetInner) SetEndLocationId(v int64)`

SetEndLocationId sets EndLocationId field to given value.

### HasEndLocationId

`func (o *CharactersCharacterIdContractsGetInner) HasEndLocationId() bool`

HasEndLocationId returns a boolean if a field has been set.

### GetForCorporation

`func (o *CharactersCharacterIdContractsGetInner) GetForCorporation() bool`

GetForCorporation returns the ForCorporation field if non-nil, zero value otherwise.

### GetForCorporationOk

`func (o *CharactersCharacterIdContractsGetInner) GetForCorporationOk() (*bool, bool)`

GetForCorporationOk returns a tuple with the ForCorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForCorporation

`func (o *CharactersCharacterIdContractsGetInner) SetForCorporation(v bool)`

SetForCorporation sets ForCorporation field to given value.


### GetIssuerCorporationId

`func (o *CharactersCharacterIdContractsGetInner) GetIssuerCorporationId() int64`

GetIssuerCorporationId returns the IssuerCorporationId field if non-nil, zero value otherwise.

### GetIssuerCorporationIdOk

`func (o *CharactersCharacterIdContractsGetInner) GetIssuerCorporationIdOk() (*int64, bool)`

GetIssuerCorporationIdOk returns a tuple with the IssuerCorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCorporationId

`func (o *CharactersCharacterIdContractsGetInner) SetIssuerCorporationId(v int64)`

SetIssuerCorporationId sets IssuerCorporationId field to given value.


### GetIssuerId

`func (o *CharactersCharacterIdContractsGetInner) GetIssuerId() int64`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *CharactersCharacterIdContractsGetInner) GetIssuerIdOk() (*int64, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *CharactersCharacterIdContractsGetInner) SetIssuerId(v int64)`

SetIssuerId sets IssuerId field to given value.


### GetPrice

`func (o *CharactersCharacterIdContractsGetInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CharactersCharacterIdContractsGetInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CharactersCharacterIdContractsGetInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CharactersCharacterIdContractsGetInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReward

`func (o *CharactersCharacterIdContractsGetInner) GetReward() float64`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *CharactersCharacterIdContractsGetInner) GetRewardOk() (*float64, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *CharactersCharacterIdContractsGetInner) SetReward(v float64)`

SetReward sets Reward field to given value.

### HasReward

`func (o *CharactersCharacterIdContractsGetInner) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetStartLocationId

`func (o *CharactersCharacterIdContractsGetInner) GetStartLocationId() int64`

GetStartLocationId returns the StartLocationId field if non-nil, zero value otherwise.

### GetStartLocationIdOk

`func (o *CharactersCharacterIdContractsGetInner) GetStartLocationIdOk() (*int64, bool)`

GetStartLocationIdOk returns a tuple with the StartLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocationId

`func (o *CharactersCharacterIdContractsGetInner) SetStartLocationId(v int64)`

SetStartLocationId sets StartLocationId field to given value.

### HasStartLocationId

`func (o *CharactersCharacterIdContractsGetInner) HasStartLocationId() bool`

HasStartLocationId returns a boolean if a field has been set.

### GetStatus

`func (o *CharactersCharacterIdContractsGetInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CharactersCharacterIdContractsGetInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CharactersCharacterIdContractsGetInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *CharactersCharacterIdContractsGetInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CharactersCharacterIdContractsGetInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CharactersCharacterIdContractsGetInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CharactersCharacterIdContractsGetInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CharactersCharacterIdContractsGetInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharactersCharacterIdContractsGetInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharactersCharacterIdContractsGetInner) SetType(v string)`

SetType sets Type field to given value.


### GetVolume

`func (o *CharactersCharacterIdContractsGetInner) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CharactersCharacterIdContractsGetInner) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CharactersCharacterIdContractsGetInner) SetVolume(v float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CharactersCharacterIdContractsGetInner) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


