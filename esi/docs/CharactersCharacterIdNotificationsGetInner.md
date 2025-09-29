# CharactersCharacterIdNotificationsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRead** | Pointer to **bool** |  | [optional] 
**NotificationId** | **int64** |  | 
**SenderId** | **int64** |  | 
**SenderType** | **string** |  | 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Type** | **string** |  | 

## Methods

### NewCharactersCharacterIdNotificationsGetInner

`func NewCharactersCharacterIdNotificationsGetInner(notificationId int64, senderId int64, senderType string, timestamp time.Time, type_ string, ) *CharactersCharacterIdNotificationsGetInner`

NewCharactersCharacterIdNotificationsGetInner instantiates a new CharactersCharacterIdNotificationsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdNotificationsGetInnerWithDefaults

`func NewCharactersCharacterIdNotificationsGetInnerWithDefaults() *CharactersCharacterIdNotificationsGetInner`

NewCharactersCharacterIdNotificationsGetInnerWithDefaults instantiates a new CharactersCharacterIdNotificationsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRead

`func (o *CharactersCharacterIdNotificationsGetInner) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *CharactersCharacterIdNotificationsGetInner) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *CharactersCharacterIdNotificationsGetInner) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *CharactersCharacterIdNotificationsGetInner) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetNotificationId

`func (o *CharactersCharacterIdNotificationsGetInner) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *CharactersCharacterIdNotificationsGetInner) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *CharactersCharacterIdNotificationsGetInner) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.


### GetSenderId

`func (o *CharactersCharacterIdNotificationsGetInner) GetSenderId() int64`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *CharactersCharacterIdNotificationsGetInner) GetSenderIdOk() (*int64, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *CharactersCharacterIdNotificationsGetInner) SetSenderId(v int64)`

SetSenderId sets SenderId field to given value.


### GetSenderType

`func (o *CharactersCharacterIdNotificationsGetInner) GetSenderType() string`

GetSenderType returns the SenderType field if non-nil, zero value otherwise.

### GetSenderTypeOk

`func (o *CharactersCharacterIdNotificationsGetInner) GetSenderTypeOk() (*string, bool)`

GetSenderTypeOk returns a tuple with the SenderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderType

`func (o *CharactersCharacterIdNotificationsGetInner) SetSenderType(v string)`

SetSenderType sets SenderType field to given value.


### GetText

`func (o *CharactersCharacterIdNotificationsGetInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CharactersCharacterIdNotificationsGetInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CharactersCharacterIdNotificationsGetInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CharactersCharacterIdNotificationsGetInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *CharactersCharacterIdNotificationsGetInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CharactersCharacterIdNotificationsGetInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CharactersCharacterIdNotificationsGetInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *CharactersCharacterIdNotificationsGetInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharactersCharacterIdNotificationsGetInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharactersCharacterIdNotificationsGetInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


