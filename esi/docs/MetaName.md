# MetaName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **string** | The name ESI is going by today. It was almost certainly something else last time you looked. | 
**History** | [**[]MetaNameEntry**](MetaNameEntry.md) | Every name ESI has ever gone by, newest first. It has never once been the same twice. | 

## Methods

### NewMetaName

`func NewMetaName(current string, history []MetaNameEntry, ) *MetaName`

NewMetaName instantiates a new MetaName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaNameWithDefaults

`func NewMetaNameWithDefaults() *MetaName`

NewMetaNameWithDefaults instantiates a new MetaName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *MetaName) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *MetaName) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *MetaName) SetCurrent(v string)`

SetCurrent sets Current field to given value.


### GetHistory

`func (o *MetaName) GetHistory() []MetaNameEntry`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *MetaName) GetHistoryOk() (*[]MetaNameEntry, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *MetaName) SetHistory(v []MetaNameEntry)`

SetHistory sets History field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


