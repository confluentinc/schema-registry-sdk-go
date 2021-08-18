/*
 * Confluent Schema Registry
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schemaregistry

type ExporterInfo struct {
	Name string `json:"name,omitempty"`
	Subjects []string `json:"subjects,omitempty"`
	ContextType string `json:"contextType,omitempty"`
	Context string `json:"context,omitempty"`
	Config map[string]string `json:"config,omitempty"`
}
