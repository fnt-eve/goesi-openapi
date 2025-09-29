# CorporationsCorporationIdWalletsDivisionTransactionsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **int64** |  | 
**Date** | **time.Time** | Date and time of transaction | 
**IsBuy** | **bool** |  | 
**JournalRefId** | **int64** | -1 if there is no corresponding wallet journal entry | 
**LocationId** | **int64** |  | 
**Quantity** | **int64** |  | 
**TransactionId** | **int64** | Unique transaction ID | 
**TypeId** | **int64** |  | 
**UnitPrice** | **float64** | Amount paid per unit | 

## Methods

### NewCorporationsCorporationIdWalletsDivisionTransactionsGetInner

`func NewCorporationsCorporationIdWalletsDivisionTransactionsGetInner(clientId int64, date time.Time, isBuy bool, journalRefId int64, locationId int64, quantity int64, transactionId int64, typeId int64, unitPrice float64, ) *CorporationsCorporationIdWalletsDivisionTransactionsGetInner`

NewCorporationsCorporationIdWalletsDivisionTransactionsGetInner instantiates a new CorporationsCorporationIdWalletsDivisionTransactionsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdWalletsDivisionTransactionsGetInnerWithDefaults

`func NewCorporationsCorporationIdWalletsDivisionTransactionsGetInnerWithDefaults() *CorporationsCorporationIdWalletsDivisionTransactionsGetInner`

NewCorporationsCorporationIdWalletsDivisionTransactionsGetInnerWithDefaults instantiates a new CorporationsCorporationIdWalletsDivisionTransactionsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetClientId(v int64)`

SetClientId sets ClientId field to given value.


### GetDate

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetIsBuy

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetIsBuy() bool`

GetIsBuy returns the IsBuy field if non-nil, zero value otherwise.

### GetIsBuyOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetIsBuyOk() (*bool, bool)`

GetIsBuyOk returns a tuple with the IsBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuy

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetIsBuy(v bool)`

SetIsBuy sets IsBuy field to given value.


### GetJournalRefId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetJournalRefId() int64`

GetJournalRefId returns the JournalRefId field if non-nil, zero value otherwise.

### GetJournalRefIdOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetJournalRefIdOk() (*int64, bool)`

GetJournalRefIdOk returns a tuple with the JournalRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalRefId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetJournalRefId(v int64)`

SetJournalRefId sets JournalRefId field to given value.


### GetLocationId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetQuantity

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetTransactionId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.


### GetTypeId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetUnitPrice

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CorporationsCorporationIdWalletsDivisionTransactionsGetInner) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


