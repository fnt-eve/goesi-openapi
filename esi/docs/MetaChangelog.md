# MetaChangelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changelog** | [**map[string][]MetaChangelogEntry**](array.md) | Per date, list changes for that date | 

## Methods

### NewMetaChangelog

`func NewMetaChangelog(changelog map[string][]MetaChangelogEntry, ) *MetaChangelog`

NewMetaChangelog instantiates a new MetaChangelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaChangelogWithDefaults

`func NewMetaChangelogWithDefaults() *MetaChangelog`

NewMetaChangelogWithDefaults instantiates a new MetaChangelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangelog

`func (o *MetaChangelog) GetChangelog() map[string][]MetaChangelogEntry`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *MetaChangelog) GetChangelogOk() (*map[string][]MetaChangelogEntry, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *MetaChangelog) SetChangelog(v map[string][]MetaChangelogEntry)`

SetChangelog sets Changelog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


