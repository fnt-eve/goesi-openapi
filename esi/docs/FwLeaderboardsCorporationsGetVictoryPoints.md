# FwLeaderboardsCorporationsGetVictoryPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTotal** | [**[]FwLeaderboardsCorporationsGetVictoryPointsActiveTotalInner**](FwLeaderboardsCorporationsGetVictoryPointsActiveTotalInner.md) | Top 10 ranking of corporations active in faction warfare by total victory points. A corporation is considered \&quot;active\&quot; if they have participated in faction warfare in the past 14 days | 
**LastWeek** | [**[]FwLeaderboardsCorporationsGetVictoryPointsLastWeekInner**](FwLeaderboardsCorporationsGetVictoryPointsLastWeekInner.md) | Top 10 ranking of corporations by victory points in the past week | 
**Yesterday** | [**[]FwLeaderboardsCorporationsGetVictoryPointsYesterdayInner**](FwLeaderboardsCorporationsGetVictoryPointsYesterdayInner.md) | Top 10 ranking of corporations by victory points in the past day | 

## Methods

### NewFwLeaderboardsCorporationsGetVictoryPoints

`func NewFwLeaderboardsCorporationsGetVictoryPoints(activeTotal []FwLeaderboardsCorporationsGetVictoryPointsActiveTotalInner, lastWeek []FwLeaderboardsCorporationsGetVictoryPointsLastWeekInner, yesterday []FwLeaderboardsCorporationsGetVictoryPointsYesterdayInner, ) *FwLeaderboardsCorporationsGetVictoryPoints`

NewFwLeaderboardsCorporationsGetVictoryPoints instantiates a new FwLeaderboardsCorporationsGetVictoryPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwLeaderboardsCorporationsGetVictoryPointsWithDefaults

`func NewFwLeaderboardsCorporationsGetVictoryPointsWithDefaults() *FwLeaderboardsCorporationsGetVictoryPoints`

NewFwLeaderboardsCorporationsGetVictoryPointsWithDefaults instantiates a new FwLeaderboardsCorporationsGetVictoryPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTotal

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) GetActiveTotal() []FwLeaderboardsCorporationsGetVictoryPointsActiveTotalInner`

GetActiveTotal returns the ActiveTotal field if non-nil, zero value otherwise.

### GetActiveTotalOk

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) GetActiveTotalOk() (*[]FwLeaderboardsCorporationsGetVictoryPointsActiveTotalInner, bool)`

GetActiveTotalOk returns a tuple with the ActiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTotal

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) SetActiveTotal(v []FwLeaderboardsCorporationsGetVictoryPointsActiveTotalInner)`

SetActiveTotal sets ActiveTotal field to given value.


### GetLastWeek

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) GetLastWeek() []FwLeaderboardsCorporationsGetVictoryPointsLastWeekInner`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) GetLastWeekOk() (*[]FwLeaderboardsCorporationsGetVictoryPointsLastWeekInner, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) SetLastWeek(v []FwLeaderboardsCorporationsGetVictoryPointsLastWeekInner)`

SetLastWeek sets LastWeek field to given value.


### GetYesterday

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) GetYesterday() []FwLeaderboardsCorporationsGetVictoryPointsYesterdayInner`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) GetYesterdayOk() (*[]FwLeaderboardsCorporationsGetVictoryPointsYesterdayInner, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *FwLeaderboardsCorporationsGetVictoryPoints) SetYesterday(v []FwLeaderboardsCorporationsGetVictoryPointsYesterdayInner)`

SetYesterday sets Yesterday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


