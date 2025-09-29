# UniverseTypesTypeIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **float64** |  | [optional] 
**Description** | **string** |  | 
**DogmaAttributes** | Pointer to [**[]DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner**](DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner.md) |  | [optional] 
**DogmaEffects** | Pointer to [**[]DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner**](DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner.md) |  | [optional] 
**GraphicId** | Pointer to **int64** |  | [optional] 
**GroupId** | **int64** |  | 
**IconId** | Pointer to **int64** |  | [optional] 
**MarketGroupId** | Pointer to **int64** | This only exists for types that can be put on the market | [optional] 
**Mass** | Pointer to **float64** |  | [optional] 
**Name** | **string** |  | 
**PackagedVolume** | Pointer to **float64** |  | [optional] 
**PortionSize** | Pointer to **int64** |  | [optional] 
**Published** | **bool** |  | 
**Radius** | Pointer to **float64** |  | [optional] 
**TypeId** | **int64** |  | 
**Volume** | Pointer to **float64** |  | [optional] 

## Methods

### NewUniverseTypesTypeIdGet

`func NewUniverseTypesTypeIdGet(description string, groupId int64, name string, published bool, typeId int64, ) *UniverseTypesTypeIdGet`

NewUniverseTypesTypeIdGet instantiates a new UniverseTypesTypeIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseTypesTypeIdGetWithDefaults

`func NewUniverseTypesTypeIdGetWithDefaults() *UniverseTypesTypeIdGet`

NewUniverseTypesTypeIdGetWithDefaults instantiates a new UniverseTypesTypeIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *UniverseTypesTypeIdGet) GetCapacity() float64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *UniverseTypesTypeIdGet) GetCapacityOk() (*float64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *UniverseTypesTypeIdGet) SetCapacity(v float64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *UniverseTypesTypeIdGet) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDescription

`func (o *UniverseTypesTypeIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UniverseTypesTypeIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UniverseTypesTypeIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDogmaAttributes

`func (o *UniverseTypesTypeIdGet) GetDogmaAttributes() []DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner`

GetDogmaAttributes returns the DogmaAttributes field if non-nil, zero value otherwise.

### GetDogmaAttributesOk

`func (o *UniverseTypesTypeIdGet) GetDogmaAttributesOk() (*[]DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner, bool)`

GetDogmaAttributesOk returns a tuple with the DogmaAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDogmaAttributes

`func (o *UniverseTypesTypeIdGet) SetDogmaAttributes(v []DogmaDynamicItemsTypeIdItemIdGetDogmaAttributesInner)`

SetDogmaAttributes sets DogmaAttributes field to given value.

### HasDogmaAttributes

`func (o *UniverseTypesTypeIdGet) HasDogmaAttributes() bool`

HasDogmaAttributes returns a boolean if a field has been set.

### GetDogmaEffects

`func (o *UniverseTypesTypeIdGet) GetDogmaEffects() []DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner`

GetDogmaEffects returns the DogmaEffects field if non-nil, zero value otherwise.

### GetDogmaEffectsOk

`func (o *UniverseTypesTypeIdGet) GetDogmaEffectsOk() (*[]DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner, bool)`

GetDogmaEffectsOk returns a tuple with the DogmaEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDogmaEffects

`func (o *UniverseTypesTypeIdGet) SetDogmaEffects(v []DogmaDynamicItemsTypeIdItemIdGetDogmaEffectsInner)`

SetDogmaEffects sets DogmaEffects field to given value.

### HasDogmaEffects

`func (o *UniverseTypesTypeIdGet) HasDogmaEffects() bool`

HasDogmaEffects returns a boolean if a field has been set.

### GetGraphicId

`func (o *UniverseTypesTypeIdGet) GetGraphicId() int64`

GetGraphicId returns the GraphicId field if non-nil, zero value otherwise.

### GetGraphicIdOk

`func (o *UniverseTypesTypeIdGet) GetGraphicIdOk() (*int64, bool)`

GetGraphicIdOk returns a tuple with the GraphicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicId

`func (o *UniverseTypesTypeIdGet) SetGraphicId(v int64)`

SetGraphicId sets GraphicId field to given value.

### HasGraphicId

`func (o *UniverseTypesTypeIdGet) HasGraphicId() bool`

HasGraphicId returns a boolean if a field has been set.

### GetGroupId

`func (o *UniverseTypesTypeIdGet) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UniverseTypesTypeIdGet) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UniverseTypesTypeIdGet) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetIconId

