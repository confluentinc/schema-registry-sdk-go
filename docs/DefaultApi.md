# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncapiParsePut**](DefaultApi.md#AsyncapiParsePut) | **Put** /asyncapi/parse | 
[**AsyncapiPut**](DefaultApi.md#AsyncapiPut) | **Put** /asyncapi | 
[**CreateBusinessMetadata**](DefaultApi.md#CreateBusinessMetadata) | **Post** /catalog/v1/entity/businessmetadata | Bulk API to create multiple business metadata.
[**CreateBusinessMetadataDefs**](DefaultApi.md#CreateBusinessMetadataDefs) | **Post** /catalog/v1/types/businessmetadatadefs | Bulk create API for business metadata definitions.
[**CreateDek**](DefaultApi.md#CreateDek) | **Post** /dek-registry/v1/keks/{name}/deks | Create a dek.
[**CreateExporter**](DefaultApi.md#CreateExporter) | **Post** /exporters | Create an exporter.
[**CreateKek**](DefaultApi.md#CreateKek) | **Post** /dek-registry/v1/keks | Create a kek.
[**CreateOrUpdate**](DefaultApi.md#CreateOrUpdate) | **Post** /catalog/v1/entity | 
[**CreateTagDefs**](DefaultApi.md#CreateTagDefs) | **Post** /catalog/v1/types/tagdefs | Bulk create API for tag definitions.
[**CreateTags**](DefaultApi.md#CreateTags) | **Post** /catalog/v1/entity/tags | Bulk API to create multiple tags.
[**DeleteBusinessMetadata**](DefaultApi.md#DeleteBusinessMetadata) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata/{bmName} | Delete a business metadata on an entity.
[**DeleteBusinessMetadataDef**](DefaultApi.md#DeleteBusinessMetadataDef) | **Delete** /catalog/v1/types/businessmetadatadefs/{bmName} | Delete API for business metadata definition identified by its name.
[**DeleteByUniqueAttributes**](DefaultApi.md#DeleteByUniqueAttributes) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | 
[**DeleteDekVersion**](DefaultApi.md#DeleteDekVersion) | **Delete** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version} | Delete a dek version.
[**DeleteDekVersions**](DefaultApi.md#DeleteDekVersions) | **Delete** /dek-registry/v1/keks/{name}/deks/{subject} | Delete all versions of a dek.
[**DeleteExporter**](DefaultApi.md#DeleteExporter) | **Delete** /exporters/{name} | Delete an exporter.
[**DeleteKek**](DefaultApi.md#DeleteKek) | **Delete** /dek-registry/v1/keks/{name} | Delete a kek.
[**DeleteMode**](DefaultApi.md#DeleteMode) | **Delete** /mode | Deletes the global mode and revert to the default.
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
[**GetDek**](DefaultApi.md#GetDek) | **Get** /dek-registry/v1/keks/{name}/deks/{subject} | Get a dek by subject.
[**GetDekByVersion**](DefaultApi.md#GetDekByVersion) | **Get** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version} | Get a dek by subject and version.
[**GetDekSubjects**](DefaultApi.md#GetDekSubjects) | **Get** /dek-registry/v1/keks/{name}/deks | Get a list of dek subjects.
[**GetDekVersions**](DefaultApi.md#GetDekVersions) | **Get** /dek-registry/v1/keks/{name}/deks/{subject}/versions | List versions of dek.
[**GetExporterConfig**](DefaultApi.md#GetExporterConfig) | **Get** /exporters/{name}/config | Get the config for an exporter.
[**GetExporterInfo**](DefaultApi.md#GetExporterInfo) | **Get** /exporters/{name} | Get the info for an exporter.
[**GetExporterStatus**](DefaultApi.md#GetExporterStatus) | **Get** /exporters/{name}/status | Get the status for an exporter.
[**GetExporters**](DefaultApi.md#GetExporters) | **Get** /exporters | Get a list of exporter names.
[**GetKek**](DefaultApi.md#GetKek) | **Get** /dek-registry/v1/keks/{name} | Get a kek by name.
[**GetKekNames**](DefaultApi.md#GetKekNames) | **Get** /dek-registry/v1/keks | Get a list of kek names.
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
[**PutKek**](DefaultApi.md#PutKek) | **Put** /dek-registry/v1/keks/{name} | Alters a kek.
[**Register**](DefaultApi.md#Register) | **Post** /subjects/{subject}/versions | Register schema under a subject
[**ResetExporter**](DefaultApi.md#ResetExporter) | **Put** /exporters/{name}/reset | Reset an exporter.
[**ResumeExporter**](DefaultApi.md#ResumeExporter) | **Put** /exporters/{name}/resume | Resume an exporter.
[**SearchUsingAttribute**](DefaultApi.md#SearchUsingAttribute) | **Get** /catalog/v1/search/attribute | Retrieve data for the specified attribute search query.
[**SearchUsingBasic**](DefaultApi.md#SearchUsingBasic) | **Get** /catalog/v1/search/basic | Retrieve data for the specified fulltext query.
[**TestCompatibilityBySubjectName**](DefaultApi.md#TestCompatibilityBySubjectName) | **Post** /compatibility/subjects/{subject}/versions/{version} | Test input schema against a particular version of a subject&#39;s schema for compatibility.
[**TestCompatibilityForSubject**](DefaultApi.md#TestCompatibilityForSubject) | **Post** /compatibility/subjects/{subject}/versions | Test input schema against a subject&#39;s schemas for compatibility, based on the compatibility level of the subject configured. In other word, it will perform the same compatibility check as register for that subject
[**UndeleteDekVersion**](DefaultApi.md#UndeleteDekVersion) | **Post** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version}/undelete | Undelete a dek version.
[**UndeleteDekVersions**](DefaultApi.md#UndeleteDekVersions) | **Post** /dek-registry/v1/keks/{name}/deks/{subject}/undelete | Undelete all versions of a dek.
[**UndeleteKek**](DefaultApi.md#UndeleteKek) | **Post** /dek-registry/v1/keks/{name}/undelete | Undelete a kek.
[**UpdateBusinessMetadata**](DefaultApi.md#UpdateBusinessMetadata) | **Put** /catalog/v1/entity/businessmetadata | Bulk API to update multiple business metadata.
[**UpdateBusinessMetadataDefs**](DefaultApi.md#UpdateBusinessMetadataDefs) | **Put** /catalog/v1/types/businessmetadatadefs | Bulk update API for business metadata definitions.
[**UpdateMode**](DefaultApi.md#UpdateMode) | **Put** /mode/{subject} | Update mode for the specified subject.
[**UpdateSubjectLevelConfig**](DefaultApi.md#UpdateSubjectLevelConfig) | **Put** /config/{subject} | Update compatibility level for the specified subject.
[**UpdateTagDefs**](DefaultApi.md#UpdateTagDefs) | **Put** /catalog/v1/types/tagdefs | Bulk update API for tag definitions.
[**UpdateTags**](DefaultApi.md#UpdateTags) | **Put** /catalog/v1/entity/tags | Bulk API to update multiple tags.
[**UpdateTopLevelConfig**](DefaultApi.md#UpdateTopLevelConfig) | **Put** /config | Update global compatibility level
[**UpdateTopLevelMode**](DefaultApi.md#UpdateTopLevelMode) | **Put** /mode | Update global mode.



## AsyncapiParsePut

> AsyncapiParsePut(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AsyncapiParsePut(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AsyncapiParsePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncapiParsePutRequest struct via the builder pattern


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

> AsyncapiPut(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AsyncapiPut(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AsyncapiPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncapiPutRequest struct via the builder pattern


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

> []BusinessMetadataResponse CreateBusinessMetadata(ctx).BusinessMetadata(businessMetadata).Execute()

Bulk API to create multiple business metadata.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    businessMetadata := []openapiclient.BusinessMetadata{*openapiclient.NewBusinessMetadata()} // []BusinessMetadata | The business metadata (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateBusinessMetadata(context.Background()).BusinessMetadata(businessMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBusinessMetadata`: []BusinessMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBusinessMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadata** | [**[]BusinessMetadata**](BusinessMetadata.md) | The business metadata | 

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

> []BusinessMetadataDefResponse CreateBusinessMetadataDefs(ctx).AtlasBusinessMetadataDef(atlasBusinessMetadataDef).Execute()

Bulk create API for business metadata definitions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    atlasBusinessMetadataDef := []openapiclient.AtlasBusinessMetadataDef{*openapiclient.NewAtlasBusinessMetadataDef()} // []AtlasBusinessMetadataDef | The business metadata definitions to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateBusinessMetadataDefs(context.Background()).AtlasBusinessMetadataDef(atlasBusinessMetadataDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBusinessMetadataDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBusinessMetadataDefs`: []BusinessMetadataDefResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBusinessMetadataDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBusinessMetadataDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasBusinessMetadataDef** | [**[]AtlasBusinessMetadataDef**](AtlasBusinessMetadataDef.md) | The business metadata definitions to create | 

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


## CreateDek

> Dek CreateDek(ctx, name).CreateDekRequest(createDekRequest).Execute()

Create a dek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    createDekRequest := *openapiclient.NewCreateDekRequest() // CreateDekRequest | The create request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateDek(context.Background(), name).CreateDekRequest(createDekRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDek`: Dek
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDekRequest** | [**CreateDekRequest**](CreateDekRequest.md) | The create request | 

### Return type

[**Dek**](Dek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExporter

> CreateExporterResponse CreateExporter(ctx).Body(body).Execute()

Create an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCreateExporterRequest() // CreateExporterRequest | Info

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateExporter(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExporter`: CreateExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateExporter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateExporterRequest**](CreateExporterRequest.md) | Info | 

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


## CreateKek

> Kek CreateKek(ctx).CreateKekRequest(createKekRequest).Execute()

Create a kek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createKekRequest := *openapiclient.NewCreateKekRequest() // CreateKekRequest | The create request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateKek(context.Background()).CreateKekRequest(createKekRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKek`: Kek
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateKek`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKekRequest** | [**CreateKekRequest**](CreateKekRequest.md) | The create request | 

### Return type

[**Kek**](Kek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdate

> CreateOrUpdate(ctx).AtlasEntityWithExtInfo(atlasEntityWithExtInfo).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    atlasEntityWithExtInfo := *openapiclient.NewAtlasEntityWithExtInfo() // AtlasEntityWithExtInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateOrUpdate(context.Background()).AtlasEntityWithExtInfo(atlasEntityWithExtInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasEntityWithExtInfo** | [**AtlasEntityWithExtInfo**](AtlasEntityWithExtInfo.md) |  | 

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

> []TagDefResponse CreateTagDefs(ctx).TagDef(tagDef).Execute()

Bulk create API for tag definitions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagDef := []openapiclient.TagDef{*openapiclient.NewTagDef()} // []TagDef | The tag definitions to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTagDefs(context.Background()).TagDef(tagDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTagDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagDefs`: []TagDefResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTagDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDef** | [**[]TagDef**](TagDef.md) | The tag definitions to create | 

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

> []TagResponse CreateTags(ctx).Tag(tag).Execute()

Bulk API to create multiple tags.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | The tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTags(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**[]Tag**](Tag.md) | The tags | 

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

> DeleteBusinessMetadata(ctx, typeName, qualifiedName, bmName).Execute()

Delete a business metadata on an entity.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    bmName := "bmName_example" // string | The name of the business metadata

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteBusinessMetadata(context.Background(), typeName, qualifiedName, bmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 
**bmName** | **string** | The name of the business metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> DeleteBusinessMetadataDef(ctx, bmName).Execute()

Delete API for business metadata definition identified by its name.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bmName := "bmName_example" // string | The name of the business metadata definition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteBusinessMetadataDef(context.Background(), bmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBusinessMetadataDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bmName** | **string** | The name of the business metadata definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessMetadataDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> DeleteByUniqueAttributes(ctx, typeName, qualifiedName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeName := "typeName_example" // string | 
    qualifiedName := "qualifiedName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteByUniqueAttributes(context.Background(), typeName, qualifiedName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteByUniqueAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** |  | 
**qualifiedName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByUniqueAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteDekVersion

> DeleteDekVersion(ctx, name, subject, version).Algorithm(algorithm).Permanent(permanent).Execute()

Delete a dek version.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    subject := "subject_example" // string | Subject of the dek
    version := "version_example" // string | Version of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteDekVersion(context.Background(), name, subject, version).Algorithm(algorithm).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDekVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 
**version** | **string** | Version of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDekVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **algorithm** | **string** | Algorithm of the dek | 
 **permanent** | **bool** | Whether to perform a permanent delete | 

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


## DeleteDekVersions

> DeleteDekVersions(ctx, name, subject).Algorithm(algorithm).Permanent(permanent).Execute()

Delete all versions of a dek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteDekVersions(context.Background(), name, subject).Algorithm(algorithm).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDekVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDekVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 
 **permanent** | **bool** | Whether to perform a permanent delete | 

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


## DeleteExporter

> DeleteExporter(ctx, name).Execute()

Delete an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteExporter(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteKek

> DeleteKek(ctx, name).Permanent(permanent).Execute()

Delete a kek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    permanent := true // bool | Whether to perform a permanent delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteKek(context.Background(), name).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permanent** | **bool** | Whether to perform a permanent delete | 

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


## DeleteMode

> string DeleteMode(ctx).Recursive(recursive).Execute()

Deletes the global mode and revert to the default.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    recursive := true // bool | recursive delete mode across all subjects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteMode(context.Background()).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMode`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recursive** | **bool** | recursive delete mode across all subjects. | 

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


## DeleteSchemaVersion

> int32 DeleteSchemaVersion(ctx, subject, version).Permanent(permanent).Execute()

Deletes a specific version of the schema registered under this subject. This only deletes the version and the schema ID remains intact making it still possible to decode data using the schema ID. This API is recommended to be used only in development environments or under extreme circumstances where-in, its required to delete a previously registered schema for compatibility purposes or re-register previously registered schema.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.
    permanent := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSchemaVersion(context.Background(), subject, version).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchemaVersion`: int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSchemaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchemaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permanent** | **bool** |  | 

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

> []int32 DeleteSubject(ctx, subject).Permanent(permanent).Execute()

Deletes the specified subject and its associated compatibility level if registered. It is recommended to use this API only when a topic needs to be recycled or in development environment.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | the name of the subject
    permanent := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSubject(context.Background(), subject).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubject`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | the name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permanent** | **bool** |  | 

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

> string DeleteSubjectConfig(ctx, subject).Execute()

Deletes the specified subject-level compatibility level config and revert to the global default.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | the name of the subject

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSubjectConfig(context.Background(), subject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSubjectConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubjectConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSubjectConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | the name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> string DeleteSubjectMode(ctx, subject).Recursive(recursive).Execute()

Deletes the specified subject-level mode and revert to the global default.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | the name of the subject
    recursive := true // bool | recursive delete mode for all subjects under the context if subject parameter is a context (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSubjectMode(context.Background(), subject).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSubjectMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubjectMode`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSubjectMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | the name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **bool** | recursive delete mode for all subjects under the context if subject parameter is a context | 

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

> DeleteTag(ctx, typeName, qualifiedName, tagName).Execute()

Delete a tag on an entity.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    tagName := "tagName_example" // string | The name of the tag

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTag(context.Background(), typeName, qualifiedName, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 
**tagName** | **string** | The name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> DeleteTagDef(ctx, tagName).Execute()

Delete API for tag definition identified by its name.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagName := "tagName_example" // string | The name of the tag definition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTagDef(context.Background(), tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTagDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | The name of the tag definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> string DeleteTopLevelConfig(ctx).Execute()

Delete global compatibility level



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTopLevelConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTopLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTopLevelConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTopLevelConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopLevelConfigRequest struct via the builder pattern


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

> map[string]map[string]interface{} Get(ctx).Execute()

Schema Registry Root Resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


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

> []BusinessMetadataDefResponse GetAllBusinessMetadataDefs(ctx).Prefix(prefix).Execute()

Bulk retrieval API for retrieving business metadata definitions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string | The prefix of a business metadata definition name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAllBusinessMetadataDefs(context.Background()).Prefix(prefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllBusinessMetadataDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBusinessMetadataDefs`: []BusinessMetadataDefResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllBusinessMetadataDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBusinessMetadataDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | The prefix of a business metadata definition name | 

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

> []TagDefResponse GetAllTagDefs(ctx).Prefix(prefix).Execute()

Bulk retrieval API for retrieving tag definitions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string | The prefix of a tag definition name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAllTagDefs(context.Background()).Prefix(prefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllTagDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTagDefs`: []TagDefResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllTagDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTagDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | The prefix of a tag definition name | 

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

> []BusinessMetadataResponse GetBusinessMetadata(ctx, typeName, qualifiedName).Execute()

Gets the list of business metadata for a given entity represented by a qualified name.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetBusinessMetadata(context.Background(), typeName, qualifiedName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessMetadata`: []BusinessMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBusinessMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> AtlasBusinessMetadataDef GetBusinessMetadataDefByName(ctx, bmName).Execute()

Get the business metadata definition with the given name.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bmName := "bmName_example" // string | The name of the business metadata definition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetBusinessMetadataDefByName(context.Background(), bmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBusinessMetadataDefByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessMetadataDefByName`: AtlasBusinessMetadataDef
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBusinessMetadataDefByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bmName** | **string** | The name of the business metadata definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessMetadataDefByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> AtlasEntityWithExtInfo GetByUniqueAttributes(ctx, typeName, qualifiedName).MinExtInfo(minExtInfo).IgnoreRelationships(ignoreRelationships).IncludeInternalPrefix(includeInternalPrefix).Execute()

Fetch complete definition of an entity given its type and unique attribute.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    minExtInfo := true // bool | Whether to only populate header and schema attributes (optional) (default to false)
    ignoreRelationships := true // bool | Whether to ignore relationships (optional) (default to false)
    includeInternalPrefix := "includeInternalPrefix_example" // string | If not null, include internal attributes that start with this prefix (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetByUniqueAttributes(context.Background(), typeName, qualifiedName).MinExtInfo(minExtInfo).IgnoreRelationships(ignoreRelationships).IncludeInternalPrefix(includeInternalPrefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetByUniqueAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByUniqueAttributes`: AtlasEntityWithExtInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetByUniqueAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByUniqueAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minExtInfo** | **bool** | Whether to only populate header and schema attributes | [default to false]
 **ignoreRelationships** | **bool** | Whether to ignore relationships | [default to false]
 **includeInternalPrefix** | **string** | If not null, include internal attributes that start with this prefix | 

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

> ServerClusterId GetClusterId(ctx).Execute()

Get the server metadata

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetClusterId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterId`: ServerClusterId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClusterId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterIdRequest struct via the builder pattern


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


## GetDek

> Dek GetDek(ctx, name, subject).Algorithm(algorithm).Deleted(deleted).Execute()

Get a dek by subject.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDek(context.Background(), name, subject).Algorithm(algorithm).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDek`: Dek
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 
 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

[**Dek**](Dek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDekByVersion

> Dek GetDekByVersion(ctx, name, subject, version).Algorithm(algorithm).Deleted(deleted).Execute()

Get a dek by subject and version.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    subject := "subject_example" // string | Subject of the dek
    version := "version_example" // string | Version of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDekByVersion(context.Background(), name, subject, version).Algorithm(algorithm).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDekByVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDekByVersion`: Dek
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDekByVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 
**version** | **string** | Version of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekByVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **algorithm** | **string** | Algorithm of the dek | 
 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

[**Dek**](Dek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDekSubjects

> []string GetDekSubjects(ctx, name).Deleted(deleted).Execute()

Get a list of dek subjects.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDekSubjects(context.Background(), name).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDekSubjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDekSubjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDekSubjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleted** | **bool** | Whether to include deleted keys | 

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


## GetDekVersions

> []int32 GetDekVersions(ctx, name, subject).Algorithm(algorithm).Deleted(deleted).Execute()

List versions of dek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDekVersions(context.Background(), name, subject).Algorithm(algorithm).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDekVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDekVersions`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDekVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDekVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 
 **deleted** | **bool** | Whether to include deleted keys | 

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


## GetExporterConfig

> map[string]string GetExporterConfig(ctx, name).Execute()

Get the config for an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetExporterConfig(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetExporterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExporterConfig`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetExporterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExporterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ExporterInfo GetExporterInfo(ctx, name).Execute()

Get the info for an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetExporterInfo(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetExporterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExporterInfo`: ExporterInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetExporterInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExporterInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ExporterStatus GetExporterStatus(ctx, name).Execute()

Get the status for an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetExporterStatus(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetExporterStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExporterStatus`: ExporterStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetExporterStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExporterStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []string GetExporters(ctx).Execute()

Get a list of exporter names.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetExporters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetExporters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExporters`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetExporters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportersRequest struct via the builder pattern


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


## GetKek

> Kek GetKek(ctx, name).Deleted(deleted).Execute()

Get a kek by name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKek(context.Background(), name).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKek`: Kek
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleted** | **bool** | Whether to include deleted keys | 

### Return type

[**Kek**](Kek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKekNames

> []string GetKekNames(ctx).Deleted(deleted).Execute()

Get a list of kek names.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleted := true // bool | Whether to include deleted keys (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKekNames(context.Background()).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKekNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKekNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKekNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKekNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleted** | **bool** | Whether to include deleted keys | 

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

> Mode GetMode(ctx, subject).DefaultToGlobal(defaultToGlobal).Execute()

Get mode for a subject.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    defaultToGlobal := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMode(context.Background(), subject).DefaultToGlobal(defaultToGlobal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMode`: Mode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultToGlobal** | **bool** |  | 

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

> []int32 GetReferencedBy(ctx, subject, version).Execute()

Get the schemas that reference the specified schema.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetReferencedBy(context.Background(), subject, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetReferencedBy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferencedBy`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetReferencedBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferencedByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> SchemaString GetSchema(ctx, id).Subject(subject).Format(format).FetchMaxId(fetchMaxId).Execute()

Get the schema string identified by the input ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | Globally unique identifier of the schema
    subject := "subject_example" // string |  (optional)
    format := "format_example" // string |  (optional)
    fetchMaxId := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSchema(context.Background(), id).Subject(subject).Format(format).FetchMaxId(fetchMaxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: SchemaString
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Globally unique identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** |  | 
 **format** | **string** |  | 
 **fetchMaxId** | **bool** |  | [default to false]

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

> Schema GetSchemaByVersion(ctx, subject, version).Deleted(deleted).Execute()

Get a specific version of the schema registered under this subject.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.
    deleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSchemaByVersion(context.Background(), subject, version).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSchemaByVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaByVersion`: Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSchemaByVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaByVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **bool** |  | 

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

> string GetSchemaOnly(ctx, subject, version).Deleted(deleted).Execute()

Get the schema for the specified version of this subject. The unescaped schema only is returned.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    version := "version_example" // string | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \"latest\". \"latest\" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.
    deleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSchemaOnly(context.Background(), subject, version).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSchemaOnly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaOnly`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSchemaOnly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 
**version** | **string** | Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;. \&quot;latest\&quot; returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaOnlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleted** | **bool** |  | 

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

> []string GetSchemaTypes(ctx).Execute()

Get the schema types supported by this registry.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSchemaTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSchemaTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSchemaTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaTypesRequest struct via the builder pattern


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

> []Schema GetSchemas(ctx).SubjectPrefix(subjectPrefix).Deleted(deleted).LatestOnly(latestOnly).Offset(offset).Limit(limit).Execute()

Get the schemas.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectPrefix := "subjectPrefix_example" // string |  (optional)
    deleted := true // bool |  (optional) (default to false)
    latestOnly := true // bool |  (optional) (default to false)
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSchemas(context.Background()).SubjectPrefix(subjectPrefix).Deleted(deleted).LatestOnly(latestOnly).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemas`: []Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectPrefix** | **string** |  | 
 **deleted** | **bool** |  | [default to false]
 **latestOnly** | **bool** |  | [default to false]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to -1]

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

> Config GetSubjectLevelConfig(ctx, subject).DefaultToGlobal(defaultToGlobal).Execute()

Get compatibility level for a subject.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | 
    defaultToGlobal := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSubjectLevelConfig(context.Background(), subject).DefaultToGlobal(defaultToGlobal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSubjectLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjectLevelConfig`: Config
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSubjectLevelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectLevelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultToGlobal** | **bool** |  | 

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

> []string GetSubjects(ctx, id).Subject(subject).Deleted(deleted).Execute()

Get all the subjects associated with the input ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | Globally unique identifier of the schema
    subject := "subject_example" // string |  (optional)
    deleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSubjects(context.Background(), id).Subject(subject).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSubjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSubjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Globally unique identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** |  | 
 **deleted** | **bool** |  | 

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

> TagDef GetTagDefByName(ctx, tagName).Execute()

Get the tag definition with the given name.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagName := "tagName_example" // string | The name of the tag definiton

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTagDefByName(context.Background(), tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTagDefByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagDefByName`: TagDef
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTagDefByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | The name of the tag definiton | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagDefByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []TagResponse GetTags(ctx, typeName, qualifiedName).Execute()

Gets the list of classifications for a given entity represented by a qualifed name.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTags(context.Background(), typeName, qualifiedName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> Config GetTopLevelConfig(ctx).Execute()

Get global compatibility level



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTopLevelConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTopLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopLevelConfig`: Config
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTopLevelConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopLevelConfigRequest struct via the builder pattern


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

> Mode GetTopLevelMode(ctx).Execute()

Get global mode.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTopLevelMode(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTopLevelMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopLevelMode`: Mode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTopLevelMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopLevelModeRequest struct via the builder pattern


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

> []SubjectVersion GetVersions(ctx, id).Subject(subject).Deleted(deleted).Execute()

Get all the subject-version pairs associated with the input ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | Globally unique identifier of the schema
    subject := "subject_example" // string |  (optional)
    deleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetVersions(context.Background(), id).Subject(subject).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersions`: []SubjectVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Globally unique identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** |  | 
 **deleted** | **bool** |  | 

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

> []string List(ctx).SubjectPrefix(subjectPrefix).Deleted(deleted).DeletedOnly(deletedOnly).Execute()

Get a list of registered subjects.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectPrefix := "subjectPrefix_example" // string |  (optional)
    deleted := true // bool |  (optional)
    deletedOnly := true // bool | Whether to return deleted subjects only (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.List(context.Background()).SubjectPrefix(subjectPrefix).Deleted(deleted).DeletedOnly(deletedOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectPrefix** | **string** |  | 
 **deleted** | **bool** |  | 
 **deletedOnly** | **bool** | Whether to return deleted subjects only | 

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

> []string ListContexts(ctx).ContextPrefix(contextPrefix).Execute()

Get a list of contexts.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contextPrefix := "contextPrefix_example" // string | prefix to filter contexts (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListContexts(context.Background()).ContextPrefix(contextPrefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContexts`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListContexts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextPrefix** | **string** | prefix to filter contexts | 

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

> []int32 ListVersions(ctx, subject).Deleted(deleted).DeletedOnly(deletedOnly).Execute()

Get a list of versions registered under the specified subject.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    deleted := true // bool |  (optional)
    deletedOnly := true // bool | Whether to return deleted schemas only (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListVersions(context.Background(), subject).Deleted(deleted).DeletedOnly(deletedOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersions`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleted** | **bool** |  | 
 **deletedOnly** | **bool** | Whether to return deleted schemas only | 

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

> LookUpSchemaUnderSubject(ctx, subject).Body(body).Deleted(deleted).Execute()

Check if a schema has already been registered under the specified subject. If so, this returns the schema string along with its globally unique identifier, its version under this subject and the subject name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Subject under which the schema will be registered
    body := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    deleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.LookUpSchemaUnderSubject(context.Background(), subject).Body(body).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LookUpSchemaUnderSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Subject under which the schema will be registered | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookUpSchemaUnderSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **deleted** | **bool** |  | 

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

> PartialUpdateByUniqueAttributes(ctx).AtlasEntityWithExtInfo(atlasEntityWithExtInfo).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    atlasEntityWithExtInfo := *openapiclient.NewAtlasEntityWithExtInfo() // AtlasEntityWithExtInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PartialUpdateByUniqueAttributes(context.Background()).AtlasEntityWithExtInfo(atlasEntityWithExtInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PartialUpdateByUniqueAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateByUniqueAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasEntityWithExtInfo** | [**AtlasEntityWithExtInfo**](AtlasEntityWithExtInfo.md) |  | 

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

> UpdateExporterResponse PauseExporter(ctx, name).Execute()

Pause an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PauseExporter(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PauseExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseExporter`: UpdateExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PauseExporter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> map[string]string Post(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Post(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Post`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Post`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostRequest struct via the builder pattern


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

> UpdateExporterResponse PutExporter(ctx, name).Body(body).Execute()

Alters an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the exporter
    body := *openapiclient.NewUpdateExporterRequest() // UpdateExporterRequest | Info

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PutExporter(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutExporter`: UpdateExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutExporter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateExporterRequest**](UpdateExporterRequest.md) | Info | 

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

> UpdateExporterResponse PutExporterConfig(ctx, name).Body(body).Execute()

Alters the config of an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the exporter
    body := map[string]string{"key": "Inner_example"} // map[string]string | Config

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PutExporterConfig(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutExporterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutExporterConfig`: UpdateExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutExporterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExporterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]string** | Config | 

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


## PutKek

> Kek PutKek(ctx, name).UpdateKekRequest(updateKekRequest).Execute()

Alters a kek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    updateKekRequest := *openapiclient.NewUpdateKekRequest() // UpdateKekRequest | The update request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PutKek(context.Background(), name).UpdateKekRequest(updateKekRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutKek`: Kek
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutKek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKekRequest** | [**UpdateKekRequest**](UpdateKekRequest.md) | The update request | 

### Return type

[**Kek**](Kek.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json, application/json, application/octet-stream
- **Accept**: application/vnd.schemaregistry.v1+json, application/vnd.schemaregistry+json; qs=0.9, application/json; qs=0.5

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> RegisterSchemaResponse Register(ctx, subject).Body(body).Normalize(normalize).Execute()

Register schema under a subject



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the subject
    body := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    normalize := true // bool | Whether to register the normalized schema (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Register(context.Background(), subject).Body(body).Normalize(normalize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: RegisterSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Register`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **normalize** | **bool** | Whether to register the normalized schema | 

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

> UpdateExporterResponse ResetExporter(ctx, name).Execute()

Reset an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResetExporter(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetExporter`: UpdateExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResetExporter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> UpdateExporterResponse ResumeExporter(ctx, name).Execute()

Resume an exporter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the exporter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResumeExporter(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResumeExporter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeExporter`: UpdateExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResumeExporter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exporter | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeExporterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> SearchResult SearchUsingAttribute(ctx).Type_(type_).Types(types).Attr(attr).Attrs(attrs).AttrName(attrName).AttrValuePrefix(attrValuePrefix).Tag(tag).TimeRangeType(timeRangeType).TimeRangeAttr(timeRangeAttr).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()

Retrieve data for the specified attribute search query.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := []string{"Inner_example"} // []string | Limit the result to only entities of specified types (optional)
    types := "types_example" // string |  (optional)
    attr := []string{"Inner_example"} // []string | One of more additional attributes to return in the response (optional)
    attrs := "attrs_example" // string |  (optional)
    attrName := []string{"Inner_example"} // []string | The attribute to search (optional)
    attrValuePrefix := []string{"Inner_example"} // []string | The prefix for the attribute value to search (optional)
    tag := []string{"Inner_example"} // []string | Limit the result to only entities tagged with the given tag (optional)
    timeRangeType := "timeRangeType_example" // string | The type of time range search, default is CUSTOM (optional)
    timeRangeAttr := "timeRangeAttr_example" // string | The attribute for a time range search (optional)
    timeRangeStart := int64(789) // int64 | The start for a custom time range search in ms since the epoch (optional)
    timeRangeEnd := int64(789) // int64 | The end for a custom time range search in ms since the epoch (optional)
    sortBy := "sortBy_example" // string | An attribute to sort by (optional)
    sortOrder := "sortOrder_example" // string | Sort order, either ASCENDING (default) or DESCENDING (optional)
    deleted := true // bool | Whether to include deleted entities (optional)
    limit := int32(56) // int32 | Limit the result set to only include the specified number of entries (optional)
    offset := int32(56) // int32 | Start offset of the result set (useful for pagination) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SearchUsingAttribute(context.Background()).Type_(type_).Types(types).Attr(attr).Attrs(attrs).AttrName(attrName).AttrValuePrefix(attrValuePrefix).Tag(tag).TimeRangeType(timeRangeType).TimeRangeAttr(timeRangeAttr).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchUsingAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsingAttribute`: SearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchUsingAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsingAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **[]string** | Limit the result to only entities of specified types | 
 **types** | **string** |  | 
 **attr** | **[]string** | One of more additional attributes to return in the response | 
 **attrs** | **string** |  | 
 **attrName** | **[]string** | The attribute to search | 
 **attrValuePrefix** | **[]string** | The prefix for the attribute value to search | 
 **tag** | **[]string** | Limit the result to only entities tagged with the given tag | 
 **timeRangeType** | **string** | The type of time range search, default is CUSTOM | 
 **timeRangeAttr** | **string** | The attribute for a time range search | 
 **timeRangeStart** | **int64** | The start for a custom time range search in ms since the epoch | 
 **timeRangeEnd** | **int64** | The end for a custom time range search in ms since the epoch | 
 **sortBy** | **string** | An attribute to sort by | 
 **sortOrder** | **string** | Sort order, either ASCENDING (default) or DESCENDING | 
 **deleted** | **bool** | Whether to include deleted entities | 
 **limit** | **int32** | Limit the result set to only include the specified number of entries | 
 **offset** | **int32** | Start offset of the result set (useful for pagination) | 

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

> SearchResult SearchUsingBasic(ctx).Query(query).Type_(type_).Types(types).Attr(attr).Attrs(attrs).Tag(tag).TimeRangeType(timeRangeType).TimeRangeAttr(timeRangeAttr).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()

Retrieve data for the specified fulltext query.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    query := "query_example" // string | The full-text query (optional)
    type_ := []string{"Inner_example"} // []string | Limit the result to only entities of specified types (optional)
    types := "types_example" // string |  (optional)
    attr := []string{"Inner_example"} // []string | One of more additional attributes to return in the response (optional)
    attrs := "attrs_example" // string |  (optional)
    tag := []string{"Inner_example"} // []string | Limit the result to only entities with the given tag (optional)
    timeRangeType := "timeRangeType_example" // string | The type of time range search, default is CUSTOM (optional)
    timeRangeAttr := "timeRangeAttr_example" // string | The attribute for a time range search (optional)
    timeRangeStart := int64(789) // int64 | The start for a custom time range search in ms since the epoch (optional)
    timeRangeEnd := int64(789) // int64 | The end for a custom time range search in ms since the epoch (optional)
    sortBy := "sortBy_example" // string | An attribute to sort by (optional)
    sortOrder := "sortOrder_example" // string | Sort order, either ASCENDING (default) or DESCENDING (optional)
    deleted := true // bool | Whether to include deleted entities (optional)
    limit := int32(56) // int32 | Limit the result set to only include the specified number of entries (optional)
    offset := int32(56) // int32 | Start offset of the result set (useful for pagination) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SearchUsingBasic(context.Background()).Query(query).Type_(type_).Types(types).Attr(attr).Attrs(attrs).Tag(tag).TimeRangeType(timeRangeType).TimeRangeAttr(timeRangeAttr).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).SortBy(sortBy).SortOrder(sortOrder).Deleted(deleted).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchUsingBasic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsingBasic`: SearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchUsingBasic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsingBasicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The full-text query | 
 **type_** | **[]string** | Limit the result to only entities of specified types | 
 **types** | **string** |  | 
 **attr** | **[]string** | One of more additional attributes to return in the response | 
 **attrs** | **string** |  | 
 **tag** | **[]string** | Limit the result to only entities with the given tag | 
 **timeRangeType** | **string** | The type of time range search, default is CUSTOM | 
 **timeRangeAttr** | **string** | The attribute for a time range search | 
 **timeRangeStart** | **int64** | The start for a custom time range search in ms since the epoch | 
 **timeRangeEnd** | **int64** | The end for a custom time range search in ms since the epoch | 
 **sortBy** | **string** | An attribute to sort by | 
 **sortOrder** | **string** | Sort order, either ASCENDING (default) or DESCENDING | 
 **deleted** | **bool** | Whether to include deleted entities | 
 **limit** | **int32** | Limit the result set to only include the specified number of entries | 
 **offset** | **int32** | Start offset of the result set (useful for pagination) | 

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

> CompatibilityCheckResponse TestCompatibilityBySubjectName(ctx, subject, version).Body(body).ContentType(contentType).Accept(accept).Verbose(verbose).Execute()

Test input schema against a particular version of a subject's schema for compatibility.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Subject of the schema version against which compatibility is to be tested
    version := "version_example" // string | Version of the subject's schema against which compatibility is to be tested. Valid values for versionId are between [1,2^31-1] or the string \"latest\".\"latest\" checks compatibility of the input schema with the last registered schema under the specified subject
    body := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    contentType := "contentType_example" // string |  (optional)
    accept := "accept_example" // string |  (optional)
    verbose := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestCompatibilityBySubjectName(context.Background(), subject, version).Body(body).ContentType(contentType).Accept(accept).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestCompatibilityBySubjectName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCompatibilityBySubjectName`: CompatibilityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestCompatibilityBySubjectName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Subject of the schema version against which compatibility is to be tested | 
**version** | **string** | Version of the subject&#39;s schema against which compatibility is to be tested. Valid values for versionId are between [1,2^31-1] or the string \&quot;latest\&quot;.\&quot;latest\&quot; checks compatibility of the input schema with the last registered schema under the specified subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestCompatibilityBySubjectNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 
 **verbose** | **bool** |  | 

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

> CompatibilityCheckResponse TestCompatibilityForSubject(ctx, subject).Body(body).ContentType(contentType).Accept(accept).Verbose(verbose).Execute()

Test input schema against a subject's schemas for compatibility, based on the compatibility level of the subject configured. In other word, it will perform the same compatibility check as register for that subject



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Subject of the schema version against which compatibility is to be tested
    body := *openapiclient.NewRegisterSchemaRequest() // RegisterSchemaRequest | Schema
    contentType := "contentType_example" // string |  (optional)
    accept := "accept_example" // string |  (optional)
    verbose := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestCompatibilityForSubject(context.Background(), subject).Body(body).ContentType(contentType).Accept(accept).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestCompatibilityForSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCompatibilityForSubject`: CompatibilityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestCompatibilityForSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Subject of the schema version against which compatibility is to be tested | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestCompatibilityForSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RegisterSchemaRequest**](RegisterSchemaRequest.md) | Schema | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 
 **verbose** | **bool** |  | 

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


## UndeleteDekVersion

> UndeleteDekVersion(ctx, name, subject, version).Algorithm(algorithm).Execute()

Undelete a dek version.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    subject := "subject_example" // string | Subject of the dek
    version := "version_example" // string | Version of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UndeleteDekVersion(context.Background(), name, subject, version).Algorithm(algorithm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UndeleteDekVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 
**version** | **string** | Version of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteDekVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **algorithm** | **string** | Algorithm of the dek | 

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


## UndeleteDekVersions

> UndeleteDekVersions(ctx, name, subject).Algorithm(algorithm).Execute()

Undelete all versions of a dek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek
    subject := "subject_example" // string | Subject of the dek
    algorithm := "algorithm_example" // string | Algorithm of the dek (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UndeleteDekVersions(context.Background(), name, subject).Algorithm(algorithm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UndeleteDekVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 
**subject** | **string** | Subject of the dek | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteDekVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **algorithm** | **string** | Algorithm of the dek | 

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


## UndeleteKek

> UndeleteKek(ctx, name).Execute()

Undelete a kek.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the kek

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UndeleteKek(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UndeleteKek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the kek | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteKekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateBusinessMetadata

> []BusinessMetadataResponse UpdateBusinessMetadata(ctx).BusinessMetadata(businessMetadata).Execute()

Bulk API to update multiple business metadata.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    businessMetadata := []openapiclient.BusinessMetadata{*openapiclient.NewBusinessMetadata()} // []BusinessMetadata | The business metadata (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateBusinessMetadata(context.Background()).BusinessMetadata(businessMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBusinessMetadata`: []BusinessMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBusinessMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadata** | [**[]BusinessMetadata**](BusinessMetadata.md) | The business metadata | 

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

> []BusinessMetadataDefResponse UpdateBusinessMetadataDefs(ctx).AtlasBusinessMetadataDef(atlasBusinessMetadataDef).Execute()

Bulk update API for business metadata definitions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    atlasBusinessMetadataDef := []openapiclient.AtlasBusinessMetadataDef{*openapiclient.NewAtlasBusinessMetadataDef()} // []AtlasBusinessMetadataDef | The business metadata definitions to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateBusinessMetadataDefs(context.Background()).AtlasBusinessMetadataDef(atlasBusinessMetadataDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBusinessMetadataDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBusinessMetadataDefs`: []BusinessMetadataDefResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBusinessMetadataDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessMetadataDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlasBusinessMetadataDef** | [**[]AtlasBusinessMetadataDef**](AtlasBusinessMetadataDef.md) | The business metadata definitions to update | 

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

> ModeUpdateRequest UpdateMode(ctx, subject).Body(body).Execute()

Update mode for the specified subject.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    body := *openapiclient.NewModeUpdateRequest() // ModeUpdateRequest | Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMode(context.Background(), subject).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMode`: ModeUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ModeUpdateRequest**](ModeUpdateRequest.md) | Update Request | 

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

> ConfigUpdateRequest UpdateSubjectLevelConfig(ctx, subject).Body(body).Execute()

Update compatibility level for the specified subject.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Name of the Subject
    body := *openapiclient.NewConfigUpdateRequest() // ConfigUpdateRequest | Config Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSubjectLevelConfig(context.Background(), subject).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSubjectLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubjectLevelConfig`: ConfigUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSubjectLevelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Name of the Subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubjectLevelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfigUpdateRequest**](ConfigUpdateRequest.md) | Config Update Request | 

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

> []TagDefResponse UpdateTagDefs(ctx).TagDef(tagDef).Execute()

Bulk update API for tag definitions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagDef := []openapiclient.TagDef{*openapiclient.NewTagDef()} // []TagDef | The tag definitions to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTagDefs(context.Background()).TagDef(tagDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTagDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagDefs`: []TagDefResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTagDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDef** | [**[]TagDef**](TagDef.md) | The tag definitions to update | 

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

> []TagResponse UpdateTags(ctx).Tag(tag).Execute()

Bulk API to update multiple tags.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | The tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTags(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**[]Tag**](Tag.md) | The tags | 

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

> ConfigUpdateRequest UpdateTopLevelConfig(ctx).ConfigUpdateRequest(configUpdateRequest).Execute()

Update global compatibility level



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configUpdateRequest := *openapiclient.NewConfigUpdateRequest() // ConfigUpdateRequest | Config Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTopLevelConfig(context.Background()).ConfigUpdateRequest(configUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTopLevelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopLevelConfig`: ConfigUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTopLevelConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopLevelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configUpdateRequest** | [**ConfigUpdateRequest**](ConfigUpdateRequest.md) | Config Update Request | 

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

> ModeUpdateRequest UpdateTopLevelMode(ctx).Body(body).Execute()

Update global mode.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewModeUpdateRequest() // ModeUpdateRequest | Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTopLevelMode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTopLevelMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopLevelMode`: ModeUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTopLevelMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopLevelModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ModeUpdateRequest**](ModeUpdateRequest.md) | Update Request | 

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

