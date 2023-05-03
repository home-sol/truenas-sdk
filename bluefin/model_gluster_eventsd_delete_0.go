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

// checks if the GlusterEventsdDelete0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlusterEventsdDelete0{}

// GlusterEventsdDelete0 struct for GlusterEventsdDelete0
type GlusterEventsdDelete0 struct {
	// Delete `url` webhook
	Url *string `json:"url,omitempty"`
}

// NewGlusterEventsdDelete0 instantiates a new GlusterEventsdDelete0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterEventsdDelete0() *GlusterEventsdDelete0 {
	this := GlusterEventsdDelete0{}
	return &this
}

// NewGlusterEventsdDelete0WithDefaults instantiates a new GlusterEventsdDelete0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterEventsdDelete0WithDefaults() *GlusterEventsdDelete0 {
	this := GlusterEventsdDelete0{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GlusterEventsdDelete0) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterEventsdDelete0) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GlusterEventsdDelete0) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GlusterEventsdDelete0) SetUrl(v string) {
	o.Url = &v
}

func (o GlusterEventsdDelete0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlusterEventsdDelete0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGlusterEventsdDelete0 struct {
	value *GlusterEventsdDelete0
	isSet bool
}

func (v NullableGlusterEventsdDelete0) Get() *GlusterEventsdDelete0 {
	return v.value
}

func (v *NullableGlusterEventsdDelete0) Set(val *GlusterEventsdDelete0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterEventsdDelete0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterEventsdDelete0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterEventsdDelete0(val *GlusterEventsdDelete0) *NullableGlusterEventsdDelete0 {
	return &NullableGlusterEventsdDelete0{value: val, isSet: true}
}

func (v NullableGlusterEventsdDelete0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterEventsdDelete0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}