# FwLeaderboardsCharactersGetKills

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTotal** | [**[]FwLeaderboardsCharactersGetKillsActiveTotalInner**](FwLeaderboardsCharactersGetKillsActiveTotalInner.md) | Top 100 ranking of pilots active in faction warfare by total kills. A pilot is considered \&quot;active\&quot; if they have participated in faction warfare in the past 14 days | 
**LastWeek** | [**[]FwLeaderboardsCharactersGetKillsLastWeekInner**](FwLeaderboardsCharactersGetKillsLastWeekInner.md) | Top 100 ranking of pilots by kills in the past week | 
**Yesterday** | [**[]FwLeaderboardsCharactersGetKillsYesterdayInner**](FwLeaderboardsCharactersGetKillsYesterdayInner.md) | Top 100 ranking of pilots by kills in the past day | 

## Methods

### NewFwLeaderboardsCharactersGetKills

`func NewFwLeaderboardsCharactersGetKills(activeTotal []FwLeaderboardsCharactersGetKillsActiveTotalInner, lastWeek []FwLeaderboardsCharactersGetKillsLastWeekInner, yesterday []FwLeaderboardsCharactersGetKillsYesterdayInner, ) *FwLeaderboardsCharactersGetKills`

NewFwLeaderboardsCharactersGetKills instantiates a new FwLeaderboardsCharactersGetKills object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwLeaderboardsCharactersGetKillsWithDefaults

`func NewFwLeaderboardsCharactersGetKillsWithDefaults() *FwLeaderboardsCharactersGetKills`

NewFwLeaderboardsCharactersGetKillsWithDefaults instantiates a new FwLeaderboardsCharactersGetKills object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTotal

`func (o *FwLeaderboardsCharactersGetKills) GetActiveTotal() []FwLeaderboardsCharactersGetKillsActiveTotalInner`

GetActiveTotal returns the ActiveTotal field if non-nil, zero value otherwise.

### GetActiveTotalOk

`func (o *FwLeaderboardsCharactersGetKills) GetActiveTotalOk() (*[]FwLeaderboardsCharactersGetKillsActiveTotalInner, bool)`

GetActiveTotalOk returns a tuple with the ActiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTotal

`func (o *FwLeaderboardsCharactersGetKills) SetActiveTotal(v []FwLeaderboardsCharactersGetKillsActiveTotalInner)`

SetActiveTotal sets ActiveTotal field to given value.


### GetLastWeek

`func (o *FwLeaderboardsCharactersGetKills) GetLastWeek() []FwLeaderboardsCharactersGetKillsLastWeekInner`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *FwLeaderboardsCharactersGetKills) GetLastWeekOk() (*[]FwLeaderboardsCharactersGetKillsLastWeekInner, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *FwLeaderboardsCharactersGetKills) SetLastWeek(v []FwLeaderboardsCharactersGetKillsLastWeekInner)`

SetLastWeek sets LastWeek field to given value.


### GetYesterday

`func (o *FwLeaderboardsCharactersGetKills) GetYesterday() []FwLeaderboardsCharactersGetKillsYesterdayInner`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *FwLeaderboardsCharactersGetKills) GetYesterdayOk() (*[]FwLeaderboardsCharactersGetKillsYesterdayInner, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *FwLeaderboardsCharactersGetKills) SetYesterday(v []FwLeaderboardsCharactersGetKillsYesterdayInner)`

SetYesterday sets Yesterday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


