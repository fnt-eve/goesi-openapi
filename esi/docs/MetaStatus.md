# MetaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | [**[]MetaStatusRoutestatus**](MetaStatusRoutestatus.md) | List of all API routes and their health status | 

## Methods

### NewMetaStatus

`func NewMetaStatus(routes []MetaStatusRoutestatus, ) *MetaStatus`

NewMetaStatus instantiates a new MetaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaStatusWithDefaults

`func NewMetaStatusWithDefaults() *MetaStatus`

NewMetaStatusWithDefaults instantiates a new MetaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *MetaStatus) GetRoutes() []MetaStatusRoutestatus`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *MetaStatus) GetRoutesOk() (*[]MetaStatusRoutestatus, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *MetaStatus) SetRoutes(v []MetaStatusRoutestatus)`

SetRoutes sets Routes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


