# CorporationsCorporationIdWalletsDivisionJournalGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | The amount of ISK given or taken from the wallet as a result of the given transaction. Positive when ISK is deposited into the wallet and negative when ISK is withdrawn | [optional] 
**Balance** | Pointer to **float64** | Wallet balance after transaction occurred | [optional] 
**ContextId** | Pointer to **int64** | An ID that gives extra context to the particular transaction. Because of legacy reasons the context is completely different per ref_type and means different things. It is also possible to not have a context_id | [optional] 
**ContextIdType** | Pointer to **string** | The type of the given context_id if present | [optional] 
**Date** | **time.Time** | Date and time of transaction | 
**Description** | **string** | The reason for the transaction, mirrors what is seen in the client | 
**FirstPartyId** | Pointer to **int64** | The id of the first party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name | [optional] 
**Id** | **int64** | Unique journal reference ID | 
**Reason** | Pointer to **string** | The user stated reason for the transaction. Only applies to some ref_types | [optional] 
**RefType** | **string** | \&quot;The transaction type for the given. transaction. Different transaction types will populate different attributes. Note: If you have an existing XML API application that is using ref_types, you will need to know which string ESI ref_type maps to which integer. You can look at the following file to see string-&gt;int mappings: https://github.com/ccpgames/eve-glue/blob/master/eve_glue/wallet_journal_ref.py\&quot; | 
**SecondPartyId** | Pointer to **int64** | The id of the second party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name | [optional] 
**Tax** | Pointer to **float64** | Tax amount received. Only applies to tax related transactions | [optional] 
**TaxReceiverId** | Pointer to **int64** | The corporation ID receiving any tax paid. Only applies to tax related transactions | [optional] 

## Methods

### NewCorporationsCorporationIdWalletsDivisionJournalGetInner

`func NewCorporationsCorporationIdWalletsDivisionJournalGetInner(date time.Time, description string, id int64, refType string, ) *CorporationsCorporationIdWalletsDivisionJournalGetInner`

NewCorporationsCorporationIdWalletsDivisionJournalGetInner instantiates a new CorporationsCorporationIdWalletsDivisionJournalGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdWalletsDivisionJournalGetInnerWithDefaults

`func NewCorporationsCorporationIdWalletsDivisionJournalGetInnerWithDefaults() *CorporationsCorporationIdWalletsDivisionJournalGetInner`

NewCorporationsCorporationIdWalletsDivisionJournalGetInnerWithDefaults instantiates a new CorporationsCorporationIdWalletsDivisionJournalGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalance

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetContextId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetContextId() int64`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetContextIdOk() (*int64, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetContextId(v int64)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasContextId() bool`

HasContextId returns a boolean if a field has been set.

### GetContextIdType

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetContextIdType() string`

GetContextIdType returns the ContextIdType field if non-nil, zero value otherwise.

### GetContextIdTypeOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetContextIdTypeOk() (*string, bool)`

GetContextIdTypeOk returns a tuple with the ContextIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextIdType

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetContextIdType(v string)`

SetContextIdType sets ContextIdType field to given value.

### HasContextIdType

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasContextIdType() bool`

HasContextIdType returns a boolean if a field has been set.

### GetDate

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetDescription

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFirstPartyId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetFirstPartyId() int64`

GetFirstPartyId returns the FirstPartyId field if non-nil, zero value otherwise.

### GetFirstPartyIdOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetFirstPartyIdOk() (*int64, bool)`

GetFirstPartyIdOk returns a tuple with the FirstPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPartyId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetFirstPartyId(v int64)`

SetFirstPartyId sets FirstPartyId field to given value.

### HasFirstPartyId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasFirstPartyId() bool`

HasFirstPartyId returns a boolean if a field has been set.

### GetId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetId(v int64)`

SetId sets Id field to given value.


### GetReason

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRefType

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetSecondPartyId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetSecondPartyId() int64`

GetSecondPartyId returns the SecondPartyId field if non-nil, zero value otherwise.

### GetSecondPartyIdOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetSecondPartyIdOk() (*int64, bool)`

GetSecondPartyIdOk returns a tuple with the SecondPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondPartyId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetSecondPartyId(v int64)`

SetSecondPartyId sets SecondPartyId field to given value.

### HasSecondPartyId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasSecondPartyId() bool`

HasSecondPartyId returns a boolean if a field has been set.

### GetTax

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetTax() float64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetTaxOk() (*float64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetTax(v float64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxReceiverId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetTaxReceiverId() int64`

GetTaxReceiverId returns the TaxReceiverId field if non-nil, zero value otherwise.

### GetTaxReceiverIdOk

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) GetTaxReceiverIdOk() (*int64, bool)`

GetTaxReceiverIdOk returns a tuple with the TaxReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxReceiverId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) SetTaxReceiverId(v int64)`

SetTaxReceiverId sets TaxReceiverId field to given value.

### HasTaxReceiverId

`func (o *CorporationsCorporationIdWalletsDivisionJournalGetInner) HasTaxReceiverId() bool`

HasTaxReceiverId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


