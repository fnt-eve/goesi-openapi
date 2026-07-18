# MetaNameEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | The date ESI started insisting on this name (UTC-11). | 
**Name** | **string** | What those three letters allegedly stood for at the time. | 

## Methods

### NewMetaNameEntry

`func NewMetaNameEntry(date string, name string, ) *MetaNameEntry`

NewMetaNameEntry instantiates a new MetaNameEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaNameEntryWithDefaults

`func NewMetaNameEntryWithDefaults() *MetaNameEntry`

NewMetaNameEntryWithDefaults instantiates a new MetaNameEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *MetaNameEntry) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MetaNameEntry) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MetaNameEntry) SetDate(v string)`

SetDate sets Date field to given value.


### GetName

`func (o *MetaNameEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaNameEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaNameEntry) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


