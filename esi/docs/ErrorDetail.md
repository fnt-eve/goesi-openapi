# ErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** | Where the error occurred, e.g. &#39;body.items[3].tags&#39; or &#39;path.thing-id&#39; | [optional] 
**Message** | Pointer to **string** | Error message text | [optional] 
**Value** | Pointer to **interface{}** | The value at the given location | [optional] 

## Methods

### NewErrorDetail

`func NewErrorDetail() *ErrorDetail`

NewErrorDetail instantiates a new ErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailWithDefaults

`func NewErrorDetailWithDefaults() *ErrorDetail`

NewErrorDetailWithDefaults instantiates a new ErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ErrorDetail) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ErrorDetail) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ErrorDetail) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ErrorDetail) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetValue

`func (o *ErrorDetail) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorDetail) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorDetail) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorDetail) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ErrorDetail) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ErrorDetail) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


