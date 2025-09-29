# CharactersCharacterIdIndustryJobsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **int64** | Job activity ID | 
**BlueprintId** | **int64** |  | 
**BlueprintLocationId** | **int64** | Location ID of the location from which the blueprint was installed. Normally a station ID, but can also be an asset (e.g. container) or corporation facility | 
**BlueprintTypeId** | **int64** |  | 
**CompletedCharacterId** | Pointer to **int64** | ID of the character which completed this job | [optional] 
**CompletedDate** | Pointer to **time.Time** | Date and time when this job was completed | [optional] 
**Cost** | Pointer to **float64** | The sume of job installation fee and industry facility tax | [optional] 
**Duration** | **int64** | Job duration in seconds | 
**EndDate** | **time.Time** | Date and time when this job finished | 
**FacilityId** | **int64** | ID of the facility where this job is running | 
**InstallerId** | **int64** | ID of the character which installed this job | 
**JobId** | **int64** | Unique job ID | 
**LicensedRuns** | Pointer to **int64** | Number of runs blueprint is licensed for | [optional] 
**OutputLocationId** | **int64** | Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility | 
**PauseDate** | Pointer to **time.Time** | Date and time when this job was paused (i.e. time when the facility where this job was installed went offline) | [optional] 
**Probability** | Pointer to **float64** | Chance of success for invention | [optional] 
**ProductTypeId** | Pointer to **int64** | Type ID of product (manufactured, copied or invented) | [optional] 
**Runs** | **int64** | Number of runs for a manufacturing job, or number of copies to make for a blueprint copy | 
**StartDate** | **time.Time** | Date and time when this job started | 
**StationId** | **int64** | ID of the station where industry facility is located | 
**Status** | **string** |  | 
**SuccessfulRuns** | Pointer to **int64** | Number of successful runs for this job. Equal to runs unless this is an invention job | [optional] 

## Methods

### NewCharactersCharacterIdIndustryJobsGetInner

`func NewCharactersCharacterIdIndustryJobsGetInner(activityId int64, blueprintId int64, blueprintLocationId int64, blueprintTypeId int64, duration int64, endDate time.Time, facilityId int64, installerId int64, jobId int64, outputLocationId int64, runs int64, startDate time.Time, stationId int64, status string, ) *CharactersCharacterIdIndustryJobsGetInner`

NewCharactersCharacterIdIndustryJobsGetInner instantiates a new CharactersCharacterIdIndustryJobsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharactersCharacterIdIndustryJobsGetInnerWithDefaults

`func NewCharactersCharacterIdIndustryJobsGetInnerWithDefaults() *CharactersCharacterIdIndustryJobsGetInner`

NewCharactersCharacterIdIndustryJobsGetInnerWithDefaults instantiates a new CharactersCharacterIdIndustryJobsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.


### GetBlueprintId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetBlueprintId() int64`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetBlueprintIdOk() (*int64, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetBlueprintId(v int64)`

SetBlueprintId sets BlueprintId field to given value.


### GetBlueprintLocationId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetBlueprintLocationId() int64`

GetBlueprintLocationId returns the BlueprintLocationId field if non-nil, zero value otherwise.

### GetBlueprintLocationIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetBlueprintLocationIdOk() (*int64, bool)`

GetBlueprintLocationIdOk returns a tuple with the BlueprintLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintLocationId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetBlueprintLocationId(v int64)`

SetBlueprintLocationId sets BlueprintLocationId field to given value.


### GetBlueprintTypeId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetBlueprintTypeId() int64`

GetBlueprintTypeId returns the BlueprintTypeId field if non-nil, zero value otherwise.

### GetBlueprintTypeIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetBlueprintTypeIdOk() (*int64, bool)`

GetBlueprintTypeIdOk returns a tuple with the BlueprintTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintTypeId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetBlueprintTypeId(v int64)`

SetBlueprintTypeId sets BlueprintTypeId field to given value.


### GetCompletedCharacterId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetCompletedCharacterId() int64`

GetCompletedCharacterId returns the CompletedCharacterId field if non-nil, zero value otherwise.

### GetCompletedCharacterIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetCompletedCharacterIdOk() (*int64, bool)`

GetCompletedCharacterIdOk returns a tuple with the CompletedCharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCharacterId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetCompletedCharacterId(v int64)`

SetCompletedCharacterId sets CompletedCharacterId field to given value.

### HasCompletedCharacterId

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasCompletedCharacterId() bool`

HasCompletedCharacterId returns a boolean if a field has been set.

### GetCompletedDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCost

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDuration

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetEndDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetFacilityId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetFacilityId() int64`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetFacilityIdOk() (*int64, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetFacilityId(v int64)`

SetFacilityId sets FacilityId field to given value.


### GetInstallerId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetInstallerId() int64`

GetInstallerId returns the InstallerId field if non-nil, zero value otherwise.

### GetInstallerIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetInstallerIdOk() (*int64, bool)`

GetInstallerIdOk returns a tuple with the InstallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetInstallerId(v int64)`

SetInstallerId sets InstallerId field to given value.


### GetJobId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetJobId(v int64)`

SetJobId sets JobId field to given value.


### GetLicensedRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetLicensedRuns() int64`

GetLicensedRuns returns the LicensedRuns field if non-nil, zero value otherwise.

### GetLicensedRunsOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetLicensedRunsOk() (*int64, bool)`

GetLicensedRunsOk returns a tuple with the LicensedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetLicensedRuns(v int64)`

SetLicensedRuns sets LicensedRuns field to given value.

### HasLicensedRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasLicensedRuns() bool`

HasLicensedRuns returns a boolean if a field has been set.

### GetOutputLocationId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetOutputLocationId() int64`

GetOutputLocationId returns the OutputLocationId field if non-nil, zero value otherwise.

### GetOutputLocationIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetOutputLocationIdOk() (*int64, bool)`

GetOutputLocationIdOk returns a tuple with the OutputLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLocationId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetOutputLocationId(v int64)`

SetOutputLocationId sets OutputLocationId field to given value.


### GetPauseDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetPauseDate() time.Time`

GetPauseDate returns the PauseDate field if non-nil, zero value otherwise.

### GetPauseDateOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetPauseDateOk() (*time.Time, bool)`

GetPauseDateOk returns a tuple with the PauseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetPauseDate(v time.Time)`

SetPauseDate sets PauseDate field to given value.

### HasPauseDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasPauseDate() bool`

HasPauseDate returns a boolean if a field has been set.

### GetProbability

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetProbability() float64`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetProbabilityOk() (*float64, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetProbability(v float64)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### GetProductTypeId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetProductTypeId() int64`

GetProductTypeId returns the ProductTypeId field if non-nil, zero value otherwise.

### GetProductTypeIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetProductTypeIdOk() (*int64, bool)`

GetProductTypeIdOk returns a tuple with the ProductTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetProductTypeId(v int64)`

SetProductTypeId sets ProductTypeId field to given value.

### HasProductTypeId

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasProductTypeId() bool`

HasProductTypeId returns a boolean if a field has been set.

### GetRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetRuns() int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetRunsOk() (*int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetRuns(v int64)`

SetRuns sets Runs field to given value.


### GetStartDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetStationId

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetStationId() int64`

GetStationId returns the StationId field if non-nil, zero value otherwise.

### GetStationIdOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetStationIdOk() (*int64, bool)`

GetStationIdOk returns a tuple with the StationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationId

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetStationId(v int64)`

SetStationId sets StationId field to given value.


### GetStatus

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuccessfulRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetSuccessfulRuns() int64`

GetSuccessfulRuns returns the SuccessfulRuns field if non-nil, zero value otherwise.

### GetSuccessfulRunsOk

`func (o *CharactersCharacterIdIndustryJobsGetInner) GetSuccessfulRunsOk() (*int64, bool)`

GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) SetSuccessfulRuns(v int64)`

SetSuccessfulRuns sets SuccessfulRuns field to given value.

### HasSuccessfulRuns

`func (o *CharactersCharacterIdIndustryJobsGetInner) HasSuccessfulRuns() bool`

HasSuccessfulRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


