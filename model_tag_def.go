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

// TagDef struct for TagDef
type TagDef struct {
	Category *string `json:"category,omitempty"`
	Guid *string `json:"guid,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	CreateTime *int64 `json:"createTime,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
	Version *int64 `json:"version,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	TypeVersion *string `json:"typeVersion,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
	Options *map[string]string `json:"options,omitempty"`
	AttributeDefs *[]AtlasAttributeDef `json:"attributeDefs,omitempty"`
	SuperTypes *[]string `json:"superTypes,omitempty"`
	EntityTypes *[]string `json:"entityTypes,omitempty"`
	SubTypes *[]string `json:"subTypes,omitempty"`
}

// NewTagDef instantiates a new TagDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagDef() *TagDef {
	this := TagDef{}
	return &this
}

// NewTagDefWithDefaults instantiates a new TagDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagDefWithDefaults() *TagDef {
	this := TagDef{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *TagDef) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *TagDef) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *TagDef) SetCategory(v string) {
	o.Category = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *TagDef) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *TagDef) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *TagDef) SetGuid(v string) {
	o.Guid = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *TagDef) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TagDef) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *TagDef) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *TagDef) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *TagDef) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *TagDef) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *TagDef) GetCreateTime() int64 {
	if o == nil || o.CreateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetCreateTimeOk() (*int64, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *TagDef) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *TagDef) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *TagDef) GetUpdateTime() int64 {
	if o == nil || o.UpdateTime == nil {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *TagDef) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *TagDef) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TagDef) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TagDef) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *TagDef) SetVersion(v int64) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagDef) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagDef) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagDef) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagDef) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagDef) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagDef) SetDescription(v string) {
	o.Description = &v
}

// GetTypeVersion returns the TypeVersion field value if set, zero value otherwise.
func (o *TagDef) GetTypeVersion() string {
	if o == nil || o.TypeVersion == nil {
		var ret string
		return ret
	}
	return *o.TypeVersion
}

// GetTypeVersionOk returns a tuple with the TypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetTypeVersionOk() (*string, bool) {
	if o == nil || o.TypeVersion == nil {
		return nil, false
	}
	return o.TypeVersion, true
}

// HasTypeVersion returns a boolean if a field has been set.
func (o *TagDef) HasTypeVersion() bool {
	if o != nil && o.TypeVersion != nil {
		return true
	}

	return false
}

// SetTypeVersion gets a reference to the given string and assigns it to the TypeVersion field.
func (o *TagDef) SetTypeVersion(v string) {
	o.TypeVersion = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *TagDef) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *TagDef) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *TagDef) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TagDef) GetOptions() map[string]string {
	if o == nil || o.Options == nil {
		var ret map[string]string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetOptionsOk() (*map[string]string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TagDef) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]string and assigns it to the Options field.
func (o *TagDef) SetOptions(v map[string]string) {
	o.Options = &v
}

// GetAttributeDefs returns the AttributeDefs field value if set, zero value otherwise.
func (o *TagDef) GetAttributeDefs() []AtlasAttributeDef {
	if o == nil || o.AttributeDefs == nil {
		var ret []AtlasAttributeDef
		return ret
	}
	return *o.AttributeDefs
}

// GetAttributeDefsOk returns a tuple with the AttributeDefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetAttributeDefsOk() (*[]AtlasAttributeDef, bool) {
	if o == nil || o.AttributeDefs == nil {
		return nil, false
	}
	return o.AttributeDefs, true
}

// HasAttributeDefs returns a boolean if a field has been set.
func (o *TagDef) HasAttributeDefs() bool {
	if o != nil && o.AttributeDefs != nil {
		return true
	}

	return false
}

// SetAttributeDefs gets a reference to the given []AtlasAttributeDef and assigns it to the AttributeDefs field.
func (o *TagDef) SetAttributeDefs(v []AtlasAttributeDef) {
	o.AttributeDefs = &v
}

// GetSuperTypes returns the SuperTypes field value if set, zero value otherwise.
func (o *TagDef) GetSuperTypes() []string {
	if o == nil || o.SuperTypes == nil {
		var ret []string
		return ret
	}
	return *o.SuperTypes
}

