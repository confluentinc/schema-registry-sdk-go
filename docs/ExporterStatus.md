# ExporterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Ts** | Pointer to **int64** |  | [optional] 
**Trace** | Pointer to **string** |  | [optional] 

## Methods

### NewExporterStatus

`func NewExporterStatus() *ExporterStatus`

NewExporterStatus instantiates a new ExporterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExporterStatusWithDefaults

`func NewExporterStatusWithDefaults() *ExporterStatus`

NewExporterStatusWithDefaults instantiates a new ExporterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExporterStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExporterStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExporterStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExporterStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ExporterStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExporterStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExporterStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ExporterStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOffset

`func (o *ExporterStatus) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ExporterStatus) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ExporterStatus) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ExporterStatus) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTs

`func (o *ExporterStatus) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ExporterStatus) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ExporterStatus) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ExporterStatus) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTrace

`func (o *ExporterStatus) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ExporterStatus) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ExporterStatus) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ExporterStatus) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


