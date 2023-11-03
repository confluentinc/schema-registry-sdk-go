# AtlasTermAssignmentHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermGuid** | Pointer to **string** |  | [optional] 
**RelationGuid** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayText** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Steward** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAtlasTermAssignmentHeader

`func NewAtlasTermAssignmentHeader() *AtlasTermAssignmentHeader`

NewAtlasTermAssignmentHeader instantiates a new AtlasTermAssignmentHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasTermAssignmentHeaderWithDefaults

`func NewAtlasTermAssignmentHeaderWithDefaults() *AtlasTermAssignmentHeader`

NewAtlasTermAssignmentHeaderWithDefaults instantiates a new AtlasTermAssignmentHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermGuid

`func (o *AtlasTermAssignmentHeader) GetTermGuid() string`

GetTermGuid returns the TermGuid field if non-nil, zero value otherwise.

### GetTermGuidOk

`func (o *AtlasTermAssignmentHeader) GetTermGuidOk() (*string, bool)`

GetTermGuidOk returns a tuple with the TermGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermGuid

`func (o *AtlasTermAssignmentHeader) SetTermGuid(v string)`

SetTermGuid sets TermGuid field to given value.

### HasTermGuid

`func (o *AtlasTermAssignmentHeader) HasTermGuid() bool`

HasTermGuid returns a boolean if a field has been set.

### GetRelationGuid

`func (o *AtlasTermAssignmentHeader) GetRelationGuid() string`

GetRelationGuid returns the RelationGuid field if non-nil, zero value otherwise.

### GetRelationGuidOk

`func (o *AtlasTermAssignmentHeader) GetRelationGuidOk() (*string, bool)`

GetRelationGuidOk returns a tuple with the RelationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationGuid

`func (o *AtlasTermAssignmentHeader) SetRelationGuid(v string)`

SetRelationGuid sets RelationGuid field to given value.

### HasRelationGuid

`func (o *AtlasTermAssignmentHeader) HasRelationGuid() bool`

HasRelationGuid returns a boolean if a field has been set.

### GetDescription

`func (o *AtlasTermAssignmentHeader) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtlasTermAssignmentHeader) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtlasTermAssignmentHeader) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtlasTermAssignmentHeader) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayText

`func (o *AtlasTermAssignmentHeader) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *AtlasTermAssignmentHeader) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *AtlasTermAssignmentHeader) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *AtlasTermAssignmentHeader) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.

### GetExpression

`func (o *AtlasTermAssignmentHeader) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *AtlasTermAssignmentHeader) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *AtlasTermAssignmentHeader) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *AtlasTermAssignmentHeader) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AtlasTermAssignmentHeader) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AtlasTermAssignmentHeader) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AtlasTermAssignmentHeader) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AtlasTermAssignmentHeader) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetSteward

`func (o *AtlasTermAssignmentHeader) GetSteward() string`

GetSteward returns the Steward field if non-nil, zero value otherwise.

### GetStewardOk

`func (o *AtlasTermAssignmentHeader) GetStewardOk() (*string, bool)`

GetStewardOk returns a tuple with the Steward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteward

`func (o *AtlasTermAssignmentHeader) SetSteward(v string)`

SetSteward sets Steward field to given value.

### HasSteward

`func (o *AtlasTermAssignmentHeader) HasSteward() bool`

HasSteward returns a boolean if a field has been set.

### GetSource

`func (o *AtlasTermAssignmentHeader) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AtlasTermAssignmentHeader) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AtlasTermAssignmentHeader) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AtlasTermAssignmentHeader) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetConfidence

`func (o *AtlasTermAssignmentHeader) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AtlasTermAssignmentHeader) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AtlasTermAssignmentHeader) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *AtlasTermAssignmentHeader) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetStatus

`func (o *AtlasTermAssignmentHeader) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AtlasTermAssignmentHeader) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AtlasTermAssignmentHeader) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AtlasTermAssignmentHeader) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


