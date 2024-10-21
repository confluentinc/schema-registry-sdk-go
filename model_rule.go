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
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// Rule Rule
type Rule struct {
	// Rule name
	Name *string `json:"name,omitempty"`
	// Rule doc
	Doc *string `json:"doc,omitempty"`
	// Rule kind
	Kind *string `json:"kind,omitempty"`
	// Rule mode
	Mode *string `json:"mode,omitempty"`
	// Rule type
	Type *string `json:"type,omitempty"`
	// The tags to which this rule applies
	Tags *[]string `json:"tags,omitempty"`
	// Optional params for the rule
	Params *map[string]string `json:"params,omitempty"`
	// Rule expression
	Expr *string `json:"expr,omitempty"`
	// Rule action on success
	OnSuccess *string `json:"onSuccess,omitempty"`
	// Rule action on failure
	OnFailure *string `json:"onFailure,omitempty"`
	// Whether the rule is disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// NewRule instantiates a new Rule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRule() *Rule {
	this := Rule{}
	return &this
}

// NewRuleWithDefaults instantiates a new Rule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWithDefaults() *Rule {
	this := Rule{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Rule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Rule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Rule) SetName(v string) {
	o.Name = &v
}

// GetDoc returns the Doc field value if set, zero value otherwise.
func (o *Rule) GetDoc() string {
	if o == nil || o.Doc == nil {
		var ret string
		return ret
	}
	return *o.Doc
}

// GetDocOk returns a tuple with the Doc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetDocOk() (*string, bool) {
	if o == nil || o.Doc == nil {
		return nil, false
	}
	return o.Doc, true
}

// HasDoc returns a boolean if a field has been set.
func (o *Rule) HasDoc() bool {
	if o != nil && o.Doc != nil {
		return true
	}

	return false
}

// SetDoc gets a reference to the given string and assigns it to the Doc field.
func (o *Rule) SetDoc(v string) {
	o.Doc = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Rule) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Rule) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Rule) SetKind(v string) {
	o.Kind = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Rule) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Rule) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *Rule) SetMode(v string) {
	o.Mode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Rule) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Rule) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Rule) SetType(v string) {
	o.Type = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Rule) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Rule) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Rule) SetTags(v []string) {
	o.Tags = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *Rule) GetParams() map[string]string {
	if o == nil || o.Params == nil {
		var ret map[string]string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetParamsOk() (*map[string]string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *Rule) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *Rule) SetParams(v map[string]string) {
	o.Params = &v
}

// GetExpr returns the Expr field value if set, zero value otherwise.
func (o *Rule) GetExpr() string {
	if o == nil || o.Expr == nil {
		var ret string
		return ret
	}
	return *o.Expr
}

// GetExprOk returns a tuple with the Expr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetExprOk() (*string, bool) {
	if o == nil || o.Expr == nil {
		return nil, false
	}
	return o.Expr, true
}

// HasExpr returns a boolean if a field has been set.
func (o *Rule) HasExpr() bool {
	if o != nil && o.Expr != nil {
		return true
	}

	return false
}

// SetExpr gets a reference to the given string and assigns it to the Expr field.
func (o *Rule) SetExpr(v string) {
	o.Expr = &v
}

// GetOnSuccess returns the OnSuccess field value if set, zero value otherwise.
func (o *Rule) GetOnSuccess() string {
	if o == nil || o.OnSuccess == nil {
		var ret string
		return ret
	}
	return *o.OnSuccess
}

// GetOnSuccessOk returns a tuple with the OnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetOnSuccessOk() (*string, bool) {
	if o == nil || o.OnSuccess == nil {
		return nil, false
	}
	return o.OnSuccess, true
}

// HasOnSuccess returns a boolean if a field has been set.
func (o *Rule) HasOnSuccess() bool {
	if o != nil && o.OnSuccess != nil {
		return true
	}

	return false
}

// SetOnSuccess gets a reference to the given string and assigns it to the OnSuccess field.
func (o *Rule) SetOnSuccess(v string) {
	o.OnSuccess = &v
}

// GetOnFailure returns the OnFailure field value if set, zero value otherwise.
func (o *Rule) GetOnFailure() string {
	if o == nil || o.OnFailure == nil {
		var ret string
		return ret
	}
	return *o.OnFailure
}

// GetOnFailureOk returns a tuple with the OnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetOnFailureOk() (*string, bool) {
	if o == nil || o.OnFailure == nil {
		return nil, false
	}
	return o.OnFailure, true
}

// HasOnFailure returns a boolean if a field has been set.
func (o *Rule) HasOnFailure() bool {
	if o != nil && o.OnFailure != nil {
		return true
	}

	return false
}

// SetOnFailure gets a reference to the given string and assigns it to the OnFailure field.
func (o *Rule) SetOnFailure(v string) {
	o.OnFailure = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Rule) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *Rule) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Rule) SetDisabled(v bool) {
	o.Disabled = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *Rule) Redact() {
    o.recurseRedact(o.Name)
    o.recurseRedact(o.Doc)
    o.recurseRedact(o.Kind)
    o.recurseRedact(o.Mode)
    o.recurseRedact(o.Type)
    o.recurseRedact(o.Tags)
    o.recurseRedact(o.Params)
    o.recurseRedact(o.Expr)
    o.recurseRedact(o.OnSuccess)
    o.recurseRedact(o.OnFailure)
    o.recurseRedact(o.Disabled)
}

func (o *Rule) recurseRedact(v interface{}) {
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

func (o Rule) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o Rule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Doc != nil {
		toSerialize["doc"] = o.Doc
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.Expr != nil {
		toSerialize["expr"] = o.Expr
	}
	if o.OnSuccess != nil {
		toSerialize["onSuccess"] = o.OnSuccess
	}
	if o.OnFailure != nil {
		toSerialize["onFailure"] = o.OnFailure
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableRule struct {
	value *Rule
	isSet bool
}

func (v NullableRule) Get() *Rule {
	return v.value
}

func (v *NullableRule) Set(val *Rule) {
	v.value = val
	v.isSet = true
}

func (v NullableRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRule(val *Rule) *NullableRule {
	return &NullableRule{value: val, isSet: true}
}

func (v NullableRule) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


