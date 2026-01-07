# UpdateExporterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subjects** | Pointer to **[]string** |  | [optional] 
**ContextType** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**KekRenameFormat** | Pointer to **string** |  | [optional] 
**SubjectRenameFormat** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateExporterRequest

`func NewUpdateExporterRequest() *UpdateExporterRequest`

NewUpdateExporterRequest instantiates a new UpdateExporterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExporterRequestWithDefaults

`func NewUpdateExporterRequestWithDefaults() *UpdateExporterRequest`

NewUpdateExporterRequestWithDefaults instantiates a new UpdateExporterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjects

`func (o *UpdateExporterRequest) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *UpdateExporterRequest) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *UpdateExporterRequest) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *UpdateExporterRequest) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetContextType

`func (o *UpdateExporterRequest) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *UpdateExporterRequest) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *UpdateExporterRequest) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *UpdateExporterRequest) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetContext

`func (o *UpdateExporterRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UpdateExporterRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UpdateExporterRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *UpdateExporterRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetKekRenameFormat

`func (o *UpdateExporterRequest) GetKekRenameFormat() string`

GetKekRenameFormat returns the KekRenameFormat field if non-nil, zero value otherwise.

### GetKekRenameFormatOk

`func (o *UpdateExporterRequest) GetKekRenameFormatOk() (*string, bool)`

GetKekRenameFormatOk returns a tuple with the KekRenameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKekRenameFormat

`func (o *UpdateExporterRequest) SetKekRenameFormat(v string)`

SetKekRenameFormat sets KekRenameFormat field to given value.

### HasKekRenameFormat

`func (o *UpdateExporterRequest) HasKekRenameFormat() bool`

HasKekRenameFormat returns a boolean if a field has been set.

### GetSubjectRenameFormat

`func (o *UpdateExporterRequest) GetSubjectRenameFormat() string`

GetSubjectRenameFormat returns the SubjectRenameFormat field if non-nil, zero value otherwise.

### GetSubjectRenameFormatOk

`func (o *UpdateExporterRequest) GetSubjectRenameFormatOk() (*string, bool)`

GetSubjectRenameFormatOk returns a tuple with the SubjectRenameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRenameFormat

`func (o *UpdateExporterRequest) SetSubjectRenameFormat(v string)`

SetSubjectRenameFormat sets SubjectRenameFormat field to given value.

### HasSubjectRenameFormat

`func (o *UpdateExporterRequest) HasSubjectRenameFormat() bool`

HasSubjectRenameFormat returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateExporterRequest) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateExporterRequest) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateExporterRequest) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateExporterRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


