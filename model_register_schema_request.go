// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Confluent Schema Registry

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemaregistry

import (
	"encoding/json"
)

import (
	"reflect"
)

// RegisterSchemaRequest Schema register request
type RegisterSchemaRequest struct {
	// Version number
	Version *int32 `json:"version,omitempty"`
	// Globally unique identifier of the schema
	Id *int32 `json:"id,omitempty"`
	// Schema type
	SchemaType *string `json:"schemaType,omitempty"`
	// References to other schemas
	References *[]SchemaReference `json:"references,omitempty"`
	Metadata NullableMetadata `json:"metadata,omitempty"`
	RuleSet NullableRuleSet `json:"ruleSet,omitempty"`
	// Schema definition string
	Schema *string `json:"schema,omitempty"`
}

// NewRegisterSchemaRequest instantiates a new RegisterSchemaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterSchemaRequest() *RegisterSchemaRequest {
	this := RegisterSchemaRequest{}
	return &this
}

// NewRegisterSchemaRequestWithDefaults instantiates a new RegisterSchemaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterSchemaRequestWithDefaults() *RegisterSchemaRequest {
	this := RegisterSchemaRequest{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RegisterSchemaRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSchemaRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RegisterSchemaRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *RegisterSchemaRequest) SetVersion(v int32) {
	o.Version = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegisterSchemaRequest) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSchemaRequest) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegisterSchemaRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RegisterSchemaRequest) SetId(v int32) {
	o.Id = &v
}

// GetSchemaType returns the SchemaType field value if set, zero value otherwise.
func (o *RegisterSchemaRequest) GetSchemaType() string {
	if o == nil || o.SchemaType == nil {
		var ret string
		return ret
	}
	return *o.SchemaType
}

// GetSchemaTypeOk returns a tuple with the SchemaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSchemaRequest) GetSchemaTypeOk() (*string, bool) {
	if o == nil || o.SchemaType == nil {
		return nil, false
	}
	return o.SchemaType, true
}

// HasSchemaType returns a boolean if a field has been set.
func (o *RegisterSchemaRequest) HasSchemaType() bool {
	if o != nil && o.SchemaType != nil {
		return true
	}

	return false
}

// SetSchemaType gets a reference to the given string and assigns it to the SchemaType field.
func (o *RegisterSchemaRequest) SetSchemaType(v string) {
	o.SchemaType = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *RegisterSchemaRequest) GetReferences() []SchemaReference {
	if o == nil || o.References == nil {
		var ret []SchemaReference
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSchemaRequest) GetReferencesOk() (*[]SchemaReference, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *RegisterSchemaRequest) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []SchemaReference and assigns it to the References field.
func (o *RegisterSchemaRequest) SetReferences(v []SchemaReference) {
	o.References = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterSchemaRequest) GetMetadata() Metadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret Metadata
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterSchemaRequest) GetMetadataOk() (*Metadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *RegisterSchemaRequest) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableMetadata and assigns it to the Metadata field.
func (o *RegisterSchemaRequest) SetMetadata(v Metadata) {
	o.Metadata.Set(&v)
}
// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *RegisterSchemaRequest) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *RegisterSchemaRequest) UnsetMetadata() {
	o.Metadata.Unset()
}

// GetRuleSet returns the RuleSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterSchemaRequest) GetRuleSet() RuleSet {
	if o == nil || o.RuleSet.Get() == nil {
		var ret RuleSet
		return ret
	}
	return *o.RuleSet.Get()
}

// GetRuleSetOk returns a tuple with the RuleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterSchemaRequest) GetRuleSetOk() (*RuleSet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RuleSet.Get(), o.RuleSet.IsSet()
}

// HasRuleSet returns a boolean if a field has been set.
func (o *RegisterSchemaRequest) HasRuleSet() bool {
	if o != nil && o.RuleSet.IsSet() {
		return true
	}

	return false
}

// SetRuleSet gets a reference to the given NullableRuleSet and assigns it to the RuleSet field.
func (o *RegisterSchemaRequest) SetRuleSet(v RuleSet) {
	o.RuleSet.Set(&v)
}
// SetRuleSetNil sets the value for RuleSet to be an explicit nil
func (o *RegisterSchemaRequest) SetRuleSetNil() {
	o.RuleSet.Set(nil)
}

// UnsetRuleSet ensures that no value is present for RuleSet, not even an explicit nil
func (o *RegisterSchemaRequest) UnsetRuleSet() {
	o.RuleSet.Unset()
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *RegisterSchemaRequest) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSchemaRequest) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *RegisterSchemaRequest) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *RegisterSchemaRequest) SetSchema(v string) {
	o.Schema = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *RegisterSchemaRequest) Redact() {
    o.recurseRedact(o.Version)
    o.recurseRedact(o.Id)
    o.recurseRedact(o.SchemaType)
    o.recurseRedact(o.References)
    o.recurseRedact(o.Metadata)
    o.recurseRedact(o.RuleSet)
    o.recurseRedact(o.Schema)
}

func (o *RegisterSchemaRequest) recurseRedact(v interface{}) {
    type redactor interface {
        Redact()
    }
    if r, ok := v.(redactor); ok {
        r.Redact()
    } else {
        val := reflect.ValueOf(v)
        if val.Kind() == reflect.Ptr {
            val = val.Elem()
        }
        switch val.Kind() {
        case reflect.Slice, reflect.Array:
            for i := 0; i < val.Len(); i++ {
                // support data types declared without pointers
                o.recurseRedact(val.Index(i).Interface())
                // ... and data types that were declared without but need pointers (for Redact)
                if val.Index(i).CanAddr() {
                    o.recurseRedact(val.Index(i).Addr().Interface())
                }
            }
        }
    }
}

func (o RegisterSchemaRequest) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o RegisterSchemaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SchemaType != nil {
		toSerialize["schemaType"] = o.SchemaType
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.Metadata.IsSet() {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	if o.RuleSet.IsSet() {
		toSerialize["ruleSet"] = o.RuleSet.Get()
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterSchemaRequest struct {
	value *RegisterSchemaRequest
	isSet bool
}

func (v NullableRegisterSchemaRequest) Get() *RegisterSchemaRequest {
	return v.value
}

func (v *NullableRegisterSchemaRequest) Set(val *RegisterSchemaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterSchemaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterSchemaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterSchemaRequest(val *RegisterSchemaRequest) *NullableRegisterSchemaRequest {
	return &NullableRegisterSchemaRequest{value: val, isSet: true}
}

func (v NullableRegisterSchemaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterSchemaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


