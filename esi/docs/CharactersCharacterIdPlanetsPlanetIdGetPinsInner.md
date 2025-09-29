# CharactersCharacterIdPlanetsPlanetIdGetPinsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]CharactersCharacterIdPlanetsPlanetIdGetPinsInnerContentsInner**](CharactersCharacterIdPlanetsPlanetIdGetPinsInnerContentsInner.md) |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**ExtractorDetails** | Pointer to [**CharactersCharacterIdPlanetsPlanetIdGetPinsInnerExtractorDetails**](CharactersCharacterIdPlanetsPlanetIdGetPinsInnerExtractorDetails.md) |  | [optional] 
**FactoryDetails** | Pointer to [**CharactersCharacterIdPlanetsPlanetIdGetPinsInnerFactoryDetails**](CharactersCharacterIdPlanetsPlanetIdGetPinsInnerFactoryDetails.md) |  | [optional] 
**InstallTime** | Pointer to **time.Time** |  | [optional] 
**LastCycleStart** | Pointer to **time.Time** |  | [optional] 
**Latitude** | **float64** |  | 
**Longitude** | **float64** |  | 
**PinId** | **int64** |  | 
**SchematicId** | Pointer to **int64** |  | [optional] 
**TypeId** | **int64** |  | 

## Methods

### NewCharactersCharacterIdPlanetsPlanetIdGetPinsInner

`func NewCharactersCharacterIdPlanetsPlanetIdGetPinsInner(latitude float64, longitude float64, pinId int64, typeId int64, ) *CharactersCharacterIdPlanetsPlanetIdGetPinsInner`

NewCharactersCharacterIdPlanetsPlanetIdGetPinsInner instantiates a new CharactersCharacterIdPlanetsPlanetIdGetPinsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdPlanetsPlanetIdGetPinsInnerWithDefaults

`func NewCharactersCharacterIdPlanetsPlanetIdGetPinsInnerWithDefaults() *CharactersCharacterIdPlanetsPlanetIdGetPinsInner`

NewCharactersCharacterIdPlanetsPlanetIdGetPinsInnerWithDefaults instantiates a new CharactersCharacterIdPlanetsPlanetIdGetPinsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetContents() []CharactersCharacterIdPlanetsPlanetIdGetPinsInnerContentsInner`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetContentsOk() (*[]CharactersCharacterIdPlanetsPlanetIdGetPinsInnerContentsInner, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetContents(v []CharactersCharacterIdPlanetsPlanetIdGetPinsInnerContentsInner)`

SetContents sets Contents field to given value.

### HasContents

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetExpiryTime

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetExtractorDetails

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetExtractorDetails() CharactersCharacterIdPlanetsPlanetIdGetPinsInnerExtractorDetails`

GetExtractorDetails returns the ExtractorDetails field if non-nil, zero value otherwise.

### GetExtractorDetailsOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetExtractorDetailsOk() (*CharactersCharacterIdPlanetsPlanetIdGetPinsInnerExtractorDetails, bool)`

GetExtractorDetailsOk returns a tuple with the ExtractorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractorDetails

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetExtractorDetails(v CharactersCharacterIdPlanetsPlanetIdGetPinsInnerExtractorDetails)`

SetExtractorDetails sets ExtractorDetails field to given value.

### HasExtractorDetails

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) HasExtractorDetails() bool`

HasExtractorDetails returns a boolean if a field has been set.

### GetFactoryDetails

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetFactoryDetails() CharactersCharacterIdPlanetsPlanetIdGetPinsInnerFactoryDetails`

GetFactoryDetails returns the FactoryDetails field if non-nil, zero value otherwise.

### GetFactoryDetailsOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetFactoryDetailsOk() (*CharactersCharacterIdPlanetsPlanetIdGetPinsInnerFactoryDetails, bool)`

GetFactoryDetailsOk returns a tuple with the FactoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryDetails

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetFactoryDetails(v CharactersCharacterIdPlanetsPlanetIdGetPinsInnerFactoryDetails)`

SetFactoryDetails sets FactoryDetails field to given value.

### HasFactoryDetails

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) HasFactoryDetails() bool`

HasFactoryDetails returns a boolean if a field has been set.

### GetInstallTime

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetInstallTime() time.Time`

GetInstallTime returns the InstallTime field if non-nil, zero value otherwise.

### GetInstallTimeOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetInstallTimeOk() (*time.Time, bool)`

GetInstallTimeOk returns a tuple with the InstallTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTime

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetInstallTime(v time.Time)`

SetInstallTime sets InstallTime field to given value.

### HasInstallTime

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) HasInstallTime() bool`

HasInstallTime returns a boolean if a field has been set.

### GetLastCycleStart

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetLastCycleStart() time.Time`

GetLastCycleStart returns the LastCycleStart field if non-nil, zero value otherwise.

### GetLastCycleStartOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetLastCycleStartOk() (*time.Time, bool)`

GetLastCycleStartOk returns a tuple with the LastCycleStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCycleStart

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetLastCycleStart(v time.Time)`

SetLastCycleStart sets LastCycleStart field to given value.

### HasLastCycleStart

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) HasLastCycleStart() bool`

HasLastCycleStart returns a boolean if a field has been set.

### GetLatitude

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetPinId

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetPinId() int64`

GetPinId returns the PinId field if non-nil, zero value otherwise.

### GetPinIdOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetPinIdOk() (*int64, bool)`

GetPinIdOk returns a tuple with the PinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinId

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetPinId(v int64)`

SetPinId sets PinId field to given value.


### GetSchematicId

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetSchematicId() int64`

GetSchematicId returns the SchematicId field if non-nil, zero value otherwise.

### GetSchematicIdOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetSchematicIdOk() (*int64, bool)`

GetSchematicIdOk returns a tuple with the SchematicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchematicId

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetSchematicId(v int64)`

SetSchematicId sets SchematicId field to given value.

### HasSchematicId

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) HasSchematicId() bool`

HasSchematicId returns a boolean if a field has been set.

### GetTypeId

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CharactersCharacterIdPlanetsPlanetIdGetPinsInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


