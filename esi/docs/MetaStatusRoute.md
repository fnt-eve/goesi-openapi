# MetaStatusRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | Route&#39;s HTTP method | 
**Path** | **string** | Route&#39;s HTTP path | 
**Status** | **string** | Route&#39;s health status | 

## Methods

### NewMetaStatusRoute

`func NewMetaStatusRoute(method string, path string, status string, ) *MetaStatusRoute`

NewMetaStatusRoute instantiates a new MetaStatusRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaStatusRouteWithDefaults

`func NewMetaStatusRouteWithDefaults() *MetaStatusRoute`

NewMetaStatusRouteWithDefaults instantiates a new MetaStatusRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *MetaStatusRoute) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MetaStatusRoute) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MetaStatusRoute) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPath

`func (o *MetaStatusRoute) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MetaStatusRoute) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MetaStatusRoute) SetPath(v string)`

SetPath sets Path field to given value.


### GetStatus

`func (o *MetaStatusRoute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetaStatusRoute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetaStatusRoute) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


