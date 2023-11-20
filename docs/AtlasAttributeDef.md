# AtlasAttributeDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 
**IsOptional** | Pointer to **bool** |  | [optional] 
**Cardinality** | Pointer to **string** |  | [optional] 
**ValuesMinCount** | Pointer to **int32** |  | [optional] 
**ValuesMaxCount** | Pointer to **int32** |  | [optional] 
**IsUnique** | Pointer to **bool** |  | [optional] 
**IsIndexable** | Pointer to **bool** |  | [optional] 
**IncludeInNotification** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SearchWeight** | Pointer to **int32** |  | [optional] 
**IndexType** | Pointer to **string** |  | [optional] 
**Constraints** | Pointer to [**[]AtlasConstraintDef**](AtlasConstraintDef.md) |  | [optional] 
**Options** | Pointer to **map[string]string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewAtlasAttributeDef

`func NewAtlasAttributeDef() *AtlasAttributeDef`

NewAtlasAttributeDef instantiates a new AtlasAttributeDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasAttributeDefWithDefaults

`func NewAtlasAttributeDefWithDefaults() *AtlasAttributeDef`

NewAtlasAttributeDefWithDefaults instantiates a new AtlasAttributeDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AtlasAttributeDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtlasAttributeDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtlasAttributeDef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtlasAttributeDef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeName

`func (o *AtlasAttributeDef) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AtlasAttributeDef) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AtlasAttributeDef) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *AtlasAttributeDef) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetIsOptional

`func (o *AtlasAttributeDef) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *AtlasAttributeDef) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *AtlasAttributeDef) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *AtlasAttributeDef) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetCardinality

`func (o *AtlasAttributeDef) GetCardinality() string`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *AtlasAttributeDef) GetCardinalityOk() (*string, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *AtlasAttributeDef) SetCardinality(v string)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *AtlasAttributeDef) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### GetValuesMinCount

`func (o *AtlasAttributeDef) GetValuesMinCount() int32`

GetValuesMinCount returns the ValuesMinCount field if non-nil, zero value otherwise.

### GetValuesMinCountOk

`func (o *AtlasAttributeDef) GetValuesMinCountOk() (*int32, bool)`

GetValuesMinCountOk returns a tuple with the ValuesMinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesMinCount

`func (o *AtlasAttributeDef) SetValuesMinCount(v int32)`

SetValuesMinCount sets ValuesMinCount field to given value.

### HasValuesMinCount

`func (o *AtlasAttributeDef) HasValuesMinCount() bool`

HasValuesMinCount returns a boolean if a field has been set.

### GetValuesMaxCount

`func (o *AtlasAttributeDef) GetValuesMaxCount() int32`

GetValuesMaxCount returns the ValuesMaxCount field if non-nil, zero value otherwise.

### GetValuesMaxCountOk

`func (o *AtlasAttributeDef) GetValuesMaxCountOk() (*int32, bool)`

GetValuesMaxCountOk returns a tuple with the ValuesMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesMaxCount

`func (o *AtlasAttributeDef) SetValuesMaxCount(v int32)`

SetValuesMaxCount sets ValuesMaxCount field to given value.

### HasValuesMaxCount

`func (o *AtlasAttributeDef) HasValuesMaxCount() bool`

HasValuesMaxCount returns a boolean if a field has been set.

### GetIsUnique

`func (o *AtlasAttributeDef) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *AtlasAttributeDef) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *AtlasAttributeDef) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *AtlasAttributeDef) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetIsIndexable

`func (o *AtlasAttributeDef) GetIsIndexable() bool`

GetIsIndexable returns the IsIndexable field if non-nil, zero value otherwise.

### GetIsIndexableOk

`func (o *AtlasAttributeDef) GetIsIndexableOk() (*bool, bool)`

GetIsIndexableOk returns a tuple with the IsIndexable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndexable

`func (o *AtlasAttributeDef) SetIsIndexable(v bool)`

SetIsIndexable sets IsIndexable field to given value.

### HasIsIndexable

`func (o *AtlasAttributeDef) HasIsIndexable() bool`

HasIsIndexable returns a boolean if a field has been set.

### GetIncludeInNotification

`func (o *AtlasAttributeDef) GetIncludeInNotification() bool`

GetIncludeInNotification returns the IncludeInNotification field if non-nil, zero value otherwise.

### GetIncludeInNotificationOk

`func (o *AtlasAttributeDef) GetIncludeInNotificationOk() (*bool, bool)`

GetIncludeInNotificationOk returns a tuple with the IncludeInNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInNotification

`func (o *AtlasAttributeDef) SetIncludeInNotification(v bool)`

SetIncludeInNotification sets IncludeInNotification field to given value.

### HasIncludeInNotification

`func (o *AtlasAttributeDef) HasIncludeInNotification() bool`

HasIncludeInNotification returns a boolean if a field has been set.

### GetDefaultValue

`func (o *AtlasAttributeDef) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AtlasAttributeDef) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AtlasAttributeDef) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AtlasAttributeDef) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescription

`func (o *AtlasAttributeDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtlasAttributeDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtlasAttributeDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtlasAttributeDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSearchWeight

`func (o *AtlasAttributeDef) GetSearchWeight() int32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *AtlasAttributeDef) GetSearchWeightOk() (*int32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *AtlasAttributeDef) SetSearchWeight(v int32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *AtlasAttributeDef) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetIndexType

`func (o *AtlasAttributeDef) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *AtlasAttributeDef) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *AtlasAttributeDef) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *AtlasAttributeDef) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetConstraints

`func (o *AtlasAttributeDef) GetConstraints() []AtlasConstraintDef`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AtlasAttributeDef) GetConstraintsOk() (*[]AtlasConstraintDef, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AtlasAttributeDef) SetConstraints(v []AtlasConstraintDef)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AtlasAttributeDef) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetOptions

`func (o *AtlasAttributeDef) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AtlasAttributeDef) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AtlasAttributeDef) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AtlasAttributeDef) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDisplayName

`func (o *AtlasAttributeDef) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AtlasAttributeDef) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AtlasAttributeDef) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AtlasAttributeDef) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


