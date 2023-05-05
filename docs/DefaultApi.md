# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncapiParsePut**](DefaultApi.md#AsyncapiParsePut) | **Put** /asyncapi/parse | 
[**AsyncapiPut**](DefaultApi.md#AsyncapiPut) | **Put** /asyncapi | 
[**CreateBusinessMetadata**](DefaultApi.md#CreateBusinessMetadata) | **Post** /catalog/v1/entity/businessmetadata | Bulk API to create multiple business metadata.
[**CreateBusinessMetadataDefs**](DefaultApi.md#CreateBusinessMetadataDefs) | **Post** /catalog/v1/types/businessmetadatadefs | Bulk create API for business metadata definitions.
[**CreateExporter**](DefaultApi.md#CreateExporter) | **Post** /exporters | Create an exporter.
[**CreateOrUpdate**](DefaultApi.md#CreateOrUpdate) | **Post** /catalog/v1/entity | 
[**CreateTagDefs**](DefaultApi.md#CreateTagDefs) | **Post** /catalog/v1/types/tagdefs | Bulk create API for tag definitions.
[**CreateTags**](DefaultApi.md#CreateTags) | **Post** /catalog/v1/entity/tags | Bulk API to create multiple tags.
[**DeleteBusinessMetadata**](DefaultApi.md#DeleteBusinessMetadata) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata/{bmName} | Delete a business metadata on an entity.
[**DeleteBusinessMetadataDef**](DefaultApi.md#DeleteBusinessMetadataDef) | **Delete** /catalog/v1/types/businessmetadatadefs/{bmName} | Delete API for business metadata definition identified by its name.
[**DeleteByUniqueAttributes**](DefaultApi.md#DeleteByUniqueAttributes) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | 
[**DeleteExporter**](DefaultApi.md#DeleteExporter) | **Delete** /exporters/{name} | Delete an exporter.
[**DeleteSchemaVersion**](DefaultApi.md#DeleteSchemaVersion) | **Delete** /subjects/{subject}/versions/{version} | Deletes a specific version of the schema registered under this subject. This only deletes the version and the schema ID remains intact making it still possible to decode data using the schema ID. This API is recommended to be used only in development environments or under extreme circumstances where-in, its required to delete a previously registered schema for compatibility purposes or re-register previously registered schema.
[**DeleteSubject**](DefaultApi.md#DeleteSubject) | **Delete** /subjects/{subject} | Deletes the specified subject and its associated compatibility level if registered. It is recommended to use this API only when a topic needs to be recycled or in development environment.
[**DeleteSubjectConfig**](DefaultApi.md#DeleteSubjectConfig) | **Delete** /config/{subject} | Deletes the specified subject-level compatibility level config and revert to the global default.
[**DeleteSubjectMode**](DefaultApi.md#DeleteSubjectMode) | **Delete** /mode/{subject} | Deletes the specified subject-level mode and revert to the global default.
[**DeleteTag**](DefaultApi.md#DeleteTag) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags/{tagName} | Delete a tag on an entity.
[**DeleteTagDef**](DefaultApi.md#DeleteTagDef) | **Delete** /catalog/v1/types/tagdefs/{tagName} | Delete API for tag definition identified by its name.
[**DeleteTopLevelConfig**](DefaultApi.md#DeleteTopLevelConfig) | **Delete** /config | Delete global compatibility level
[**Get**](DefaultApi.md#Get) | **Get** / | Schema Registry Root Resource
[**GetAllBusinessMetadataDefs**](DefaultApi.md#GetAllBusinessMetadataDefs) | **Get** /catalog/v1/types/businessmetadatadefs | Bulk retrieval API for retrieving business metadata definitions.
[**GetAllTagDefs**](DefaultApi.md#GetAllTagDefs) | **Get** /catalog/v1/types/tagdefs | Bulk retrieval API for retrieving tag definitions.
[**GetBusinessMetadata**](DefaultApi.md#GetBusinessMetadata) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata | Gets the list of business metadata for a given entity represented by a qualified name.
[**GetBusinessMetadataDefByName**](DefaultApi.md#GetBusinessMetadataDefByName) | **Get** /catalog/v1/types/businessmetadatadefs/{bmName} | Get the business metadata definition with the given name.
[**GetByUniqueAttributes**](DefaultApi.md#GetByUniqueAttributes) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | Fetch complete definition of an entity given its type and unique attribute.
[**GetClusterId**](DefaultApi.md#GetClusterId) | **Get** /v1/metadata/id | Get the server metadata
[**GetExporterConfig**](DefaultApi.md#GetExporterConfig) | **Get** /exporters/{name}/config | Get the config for an exporter.
[**GetExporterInfo**](DefaultApi.md#GetExporterInfo) | **Get** /exporters/{name} | Get the info for an exporter.
[**GetExporterStatus**](DefaultApi.md#GetExporterStatus) | **Get** /exporters/{name}/status | Get the status for an exporter.
[**GetExporters**](DefaultApi.md#GetExporters) | **Get** /exporters | Get a list of exporter names.
[**GetMode**](DefaultApi.md#GetMode) | **Get** /mode/{subject} | Get mode for a subject.
[**GetReferencedBy**](DefaultApi.md#GetReferencedBy) | **Get** /subjects/{subject}/versions/{version}/referencedby | Get the schemas that reference the specified schema.
[**GetSchema**](DefaultApi.md#GetSchema) | **Get** /schemas/ids/{id} | Get the schema string identified by the input ID.
[**GetSchemaByVersion**](DefaultApi.md#GetSchemaByVersion) | **Get** /subjects/{subject}/versions/{version} | Get a specific version of the schema registered under this subject.
[**GetSchemaOnly**](DefaultApi.md#GetSchemaOnly) | **Get** /subjects/{subject}/versions/{version}/schema | Get the schema for the specified version of this subject. The unescaped schema only is returned.
[**GetSchemaTypes**](DefaultApi.md#GetSchemaTypes) | **Get** /schemas/types | Get the schema types supported by this registry.
[**GetSchemas**](DefaultApi.md#GetSchemas) | **Get** /schemas | Get the schemas.
[**GetSubjectLevelConfig**](DefaultApi.md#GetSubjectLevelConfig) | **Get** /config/{subject} | Get compatibility level for a subject.
[**GetSubjects**](DefaultApi.md#GetSubjects) | **Get** /schemas/ids/{id}/subjects | Get all the subjects associated with the input ID.
[**GetTagDefByName**](DefaultApi.md#GetTagDefByName) | **Get** /catalog/v1/types/tagdefs/{tagName} | Get the tag definition with the given name.
[**GetTags**](DefaultApi.md#GetTags) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags | Gets the list of classifications for a given entity represented by a qualifed name.
[**GetTopLevelConfig**](DefaultApi.md#GetTopLevelConfig) | **Get** /config | Get global compatibility level
[**GetTopLevelMode**](DefaultApi.md#GetTopLevelMode) | **Get** /mode | Get global mode.
[**GetVersions**](DefaultApi.md#GetVersions) | **Get** /schemas/ids/{id}/versions | Get all the subject-version pairs associated with the input ID.
[**List**](DefaultApi.md#List) | **Get** /subjects | Get a list of registered subjects.
[**ListContexts**](DefaultApi.md#ListContexts) | **Get** /contexts | Get a list of contexts.
[**ListVersions**](DefaultApi.md#ListVersions) | **Get** /subjects/{subject}/versions | Get a list of versions registered under the specified subject.
[**LookUpSchemaUnderSubject**](DefaultApi.md#LookUpSchemaUnderSubject) | **Post** /subjects/{subject} | Check if a schema has already been registered under the specified subject. If so, this returns the schema string along with its globally unique identifier, its version under this subject and the subject name.
[**PartialUpdateByUniqueAttributes**](DefaultApi.md#PartialUpdateByUniqueAttributes) | **Put** /catalog/v1/entity | 
[**PauseExporter**](DefaultApi.md#PauseExporter) | **Put** /exporters/{name}/pause | Pause an exporter.
[**Post**](DefaultApi.md#Post) | **Post** / | 
[**PutExporter**](DefaultApi.md#PutExporter) | **Put** /exporters/{name} | Alters an exporter.
[**PutExporterConfig**](DefaultApi.md#PutExporterConfig) | **Put** /exporters/{name}/config | Alters the config of an exporter.
[**Register**](DefaultApi.md#Register) | **Post** /subjects/{subject}/versions | Register schema under a subject
[**ResetExporter**](DefaultApi.md#ResetExporter) | **Put** /exporters/{name}/reset | Reset an exporter.
[**ResumeExporter**](DefaultApi.md#ResumeExporter) | **Put** /exporters/{name}/resume | Resume an exporter.
[**SearchUsingAttribute**](DefaultApi.md#SearchUsingAttribute) | **Get** /catalog/v1/search/attribute | Retrieve data for the specified attribute search query.
[**SearchUsingBasic**](DefaultApi.md#SearchUsingBasic) | **Get** /catalog/v1/search/basic | Retrieve data for the specified fulltext query.
[**TestCompatibilityBySubjectName**](DefaultApi.md#TestCompatibilityBySubjectName) | **Post** /compatibility/subjects/{subject}/versions/{version} | Test input schema against a particular version of a subject&#39;s schema for compatibility.
[**TestCompatibilityForSubject**](DefaultApi.md#TestCompatibilityForSubject) | **Post** /compatibility/subjects/{subject}/versions | Test input schema against a subject&#39;s schemas for compatibility, based on the compatibility level of the subject configured. In other word, it will perform the same compatibility check as register for that subject
[**UpdateBusinessMetadata**](DefaultApi.md#UpdateBusinessMetadata) | **Put** /catalog/v1/entity/businessmetadata | Bulk API to update multiple business metadata.
[**UpdateBusinessMetadataDefs**](DefaultApi.md#UpdateBusinessMetadataDefs) | **Put** /catalog/v1/types/businessmetadatadefs | Bulk update API for business metadata definitions.
[**UpdateMode**](DefaultApi.md#UpdateMode) | **Put** /mode/{subject} | Update mode for the specified subject.
[**UpdateSubjectLevelConfig**](DefaultApi.md#UpdateSubjectLevelConfig) | **Put** /config/{subject} | Update compatibility level for the specified subject.
[**UpdateTagDefs**](DefaultApi.md#UpdateTagDefs) | **Put** /catalog/v1/types/tagdefs | Bulk update API for tag definitions.
[**UpdateTags**](DefaultApi.md#UpdateTags) | **Put** /catalog/v1/entity/tags | Bulk API to update multiple tags.
[**UpdateTopLevelConfig**](DefaultApi.md#UpdateTopLevelConfig) | **Put** /config | Update global compatibility level
[**UpdateTopLevelMode**](DefaultApi.md#UpdateTopLevelMode) | **Put** /mode | Update global mode.



## AsyncapiParsePut

> AsyncapiParsePut(ctx, )



Get number of times the cli tool is used to import and parse the spec file

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AsyncapiPut

> AsyncapiPut(ctx, )



Get number of times the cli tool is used to export/produce the spec file

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBusinessMetadata

> []BusinessMetadataResponse CreateBusinessMetadata(ctx, optional)

Bulk API to create multiple business metadata.

Bulk API to create multiple business metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateBusinessMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBusinessMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadata** | [**optional.Interface of []BusinessMetadata**](BusinessMetadata.md)| The business metadata | 

### Return type

[**[]BusinessMetadataResponse**](BusinessMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBusinessMetadataDefs

> []BusinessMetadataDefResponse CreateBusinessMetadataDefs(ctx, optional)

Bulk create API for business metadata definitions.

Bulk create API for business metadata definitions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateBusinessMetadataDefsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBusinessMetadataDefsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasBusinessMetadataDef** | [**optional.Interface of []AtlasBusinessMetadataDef**](AtlasBusinessMetadataDef.md)| The business metadata definitions to create | 

### Return type

[**[]BusinessMetadataDefResponse**](BusinessMetadataDefResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExporter

> CreateExporterResponse CreateExporter(ctx, body)

Create an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateExporterRequest**](CreateExporterRequest.md)| Info | 

### Return type

[**CreateExporterResponse**](CreateExporterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdate

> CreateOrUpdate(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateOrUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasEntityWithExtInfo** | [**optional.Interface of AtlasEntityWithExtInfo**](AtlasEntityWithExtInfo.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagDefs

> []TagDefResponse CreateTagDefs(ctx, optional)

Bulk create API for tag definitions.

Bulk create API for tag definitions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTagDefsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTagDefsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDef** | [**optional.Interface of []TagDef**](TagDef.md)| The tag definitions to create | 

### Return type

[**[]TagDefResponse**](TagDefResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTags

> []TagResponse CreateTags(ctx, optional)

Bulk API to create multiple tags.

Bulk API to create multiple tags.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**optional.Interface of []Tag**](Tag.md)| The tags | 

### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessMetadata

> DeleteBusinessMetadata(ctx, typeName, qualifiedName, bmName)

Delete a business metadata on an entity.

Delete a business metadata on an entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string**| The type of the entity | 
**qualifiedName** | **string**| The qualified name of the entity | 
**bmName** | **string**| The name of the business metadata | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessMetadataDef

> DeleteBusinessMetadataDef(ctx, bmName)

Delete API for business metadata definition identified by its name.

Delete API for business metadata definition identified by its name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bmName** | **string**| The name of the business metadata definition | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByUniqueAttributes

> DeleteByUniqueAttributes(ctx, typeName, qualifiedName)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string**|  | 
**qualifiedName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExporter

> DeleteExporter(ctx, name)

Delete an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the exporter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchemaVersion

> int32 DeleteSchemaVersion(ctx, subject, version, optional)

Deletes a specific version of the schema registered under this subject. This only deletes the version and the schema ID remains intact making it still possible to decode data using the schema ID. This API is recommended to be used only in development environments or under extreme circumstances where-in, its required to delete a previously registered schema for compatibility purposes or re-register previously registered schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
**version** | **string**| Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 
 **optional** | ***DeleteSchemaVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSchemaVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permanent** | **optional.Bool**|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubject

> []int32 DeleteSubject(ctx, subject, optional)

Deletes the specified subject and its associated compatibility level if registered. It is recommended to use this API only when a topic needs to be recycled or in development environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| the name of the subject | 
 **optional** | ***DeleteSubjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSubjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permanent** | **optional.Bool**|  | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubjectConfig

> string DeleteSubjectConfig(ctx, subject)

Deletes the specified subject-level compatibility level config and revert to the global default.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| the name of the subject | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubjectMode

> string DeleteSubjectMode(ctx, subject)

Deletes the specified subject-level mode and revert to the global default.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| the name of the subject | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, typeName, qualifiedName, tagName)

Delete a tag on an entity.

Delete a tag on an entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string**| The type of the entity | 
**qualifiedName** | **string**| The qualified name of the entity | 
**tagName** | **string**| The name of the tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagDef

> DeleteTagDef(ctx, tagName)

Delete API for tag definition identified by its name.

Delete API for tag definition identified by its name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string**| The name of the tag definition | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTopLevelConfig

> string DeleteTopLevelConfig(ctx, )

Delete global compatibility level

Deletes the global compatibility level config and reverts to the default.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> map[string]map[string]interface{} Get(ctx, )

Schema Registry Root Resource

The Root resource is a no-op.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBusinessMetadataDefs

> []BusinessMetadataDefResponse GetAllBusinessMetadataDefs(ctx, optional)

Bulk retrieval API for retrieving business metadata definitions.

Bulk retrieval API for retrieving business metadata definitions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllBusinessMetadataDefsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllBusinessMetadataDefsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| The prefix of a business metadata definition name | 

### Return type

[**[]BusinessMetadataDefResponse**](BusinessMetadataDefResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTagDefs

> []TagDefResponse GetAllTagDefs(ctx, optional)

Bulk retrieval API for retrieving tag definitions.

Bulk retrieval API for retrieving tag definitions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTagDefsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllTagDefsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| The prefix of a tag definition name | 

### Return type

[**[]TagDefResponse**](TagDefResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessMetadata

> []BusinessMetadataResponse GetBusinessMetadata(ctx, typeName, qualifiedName)

Gets the list of business metadata for a given entity represented by a qualified name.

Gets the list of business metadata for a given entity represented by a qualified name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string**| The type of the entity | 
**qualifiedName** | **string**| The qualified name of the entity | 

### Return type

[**[]BusinessMetadataResponse**](BusinessMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessMetadataDefByName

> AtlasBusinessMetadataDef GetBusinessMetadataDefByName(ctx, bmName)

Get the business metadata definition with the given name.

Get the business metadata definition with the given name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bmName** | **string**| The name of the business metadata definition | 

### Return type

[**AtlasBusinessMetadataDef**](AtlasBusinessMetadataDef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByUniqueAttributes

> AtlasEntityWithExtInfo GetByUniqueAttributes(ctx, typeName, qualifiedName, optional)

Fetch complete definition of an entity given its type and unique attribute.

Fetch complete definition of an entity given its type and unique attribute.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string**| The type of the entity | 
**qualifiedName** | **string**| The qualified name of the entity | 
 **optional** | ***GetByUniqueAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetByUniqueAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minExtInfo** | **optional.Bool**| Whether to only populate header and schema attributes | [default to false]
 **ignoreRelationships** | **optional.Bool**| Whether to ignore relationships | [default to false]
 **includeInternalPrefix** | **optional.String**| If not null, include internal attributes that start with this prefix | 

### Return type

[**AtlasEntityWithExtInfo**](AtlasEntityWithExtInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterId

> ServerClusterId GetClusterId(ctx, )

Get the server metadata

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServerClusterId**](ServerClusterId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExporterConfig

> map[string]string GetExporterConfig(ctx, name)

Get the config for an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExporterInfo

> ExporterInfo GetExporterInfo(ctx, name)

Get the info for an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

[**ExporterInfo**](ExporterInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExporterStatus

> ExporterStatus GetExporterStatus(ctx, name)

Get the status for an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

[**ExporterStatus**](ExporterStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExporters

> []string GetExporters(ctx, )

Get a list of exporter names.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMode

> Mode GetMode(ctx, subject, optional)

Get mode for a subject.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
 **optional** | ***GetModeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetModeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultToGlobal** | **optional.Bool**|  | 

### Return type

[**Mode**](Mode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferencedBy

> []int32 GetReferencedBy(ctx, subject, version)

Get the schemas that reference the specified schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
**version** | **string**| Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> SchemaString GetSchema(ctx, id, optional)

Get the schema string identified by the input ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Globally unique identifier of the schema | 
 **optional** | ***GetSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSchemaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **optional.String**|  | 
 **format** | **optional.String**|  | 
 **fetchMaxId** | **optional.Bool**|  | [default to false]

### Return type

[**SchemaString**](SchemaString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaByVersion

> Schema GetSchemaByVersion(ctx, subject, version, optional)

Get a specific version of the schema registered under this subject.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
**version** | **string**| Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 
 **optional** | ***GetSchemaByVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSchemaByVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **optional.Bool**|  | 

### Return type

[**Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaOnly

> string GetSchemaOnly(ctx, subject, version, optional)

Get the schema for the specified version of this subject. The unescaped schema only is returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
**version** | **string**| Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 
 **optional** | ***GetSchemaOnlyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSchemaOnlyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **optional.Bool**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaTypes

> []string GetSchemaTypes(ctx, )

Get the schema types supported by this registry.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemas

> []Schema GetSchemas(ctx, optional)

Get the schemas.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSchemasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSchemasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectPrefix** | **optional.String**|  | 
 **deleted** | **optional.Bool**|  | [default to false]
 **latestOnly** | **optional.Bool**|  | [default to false]
 **offset** | **optional.Int32**|  | [default to 0]
 **limit** | **optional.Int32**|  | [default to -1]

### Return type

[**[]Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubjectLevelConfig

> Config GetSubjectLevelConfig(ctx, subject, optional)

Get compatibility level for a subject.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**|  | 
 **optional** | ***GetSubjectLevelConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubjectLevelConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultToGlobal** | **optional.Bool**|  | 

### Return type

[**Config**](Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubjects

> []string GetSubjects(ctx, id, optional)

Get all the subjects associated with the input ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Globally unique identifier of the schema | 
 **optional** | ***GetSubjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubjectsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **optional.String**|  | 
 **deleted** | **optional.Bool**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagDefByName

> TagDef GetTagDefByName(ctx, tagName)

Get the tag definition with the given name.

Get the tag definition with the given name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string**| The name of the tag definiton | 

### Return type

[**TagDef**](TagDef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> []TagResponse GetTags(ctx, typeName, qualifiedName)

Gets the list of classifications for a given entity represented by a qualifed name.

Gets the list of classifications for a given entity represented by a qualifed name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string**| The type of the entity | 
**qualifiedName** | **string**| The qualified name of the entity | 

### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopLevelConfig

> Config GetTopLevelConfig(ctx, )

Get global compatibility level

Retrieves the global compatibility level.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Config**](Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopLevelMode

> Mode GetTopLevelMode(ctx, )

Get global mode.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Mode**](Mode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersions

> []SubjectVersion GetVersions(ctx, id, optional)

Get all the subject-version pairs associated with the input ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Globally unique identifier of the schema | 
 **optional** | ***GetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **optional.String**|  | 
 **deleted** | **optional.Bool**|  | 

### Return type

[**[]SubjectVersion**](SubjectVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []string List(ctx, optional)

Get a list of registered subjects.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectPrefix** | **optional.String**|  | 
 **deleted** | **optional.Bool**|  | 
 **deletedOnly** | **optional.Bool**| Whether to return deleted subjects only | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContexts

> []string ListContexts(ctx, )

Get a list of contexts.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersions

> []int32 ListVersions(ctx, subject, optional)

Get a list of versions registered under the specified subject.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
 **optional** | ***ListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleted** | **optional.Bool**|  | 
 **deletedOnly** | **optional.Bool**| Whether to return deleted schemas only | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookUpSchemaUnderSubject

> LookUpSchemaUnderSubject(ctx, subject, body, optional)

Check if a schema has already been registered under the specified subject. If so, this returns the schema string along with its globally unique identifier, its version under this subject and the subject name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Subject under which the schema will be registered | 
**body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md)| Schema | 
 **optional** | ***LookUpSchemaUnderSubjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LookUpSchemaUnderSubjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateByUniqueAttributes

> PartialUpdateByUniqueAttributes(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartialUpdateByUniqueAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PartialUpdateByUniqueAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasEntityWithExtInfo** | [**optional.Interface of AtlasEntityWithExtInfo**](AtlasEntityWithExtInfo.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseExporter

> UpdateExporterResponse PauseExporter(ctx, name)

Pause an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the exporter | 

### Return type

[**UpdateExporterResponse**](UpdateExporterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> map[string]string Post(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutExporter

> UpdateExporterResponse PutExporter(ctx, name, body)

Alters an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the exporter | 
**body** | [**UpdateExporterRequest**](UpdateExporterRequest.md)| Info | 

### Return type

[**UpdateExporterResponse**](UpdateExporterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutExporterConfig

> UpdateExporterResponse PutExporterConfig(ctx, name, body)

Alters the config of an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the exporter | 
**body** | [**map[string]string**](string.md)| Config | 

### Return type

[**UpdateExporterResponse**](UpdateExporterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> RegisterSchemaResponse Register(ctx, subject, body, optional)

Register schema under a subject

Register a new schema under the specified subject. If successfully registered, this returns the unique identifier of this schema in the registry. The returned identifier should be used to retrieve this schema from the schemas resource and is different from the schema's version which is associated with the subject. If the same schema is registered under a different subject, the same identifier will be returned. However, the version of the schema may be different under different subjects. A schema should be compatible with the previously registered schema or schemas (if there are any) as per the configured compatibility level. The configured compatibility level can be obtained by issuing a GET http:get:: /config/(string: subject). If that returns null, then GET http:get:: /config When there are multiple instances of Schema Registry running in the same cluster, the schema registration request will be forwarded to one of the instances designated as the primary. If the primary is not available, the client will get an error code indicating that the forwarding has failed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the subject | 
**body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md)| Schema | 
 **optional** | ***RegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **normalize** | **optional.Bool**| Whether to register the normalized schema | 

### Return type

[**RegisterSchemaResponse**](RegisterSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetExporter

> UpdateExporterResponse ResetExporter(ctx, name)

Reset an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the exporter | 

### Return type

[**UpdateExporterResponse**](UpdateExporterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeExporter

> UpdateExporterResponse ResumeExporter(ctx, name)

Resume an exporter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the exporter | 

### Return type

[**UpdateExporterResponse**](UpdateExporterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsingAttribute

> SearchResult SearchUsingAttribute(ctx, optional)

Retrieve data for the specified attribute search query.

Retrieve data for the specified attribute search query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchUsingAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUsingAttributeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of []string**](string.md)| Limit the result to only entities of specified types | 
 **types** | **optional.String**|  | 
 **attr** | [**optional.Interface of []string**](string.md)| One of more additional attributes to return in the response | 
 **attrs** | **optional.String**|  | 
 **attrName** | [**optional.Interface of []string**](string.md)| The attribute to search | 
 **attrValuePrefix** | [**optional.Interface of []string**](string.md)| The prefix for the attribute value to search | 
 **tag** | [**optional.Interface of []string**](string.md)| Limit the result to only entities tagged with the given tag | 
 **timeRangeType** | **optional.String**| The type of time range search, default is CUSTOM | 
 **timeRangeAttr** | **optional.String**| The attribute for a time range search | 
 **timeRangeStart** | **optional.Int64**| The start for a custom time range search in ms since the epoch | 
 **timeRangeEnd** | **optional.Int64**| The end for a custom time range search in ms since the epoch | 
 **sortBy** | **optional.String**| An attribute to sort by | 
 **sortOrder** | **optional.String**| Sort order, either ASCENDING (default) or DESCENDING | 
 **deleted** | **optional.Bool**| Whether to include deleted entities | 
 **limit** | **optional.Int32**| Limit the result set to only include the specified number of entries | 
 **offset** | **optional.Int32**| Start offset of the result set (useful for pagination) | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsingBasic

> SearchResult SearchUsingBasic(ctx, optional)

Retrieve data for the specified fulltext query.

Retrieve data for the specified fulltext query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchUsingBasicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUsingBasicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The full-text query | 
 **type_** | [**optional.Interface of []string**](string.md)| Limit the result to only entities of specified types | 
 **types** | **optional.String**|  | 
 **attr** | [**optional.Interface of []string**](string.md)| One of more additional attributes to return in the response | 
 **attrs** | **optional.String**|  | 
 **tag** | [**optional.Interface of []string**](string.md)| Limit the result to only entities with the given tag | 
 **timeRangeType** | **optional.String**| The type of time range search, default is CUSTOM | 
 **timeRangeAttr** | **optional.String**| The attribute for a time range search | 
 **timeRangeStart** | **optional.Int64**| The start for a custom time range search in ms since the epoch | 
 **timeRangeEnd** | **optional.Int64**| The end for a custom time range search in ms since the epoch | 
 **sortBy** | **optional.String**| An attribute to sort by | 
 **sortOrder** | **optional.String**| Sort order, either ASCENDING (default) or DESCENDING | 
 **deleted** | **optional.Bool**| Whether to include deleted entities | 
 **limit** | **optional.Int32**| Limit the result set to only include the specified number of entries | 
 **offset** | **optional.Int32**| Start offset of the result set (useful for pagination) | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCompatibilityBySubjectName

> CompatibilityCheckResponse TestCompatibilityBySubjectName(ctx, subject, version, body, optional)

Test input schema against a particular version of a subject's schema for compatibility.

the compatibility level applied for the check is the configured compatibility level for the subject (http:get:: /config/(string: subject)). If this subject's compatibility level was never changed, then the global compatibility level applies (http:get:: /config).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Subject of the schema version against which compatibility is to be tested | 
**version** | **string**| Version of the subject&#39;s schema against which compatibility is to be tested. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;.\&quot;latest\&quot; checks compatibility of the input schema with the last registered schema under the specified subject | 
**body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md)| Schema | 
 **optional** | ***TestCompatibilityBySubjectNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestCompatibilityBySubjectNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **optional.String**|  | 
 **accept** | **optional.String**|  | 
 **verbose** | **optional.Bool**|  | 

### Return type

[**CompatibilityCheckResponse**](CompatibilityCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCompatibilityForSubject

> CompatibilityCheckResponse TestCompatibilityForSubject(ctx, subject, body, optional)

Test input schema against a subject's schemas for compatibility, based on the compatibility level of the subject configured. In other word, it will perform the same compatibility check as register for that subject

the compatibility level applied for the check is the configured compatibility level for the subject (http:get:: /config/(string: subject)). If this subject's compatibility level was never changed, then the global compatibility level applies (http:get:: /config).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Subject of the schema version against which compatibility is to be tested | 
**body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md)| Schema | 
 **optional** | ***TestCompatibilityForSubjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestCompatibilityForSubjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **optional.String**|  | 
 **accept** | **optional.String**|  | 
 **verbose** | **optional.Bool**|  | 

### Return type

[**CompatibilityCheckResponse**](CompatibilityCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessMetadata

> []BusinessMetadataResponse UpdateBusinessMetadata(ctx, optional)

Bulk API to update multiple business metadata.

Bulk API to update multiple business metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateBusinessMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBusinessMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadata** | [**optional.Interface of []BusinessMetadata**](BusinessMetadata.md)| The business metadata | 

### Return type

[**[]BusinessMetadataResponse**](BusinessMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessMetadataDefs

> []BusinessMetadataDefResponse UpdateBusinessMetadataDefs(ctx, optional)

Bulk update API for business metadata definitions.

Bulk update API for business metadata definitions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateBusinessMetadataDefsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBusinessMetadataDefsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasBusinessMetadataDef** | [**optional.Interface of []AtlasBusinessMetadataDef**](AtlasBusinessMetadataDef.md)| The business metadata definitions to update | 

### Return type

[**[]BusinessMetadataDefResponse**](BusinessMetadataDefResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMode

> ModeUpdateRequest UpdateMode(ctx, subject, body)

Update mode for the specified subject.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
**body** | [**ModeUpdateRequest**](ModeUpdateRequest.md)| Update Request | 

### Return type

[**ModeUpdateRequest**](ModeUpdateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubjectLevelConfig

> ConfigUpdateRequest UpdateSubjectLevelConfig(ctx, subject, body)

Update compatibility level for the specified subject.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string**| Name of the Subject | 
**body** | [**ConfigUpdateRequest**](ConfigUpdateRequest.md)| Config Update Request | 

### Return type

[**ConfigUpdateRequest**](ConfigUpdateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagDefs

> []TagDefResponse UpdateTagDefs(ctx, optional)

Bulk update API for tag definitions.

Bulk update API for tag definitions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateTagDefsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTagDefsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDef** | [**optional.Interface of []TagDef**](TagDef.md)| The tag definitions to update | 

### Return type

[**[]TagDefResponse**](TagDefResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTags

> []TagResponse UpdateTags(ctx, optional)

Bulk API to update multiple tags.

Bulk API to update multiple tags.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**optional.Interface of []Tag**](Tag.md)| The tags | 

### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopLevelConfig

> ConfigUpdateRequest UpdateTopLevelConfig(ctx, configUpdateRequest)

Update global compatibility level

Updates the global compatibility level. On success, echoes the original request back to the client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configUpdateRequest** | [**ConfigUpdateRequest**](ConfigUpdateRequest.md)| Config Update Request | 

### Return type

[**ConfigUpdateRequest**](ConfigUpdateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopLevelMode

> ModeUpdateRequest UpdateTopLevelMode(ctx, body)

Update global mode.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ModeUpdateRequest**](ModeUpdateRequest.md)| Update Request | 

### Return type

[**ModeUpdateRequest**](ModeUpdateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

