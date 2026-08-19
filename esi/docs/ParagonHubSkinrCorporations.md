# ParagonHubSkinrCorporations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to [**Cursor**](Cursor.md) |  | [optional] 
**Listings** | [**[]ParagonHubSkinrInternalItem**](ParagonHubSkinrInternalItem.md) | Page of SKINR listings targeted at the corporation | 

## Methods

### NewParagonHubSkinrCorporations

`func NewParagonHubSkinrCorporations(listings []ParagonHubSkinrInternalItem, ) *ParagonHubSkinrCorporations`

NewParagonHubSkinrCorporations instantiates a new ParagonHubSkinrCorporations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParagonHubSkinrCorporationsWithDefaults

`func NewParagonHubSkinrCorporationsWithDefaults() *ParagonHubSkinrCorporations`

NewParagonHubSkinrCorporationsWithDefaults instantiates a new ParagonHubSkinrCorporations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ParagonHubSkinrCorporations) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ParagonHubSkinrCorporations) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ParagonHubSkinrCorporations) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ParagonHubSkinrCorporations) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetListings

`func (o *ParagonHubSkinrCorporations) GetListings() []ParagonHubSkinrInternalItem`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *ParagonHubSkinrCorporations) GetListingsOk() (*[]ParagonHubSkinrInternalItem, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *ParagonHubSkinrCorporations) SetListings(v []ParagonHubSkinrInternalItem)`

SetListings sets Listings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


