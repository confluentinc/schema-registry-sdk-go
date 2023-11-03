# AtlasClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**EntityGuid** | Pointer to **string** |  | [optional] 
**EntityStatus** | Pointer to **string** |  | [optional] 
**Propagate** | Pointer to **bool** |  | [optional] 
**ValidityPeriods** | Pointer to [**[]TimeBoundary**](TimeBoundary.md) |  | [optional] 
**RemovePropagationsOnEntityDelete** | Pointer to **bool** |  | [optional] 

## Methods

### NewAtlasClassification

`func NewAtlasClassification() *AtlasClassification`

NewAtlasClassification instantiates a new AtlasClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasClassificationWithDefaults

`func NewAtlasClassificationWithDefaults() *AtlasClassification`

NewAtlasClassificationWithDefaults instantiates a new AtlasClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *AtlasClassification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AtlasClassification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AtlasClassification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *AtlasClassification) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *AtlasClassification) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AtlasClassification) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AtlasClassification) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AtlasClassification) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntityGuid

`func (o *AtlasClassification) GetEntityGuid() string`

GetEntityGuid returns the EntityGuid field if non-nil, zero value otherwise.

### GetEntityGuidOk

`func (o *AtlasClassification) GetEntityGuidOk() (*string, bool)`

GetEntityGuidOk returns a tuple with the EntityGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityGuid

`func (o *AtlasClassification) SetEntityGuid(v string)`

SetEntityGuid sets EntityGuid field to given value.

### HasEntityGuid

`func (o *AtlasClassification) HasEntityGuid() bool`

HasEntityGuid returns a boolean if a field has been set.

### GetEntityStatus

`func (o *AtlasClassification) GetEntityStatus() string`

GetEntityStatus returns the EntityStatus field if non-nil, zero value otherwise.

### GetEntityStatusOk

`func (o *AtlasClassification) GetEntityStatusOk() (*string, bool)`

GetEntityStatusOk returns a tuple with the EntityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityStatus

`func (o *AtlasClassification) SetEntityStatus(v string)`

SetEntityStatus sets EntityStatus field to given value.

### HasEntityStatus

`func (o *AtlasClassification) HasEntityStatus() bool`

HasEntityStatus returns a boolean if a field has been set.

### GetPropagate

`func (o *AtlasClassification) GetPropagate() bool`

GetPropagate returns the Propagate field if non-nil, zero value otherwise.

### GetPropagateOk

`func (o *AtlasClassification) GetPropagateOk() (*bool, bool)`

GetPropagateOk returns a tuple with the Propagate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagate

`func (o *AtlasClassification) SetPropagate(v bool)`

SetPropagate sets Propagate field to given value.

### HasPropagate

`func (o *AtlasClassification) HasPropagate() bool`

HasPropagate returns a boolean if a field has been set.

### GetValidityPeriods

`func (o *AtlasClassification) GetValidityPeriods() []TimeBoundary`

GetValidityPeriods returns the ValidityPeriods field if non-nil, zero value otherwise.

### GetValidityPeriodsOk

`func (o *AtlasClassification) GetValidityPeriodsOk() (*[]TimeBoundary, bool)`

GetValidityPeriodsOk returns a tuple with the ValidityPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriods

`func (o *AtlasClassification) SetValidityPeriods(v []TimeBoundary)`

SetValidityPeriods sets ValidityPeriods field to given value.

### HasValidityPeriods

`func (o *AtlasClassification) HasValidityPeriods() bool`

HasValidityPeriods returns a boolean if a field has been set.

### GetRemovePropagationsOnEntityDelete

`func (o *AtlasClassification) GetRemovePropagationsOnEntityDelete() bool`

GetRemovePropagationsOnEntityDelete returns the RemovePropagationsOnEntityDelete field if non-nil, zero value otherwise.

### GetRemovePropagationsOnEntityDeleteOk

`func (o *AtlasClassification) GetRemovePropagationsOnEntityDeleteOk() (*bool, bool)`

GetRemovePropagationsOnEntityDeleteOk returns a tuple with the RemovePropagationsOnEntityDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePropagationsOnEntityDelete

`func (o *AtlasClassification) SetRemovePropagationsOnEntityDelete(v bool)`

SetRemovePropagationsOnEntityDelete sets RemovePropagationsOnEntityDelete field to given value.

### HasRemovePropagationsOnEntityDelete

`func (o *AtlasClassification) HasRemovePropagationsOnEntityDelete() bool`

HasRemovePropagationsOnEntityDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


