/*
 * Confluent Schema Registry
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schemaregistry
// AtlasAttributeDef struct for AtlasAttributeDef
type AtlasAttributeDef struct {
	Name string `json:"name,omitempty"`
	TypeName string `json:"typeName,omitempty"`
	IsOptional bool `json:"isOptional,omitempty"`
	Cardinality string `json:"cardinality,omitempty"`
	ValuesMinCount int32 `json:"valuesMinCount,omitempty"`
	ValuesMaxCount int32 `json:"valuesMaxCount,omitempty"`
	IsUnique bool `json:"isUnique,omitempty"`
	IsIndexable bool `json:"isIndexable,omitempty"`
	IncludeInNotification bool `json:"includeInNotification,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Description string `json:"description,omitempty"`
	SearchWeight int32 `json:"searchWeight,omitempty"`
	IndexType string `json:"indexType,omitempty"`
	Constraints []AtlasConstraintDef `json:"constraints,omitempty"`
	Options map[string]string `json:"options,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
