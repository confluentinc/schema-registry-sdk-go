# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**SchemaType** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]SchemaReference**](SchemaReference.md) |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**Ruleset** | Pointer to [**NullableRuleSet**](RuleSet.md) |  | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *Schema) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Schema) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Schema) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Schema) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *Schema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Schema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Schema) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Schema) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *Schema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Schema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaType

`func (o *Schema) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *Schema) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *Schema) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *Schema) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetReferences

`func (o *Schema) GetReferences() []SchemaReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Schema) GetReferencesOk() (*[]SchemaReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Schema) SetReferences(v []SchemaReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Schema) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSchema

`func (o *Schema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Schema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Schema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Schema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetMetadata

`func (o *Schema) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Schema) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Schema) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Schema) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Schema) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Schema) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRuleset

`func (o *Schema) GetRuleset() RuleSet`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *Schema) GetRulesetOk() (*RuleSet, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *Schema) SetRuleset(v RuleSet)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *Schema) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### SetRulesetNil

`func (o *Schema) SetRulesetNil(b bool)`

 SetRulesetNil sets the value for Ruleset to be an explicit nil

### UnsetRuleset
`func (o *Schema) UnsetRuleset()`

UnsetRuleset ensures that no value is present for Ruleset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


