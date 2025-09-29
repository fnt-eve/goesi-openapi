# CorporationsCorporationIdContactsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | **int64** |  | 
**ContactType** | **string** |  | 
**IsWatched** | Pointer to **bool** | Whether this contact is being watched | [optional] 
**LabelIds** | Pointer to **[]int64** |  | [optional] 
**Standing** | **float64** | Standing of the contact | 

## Methods

### NewCorporationsCorporationIdContactsGetInner

`func NewCorporationsCorporationIdContactsGetInner(contactId int64, contactType string, standing float64, ) *CorporationsCorporationIdContactsGetInner`

NewCorporationsCorporationIdContactsGetInner instantiates a new CorporationsCorporationIdContactsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationsCorporationIdContactsGetInnerWithDefaults

`func NewCorporationsCorporationIdContactsGetInnerWithDefaults() *CorporationsCorporationIdContactsGetInner`

NewCorporationsCorporationIdContactsGetInnerWithDefaults instantiates a new CorporationsCorporationIdContactsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *CorporationsCorporationIdContactsGetInner) GetContactId() int64`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *CorporationsCorporationIdContactsGetInner) GetContactIdOk() (*int64, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *CorporationsCorporationIdContactsGetInner) SetContactId(v int64)`

SetContactId sets ContactId field to given value.


### GetContactType

`func (o *CorporationsCorporationIdContactsGetInner) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *CorporationsCorporationIdContactsGetInner) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *CorporationsCorporationIdContactsGetInner) SetContactType(v string)`

SetContactType sets ContactType field to given value.


### GetIsWatched

`func (o *CorporationsCorporationIdContactsGetInner) GetIsWatched() bool`

GetIsWatched returns the IsWatched field if non-nil, zero value otherwise.

### GetIsWatchedOk

`func (o *CorporationsCorporationIdContactsGetInner) GetIsWatchedOk() (*bool, bool)`

GetIsWatchedOk returns a tuple with the IsWatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWatched

`func (o *CorporationsCorporationIdContactsGetInner) SetIsWatched(v bool)`

SetIsWatched sets IsWatched field to given value.

### HasIsWatched

`func (o *CorporationsCorporationIdContactsGetInner) HasIsWatched() bool`

HasIsWatched returns a boolean if a field has been set.

### GetLabelIds

`func (o *CorporationsCorporationIdContactsGetInner) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CorporationsCorporationIdContactsGetInner) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CorporationsCorporationIdContactsGetInner) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CorporationsCorporationIdContactsGetInner) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetStanding

`func (o *CorporationsCorporationIdContactsGetInner) GetStanding() float64`

GetStanding returns the Standing field if non-nil, zero value otherwise.

### GetStandingOk

`func (o *CorporationsCorporationIdContactsGetInner) GetStandingOk() (*float64, bool)`

GetStandingOk returns a tuple with the Standing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanding

`func (o *CorporationsCorporationIdContactsGetInner) SetStanding(v float64)`

SetStanding sets Standing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


