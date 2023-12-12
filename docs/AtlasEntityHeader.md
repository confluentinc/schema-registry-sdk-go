# AtlasEntityHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DisplayText** | Pointer to **string** |  | [optional] 
**ClassificationNames** | Pointer to **[]string** |  | [optional] 
**Classifications** | Pointer to [**[]AtlasClassification**](AtlasClassification.md) |  | [optional] 
**MeaningNames** | Pointer to **[]string** |  | [optional] 
**Meanings** | Pointer to [**[]AtlasTermAssignmentHeader**](AtlasTermAssignmentHeader.md) |  | [optional] 
**IsIncomplete** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAtlasEntityHeader

`func NewAtlasEntityHeader() *AtlasEntityHeader`

NewAtlasEntityHeader instantiates a new AtlasEntityHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasEntityHeaderWithDefaults

`func NewAtlasEntityHeaderWithDefaults() *AtlasEntityHeader`

NewAtlasEntityHeaderWithDefaults instantiates a new AtlasEntityHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *AtlasEntityHeader) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AtlasEntityHeader) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AtlasEntityHeader) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *AtlasEntityHeader) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *AtlasEntityHeader) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AtlasEntityHeader) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AtlasEntityHeader) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AtlasEntityHeader) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGuid

`func (o *AtlasEntityHeader) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AtlasEntityHeader) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AtlasEntityHeader) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AtlasEntityHeader) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetStatus

`func (o *AtlasEntityHeader) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AtlasEntityHeader) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AtlasEntityHeader) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AtlasEntityHeader) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDisplayText

`func (o *AtlasEntityHeader) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *AtlasEntityHeader) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *AtlasEntityHeader) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *AtlasEntityHeader) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.

### GetClassificationNames

`func (o *AtlasEntityHeader) GetClassificationNames() []string`

GetClassificationNames returns the ClassificationNames field if non-nil, zero value otherwise.

### GetClassificationNamesOk

`func (o *AtlasEntityHeader) GetClassificationNamesOk() (*[]string, bool)`

GetClassificationNamesOk returns a tuple with the ClassificationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationNames

`func (o *AtlasEntityHeader) SetClassificationNames(v []string)`

SetClassificationNames sets ClassificationNames field to given value.

### HasClassificationNames

`func (o *AtlasEntityHeader) HasClassificationNames() bool`

HasClassificationNames returns a boolean if a field has been set.

### GetClassifications

`func (o *AtlasEntityHeader) GetClassifications() []AtlasClassification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *AtlasEntityHeader) GetClassificationsOk() (*[]AtlasClassification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *AtlasEntityHeader) SetClassifications(v []AtlasClassification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *AtlasEntityHeader) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetMeaningNames

`func (o *AtlasEntityHeader) GetMeaningNames() []string`

GetMeaningNames returns the MeaningNames field if non-nil, zero value otherwise.

### GetMeaningNamesOk

`func (o *AtlasEntityHeader) GetMeaningNamesOk() (*[]string, bool)`

GetMeaningNamesOk returns a tuple with the MeaningNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeaningNames

`func (o *AtlasEntityHeader) SetMeaningNames(v []string)`

SetMeaningNames sets MeaningNames field to given value.

### HasMeaningNames

`func (o *AtlasEntityHeader) HasMeaningNames() bool`

HasMeaningNames returns a boolean if a field has been set.

### GetMeanings

`func (o *AtlasEntityHeader) GetMeanings() []AtlasTermAssignmentHeader`

GetMeanings returns the Meanings field if non-nil, zero value otherwise.

### GetMeaningsOk

`func (o *AtlasEntityHeader) GetMeaningsOk() (*[]AtlasTermAssignmentHeader, bool)`

GetMeaningsOk returns a tuple with the Meanings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanings

`func (o *AtlasEntityHeader) SetMeanings(v []AtlasTermAssignmentHeader)`

SetMeanings sets Meanings field to given value.

### HasMeanings

`func (o *AtlasEntityHeader) HasMeanings() bool`

HasMeanings returns a boolean if a field has been set.

### GetIsIncomplete

`func (o *AtlasEntityHeader) GetIsIncomplete() bool`

GetIsIncomplete returns the IsIncomplete field if non-nil, zero value otherwise.

### GetIsIncompleteOk

`func (o *AtlasEntityHeader) GetIsIncompleteOk() (*bool, bool)`

GetIsIncompleteOk returns a tuple with the IsIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncomplete

`func (o *AtlasEntityHeader) SetIsIncomplete(v bool)`

SetIsIncomplete sets IsIncomplete field to given value.

### HasIsIncomplete

`func (o *AtlasEntityHeader) HasIsIncomplete() bool`

HasIsIncomplete returns a boolean if a field has been set.

### GetLabels

`func (o *AtlasEntityHeader) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AtlasEntityHeader) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AtlasEntityHeader) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AtlasEntityHeader) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


