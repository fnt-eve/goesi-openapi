# CharactersCharacterIdMailGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int64** | From whom the mail was sent | [optional] 
**IsRead** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **[]int64** |  | [optional] 
**MailId** | Pointer to **int64** |  | [optional] 
**Recipients** | Pointer to [**[]PostCharactersCharacterIdMailRequestRecipientsInner**](PostCharactersCharacterIdMailRequestRecipientsInner.md) | Recipients of the mail | [optional] 
**Subject** | Pointer to **string** | Mail subject | [optional] 
**Timestamp** | Pointer to **time.Time** | When the mail was sent | [optional] 

## Methods

### NewCharactersCharacterIdMailGetInner

`func NewCharactersCharacterIdMailGetInner() *CharactersCharacterIdMailGetInner`

NewCharactersCharacterIdMailGetInner instantiates a new CharactersCharacterIdMailGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdMailGetInnerWithDefaults

`func NewCharactersCharacterIdMailGetInnerWithDefaults() *CharactersCharacterIdMailGetInner`

NewCharactersCharacterIdMailGetInnerWithDefaults instantiates a new CharactersCharacterIdMailGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CharactersCharacterIdMailGetInner) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CharactersCharacterIdMailGetInner) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CharactersCharacterIdMailGetInner) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CharactersCharacterIdMailGetInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIsRead

`func (o *CharactersCharacterIdMailGetInner) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *CharactersCharacterIdMailGetInner) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *CharactersCharacterIdMailGetInner) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *CharactersCharacterIdMailGetInner) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetLabels

`func (o *CharactersCharacterIdMailGetInner) GetLabels() []int64`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CharactersCharacterIdMailGetInner) GetLabelsOk() (*[]int64, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CharactersCharacterIdMailGetInner) SetLabels(v []int64)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CharactersCharacterIdMailGetInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMailId

`func (o *CharactersCharacterIdMailGetInner) GetMailId() int64`

GetMailId returns the MailId field if non-nil, zero value otherwise.

### GetMailIdOk

`func (o *CharactersCharacterIdMailGetInner) GetMailIdOk() (*int64, bool)`

GetMailIdOk returns a tuple with the MailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailId

`func (o *CharactersCharacterIdMailGetInner) SetMailId(v int64)`

SetMailId sets MailId field to given value.

### HasMailId

`func (o *CharactersCharacterIdMailGetInner) HasMailId() bool`

HasMailId returns a boolean if a field has been set.

### GetRecipients

`func (o *CharactersCharacterIdMailGetInner) GetRecipients() []PostCharactersCharacterIdMailRequestRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CharactersCharacterIdMailGetInner) GetRecipientsOk() (*[]PostCharactersCharacterIdMailRequestRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CharactersCharacterIdMailGetInner) SetRecipients(v []PostCharactersCharacterIdMailRequestRecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CharactersCharacterIdMailGetInner) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSubject

`func (o *CharactersCharacterIdMailGetInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CharactersCharacterIdMailGetInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CharactersCharacterIdMailGetInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CharactersCharacterIdMailGetInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTimestamp

`func (o *CharactersCharacterIdMailGetInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CharactersCharacterIdMailGetInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CharactersCharacterIdMailGetInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CharactersCharacterIdMailGetInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


