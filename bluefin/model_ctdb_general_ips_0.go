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

// checks if the CtdbGeneralIps0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CtdbGeneralIps0{}

// CtdbGeneralIps0 struct for CtdbGeneralIps0
type CtdbGeneralIps0 struct {
	AllNodes *bool `json:"all_nodes,omitempty"`
}

// NewCtdbGeneralIps0 instantiates a new CtdbGeneralIps0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCtdbGeneralIps0() *CtdbGeneralIps0 {
	this := CtdbGeneralIps0{}
	var allNodes bool
	this.AllNodes = &allNodes
	return &this
}

// NewCtdbGeneralIps0WithDefaults instantiates a new CtdbGeneralIps0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCtdbGeneralIps0WithDefaults() *CtdbGeneralIps0 {
	this := CtdbGeneralIps0{}
	var allNodes bool
	this.AllNodes = &allNodes
	return &this
}

// GetAllNodes returns the AllNodes field value if set, zero value otherwise.
func (o *CtdbGeneralIps0) GetAllNodes() bool {
	if o == nil || IsNil(o.AllNodes) {
		var ret bool
		return ret
	}
	return *o.AllNodes
}

// GetAllNodesOk returns a tuple with the AllNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CtdbGeneralIps0) GetAllNodesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllNodes) {
		return nil, false
	}
	return o.AllNodes, true
}

// HasAllNodes returns a boolean if a field has been set.
func (o *CtdbGeneralIps0) HasAllNodes() bool {
	if o != nil && !IsNil(o.AllNodes) {
		return true
	}

	return false
}

// SetAllNodes gets a reference to the given bool and assigns it to the AllNodes field.
func (o *CtdbGeneralIps0) SetAllNodes(v bool) {
	o.AllNodes = &v
}

func (o CtdbGeneralIps0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CtdbGeneralIps0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllNodes) {
		toSerialize["all_nodes"] = o.AllNodes
	}
	return toSerialize, nil
}

type NullableCtdbGeneralIps0 struct {
	value *CtdbGeneralIps0
	isSet bool
}

func (v NullableCtdbGeneralIps0) Get() *CtdbGeneralIps0 {
	return v.value
}

func (v *NullableCtdbGeneralIps0) Set(val *CtdbGeneralIps0) {
	v.value = val
	v.isSet = true
}

func (v NullableCtdbGeneralIps0) IsSet() bool {
	return v.isSet
}

func (v *NullableCtdbGeneralIps0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCtdbGeneralIps0(val *CtdbGeneralIps0) *NullableCtdbGeneralIps0 {
	return &NullableCtdbGeneralIps0{value: val, isSet: true}
}

func (v NullableCtdbGeneralIps0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCtdbGeneralIps0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}