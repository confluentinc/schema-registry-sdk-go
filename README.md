# Go API client for schemaregistry

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Mocking	
We use a [custom fork](https://github.com/DABH/openapi-generator) of the openapi generator to generate interfaces that can be used to generate mocks. Once the interfaces are updated, use the `make mock` command to generate new mocks.

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./schemaregistry"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**DeleteSchemaVersion**](docs/DefaultApi.md#deleteschemaversion) | **Delete** /subjects/{subject}/versions/{version} | Deletes a specific version of the schema registered under this subject. This only deletes the version and the schema ID remains intact making it still possible to decode data using the schema ID. This API is recommended to be used only in development environments or under extreme circumstances where-in, its required to delete a previously registered schema for compatibility purposes or re-register previously registered schema.
*DefaultApi* | [**DeleteSubject**](docs/DefaultApi.md#deletesubject) | **Delete** /subjects/{subject} | Deletes the specified subject and its associated compatibility level if registered. It is recommended to use this API only when a topic needs to be recycled or in development environment.
*DefaultApi* | [**Get**](docs/DefaultApi.md#get) | **Get** / | Schema Registry Root Resource
*DefaultApi* | [**GetClusterId**](docs/DefaultApi.md#getclusterid) | **Get** /v1/metadata/id | Get the server metadata
*DefaultApi* | [**GetMode**](docs/DefaultApi.md#getmode) | **Get** /mode/{subject} | 
*DefaultApi* | [**GetReferencedBy**](docs/DefaultApi.md#getreferencedby) | **Get** /subjects/{subject}/versions/{version}/referencedby | Get the schemas that reference the specified schema.
*DefaultApi* | [**GetSchema**](docs/DefaultApi.md#getschema) | **Get** /schemas/ids/{id} | Get the schema string identified by the input ID.
*DefaultApi* | [**GetSchemaByVersion**](docs/DefaultApi.md#getschemabyversion) | **Get** /subjects/{subject}/versions/{version} | Get a specific version of the schema registered under this subject.
*DefaultApi* | [**GetSchemaOnly**](docs/DefaultApi.md#getschemaonly) | **Get** /subjects/{subject}/versions/{version}/schema | Get the schema for the specified version of this subject. The unescaped schema only is returned.
*DefaultApi* | [**GetSchemaTypes**](docs/DefaultApi.md#getschematypes) | **Get** /schemas/types | Get the schema types supported by this registry.
*DefaultApi* | [**GetSchemas**](docs/DefaultApi.md#getschemas) | **Get** /schemas | Get the schemas.
*DefaultApi* | [**GetSubjectLevelConfig**](docs/DefaultApi.md#getsubjectlevelconfig) | **Get** /config/{subject} | Get compatibility level for a subject.
*DefaultApi* | [**GetSubjects**](docs/DefaultApi.md#getsubjects) | **Get** /schemas/ids/{id}/subjects | Get all the subjects associated with the input ID.
*DefaultApi* | [**GetTopLevelConfig**](docs/DefaultApi.md#gettoplevelconfig) | **Get** /config | Get global compatibility level.
*DefaultApi* | [**GetTopLevelMode**](docs/DefaultApi.md#gettoplevelmode) | **Get** /mode | 
*DefaultApi* | [**GetVersions**](docs/DefaultApi.md#getversions) | **Get** /schemas/ids/{id}/versions | Get all the subject-version pairs associated with the input ID.
*DefaultApi* | [**List**](docs/DefaultApi.md#list) | **Get** /subjects | Get a list of registered subjects.
*DefaultApi* | [**ListVersions**](docs/DefaultApi.md#listversions) | **Get** /subjects/{subject}/versions | Get a list of versions registered under the specified subject.
*DefaultApi* | [**LookUpSchemaUnderSubject**](docs/DefaultApi.md#lookupschemaundersubject) | **Post** /subjects/{subject} | Check if a schema has already been registered under the specified subject. If so, this returns the schema string along with its globally unique identifier, its version under this subject and the subject name.
*DefaultApi* | [**Post**](docs/DefaultApi.md#post) | **Post** / | 
*DefaultApi* | [**Register**](docs/DefaultApi.md#register) | **Post** /subjects/{subject}/versions | Register a new schema under the specified subject. If successfully registered, this returns the unique identifier of this schema in the registry. The returned identifier should be used to retrieve this schema from the schemas resource and is different from the schema&#39;s version which is associated with the subject. If the same schema is registered under a different subject, the same identifier will be returned. However, the version of the schema may be different under different subjects. A schema should be compatible with the previously registered schema or schemas (if there are any) as per the configured compatibility level. The configured compatibility level can be obtained by issuing a GET http:get:: /config/(string: subject). If that returns null, then GET http:get:: /config When there are multiple instances of Schema Registry running in the same cluster, the schema registration request will be forwarded to one of the instances designated as the primary. If the primary is not available, the client will get an error code indicating that the forwarding has failed.
*DefaultApi* | [**TestCompatibilityBySubjectName**](docs/DefaultApi.md#testcompatibilitybysubjectname) | **Post** /compatibility/subjects/{subject}/versions/{version} | Test input schema against a particular version of a subject&#39;s schema for compatibility.
*DefaultApi* | [**UpdateMode**](docs/DefaultApi.md#updatemode) | **Put** /mode/{subject} | 
*DefaultApi* | [**UpdateSubjectLevelConfig**](docs/DefaultApi.md#updatesubjectlevelconfig) | **Put** /config/{subject} | Update compatibility level for the specified subject.
*DefaultApi* | [**UpdateTopLevelConfig**](docs/DefaultApi.md#updatetoplevelconfig) | **Put** /config | Update global compatibility level.
*DefaultApi* | [**UpdateTopLevelMode**](docs/DefaultApi.md#updatetoplevelmode) | **Put** /mode | 


## Documentation For Models

 - [CompatibilityCheckResponse](docs/CompatibilityCheckResponse.md)
 - [Config](docs/Config.md)
 - [ConfigUpdateRequest](docs/ConfigUpdateRequest.md)
 - [ModeGetResponse](docs/ModeGetResponse.md)
 - [ModeUpdateRequest](docs/ModeUpdateRequest.md)
 - [RegisterSchemaRequest](docs/RegisterSchemaRequest.md)
 - [RegisterSchemaResponse](docs/RegisterSchemaResponse.md)
 - [Schema](docs/Schema.md)
 - [SchemaReference](docs/SchemaReference.md)
 - [SchemaString](docs/SchemaString.md)
 - [ServerClusterId](docs/ServerClusterId.md)
 - [SubjectVersion](docs/SubjectVersion.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Author



