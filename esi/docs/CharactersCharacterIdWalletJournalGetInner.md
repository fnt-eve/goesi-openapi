# CharactersCharacterIdWalletJournalGetInner

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
**RefType** | **string** | \&quot;The transaction type for the given. transaction. Different transaction types will populate different attributes.\&quot; | 
**SecondPartyId** | Pointer to **int64** | The id of the second party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name | [optional] 
**Tax** | Pointer to **float64** | Tax amount received. Only applies to tax related transactions | [optional] 
**TaxReceiverId** | Pointer to **int64** | The corporation ID receiving any tax paid. Only applies to tax related transactions | [optional] 

## Methods

### NewCharactersCharacterIdWalletJournalGetInner

`func NewCharactersCharacterIdWalletJournalGetInner(date time.Time, description string, id int64, refType string, ) *CharactersCharacterIdWalletJournalGetInner`

NewCharactersCharacterIdWalletJournalGetInner instantiates a new CharactersCharacterIdWalletJournalGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdWalletJournalGetInnerWithDefaults

`func NewCharactersCharacterIdWalletJournalGetInnerWithDefaults() *CharactersCharacterIdWalletJournalGetInner`

NewCharactersCharacterIdWalletJournalGetInnerWithDefaults instantiates a new CharactersCharacterIdWalletJournalGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CharactersCharacterIdWalletJournalGetInner) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CharactersCharacterIdWalletJournalGetInner) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CharactersCharacterIdWalletJournalGetInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalance

`func (o *CharactersCharacterIdWalletJournalGetInner) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CharactersCharacterIdWalletJournalGetInner) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CharactersCharacterIdWalletJournalGetInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetContextId

`func (o *CharactersCharacterIdWalletJournalGetInner) GetContextId() int64`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetContextIdOk() (*int64, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *CharactersCharacterIdWalletJournalGetInner) SetContextId(v int64)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *CharactersCharacterIdWalletJournalGetInner) HasContextId() bool`

HasContextId returns a boolean if a field has been set.

### GetContextIdType

`func (o *CharactersCharacterIdWalletJournalGetInner) GetContextIdType() string`

GetContextIdType returns the ContextIdType field if non-nil, zero value otherwise.

### GetContextIdTypeOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetContextIdTypeOk() (*string, bool)`

GetContextIdTypeOk returns a tuple with the ContextIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextIdType

`func (o *CharactersCharacterIdWalletJournalGetInner) SetContextIdType(v string)`

SetContextIdType sets ContextIdType field to given value.

### HasContextIdType

`func (o *CharactersCharacterIdWalletJournalGetInner) HasContextIdType() bool`

HasContextIdType returns a boolean if a field has been set.

### GetDate

`func (o *CharactersCharacterIdWalletJournalGetInner) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CharactersCharacterIdWalletJournalGetInner) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetDescription

`func (o *CharactersCharacterIdWalletJournalGetInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CharactersCharacterIdWalletJournalGetInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFirstPartyId

`func (o *CharactersCharacterIdWalletJournalGetInner) GetFirstPartyId() int64`

GetFirstPartyId returns the FirstPartyId field if non-nil, zero value otherwise.

### GetFirstPartyIdOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetFirstPartyIdOk() (*int64, bool)`

GetFirstPartyIdOk returns a tuple with the FirstPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPartyId

`func (o *CharactersCharacterIdWalletJournalGetInner) SetFirstPartyId(v int64)`

SetFirstPartyId sets FirstPartyId field to given value.

### HasFirstPartyId

`func (o *CharactersCharacterIdWalletJournalGetInner) HasFirstPartyId() bool`

HasFirstPartyId returns a boolean if a field has been set.

### GetId

`func (o *CharactersCharacterIdWalletJournalGetInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CharactersCharacterIdWalletJournalGetInner) SetId(v int64)`

SetId sets Id field to given value.


### GetReason

`func (o *CharactersCharacterIdWalletJournalGetInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CharactersCharacterIdWalletJournalGetInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CharactersCharacterIdWalletJournalGetInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRefType

`func (o *CharactersCharacterIdWalletJournalGetInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *CharactersCharacterIdWalletJournalGetInner) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetSecondPartyId

`func (o *CharactersCharacterIdWalletJournalGetInner) GetSecondPartyId() int64`

GetSecondPartyId returns the SecondPartyId field if non-nil, zero value otherwise.

### GetSecondPartyIdOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetSecondPartyIdOk() (*int64, bool)`

GetSecondPartyIdOk returns a tuple with the SecondPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondPartyId

`func (o *CharactersCharacterIdWalletJournalGetInner) SetSecondPartyId(v int64)`

SetSecondPartyId sets SecondPartyId field to given value.

### HasSecondPartyId

`func (o *CharactersCharacterIdWalletJournalGetInner) HasSecondPartyId() bool`

HasSecondPartyId returns a boolean if a field has been set.

### GetTax

`func (o *CharactersCharacterIdWalletJournalGetInner) GetTax() float64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetTaxOk() (*float64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *CharactersCharacterIdWalletJournalGetInner) SetTax(v float64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *CharactersCharacterIdWalletJournalGetInner) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxReceiverId

`func (o *CharactersCharacterIdWalletJournalGetInner) GetTaxReceiverId() int64`

GetTaxReceiverId returns the TaxReceiverId field if non-nil, zero value otherwise.

### GetTaxReceiverIdOk

`func (o *CharactersCharacterIdWalletJournalGetInner) GetTaxReceiverIdOk() (*int64, bool)`

GetTaxReceiverIdOk returns a tuple with the TaxReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxReceiverId

`func (o *CharactersCharacterIdWalletJournalGetInner) SetTaxReceiverId(v int64)`

SetTaxReceiverId sets TaxReceiverId field to given value.

### HasTaxReceiverId

`func (o *CharactersCharacterIdWalletJournalGetInner) HasTaxReceiverId() bool`

HasTaxReceiverId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


