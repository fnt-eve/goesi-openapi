# DogmaEffectsEffectIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisallowAutoRepeat** | Pointer to **bool** |  | [optional] 
**DischargeAttributeId** | Pointer to **int64** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DurationAttributeId** | Pointer to **int64** |  | [optional] 
**EffectCategory** | Pointer to **int64** |  | [optional] 
**EffectId** | **int64** |  | 
**ElectronicChance** | Pointer to **bool** |  | [optional] 
**FalloffAttributeId** | Pointer to **int64** |  | [optional] 
**IconId** | Pointer to **int64** |  | [optional] 
**IsAssistance** | Pointer to **bool** |  | [optional] 
**IsOffensive** | Pointer to **bool** |  | [optional] 
**IsWarpSafe** | Pointer to **bool** |  | [optional] 
**Modifiers** | Pointer to [**[]DogmaEffectsEffectIdGetModifiersInner**](DogmaEffectsEffectIdGetModifiersInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PostExpression** | Pointer to **int64** |  | [optional] 
**PreExpression** | Pointer to **int64** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**RangeAttributeId** | Pointer to **int64** |  | [optional] 
**RangeChance** | Pointer to **bool** |  | [optional] 
**TrackingSpeedAttributeId** | Pointer to **int64** |  | [optional] 

## Methods

### NewDogmaEffectsEffectIdGet

`func NewDogmaEffectsEffectIdGet(effectId int64, ) *DogmaEffectsEffectIdGet`

NewDogmaEffectsEffectIdGet instantiates a new DogmaEffectsEffectIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDogmaEffectsEffectIdGetWithDefaults

`func NewDogmaEffectsEffectIdGetWithDefaults() *DogmaEffectsEffectIdGet`

NewDogmaEffectsEffectIdGetWithDefaults instantiates a new DogmaEffectsEffectIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DogmaEffectsEffectIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DogmaEffectsEffectIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DogmaEffectsEffectIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DogmaEffectsEffectIdGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisallowAutoRepeat

`func (o *DogmaEffectsEffectIdGet) GetDisallowAutoRepeat() bool`

GetDisallowAutoRepeat returns the DisallowAutoRepeat field if non-nil, zero value otherwise.

### GetDisallowAutoRepeatOk

`func (o *DogmaEffectsEffectIdGet) GetDisallowAutoRepeatOk() (*bool, bool)`

GetDisallowAutoRepeatOk returns a tuple with the DisallowAutoRepeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowAutoRepeat

`func (o *DogmaEffectsEffectIdGet) SetDisallowAutoRepeat(v bool)`

SetDisallowAutoRepeat sets DisallowAutoRepeat field to given value.

### HasDisallowAutoRepeat

`func (o *DogmaEffectsEffectIdGet) HasDisallowAutoRepeat() bool`

HasDisallowAutoRepeat returns a boolean if a field has been set.

### GetDischargeAttributeId

`func (o *DogmaEffectsEffectIdGet) GetDischargeAttributeId() int64`

GetDischargeAttributeId returns the DischargeAttributeId field if non-nil, zero value otherwise.

### GetDischargeAttributeIdOk

`func (o *DogmaEffectsEffectIdGet) GetDischargeAttributeIdOk() (*int64, bool)`

GetDischargeAttributeIdOk returns a tuple with the DischargeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDischargeAttributeId

`func (o *DogmaEffectsEffectIdGet) SetDischargeAttributeId(v int64)`

SetDischargeAttributeId sets DischargeAttributeId field to given value.

### HasDischargeAttributeId

`func (o *DogmaEffectsEffectIdGet) HasDischargeAttributeId() bool`

HasDischargeAttributeId returns a boolean if a field has been set.

### GetDisplayName

`func (o *DogmaEffectsEffectIdGet) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DogmaEffectsEffectIdGet) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DogmaEffectsEffectIdGet) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DogmaEffectsEffectIdGet) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationAttributeId

`func (o *DogmaEffectsEffectIdGet) GetDurationAttributeId() int64`

GetDurationAttributeId returns the DurationAttributeId field if non-nil, zero value otherwise.

### GetDurationAttributeIdOk

`func (o *DogmaEffectsEffectIdGet) GetDurationAttributeIdOk() (*int64, bool)`

GetDurationAttributeIdOk returns a tuple with the DurationAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationAttributeId

`func (o *DogmaEffectsEffectIdGet) SetDurationAttributeId(v int64)`

SetDurationAttributeId sets DurationAttributeId field to given value.

### HasDurationAttributeId

`func (o *DogmaEffectsEffectIdGet) HasDurationAttributeId() bool`

HasDurationAttributeId returns a boolean if a field has been set.

### GetEffectCategory

`func (o *DogmaEffectsEffectIdGet) GetEffectCategory() int64`

GetEffectCategory returns the EffectCategory field if non-nil, zero value otherwise.

### GetEffectCategoryOk

`func (o *DogmaEffectsEffectIdGet) GetEffectCategoryOk() (*int64, bool)`

GetEffectCategoryOk returns a tuple with the EffectCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectCategory

`func (o *DogmaEffectsEffectIdGet) SetEffectCategory(v int64)`

SetEffectCategory sets EffectCategory field to given value.

### HasEffectCategory

`func (o *DogmaEffectsEffectIdGet) HasEffectCategory() bool`

HasEffectCategory returns a boolean if a field has been set.

### GetEffectId

`func (o *DogmaEffectsEffectIdGet) GetEffectId() int64`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *DogmaEffectsEffectIdGet) GetEffectIdOk() (*int64, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *DogmaEffectsEffectIdGet) SetEffectId(v int64)`

SetEffectId sets EffectId field to given value.


### GetElectronicChance

`func (o *DogmaEffectsEffectIdGet) GetElectronicChance() bool`

GetElectronicChance returns the ElectronicChance field if non-nil, zero value otherwise.

### GetElectronicChanceOk

`func (o *DogmaEffectsEffectIdGet) GetElectronicChanceOk() (*bool, bool)`

GetElectronicChanceOk returns a tuple with the ElectronicChance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicChance

`func (o *DogmaEffectsEffectIdGet) SetElectronicChance(v bool)`

SetElectronicChance sets ElectronicChance field to given value.

### HasElectronicChance

`func (o *DogmaEffectsEffectIdGet) HasElectronicChance() bool`

HasElectronicChance returns a boolean if a field has been set.

### GetFalloffAttributeId

`func (o *DogmaEffectsEffectIdGet) GetFalloffAttributeId() int64`

GetFalloffAttributeId returns the FalloffAttributeId field if non-nil, zero value otherwise.

### GetFalloffAttributeIdOk

`func (o *DogmaEffectsEffectIdGet) GetFalloffAttributeIdOk() (*int64, bool)`

GetFalloffAttributeIdOk returns a tuple with the FalloffAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalloffAttributeId

`func (o *DogmaEffectsEffectIdGet) SetFalloffAttributeId(v int64)`

SetFalloffAttributeId sets FalloffAttributeId field to given value.

### HasFalloffAttributeId

`func (o *DogmaEffectsEffectIdGet) HasFalloffAttributeId() bool`

HasFalloffAttributeId returns a boolean if a field has been set.

### GetIconId

`func (o *DogmaEffectsEffectIdGet) GetIconId() int64`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *DogmaEffectsEffectIdGet) GetIconIdOk() (*int64, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *DogmaEffectsEffectIdGet) SetIconId(v int64)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *DogmaEffectsEffectIdGet) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### GetIsAssistance

