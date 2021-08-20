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
	"time"
)

// V202101beta1MeshMetrics struct for V202101beta1MeshMetrics
type V202101beta1MeshMetrics struct {
	Time       *time.Time              `json:"time,omitempty"`
	Latency    *V202101beta1MeshMetric `json:"latency,omitempty"`
	PacketLoss *V202101beta1MeshMetric `json:"packetLoss,omitempty"`
	Jitter     *V202101beta1MeshMetric `json:"jitter,omitempty"`
}

// NewV202101beta1MeshMetrics instantiates a new V202101beta1MeshMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1MeshMetrics() *V202101beta1MeshMetrics {
	this := V202101beta1MeshMetrics{}
	return &this
}

// NewV202101beta1MeshMetricsWithDefaults instantiates a new V202101beta1MeshMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1MeshMetricsWithDefaults() *V202101beta1MeshMetrics {
	this := V202101beta1MeshMetrics{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V202101beta1MeshMetrics) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1MeshMetrics) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V202101beta1MeshMetrics) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *V202101beta1MeshMetrics) SetTime(v time.Time) {
	o.Time = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *V202101beta1MeshMetrics) GetLatency() V202101beta1MeshMetric {
	if o == nil || o.Latency == nil {
		var ret V202101beta1MeshMetric
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1MeshMetrics) GetLatencyOk() (*V202101beta1MeshMetric, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *V202101beta1MeshMetrics) HasLatency() bool {
	if o != nil && o.Latency != nil {
		return true
	}

	return false
}

// SetLatency gets a reference to the given V202101beta1MeshMetric and assigns it to the Latency field.
func (o *V202101beta1MeshMetrics) SetLatency(v V202101beta1MeshMetric) {
	o.Latency = &v
}

// GetPacketLoss returns the PacketLoss field value if set, zero value otherwise.
func (o *V202101beta1MeshMetrics) GetPacketLoss() V202101beta1MeshMetric {
	if o == nil || o.PacketLoss == nil {
		var ret V202101beta1MeshMetric
		return ret
	}
	return *o.PacketLoss
}

// GetPacketLossOk returns a tuple with the PacketLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1MeshMetrics) GetPacketLossOk() (*V202101beta1MeshMetric, bool) {
	if o == nil || o.PacketLoss == nil {
		return nil, false
	}
	return o.PacketLoss, true
}

// HasPacketLoss returns a boolean if a field has been set.
func (o *V202101beta1MeshMetrics) HasPacketLoss() bool {
	if o != nil && o.PacketLoss != nil {
		return true
	}

	return false
}

// SetPacketLoss gets a reference to the given V202101beta1MeshMetric and assigns it to the PacketLoss field.
func (o *V202101beta1MeshMetrics) SetPacketLoss(v V202101beta1MeshMetric) {
	o.PacketLoss = &v
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *V202101beta1MeshMetrics) GetJitter() V202101beta1MeshMetric {
	if o == nil || o.Jitter == nil {
		var ret V202101beta1MeshMetric
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1MeshMetrics) GetJitterOk() (*V202101beta1MeshMetric, bool) {
	if o == nil || o.Jitter == nil {
		return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *V202101beta1MeshMetrics) HasJitter() bool {
	if o != nil && o.Jitter != nil {
		return true
	}

	return false
}

// SetJitter gets a reference to the given V202101beta1MeshMetric and assigns it to the Jitter field.
func (o *V202101beta1MeshMetrics) SetJitter(v V202101beta1MeshMetric) {
	o.Jitter = &v
}

func (o V202101beta1MeshMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Latency != nil {
		toSerialize["latency"] = o.Latency
	}
	if o.PacketLoss != nil {
		toSerialize["packetLoss"] = o.PacketLoss
	}
	if o.Jitter != nil {
		toSerialize["jitter"] = o.Jitter
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1MeshMetrics struct {
	value *V202101beta1MeshMetrics
	isSet bool
}

func (v NullableV202101beta1MeshMetrics) Get() *V202101beta1MeshMetrics {
	return v.value
}

func (v *NullableV202101beta1MeshMetrics) Set(val *V202101beta1MeshMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1MeshMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1MeshMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1MeshMetrics(val *V202101beta1MeshMetrics) *NullableV202101beta1MeshMetrics {
	return &NullableV202101beta1MeshMetrics{value: val, isSet: true}
}

func (v NullableV202101beta1MeshMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1MeshMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}