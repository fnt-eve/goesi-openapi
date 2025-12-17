# CorporationsFreelanceJobsListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to [**Cursor**](Cursor.md) |  | [optional] 
**FreelanceJobs** | [**[]FreelanceJobsDetailFreelancejob**](FreelanceJobsDetailFreelancejob.md) | List of freelance jobs | 

## Methods

### NewCorporationsFreelanceJobsListing

`func NewCorporationsFreelanceJobsListing(freelanceJobs []FreelanceJobsDetailFreelancejob, ) *CorporationsFreelanceJobsListing`

NewCorporationsFreelanceJobsListing instantiates a new CorporationsFreelanceJobsListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsFreelanceJobsListingWithDefaults

`func NewCorporationsFreelanceJobsListingWithDefaults() *CorporationsFreelanceJobsListing`

NewCorporationsFreelanceJobsListingWithDefaults instantiates a new CorporationsFreelanceJobsListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *CorporationsFreelanceJobsListing) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *CorporationsFreelanceJobsListing) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *CorporationsFreelanceJobsListing) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *CorporationsFreelanceJobsListing) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetFreelanceJobs

`func (o *CorporationsFreelanceJobsListing) GetFreelanceJobs() []FreelanceJobsDetailFreelancejob`

GetFreelanceJobs returns the FreelanceJobs field if non-nil, zero value otherwise.

### GetFreelanceJobsOk

`func (o *CorporationsFreelanceJobsListing) GetFreelanceJobsOk() (*[]FreelanceJobsDetailFreelancejob, bool)`

GetFreelanceJobsOk returns a tuple with the FreelanceJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreelanceJobs

`func (o *CorporationsFreelanceJobsListing) SetFreelanceJobs(v []FreelanceJobsDetailFreelancejob)`

SetFreelanceJobs sets FreelanceJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


