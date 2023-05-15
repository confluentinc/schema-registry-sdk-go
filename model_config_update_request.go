/*
 * Confluent Schema Registry
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schemaregistry
// ConfigUpdateRequest Config update request
type ConfigUpdateRequest struct {
	// Compatibility Level
	Compatibility string `json:"compatibility,omitempty"`
	CompatibilityGroup string `json:"compatibilityGroup,omitempty"`
	DefaultMetadata *Metadata `json:"defaultMetadata,omitempty"`
	OverrideMetadata *Metadata `json:"overrideMetadata,omitempty"`
	DefaultRuleSet *RuleSet `json:"defaultRuleSet,omitempty"`
	OverrideRuleSet *RuleSet `json:"overrideRuleSet,omitempty"`
}
