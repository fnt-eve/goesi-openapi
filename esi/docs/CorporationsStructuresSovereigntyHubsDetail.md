# CorporationsStructuresSovereigntyHubsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuelAccessListId** | Pointer to **int64** |  | [optional] 
**Id** | **int64** |  | 
**ReagentBay** | [**CorporationsStructuresSovereigntyHubsDetailReagentbay**](CorporationsStructuresSovereigntyHubsDetailReagentbay.md) |  | 
**Resources** | [**CorporationsStructuresSovereigntyHubsDetailResources**](CorporationsStructuresSovereigntyHubsDetailResources.md) |  | 
**SolarSystemId** | **int64** |  | 
**Upgrades** | [**[]CorporationsStructuresSovereigntyHubsDetailUpgrade**](CorporationsStructuresSovereigntyHubsDetailUpgrade.md) | Sovereignty Hub&#39;s installed upgrades | 
**VulnerabilityWindow** | Pointer to [**CorporationsStructuresSovereigntyHubsDetailVulnerabilitywindow**](CorporationsStructuresSovereigntyHubsDetailVulnerabilitywindow.md) |  | [optional] 
**WorkforceTransport** | [**CorporationsStructuresSovereigntyHubsDetailTransport**](CorporationsStructuresSovereigntyHubsDetailTransport.md) |  | 

## Methods

### NewCorporationsStructuresSovereigntyHubsDetail

`func NewCorporationsStructuresSovereigntyHubsDetail(id int64, reagentBay CorporationsStructuresSovereigntyHubsDetailReagentbay, resources CorporationsStructuresSovereigntyHubsDetailResources, solarSystemId int64, upgrades []CorporationsStructuresSovereigntyHubsDetailUpgrade, workforceTransport CorporationsStructuresSovereigntyHubsDetailTransport, ) *CorporationsStructuresSovereigntyHubsDetail`

NewCorporationsStructuresSovereigntyHubsDetail instantiates a new CorporationsStructuresSovereigntyHubsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsStructuresSovereigntyHubsDetailWithDefaults

`func NewCorporationsStructuresSovereigntyHubsDetailWithDefaults() *CorporationsStructuresSovereigntyHubsDetail`

NewCorporationsStructuresSovereigntyHubsDetailWithDefaults instantiates a new CorporationsStructuresSovereigntyHubsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuelAccessListId

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetFuelAccessListId() int64`

GetFuelAccessListId returns the FuelAccessListId field if non-nil, zero value otherwise.

### GetFuelAccessListIdOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetFuelAccessListIdOk() (*int64, bool)`

GetFuelAccessListIdOk returns a tuple with the FuelAccessListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelAccessListId

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetFuelAccessListId(v int64)`

SetFuelAccessListId sets FuelAccessListId field to given value.

### HasFuelAccessListId

`func (o *CorporationsStructuresSovereigntyHubsDetail) HasFuelAccessListId() bool`

HasFuelAccessListId returns a boolean if a field has been set.

### GetId

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetId(v int64)`

SetId sets Id field to given value.


### GetReagentBay

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetReagentBay() CorporationsStructuresSovereigntyHubsDetailReagentbay`

GetReagentBay returns the ReagentBay field if non-nil, zero value otherwise.

### GetReagentBayOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetReagentBayOk() (*CorporationsStructuresSovereigntyHubsDetailReagentbay, bool)`

GetReagentBayOk returns a tuple with the ReagentBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReagentBay

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetReagentBay(v CorporationsStructuresSovereigntyHubsDetailReagentbay)`

SetReagentBay sets ReagentBay field to given value.


### GetResources

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetResources() CorporationsStructuresSovereigntyHubsDetailResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetResourcesOk() (*CorporationsStructuresSovereigntyHubsDetailResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetResources(v CorporationsStructuresSovereigntyHubsDetailResources)`

SetResources sets Resources field to given value.


### GetSolarSystemId

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetSolarSystemId() int64`

GetSolarSystemId returns the SolarSystemId field if non-nil, zero value otherwise.

### GetSolarSystemIdOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetSolarSystemIdOk() (*int64, bool)`

GetSolarSystemIdOk returns a tuple with the SolarSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarSystemId

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetSolarSystemId(v int64)`

SetSolarSystemId sets SolarSystemId field to given value.


### GetUpgrades

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetUpgrades() []CorporationsStructuresSovereigntyHubsDetailUpgrade`

GetUpgrades returns the Upgrades field if non-nil, zero value otherwise.

### GetUpgradesOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetUpgradesOk() (*[]CorporationsStructuresSovereigntyHubsDetailUpgrade, bool)`

GetUpgradesOk returns a tuple with the Upgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrades

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetUpgrades(v []CorporationsStructuresSovereigntyHubsDetailUpgrade)`

SetUpgrades sets Upgrades field to given value.


### GetVulnerabilityWindow

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetVulnerabilityWindow() CorporationsStructuresSovereigntyHubsDetailVulnerabilitywindow`

GetVulnerabilityWindow returns the VulnerabilityWindow field if non-nil, zero value otherwise.

### GetVulnerabilityWindowOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetVulnerabilityWindowOk() (*CorporationsStructuresSovereigntyHubsDetailVulnerabilitywindow, bool)`

GetVulnerabilityWindowOk returns a tuple with the VulnerabilityWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityWindow

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetVulnerabilityWindow(v CorporationsStructuresSovereigntyHubsDetailVulnerabilitywindow)`

SetVulnerabilityWindow sets VulnerabilityWindow field to given value.

### HasVulnerabilityWindow

`func (o *CorporationsStructuresSovereigntyHubsDetail) HasVulnerabilityWindow() bool`

HasVulnerabilityWindow returns a boolean if a field has been set.

### GetWorkforceTransport

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetWorkforceTransport() CorporationsStructuresSovereigntyHubsDetailTransport`

GetWorkforceTransport returns the WorkforceTransport field if non-nil, zero value otherwise.

### GetWorkforceTransportOk

`func (o *CorporationsStructuresSovereigntyHubsDetail) GetWorkforceTransportOk() (*CorporationsStructuresSovereigntyHubsDetailTransport, bool)`

GetWorkforceTransportOk returns a tuple with the WorkforceTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkforceTransport

`func (o *CorporationsStructuresSovereigntyHubsDetail) SetWorkforceTransport(v CorporationsStructuresSovereigntyHubsDetailTransport)`

SetWorkforceTransport sets WorkforceTransport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


