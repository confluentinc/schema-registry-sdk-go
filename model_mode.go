/*
 * Confluent Schema Registry
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemaregistry

import (
	"encoding/json"
)

// Mode struct for Mode
type Mode struct {
	Mode *string `json:"mode,omitempty"`
}

// NewMode instantiates a new Mode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMode() *Mode {
	this := Mode{}
	return &this
}

// NewModeWithDefaults instantiates a new Mode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModeWithDefaults() *Mode {
	this := Mode{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Mode) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mode) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Mode) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *Mode) SetMode(v string) {
	o.Mode = &v
}

func (o Mode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableMode struct {
	value *Mode
	isSet bool
}

func (v NullableMode) Get() *Mode {
	return v.value
}

func (v *NullableMode) Set(val *Mode) {
	v.value = val
	v.isSet = true
}

func (v NullableMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMode(val *Mode) *NullableMode {
	return &NullableMode{value: val, isSet: true}
}

func (v NullableMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