`func (o *UniverseTypesTypeIdGet) GetIconId() int64`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *UniverseTypesTypeIdGet) GetIconIdOk() (*int64, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *UniverseTypesTypeIdGet) SetIconId(v int64)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *UniverseTypesTypeIdGet) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### GetMarketGroupId

`func (o *UniverseTypesTypeIdGet) GetMarketGroupId() int64`

GetMarketGroupId returns the MarketGroupId field if non-nil, zero value otherwise.

### GetMarketGroupIdOk

`func (o *UniverseTypesTypeIdGet) GetMarketGroupIdOk() (*int64, bool)`

GetMarketGroupIdOk returns a tuple with the MarketGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketGroupId

`func (o *UniverseTypesTypeIdGet) SetMarketGroupId(v int64)`

SetMarketGroupId sets MarketGroupId field to given value.

### HasMarketGroupId

`func (o *UniverseTypesTypeIdGet) HasMarketGroupId() bool`

HasMarketGroupId returns a boolean if a field has been set.

### GetMass

`func (o *UniverseTypesTypeIdGet) GetMass() float64`

GetMass returns the Mass field if non-nil, zero value otherwise.

### GetMassOk

`func (o *UniverseTypesTypeIdGet) GetMassOk() (*float64, bool)`

GetMassOk returns a tuple with the Mass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMass

`func (o *UniverseTypesTypeIdGet) SetMass(v float64)`

SetMass sets Mass field to given value.

### HasMass

`func (o *UniverseTypesTypeIdGet) HasMass() bool`

HasMass returns a boolean if a field has been set.

### GetName

`func (o *UniverseTypesTypeIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniverseTypesTypeIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniverseTypesTypeIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetPackagedVolume

`func (o *UniverseTypesTypeIdGet) GetPackagedVolume() float64`

GetPackagedVolume returns the PackagedVolume field if non-nil, zero value otherwise.

### GetPackagedVolumeOk

`func (o *UniverseTypesTypeIdGet) GetPackagedVolumeOk() (*float64, bool)`

GetPackagedVolumeOk returns a tuple with the PackagedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagedVolume

`func (o *UniverseTypesTypeIdGet) SetPackagedVolume(v float64)`

SetPackagedVolume sets PackagedVolume field to given value.

### HasPackagedVolume

`func (o *UniverseTypesTypeIdGet) HasPackagedVolume() bool`

HasPackagedVolume returns a boolean if a field has been set.

### GetPortionSize

`func (o *UniverseTypesTypeIdGet) GetPortionSize() int64`

GetPortionSize returns the PortionSize field if non-nil, zero value otherwise.

### GetPortionSizeOk

`func (o *UniverseTypesTypeIdGet) GetPortionSizeOk() (*int64, bool)`

GetPortionSizeOk returns a tuple with the PortionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortionSize

`func (o *UniverseTypesTypeIdGet) SetPortionSize(v int64)`

SetPortionSize sets PortionSize field to given value.

### HasPortionSize

`func (o *UniverseTypesTypeIdGet) HasPortionSize() bool`

HasPortionSize returns a boolean if a field has been set.

### GetPublished

`func (o *UniverseTypesTypeIdGet) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *UniverseTypesTypeIdGet) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *UniverseTypesTypeIdGet) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetRadius

`func (o *UniverseTypesTypeIdGet) GetRadius() float64`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *UniverseTypesTypeIdGet) GetRadiusOk() (*float64, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *UniverseTypesTypeIdGet) SetRadius(v float64)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *UniverseTypesTypeIdGet) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetTypeId

`func (o *UniverseTypesTypeIdGet) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *UniverseTypesTypeIdGet) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *UniverseTypesTypeIdGet) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetVolume

`func (o *UniverseTypesTypeIdGet) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *UniverseTypesTypeIdGet) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *UniverseTypesTypeIdGet) SetVolume(v float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *UniverseTypesTypeIdGet) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