`func (o *DogmaEffectsEffectIdGet) GetIsAssistance() bool`

GetIsAssistance returns the IsAssistance field if non-nil, zero value otherwise.

### GetIsAssistanceOk

`func (o *DogmaEffectsEffectIdGet) GetIsAssistanceOk() (*bool, bool)`

GetIsAssistanceOk returns a tuple with the IsAssistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssistance

`func (o *DogmaEffectsEffectIdGet) SetIsAssistance(v bool)`

SetIsAssistance sets IsAssistance field to given value.

### HasIsAssistance

`func (o *DogmaEffectsEffectIdGet) HasIsAssistance() bool`

HasIsAssistance returns a boolean if a field has been set.

### GetIsOffensive

`func (o *DogmaEffectsEffectIdGet) GetIsOffensive() bool`

GetIsOffensive returns the IsOffensive field if non-nil, zero value otherwise.

### GetIsOffensiveOk

`func (o *DogmaEffectsEffectIdGet) GetIsOffensiveOk() (*bool, bool)`

GetIsOffensiveOk returns a tuple with the IsOffensive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffensive

`func (o *DogmaEffectsEffectIdGet) SetIsOffensive(v bool)`

SetIsOffensive sets IsOffensive field to given value.

### HasIsOffensive

`func (o *DogmaEffectsEffectIdGet) HasIsOffensive() bool`

HasIsOffensive returns a boolean if a field has been set.

### GetIsWarpSafe

`func (o *DogmaEffectsEffectIdGet) GetIsWarpSafe() bool`

GetIsWarpSafe returns the IsWarpSafe field if non-nil, zero value otherwise.

### GetIsWarpSafeOk

`func (o *DogmaEffectsEffectIdGet) GetIsWarpSafeOk() (*bool, bool)`

GetIsWarpSafeOk returns a tuple with the IsWarpSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWarpSafe

`func (o *DogmaEffectsEffectIdGet) SetIsWarpSafe(v bool)`

SetIsWarpSafe sets IsWarpSafe field to given value.

