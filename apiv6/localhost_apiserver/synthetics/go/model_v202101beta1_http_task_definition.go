/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

type V202101beta1HttpTaskDefinition struct {
	Target string `json:"target,omitempty"`

	Period int64 `json:"period,omitempty"`

	Expiry int64 `json:"expiry,omitempty"`
}