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

// UpdateExporterResponse struct for UpdateExporterResponse
type UpdateExporterResponse struct {
	Name *string `json:"name,omitempty"`
}

// NewUpdateExporterResponse instantiates a new UpdateExporterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExporterResponse() *UpdateExporterResponse {
	this := UpdateExporterResponse{}
	return &this
}

// NewUpdateExporterResponseWithDefaults instantiates a new UpdateExporterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExporterResponseWithDefaults() *UpdateExporterResponse {
	this := UpdateExporterResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateExporterResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExporterResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateExporterResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateExporterResponse) SetName(v string) {
	o.Name = &v
}

func (o UpdateExporterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateExporterResponse struct {
	value *UpdateExporterResponse
	isSet bool
}

func (v NullableUpdateExporterResponse) Get() *UpdateExporterResponse {
	return v.value
}

func (v *NullableUpdateExporterResponse) Set(val *UpdateExporterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExporterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExporterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExporterResponse(val *UpdateExporterResponse) *NullableUpdateExporterResponse {
	return &NullableUpdateExporterResponse{value: val, isSet: true}
}

func (v NullableUpdateExporterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExporterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


