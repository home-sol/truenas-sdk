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

// checks if the SnapshotSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotSpec{}

// SnapshotSpec struct for SnapshotSpec
type SnapshotSpec struct {
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

// NewSnapshotSpec instantiates a new SnapshotSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotSpec() *SnapshotSpec {
	this := SnapshotSpec{}
	return &this
}

// NewSnapshotSpecWithDefaults instantiates a new SnapshotSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotSpecWithDefaults() *SnapshotSpec {
	this := SnapshotSpec{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SnapshotSpec) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSpec) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SnapshotSpec) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *SnapshotSpec) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SnapshotSpec) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSpec) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SnapshotSpec) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *SnapshotSpec) SetEnd(v string) {
	o.End = &v
}

func (o SnapshotSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableSnapshotSpec struct {
	value *SnapshotSpec
	isSet bool
}

func (v NullableSnapshotSpec) Get() *SnapshotSpec {
	return v.value
}

func (v *NullableSnapshotSpec) Set(val *SnapshotSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotSpec(val *SnapshotSpec) *NullableSnapshotSpec {
	return &NullableSnapshotSpec{value: val, isSet: true}
}

func (v NullableSnapshotSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
