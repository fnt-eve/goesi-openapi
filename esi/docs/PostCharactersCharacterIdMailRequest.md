# PostCharactersCharacterIdMailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedCost** | Pointer to **int64** |  | [optional] [default to 0]
**Body** | **string** |  | 
**Recipients** | [**[]PostCharactersCharacterIdMailRequestRecipientsInner**](PostCharactersCharacterIdMailRequestRecipientsInner.md) |  | 
**Subject** | **string** |  | 

## Methods

### NewPostCharactersCharacterIdMailRequest

`func NewPostCharactersCharacterIdMailRequest(body string, recipients []PostCharactersCharacterIdMailRequestRecipientsInner, subject string, ) *PostCharactersCharacterIdMailRequest`

NewPostCharactersCharacterIdMailRequest instantiates a new PostCharactersCharacterIdMailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCharactersCharacterIdMailRequestWithDefaults

`func NewPostCharactersCharacterIdMailRequestWithDefaults() *PostCharactersCharacterIdMailRequest`

NewPostCharactersCharacterIdMailRequestWithDefaults instantiates a new PostCharactersCharacterIdMailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedCost

`func (o *PostCharactersCharacterIdMailRequest) GetApprovedCost() int64`

GetApprovedCost returns the ApprovedCost field if non-nil, zero value otherwise.

### GetApprovedCostOk

`func (o *PostCharactersCharacterIdMailRequest) GetApprovedCostOk() (*int64, bool)`

GetApprovedCostOk returns a tuple with the ApprovedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedCost

`func (o *PostCharactersCharacterIdMailRequest) SetApprovedCost(v int64)`

SetApprovedCost sets ApprovedCost field to given value.

### HasApprovedCost

`func (o *PostCharactersCharacterIdMailRequest) HasApprovedCost() bool`

HasApprovedCost returns a boolean if a field has been set.

### GetBody

`func (o *PostCharactersCharacterIdMailRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PostCharactersCharacterIdMailRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PostCharactersCharacterIdMailRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetRecipients

`func (o *PostCharactersCharacterIdMailRequest) GetRecipients() []PostCharactersCharacterIdMailRequestRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PostCharactersCharacterIdMailRequest) GetRecipientsOk() (*[]PostCharactersCharacterIdMailRequestRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PostCharactersCharacterIdMailRequest) SetRecipients(v []PostCharactersCharacterIdMailRequestRecipientsInner)`

SetRecipients sets Recipients field to given value.


### GetSubject

`func (o *PostCharactersCharacterIdMailRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PostCharactersCharacterIdMailRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PostCharactersCharacterIdMailRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


