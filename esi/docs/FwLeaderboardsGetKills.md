# FwLeaderboardsGetKills

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTotal** | [**[]FwLeaderboardsGetKillsActiveTotalInner**](FwLeaderboardsGetKillsActiveTotalInner.md) | Top 4 ranking of factions active in faction warfare by total kills. A faction is considered \&quot;active\&quot; if they have participated in faction warfare in the past 14 days | 
**LastWeek** | [**[]FwLeaderboardsGetKillsLastWeekInner**](FwLeaderboardsGetKillsLastWeekInner.md) | Top 4 ranking of factions by kills in the past week | 
**Yesterday** | [**[]FwLeaderboardsGetKillsYesterdayInner**](FwLeaderboardsGetKillsYesterdayInner.md) | Top 4 ranking of factions by kills in the past day | 

## Methods

### NewFwLeaderboardsGetKills

`func NewFwLeaderboardsGetKills(activeTotal []FwLeaderboardsGetKillsActiveTotalInner, lastWeek []FwLeaderboardsGetKillsLastWeekInner, yesterday []FwLeaderboardsGetKillsYesterdayInner, ) *FwLeaderboardsGetKills`

NewFwLeaderboardsGetKills instantiates a new FwLeaderboardsGetKills object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwLeaderboardsGetKillsWithDefaults

`func NewFwLeaderboardsGetKillsWithDefaults() *FwLeaderboardsGetKills`

NewFwLeaderboardsGetKillsWithDefaults instantiates a new FwLeaderboardsGetKills object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTotal

`func (o *FwLeaderboardsGetKills) GetActiveTotal() []FwLeaderboardsGetKillsActiveTotalInner`

GetActiveTotal returns the ActiveTotal field if non-nil, zero value otherwise.

### GetActiveTotalOk

`func (o *FwLeaderboardsGetKills) GetActiveTotalOk() (*[]FwLeaderboardsGetKillsActiveTotalInner, bool)`

GetActiveTotalOk returns a tuple with the ActiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTotal

`func (o *FwLeaderboardsGetKills) SetActiveTotal(v []FwLeaderboardsGetKillsActiveTotalInner)`

SetActiveTotal sets ActiveTotal field to given value.


### GetLastWeek

`func (o *FwLeaderboardsGetKills) GetLastWeek() []FwLeaderboardsGetKillsLastWeekInner`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *FwLeaderboardsGetKills) GetLastWeekOk() (*[]FwLeaderboardsGetKillsLastWeekInner, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *FwLeaderboardsGetKills) SetLastWeek(v []FwLeaderboardsGetKillsLastWeekInner)`

SetLastWeek sets LastWeek field to given value.


### GetYesterday

`func (o *FwLeaderboardsGetKills) GetYesterday() []FwLeaderboardsGetKillsYesterdayInner`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *FwLeaderboardsGetKills) GetYesterdayOk() (*[]FwLeaderboardsGetKillsYesterdayInner, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *FwLeaderboardsGetKills) SetYesterday(v []FwLeaderboardsGetKillsYesterdayInner)`

SetYesterday sets Yesterday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


