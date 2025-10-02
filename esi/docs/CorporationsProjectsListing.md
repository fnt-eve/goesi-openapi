# CorporationsProjectsListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to [**Cursor**](Cursor.md) |  | [optional] 
**Projects** | [**[]CorporationsProjectsDetailProject**](CorporationsProjectsDetailProject.md) | List of projects | 

## Methods

### NewCorporationsProjectsListing

`func NewCorporationsProjectsListing(projects []CorporationsProjectsDetailProject, ) *CorporationsProjectsListing`

NewCorporationsProjectsListing instantiates a new CorporationsProjectsListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsProjectsListingWithDefaults

`func NewCorporationsProjectsListingWithDefaults() *CorporationsProjectsListing`

NewCorporationsProjectsListingWithDefaults instantiates a new CorporationsProjectsListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *CorporationsProjectsListing) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *CorporationsProjectsListing) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *CorporationsProjectsListing) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *CorporationsProjectsListing) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetProjects

`func (o *CorporationsProjectsListing) GetProjects() []CorporationsProjectsDetailProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CorporationsProjectsListing) GetProjectsOk() (*[]CorporationsProjectsDetailProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CorporationsProjectsListing) SetProjects(v []CorporationsProjectsDetailProject)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


