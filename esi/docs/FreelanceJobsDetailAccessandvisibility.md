# FreelanceJobsDetailAccessandvisibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProtected** | **bool** | Whether the job is protected by an ACL | 
**BroadcastLocations** | Pointer to [**[]FreelanceJobsDetailBroadcastlocations**](FreelanceJobsDetailBroadcastlocations.md) | Solar systems where the job is broadcast | [optional] 
**Restrictions** | Pointer to [**FreelanceJobsDetailRestrictions**](FreelanceJobsDetailRestrictions.md) |  | [optional] 

## Methods

### NewFreelanceJobsDetailAccessandvisibility

`func NewFreelanceJobsDetailAccessandvisibility(aclProtected bool, ) *FreelanceJobsDetailAccessandvisibility`

NewFreelanceJobsDetailAccessandvisibility instantiates a new FreelanceJobsDetailAccessandvisibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreelanceJobsDetailAccessandvisibilityWithDefaults

`func NewFreelanceJobsDetailAccessandvisibilityWithDefaults() *FreelanceJobsDetailAccessandvisibility`

NewFreelanceJobsDetailAccessandvisibilityWithDefaults instantiates a new FreelanceJobsDetailAccessandvisibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProtected

`func (o *FreelanceJobsDetailAccessandvisibility) GetAclProtected() bool`

GetAclProtected returns the AclProtected field if non-nil, zero value otherwise.

### GetAclProtectedOk

`func (o *FreelanceJobsDetailAccessandvisibility) GetAclProtectedOk() (*bool, bool)`

GetAclProtectedOk returns a tuple with the AclProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProtected

`func (o *FreelanceJobsDetailAccessandvisibility) SetAclProtected(v bool)`

SetAclProtected sets AclProtected field to given value.


### GetBroadcastLocations

`func (o *FreelanceJobsDetailAccessandvisibility) GetBroadcastLocations() []FreelanceJobsDetailBroadcastlocations`

GetBroadcastLocations returns the BroadcastLocations field if non-nil, zero value otherwise.

### GetBroadcastLocationsOk

`func (o *FreelanceJobsDetailAccessandvisibility) GetBroadcastLocationsOk() (*[]FreelanceJobsDetailBroadcastlocations, bool)`

GetBroadcastLocationsOk returns a tuple with the BroadcastLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastLocations

`func (o *FreelanceJobsDetailAccessandvisibility) SetBroadcastLocations(v []FreelanceJobsDetailBroadcastlocations)`

SetBroadcastLocations sets BroadcastLocations field to given value.

### HasBroadcastLocations

`func (o *FreelanceJobsDetailAccessandvisibility) HasBroadcastLocations() bool`

HasBroadcastLocations returns a boolean if a field has been set.

### GetRestrictions

`func (o *FreelanceJobsDetailAccessandvisibility) GetRestrictions() FreelanceJobsDetailRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *FreelanceJobsDetailAccessandvisibility) GetRestrictionsOk() (*FreelanceJobsDetailRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *FreelanceJobsDetailAccessandvisibility) SetRestrictions(v FreelanceJobsDetailRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *FreelanceJobsDetailAccessandvisibility) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


