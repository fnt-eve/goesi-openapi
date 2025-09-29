# CorporationsCorporationIdBlueprintsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **int64** | Unique ID for this item. | 
**LocationFlag** | **string** | Type of the location_id | 
**LocationId** | **int64** | References a station, a ship or an item_id if this blueprint is located within a container. | 
**MaterialEfficiency** | **int64** | Material Efficiency Level of the blueprint. | 
**Quantity** | **int64** | A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet). | 
**Runs** | **int64** | Number of runs remaining if the blueprint is a copy, -1 if it is an original. | 
**TimeEfficiency** | **int64** | Time Efficiency Level of the blueprint. | 
**TypeId** | **int64** |  | 

## Methods

### NewCorporationsCorporationIdBlueprintsGetInner

`func NewCorporationsCorporationIdBlueprintsGetInner(itemId int64, locationFlag string, locationId int64, materialEfficiency int64, quantity int64, runs int64, timeEfficiency int64, typeId int64, ) *CorporationsCorporationIdBlueprintsGetInner`

NewCorporationsCorporationIdBlueprintsGetInner instantiates a new CorporationsCorporationIdBlueprintsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdBlueprintsGetInnerWithDefaults

`func NewCorporationsCorporationIdBlueprintsGetInnerWithDefaults() *CorporationsCorporationIdBlueprintsGetInner`

NewCorporationsCorporationIdBlueprintsGetInnerWithDefaults instantiates a new CorporationsCorporationIdBlueprintsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.


### GetLocationFlag

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetLocationFlag() string`

GetLocationFlag returns the LocationFlag field if non-nil, zero value otherwise.

### GetLocationFlagOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetLocationFlagOk() (*string, bool)`

GetLocationFlagOk returns a tuple with the LocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationFlag

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetLocationFlag(v string)`

SetLocationFlag sets LocationFlag field to given value.


### GetLocationId

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetMaterialEfficiency

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetMaterialEfficiency() int64`

GetMaterialEfficiency returns the MaterialEfficiency field if non-nil, zero value otherwise.

### GetMaterialEfficiencyOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetMaterialEfficiencyOk() (*int64, bool)`

GetMaterialEfficiencyOk returns a tuple with the MaterialEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialEfficiency

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetMaterialEfficiency(v int64)`

SetMaterialEfficiency sets MaterialEfficiency field to given value.


### GetQuantity

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetRuns

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetRuns() int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetRunsOk() (*int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetRuns(v int64)`

SetRuns sets Runs field to given value.


### GetTimeEfficiency

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetTimeEfficiency() int64`

GetTimeEfficiency returns the TimeEfficiency field if non-nil, zero value otherwise.

### GetTimeEfficiencyOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetTimeEfficiencyOk() (*int64, bool)`

GetTimeEfficiencyOk returns a tuple with the TimeEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEfficiency

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetTimeEfficiency(v int64)`

SetTimeEfficiency sets TimeEfficiency field to given value.


### GetTypeId

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CorporationsCorporationIdBlueprintsGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CorporationsCorporationIdBlueprintsGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


