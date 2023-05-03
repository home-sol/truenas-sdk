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

// checks if the GlusterPeerStatus0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlusterPeerStatus0{}

// GlusterPeerStatus0 struct for GlusterPeerStatus0
type GlusterPeerStatus0 struct {
	Localhost *bool `json:"localhost,omitempty"`
}

// NewGlusterPeerStatus0 instantiates a new GlusterPeerStatus0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterPeerStatus0() *GlusterPeerStatus0 {
	this := GlusterPeerStatus0{}
	var localhost bool
	this.Localhost = &localhost
	return &this
}

// NewGlusterPeerStatus0WithDefaults instantiates a new GlusterPeerStatus0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterPeerStatus0WithDefaults() *GlusterPeerStatus0 {
	this := GlusterPeerStatus0{}
	var localhost bool
	this.Localhost = &localhost
	return &this
}

// GetLocalhost returns the Localhost field value if set, zero value otherwise.
func (o *GlusterPeerStatus0) GetLocalhost() bool {
	if o == nil || IsNil(o.Localhost) {
		var ret bool
		return ret
	}
	return *o.Localhost
}

// GetLocalhostOk returns a tuple with the Localhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterPeerStatus0) GetLocalhostOk() (*bool, bool) {
	if o == nil || IsNil(o.Localhost) {
		return nil, false
	}
	return o.Localhost, true
}

// HasLocalhost returns a boolean if a field has been set.
func (o *GlusterPeerStatus0) HasLocalhost() bool {
	if o != nil && !IsNil(o.Localhost) {
		return true
	}

	return false
}

// SetLocalhost gets a reference to the given bool and assigns it to the Localhost field.
func (o *GlusterPeerStatus0) SetLocalhost(v bool) {
	o.Localhost = &v
}

func (o GlusterPeerStatus0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlusterPeerStatus0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Localhost) {
		toSerialize["localhost"] = o.Localhost
	}
	return toSerialize, nil
}

type NullableGlusterPeerStatus0 struct {
	value *GlusterPeerStatus0
	isSet bool
}

func (v NullableGlusterPeerStatus0) Get() *GlusterPeerStatus0 {
	return v.value
}

func (v *NullableGlusterPeerStatus0) Set(val *GlusterPeerStatus0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterPeerStatus0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterPeerStatus0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterPeerStatus0(val *GlusterPeerStatus0) *NullableGlusterPeerStatus0 {
	return &NullableGlusterPeerStatus0{value: val, isSet: true}
}

func (v NullableGlusterPeerStatus0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterPeerStatus0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
