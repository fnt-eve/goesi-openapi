# CorporationsCorporationIdIndustryJobsGetInner

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
**LocationId** | **int64** | ID of the location for the industry facility | 
**OutputLocationId** | **int64** | Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility | 
**PauseDate** | Pointer to **time.Time** | Date and time when this job was paused (i.e. time when the facility where this job was installed went offline) | [optional] 
**Probability** | Pointer to **float64** | Chance of success for invention | [optional] 
**ProductTypeId** | Pointer to **int64** | Type ID of product (manufactured, copied or invented) | [optional] 
**Runs** | **int64** | Number of runs for a manufacturing job, or number of copies to make for a blueprint copy | 
**StartDate** | **time.Time** | Date and time when this job started | 
**Status** | **string** |  | 
**SuccessfulRuns** | Pointer to **int64** | Number of successful runs for this job. Equal to runs unless this is an invention job | [optional] 

## Methods

### NewCorporationsCorporationIdIndustryJobsGetInner

`func NewCorporationsCorporationIdIndustryJobsGetInner(activityId int64, blueprintId int64, blueprintLocationId int64, blueprintTypeId int64, duration int64, endDate time.Time, facilityId int64, installerId int64, jobId int64, locationId int64, outputLocationId int64, runs int64, startDate time.Time, status string, ) *CorporationsCorporationIdIndustryJobsGetInner`

NewCorporationsCorporationIdIndustryJobsGetInner instantiates a new CorporationsCorporationIdIndustryJobsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdIndustryJobsGetInnerWithDefaults

`func NewCorporationsCorporationIdIndustryJobsGetInnerWithDefaults() *CorporationsCorporationIdIndustryJobsGetInner`

NewCorporationsCorporationIdIndustryJobsGetInnerWithDefaults instantiates a new CorporationsCorporationIdIndustryJobsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.


### GetBlueprintId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetBlueprintId() int64`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetBlueprintIdOk() (*int64, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetBlueprintId(v int64)`

SetBlueprintId sets BlueprintId field to given value.


### GetBlueprintLocationId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetBlueprintLocationId() int64`

GetBlueprintLocationId returns the BlueprintLocationId field if non-nil, zero value otherwise.

### GetBlueprintLocationIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetBlueprintLocationIdOk() (*int64, bool)`

GetBlueprintLocationIdOk returns a tuple with the BlueprintLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintLocationId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetBlueprintLocationId(v int64)`

SetBlueprintLocationId sets BlueprintLocationId field to given value.


### GetBlueprintTypeId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetBlueprintTypeId() int64`

GetBlueprintTypeId returns the BlueprintTypeId field if non-nil, zero value otherwise.

### GetBlueprintTypeIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetBlueprintTypeIdOk() (*int64, bool)`

GetBlueprintTypeIdOk returns a tuple with the BlueprintTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintTypeId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetBlueprintTypeId(v int64)`

SetBlueprintTypeId sets BlueprintTypeId field to given value.


### GetCompletedCharacterId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetCompletedCharacterId() int64`

GetCompletedCharacterId returns the CompletedCharacterId field if non-nil, zero value otherwise.

### GetCompletedCharacterIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetCompletedCharacterIdOk() (*int64, bool)`

GetCompletedCharacterIdOk returns a tuple with the CompletedCharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCharacterId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetCompletedCharacterId(v int64)`

SetCompletedCharacterId sets CompletedCharacterId field to given value.

### HasCompletedCharacterId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasCompletedCharacterId() bool`

HasCompletedCharacterId returns a boolean if a field has been set.

### GetCompletedDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCost

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDuration

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetEndDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetFacilityId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetFacilityId() int64`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetFacilityIdOk() (*int64, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetFacilityId(v int64)`

SetFacilityId sets FacilityId field to given value.


### GetInstallerId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetInstallerId() int64`

GetInstallerId returns the InstallerId field if non-nil, zero value otherwise.

### GetInstallerIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetInstallerIdOk() (*int64, bool)`

GetInstallerIdOk returns a tuple with the InstallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetInstallerId(v int64)`

SetInstallerId sets InstallerId field to given value.


### GetJobId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetJobId(v int64)`

SetJobId sets JobId field to given value.


### GetLicensedRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetLicensedRuns() int64`

GetLicensedRuns returns the LicensedRuns field if non-nil, zero value otherwise.

### GetLicensedRunsOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetLicensedRunsOk() (*int64, bool)`

GetLicensedRunsOk returns a tuple with the LicensedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetLicensedRuns(v int64)`

SetLicensedRuns sets LicensedRuns field to given value.

### HasLicensedRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasLicensedRuns() bool`

HasLicensedRuns returns a boolean if a field has been set.

### GetLocationId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetLocationId() int64`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetLocationIdOk() (*int64, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetLocationId(v int64)`

SetLocationId sets LocationId field to given value.


### GetOutputLocationId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetOutputLocationId() int64`

GetOutputLocationId returns the OutputLocationId field if non-nil, zero value otherwise.

### GetOutputLocationIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetOutputLocationIdOk() (*int64, bool)`

GetOutputLocationIdOk returns a tuple with the OutputLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLocationId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetOutputLocationId(v int64)`

SetOutputLocationId sets OutputLocationId field to given value.


### GetPauseDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetPauseDate() time.Time`

GetPauseDate returns the PauseDate field if non-nil, zero value otherwise.

### GetPauseDateOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetPauseDateOk() (*time.Time, bool)`

GetPauseDateOk returns a tuple with the PauseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetPauseDate(v time.Time)`

SetPauseDate sets PauseDate field to given value.

### HasPauseDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasPauseDate() bool`

HasPauseDate returns a boolean if a field has been set.

### GetProbability

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetProbability() float64`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetProbabilityOk() (*float64, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetProbability(v float64)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### GetProductTypeId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetProductTypeId() int64`

GetProductTypeId returns the ProductTypeId field if non-nil, zero value otherwise.

### GetProductTypeIdOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetProductTypeIdOk() (*int64, bool)`

GetProductTypeIdOk returns a tuple with the ProductTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetProductTypeId(v int64)`

SetProductTypeId sets ProductTypeId field to given value.

### HasProductTypeId

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasProductTypeId() bool`

HasProductTypeId returns a boolean if a field has been set.

### GetRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetRuns() int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetRunsOk() (*int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetRuns(v int64)`

SetRuns sets Runs field to given value.


### GetStartDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetStatus

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuccessfulRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetSuccessfulRuns() int64`

GetSuccessfulRuns returns the SuccessfulRuns field if non-nil, zero value otherwise.

### GetSuccessfulRunsOk

`func (o *CorporationsCorporationIdIndustryJobsGetInner) GetSuccessfulRunsOk() (*int64, bool)`

GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) SetSuccessfulRuns(v int64)`

SetSuccessfulRuns sets SuccessfulRuns field to given value.

### HasSuccessfulRuns

`func (o *CorporationsCorporationIdIndustryJobsGetInner) HasSuccessfulRuns() bool`

HasSuccessfulRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


