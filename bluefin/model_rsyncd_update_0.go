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

// checks if the RsyncdUpdate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RsyncdUpdate0{}

// RsyncdUpdate0 struct for RsyncdUpdate0
type RsyncdUpdate0 struct {
	Port      *int32  `json:"port,omitempty"`
	Auxiliary *string `json:"auxiliary,omitempty"`
}

// NewRsyncdUpdate0 instantiates a new RsyncdUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRsyncdUpdate0() *RsyncdUpdate0 {
	this := RsyncdUpdate0{}
	return &this
}

// NewRsyncdUpdate0WithDefaults instantiates a new RsyncdUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRsyncdUpdate0WithDefaults() *RsyncdUpdate0 {
	this := RsyncdUpdate0{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RsyncdUpdate0) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsyncdUpdate0) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RsyncdUpdate0) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RsyncdUpdate0) SetPort(v int32) {
	o.Port = &v
}

// GetAuxiliary returns the Auxiliary field value if set, zero value otherwise.
func (o *RsyncdUpdate0) GetAuxiliary() string {
	if o == nil || IsNil(o.Auxiliary) {
		var ret string
		return ret
	}
	return *o.Auxiliary
}

// GetAuxiliaryOk returns a tuple with the Auxiliary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsyncdUpdate0) GetAuxiliaryOk() (*string, bool) {
	if o == nil || IsNil(o.Auxiliary) {
		return nil, false
	}
	return o.Auxiliary, true
}

// HasAuxiliary returns a boolean if a field has been set.
func (o *RsyncdUpdate0) HasAuxiliary() bool {
	if o != nil && !IsNil(o.Auxiliary) {
		return true
	}

	return false
}

// SetAuxiliary gets a reference to the given string and assigns it to the Auxiliary field.
func (o *RsyncdUpdate0) SetAuxiliary(v string) {
	o.Auxiliary = &v
}

func (o RsyncdUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RsyncdUpdate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Auxiliary) {
		toSerialize["auxiliary"] = o.Auxiliary
	}
	return toSerialize, nil
}

type NullableRsyncdUpdate0 struct {
	value *RsyncdUpdate0
	isSet bool
}

func (v NullableRsyncdUpdate0) Get() *RsyncdUpdate0 {
	return v.value
}

func (v *NullableRsyncdUpdate0) Set(val *RsyncdUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableRsyncdUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableRsyncdUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRsyncdUpdate0(val *RsyncdUpdate0) *NullableRsyncdUpdate0 {
	return &NullableRsyncdUpdate0{value: val, isSet: true}
}

func (v NullableRsyncdUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRsyncdUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}