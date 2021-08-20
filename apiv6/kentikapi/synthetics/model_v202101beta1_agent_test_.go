/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synthetics

import (
	"encoding/json"
)

// V202101beta1AgentTest struct for V202101beta1AgentTest
type V202101beta1AgentTest struct {
	Target *string `json:"target,omitempty"`
}

// NewV202101beta1AgentTest instantiates a new V202101beta1AgentTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1AgentTest() *V202101beta1AgentTest {
	this := V202101beta1AgentTest{}
	return &this
}

// NewV202101beta1AgentTestWithDefaults instantiates a new V202101beta1AgentTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1AgentTestWithDefaults() *V202101beta1AgentTest {
	this := V202101beta1AgentTest{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *V202101beta1AgentTest) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1AgentTest) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *V202101beta1AgentTest) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *V202101beta1AgentTest) SetTarget(v string) {
	o.Target = &v
}

func (o V202101beta1AgentTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1AgentTest struct {
	value *V202101beta1AgentTest
	isSet bool
}

func (v NullableV202101beta1AgentTest) Get() *V202101beta1AgentTest {
	return v.value
}

func (v *NullableV202101beta1AgentTest) Set(val *V202101beta1AgentTest) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1AgentTest) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1AgentTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1AgentTest(val *V202101beta1AgentTest) *NullableV202101beta1AgentTest {
	return &NullableV202101beta1AgentTest{value: val, isSet: true}
}

func (v NullableV202101beta1AgentTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1AgentTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}