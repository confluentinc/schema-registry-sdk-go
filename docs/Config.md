# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**AliasForDeks** | Pointer to **string** |  | [optional] 
**Normalize** | Pointer to **bool** |  | [optional] 
**ValidateFields** | Pointer to **bool** |  | [optional] 
**ValidateNewSchemas** | Pointer to **bool** |  | [optional] 
**ValidateRules** | Pointer to **bool** |  | [optional] 
**CompatibilityLevel** | Pointer to **string** | Compatibility Level | [optional] 
**CompatibilityPolicy** | Pointer to **string** |  | [optional] 
**CompatibilityGroup** | Pointer to **string** |  | [optional] 
**DefaultMetadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**OverrideMetadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**DefaultRuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 
**OverrideRuleSet** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 

## Methods

### NewConfig

`func NewConfig() *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *Config) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Config) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Config) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Config) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAliasForDeks

`func (o *Config) GetAliasForDeks() string`

GetAliasForDeks returns the AliasForDeks field if non-nil, zero value otherwise.

### GetAliasForDeksOk

`func (o *Config) GetAliasForDeksOk() (*string, bool)`

GetAliasForDeksOk returns a tuple with the AliasForDeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasForDeks

`func (o *Config) SetAliasForDeks(v string)`

SetAliasForDeks sets AliasForDeks field to given value.

### HasAliasForDeks

`func (o *Config) HasAliasForDeks() bool`

HasAliasForDeks returns a boolean if a field has been set.

### GetNormalize

`func (o *Config) GetNormalize() bool`

GetNormalize returns the Normalize field if non-nil, zero value otherwise.

### GetNormalizeOk

`func (o *Config) GetNormalizeOk() (*bool, bool)`

GetNormalizeOk returns a tuple with the Normalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalize

`func (o *Config) SetNormalize(v bool)`

SetNormalize sets Normalize field to given value.

### HasNormalize

`func (o *Config) HasNormalize() bool`

HasNormalize returns a boolean if a field has been set.

### GetValidateFields

`func (o *Config) GetValidateFields() bool`

GetValidateFields returns the ValidateFields field if non-nil, zero value otherwise.

### GetValidateFieldsOk

`func (o *Config) GetValidateFieldsOk() (*bool, bool)`

GetValidateFieldsOk returns a tuple with the ValidateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateFields

`func (o *Config) SetValidateFields(v bool)`

SetValidateFields sets ValidateFields field to given value.

### HasValidateFields

`func (o *Config) HasValidateFields() bool`

HasValidateFields returns a boolean if a field has been set.

### GetValidateNewSchemas

`func (o *Config) GetValidateNewSchemas() bool`

GetValidateNewSchemas returns the ValidateNewSchemas field if non-nil, zero value otherwise.

### GetValidateNewSchemasOk

`func (o *Config) GetValidateNewSchemasOk() (*bool, bool)`

GetValidateNewSchemasOk returns a tuple with the ValidateNewSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateNewSchemas

`func (o *Config) SetValidateNewSchemas(v bool)`

SetValidateNewSchemas sets ValidateNewSchemas field to given value.

### HasValidateNewSchemas

`func (o *Config) HasValidateNewSchemas() bool`

HasValidateNewSchemas returns a boolean if a field has been set.

### GetValidateRules

`func (o *Config) GetValidateRules() bool`

GetValidateRules returns the ValidateRules field if non-nil, zero value otherwise.

### GetValidateRulesOk

`func (o *Config) GetValidateRulesOk() (*bool, bool)`

GetValidateRulesOk returns a tuple with the ValidateRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRules

`func (o *Config) SetValidateRules(v bool)`

SetValidateRules sets ValidateRules field to given value.

### HasValidateRules

`func (o *Config) HasValidateRules() bool`

HasValidateRules returns a boolean if a field has been set.

### GetCompatibilityLevel

`func (o *Config) GetCompatibilityLevel() string`

GetCompatibilityLevel returns the CompatibilityLevel field if non-nil, zero value otherwise.

### GetCompatibilityLevelOk

`func (o *Config) GetCompatibilityLevelOk() (*string, bool)`

GetCompatibilityLevelOk returns a tuple with the CompatibilityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityLevel

`func (o *Config) SetCompatibilityLevel(v string)`

SetCompatibilityLevel sets CompatibilityLevel field to given value.

### HasCompatibilityLevel

`func (o *Config) HasCompatibilityLevel() bool`

HasCompatibilityLevel returns a boolean if a field has been set.

### GetCompatibilityPolicy

`func (o *Config) GetCompatibilityPolicy() string`

GetCompatibilityPolicy returns the CompatibilityPolicy field if non-nil, zero value otherwise.

### GetCompatibilityPolicyOk

`func (o *Config) GetCompatibilityPolicyOk() (*string, bool)`

GetCompatibilityPolicyOk returns a tuple with the CompatibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityPolicy

`func (o *Config) SetCompatibilityPolicy(v string)`

SetCompatibilityPolicy sets CompatibilityPolicy field to given value.

### HasCompatibilityPolicy

`func (o *Config) HasCompatibilityPolicy() bool`

HasCompatibilityPolicy returns a boolean if a field has been set.

### GetCompatibilityGroup

`func (o *Config) GetCompatibilityGroup() string`

GetCompatibilityGroup returns the CompatibilityGroup field if non-nil, zero value otherwise.

### GetCompatibilityGroupOk

`func (o *Config) GetCompatibilityGroupOk() (*string, bool)`

GetCompatibilityGroupOk returns a tuple with the CompatibilityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityGroup

`func (o *Config) SetCompatibilityGroup(v string)`

SetCompatibilityGroup sets CompatibilityGroup field to given value.

### HasCompatibilityGroup

`func (o *Config) HasCompatibilityGroup() bool`

HasCompatibilityGroup returns a boolean if a field has been set.

### GetDefaultMetadata

`func (o *Config) GetDefaultMetadata() Metadata`

GetDefaultMetadata returns the DefaultMetadata field if non-nil, zero value otherwise.

### GetDefaultMetadataOk

`func (o *Config) GetDefaultMetadataOk() (*Metadata, bool)`

GetDefaultMetadataOk returns a tuple with the DefaultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMetadata

`func (o *Config) SetDefaultMetadata(v Metadata)`

SetDefaultMetadata sets DefaultMetadata field to given value.

### HasDefaultMetadata

`func (o *Config) HasDefaultMetadata() bool`

HasDefaultMetadata returns a boolean if a field has been set.

### SetDefaultMetadataNil

`func (o *Config) SetDefaultMetadataNil(b bool)`

 SetDefaultMetadataNil sets the value for DefaultMetadata to be an explicit nil

### UnsetDefaultMetadata
`func (o *Config) UnsetDefaultMetadata()`

UnsetDefaultMetadata ensures that no value is present for DefaultMetadata, not even an explicit nil
### GetOverrideMetadata

`func (o *Config) GetOverrideMetadata() Metadata`

GetOverrideMetadata returns the OverrideMetadata field if non-nil, zero value otherwise.

### GetOverrideMetadataOk

`func (o *Config) GetOverrideMetadataOk() (*Metadata, bool)`

GetOverrideMetadataOk returns a tuple with the OverrideMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideMetadata

`func (o *Config) SetOverrideMetadata(v Metadata)`

SetOverrideMetadata sets OverrideMetadata field to given value.

### HasOverrideMetadata

`func (o *Config) HasOverrideMetadata() bool`

HasOverrideMetadata returns a boolean if a field has been set.

### SetOverrideMetadataNil

`func (o *Config) SetOverrideMetadataNil(b bool)`

 SetOverrideMetadataNil sets the value for OverrideMetadata to be an explicit nil

### UnsetOverrideMetadata
`func (o *Config) UnsetOverrideMetadata()`

UnsetOverrideMetadata ensures that no value is present for OverrideMetadata, not even an explicit nil
### GetDefaultRuleSet

`func (o *Config) GetDefaultRuleSet() RuleSet`

GetDefaultRuleSet returns the DefaultRuleSet field if non-nil, zero value otherwise.

### GetDefaultRuleSetOk

`func (o *Config) GetDefaultRuleSetOk() (*RuleSet, bool)`

GetDefaultRuleSetOk returns a tuple with the DefaultRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRuleSet

`func (o *Config) SetDefaultRuleSet(v RuleSet)`

SetDefaultRuleSet sets DefaultRuleSet field to given value.

### HasDefaultRuleSet

`func (o *Config) HasDefaultRuleSet() bool`

HasDefaultRuleSet returns a boolean if a field has been set.

### SetDefaultRuleSetNil

`func (o *Config) SetDefaultRuleSetNil(b bool)`

 SetDefaultRuleSetNil sets the value for DefaultRuleSet to be an explicit nil

### UnsetDefaultRuleSet
`func (o *Config) UnsetDefaultRuleSet()`

UnsetDefaultRuleSet ensures that no value is present for DefaultRuleSet, not even an explicit nil
### GetOverrideRuleSet

`func (o *Config) GetOverrideRuleSet() RuleSet`

GetOverrideRuleSet returns the OverrideRuleSet field if non-nil, zero value otherwise.

### GetOverrideRuleSetOk

`func (o *Config) GetOverrideRuleSetOk() (*RuleSet, bool)`

GetOverrideRuleSetOk returns a tuple with the OverrideRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRuleSet

`func (o *Config) SetOverrideRuleSet(v RuleSet)`

SetOverrideRuleSet sets OverrideRuleSet field to given value.

### HasOverrideRuleSet

`func (o *Config) HasOverrideRuleSet() bool`

HasOverrideRuleSet returns a boolean if a field has been set.

### SetOverrideRuleSetNil

`func (o *Config) SetOverrideRuleSetNil(b bool)`

 SetOverrideRuleSetNil sets the value for OverrideRuleSet to be an explicit nil

### UnsetOverrideRuleSet
`func (o *Config) UnsetOverrideRuleSet()`

UnsetOverrideRuleSet ensures that no value is present for OverrideRuleSet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


