# CorporationsStructuresSkyhooksDetailReagent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCycle** | **time.Time** | Moment the &#39;SecureStock&#39;/&#39;UnsecuredStock&#39; value had its last cycle; use SDE to calculate the current values | 
**SecuredStock** | **int64** | Secured stock of the reagent at the time of &#39;last_cycle&#39; | 
**TypeId** | **int64** |  | 
**UnsecuredStock** | **int64** | Unsecured stock of the reagent at the time of &#39;last_cycle&#39; | 

## Methods

### NewCorporationsStructuresSkyhooksDetailReagent

`func NewCorporationsStructuresSkyhooksDetailReagent(lastCycle time.Time, securedStock int64, typeId int64, unsecuredStock int64, ) *CorporationsStructuresSkyhooksDetailReagent`

NewCorporationsStructuresSkyhooksDetailReagent instantiates a new CorporationsStructuresSkyhooksDetailReagent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsStructuresSkyhooksDetailReagentWithDefaults

`func NewCorporationsStructuresSkyhooksDetailReagentWithDefaults() *CorporationsStructuresSkyhooksDetailReagent`

NewCorporationsStructuresSkyhooksDetailReagentWithDefaults instantiates a new CorporationsStructuresSkyhooksDetailReagent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCycle

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetLastCycle() time.Time`

GetLastCycle returns the LastCycle field if non-nil, zero value otherwise.

### GetLastCycleOk

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetLastCycleOk() (*time.Time, bool)`

GetLastCycleOk returns a tuple with the LastCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCycle

`func (o *CorporationsStructuresSkyhooksDetailReagent) SetLastCycle(v time.Time)`

SetLastCycle sets LastCycle field to given value.


### GetSecuredStock

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetSecuredStock() int64`

GetSecuredStock returns the SecuredStock field if non-nil, zero value otherwise.

### GetSecuredStockOk

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetSecuredStockOk() (*int64, bool)`

GetSecuredStockOk returns a tuple with the SecuredStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredStock

`func (o *CorporationsStructuresSkyhooksDetailReagent) SetSecuredStock(v int64)`

SetSecuredStock sets SecuredStock field to given value.


### GetTypeId

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsStructuresSkyhooksDetailReagent) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetUnsecuredStock

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetUnsecuredStock() int64`

GetUnsecuredStock returns the UnsecuredStock field if non-nil, zero value otherwise.

### GetUnsecuredStockOk

`func (o *CorporationsStructuresSkyhooksDetailReagent) GetUnsecuredStockOk() (*int64, bool)`

GetUnsecuredStockOk returns a tuple with the UnsecuredStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsecuredStock

`func (o *CorporationsStructuresSkyhooksDetailReagent) SetUnsecuredStock(v int64)`

SetUnsecuredStock sets UnsecuredStock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


