# AtlasEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**HomeId** | Pointer to **string** |  | [optional] 
**IsProxy** | Pointer to **bool** |  | [optional] 
**IsIncomplete** | Pointer to **bool** |  | [optional] 
**ProvenanceType** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**RelationshipAttributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Classifications** | Pointer to [**[]AtlasClassification**](AtlasClassification.md) |  | [optional] 
**Meanings** | Pointer to [**[]AtlasTermAssignmentHeader**](AtlasTermAssignmentHeader.md) |  | [optional] 
**CustomAttributes** | Pointer to **map[string]string** |  | [optional] 
**BusinessAttributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**PendingTasks** | Pointer to **[]string** |  | [optional] 
**Proxy** | Pointer to **bool** |  | [optional] 

## Methods

### NewAtlasEntity

`func NewAtlasEntity() *AtlasEntity`

NewAtlasEntity instantiates a new AtlasEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasEntityWithDefaults

`func NewAtlasEntityWithDefaults() *AtlasEntity`

NewAtlasEntityWithDefaults instantiates a new AtlasEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *AtlasEntity) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AtlasEntity) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AtlasEntity) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *AtlasEntity) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *AtlasEntity) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AtlasEntity) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AtlasEntity) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AtlasEntity) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGuid

`func (o *AtlasEntity) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AtlasEntity) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AtlasEntity) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AtlasEntity) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHomeId

`func (o *AtlasEntity) GetHomeId() string`

GetHomeId returns the HomeId field if non-nil, zero value otherwise.

### GetHomeIdOk

`func (o *AtlasEntity) GetHomeIdOk() (*string, bool)`

GetHomeIdOk returns a tuple with the HomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeId

`func (o *AtlasEntity) SetHomeId(v string)`

SetHomeId sets HomeId field to given value.

### HasHomeId

`func (o *AtlasEntity) HasHomeId() bool`

HasHomeId returns a boolean if a field has been set.

### GetIsProxy

`func (o *AtlasEntity) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *AtlasEntity) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *AtlasEntity) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *AtlasEntity) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetIsIncomplete

`func (o *AtlasEntity) GetIsIncomplete() bool`

GetIsIncomplete returns the IsIncomplete field if non-nil, zero value otherwise.

### GetIsIncompleteOk

`func (o *AtlasEntity) GetIsIncompleteOk() (*bool, bool)`

GetIsIncompleteOk returns a tuple with the IsIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncomplete

`func (o *AtlasEntity) SetIsIncomplete(v bool)`

SetIsIncomplete sets IsIncomplete field to given value.

### HasIsIncomplete

`func (o *AtlasEntity) HasIsIncomplete() bool`

HasIsIncomplete returns a boolean if a field has been set.

### GetProvenanceType

`func (o *AtlasEntity) GetProvenanceType() int32`

GetProvenanceType returns the ProvenanceType field if non-nil, zero value otherwise.

### GetProvenanceTypeOk

`func (o *AtlasEntity) GetProvenanceTypeOk() (*int32, bool)`

GetProvenanceTypeOk returns a tuple with the ProvenanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceType

`func (o *AtlasEntity) SetProvenanceType(v int32)`

SetProvenanceType sets ProvenanceType field to given value.

### HasProvenanceType

`func (o *AtlasEntity) HasProvenanceType() bool`

HasProvenanceType returns a boolean if a field has been set.

### GetStatus

`func (o *AtlasEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AtlasEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AtlasEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AtlasEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AtlasEntity) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AtlasEntity) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AtlasEntity) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AtlasEntity) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AtlasEntity) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AtlasEntity) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AtlasEntity) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AtlasEntity) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *AtlasEntity) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AtlasEntity) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AtlasEntity) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AtlasEntity) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AtlasEntity) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AtlasEntity) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AtlasEntity) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AtlasEntity) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *AtlasEntity) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AtlasEntity) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AtlasEntity) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AtlasEntity) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelationshipAttributes

`func (o *AtlasEntity) GetRelationshipAttributes() map[string]interface{}`

GetRelationshipAttributes returns the RelationshipAttributes field if non-nil, zero value otherwise.

### GetRelationshipAttributesOk

`func (o *AtlasEntity) GetRelationshipAttributesOk() (*map[string]interface{}, bool)`

GetRelationshipAttributesOk returns a tuple with the RelationshipAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipAttributes

`func (o *AtlasEntity) SetRelationshipAttributes(v map[string]interface{})`

SetRelationshipAttributes sets RelationshipAttributes field to given value.

### HasRelationshipAttributes

`func (o *AtlasEntity) HasRelationshipAttributes() bool`

HasRelationshipAttributes returns a boolean if a field has been set.

### GetClassifications

`func (o *AtlasEntity) GetClassifications() []AtlasClassification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *AtlasEntity) GetClassificationsOk() (*[]AtlasClassification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *AtlasEntity) SetClassifications(v []AtlasClassification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *AtlasEntity) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetMeanings

`func (o *AtlasEntity) GetMeanings() []AtlasTermAssignmentHeader`

GetMeanings returns the Meanings field if non-nil, zero value otherwise.

### GetMeaningsOk

`func (o *AtlasEntity) GetMeaningsOk() (*[]AtlasTermAssignmentHeader, bool)`

GetMeaningsOk returns a tuple with the Meanings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanings

`func (o *AtlasEntity) SetMeanings(v []AtlasTermAssignmentHeader)`

SetMeanings sets Meanings field to given value.

### HasMeanings

`func (o *AtlasEntity) HasMeanings() bool`

HasMeanings returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *AtlasEntity) GetCustomAttributes() map[string]string`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *AtlasEntity) GetCustomAttributesOk() (*map[string]string, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *AtlasEntity) SetCustomAttributes(v map[string]string)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *AtlasEntity) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### GetBusinessAttributes

`func (o *AtlasEntity) GetBusinessAttributes() map[string]map[string]interface{}`

GetBusinessAttributes returns the BusinessAttributes field if non-nil, zero value otherwise.

### GetBusinessAttributesOk

`func (o *AtlasEntity) GetBusinessAttributesOk() (*map[string]map[string]interface{}, bool)`

GetBusinessAttributesOk returns a tuple with the BusinessAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAttributes

`func (o *AtlasEntity) SetBusinessAttributes(v map[string]map[string]interface{})`

SetBusinessAttributes sets BusinessAttributes field to given value.

### HasBusinessAttributes

`func (o *AtlasEntity) HasBusinessAttributes() bool`

HasBusinessAttributes returns a boolean if a field has been set.

### GetLabels

`func (o *AtlasEntity) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AtlasEntity) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AtlasEntity) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AtlasEntity) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPendingTasks

`func (o *AtlasEntity) GetPendingTasks() []string`

GetPendingTasks returns the PendingTasks field if non-nil, zero value otherwise.

### GetPendingTasksOk

`func (o *AtlasEntity) GetPendingTasksOk() (*[]string, bool)`

GetPendingTasksOk returns a tuple with the PendingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTasks

`func (o *AtlasEntity) SetPendingTasks(v []string)`

SetPendingTasks sets PendingTasks field to given value.

### HasPendingTasks

`func (o *AtlasEntity) HasPendingTasks() bool`

HasPendingTasks returns a boolean if a field has been set.

### GetProxy

`func (o *AtlasEntity) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AtlasEntity) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AtlasEntity) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AtlasEntity) HasProxy() bool`

HasProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


