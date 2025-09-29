# CorporationsCorporationIdStarbasesStarbaseIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAllianceMembers** | **bool** |  | 
**AllowCorporationMembers** | **bool** |  | 
**Anchor** | **string** | Who can anchor starbase (POS) and its structures | 
**AttackIfAtWar** | **bool** |  | 
**AttackIfOtherSecurityStatusDropping** | **bool** |  | 
**AttackSecurityStatusThreshold** | Pointer to **float64** | Starbase (POS) will attack if target&#39;s security standing is lower than this value | [optional] 
**AttackStandingThreshold** | Pointer to **float64** | Starbase (POS) will attack if target&#39;s standing is lower than this value | [optional] 
**FuelBayTake** | **string** | Who can take fuel blocks out of the starbase (POS)&#39;s fuel bay | 
**FuelBayView** | **string** | Who can view the starbase (POS)&#39;s fule bay. Characters either need to have required role or belong to the starbase (POS) owner&#39;s corporation or alliance, as described by the enum, all other access settings follows the same scheme | 
**Fuels** | Pointer to [**[]CorporationsCorporationIdStarbasesStarbaseIdGetFuelsInner**](CorporationsCorporationIdStarbasesStarbaseIdGetFuelsInner.md) | Fuel blocks and other things that will be consumed when operating a starbase (POS) | [optional] 
**Offline** | **string** | Who can offline starbase (POS) and its structures | 
**Online** | **string** | Who can online starbase (POS) and its structures | 
**Unanchor** | **string** | Who can unanchor starbase (POS) and its structures | 
**UseAllianceStandings** | **bool** | True if the starbase (POS) is using alliance standings, otherwise using corporation&#39;s | 

## Methods

### NewCorporationsCorporationIdStarbasesStarbaseIdGet

`func NewCorporationsCorporationIdStarbasesStarbaseIdGet(allowAllianceMembers bool, allowCorporationMembers bool, anchor string, attackIfAtWar bool, attackIfOtherSecurityStatusDropping bool, fuelBayTake string, fuelBayView string, offline string, online string, unanchor string, useAllianceStandings bool, ) *CorporationsCorporationIdStarbasesStarbaseIdGet`

NewCorporationsCorporationIdStarbasesStarbaseIdGet instantiates a new CorporationsCorporationIdStarbasesStarbaseIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdStarbasesStarbaseIdGetWithDefaults

`func NewCorporationsCorporationIdStarbasesStarbaseIdGetWithDefaults() *CorporationsCorporationIdStarbasesStarbaseIdGet`

NewCorporationsCorporationIdStarbasesStarbaseIdGetWithDefaults instantiates a new CorporationsCorporationIdStarbasesStarbaseIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAllianceMembers

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAllowAllianceMembers() bool`

GetAllowAllianceMembers returns the AllowAllianceMembers field if non-nil, zero value otherwise.

### GetAllowAllianceMembersOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAllowAllianceMembersOk() (*bool, bool)`

GetAllowAllianceMembersOk returns a tuple with the AllowAllianceMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllianceMembers

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetAllowAllianceMembers(v bool)`

SetAllowAllianceMembers sets AllowAllianceMembers field to given value.


### GetAllowCorporationMembers

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAllowCorporationMembers() bool`

GetAllowCorporationMembers returns the AllowCorporationMembers field if non-nil, zero value otherwise.

### GetAllowCorporationMembersOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAllowCorporationMembersOk() (*bool, bool)`

