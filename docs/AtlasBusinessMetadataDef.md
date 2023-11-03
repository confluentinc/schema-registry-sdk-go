# AtlasBusinessMetadataDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TypeVersion** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **map[string]string** |  | [optional] 
**AttributeDefs** | Pointer to [**[]AtlasAttributeDef**](AtlasAttributeDef.md) |  | [optional] 

## Methods

### NewAtlasBusinessMetadataDef

`func NewAtlasBusinessMetadataDef() *AtlasBusinessMetadataDef`

NewAtlasBusinessMetadataDef instantiates a new AtlasBusinessMetadataDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasBusinessMetadataDefWithDefaults

`func NewAtlasBusinessMetadataDefWithDefaults() *AtlasBusinessMetadataDef`

NewAtlasBusinessMetadataDefWithDefaults instantiates a new AtlasBusinessMetadataDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AtlasBusinessMetadataDef) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AtlasBusinessMetadataDef) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AtlasBusinessMetadataDef) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AtlasBusinessMetadataDef) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGuid

`func (o *AtlasBusinessMetadataDef) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AtlasBusinessMetadataDef) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AtlasBusinessMetadataDef) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AtlasBusinessMetadataDef) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AtlasBusinessMetadataDef) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AtlasBusinessMetadataDef) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AtlasBusinessMetadataDef) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AtlasBusinessMetadataDef) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AtlasBusinessMetadataDef) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AtlasBusinessMetadataDef) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AtlasBusinessMetadataDef) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AtlasBusinessMetadataDef) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *AtlasBusinessMetadataDef) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AtlasBusinessMetadataDef) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AtlasBusinessMetadataDef) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AtlasBusinessMetadataDef) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AtlasBusinessMetadataDef) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AtlasBusinessMetadataDef) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AtlasBusinessMetadataDef) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AtlasBusinessMetadataDef) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *AtlasBusinessMetadataDef) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AtlasBusinessMetadataDef) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AtlasBusinessMetadataDef) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AtlasBusinessMetadataDef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *AtlasBusinessMetadataDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtlasBusinessMetadataDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtlasBusinessMetadataDef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtlasBusinessMetadataDef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AtlasBusinessMetadataDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtlasBusinessMetadataDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtlasBusinessMetadataDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtlasBusinessMetadataDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTypeVersion

`func (o *AtlasBusinessMetadataDef) GetTypeVersion() string`

GetTypeVersion returns the TypeVersion field if non-nil, zero value otherwise.

### GetTypeVersionOk

`func (o *AtlasBusinessMetadataDef) GetTypeVersionOk() (*string, bool)`

GetTypeVersionOk returns a tuple with the TypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVersion

`func (o *AtlasBusinessMetadataDef) SetTypeVersion(v string)`

SetTypeVersion sets TypeVersion field to given value.

### HasTypeVersion

`func (o *AtlasBusinessMetadataDef) HasTypeVersion() bool`

HasTypeVersion returns a boolean if a field has been set.

### GetServiceType

`func (o *AtlasBusinessMetadataDef) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AtlasBusinessMetadataDef) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AtlasBusinessMetadataDef) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *AtlasBusinessMetadataDef) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetOptions

`func (o *AtlasBusinessMetadataDef) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AtlasBusinessMetadataDef) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AtlasBusinessMetadataDef) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AtlasBusinessMetadataDef) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAttributeDefs

`func (o *AtlasBusinessMetadataDef) GetAttributeDefs() []AtlasAttributeDef`

GetAttributeDefs returns the AttributeDefs field if non-nil, zero value otherwise.

### GetAttributeDefsOk

`func (o *AtlasBusinessMetadataDef) GetAttributeDefsOk() (*[]AtlasAttributeDef, bool)`

GetAttributeDefsOk returns a tuple with the AttributeDefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDefs

`func (o *AtlasBusinessMetadataDef) SetAttributeDefs(v []AtlasAttributeDef)`

SetAttributeDefs sets AttributeDefs field to given value.

### HasAttributeDefs

`func (o *AtlasBusinessMetadataDef) HasAttributeDefs() bool`

HasAttributeDefs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


