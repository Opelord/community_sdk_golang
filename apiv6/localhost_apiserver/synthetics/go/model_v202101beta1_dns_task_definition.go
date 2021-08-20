/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

type V202101beta1DnsTaskDefinition struct {
	Target string `json:"target,omitempty"`

	Period int64 `json:"period,omitempty"`

	Expiry int64 `json:"expiry,omitempty"`

	Count int64 `json:"count,omitempty"`

	Port int64 `json:"port,omitempty"`

	Type string `json:"type,omitempty"`

	Resolver string `json:"resolver,omitempty"`
}