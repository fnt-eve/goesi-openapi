# InsurancePricesGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Levels** | [**[]InsurancePricesGetInnerLevelsInner**](InsurancePricesGetInnerLevelsInner.md) | A list of a available insurance levels for this ship type | 
**TypeId** | **int64** |  | 

## Methods

### NewInsurancePricesGetInner

`func NewInsurancePricesGetInner(levels []InsurancePricesGetInnerLevelsInner, typeId int64, ) *InsurancePricesGetInner`

NewInsurancePricesGetInner instantiates a new InsurancePricesGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsurancePricesGetInnerWithDefaults

`func NewInsurancePricesGetInnerWithDefaults() *InsurancePricesGetInner`

NewInsurancePricesGetInnerWithDefaults instantiates a new InsurancePricesGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevels

`func (o *InsurancePricesGetInner) GetLevels() []InsurancePricesGetInnerLevelsInner`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *InsurancePricesGetInner) GetLevelsOk() (*[]InsurancePricesGetInnerLevelsInner, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *InsurancePricesGetInner) SetLevels(v []InsurancePricesGetInnerLevelsInner)`

SetLevels sets Levels field to given value.


### GetTypeId

`func (o *InsurancePricesGetInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *InsurancePricesGetInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *InsurancePricesGetInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


