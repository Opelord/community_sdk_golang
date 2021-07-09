/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudexport

import (
	"encoding/json"
)

// ProtobufAny struct for ProtobufAny
type ProtobufAny struct {
	TypeUrl *string `json:"typeUrl,omitempty"`
	Value   *string `json:"value,omitempty"`
}

// NewProtobufAny instantiates a new ProtobufAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtobufAny() *ProtobufAny {
	this := ProtobufAny{}
	return &this
}

// NewProtobufAnyWithDefaults instantiates a new ProtobufAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtobufAnyWithDefaults() *ProtobufAny {
	this := ProtobufAny{}
	return &this
}

// GetTypeUrl returns the TypeUrl field value if set, zero value otherwise.
func (o *ProtobufAny) GetTypeUrl() string {
	if o == nil || o.TypeUrl == nil {
		var ret string
		return ret
	}
	return *o.TypeUrl
}

// GetTypeUrlOk returns a tuple with the TypeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtobufAny) GetTypeUrlOk() (*string, bool) {
	if o == nil || o.TypeUrl == nil {
		return nil, false
	}
	return o.TypeUrl, true
}

// HasTypeUrl returns a boolean if a field has been set.
func (o *ProtobufAny) HasTypeUrl() bool {
	if o != nil && o.TypeUrl != nil {
		return true
	}

	return false
}

// SetTypeUrl gets a reference to the given string and assigns it to the TypeUrl field.
func (o *ProtobufAny) SetTypeUrl(v string) {
	o.TypeUrl = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProtobufAny) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtobufAny) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProtobufAny) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProtobufAny) SetValue(v string) {
	o.Value = &v
}

func (o ProtobufAny) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeUrl != nil {
		toSerialize["typeUrl"] = o.TypeUrl
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableProtobufAny struct {
	value *ProtobufAny
	isSet bool
}

func (v NullableProtobufAny) Get() *ProtobufAny {
	return v.value
}

func (v *NullableProtobufAny) Set(val *ProtobufAny) {
	v.value = val
	v.isSet = true
}

func (v NullableProtobufAny) IsSet() bool {
	return v.isSet
}

func (v *NullableProtobufAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtobufAny(val *ProtobufAny) *NullableProtobufAny {
	return &NullableProtobufAny{value: val, isSet: true}
}

func (v NullableProtobufAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtobufAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
