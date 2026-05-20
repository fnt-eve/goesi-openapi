# CharactersAccessListsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The Access List&#39;s description | 
**Id** | **int64** |  | 
**Membership** | [**CharactersAccessListsDetailMembership**](CharactersAccessListsDetailMembership.md) |  | 
**Name** | **string** | The Access List&#39;s name | 

## Methods

### NewCharactersAccessListsDetail

`func NewCharactersAccessListsDetail(description string, id int64, membership CharactersAccessListsDetailMembership, name string, ) *CharactersAccessListsDetail`

NewCharactersAccessListsDetail instantiates a new CharactersAccessListsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersAccessListsDetailWithDefaults

`func NewCharactersAccessListsDetailWithDefaults() *CharactersAccessListsDetail`

NewCharactersAccessListsDetailWithDefaults instantiates a new CharactersAccessListsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CharactersAccessListsDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CharactersAccessListsDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CharactersAccessListsDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *CharactersAccessListsDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CharactersAccessListsDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CharactersAccessListsDetail) SetId(v int64)`

SetId sets Id field to given value.


### GetMembership

`func (o *CharactersAccessListsDetail) GetMembership() CharactersAccessListsDetailMembership`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *CharactersAccessListsDetail) GetMembershipOk() (*CharactersAccessListsDetailMembership, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *CharactersAccessListsDetail) SetMembership(v CharactersAccessListsDetailMembership)`

SetMembership sets Membership field to given value.


### GetName

`func (o *CharactersAccessListsDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharactersAccessListsDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharactersAccessListsDetail) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


