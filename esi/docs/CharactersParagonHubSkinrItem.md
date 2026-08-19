# CharactersParagonHubSkinrItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | When the listing was created | 
**Expires** | **time.Time** | When the listing expires | 
**Id** | **string** |  | 
**LastModified** | **time.Time** | When this listing was last retrieved from the game | 
**Price** | [**CharactersParagonHubSkinrItemPrice**](CharactersParagonHubSkinrItemPrice.md) |  | 
**Quantity** | **int64** | How many licenses remain in this listing | 
**SellerId** | **int64** |  | 
**SkinrId** | **string** | SKINR license identifier | 
**State** | **string** | Lifecycle state of the listing | 
**Target** | [**CharactersParagonHubSkinrItemTarget**](CharactersParagonHubSkinrItemTarget.md) |  | 

## Methods

### NewCharactersParagonHubSkinrItem

`func NewCharactersParagonHubSkinrItem(created time.Time, expires time.Time, id string, lastModified time.Time, price CharactersParagonHubSkinrItemPrice, quantity int64, sellerId int64, skinrId string, state string, target CharactersParagonHubSkinrItemTarget, ) *CharactersParagonHubSkinrItem`

NewCharactersParagonHubSkinrItem instantiates a new CharactersParagonHubSkinrItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersParagonHubSkinrItemWithDefaults

`func NewCharactersParagonHubSkinrItemWithDefaults() *CharactersParagonHubSkinrItem`

NewCharactersParagonHubSkinrItemWithDefaults instantiates a new CharactersParagonHubSkinrItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CharactersParagonHubSkinrItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CharactersParagonHubSkinrItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CharactersParagonHubSkinrItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *CharactersParagonHubSkinrItem) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CharactersParagonHubSkinrItem) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CharactersParagonHubSkinrItem) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetId

`func (o *CharactersParagonHubSkinrItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CharactersParagonHubSkinrItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CharactersParagonHubSkinrItem) SetId(v string)`

SetId sets Id field to given value.


### GetLastModified

`func (o *CharactersParagonHubSkinrItem) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CharactersParagonHubSkinrItem) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CharactersParagonHubSkinrItem) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetPrice

`func (o *CharactersParagonHubSkinrItem) GetPrice() CharactersParagonHubSkinrItemPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CharactersParagonHubSkinrItem) GetPriceOk() (*CharactersParagonHubSkinrItemPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CharactersParagonHubSkinrItem) SetPrice(v CharactersParagonHubSkinrItemPrice)`

SetPrice sets Price field to given value.


### GetQuantity

`func (o *CharactersParagonHubSkinrItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CharactersParagonHubSkinrItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CharactersParagonHubSkinrItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetSellerId

`func (o *CharactersParagonHubSkinrItem) GetSellerId() int64`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *CharactersParagonHubSkinrItem) GetSellerIdOk() (*int64, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *CharactersParagonHubSkinrItem) SetSellerId(v int64)`

SetSellerId sets SellerId field to given value.


### GetSkinrId

`func (o *CharactersParagonHubSkinrItem) GetSkinrId() string`

GetSkinrId returns the SkinrId field if non-nil, zero value otherwise.

### GetSkinrIdOk

`func (o *CharactersParagonHubSkinrItem) GetSkinrIdOk() (*string, bool)`

GetSkinrIdOk returns a tuple with the SkinrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkinrId

`func (o *CharactersParagonHubSkinrItem) SetSkinrId(v string)`

SetSkinrId sets SkinrId field to given value.


### GetState

`func (o *CharactersParagonHubSkinrItem) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CharactersParagonHubSkinrItem) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CharactersParagonHubSkinrItem) SetState(v string)`

SetState sets State field to given value.


### GetTarget

`func (o *CharactersParagonHubSkinrItem) GetTarget() CharactersParagonHubSkinrItemTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CharactersParagonHubSkinrItem) GetTargetOk() (*CharactersParagonHubSkinrItemTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CharactersParagonHubSkinrItem) SetTarget(v CharactersParagonHubSkinrItemTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


