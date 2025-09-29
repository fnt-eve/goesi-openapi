# FwLeaderboardsGetVictoryPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTotal** | [**[]FwLeaderboardsGetVictoryPointsActiveTotalInner**](FwLeaderboardsGetVictoryPointsActiveTotalInner.md) | Top 4 ranking of factions active in faction warfare by total victory points. A faction is considered \&quot;active\&quot; if they have participated in faction warfare in the past 14 days | 
**LastWeek** | [**[]FwLeaderboardsGetVictoryPointsLastWeekInner**](FwLeaderboardsGetVictoryPointsLastWeekInner.md) | Top 4 ranking of factions by victory points in the past week | 
**Yesterday** | [**[]FwLeaderboardsGetVictoryPointsYesterdayInner**](FwLeaderboardsGetVictoryPointsYesterdayInner.md) | Top 4 ranking of factions by victory points in the past day | 

## Methods

### NewFwLeaderboardsGetVictoryPoints

`func NewFwLeaderboardsGetVictoryPoints(activeTotal []FwLeaderboardsGetVictoryPointsActiveTotalInner, lastWeek []FwLeaderboardsGetVictoryPointsLastWeekInner, yesterday []FwLeaderboardsGetVictoryPointsYesterdayInner, ) *FwLeaderboardsGetVictoryPoints`

NewFwLeaderboardsGetVictoryPoints instantiates a new FwLeaderboardsGetVictoryPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwLeaderboardsGetVictoryPointsWithDefaults

`func NewFwLeaderboardsGetVictoryPointsWithDefaults() *FwLeaderboardsGetVictoryPoints`

NewFwLeaderboardsGetVictoryPointsWithDefaults instantiates a new FwLeaderboardsGetVictoryPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTotal

`func (o *FwLeaderboardsGetVictoryPoints) GetActiveTotal() []FwLeaderboardsGetVictoryPointsActiveTotalInner`

GetActiveTotal returns the ActiveTotal field if non-nil, zero value otherwise.

### GetActiveTotalOk

`func (o *FwLeaderboardsGetVictoryPoints) GetActiveTotalOk() (*[]FwLeaderboardsGetVictoryPointsActiveTotalInner, bool)`

GetActiveTotalOk returns a tuple with the ActiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTotal

`func (o *FwLeaderboardsGetVictoryPoints) SetActiveTotal(v []FwLeaderboardsGetVictoryPointsActiveTotalInner)`

SetActiveTotal sets ActiveTotal field to given value.


### GetLastWeek

`func (o *FwLeaderboardsGetVictoryPoints) GetLastWeek() []FwLeaderboardsGetVictoryPointsLastWeekInner`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *FwLeaderboardsGetVictoryPoints) GetLastWeekOk() (*[]FwLeaderboardsGetVictoryPointsLastWeekInner, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *FwLeaderboardsGetVictoryPoints) SetLastWeek(v []FwLeaderboardsGetVictoryPointsLastWeekInner)`

SetLastWeek sets LastWeek field to given value.


### GetYesterday

`func (o *FwLeaderboardsGetVictoryPoints) GetYesterday() []FwLeaderboardsGetVictoryPointsYesterdayInner`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *FwLeaderboardsGetVictoryPoints) GetYesterdayOk() (*[]FwLeaderboardsGetVictoryPointsYesterdayInner, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *FwLeaderboardsGetVictoryPoints) SetYesterday(v []FwLeaderboardsGetVictoryPointsYesterdayInner)`

SetYesterday sets Yesterday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


