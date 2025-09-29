# FwStatsGetInnerKills

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastWeek** | **int64** | Last week&#39;s total number of kills against enemy factions | 
**Total** | **int64** | Total number of kills against enemy factions since faction warfare began | 
**Yesterday** | **int64** | Yesterday&#39;s total number of kills against enemy factions | 

## Methods

### NewFwStatsGetInnerKills

`func NewFwStatsGetInnerKills(lastWeek int64, total int64, yesterday int64, ) *FwStatsGetInnerKills`

NewFwStatsGetInnerKills instantiates a new FwStatsGetInnerKills object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwStatsGetInnerKillsWithDefaults

`func NewFwStatsGetInnerKillsWithDefaults() *FwStatsGetInnerKills`

NewFwStatsGetInnerKillsWithDefaults instantiates a new FwStatsGetInnerKills object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastWeek

`func (o *FwStatsGetInnerKills) GetLastWeek() int64`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *FwStatsGetInnerKills) GetLastWeekOk() (*int64, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *FwStatsGetInnerKills) SetLastWeek(v int64)`

SetLastWeek sets LastWeek field to given value.


### GetTotal

`func (o *FwStatsGetInnerKills) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FwStatsGetInnerKills) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FwStatsGetInnerKills) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetYesterday

`func (o *FwStatsGetInnerKills) GetYesterday() int64`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *FwStatsGetInnerKills) GetYesterdayOk() (*int64, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *FwStatsGetInnerKills) SetYesterday(v int64)`

SetYesterday sets Yesterday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