### HasIsWarpSafe

`func (o *DogmaEffectsEffectIdGet) HasIsWarpSafe() bool`

HasIsWarpSafe returns a boolean if a field has been set.

### GetModifiers

`func (o *DogmaEffectsEffectIdGet) GetModifiers() []DogmaEffectsEffectIdGetModifiersInner`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *DogmaEffectsEffectIdGet) GetModifiersOk() (*[]DogmaEffectsEffectIdGetModifiersInner, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *DogmaEffectsEffectIdGet) SetModifiers(v []DogmaEffectsEffectIdGetModifiersInner)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *DogmaEffectsEffectIdGet) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetName

`func (o *DogmaEffectsEffectIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DogmaEffectsEffectIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DogmaEffectsEffectIdGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DogmaEffectsEffectIdGet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostExpression

`func (o *DogmaEffectsEffectIdGet) GetPostExpression() int64`

GetPostExpression returns the PostExpression field if non-nil, zero value otherwise.

### GetPostExpressionOk

`func (o *DogmaEffectsEffectIdGet) GetPostExpressionOk() (*int64, bool)`

GetPostExpressionOk returns a tuple with the PostExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostExpression

`func (o *DogmaEffectsEffectIdGet) SetPostExpression(v int64)`

SetPostExpression sets PostExpression field to given value.

### HasPostExpression

`func (o *DogmaEffectsEffectIdGet) HasPostExpression() bool`

HasPostExpression returns a boolean if a field has been set.

### GetPreExpression

`func (o *DogmaEffectsEffectIdGet) GetPreExpression() int64`

GetPreExpression returns the PreExpression field if non-nil, zero value otherwise.

### GetPreExpressionOk

`func (o *DogmaEffectsEffectIdGet) GetPreExpressionOk() (*int64, bool)`

GetPreExpressionOk returns a tuple with the PreExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreExpression

`func (o *DogmaEffectsEffectIdGet) SetPreExpression(v int64)`

SetPreExpression sets PreExpression field to given value.

### HasPreExpression

`func (o *DogmaEffectsEffectIdGet) HasPreExpression() bool`

HasPreExpression returns a boolean if a field has been set.

### GetPublished

`func (o *DogmaEffectsEffectIdGet) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *DogmaEffectsEffectIdGet) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *DogmaEffectsEffectIdGet) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *DogmaEffectsEffectIdGet) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRangeAttributeId

`func (o *DogmaEffectsEffectIdGet) GetRangeAttributeId() int64`

GetRangeAttributeId returns the RangeAttributeId field if non-nil, zero value otherwise.

### GetRangeAttributeIdOk

`func (o *DogmaEffectsEffectIdGet) GetRangeAttributeIdOk() (*int64, bool)`

GetRangeAttributeIdOk returns a tuple with the RangeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeAttributeId

`func (o *DogmaEffectsEffectIdGet) SetRangeAttributeId(v int64)`

SetRangeAttributeId sets RangeAttributeId field to given value.

### HasRangeAttributeId

`func (o *DogmaEffectsEffectIdGet) HasRangeAttributeId() bool`

HasRangeAttributeId returns a boolean if a field has been set.

### GetRangeChance

`func (o *DogmaEffectsEffectIdGet) GetRangeChance() bool`

GetRangeChance returns the RangeChance field if non-nil, zero value otherwise.

### GetRangeChanceOk

`func (o *DogmaEffectsEffectIdGet) GetRangeChanceOk() (*bool, bool)`

GetRangeChanceOk returns a tuple with the RangeChance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeChance

`func (o *DogmaEffectsEffectIdGet) SetRangeChance(v bool)`

SetRangeChance sets RangeChance field to given value.

### HasRangeChance

`func (o *DogmaEffectsEffectIdGet) HasRangeChance() bool`

HasRangeChance returns a boolean if a field has been set.

### GetTrackingSpeedAttributeId

`func (o *DogmaEffectsEffectIdGet) GetTrackingSpeedAttributeId() int64`

GetTrackingSpeedAttributeId returns the TrackingSpeedAttributeId field if non-nil, zero value otherwise.

### GetTrackingSpeedAttributeIdOk

`func (o *DogmaEffectsEffectIdGet) GetTrackingSpeedAttributeIdOk() (*int64, bool)`

GetTrackingSpeedAttributeIdOk returns a tuple with the TrackingSpeedAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSpeedAttributeId

`func (o *DogmaEffectsEffectIdGet) SetTrackingSpeedAttributeId(v int64)`

SetTrackingSpeedAttributeId sets TrackingSpeedAttributeId field to given value.

### HasTrackingSpeedAttributeId

`func (o *DogmaEffectsEffectIdGet) HasTrackingSpeedAttributeId() bool`

HasTrackingSpeedAttributeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


