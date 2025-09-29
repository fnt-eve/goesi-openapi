# PutFleetsFleetIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFreeMove** | Pointer to **bool** | Should free-move be enabled in the fleet | [optional] 
**Motd** | Pointer to **string** | New fleet MOTD in CCP flavoured HTML | [optional] 

## Methods

### NewPutFleetsFleetIdRequest

`func NewPutFleetsFleetIdRequest() *PutFleetsFleetIdRequest`

NewPutFleetsFleetIdRequest instantiates a new PutFleetsFleetIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFleetsFleetIdRequestWithDefaults

`func NewPutFleetsFleetIdRequestWithDefaults() *PutFleetsFleetIdRequest`

NewPutFleetsFleetIdRequestWithDefaults instantiates a new PutFleetsFleetIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFreeMove

`func (o *PutFleetsFleetIdRequest) GetIsFreeMove() bool`

GetIsFreeMove returns the IsFreeMove field if non-nil, zero value otherwise.

### GetIsFreeMoveOk

`func (o *PutFleetsFleetIdRequest) GetIsFreeMoveOk() (*bool, bool)`

GetIsFreeMoveOk returns a tuple with the IsFreeMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeMove

`func (o *PutFleetsFleetIdRequest) SetIsFreeMove(v bool)`

SetIsFreeMove sets IsFreeMove field to given value.

### HasIsFreeMove

`func (o *PutFleetsFleetIdRequest) HasIsFreeMove() bool`

HasIsFreeMove returns a boolean if a field has been set.

### GetMotd

`func (o *PutFleetsFleetIdRequest) GetMotd() string`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *PutFleetsFleetIdRequest) GetMotdOk() (*string, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *PutFleetsFleetIdRequest) SetMotd(v string)`

SetMotd sets Motd field to given value.

### HasMotd

`func (o *PutFleetsFleetIdRequest) HasMotd() bool`

HasMotd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