// GetSuperTypesOk returns a tuple with the SuperTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetSuperTypesOk() (*[]string, bool) {
	if o == nil || o.SuperTypes == nil {
		return nil, false
	}
	return o.SuperTypes, true
}

// HasSuperTypes returns a boolean if a field has been set.
func (o *TagDef) HasSuperTypes() bool {
	if o != nil && o.SuperTypes != nil {
		return true
	}

	return false
}

// SetSuperTypes gets a reference to the given []string and assigns it to the SuperTypes field.
func (o *TagDef) SetSuperTypes(v []string) {
	o.SuperTypes = &v
}

// GetEntityTypes returns the EntityTypes field value if set, zero value otherwise.
func (o *TagDef) GetEntityTypes() []string {
	if o == nil || o.EntityTypes == nil {
		var ret []string
		return ret
	}
	return *o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetEntityTypesOk() (*[]string, bool) {
	if o == nil || o.EntityTypes == nil {
		return nil, false
	}
	return o.EntityTypes, true
}

// HasEntityTypes returns a boolean if a field has been set.
func (o *TagDef) HasEntityTypes() bool {
	if o != nil && o.EntityTypes != nil {
		return true
	}

	return false
}

// SetEntityTypes gets a reference to the given []string and assigns it to the EntityTypes field.
func (o *TagDef) SetEntityTypes(v []string) {
	o.EntityTypes = &v
}

// GetSubTypes returns the SubTypes field value if set, zero value otherwise.
func (o *TagDef) GetSubTypes() []string {
	if o == nil || o.SubTypes == nil {
		var ret []string
		return ret
	}
	return *o.SubTypes
}

// GetSubTypesOk returns a tuple with the SubTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDef) GetSubTypesOk() (*[]string, bool) {
	if o == nil || o.SubTypes == nil {
		return nil, false
	}
	return o.SubTypes, true
}

// HasSubTypes returns a boolean if a field has been set.
func (o *TagDef) HasSubTypes() bool {
	if o != nil && o.SubTypes != nil {
		return true
	}

	return false
}

// SetSubTypes gets a reference to the given []string and assigns it to the SubTypes field.
func (o *TagDef) SetSubTypes(v []string) {
	o.SubTypes = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *TagDef) Redact() {
    o.recurseRedact(o.Category)
    o.recurseRedact(o.Guid)
    o.recurseRedact(o.CreatedBy)
    o.recurseRedact(o.UpdatedBy)
    o.recurseRedact(o.CreateTime)
    o.recurseRedact(o.UpdateTime)
    o.recurseRedact(o.Version)
    o.recurseRedact(o.Name)
    o.recurseRedact(o.Description)
    o.recurseRedact(o.TypeVersion)
    o.recurseRedact(o.ServiceType)
    o.recurseRedact(o.Options)
    o.recurseRedact(o.AttributeDefs)
    o.recurseRedact(o.SuperTypes)
    o.recurseRedact(o.EntityTypes)
    o.recurseRedact(o.SubTypes)
}

func (o *TagDef) recurseRedact(v interface{}) {
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

func (o TagDef) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o TagDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TypeVersion != nil {
		toSerialize["typeVersion"] = o.TypeVersion
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.AttributeDefs != nil {
		toSerialize["attributeDefs"] = o.AttributeDefs
	}
	if o.SuperTypes != nil {
		toSerialize["superTypes"] = o.SuperTypes
	}
	if o.EntityTypes != nil {
		toSerialize["entityTypes"] = o.EntityTypes
	}
	if o.SubTypes != nil {
		toSerialize["subTypes"] = o.SubTypes
	}
	return json.Marshal(toSerialize)
}

type NullableTagDef struct {
	value *TagDef
	isSet bool
}

func (v NullableTagDef) Get() *TagDef {
	return v.value
}

func (v *NullableTagDef) Set(val *TagDef) {
	v.value = val
	v.isSet = true
}

func (v NullableTagDef) IsSet() bool {
	return v.isSet
}

func (v *NullableTagDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagDef(val *TagDef) *NullableTagDef {
	return &NullableTagDef{value: val, isSet: true}
}

func (v NullableTagDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


