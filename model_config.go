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

// Config Config
type Config struct {
	// Compatibility Level
	CompatibilityLevel *string `json:"compatibilityLevel,omitempty"`
	CompatibilityGroup *string `json:"compatibilityGroup,omitempty"`
	DefaultMetadata NullableMetadata `json:"defaultMetadata,omitempty"`
	OverrideMetadata NullableMetadata `json:"overrideMetadata,omitempty"`
	DefaultRuleSet NullableRuleSet `json:"defaultRuleSet,omitempty"`
	OverrideRuleSet NullableRuleSet `json:"overrideRuleSet,omitempty"`
}

// NewConfig instantiates a new Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfig() *Config {
	this := Config{}
	return &this
}

// NewConfigWithDefaults instantiates a new Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigWithDefaults() *Config {
	this := Config{}
	return &this
}

// GetCompatibilityLevel returns the CompatibilityLevel field value if set, zero value otherwise.
func (o *Config) GetCompatibilityLevel() string {
	if o == nil || o.CompatibilityLevel == nil {
		var ret string
		return ret
	}
	return *o.CompatibilityLevel
}

// GetCompatibilityLevelOk returns a tuple with the CompatibilityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config) GetCompatibilityLevelOk() (*string, bool) {
	if o == nil || o.CompatibilityLevel == nil {
		return nil, false
	}
	return o.CompatibilityLevel, true
}

// HasCompatibilityLevel returns a boolean if a field has been set.
func (o *Config) HasCompatibilityLevel() bool {
	if o != nil && o.CompatibilityLevel != nil {
		return true
	}

	return false
}

// SetCompatibilityLevel gets a reference to the given string and assigns it to the CompatibilityLevel field.
func (o *Config) SetCompatibilityLevel(v string) {
	o.CompatibilityLevel = &v
}

// GetCompatibilityGroup returns the CompatibilityGroup field value if set, zero value otherwise.
func (o *Config) GetCompatibilityGroup() string {
	if o == nil || o.CompatibilityGroup == nil {
		var ret string
		return ret
	}
	return *o.CompatibilityGroup
}

// GetCompatibilityGroupOk returns a tuple with the CompatibilityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config) GetCompatibilityGroupOk() (*string, bool) {
	if o == nil || o.CompatibilityGroup == nil {
		return nil, false
	}
	return o.CompatibilityGroup, true
}

// HasCompatibilityGroup returns a boolean if a field has been set.
func (o *Config) HasCompatibilityGroup() bool {
	if o != nil && o.CompatibilityGroup != nil {
		return true
	}

	return false
}

// SetCompatibilityGroup gets a reference to the given string and assigns it to the CompatibilityGroup field.
func (o *Config) SetCompatibilityGroup(v string) {
	o.CompatibilityGroup = &v
}

// GetDefaultMetadata returns the DefaultMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetDefaultMetadata() Metadata {
	if o == nil || o.DefaultMetadata.Get() == nil {
		var ret Metadata
		return ret
	}
	return *o.DefaultMetadata.Get()
}

// GetDefaultMetadataOk returns a tuple with the DefaultMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetDefaultMetadataOk() (*Metadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultMetadata.Get(), o.DefaultMetadata.IsSet()
}

// HasDefaultMetadata returns a boolean if a field has been set.
func (o *Config) HasDefaultMetadata() bool {
	if o != nil && o.DefaultMetadata.IsSet() {
		return true
	}

	return false
}

// SetDefaultMetadata gets a reference to the given NullableMetadata and assigns it to the DefaultMetadata field.
func (o *Config) SetDefaultMetadata(v Metadata) {
	o.DefaultMetadata.Set(&v)
}
// SetDefaultMetadataNil sets the value for DefaultMetadata to be an explicit nil
func (o *Config) SetDefaultMetadataNil() {
	o.DefaultMetadata.Set(nil)
}

// UnsetDefaultMetadata ensures that no value is present for DefaultMetadata, not even an explicit nil
func (o *Config) UnsetDefaultMetadata() {
	o.DefaultMetadata.Unset()
}

// GetOverrideMetadata returns the OverrideMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetOverrideMetadata() Metadata {
	if o == nil || o.OverrideMetadata.Get() == nil {
		var ret Metadata
		return ret
	}
	return *o.OverrideMetadata.Get()
}

// GetOverrideMetadataOk returns a tuple with the OverrideMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetOverrideMetadataOk() (*Metadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverrideMetadata.Get(), o.OverrideMetadata.IsSet()
}

// HasOverrideMetadata returns a boolean if a field has been set.
func (o *Config) HasOverrideMetadata() bool {
	if o != nil && o.OverrideMetadata.IsSet() {
		return true
	}

	return false
}

