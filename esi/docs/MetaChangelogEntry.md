# MetaChangelogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityDate** | **string** |  | 
**Description** | **string** | Description | 
**IsBreaking** | **bool** | Whether this is a breaking change | 
**Method** | **string** | HTTP method of the route | 
**Path** | **string** | Path of the route | 

## Methods

### NewMetaChangelogEntry

`func NewMetaChangelogEntry(compatibilityDate string, description string, isBreaking bool, method string, path string, ) *MetaChangelogEntry`

NewMetaChangelogEntry instantiates a new MetaChangelogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaChangelogEntryWithDefaults

`func NewMetaChangelogEntryWithDefaults() *MetaChangelogEntry`

NewMetaChangelogEntryWithDefaults instantiates a new MetaChangelogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityDate

`func (o *MetaChangelogEntry) GetCompatibilityDate() string`

GetCompatibilityDate returns the CompatibilityDate field if non-nil, zero value otherwise.

### GetCompatibilityDateOk

`func (o *MetaChangelogEntry) GetCompatibilityDateOk() (*string, bool)`

GetCompatibilityDateOk returns a tuple with the CompatibilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityDate

`func (o *MetaChangelogEntry) SetCompatibilityDate(v string)`

SetCompatibilityDate sets CompatibilityDate field to given value.


### GetDescription

`func (o *MetaChangelogEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetaChangelogEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetaChangelogEntry) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsBreaking

`func (o *MetaChangelogEntry) GetIsBreaking() bool`

GetIsBreaking returns the IsBreaking field if non-nil, zero value otherwise.

### GetIsBreakingOk

`func (o *MetaChangelogEntry) GetIsBreakingOk() (*bool, bool)`

GetIsBreakingOk returns a tuple with the IsBreaking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreaking

`func (o *MetaChangelogEntry) SetIsBreaking(v bool)`

SetIsBreaking sets IsBreaking field to given value.


### GetMethod

`func (o *MetaChangelogEntry) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MetaChangelogEntry) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MetaChangelogEntry) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPath

`func (o *MetaChangelogEntry) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MetaChangelogEntry) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MetaChangelogEntry) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


