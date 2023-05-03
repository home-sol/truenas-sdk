/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"encoding/json"
)

// checks if the ChartReleasePodLogs1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChartReleasePodLogs1{}

// ChartReleasePodLogs1 struct for ChartReleasePodLogs1
type ChartReleasePodLogs1 struct {
	LimitBytes    NullableInt32 `json:"limit_bytes,omitempty"`
	TailLines     NullableInt32 `json:"tail_lines,omitempty"`
	PodName       *string       `json:"pod_name,omitempty"`
	ContainerName *string       `json:"container_name,omitempty"`
}

// NewChartReleasePodLogs1 instantiates a new ChartReleasePodLogs1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartReleasePodLogs1() *ChartReleasePodLogs1 {
	this := ChartReleasePodLogs1{}
	var tailLines int32 = 500
	this.TailLines = *NewNullableInt32(&tailLines)
	return &this
}

// NewChartReleasePodLogs1WithDefaults instantiates a new ChartReleasePodLogs1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartReleasePodLogs1WithDefaults() *ChartReleasePodLogs1 {
	this := ChartReleasePodLogs1{}
	var tailLines int32 = 500
	this.TailLines = *NewNullableInt32(&tailLines)
	return &this
}

// GetLimitBytes returns the LimitBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChartReleasePodLogs1) GetLimitBytes() int32 {
	if o == nil || IsNil(o.LimitBytes.Get()) {
		var ret int32
		return ret
	}
	return *o.LimitBytes.Get()
}

// GetLimitBytesOk returns a tuple with the LimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChartReleasePodLogs1) GetLimitBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitBytes.Get(), o.LimitBytes.IsSet()
}

// HasLimitBytes returns a boolean if a field has been set.
func (o *ChartReleasePodLogs1) HasLimitBytes() bool {
	if o != nil && o.LimitBytes.IsSet() {
		return true
	}

	return false
}

// SetLimitBytes gets a reference to the given NullableInt32 and assigns it to the LimitBytes field.
func (o *ChartReleasePodLogs1) SetLimitBytes(v int32) {
	o.LimitBytes.Set(&v)
}

// SetLimitBytesNil sets the value for LimitBytes to be an explicit nil
func (o *ChartReleasePodLogs1) SetLimitBytesNil() {
	o.LimitBytes.Set(nil)
}

// UnsetLimitBytes ensures that no value is present for LimitBytes, not even an explicit nil
func (o *ChartReleasePodLogs1) UnsetLimitBytes() {
	o.LimitBytes.Unset()
}

// GetTailLines returns the TailLines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChartReleasePodLogs1) GetTailLines() int32 {
	if o == nil || IsNil(o.TailLines.Get()) {
		var ret int32
		return ret
	}
	return *o.TailLines.Get()
}

// GetTailLinesOk returns a tuple with the TailLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChartReleasePodLogs1) GetTailLinesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TailLines.Get(), o.TailLines.IsSet()
}

// HasTailLines returns a boolean if a field has been set.
func (o *ChartReleasePodLogs1) HasTailLines() bool {
	if o != nil && o.TailLines.IsSet() {
		return true
	}

	return false
}

// SetTailLines gets a reference to the given NullableInt32 and assigns it to the TailLines field.
func (o *ChartReleasePodLogs1) SetTailLines(v int32) {
	o.TailLines.Set(&v)
}

// SetTailLinesNil sets the value for TailLines to be an explicit nil
func (o *ChartReleasePodLogs1) SetTailLinesNil() {
	o.TailLines.Set(nil)
}

// UnsetTailLines ensures that no value is present for TailLines, not even an explicit nil
func (o *ChartReleasePodLogs1) UnsetTailLines() {
	o.TailLines.Unset()
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *ChartReleasePodLogs1) GetPodName() string {
	if o == nil || IsNil(o.PodName) {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartReleasePodLogs1) GetPodNameOk() (*string, bool) {
	if o == nil || IsNil(o.PodName) {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *ChartReleasePodLogs1) HasPodName() bool {
	if o != nil && !IsNil(o.PodName) {
		return true
	}

	return false
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *ChartReleasePodLogs1) SetPodName(v string) {
	o.PodName = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *ChartReleasePodLogs1) GetContainerName() string {
	if o == nil || IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartReleasePodLogs1) GetContainerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *ChartReleasePodLogs1) HasContainerName() bool {
	if o != nil && !IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *ChartReleasePodLogs1) SetContainerName(v string) {
	o.ContainerName = &v
}

func (o ChartReleasePodLogs1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChartReleasePodLogs1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LimitBytes.IsSet() {
		toSerialize["limit_bytes"] = o.LimitBytes.Get()
	}
	if o.TailLines.IsSet() {
		toSerialize["tail_lines"] = o.TailLines.Get()
	}
	if !IsNil(o.PodName) {
		toSerialize["pod_name"] = o.PodName
	}
	if !IsNil(o.ContainerName) {
		toSerialize["container_name"] = o.ContainerName
	}
	return toSerialize, nil
}

type NullableChartReleasePodLogs1 struct {
	value *ChartReleasePodLogs1
	isSet bool
}

func (v NullableChartReleasePodLogs1) Get() *ChartReleasePodLogs1 {
	return v.value
}

func (v *NullableChartReleasePodLogs1) Set(val *ChartReleasePodLogs1) {
	v.value = val
	v.isSet = true
}

func (v NullableChartReleasePodLogs1) IsSet() bool {
	return v.isSet
}

func (v *NullableChartReleasePodLogs1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartReleasePodLogs1(val *ChartReleasePodLogs1) *NullableChartReleasePodLogs1 {
	return &NullableChartReleasePodLogs1{value: val, isSet: true}
}

func (v NullableChartReleasePodLogs1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartReleasePodLogs1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
