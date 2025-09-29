# UniverseStructuresStructureIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The full name of the structure | 
**OwnerId** | **int64** | The ID of the corporation who owns this particular structure | 
**Position** | Pointer to [**UniverseStructuresStructureIdGetPosition**](UniverseStructuresStructureIdGetPosition.md) |  | [optional] 
**SolarSystemId** | **int64** |  | 
**TypeId** | Pointer to **int64** |  | [optional] 

## Methods

### NewUniverseStructuresStructureIdGet

`func NewUniverseStructuresStructureIdGet(name string, ownerId int64, solarSystemId int64, ) *UniverseStructuresStructureIdGet`

NewUniverseStructuresStructureIdGet instantiates a new UniverseStructuresStructureIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseStructuresStructureIdGetWithDefaults

`func NewUniverseStructuresStructureIdGetWithDefaults() *UniverseStructuresStructureIdGet`

NewUniverseStructuresStructureIdGetWithDefaults instantiates a new UniverseStructuresStructureIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UniverseStructuresStructureIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniverseStructuresStructureIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniverseStructuresStructureIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *UniverseStructuresStructureIdGet) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UniverseStructuresStructureIdGet) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UniverseStructuresStructureIdGet) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.


### GetPosition

`func (o *UniverseStructuresStructureIdGet) GetPosition() UniverseStructuresStructureIdGetPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UniverseStructuresStructureIdGet) GetPositionOk() (*UniverseStructuresStructureIdGetPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UniverseStructuresStructureIdGet) SetPosition(v UniverseStructuresStructureIdGetPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UniverseStructuresStructureIdGet) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSolarSystemId

`func (o *UniverseStructuresStructureIdGet) GetSolarSystemId() int64`

GetSolarSystemId returns the SolarSystemId field if non-nil, zero value otherwise.

### GetSolarSystemIdOk

`func (o *UniverseStructuresStructureIdGet) GetSolarSystemIdOk() (*int64, bool)`

GetSolarSystemIdOk returns a tuple with the SolarSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarSystemId

`func (o *UniverseStructuresStructureIdGet) SetSolarSystemId(v int64)`

SetSolarSystemId sets SolarSystemId field to given value.


### GetTypeId

`func (o *UniverseStructuresStructureIdGet) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *UniverseStructuresStructureIdGet) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *UniverseStructuresStructureIdGet) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *UniverseStructuresStructureIdGet) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


