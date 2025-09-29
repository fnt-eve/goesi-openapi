# FwLeaderboardsCharactersGetVictoryPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTotal** | [**[]FwLeaderboardsCharactersGetVictoryPointsActiveTotalInner**](FwLeaderboardsCharactersGetVictoryPointsActiveTotalInner.md) | Top 100 ranking of pilots active in faction warfare by total victory points. A pilot is considered \&quot;active\&quot; if they have participated in faction warfare in the past 14 days | 
**LastWeek** | [**[]FwLeaderboardsCharactersGetVictoryPointsLastWeekInner**](FwLeaderboardsCharactersGetVictoryPointsLastWeekInner.md) | Top 100 ranking of pilots by victory points in the past week | 
**Yesterday** | [**[]FwLeaderboardsCharactersGetVictoryPointsYesterdayInner**](FwLeaderboardsCharactersGetVictoryPointsYesterdayInner.md) | Top 100 ranking of pilots by victory points in the past day | 

## Methods

### NewFwLeaderboardsCharactersGetVictoryPoints

`func NewFwLeaderboardsCharactersGetVictoryPoints(activeTotal []FwLeaderboardsCharactersGetVictoryPointsActiveTotalInner, lastWeek []FwLeaderboardsCharactersGetVictoryPointsLastWeekInner, yesterday []FwLeaderboardsCharactersGetVictoryPointsYesterdayInner, ) *FwLeaderboardsCharactersGetVictoryPoints`

NewFwLeaderboardsCharactersGetVictoryPoints instantiates a new FwLeaderboardsCharactersGetVictoryPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwLeaderboardsCharactersGetVictoryPointsWithDefaults

`func NewFwLeaderboardsCharactersGetVictoryPointsWithDefaults() *FwLeaderboardsCharactersGetVictoryPoints`

NewFwLeaderboardsCharactersGetVictoryPointsWithDefaults instantiates a new FwLeaderboardsCharactersGetVictoryPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTotal

`func (o *FwLeaderboardsCharactersGetVictoryPoints) GetActiveTotal() []FwLeaderboardsCharactersGetVictoryPointsActiveTotalInner`

GetActiveTotal returns the ActiveTotal field if non-nil, zero value otherwise.

### GetActiveTotalOk

`func (o *FwLeaderboardsCharactersGetVictoryPoints) GetActiveTotalOk() (*[]FwLeaderboardsCharactersGetVictoryPointsActiveTotalInner, bool)`

GetActiveTotalOk returns a tuple with the ActiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTotal

`func (o *FwLeaderboardsCharactersGetVictoryPoints) SetActiveTotal(v []FwLeaderboardsCharactersGetVictoryPointsActiveTotalInner)`

SetActiveTotal sets ActiveTotal field to given value.


### GetLastWeek

`func (o *FwLeaderboardsCharactersGetVictoryPoints) GetLastWeek() []FwLeaderboardsCharactersGetVictoryPointsLastWeekInner`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *FwLeaderboardsCharactersGetVictoryPoints) GetLastWeekOk() (*[]FwLeaderboardsCharactersGetVictoryPointsLastWeekInner, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *FwLeaderboardsCharactersGetVictoryPoints) SetLastWeek(v []FwLeaderboardsCharactersGetVictoryPointsLastWeekInner)`

SetLastWeek sets LastWeek field to given value.


### GetYesterday

`func (o *FwLeaderboardsCharactersGetVictoryPoints) GetYesterday() []FwLeaderboardsCharactersGetVictoryPointsYesterdayInner`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *FwLeaderboardsCharactersGetVictoryPoints) GetYesterdayOk() (*[]FwLeaderboardsCharactersGetVictoryPointsYesterdayInner, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *FwLeaderboardsCharactersGetVictoryPoints) SetYesterday(v []FwLeaderboardsCharactersGetVictoryPointsYesterdayInner)`

SetYesterday sets Yesterday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


