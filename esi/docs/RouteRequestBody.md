# RouteRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvoidSystems** | Pointer to **[]int64** | Systems to avoid | [optional] 
**Connections** | Pointer to [**[]RouteConnection**](RouteConnection.md) | Additional one-way connections (like Jump Bridges) between systems | [optional] 
**Preference** | Pointer to **string** | Preference for the route | [optional] [default to "Shorter"]
**SecurityPenalty** | Pointer to **int64** | Strictness of the path preference | [optional] [default to 50]

## Methods

### NewRouteRequestBody

`func NewRouteRequestBody() *RouteRequestBody`

NewRouteRequestBody instantiates a new RouteRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteRequestBodyWithDefaults

`func NewRouteRequestBodyWithDefaults() *RouteRequestBody`

NewRouteRequestBodyWithDefaults instantiates a new RouteRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvoidSystems

`func (o *RouteRequestBody) GetAvoidSystems() []int64`

GetAvoidSystems returns the AvoidSystems field if non-nil, zero value otherwise.

### GetAvoidSystemsOk

`func (o *RouteRequestBody) GetAvoidSystemsOk() (*[]int64, bool)`

GetAvoidSystemsOk returns a tuple with the AvoidSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvoidSystems

`func (o *RouteRequestBody) SetAvoidSystems(v []int64)`

SetAvoidSystems sets AvoidSystems field to given value.

### HasAvoidSystems

`func (o *RouteRequestBody) HasAvoidSystems() bool`

HasAvoidSystems returns a boolean if a field has been set.

### GetConnections

`func (o *RouteRequestBody) GetConnections() []RouteConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *RouteRequestBody) GetConnectionsOk() (*[]RouteConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *RouteRequestBody) SetConnections(v []RouteConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *RouteRequestBody) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetPreference

`func (o *RouteRequestBody) GetPreference() string`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *RouteRequestBody) GetPreferenceOk() (*string, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *RouteRequestBody) SetPreference(v string)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *RouteRequestBody) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetSecurityPenalty

`func (o *RouteRequestBody) GetSecurityPenalty() int64`

GetSecurityPenalty returns the SecurityPenalty field if non-nil, zero value otherwise.

### GetSecurityPenaltyOk

`func (o *RouteRequestBody) GetSecurityPenaltyOk() (*int64, bool)`

GetSecurityPenaltyOk returns a tuple with the SecurityPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPenalty

`func (o *RouteRequestBody) SetSecurityPenalty(v int64)`

SetSecurityPenalty sets SecurityPenalty field to given value.

### HasSecurityPenalty

`func (o *RouteRequestBody) HasSecurityPenalty() bool`

HasSecurityPenalty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


