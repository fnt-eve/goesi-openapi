# MilitaryCampaignsObjectivesListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to [**Cursor**](Cursor.md) |  | [optional] 
**Objectives** | [**[]MilitaryCampaignsObjectivesDetailObjective**](MilitaryCampaignsObjectivesDetailObjective.md) | List of military campaign objectives | 

## Methods

### NewMilitaryCampaignsObjectivesListing

`func NewMilitaryCampaignsObjectivesListing(objectives []MilitaryCampaignsObjectivesDetailObjective, ) *MilitaryCampaignsObjectivesListing`

NewMilitaryCampaignsObjectivesListing instantiates a new MilitaryCampaignsObjectivesListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilitaryCampaignsObjectivesListingWithDefaults

`func NewMilitaryCampaignsObjectivesListingWithDefaults() *MilitaryCampaignsObjectivesListing`

NewMilitaryCampaignsObjectivesListingWithDefaults instantiates a new MilitaryCampaignsObjectivesListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *MilitaryCampaignsObjectivesListing) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *MilitaryCampaignsObjectivesListing) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *MilitaryCampaignsObjectivesListing) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *MilitaryCampaignsObjectivesListing) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetObjectives

`func (o *MilitaryCampaignsObjectivesListing) GetObjectives() []MilitaryCampaignsObjectivesDetailObjective`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *MilitaryCampaignsObjectivesListing) GetObjectivesOk() (*[]MilitaryCampaignsObjectivesDetailObjective, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *MilitaryCampaignsObjectivesListing) SetObjectives(v []MilitaryCampaignsObjectivesDetailObjective)`

SetObjectives sets Objectives field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


