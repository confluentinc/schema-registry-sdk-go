# ConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compatibility** | Pointer to **string** | Compatibility Level | [optional] 
**CompatibilityGroup** | Pointer to **string** |  | [optional] 
**DefaultMetadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**OverrideMetadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**DefaultRuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 
**OverrideRuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 

## Methods

### NewConfigUpdateRequest

`func NewConfigUpdateRequest() *ConfigUpdateRequest`

NewConfigUpdateRequest instantiates a new ConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigUpdateRequestWithDefaults

`func NewConfigUpdateRequestWithDefaults() *ConfigUpdateRequest`

NewConfigUpdateRequestWithDefaults instantiates a new ConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibility

`func (o *ConfigUpdateRequest) GetCompatibility() string`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *ConfigUpdateRequest) GetCompatibilityOk() (*string, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *ConfigUpdateRequest) SetCompatibility(v string)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *ConfigUpdateRequest) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### GetCompatibilityGroup

`func (o *ConfigUpdateRequest) GetCompatibilityGroup() string`

GetCompatibilityGroup returns the CompatibilityGroup field if non-nil, zero value otherwise.

### GetCompatibilityGroupOk

`func (o *ConfigUpdateRequest) GetCompatibilityGroupOk() (*string, bool)`

GetCompatibilityGroupOk returns a tuple with the CompatibilityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityGroup

`func (o *ConfigUpdateRequest) SetCompatibilityGroup(v string)`

SetCompatibilityGroup sets CompatibilityGroup field to given value.

### HasCompatibilityGroup

`func (o *ConfigUpdateRequest) HasCompatibilityGroup() bool`

HasCompatibilityGroup returns a boolean if a field has been set.

### GetDefaultMetadata

`func (o *ConfigUpdateRequest) GetDefaultMetadata() Metadata`

GetDefaultMetadata returns the DefaultMetadata field if non-nil, zero value otherwise.

### GetDefaultMetadataOk

`func (o *ConfigUpdateRequest) GetDefaultMetadataOk() (*Metadata, bool)`

GetDefaultMetadataOk returns a tuple with the DefaultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMetadata

`func (o *ConfigUpdateRequest) SetDefaultMetadata(v Metadata)`

SetDefaultMetadata sets DefaultMetadata field to given value.

### HasDefaultMetadata

`func (o *ConfigUpdateRequest) HasDefaultMetadata() bool`

HasDefaultMetadata returns a boolean if a field has been set.

### SetDefaultMetadataNil

`func (o *ConfigUpdateRequest) SetDefaultMetadataNil(b bool)`

 SetDefaultMetadataNil sets the value for DefaultMetadata to be an explicit nil

### UnsetDefaultMetadata
`func (o *ConfigUpdateRequest) UnsetDefaultMetadata()`

UnsetDefaultMetadata ensures that no value is present for DefaultMetadata, not even an explicit nil
### GetOverrideMetadata

`func (o *ConfigUpdateRequest) GetOverrideMetadata() Metadata`

GetOverrideMetadata returns the OverrideMetadata field if non-nil, zero value otherwise.

### GetOverrideMetadataOk

`func (o *ConfigUpdateRequest) GetOverrideMetadataOk() (*Metadata, bool)`

GetOverrideMetadataOk returns a tuple with the OverrideMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideMetadata

`func (o *ConfigUpdateRequest) SetOverrideMetadata(v Metadata)`

SetOverrideMetadata sets OverrideMetadata field to given value.

### HasOverrideMetadata

`func (o *ConfigUpdateRequest) HasOverrideMetadata() bool`

HasOverrideMetadata returns a boolean if a field has been set.

### SetOverrideMetadataNil

`func (o *ConfigUpdateRequest) SetOverrideMetadataNil(b bool)`

 SetOverrideMetadataNil sets the value for OverrideMetadata to be an explicit nil

### UnsetOverrideMetadata
`func (o *ConfigUpdateRequest) UnsetOverrideMetadata()`

UnsetOverrideMetadata ensures that no value is present for OverrideMetadata, not even an explicit nil
### GetDefaultRuleSet

`func (o *ConfigUpdateRequest) GetDefaultRuleSet() RuleSet`

GetDefaultRuleSet returns the DefaultRuleSet field if non-nil, zero value otherwise.

### GetDefaultRuleSetOk

`func (o *ConfigUpdateRequest) GetDefaultRuleSetOk() (*RuleSet, bool)`

GetDefaultRuleSetOk returns a tuple with the DefaultRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRuleSet

`func (o *ConfigUpdateRequest) SetDefaultRuleSet(v RuleSet)`

SetDefaultRuleSet sets DefaultRuleSet field to given value.

### HasDefaultRuleSet

`func (o *ConfigUpdateRequest) HasDefaultRuleSet() bool`

HasDefaultRuleSet returns a boolean if a field has been set.

### SetDefaultRuleSetNil

`func (o *ConfigUpdateRequest) SetDefaultRuleSetNil(b bool)`

 SetDefaultRuleSetNil sets the value for DefaultRuleSet to be an explicit nil

### UnsetDefaultRuleSet
`func (o *ConfigUpdateRequest) UnsetDefaultRuleSet()`

UnsetDefaultRuleSet ensures that no value is present for DefaultRuleSet, not even an explicit nil
### GetOverrideRuleSet

`func (o *ConfigUpdateRequest) GetOverrideRuleSet() RuleSet`

GetOverrideRuleSet returns the OverrideRuleSet field if non-nil, zero value otherwise.

### GetOverrideRuleSetOk

`func (o *ConfigUpdateRequest) GetOverrideRuleSetOk() (*RuleSet, bool)`

GetOverrideRuleSetOk returns a tuple with the OverrideRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRuleSet

`func (o *ConfigUpdateRequest) SetOverrideRuleSet(v RuleSet)`

SetOverrideRuleSet sets OverrideRuleSet field to given value.

### HasOverrideRuleSet

`func (o *ConfigUpdateRequest) HasOverrideRuleSet() bool`

HasOverrideRuleSet returns a boolean if a field has been set.

### SetOverrideRuleSetNil

`func (o *ConfigUpdateRequest) SetOverrideRuleSetNil(b bool)`

 SetOverrideRuleSetNil sets the value for OverrideRuleSet to be an explicit nil

### UnsetOverrideRuleSet
`func (o *ConfigUpdateRequest) UnsetOverrideRuleSet()`

UnsetOverrideRuleSet ensures that no value is present for OverrideRuleSet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