GetAllowCorporationMembersOk returns a tuple with the AllowCorporationMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCorporationMembers

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetAllowCorporationMembers(v bool)`

SetAllowCorporationMembers sets AllowCorporationMembers field to given value.


### GetAnchor

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAnchor() string`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAnchorOk() (*string, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetAnchor(v string)`

SetAnchor sets Anchor field to given value.


### GetAttackIfAtWar

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackIfAtWar() bool`

GetAttackIfAtWar returns the AttackIfAtWar field if non-nil, zero value otherwise.

### GetAttackIfAtWarOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackIfAtWarOk() (*bool, bool)`

GetAttackIfAtWarOk returns a tuple with the AttackIfAtWar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackIfAtWar

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetAttackIfAtWar(v bool)`

SetAttackIfAtWar sets AttackIfAtWar field to given value.


### GetAttackIfOtherSecurityStatusDropping

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackIfOtherSecurityStatusDropping() bool`

GetAttackIfOtherSecurityStatusDropping returns the AttackIfOtherSecurityStatusDropping field if non-nil, zero value otherwise.

### GetAttackIfOtherSecurityStatusDroppingOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackIfOtherSecurityStatusDroppingOk() (*bool, bool)`

GetAttackIfOtherSecurityStatusDroppingOk returns a tuple with the AttackIfOtherSecurityStatusDropping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackIfOtherSecurityStatusDropping

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetAttackIfOtherSecurityStatusDropping(v bool)`

SetAttackIfOtherSecurityStatusDropping sets AttackIfOtherSecurityStatusDropping field to given value.


### GetAttackSecurityStatusThreshold

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackSecurityStatusThreshold() float64`

GetAttackSecurityStatusThreshold returns the AttackSecurityStatusThreshold field if non-nil, zero value otherwise.

### GetAttackSecurityStatusThresholdOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackSecurityStatusThresholdOk() (*float64, bool)`

GetAttackSecurityStatusThresholdOk returns a tuple with the AttackSecurityStatusThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackSecurityStatusThreshold

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetAttackSecurityStatusThreshold(v float64)`

SetAttackSecurityStatusThreshold sets AttackSecurityStatusThreshold field to given value.

### HasAttackSecurityStatusThreshold

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) HasAttackSecurityStatusThreshold() bool`

HasAttackSecurityStatusThreshold returns a boolean if a field has been set.

### GetAttackStandingThreshold

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackStandingThreshold() float64`

GetAttackStandingThreshold returns the AttackStandingThreshold field if non-nil, zero value otherwise.

### GetAttackStandingThresholdOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetAttackStandingThresholdOk() (*float64, bool)`

GetAttackStandingThresholdOk returns a tuple with the AttackStandingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackStandingThreshold

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetAttackStandingThreshold(v float64)`

SetAttackStandingThreshold sets AttackStandingThreshold field to given value.

### HasAttackStandingThreshold

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) HasAttackStandingThreshold() bool`

HasAttackStandingThreshold returns a boolean if a field has been set.

### GetFuelBayTake

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetFuelBayTake() string`

GetFuelBayTake returns the FuelBayTake field if non-nil, zero value otherwise.

### GetFuelBayTakeOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetFuelBayTakeOk() (*string, bool)`

GetFuelBayTakeOk returns a tuple with the FuelBayTake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelBayTake

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetFuelBayTake(v string)`

SetFuelBayTake sets FuelBayTake field to given value.


### GetFuelBayView

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetFuelBayView() string`

GetFuelBayView returns the FuelBayView field if non-nil, zero value otherwise.

### GetFuelBayViewOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetFuelBayViewOk() (*string, bool)`

GetFuelBayViewOk returns a tuple with the FuelBayView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelBayView

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetFuelBayView(v string)`

SetFuelBayView sets FuelBayView field to given value.


### GetFuels

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetFuels() []CorporationsCorporationIdStarbasesStarbaseIdGetFuelsInner`

GetFuels returns the Fuels field if non-nil, zero value otherwise.

### GetFuelsOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetFuelsOk() (*[]CorporationsCorporationIdStarbasesStarbaseIdGetFuelsInner, bool)`

GetFuelsOk returns a tuple with the Fuels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuels

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetFuels(v []CorporationsCorporationIdStarbasesStarbaseIdGetFuelsInner)`

SetFuels sets Fuels field to given value.

### HasFuels

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) HasFuels() bool`

HasFuels returns a boolean if a field has been set.

### GetOffline

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetOffline() string`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetOfflineOk() (*string, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetOffline(v string)`

SetOffline sets Offline field to given value.


### GetOnline

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetOnline() string`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetOnlineOk() (*string, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetOnline(v string)`

SetOnline sets Online field to given value.


### GetUnanchor

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetUnanchor() string`

GetUnanchor returns the Unanchor field if non-nil, zero value otherwise.

### GetUnanchorOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetUnanchorOk() (*string, bool)`

GetUnanchorOk returns a tuple with the Unanchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnanchor

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetUnanchor(v string)`

SetUnanchor sets Unanchor field to given value.


### GetUseAllianceStandings

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetUseAllianceStandings() bool`

GetUseAllianceStandings returns the UseAllianceStandings field if non-nil, zero value otherwise.

### GetUseAllianceStandingsOk

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) GetUseAllianceStandingsOk() (*bool, bool)`

GetUseAllianceStandingsOk returns a tuple with the UseAllianceStandings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllianceStandings

`func (o *CorporationsCorporationIdStarbasesStarbaseIdGet) SetUseAllianceStandings(v bool)`

SetUseAllianceStandings sets UseAllianceStandings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