// SetOverrideMetadata gets a reference to the given NullableMetadata and assigns it to the OverrideMetadata field.
func (o *Config) SetOverrideMetadata(v Metadata) {
	o.OverrideMetadata.Set(&v)
}
// SetOverrideMetadataNil sets the value for OverrideMetadata to be an explicit nil
func (o *Config) SetOverrideMetadataNil() {
	o.OverrideMetadata.Set(nil)
}

// UnsetOverrideMetadata ensures that no value is present for OverrideMetadata, not even an explicit nil
func (o *Config) UnsetOverrideMetadata() {
	o.OverrideMetadata.Unset()
}

// GetDefaultRuleSet returns the DefaultRuleSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetDefaultRuleSet() RuleSet {
	if o == nil || o.DefaultRuleSet.Get() == nil {
		var ret RuleSet
		return ret
	}
	return *o.DefaultRuleSet.Get()
}

// GetDefaultRuleSetOk returns a tuple with the DefaultRuleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetDefaultRuleSetOk() (*RuleSet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultRuleSet.Get(), o.DefaultRuleSet.IsSet()
}

// HasDefaultRuleSet returns a boolean if a field has been set.
func (o *Config) HasDefaultRuleSet() bool {
	if o != nil && o.DefaultRuleSet.IsSet() {
		return true
	}

	return false
}

// SetDefaultRuleSet gets a reference to the given NullableRuleSet and assigns it to the DefaultRuleSet field.
func (o *Config) SetDefaultRuleSet(v RuleSet) {
	o.DefaultRuleSet.Set(&v)
}
// SetDefaultRuleSetNil sets the value for DefaultRuleSet to be an explicit nil
func (o *Config) SetDefaultRuleSetNil() {
	o.DefaultRuleSet.Set(nil)
}

// UnsetDefaultRuleSet ensures that no value is present for DefaultRuleSet, not even an explicit nil
func (o *Config) UnsetDefaultRuleSet() {
	o.DefaultRuleSet.Unset()
}

// GetOverrideRuleSet returns the OverrideRuleSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetOverrideRuleSet() RuleSet {
	if o == nil || o.OverrideRuleSet.Get() == nil {
		var ret RuleSet
		return ret
	}
	return *o.OverrideRuleSet.Get()
}

// GetOverrideRuleSetOk returns a tuple with the OverrideRuleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetOverrideRuleSetOk() (*RuleSet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverrideRuleSet.Get(), o.OverrideRuleSet.IsSet()
}

// HasOverrideRuleSet returns a boolean if a field has been set.
func (o *Config) HasOverrideRuleSet() bool {
	if o != nil && o.OverrideRuleSet.IsSet() {
		return true
	}

	return false
}

// SetOverrideRuleSet gets a reference to the given NullableRuleSet and assigns it to the OverrideRuleSet field.
func (o *Config) SetOverrideRuleSet(v RuleSet) {
	o.OverrideRuleSet.Set(&v)
}
// SetOverrideRuleSetNil sets the value for OverrideRuleSet to be an explicit nil
func (o *Config) SetOverrideRuleSetNil() {
	o.OverrideRuleSet.Set(nil)
}

// UnsetOverrideRuleSet ensures that no value is present for OverrideRuleSet, not even an explicit nil
func (o *Config) UnsetOverrideRuleSet() {
	o.OverrideRuleSet.Unset()
}

// Redact resets all sensitive fields to their zero value.
func (o *Config) Redact() {
    o.recurseRedact(o.CompatibilityLevel)
    o.recurseRedact(o.CompatibilityGroup)
    o.recurseRedact(o.DefaultMetadata)
    o.recurseRedact(o.OverrideMetadata)
    o.recurseRedact(o.DefaultRuleSet)
    o.recurseRedact(o.OverrideRuleSet)
}

func (o *Config) recurseRedact(v interface{}) {
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

func (o Config) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o Config) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompatibilityLevel != nil {
		toSerialize["compatibilityLevel"] = o.CompatibilityLevel
	}
	if o.CompatibilityGroup != nil {
		toSerialize["compatibilityGroup"] = o.CompatibilityGroup
	}
	if o.DefaultMetadata.IsSet() {
		toSerialize["defaultMetadata"] = o.DefaultMetadata.Get()
	}
	if o.OverrideMetadata.IsSet() {
		toSerialize["overrideMetadata"] = o.OverrideMetadata.Get()
	}
	if o.DefaultRuleSet.IsSet() {
		toSerialize["defaultRuleSet"] = o.DefaultRuleSet.Get()
	}
	if o.OverrideRuleSet.IsSet() {
		toSerialize["overrideRuleSet"] = o.OverrideRuleSet.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConfig struct {
	value *Config
	isSet bool
}

func (v NullableConfig) Get() *Config {
	return v.value
}

func (v *NullableConfig) Set(val *Config) {
	v.value = val
	v.isSet = true
}

func (v NullableConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfig(val *Config) *NullableConfig {
	return &NullableConfig{value: val, isSet: true}
}

func (v NullableConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


