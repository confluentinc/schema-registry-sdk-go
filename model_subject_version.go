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

// SubjectVersion struct for SubjectVersion
type SubjectVersion struct {
	Subject *string `json:"subject,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewSubjectVersion instantiates a new SubjectVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectVersion() *SubjectVersion {
	this := SubjectVersion{}
	return &this
}

// NewSubjectVersionWithDefaults instantiates a new SubjectVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectVersionWithDefaults() *SubjectVersion {
	this := SubjectVersion{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SubjectVersion) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectVersion) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SubjectVersion) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SubjectVersion) SetSubject(v string) {
	o.Subject = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SubjectVersion) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectVersion) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SubjectVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SubjectVersion) SetVersion(v int32) {
	o.Version = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SubjectVersion) Redact() {
    o.recurseRedact(o.Subject)
    o.recurseRedact(o.Version)
}

func (o *SubjectVersion) recurseRedact(v interface{}) {
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

func (o SubjectVersion) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SubjectVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSubjectVersion struct {
	value *SubjectVersion
	isSet bool
}

func (v NullableSubjectVersion) Get() *SubjectVersion {
	return v.value
}

func (v *NullableSubjectVersion) Set(val *SubjectVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectVersion(val *SubjectVersion) *NullableSubjectVersion {
	return &NullableSubjectVersion{value: val, isSet: true}
}

func (v NullableSubjectVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


