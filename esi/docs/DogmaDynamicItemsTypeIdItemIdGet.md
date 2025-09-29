# DogmaDynamicItemsTypeIdItemIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | **int64** | The ID of the character who created the item | 
**DogmaAttributes** | [**[]DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner**](DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner.md) |  | 
**DogmaEffects** | [**[]DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner**](DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner.md) |  | 
**MutatorTypeId** | **int64** | The type ID of the mutator used to generate the dynamic item. | 
**SourceTypeId** | **int64** | The type ID of the source item the mutator was applied to create the dynamic item. | 

## Methods

### NewDogmaDynamicItemsTypeIdItemIdGet

`func NewDogmaDynamicItemsTypeIdItemIdGet(createdBy int64, dogmaAttributes []DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner, dogmaEffects []DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner, mutatorTypeId int64, sourceTypeId int64, ) *DogmaDynamicItemsTypeIdItemIdGet`

NewDogmaDynamicItemsTypeIdItemIdGet instantiates a new DogmaDynamicItemsTypeIdItemIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDogmaDynamicItemsTypeIdItemIdGetWithDefaults

`func NewDogmaDynamicItemsTypeIdItemIdGetWithDefaults() *DogmaDynamicItemsTypeIdItemIdGet`

NewDogmaDynamicItemsTypeIdItemIdGetWithDefaults instantiates a new DogmaDynamicItemsTypeIdItemIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DogmaDynamicItemsTypeIdItemIdGet) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetDogmaAttributes

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetDogmaAttributes() []DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner`

GetDogmaAttributes returns the DogmaAttributes field if non-nil, zero value otherwise.

### GetDogmaAttributesOk

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetDogmaAttributesOk() (*[]DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner, bool)`

GetDogmaAttributesOk returns a tuple with the DogmaAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDogmaAttributes

`func (o *DogmaDynamicItemsTypeIdItemIdGet) SetDogmaAttributes(v []DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner)`

SetDogmaAttributes sets DogmaAttributes field to given value.


### GetDogmaEffects

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetDogmaEffects() []DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner`

GetDogmaEffects returns the DogmaEffects field if non-nil, zero value otherwise.

### GetDogmaEffectsOk

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetDogmaEffectsOk() (*[]DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner, bool)`

GetDogmaEffectsOk returns a tuple with the DogmaEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDogmaEffects

`func (o *DogmaDynamicItemsTypeIdItemIdGet) SetDogmaEffects(v []DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner)`

SetDogmaEffects sets DogmaEffects field to given value.


### GetMutatorTypeId

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetMutatorTypeId() int64`

GetMutatorTypeId returns the MutatorTypeId field if non-nil, zero value otherwise.

### GetMutatorTypeIdOk

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetMutatorTypeIdOk() (*int64, bool)`

GetMutatorTypeIdOk returns a tuple with the MutatorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutatorTypeId

`func (o *DogmaDynamicItemsTypeIdItemIdGet) SetMutatorTypeId(v int64)`

SetMutatorTypeId sets MutatorTypeId field to given value.


### GetSourceTypeId

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetSourceTypeId() int64`

GetSourceTypeId returns the SourceTypeId field if non-nil, zero value otherwise.

### GetSourceTypeIdOk

`func (o *DogmaDynamicItemsTypeIdItemIdGet) GetSourceTypeIdOk() (*int64, bool)`

GetSourceTypeIdOk returns a tuple with the SourceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypeId

`func (o *DogmaDynamicItemsTypeIdItemIdGet) SetSourceTypeId(v int64)`

SetSourceTypeId sets SourceTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


