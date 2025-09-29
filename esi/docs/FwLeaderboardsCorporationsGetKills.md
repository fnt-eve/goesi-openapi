# FwLeaderboardsCorporationsGetKills

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTotal** | [**[]FwLeaderboardsCorporationsGetKillsActiveTotalInner**](FwLeaderboardsCorporationsGetKillsActiveTotalInner.md) | Top 10 ranking of corporations active in faction warfare by total kills. A corporation is considered \&quot;active\&quot; if they have participated in faction warfare in the past 14 days | 
**LastWeek** | [**[]FwLeaderboardsCorporationsGetKillsLastWeekInner**](FwLeaderboardsCorporationsGetKillsLastWeekInner.md) | Top 10 ranking of corporations by kills in the past week | 
**Yesterday** | [**[]FwLeaderboardsCorporationsGetKillsYesterdayInner**](FwLeaderboardsCorporationsGetKillsYesterdayInner.md) | Top 10 ranking of corporations by kills in the past day | 

## Methods

### NewFwLeaderboardsCorporationsGetKills

`func NewFwLeaderboardsCorporationsGetKills(activeTotal []FwLeaderboardsCorporationsGetKillsActiveTotalInner, lastWeek []FwLeaderboardsCorporationsGetKillsLastWeekInner, yesterday []FwLeaderboardsCorporationsGetKillsYesterdayInner, ) *FwLeaderboardsCorporationsGetKills`

NewFwLeaderboardsCorporationsGetKills instantiates a new FwLeaderboardsCorporationsGetKills object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwLeaderboardsCorporationsGetKillsWithDefaults

`func NewFwLeaderboardsCorporationsGetKillsWithDefaults() *FwLeaderboardsCorporationsGetKills`

NewFwLeaderboardsCorporationsGetKillsWithDefaults instantiates a new FwLeaderboardsCorporationsGetKills object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTotal

`func (o *FwLeaderboardsCorporationsGetKills) GetActiveTotal() []FwLeaderboardsCorporationsGetKillsActiveTotalInner`

GetActiveTotal returns the ActiveTotal field if non-nil, zero value otherwise.

### GetActiveTotalOk

`func (o *FwLeaderboardsCorporationsGetKills) GetActiveTotalOk() (*[]FwLeaderboardsCorporationsGetKillsActiveTotalInner, bool)`

GetActiveTotalOk returns a tuple with the ActiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTotal

`func (o *FwLeaderboardsCorporationsGetKills) SetActiveTotal(v []FwLeaderboardsCorporationsGetKillsActiveTotalInner)`

SetActiveTotal sets ActiveTotal field to given value.


### GetLastWeek

`func (o *FwLeaderboardsCorporationsGetKills) GetLastWeek() []FwLeaderboardsCorporationsGetKillsLastWeekInner`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *FwLeaderboardsCorporationsGetKills) GetLastWeekOk() (*[]FwLeaderboardsCorporationsGetKillsLastWeekInner, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *FwLeaderboardsCorporationsGetKills) SetLastWeek(v []FwLeaderboardsCorporationsGetKillsLastWeekInner)`

SetLastWeek sets LastWeek field to given value.


### GetYesterday

`func (o *FwLeaderboardsCorporationsGetKills) GetYesterday() []FwLeaderboardsCorporationsGetKillsYesterdayInner`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *FwLeaderboardsCorporationsGetKills) GetYesterdayOk() (*[]FwLeaderboardsCorporationsGetKillsYesterdayInner, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *FwLeaderboardsCorporationsGetKills) SetYesterday(v []FwLeaderboardsCorporationsGetKillsYesterdayInner)`

SetYesterday sets Yesterday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


