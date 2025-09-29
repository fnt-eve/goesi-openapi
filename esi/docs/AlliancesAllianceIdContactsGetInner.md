# AlliancesAllianceIdContactsGetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | **int64** |  | 
**ContactType** | **string** |  | 
**LabelIds** | Pointer to **[]int64** |  | [optional] 
**Standing** | **float64** | Standing of the contact | 

## Methods

### NewAlliancesAllianceIdContactsGetInner

`func NewAlliancesAllianceIdContactsGetInner(contactId int64, contactType string, standing float64, ) *AlliancesAllianceIdContactsGetInner`

NewAlliancesAllianceIdContactsGetInner instantiates a new AlliancesAllianceIdContactsGetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlliancesAllianceIdContactsGetInnerWithDefaults

`func NewAlliancesAllianceIdContactsGetInnerWithDefaults() *AlliancesAllianceIdContactsGetInner`

NewAlliancesAllianceIdContactsGetInnerWithDefaults instantiates a new AlliancesAllianceIdContactsGetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *AlliancesAllianceIdContactsGetInner) GetContactId() int64`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *AlliancesAllianceIdContactsGetInner) GetContactIdOk() (*int64, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *AlliancesAllianceIdContactsGetInner) SetContactId(v int64)`

SetContactId sets ContactId field to given value.


### GetContactType

`func (o *AlliancesAllianceIdContactsGetInner) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *AlliancesAllianceIdContactsGetInner) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *AlliancesAllianceIdContactsGetInner) SetContactType(v string)`

SetContactType sets ContactType field to given value.


### GetLabelIds

`func (o *AlliancesAllianceIdContactsGetInner) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *AlliancesAllianceIdContactsGetInner) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *AlliancesAllianceIdContactsGetInner) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *AlliancesAllianceIdContactsGetInner) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetStanding

`func (o *AlliancesAllianceIdContactsGetInner) GetStanding() float64`

GetStanding returns the Standing field if non-nil, zero value otherwise.

### GetStandingOk

`func (o *AlliancesAllianceIdContactsGetInner) GetStandingOk() (*float64, bool)`

GetStandingOk returns a tuple with the Standing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanding

`func (o *AlliancesAllianceIdContactsGetInner) SetStanding(v float64)`

SetStanding sets Standing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


