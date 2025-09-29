# ContractsPublicItemsContractIdGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBlueprintCopy** | Pointer to **bool** |  | [optional] 
**IsIncluded** | **bool** | true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract | 
**ItemId** | Pointer to **int64** | Unique ID for the item being sold. Not present if item is being requested by contract rather than sold with contract | [optional] 
**MaterialEfficiency** | Pointer to **int64** | Material Efficiency Level of the blueprint | [optional] 
**Quantity** | **int64** | Number of items in the stack | 
**RecordId** | **int64** | Unique ID for the item, used by the contract system | 
**Runs** | Pointer to **int64** | Number of runs remaining if the blueprint is a copy, -1 if it is an original | [optional] 
**TimeEfficiency** | Pointer to **int64** | Time Efficiency Level of the blueprint | [optional] 
**TypeId** | **int64** | Type ID for item | 

## Methods

### NewContractsPublicItemsContractIdGetInner

`func NewContractsPublicItemsContractIdGetInner(isIncluded bool, quantity int64, recordId int64, typeId int64, ) *ContractsPublicItemsContractIdGetInner`

NewContractsPublicItemsContractIdGetInner instantiates a new ContractsPublicItemsContractIdGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsPublicItemsContractIdGetInnerWithDefaults

`func NewContractsPublicItemsContractIdGetInnerWithDefaults() *ContractsPublicItemsContractIdGetInner`

NewContractsPublicItemsContractIdGetInnerWithDefaults instantiates a new ContractsPublicItemsContractIdGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBlueprintCopy

`func (o *ContractsPublicItemsContractIdGetInner) GetIsBlueprintCopy() bool`

GetIsBlueprintCopy returns the IsBlueprintCopy field if non-nil, zero value otherwise.

### GetIsBlueprintCopyOk

`func (o *ContractsPublicItemsContractIdGetInner) GetIsBlueprintCopyOk() (*bool, bool)`

GetIsBlueprintCopyOk returns a tuple with the IsBlueprintCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlueprintCopy

`func (o *ContractsPublicItemsContractIdGetInner) SetIsBlueprintCopy(v bool)`

SetIsBlueprintCopy sets IsBlueprintCopy field to given value.

### HasIsBlueprintCopy

`func (o *ContractsPublicItemsContractIdGetInner) HasIsBlueprintCopy() bool`

HasIsBlueprintCopy returns a boolean if a field has been set.

### GetIsIncluded

`func (o *ContractsPublicItemsContractIdGetInner) GetIsIncluded() bool`

GetIsIncluded returns the IsIncluded field if non-nil, zero value otherwise.

### GetIsIncludedOk

`func (o *ContractsPublicItemsContractIdGetInner) GetIsIncludedOk() (*bool, bool)`

GetIsIncludedOk returns a tuple with the IsIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncluded

`func (o *ContractsPublicItemsContractIdGetInner) SetIsIncluded(v bool)`

SetIsIncluded sets IsIncluded field to given value.


### GetItemId

`func (o *ContractsPublicItemsContractIdGetInner) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ContractsPublicItemsContractIdGetInner) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ContractsPublicItemsContractIdGetInner) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ContractsPublicItemsContractIdGetInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetMaterialEfficiency

`func (o *ContractsPublicItemsContractIdGetInner) GetMaterialEfficiency() int64`

GetMaterialEfficiency returns the MaterialEfficiency field if non-nil, zero value otherwise.

### GetMaterialEfficiencyOk

`func (o *ContractsPublicItemsContractIdGetInner) GetMaterialEfficiencyOk() (*int64, bool)`

GetMaterialEfficiencyOk returns a tuple with the MaterialEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialEfficiency

`func (o *ContractsPublicItemsContractIdGetInner) SetMaterialEfficiency(v int64)`

SetMaterialEfficiency sets MaterialEfficiency field to given value.

### HasMaterialEfficiency

`func (o *ContractsPublicItemsContractIdGetInner) HasMaterialEfficiency() bool`

HasMaterialEfficiency returns a boolean if a field has been set.

### GetQuantity

`func (o *ContractsPublicItemsContractIdGetInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ContractsPublicItemsContractIdGetInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ContractsPublicItemsContractIdGetInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetRecordId

`func (o *ContractsPublicItemsContractIdGetInner) GetRecordId() int64`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *ContractsPublicItemsContractIdGetInner) GetRecordIdOk() (*int64, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *ContractsPublicItemsContractIdGetInner) SetRecordId(v int64)`

SetRecordId sets RecordId field to given value.


### GetRuns

`func (o *ContractsPublicItemsContractIdGetInner) GetRuns() int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ContractsPublicItemsContractIdGetInner) GetRunsOk() (*int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ContractsPublicItemsContractIdGetInner) SetRuns(v int64)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *ContractsPublicItemsContractIdGetInner) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetTimeEfficiency

`func (o *ContractsPublicItemsContractIdGetInner) GetTimeEfficiency() int64`

GetTimeEfficiency returns the TimeEfficiency field if non-nil, zero value otherwise.

### GetTimeEfficiencyOk

`func (o *ContractsPublicItemsContractIdGetInner) GetTimeEfficiencyOk() (*int64, bool)`

GetTimeEfficiencyOk returns a tuple with the TimeEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEfficiency

`func (o *ContractsPublicItemsContractIdGetInner) SetTimeEfficiency(v int64)`

SetTimeEfficiency sets TimeEfficiency field to given value.

### HasTimeEfficiency

`func (o *ContractsPublicItemsContractIdGetInner) HasTimeEfficiency() bool`

HasTimeEfficiency returns a boolean if a field has been set.

### GetTypeId

`func (o *ContractsPublicItemsContractIdGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ContractsPublicItemsContractIdGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ContractsPublicItemsContractIdGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


