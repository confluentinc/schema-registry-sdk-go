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

// SchemaReference struct for SchemaReference
type SchemaReference struct {
	Name *string `json:"name,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewSchemaReference instantiates a new SchemaReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaReference() *SchemaReference {
	this := SchemaReference{}
	return &this
}

// NewSchemaReferenceWithDefaults instantiates a new SchemaReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaReferenceWithDefaults() *SchemaReference {
	this := SchemaReference{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemaReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemaReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemaReference) SetName(v string) {
	o.Name = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SchemaReference) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SchemaReference) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SchemaReference) SetSubject(v string) {
	o.Subject = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SchemaReference) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SchemaReference) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SchemaReference) SetVersion(v int32) {
	o.Version = &v
}

func (o SchemaReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaReference struct {
	value *SchemaReference
	isSet bool
}

func (v NullableSchemaReference) Get() *SchemaReference {
	return v.value
}

func (v *NullableSchemaReference) Set(val *SchemaReference) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaReference) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaReference(val *SchemaReference) *NullableSchemaReference {
	return &NullableSchemaReference{value: val, isSet: true}
}

func (v NullableSchemaReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


