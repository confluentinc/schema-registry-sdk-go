# ExporterInfo

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

### NewExporterInfo

`func NewExporterInfo() *ExporterInfo`

NewExporterInfo instantiates a new ExporterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExporterInfoWithDefaults

`func NewExporterInfoWithDefaults() *ExporterInfo`

NewExporterInfoWithDefaults instantiates a new ExporterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExporterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExporterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExporterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExporterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubjects

`func (o *ExporterInfo) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *ExporterInfo) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *ExporterInfo) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *ExporterInfo) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetContextType

`func (o *ExporterInfo) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *ExporterInfo) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *ExporterInfo) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *ExporterInfo) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetContext

`func (o *ExporterInfo) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ExporterInfo) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ExporterInfo) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ExporterInfo) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSubjectRenameFormat

`func (o *ExporterInfo) GetSubjectRenameFormat() string`

GetSubjectRenameFormat returns the SubjectRenameFormat field if non-nil, zero value otherwise.

### GetSubjectRenameFormatOk

`func (o *ExporterInfo) GetSubjectRenameFormatOk() (*string, bool)`

GetSubjectRenameFormatOk returns a tuple with the SubjectRenameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRenameFormat

`func (o *ExporterInfo) SetSubjectRenameFormat(v string)`

SetSubjectRenameFormat sets SubjectRenameFormat field to given value.

### HasSubjectRenameFormat

`func (o *ExporterInfo) HasSubjectRenameFormat() bool`

HasSubjectRenameFormat returns a boolean if a field has been set.

### GetConfig

`func (o *ExporterInfo) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExporterInfo) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExporterInfo) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExporterInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


