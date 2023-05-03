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

// checks if the GlusterFuseIsMounted0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlusterFuseIsMounted0{}

// GlusterFuseIsMounted0 struct for GlusterFuseIsMounted0
type GlusterFuseIsMounted0 struct {
	Name *string `json:"name,omitempty"`
}

// NewGlusterFuseIsMounted0 instantiates a new GlusterFuseIsMounted0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterFuseIsMounted0() *GlusterFuseIsMounted0 {
	this := GlusterFuseIsMounted0{}
	return &this
}

// NewGlusterFuseIsMounted0WithDefaults instantiates a new GlusterFuseIsMounted0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterFuseIsMounted0WithDefaults() *GlusterFuseIsMounted0 {
	this := GlusterFuseIsMounted0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlusterFuseIsMounted0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterFuseIsMounted0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlusterFuseIsMounted0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlusterFuseIsMounted0) SetName(v string) {
	o.Name = &v
}

func (o GlusterFuseIsMounted0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlusterFuseIsMounted0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGlusterFuseIsMounted0 struct {
	value *GlusterFuseIsMounted0
	isSet bool
}

func (v NullableGlusterFuseIsMounted0) Get() *GlusterFuseIsMounted0 {
	return v.value
}

func (v *NullableGlusterFuseIsMounted0) Set(val *GlusterFuseIsMounted0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterFuseIsMounted0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterFuseIsMounted0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterFuseIsMounted0(val *GlusterFuseIsMounted0) *NullableGlusterFuseIsMounted0 {
	return &NullableGlusterFuseIsMounted0{value: val, isSet: true}
}

func (v NullableGlusterFuseIsMounted0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterFuseIsMounted0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
