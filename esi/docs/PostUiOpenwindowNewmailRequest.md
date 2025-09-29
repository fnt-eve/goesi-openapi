# PostUiOpenwindowNewmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** |  | 
**Recipients** | **[]int64** |  | 
**Subject** | **string** |  | 
**ToCorpOrAllianceId** | Pointer to **int64** |  | [optional] 
**ToMailingListId** | Pointer to **int64** | Corporations, alliances and mailing lists are all types of mailing groups. You may only send to one mailing group, at a time, so you may fill out either this field or the to_corp_or_alliance_ids field | [optional] 

## Methods

### NewPostUiOpenwindowNewmailRequest

`func NewPostUiOpenwindowNewmailRequest(body string, recipients []int64, subject string, ) *PostUiOpenwindowNewmailRequest`

NewPostUiOpenwindowNewmailRequest instantiates a new PostUiOpenwindowNewmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUiOpenwindowNewmailRequestWithDefaults

`func NewPostUiOpenwindowNewmailRequestWithDefaults() *PostUiOpenwindowNewmailRequest`

NewPostUiOpenwindowNewmailRequestWithDefaults instantiates a new PostUiOpenwindowNewmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PostUiOpenwindowNewmailRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PostUiOpenwindowNewmailRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PostUiOpenwindowNewmailRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetRecipients

`func (o *PostUiOpenwindowNewmailRequest) GetRecipients() []int64`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PostUiOpenwindowNewmailRequest) GetRecipientsOk() (*[]int64, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PostUiOpenwindowNewmailRequest) SetRecipients(v []int64)`

SetRecipients sets Recipients field to given value.


### GetSubject

`func (o *PostUiOpenwindowNewmailRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PostUiOpenwindowNewmailRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PostUiOpenwindowNewmailRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetToCorpOrAllianceId

`func (o *PostUiOpenwindowNewmailRequest) GetToCorpOrAllianceId() int64`

GetToCorpOrAllianceId returns the ToCorpOrAllianceId field if non-nil, zero value otherwise.

### GetToCorpOrAllianceIdOk

`func (o *PostUiOpenwindowNewmailRequest) GetToCorpOrAllianceIdOk() (*int64, bool)`

GetToCorpOrAllianceIdOk returns a tuple with the ToCorpOrAllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCorpOrAllianceId

`func (o *PostUiOpenwindowNewmailRequest) SetToCorpOrAllianceId(v int64)`

SetToCorpOrAllianceId sets ToCorpOrAllianceId field to given value.

### HasToCorpOrAllianceId

`func (o *PostUiOpenwindowNewmailRequest) HasToCorpOrAllianceId() bool`

HasToCorpOrAllianceId returns a boolean if a field has been set.

### GetToMailingListId

`func (o *PostUiOpenwindowNewmailRequest) GetToMailingListId() int64`

GetToMailingListId returns the ToMailingListId field if non-nil, zero value otherwise.

### GetToMailingListIdOk

`func (o *PostUiOpenwindowNewmailRequest) GetToMailingListIdOk() (*int64, bool)`

GetToMailingListIdOk returns a tuple with the ToMailingListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToMailingListId

`func (o *PostUiOpenwindowNewmailRequest) SetToMailingListId(v int64)`

SetToMailingListId sets ToMailingListId field to given value.

### HasToMailingListId

`func (o *PostUiOpenwindowNewmailRequest) HasToMailingListId() bool`

HasToMailingListId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


