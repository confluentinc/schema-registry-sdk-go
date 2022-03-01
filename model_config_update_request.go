/*
 * Confluent Schema Registry
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConfigUpdateRequest struct for ConfigUpdateRequest
type ConfigUpdateRequest struct {
	// Compatability Level
	Compatibility *string `json:"compatibility,omitempty"`
}

// NewConfigUpdateRequest instantiates a new ConfigUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigUpdateRequest() *ConfigUpdateRequest {
	this := ConfigUpdateRequest{}
	return &this
}

// NewConfigUpdateRequestWithDefaults instantiates a new ConfigUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigUpdateRequestWithDefaults() *ConfigUpdateRequest {
	this := ConfigUpdateRequest{}
	return &this
}

// GetCompatibility returns the Compatibility field value if set, zero value otherwise.
func (o *ConfigUpdateRequest) GetCompatibility() string {
	if o == nil || o.Compatibility == nil {
		var ret string
		return ret
	}
	return *o.Compatibility
}

// GetCompatibilityOk returns a tuple with the Compatibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigUpdateRequest) GetCompatibilityOk() (*string, bool) {
	if o == nil || o.Compatibility == nil {
		return nil, false
	}
	return o.Compatibility, true
}

// HasCompatibility returns a boolean if a field has been set.
func (o *ConfigUpdateRequest) HasCompatibility() bool {
	if o != nil && o.Compatibility != nil {
		return true
	}

	return false
}

// SetCompatibility gets a reference to the given string and assigns it to the Compatibility field.
func (o *ConfigUpdateRequest) SetCompatibility(v string) {
	o.Compatibility = &v
}

func (o ConfigUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Compatibility != nil {
		toSerialize["compatibility"] = o.Compatibility
	}
	return json.Marshal(toSerialize)
}

type NullableConfigUpdateRequest struct {
	value *ConfigUpdateRequest
	isSet bool
}

func (v NullableConfigUpdateRequest) Get() *ConfigUpdateRequest {
	return v.value
}

func (v *NullableConfigUpdateRequest) Set(val *ConfigUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigUpdateRequest(val *ConfigUpdateRequest) *NullableConfigUpdateRequest {
	return &NullableConfigUpdateRequest{value: val, isSet: true}
}

func (v NullableConfigUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


