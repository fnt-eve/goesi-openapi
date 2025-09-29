# UniverseStationsStationIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxDockableShipVolume** | **float64** |  | 
**Name** | **string** |  | 
**OfficeRentalCost** | **float64** |  | 
**Owner** | Pointer to **int64** | ID of the corporation that controls this station | [optional] 
**Position** | [**CharactersCharacterIdAssetsLocationsPostInnerPosition**](CharactersCharacterIdAssetsLocationsPostInnerPosition.md) |  | 
**RaceId** | Pointer to **int64** |  | [optional] 
**ReprocessingEfficiency** | **float64** |  | 
**ReprocessingStationsTake** | **float64** |  | 
**Services** | **[]string** |  | 
**StationId** | **int64** |  | 
**SystemId** | **int64** | The solar system this station is in | 
**TypeId** | **int64** |  | 

## Methods

### NewUniverseStationsStationIdGet

`func NewUniverseStationsStationIdGet(maxDockableShipVolume float64, name string, officeRentalCost float64, position CharactersCharacterIdAssetsLocationsPostInnerPosition, reprocessingEfficiency float64, reprocessingStationsTake float64, services []string, stationId int64, systemId int64, typeId int64, ) *UniverseStationsStationIdGet`

NewUniverseStationsStationIdGet instantiates a new UniverseStationsStationIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseStationsStationIdGetWithDefaults

`func NewUniverseStationsStationIdGetWithDefaults() *UniverseStationsStationIdGet`

NewUniverseStationsStationIdGetWithDefaults instantiates a new UniverseStationsStationIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxDockableShipVolume

`func (o *UniverseStationsStationIdGet) GetMaxDockableShipVolume() float64`

GetMaxDockableShipVolume returns the MaxDockableShipVolume field if non-nil, zero value otherwise.

### GetMaxDockableShipVolumeOk

`func (o *UniverseStationsStationIdGet) GetMaxDockableShipVolumeOk() (*float64, bool)`

GetMaxDockableShipVolumeOk returns a tuple with the MaxDockableShipVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDockableShipVolume

`func (o *UniverseStationsStationIdGet) SetMaxDockableShipVolume(v float64)`

SetMaxDockableShipVolume sets MaxDockableShipVolume field to given value.


### GetName

`func (o *UniverseStationsStationIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniverseStationsStationIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniverseStationsStationIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetOfficeRentalCost

`func (o *UniverseStationsStationIdGet) GetOfficeRentalCost() float64`

GetOfficeRentalCost returns the OfficeRentalCost field if non-nil, zero value otherwise.

### GetOfficeRentalCostOk

`func (o *UniverseStationsStationIdGet) GetOfficeRentalCostOk() (*float64, bool)`

GetOfficeRentalCostOk returns a tuple with the OfficeRentalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeRentalCost

`func (o *UniverseStationsStationIdGet) SetOfficeRentalCost(v float64)`

SetOfficeRentalCost sets OfficeRentalCost field to given value.


### GetOwner

`func (o *UniverseStationsStationIdGet) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UniverseStationsStationIdGet) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UniverseStationsStationIdGet) SetOwner(v int64)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UniverseStationsStationIdGet) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPosition

`func (o *UniverseStationsStationIdGet) GetPosition() CharactersCharacterIdAssetsLocationsPostInnerPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UniverseStationsStationIdGet) GetPositionOk() (*CharactersCharacterIdAssetsLocationsPostInnerPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UniverseStationsStationIdGet) SetPosition(v CharactersCharacterIdAssetsLocationsPostInnerPosition)`

SetPosition sets Position field to given value.


### GetRaceId

`func (o *UniverseStationsStationIdGet) GetRaceId() int64`

GetRaceId returns the RaceId field if non-nil, zero value otherwise.

### GetRaceIdOk

`func (o *UniverseStationsStationIdGet) GetRaceIdOk() (*int64, bool)`

GetRaceIdOk returns a tuple with the RaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaceId

`func (o *UniverseStationsStationIdGet) SetRaceId(v int64)`

SetRaceId sets RaceId field to given value.

### HasRaceId

`func (o *UniverseStationsStationIdGet) HasRaceId() bool`

HasRaceId returns a boolean if a field has been set.

### GetReprocessingEfficiency

`func (o *UniverseStationsStationIdGet) GetReprocessingEfficiency() float64`

GetReprocessingEfficiency returns the ReprocessingEfficiency field if non-nil, zero value otherwise.

### GetReprocessingEfficiencyOk

`func (o *UniverseStationsStationIdGet) GetReprocessingEfficiencyOk() (*float64, bool)`

GetReprocessingEfficiencyOk returns a tuple with the ReprocessingEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReprocessingEfficiency

`func (o *UniverseStationsStationIdGet) SetReprocessingEfficiency(v float64)`

SetReprocessingEfficiency sets ReprocessingEfficiency field to given value.


### GetReprocessingStationsTake

`func (o *UniverseStationsStationIdGet) GetReprocessingStationsTake() float64`

GetReprocessingStationsTake returns the ReprocessingStationsTake field if non-nil, zero value otherwise.

### GetReprocessingStationsTakeOk

`func (o *UniverseStationsStationIdGet) GetReprocessingStationsTakeOk() (*float64, bool)`

GetReprocessingStationsTakeOk returns a tuple with the ReprocessingStationsTake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReprocessingStationsTake

`func (o *UniverseStationsStationIdGet) SetReprocessingStationsTake(v float64)`

SetReprocessingStationsTake sets ReprocessingStationsTake field to given value.


### GetServices

`func (o *UniverseStationsStationIdGet) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *UniverseStationsStationIdGet) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *UniverseStationsStationIdGet) SetServices(v []string)`

SetServices sets Services field to given value.


### GetStationId

`func (o *UniverseStationsStationIdGet) GetStationId() int64`

GetStationId returns the StationId field if non-nil, zero value otherwise.

### GetStationIdOk

`func (o *UniverseStationsStationIdGet) GetStationIdOk() (*int64, bool)`

GetStationIdOk returns a tuple with the StationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationId

`func (o *UniverseStationsStationIdGet) SetStationId(v int64)`

SetStationId sets StationId field to given value.


### GetSystemId

`func (o *UniverseStationsStationIdGet) GetSystemId() int64`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *UniverseStationsStationIdGet) GetSystemIdOk() (*int64, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *UniverseStationsStationIdGet) SetSystemId(v int64)`

SetSystemId sets SystemId field to given value.


### GetTypeId

`func (o *UniverseStationsStationIdGet) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *UniverseStationsStationIdGet) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *UniverseStationsStationIdGet) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


