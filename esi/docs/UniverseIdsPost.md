# UniverseIdsPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]UniverseIdsPostAgentsInner**](UniverseIdsPostAgentsInner.md) |  | [optional] 
**Alliances** | Pointer to [**[]UniverseIdsPostAlliancesInner**](UniverseIdsPostAlliancesInner.md) |  | [optional] 
**Characters** | Pointer to [**[]UniverseIdsPostCharactersInner**](UniverseIdsPostCharactersInner.md) |  | [optional] 
**Constellations** | Pointer to [**[]UniverseIdsPostConstellationsInner**](UniverseIdsPostConstellationsInner.md) |  | [optional] 
**Corporations** | Pointer to [**[]UniverseIdsPostCorporationsInner**](UniverseIdsPostCorporationsInner.md) |  | [optional] 
**Factions** | Pointer to [**[]UniverseIdsPostFactionsInner**](UniverseIdsPostFactionsInner.md) |  | [optional] 
**InventoryTypes** | Pointer to [**[]UniverseIdsPostInventoryTypesInner**](UniverseIdsPostInventoryTypesInner.md) |  | [optional] 
**Regions** | Pointer to [**[]UniverseIdsPostRegionsInner**](UniverseIdsPostRegionsInner.md) |  | [optional] 
**Stations** | Pointer to [**[]UniverseIdsPostStationsInner**](UniverseIdsPostStationsInner.md) |  | [optional] 
**Systems** | Pointer to [**[]UniverseIdsPostSystemsInner**](UniverseIdsPostSystemsInner.md) |  | [optional] 

## Methods

### NewUniverseIdsPost

`func NewUniverseIdsPost() *UniverseIdsPost`

NewUniverseIdsPost instantiates a new UniverseIdsPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseIdsPostWithDefaults

`func NewUniverseIdsPostWithDefaults() *UniverseIdsPost`

NewUniverseIdsPostWithDefaults instantiates a new UniverseIdsPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *UniverseIdsPost) GetAgents() []UniverseIdsPostAgentsInner`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *UniverseIdsPost) GetAgentsOk() (*[]UniverseIdsPostAgentsInner, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *UniverseIdsPost) SetAgents(v []UniverseIdsPostAgentsInner)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *UniverseIdsPost) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetAlliances

`func (o *UniverseIdsPost) GetAlliances() []UniverseIdsPostAlliancesInner`

GetAlliances returns the Alliances field if non-nil, zero value otherwise.

### GetAlliancesOk

`func (o *UniverseIdsPost) GetAlliancesOk() (*[]UniverseIdsPostAlliancesInner, bool)`

GetAlliancesOk returns a tuple with the Alliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlliances

`func (o *UniverseIdsPost) SetAlliances(v []UniverseIdsPostAlliancesInner)`

SetAlliances sets Alliances field to given value.

### HasAlliances

`func (o *UniverseIdsPost) HasAlliances() bool`

HasAlliances returns a boolean if a field has been set.

### GetCharacters

`func (o *UniverseIdsPost) GetCharacters() []UniverseIdsPostCharactersInner`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *UniverseIdsPost) GetCharactersOk() (*[]UniverseIdsPostCharactersInner, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *UniverseIdsPost) SetCharacters(v []UniverseIdsPostCharactersInner)`

SetCharacters sets Characters field to given value.

### HasCharacters

`func (o *UniverseIdsPost) HasCharacters() bool`

HasCharacters returns a boolean if a field has been set.

### GetConstellations

`func (o *UniverseIdsPost) GetConstellations() []UniverseIdsPostConstellationsInner`

GetConstellations returns the Constellations field if non-nil, zero value otherwise.

### GetConstellationsOk

`func (o *UniverseIdsPost) GetConstellationsOk() (*[]UniverseIdsPostConstellationsInner, bool)`

GetConstellationsOk returns a tuple with the Constellations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstellations

`func (o *UniverseIdsPost) SetConstellations(v []UniverseIdsPostConstellationsInner)`

SetConstellations sets Constellations field to given value.

### HasConstellations

`func (o *UniverseIdsPost) HasConstellations() bool`

HasConstellations returns a boolean if a field has been set.

### GetCorporations

`func (o *UniverseIdsPost) GetCorporations() []UniverseIdsPostCorporationsInner`

GetCorporations returns the Corporations field if non-nil, zero value otherwise.

### GetCorporationsOk

`func (o *UniverseIdsPost) GetCorporationsOk() (*[]UniverseIdsPostCorporationsInner, bool)`

GetCorporationsOk returns a tuple with the Corporations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporations

`func (o *UniverseIdsPost) SetCorporations(v []UniverseIdsPostCorporationsInner)`

SetCorporations sets Corporations field to given value.

### HasCorporations

`func (o *UniverseIdsPost) HasCorporations() bool`

HasCorporations returns a boolean if a field has been set.

### GetFactions

`func (o *UniverseIdsPost) GetFactions() []UniverseIdsPostFactionsInner`

GetFactions returns the Factions field if non-nil, zero value otherwise.

### GetFactionsOk

`func (o *UniverseIdsPost) GetFactionsOk() (*[]UniverseIdsPostFactionsInner, bool)`

GetFactionsOk returns a tuple with the Factions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactions

`func (o *UniverseIdsPost) SetFactions(v []UniverseIdsPostFactionsInner)`

SetFactions sets Factions field to given value.

### HasFactions

`func (o *UniverseIdsPost) HasFactions() bool`

HasFactions returns a boolean if a field has been set.

### GetInventoryTypes

`func (o *UniverseIdsPost) GetInventoryTypes() []UniverseIdsPostInventoryTypesInner`

GetInventoryTypes returns the InventoryTypes field if non-nil, zero value otherwise.

### GetInventoryTypesOk

`func (o *UniverseIdsPost) GetInventoryTypesOk() (*[]UniverseIdsPostInventoryTypesInner, bool)`

GetInventoryTypesOk returns a tuple with the InventoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryTypes

`func (o *UniverseIdsPost) SetInventoryTypes(v []UniverseIdsPostInventoryTypesInner)`

SetInventoryTypes sets InventoryTypes field to given value.

### HasInventoryTypes

`func (o *UniverseIdsPost) HasInventoryTypes() bool`

HasInventoryTypes returns a boolean if a field has been set.

### GetRegions

`func (o *UniverseIdsPost) GetRegions() []UniverseIdsPostRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *UniverseIdsPost) GetRegionsOk() (*[]UniverseIdsPostRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *UniverseIdsPost) SetRegions(v []UniverseIdsPostRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *UniverseIdsPost) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetStations

`func (o *UniverseIdsPost) GetStations() []UniverseIdsPostStationsInner`

GetStations returns the Stations field if non-nil, zero value otherwise.

### GetStationsOk

`func (o *UniverseIdsPost) GetStationsOk() (*[]UniverseIdsPostStationsInner, bool)`

GetStationsOk returns a tuple with the Stations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStations

`func (o *UniverseIdsPost) SetStations(v []UniverseIdsPostStationsInner)`

SetStations sets Stations field to given value.

### HasStations

`func (o *UniverseIdsPost) HasStations() bool`

HasStations returns a boolean if a field has been set.

### GetSystems

`func (o *UniverseIdsPost) GetSystems() []UniverseIdsPostSystemsInner`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *UniverseIdsPost) GetSystemsOk() (*[]UniverseIdsPostSystemsInner, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *UniverseIdsPost) SetSystems(v []UniverseIdsPostSystemsInner)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *UniverseIdsPost) HasSystems() bool`

HasSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


