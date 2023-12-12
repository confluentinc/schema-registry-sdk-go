# CreateExporterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Subjects** | Pointer to **[]string** |  | [optional] 
**ContextType** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**SubjectRenameFormat** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateExporterRequest

`func NewCreateExporterRequest() *CreateExporterRequest`

NewCreateExporterRequest instantiates a new CreateExporterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExporterRequestWithDefaults

`func NewCreateExporterRequestWithDefaults() *CreateExporterRequest`

NewCreateExporterRequestWithDefaults instantiates a new CreateExporterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateExporterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExporterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExporterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateExporterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubjects

`func (o *CreateExporterRequest) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CreateExporterRequest) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CreateExporterRequest) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CreateExporterRequest) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetContextType

`func (o *CreateExporterRequest) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *CreateExporterRequest) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *CreateExporterRequest) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *CreateExporterRequest) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetContext

`func (o *CreateExporterRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CreateExporterRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CreateExporterRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CreateExporterRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSubjectRenameFormat

`func (o *CreateExporterRequest) GetSubjectRenameFormat() string`

GetSubjectRenameFormat returns the SubjectRenameFormat field if non-nil, zero value otherwise.

### GetSubjectRenameFormatOk

`func (o *CreateExporterRequest) GetSubjectRenameFormatOk() (*string, bool)`

GetSubjectRenameFormatOk returns a tuple with the SubjectRenameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRenameFormat

`func (o *CreateExporterRequest) SetSubjectRenameFormat(v string)`

SetSubjectRenameFormat sets SubjectRenameFormat field to given value.

### HasSubjectRenameFormat

`func (o *CreateExporterRequest) HasSubjectRenameFormat() bool`

HasSubjectRenameFormat returns a boolean if a field has been set.

### GetConfig

`func (o *CreateExporterRequest) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateExporterRequest) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateExporterRequest) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateExporterRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


