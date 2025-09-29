# CharactersCharacterIdMailMailIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Mail&#39;s body | [optional] 
**From** | Pointer to **int64** | From whom the mail was sent | [optional] 
**Labels** | Pointer to **[]int64** | Labels attached to the mail | [optional] 
**Read** | Pointer to **bool** | Whether the mail is flagged as read | [optional] 
**Recipients** | Pointer to [**[]PostCharactersCharacterIdMailRequestRecipientsInner**](PostCharactersCharacterIdMailRequestRecipientsInner.md) | Recipients of the mail | [optional] 
**Subject** | Pointer to **string** | Mail subject | [optional] 
**Timestamp** | Pointer to **time.Time** | When the mail was sent | [optional] 

## Methods

### NewCharactersCharacterIdMailMailIdGet

`func NewCharactersCharacterIdMailMailIdGet() *CharactersCharacterIdMailMailIdGet`

NewCharactersCharacterIdMailMailIdGet instantiates a new CharactersCharacterIdMailMailIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdMailMailIdGetWithDefaults

`func NewCharactersCharacterIdMailMailIdGetWithDefaults() *CharactersCharacterIdMailMailIdGet`

NewCharactersCharacterIdMailMailIdGetWithDefaults instantiates a new CharactersCharacterIdMailMailIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CharactersCharacterIdMailMailIdGet) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CharactersCharacterIdMailMailIdGet) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CharactersCharacterIdMailMailIdGet) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CharactersCharacterIdMailMailIdGet) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFrom

`func (o *CharactersCharacterIdMailMailIdGet) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CharactersCharacterIdMailMailIdGet) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CharactersCharacterIdMailMailIdGet) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CharactersCharacterIdMailMailIdGet) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLabels

`func (o *CharactersCharacterIdMailMailIdGet) GetLabels() []int64`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CharactersCharacterIdMailMailIdGet) GetLabelsOk() (*[]int64, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CharactersCharacterIdMailMailIdGet) SetLabels(v []int64)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CharactersCharacterIdMailMailIdGet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRead

`func (o *CharactersCharacterIdMailMailIdGet) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *CharactersCharacterIdMailMailIdGet) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *CharactersCharacterIdMailMailIdGet) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *CharactersCharacterIdMailMailIdGet) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetRecipients

`func (o *CharactersCharacterIdMailMailIdGet) GetRecipients() []PostCharactersCharacterIdMailRequestRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CharactersCharacterIdMailMailIdGet) GetRecipientsOk() (*[]PostCharactersCharacterIdMailRequestRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CharactersCharacterIdMailMailIdGet) SetRecipients(v []PostCharactersCharacterIdMailRequestRecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CharactersCharacterIdMailMailIdGet) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSubject

`func (o *CharactersCharacterIdMailMailIdGet) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CharactersCharacterIdMailMailIdGet) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CharactersCharacterIdMailMailIdGet) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CharactersCharacterIdMailMailIdGet) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTimestamp

`func (o *CharactersCharacterIdMailMailIdGet) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CharactersCharacterIdMailMailIdGet) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CharactersCharacterIdMailMailIdGet) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CharactersCharacterIdMailMailIdGet) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


