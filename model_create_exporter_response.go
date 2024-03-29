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

// CreateExporterResponse struct for CreateExporterResponse
type CreateExporterResponse struct {
	Name *string `json:"name,omitempty"`
}

// NewCreateExporterResponse instantiates a new CreateExporterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExporterResponse() *CreateExporterResponse {
	this := CreateExporterResponse{}
	return &this
}

// NewCreateExporterResponseWithDefaults instantiates a new CreateExporterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExporterResponseWithDefaults() *CreateExporterResponse {
	this := CreateExporterResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateExporterResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExporterResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateExporterResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateExporterResponse) SetName(v string) {
	o.Name = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CreateExporterResponse) Redact() {
    o.recurseRedact(o.Name)
}

func (o *CreateExporterResponse) recurseRedact(v interface{}) {
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

func (o CreateExporterResponse) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o CreateExporterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateExporterResponse struct {
	value *CreateExporterResponse
	isSet bool
}

func (v NullableCreateExporterResponse) Get() *CreateExporterResponse {
	return v.value
}

func (v *NullableCreateExporterResponse) Set(val *CreateExporterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExporterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExporterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExporterResponse(val *CreateExporterResponse) *NullableCreateExporterResponse {
	return &NullableCreateExporterResponse{value: val, isSet: true}
}

func (v NullableCreateExporterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateExporterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


