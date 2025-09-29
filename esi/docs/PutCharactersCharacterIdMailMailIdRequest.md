# PutCharactersCharacterIdMailMailIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]int64** | Labels to assign to the mail. Pre-existing labels are unassigned. | [optional] 
**Read** | Pointer to **bool** | Whether the mail is flagged as read | [optional] 

## Methods

### NewPutCharactersCharacterIdMailMailIdRequest

`func NewPutCharactersCharacterIdMailMailIdRequest() *PutCharactersCharacterIdMailMailIdRequest`

NewPutCharactersCharacterIdMailMailIdRequest instantiates a new PutCharactersCharacterIdMailMailIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutCharactersCharacterIdMailMailIdRequestWithDefaults

`func NewPutCharactersCharacterIdMailMailIdRequestWithDefaults() *PutCharactersCharacterIdMailMailIdRequest`

NewPutCharactersCharacterIdMailMailIdRequestWithDefaults instantiates a new PutCharactersCharacterIdMailMailIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *PutCharactersCharacterIdMailMailIdRequest) GetLabels() []int64`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PutCharactersCharacterIdMailMailIdRequest) GetLabelsOk() (*[]int64, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PutCharactersCharacterIdMailMailIdRequest) SetLabels(v []int64)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PutCharactersCharacterIdMailMailIdRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRead

`func (o *PutCharactersCharacterIdMailMailIdRequest) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *PutCharactersCharacterIdMailMailIdRequest) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *PutCharactersCharacterIdMailMailIdRequest) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *PutCharactersCharacterIdMailMailIdRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


