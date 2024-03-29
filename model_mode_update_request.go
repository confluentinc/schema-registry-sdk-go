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

// ModeUpdateRequest struct for ModeUpdateRequest
type ModeUpdateRequest struct {
	Mode *string `json:"mode,omitempty"`
}

// NewModeUpdateRequest instantiates a new ModeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModeUpdateRequest() *ModeUpdateRequest {
	this := ModeUpdateRequest{}
	return &this
}

// NewModeUpdateRequestWithDefaults instantiates a new ModeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModeUpdateRequestWithDefaults() *ModeUpdateRequest {
	this := ModeUpdateRequest{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ModeUpdateRequest) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeUpdateRequest) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ModeUpdateRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ModeUpdateRequest) SetMode(v string) {
	o.Mode = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ModeUpdateRequest) Redact() {
    o.recurseRedact(o.Mode)
}

func (o *ModeUpdateRequest) recurseRedact(v interface{}) {
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

func (o ModeUpdateRequest) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ModeUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableModeUpdateRequest struct {
	value *ModeUpdateRequest
	isSet bool
}

func (v NullableModeUpdateRequest) Get() *ModeUpdateRequest {
	return v.value
}

func (v *NullableModeUpdateRequest) Set(val *ModeUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModeUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModeUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModeUpdateRequest(val *ModeUpdateRequest) *NullableModeUpdateRequest {
	return &NullableModeUpdateRequest{value: val, isSet: true}
}

func (v NullableModeUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModeUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


