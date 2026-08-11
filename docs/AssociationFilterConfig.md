# AssociationFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicsToInclude** | Pointer to **[]string** |  | [optional] 
**TopicsToExclude** | Pointer to **[]string** |  | [optional] 
**SubjectRenameFormat** | Pointer to **string** |  | [optional] 
**ContextType** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**SourceSRDeployment** | Pointer to **string** |  | [optional] 
**SourceLSRC** | Pointer to **string** |  | [optional] 
**SourceSRBackingKafkaId** | Pointer to **string** |  | [optional] 
**SourceSRGroupId** | Pointer to **string** |  | [optional] 

## Methods

### NewAssociationFilterConfig

`func NewAssociationFilterConfig() *AssociationFilterConfig`

NewAssociationFilterConfig instantiates a new AssociationFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationFilterConfigWithDefaults

`func NewAssociationFilterConfigWithDefaults() *AssociationFilterConfig`

NewAssociationFilterConfigWithDefaults instantiates a new AssociationFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicsToInclude

`func (o *AssociationFilterConfig) GetTopicsToInclude() []string`

GetTopicsToInclude returns the TopicsToInclude field if non-nil, zero value otherwise.

### GetTopicsToIncludeOk

`func (o *AssociationFilterConfig) GetTopicsToIncludeOk() (*[]string, bool)`

GetTopicsToIncludeOk returns a tuple with the TopicsToInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsToInclude

`func (o *AssociationFilterConfig) SetTopicsToInclude(v []string)`

SetTopicsToInclude sets TopicsToInclude field to given value.

### HasTopicsToInclude

`func (o *AssociationFilterConfig) HasTopicsToInclude() bool`

HasTopicsToInclude returns a boolean if a field has been set.

### GetTopicsToExclude

`func (o *AssociationFilterConfig) GetTopicsToExclude() []string`

GetTopicsToExclude returns the TopicsToExclude field if non-nil, zero value otherwise.

### GetTopicsToExcludeOk

`func (o *AssociationFilterConfig) GetTopicsToExcludeOk() (*[]string, bool)`

GetTopicsToExcludeOk returns a tuple with the TopicsToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsToExclude

`func (o *AssociationFilterConfig) SetTopicsToExclude(v []string)`

SetTopicsToExclude sets TopicsToExclude field to given value.

### HasTopicsToExclude

`func (o *AssociationFilterConfig) HasTopicsToExclude() bool`

HasTopicsToExclude returns a boolean if a field has been set.

### GetSubjectRenameFormat

`func (o *AssociationFilterConfig) GetSubjectRenameFormat() string`

GetSubjectRenameFormat returns the SubjectRenameFormat field if non-nil, zero value otherwise.

### GetSubjectRenameFormatOk

`func (o *AssociationFilterConfig) GetSubjectRenameFormatOk() (*string, bool)`

GetSubjectRenameFormatOk returns a tuple with the SubjectRenameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRenameFormat

`func (o *AssociationFilterConfig) SetSubjectRenameFormat(v string)`

SetSubjectRenameFormat sets SubjectRenameFormat field to given value.

### HasSubjectRenameFormat

`func (o *AssociationFilterConfig) HasSubjectRenameFormat() bool`

HasSubjectRenameFormat returns a boolean if a field has been set.

### GetContextType

`func (o *AssociationFilterConfig) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *AssociationFilterConfig) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *AssociationFilterConfig) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *AssociationFilterConfig) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetContext

`func (o *AssociationFilterConfig) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AssociationFilterConfig) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AssociationFilterConfig) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AssociationFilterConfig) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSourceSRDeployment

`func (o *AssociationFilterConfig) GetSourceSRDeployment() string`

GetSourceSRDeployment returns the SourceSRDeployment field if non-nil, zero value otherwise.

### GetSourceSRDeploymentOk

`func (o *AssociationFilterConfig) GetSourceSRDeploymentOk() (*string, bool)`

GetSourceSRDeploymentOk returns a tuple with the SourceSRDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSRDeployment

`func (o *AssociationFilterConfig) SetSourceSRDeployment(v string)`

SetSourceSRDeployment sets SourceSRDeployment field to given value.

### HasSourceSRDeployment

`func (o *AssociationFilterConfig) HasSourceSRDeployment() bool`

HasSourceSRDeployment returns a boolean if a field has been set.

### GetSourceLSRC

`func (o *AssociationFilterConfig) GetSourceLSRC() string`

GetSourceLSRC returns the SourceLSRC field if non-nil, zero value otherwise.

### GetSourceLSRCOk

`func (o *AssociationFilterConfig) GetSourceLSRCOk() (*string, bool)`

GetSourceLSRCOk returns a tuple with the SourceLSRC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLSRC

`func (o *AssociationFilterConfig) SetSourceLSRC(v string)`

SetSourceLSRC sets SourceLSRC field to given value.

### HasSourceLSRC

`func (o *AssociationFilterConfig) HasSourceLSRC() bool`

HasSourceLSRC returns a boolean if a field has been set.

### GetSourceSRBackingKafkaId

`func (o *AssociationFilterConfig) GetSourceSRBackingKafkaId() string`

GetSourceSRBackingKafkaId returns the SourceSRBackingKafkaId field if non-nil, zero value otherwise.

### GetSourceSRBackingKafkaIdOk

`func (o *AssociationFilterConfig) GetSourceSRBackingKafkaIdOk() (*string, bool)`

GetSourceSRBackingKafkaIdOk returns a tuple with the SourceSRBackingKafkaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSRBackingKafkaId

`func (o *AssociationFilterConfig) SetSourceSRBackingKafkaId(v string)`

SetSourceSRBackingKafkaId sets SourceSRBackingKafkaId field to given value.

### HasSourceSRBackingKafkaId

`func (o *AssociationFilterConfig) HasSourceSRBackingKafkaId() bool`

HasSourceSRBackingKafkaId returns a boolean if a field has been set.

### GetSourceSRGroupId

`func (o *AssociationFilterConfig) GetSourceSRGroupId() string`

GetSourceSRGroupId returns the SourceSRGroupId field if non-nil, zero value otherwise.

### GetSourceSRGroupIdOk

`func (o *AssociationFilterConfig) GetSourceSRGroupIdOk() (*string, bool)`

GetSourceSRGroupIdOk returns a tuple with the SourceSRGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSRGroupId

`func (o *AssociationFilterConfig) SetSourceSRGroupId(v string)`

SetSourceSRGroupId sets SourceSRGroupId field to given value.

### HasSourceSRGroupId

`func (o *AssociationFilterConfig) HasSourceSRGroupId() bool`

HasSourceSRGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


