/*
 * Confluent Schema Registry
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schemaregistry
// AtlasEntityWithExtInfo struct for AtlasEntityWithExtInfo
type AtlasEntityWithExtInfo struct {
	ReferredEntities map[string]AtlasEntity `json:"referredEntities,omitempty"`
	Entity AtlasEntity `json:"entity,omitempty"`
}
