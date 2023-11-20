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

// CreateKekRequest struct for CreateKekRequest
type CreateKekRequest struct {
	Name *string `json:"name,omitempty"`
	KmsType *string `json:"kmsType,omitempty"`
	KmsKeyId *string `json:"kmsKeyId,omitempty"`
	KmsProps *map[string]string `json:"kmsProps,omitempty"`
	Doc *string `json:"doc,omitempty"`
	Shared *bool `json:"shared,omitempty"`
}

// NewCreateKekRequest instantiates a new CreateKekRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKekRequest() *CreateKekRequest {
	this := CreateKekRequest{}
	return &this
}

// NewCreateKekRequestWithDefaults instantiates a new CreateKekRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKekRequestWithDefaults() *CreateKekRequest {
	this := CreateKekRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateKekRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKekRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateKekRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateKekRequest) SetName(v string) {
	o.Name = &v
}

// GetKmsType returns the KmsType field value if set, zero value otherwise.
func (o *CreateKekRequest) GetKmsType() string {
	if o == nil || o.KmsType == nil {
		var ret string
		return ret
	}
	return *o.KmsType
}

// GetKmsTypeOk returns a tuple with the KmsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKekRequest) GetKmsTypeOk() (*string, bool) {
	if o == nil || o.KmsType == nil {
		return nil, false
	}
	return o.KmsType, true
}

// HasKmsType returns a boolean if a field has been set.
func (o *CreateKekRequest) HasKmsType() bool {
	if o != nil && o.KmsType != nil {
		return true
	}

	return false
}

// SetKmsType gets a reference to the given string and assigns it to the KmsType field.
func (o *CreateKekRequest) SetKmsType(v string) {
	o.KmsType = &v
}

// GetKmsKeyId returns the KmsKeyId field value if set, zero value otherwise.
func (o *CreateKekRequest) GetKmsKeyId() string {
	if o == nil || o.KmsKeyId == nil {
		var ret string
		return ret
	}
	return *o.KmsKeyId
}

// GetKmsKeyIdOk returns a tuple with the KmsKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKekRequest) GetKmsKeyIdOk() (*string, bool) {
	if o == nil || o.KmsKeyId == nil {
		return nil, false
	}
	return o.KmsKeyId, true
}

// HasKmsKeyId returns a boolean if a field has been set.
func (o *CreateKekRequest) HasKmsKeyId() bool {
	if o != nil && o.KmsKeyId != nil {
		return true
	}

	return false
}

// SetKmsKeyId gets a reference to the given string and assigns it to the KmsKeyId field.
func (o *CreateKekRequest) SetKmsKeyId(v string) {
	o.KmsKeyId = &v
}

// GetKmsProps returns the KmsProps field value if set, zero value otherwise.
func (o *CreateKekRequest) GetKmsProps() map[string]string {
	if o == nil || o.KmsProps == nil {
		var ret map[string]string
		return ret
	}
	return *o.KmsProps
}

// GetKmsPropsOk returns a tuple with the KmsProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKekRequest) GetKmsPropsOk() (*map[string]string, bool) {
	if o == nil || o.KmsProps == nil {
		return nil, false
	}
	return o.KmsProps, true
}

// HasKmsProps returns a boolean if a field has been set.
func (o *CreateKekRequest) HasKmsProps() bool {
	if o != nil && o.KmsProps != nil {
		return true
	}

	return false
}

// SetKmsProps gets a reference to the given map[string]string and assigns it to the KmsProps field.
func (o *CreateKekRequest) SetKmsProps(v map[string]string) {
	o.KmsProps = &v
}

// GetDoc returns the Doc field value if set, zero value otherwise.
func (o *CreateKekRequest) GetDoc() string {
	if o == nil || o.Doc == nil {
		var ret string
		return ret
	}
	return *o.Doc
}

// GetDocOk returns a tuple with the Doc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKekRequest) GetDocOk() (*string, bool) {
	if o == nil || o.Doc == nil {
		return nil, false
	}
	return o.Doc, true
}

// HasDoc returns a boolean if a field has been set.
func (o *CreateKekRequest) HasDoc() bool {
	if o != nil && o.Doc != nil {
		return true
	}

	return false
}

// SetDoc gets a reference to the given string and assigns it to the Doc field.
func (o *CreateKekRequest) SetDoc(v string) {
	o.Doc = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *CreateKekRequest) GetShared() bool {
	if o == nil || o.Shared == nil {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKekRequest) GetSharedOk() (*bool, bool) {
	if o == nil || o.Shared == nil {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *CreateKekRequest) HasShared() bool {
	if o != nil && o.Shared != nil {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *CreateKekRequest) SetShared(v bool) {
	o.Shared = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CreateKekRequest) Redact() {
    o.recurseRedact(o.Name)
    o.recurseRedact(o.KmsType)
    o.recurseRedact(o.KmsKeyId)
    o.recurseRedact(o.KmsProps)
    o.recurseRedact(o.Doc)
    o.recurseRedact(o.Shared)
}

func (o *CreateKekRequest) recurseRedact(v interface{}) {
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

func (o CreateKekRequest) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o CreateKekRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.KmsType != nil {
		toSerialize["kmsType"] = o.KmsType
	}
	if o.KmsKeyId != nil {
		toSerialize["kmsKeyId"] = o.KmsKeyId
	}
	if o.KmsProps != nil {
		toSerialize["kmsProps"] = o.KmsProps
	}
	if o.Doc != nil {
		toSerialize["doc"] = o.Doc
	}
	if o.Shared != nil {
		toSerialize["shared"] = o.Shared
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKekRequest struct {
	value *CreateKekRequest
	isSet bool
}

func (v NullableCreateKekRequest) Get() *CreateKekRequest {
	return v.value
}

func (v *NullableCreateKekRequest) Set(val *CreateKekRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKekRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKekRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKekRequest(val *CreateKekRequest) *NullableCreateKekRequest {
	return &NullableCreateKekRequest{value: val, isSet: true}
}

func (v NullableCreateKekRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKekRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


