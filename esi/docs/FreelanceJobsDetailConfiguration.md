# FreelanceJobsDetailConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | Contribution method (see SDE for valid values) | 
**Parameters** | [**map[string]FreelanceJobsDetailConfigurationParametersValue**](FreelanceJobsDetailConfigurationParametersValue.md) | Parameter values (see SDE for parameter definition) | 
**Version** | **int64** | Version of parameter definition used | 

## Methods

### NewFreelanceJobsDetailConfiguration

`func NewFreelanceJobsDetailConfiguration(method string, parameters map[string]FreelanceJobsDetailConfigurationParametersValue, version int64, ) *FreelanceJobsDetailConfiguration`

NewFreelanceJobsDetailConfiguration instantiates a new FreelanceJobsDetailConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreelanceJobsDetailConfigurationWithDefaults

`func NewFreelanceJobsDetailConfigurationWithDefaults() *FreelanceJobsDetailConfiguration`

NewFreelanceJobsDetailConfigurationWithDefaults instantiates a new FreelanceJobsDetailConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *FreelanceJobsDetailConfiguration) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FreelanceJobsDetailConfiguration) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FreelanceJobsDetailConfiguration) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParameters

`func (o *FreelanceJobsDetailConfiguration) GetParameters() map[string]FreelanceJobsDetailConfigurationParametersValue`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FreelanceJobsDetailConfiguration) GetParametersOk() (*map[string]FreelanceJobsDetailConfigurationParametersValue, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FreelanceJobsDetailConfiguration) SetParameters(v map[string]FreelanceJobsDetailConfigurationParametersValue)`

SetParameters sets Parameters field to given value.


### GetVersion

`func (o *FreelanceJobsDetailConfiguration) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FreelanceJobsDetailConfiguration) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FreelanceJobsDetailConfiguration) SetVersion(v int64)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


