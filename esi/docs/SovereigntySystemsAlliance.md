# SovereigntySystemsAlliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceId** | **int64** |  | 
**ClaimedSince** | **time.Time** | Time the claim was made | 
**CorporationId** | **int64** |  | 
**Development** | [**SovereigntySystemsDevelopment**](SovereigntySystemsDevelopment.md) |  | 
**IsCapitalSystem** | **bool** | Whether the system is the capital system of the alliance | 
**SovereigntyHub** | [**SovereigntySystemsSovereigntyhub**](SovereigntySystemsSovereigntyhub.md) |  | 

## Methods

### NewSovereigntySystemsAlliance

`func NewSovereigntySystemsAlliance(allianceId int64, claimedSince time.Time, corporationId int64, development SovereigntySystemsDevelopment, isCapitalSystem bool, sovereigntyHub SovereigntySystemsSovereigntyhub, ) *SovereigntySystemsAlliance`

NewSovereigntySystemsAlliance instantiates a new SovereigntySystemsAlliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSovereigntySystemsAllianceWithDefaults

`func NewSovereigntySystemsAllianceWithDefaults() *SovereigntySystemsAlliance`

NewSovereigntySystemsAllianceWithDefaults instantiates a new SovereigntySystemsAlliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceId

`func (o *SovereigntySystemsAlliance) GetAllianceId() int64`

GetAllianceId returns the AllianceId field if non-nil, zero value otherwise.

### GetAllianceIdOk

`func (o *SovereigntySystemsAlliance) GetAllianceIdOk() (*int64, bool)`

GetAllianceIdOk returns a tuple with the AllianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceId

`func (o *SovereigntySystemsAlliance) SetAllianceId(v int64)`

SetAllianceId sets AllianceId field to given value.


### GetClaimedSince

`func (o *SovereigntySystemsAlliance) GetClaimedSince() time.Time`

GetClaimedSince returns the ClaimedSince field if non-nil, zero value otherwise.

### GetClaimedSinceOk

`func (o *SovereigntySystemsAlliance) GetClaimedSinceOk() (*time.Time, bool)`

GetClaimedSinceOk returns a tuple with the ClaimedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedSince

`func (o *SovereigntySystemsAlliance) SetClaimedSince(v time.Time)`

SetClaimedSince sets ClaimedSince field to given value.


### GetCorporationId

`func (o *SovereigntySystemsAlliance) GetCorporationId() int64`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *SovereigntySystemsAlliance) GetCorporationIdOk() (*int64, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *SovereigntySystemsAlliance) SetCorporationId(v int64)`

SetCorporationId sets CorporationId field to given value.


### GetDevelopment

`func (o *SovereigntySystemsAlliance) GetDevelopment() SovereigntySystemsDevelopment`

GetDevelopment returns the Development field if non-nil, zero value otherwise.

### GetDevelopmentOk

`func (o *SovereigntySystemsAlliance) GetDevelopmentOk() (*SovereigntySystemsDevelopment, bool)`

GetDevelopmentOk returns a tuple with the Development field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopment

`func (o *SovereigntySystemsAlliance) SetDevelopment(v SovereigntySystemsDevelopment)`

SetDevelopment sets Development field to given value.


### GetIsCapitalSystem

`func (o *SovereigntySystemsAlliance) GetIsCapitalSystem() bool`

GetIsCapitalSystem returns the IsCapitalSystem field if non-nil, zero value otherwise.

### GetIsCapitalSystemOk

`func (o *SovereigntySystemsAlliance) GetIsCapitalSystemOk() (*bool, bool)`

GetIsCapitalSystemOk returns a tuple with the IsCapitalSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCapitalSystem

`func (o *SovereigntySystemsAlliance) SetIsCapitalSystem(v bool)`

SetIsCapitalSystem sets IsCapitalSystem field to given value.


### GetSovereigntyHub

`func (o *SovereigntySystemsAlliance) GetSovereigntyHub() SovereigntySystemsSovereigntyhub`

GetSovereigntyHub returns the SovereigntyHub field if non-nil, zero value otherwise.

### GetSovereigntyHubOk

`func (o *SovereigntySystemsAlliance) GetSovereigntyHubOk() (*SovereigntySystemsSovereigntyhub, bool)`

GetSovereigntyHubOk returns a tuple with the SovereigntyHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSovereigntyHub

`func (o *SovereigntySystemsAlliance) SetSovereigntyHub(v SovereigntySystemsSovereigntyhub)`

SetSovereigntyHub sets SovereigntyHub field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


