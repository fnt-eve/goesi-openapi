# CharactersCharacterIdNotificationsContactsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**NotificationId** | **int64** |  | 
**SendDate** | **time.Time** |  | 
**SenderCharacterId** | **int64** |  | 
**StandingLevel** | **float64** | A number representing the standing level the receiver has been added at by the sender. The standing levels are as follows: -10 -&gt; Terrible | -5 -&gt; Bad |  0 -&gt; Neutral |  5 -&gt; Good |  10 -&gt; Excellent | 

## Methods

### NewCharactersCharacterIdNotificationsContactsGetInner

`func NewCharactersCharacterIdNotificationsContactsGetInner(message string, notificationId int64, sendDate time.Time, senderCharacterId int64, standingLevel float64, ) *CharactersCharacterIdNotificationsContactsGetInner`

NewCharactersCharacterIdNotificationsContactsGetInner instantiates a new CharactersCharacterIdNotificationsContactsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdNotificationsContactsGetInnerWithDefaults

`func NewCharactersCharacterIdNotificationsContactsGetInnerWithDefaults() *CharactersCharacterIdNotificationsContactsGetInner`

NewCharactersCharacterIdNotificationsContactsGetInnerWithDefaults instantiates a new CharactersCharacterIdNotificationsContactsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CharactersCharacterIdNotificationsContactsGetInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNotificationId

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *CharactersCharacterIdNotificationsContactsGetInner) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.


### GetSendDate

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *CharactersCharacterIdNotificationsContactsGetInner) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.


### GetSenderCharacterId

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetSenderCharacterId() int64`

GetSenderCharacterId returns the SenderCharacterId field if non-nil, zero value otherwise.

### GetSenderCharacterIdOk

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetSenderCharacterIdOk() (*int64, bool)`

GetSenderCharacterIdOk returns a tuple with the SenderCharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderCharacterId

`func (o *CharactersCharacterIdNotificationsContactsGetInner) SetSenderCharacterId(v int64)`

SetSenderCharacterId sets SenderCharacterId field to given value.


### GetStandingLevel

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetStandingLevel() float64`

GetStandingLevel returns the StandingLevel field if non-nil, zero value otherwise.

### GetStandingLevelOk

`func (o *CharactersCharacterIdNotificationsContactsGetInner) GetStandingLevelOk() (*float64, bool)`

GetStandingLevelOk returns a tuple with the StandingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingLevel

`func (o *CharactersCharacterIdNotificationsContactsGetInner) SetStandingLevel(v float64)`

SetStandingLevel sets StandingLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


