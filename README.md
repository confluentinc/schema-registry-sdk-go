# Go API client for schemaregistry

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./schemaregistry"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AsyncapiParsePut**](docs/DefaultApi.md#asyncapiparseput) | **Put** /asyncapi/parse | 
*DefaultApi* | [**AsyncapiPut**](docs/DefaultApi.md#asyncapiput) | **Put** /asyncapi | 
*DefaultApi* | [**CreateBusinessMetadata**](docs/DefaultApi.md#createbusinessmetadata) | **Post** /catalog/v1/entity/businessmetadata | Bulk API to create multiple business metadata.
*DefaultApi* | [**CreateBusinessMetadataDefs**](docs/DefaultApi.md#createbusinessmetadatadefs) | **Post** /catalog/v1/types/businessmetadatadefs | Bulk create API for business metadata definitions.
*DefaultApi* | [**CreateDek**](docs/DefaultApi.md#createdek) | **Post** /dek-registry/v1/keks/{name}/deks | Create a dek.
*DefaultApi* | [**CreateExporter**](docs/DefaultApi.md#createexporter) | **Post** /exporters | Create an exporter.
*DefaultApi* | [**CreateKek**](docs/DefaultApi.md#createkek) | **Post** /dek-registry/v1/keks | Create a kek.
*DefaultApi* | [**CreateOrUpdate**](docs/DefaultApi.md#createorupdate) | **Post** /catalog/v1/entity | 
*DefaultApi* | [**CreateTagDefs**](docs/DefaultApi.md#createtagdefs) | **Post** /catalog/v1/types/tagdefs | Bulk create API for tag definitions.
*DefaultApi* | [**CreateTags**](docs/DefaultApi.md#createtags) | **Post** /catalog/v1/entity/tags | Bulk API to create multiple tags.
*DefaultApi* | [**DeleteBusinessMetadata**](docs/DefaultApi.md#deletebusinessmetadata) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata/{bmName} | Delete a business metadata on an entity.
*DefaultApi* | [**DeleteBusinessMetadataDef**](docs/DefaultApi.md#deletebusinessmetadatadef) | **Delete** /catalog/v1/types/businessmetadatadefs/{bmName} | Delete API for business metadata definition identified by its name.
*DefaultApi* | [**DeleteByUniqueAttributes**](docs/DefaultApi.md#deletebyuniqueattributes) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | 
*DefaultApi* | [**DeleteDekVersion**](docs/DefaultApi.md#deletedekversion) | **Delete** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version} | Delete a dek version.
*DefaultApi* | [**DeleteDekVersions**](docs/DefaultApi.md#deletedekversions) | **Delete** /dek-registry/v1/keks/{name}/deks/{subject} | Delete all versions of a dek.
*DefaultApi* | [**DeleteExporter**](docs/DefaultApi.md#deleteexporter) | **Delete** /exporters/{name} | Delete an exporter.
*DefaultApi* | [**DeleteKek**](docs/DefaultApi.md#deletekek) | **Delete** /dek-registry/v1/keks/{name} | Delete a kek.
*DefaultApi* | [**DeleteSchemaVersion**](docs/DefaultApi.md#deleteschemaversion) | **Delete** /subjects/{subject}/versions/{version} | Deletes a specific version of the schema registered under this subject. This only deletes the version and the schema ID remains intact making it still possible to decode data using the schema ID. This API is recommended to be used only in development environments or under extreme circumstances where-in, its required to delete a previously registered schema for compatibility purposes or re-register previously registered schema.
*DefaultApi* | [**DeleteSubject**](docs/DefaultApi.md#deletesubject) | **Delete** /subjects/{subject} | Deletes the specified subject and its associated compatibility level if registered. It is recommended to use this API only when a topic needs to be recycled or in development environment.
*DefaultApi* | [**DeleteSubjectConfig**](docs/DefaultApi.md#deletesubjectconfig) | **Delete** /config/{subject} | Deletes the specified subject-level compatibility level config and revert to the global default.
*DefaultApi* | [**DeleteSubjectMode**](docs/DefaultApi.md#deletesubjectmode) | **Delete** /mode/{subject} | Deletes the specified subject-level mode and revert to the global default.
*DefaultApi* | [**DeleteTag**](docs/DefaultApi.md#deletetag) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags/{tagName} | Delete a tag on an entity.
*DefaultApi* | [**DeleteTagDef**](docs/DefaultApi.md#deletetagdef) | **Delete** /catalog/v1/types/tagdefs/{tagName} | Delete API for tag definition identified by its name.
*DefaultApi* | [**DeleteTopLevelConfig**](docs/DefaultApi.md#deletetoplevelconfig) | **Delete** /config | Delete global compatibility level
*DefaultApi* | [**Get**](docs/DefaultApi.md#get) | **Get** / | Schema Registry Root Resource
*DefaultApi* | [**GetAllBusinessMetadataDefs**](docs/DefaultApi.md#getallbusinessmetadatadefs) | **Get** /catalog/v1/types/businessmetadatadefs | Bulk retrieval API for retrieving business metadata definitions.
*DefaultApi* | [**GetAllTagDefs**](docs/DefaultApi.md#getalltagdefs) | **Get** /catalog/v1/types/tagdefs | Bulk retrieval API for retrieving tag definitions.
*DefaultApi* | [**GetBusinessMetadata**](docs/DefaultApi.md#getbusinessmetadata) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata | Gets the list of business metadata for a given entity represented by a qualified name.
*DefaultApi* | [**GetBusinessMetadataDefByName**](docs/DefaultApi.md#getbusinessmetadatadefbyname) | **Get** /catalog/v1/types/businessmetadatadefs/{bmName} | Get the business metadata definition with the given name.
*DefaultApi* | [**GetByUniqueAttributes**](docs/DefaultApi.md#getbyuniqueattributes) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | Fetch complete definition of an entity given its type and unique attribute.
*DefaultApi* | [**GetClusterId**](docs/DefaultApi.md#getclusterid) | **Get** /v1/metadata/id | Get the server metadata
*DefaultApi* | [**GetDek**](docs/DefaultApi.md#getdek) | **Get** /dek-registry/v1/keks/{name}/deks/{subject} | Get a dek by subject.
*DefaultApi* | [**GetDekByVersion**](docs/DefaultApi.md#getdekbyversion) | **Get** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version} | Get a dek by subject and version.
*DefaultApi* | [**GetDekSubjects**](docs/DefaultApi.md#getdeksubjects) | **Get** /dek-registry/v1/keks/{name}/deks | Get a list of dek subjects.
*DefaultApi* | [**GetDekVersions**](docs/DefaultApi.md#getdekversions) | **Get** /dek-registry/v1/keks/{name}/deks/{subject}/versions | List versions of dek.
*DefaultApi* | [**GetExporterConfig**](docs/DefaultApi.md#getexporterconfig) | **Get** /exporters/{name}/config | Get the config for an exporter.
*DefaultApi* | [**GetExporterInfo**](docs/DefaultApi.md#getexporterinfo) | **Get** /exporters/{name} | Get the info for an exporter.
*DefaultApi* | [**GetExporterStatus**](docs/DefaultApi.md#getexporterstatus) | **Get** /exporters/{name}/status | Get the status for an exporter.
*DefaultApi* | [**GetExporters**](docs/DefaultApi.md#getexporters) | **Get** /exporters | Get a list of exporter names.
*DefaultApi* | [**GetKek**](docs/DefaultApi.md#getkek) | **Get** /dek-registry/v1/keks/{name} | Get a kek by name.
*DefaultApi* | [**GetKekNames**](docs/DefaultApi.md#getkeknames) | **Get** /dek-registry/v1/keks | Get a list of kek names.
*DefaultApi* | [**GetMode**](docs/DefaultApi.md#getmode) | **Get** /mode/{subject} | Get mode for a subject.
*DefaultApi* | [**GetReferencedBy**](docs/DefaultApi.md#getreferencedby) | **Get** /subjects/{subject}/versions/{version}/referencedby | Get the schemas that reference the specified schema.
*DefaultApi* | [**GetSchema**](docs/DefaultApi.md#getschema) | **Get** /schemas/ids/{id} | Get the schema string identified by the input ID.
*DefaultApi* | [**GetSchemaByVersion**](docs/DefaultApi.md#getschemabyversion) | **Get** /subjects/{subject}/versions/{version} | Get a specific version of the schema registered under this subject.
*DefaultApi* | [**GetSchemaOnly**](docs/DefaultApi.md#getschemaonly) | **Get** /subjects/{subject}/versions/{version}/schema | Get the schema for the specified version of this subject. The unescaped schema only is returned.
*DefaultApi* | [**GetSchemaTypes**](docs/DefaultApi.md#getschematypes) | **Get** /schemas/types | Get the schema types supported by this registry.
*DefaultApi* | [**GetSchemas**](docs/DefaultApi.md#getschemas) | **Get** /schemas | Get the schemas.
*DefaultApi* | [**GetSubjectLevelConfig**](docs/DefaultApi.md#getsubjectlevelconfig) | **Get** /config/{subject} | Get compatibility level for a subject.
*DefaultApi* | [**GetSubjects**](docs/DefaultApi.md#getsubjects) | **Get** /schemas/ids/{id}/subjects | Get all the subjects associated with the input ID.
*DefaultApi* | [**GetTagDefByName**](docs/DefaultApi.md#gettagdefbyname) | **Get** /catalog/v1/types/tagdefs/{tagName} | Get the tag definition with the given name.
*DefaultApi* | [**GetTags**](docs/DefaultApi.md#gettags) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags | Gets the list of classifications for a given entity represented by a qualifed name.
*DefaultApi* | [**GetTopLevelConfig**](docs/DefaultApi.md#gettoplevelconfig) | **Get** /config | Get global compatibility level
*DefaultApi* | [**GetTopLevelMode**](docs/DefaultApi.md#gettoplevelmode) | **Get** /mode | Get global mode.
*DefaultApi* | [**GetVersions**](docs/DefaultApi.md#getversions) | **Get** /schemas/ids/{id}/versions | Get all the subject-version pairs associated with the input ID.
*DefaultApi* | [**List**](docs/DefaultApi.md#list) | **Get** /subjects | Get a list of registered subjects.
*DefaultApi* | [**ListContexts**](docs/DefaultApi.md#listcontexts) | **Get** /contexts | Get a list of contexts.
*DefaultApi* | [**ListVersions**](docs/DefaultApi.md#listversions) | **Get** /subjects/{subject}/versions | Get a list of versions registered under the specified subject.
*DefaultApi* | [**LookUpSchemaUnderSubject**](docs/DefaultApi.md#lookupschemaundersubject) | **Post** /subjects/{subject} | Check if a schema has already been registered under the specified subject. If so, this returns the schema string along with its globally unique identifier, its version under this subject and the subject name.
*DefaultApi* | [**PartialUpdateByUniqueAttributes**](docs/DefaultApi.md#partialupdatebyuniqueattributes) | **Put** /catalog/v1/entity | 
*DefaultApi* | [**PauseExporter**](docs/DefaultApi.md#pauseexporter) | **Put** /exporters/{name}/pause | Pause an exporter.
*DefaultApi* | [**Post**](docs/DefaultApi.md#post) | **Post** / | 
*DefaultApi* | [**PutExporter**](docs/DefaultApi.md#putexporter) | **Put** /exporters/{name} | Alters an exporter.
*DefaultApi* | [**PutExporterConfig**](docs/DefaultApi.md#putexporterconfig) | **Put** /exporters/{name}/config | Alters the config of an exporter.
*DefaultApi* | [**PutKek**](docs/DefaultApi.md#putkek) | **Put** /dek-registry/v1/keks/{name} | Alters a kek.
*DefaultApi* | [**Register**](docs/DefaultApi.md#register) | **Post** /subjects/{subject}/versions | Register schema under a subject
*DefaultApi* | [**ResetExporter**](docs/DefaultApi.md#resetexporter) | **Put** /exporters/{name}/reset | Reset an exporter.
*DefaultApi* | [**ResumeExporter**](docs/DefaultApi.md#resumeexporter) | **Put** /exporters/{name}/resume | Resume an exporter.
*DefaultApi* | [**SearchUsingAttribute**](docs/DefaultApi.md#searchusingattribute) | **Get** /catalog/v1/search/attribute | Retrieve data for the specified attribute search query.
*DefaultApi* | [**SearchUsingBasic**](docs/DefaultApi.md#searchusingbasic) | **Get** /catalog/v1/search/basic | Retrieve data for the specified fulltext query.
*DefaultApi* | [**TestCompatibilityBySubjectName**](docs/DefaultApi.md#testcompatibilitybysubjectname) | **Post** /compatibility/subjects/{subject}/versions/{version} | Test input schema against a particular version of a subject&#39;s schema for compatibility.
*DefaultApi* | [**TestCompatibilityForSubject**](docs/DefaultApi.md#testcompatibilityforsubject) | **Post** /compatibility/subjects/{subject}/versions | Test input schema against a subject&#39;s schemas for compatibility, based on the compatibility level of the subject configured. In other word, it will perform the same compatibility check as register for that subject
*DefaultApi* | [**UndeleteDekVersion**](docs/DefaultApi.md#undeletedekversion) | **Post** /dek-registry/v1/keks/{name}/deks/{subject}/versions/{version}/undelete | Undelete a dek version.
*DefaultApi* | [**UndeleteDekVersions**](docs/DefaultApi.md#undeletedekversions) | **Post** /dek-registry/v1/keks/{name}/deks/{subject}/undelete | Undelete all versions of a dek.
*DefaultApi* | [**UndeleteKek**](docs/DefaultApi.md#undeletekek) | **Post** /dek-registry/v1/keks/{name}/undelete | Undelete a kek.
*DefaultApi* | [**UpdateBusinessMetadata**](docs/DefaultApi.md#updatebusinessmetadata) | **Put** /catalog/v1/entity/businessmetadata | Bulk API to update multiple business metadata.
*DefaultApi* | [**UpdateBusinessMetadataDefs**](docs/DefaultApi.md#updatebusinessmetadatadefs) | **Put** /catalog/v1/types/businessmetadatadefs | Bulk update API for business metadata definitions.
*DefaultApi* | [**UpdateMode**](docs/DefaultApi.md#updatemode) | **Put** /mode/{subject} | Update mode for the specified subject.
*DefaultApi* | [**UpdateSubjectLevelConfig**](docs/DefaultApi.md#updatesubjectlevelconfig) | **Put** /config/{subject} | Update compatibility level for the specified subject.
*DefaultApi* | [**UpdateTagDefs**](docs/DefaultApi.md#updatetagdefs) | **Put** /catalog/v1/types/tagdefs | Bulk update API for tag definitions.
*DefaultApi* | [**UpdateTags**](docs/DefaultApi.md#updatetags) | **Put** /catalog/v1/entity/tags | Bulk API to update multiple tags.
*DefaultApi* | [**UpdateTopLevelConfig**](docs/DefaultApi.md#updatetoplevelconfig) | **Put** /config | Update global compatibility level
*DefaultApi* | [**UpdateTopLevelMode**](docs/DefaultApi.md#updatetoplevelmode) | **Put** /mode | Update global mode.


## Documentation For Models

 - [AtlasAttributeDef](docs/AtlasAttributeDef.md)
 - [AtlasBusinessMetadataDef](docs/AtlasBusinessMetadataDef.md)
 - [AtlasClassification](docs/AtlasClassification.md)
 - [AtlasConstraintDef](docs/AtlasConstraintDef.md)
 - [AtlasEntity](docs/AtlasEntity.md)
 - [AtlasEntityHeader](docs/AtlasEntityHeader.md)
 - [AtlasEntityWithExtInfo](docs/AtlasEntityWithExtInfo.md)
 - [AtlasTermAssignmentHeader](docs/AtlasTermAssignmentHeader.md)
 - [BusinessMetadata](docs/BusinessMetadata.md)
 - [BusinessMetadataDefResponse](docs/BusinessMetadataDefResponse.md)
 - [BusinessMetadataResponse](docs/BusinessMetadataResponse.md)
 - [CompatibilityCheckResponse](docs/CompatibilityCheckResponse.md)
 - [Config](docs/Config.md)
 - [ConfigUpdateRequest](docs/ConfigUpdateRequest.md)
 - [CreateDekRequest](docs/CreateDekRequest.md)
 - [CreateExporterRequest](docs/CreateExporterRequest.md)
 - [CreateExporterResponse](docs/CreateExporterResponse.md)
 - [CreateKekRequest](docs/CreateKekRequest.md)
 - [Dek](docs/Dek.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [ExporterInfo](docs/ExporterInfo.md)
 - [ExporterStatus](docs/ExporterStatus.md)
 - [Kek](docs/Kek.md)
 - [Metadata](docs/Metadata.md)
 - [Mode](docs/Mode.md)
 - [ModeUpdateRequest](docs/ModeUpdateRequest.md)
 - [RegisterSchemaRequest](docs/RegisterSchemaRequest.md)
 - [RegisterSchemaResponse](docs/RegisterSchemaResponse.md)
 - [Rule](docs/Rule.md)
 - [RuleSet](docs/RuleSet.md)
 - [Schema](docs/Schema.md)
 - [SchemaEntity](docs/SchemaEntity.md)
 - [SchemaReference](docs/SchemaReference.md)
 - [SchemaString](docs/SchemaString.md)
 - [SchemaTags](docs/SchemaTags.md)
 - [SearchParams](docs/SearchParams.md)
 - [SearchResult](docs/SearchResult.md)
 - [ServerClusterId](docs/ServerClusterId.md)
 - [SubjectVersion](docs/SubjectVersion.md)
 - [Tag](docs/Tag.md)
 - [TagDef](docs/TagDef.md)
 - [TagDefResponse](docs/TagDefResponse.md)
 - [TagResponse](docs/TagResponse.md)
 - [TimeBoundary](docs/TimeBoundary.md)
 - [UpdateExporterRequest](docs/UpdateExporterRequest.md)
 - [UpdateExporterResponse](docs/UpdateExporterResponse.md)
 - [UpdateKekRequest](docs/UpdateKekRequest.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



