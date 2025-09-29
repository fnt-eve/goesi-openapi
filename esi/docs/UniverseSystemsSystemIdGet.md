# UniverseSystemsSystemIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstellationId** | **int64** | The constellation this solar system is in | 
**Name** | **string** |  | 
**Planets** | Pointer to [**[]UniverseSystemsSystemIdGetPlanetsInner**](UniverseSystemsSystemIdGetPlanetsInner.md) |  | [optional] 
**Position** | [**CharactersCharacterIdAssetsLocationsPostInnerPosition**](CharactersCharacterIdAssetsLocationsPostInnerPosition.md) |  | 
**SecurityClass** | Pointer to **string** |  | [optional] 
**SecurityStatus** | **float64** |  | 
**StarId** | Pointer to **int64** |  | [optional] 
**Stargates** | Pointer to **[]int64** |  | [optional] 
**Stations** | Pointer to **[]int64** |  | [optional] 
**SystemId** | **int64** |  | 

## Methods

### NewUniverseSystemsSystemIdGet

`func NewUniverseSystemsSystemIdGet(constellationId int64, name string, position CharactersCharacterIdAssetsLocationsPostInnerPosition, securityStatus float64, systemId int64, ) *UniverseSystemsSystemIdGet`

NewUniverseSystemsSystemIdGet instantiates a new UniverseSystemsSystemIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseSystemsSystemIdGetWithDefaults

`func NewUniverseSystemsSystemIdGetWithDefaults() *UniverseSystemsSystemIdGet`

NewUniverseSystemsSystemIdGetWithDefaults instantiates a new UniverseSystemsSystemIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstellationId

`func (o *UniverseSystemsSystemIdGet) GetConstellationId() int64`

GetConstellationId returns the ConstellationId field if non-nil, zero value otherwise.

### GetConstellationIdOk

`func (o *UniverseSystemsSystemIdGet) GetConstellationIdOk() (*int64, bool)`

GetConstellationIdOk returns a tuple with the ConstellationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstellationId

`func (o *UniverseSystemsSystemIdGet) SetConstellationId(v int64)`

SetConstellationId sets ConstellationId field to given value.


### GetName

`func (o *UniverseSystemsSystemIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniverseSystemsSystemIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniverseSystemsSystemIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetPlanets

`func (o *UniverseSystemsSystemIdGet) GetPlanets() []UniverseSystemsSystemIdGetPlanetsInner`

GetPlanets returns the Planets field if non-nil, zero value otherwise.

### GetPlanetsOk

`func (o *UniverseSystemsSystemIdGet) GetPlanetsOk() (*[]UniverseSystemsSystemIdGetPlanetsInner, bool)`

GetPlanetsOk returns a tuple with the Planets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanets

`func (o *UniverseSystemsSystemIdGet) SetPlanets(v []UniverseSystemsSystemIdGetPlanetsInner)`

SetPlanets sets Planets field to given value.

### HasPlanets

`func (o *UniverseSystemsSystemIdGet) HasPlanets() bool`

HasPlanets returns a boolean if a field has been set.

### GetPosition

`func (o *UniverseSystemsSystemIdGet) GetPosition() CharactersCharacterIdAssetsLocationsPostInnerPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UniverseSystemsSystemIdGet) GetPositionOk() (*CharactersCharacterIdAssetsLocationsPostInnerPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UniverseSystemsSystemIdGet) SetPosition(v CharactersCharacterIdAssetsLocationsPostInnerPosition)`

SetPosition sets Position field to given value.


### GetSecurityClass

`func (o *UniverseSystemsSystemIdGet) GetSecurityClass() string`

GetSecurityClass returns the SecurityClass field if non-nil, zero value otherwise.

### GetSecurityClassOk

`func (o *UniverseSystemsSystemIdGet) GetSecurityClassOk() (*string, bool)`

GetSecurityClassOk returns a tuple with the SecurityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityClass

`func (o *UniverseSystemsSystemIdGet) SetSecurityClass(v string)`

SetSecurityClass sets SecurityClass field to given value.

### HasSecurityClass

`func (o *UniverseSystemsSystemIdGet) HasSecurityClass() bool`

HasSecurityClass returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *UniverseSystemsSystemIdGet) GetSecurityStatus() float64`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *UniverseSystemsSystemIdGet) GetSecurityStatusOk() (*float64, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *UniverseSystemsSystemIdGet) SetSecurityStatus(v float64)`

SetSecurityStatus sets SecurityStatus field to given value.


### GetStarId

`func (o *UniverseSystemsSystemIdGet) GetStarId() int64`

GetStarId returns the StarId field if non-nil, zero value otherwise.

### GetStarIdOk

`func (o *UniverseSystemsSystemIdGet) GetStarIdOk() (*int64, bool)`

GetStarIdOk returns a tuple with the StarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarId

`func (o *UniverseSystemsSystemIdGet) SetStarId(v int64)`

SetStarId sets StarId field to given value.

### HasStarId

`func (o *UniverseSystemsSystemIdGet) HasStarId() bool`

HasStarId returns a boolean if a field has been set.

### GetStargates

`func (o *UniverseSystemsSystemIdGet) GetStargates() []int64`

GetStargates returns the Stargates field if non-nil, zero value otherwise.

### GetStargatesOk

`func (o *UniverseSystemsSystemIdGet) GetStargatesOk() (*[]int64, bool)`

GetStargatesOk returns a tuple with the Stargates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargates

`func (o *UniverseSystemsSystemIdGet) SetStargates(v []int64)`

SetStargates sets Stargates field to given value.

### HasStargates

`func (o *UniverseSystemsSystemIdGet) HasStargates() bool`

HasStargates returns a boolean if a field has been set.

### GetStations

`func (o *UniverseSystemsSystemIdGet) GetStations() []int64`

GetStations returns the Stations field if non-nil, zero value otherwise.

### GetStationsOk

`func (o *UniverseSystemsSystemIdGet) GetStationsOk() (*[]int64, bool)`

GetStationsOk returns a tuple with the Stations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStations

`func (o *UniverseSystemsSystemIdGet) SetStations(v []int64)`

SetStations sets Stations field to given value.

### HasStations

`func (o *UniverseSystemsSystemIdGet) HasStations() bool`

HasStations returns a boolean if a field has been set.

### GetSystemId

`func (o *UniverseSystemsSystemIdGet) GetSystemId() int64`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *UniverseSystemsSystemIdGet) GetSystemIdOk() (*int64, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *UniverseSystemsSystemIdGet) SetSystemId(v int64)`

SetSystemId sets SystemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


