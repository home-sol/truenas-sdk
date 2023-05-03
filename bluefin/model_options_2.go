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

// checks if the Options2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Options2{}

// Options2 If `mode` is specified then the mode will be applied to the path and files and subdirectories depending on which `options` are selected. Mode should be formatted as string representation of octal permissions bits.
type Options2 struct {
	Stripacl  *bool `json:"stripacl,omitempty"`
	Recursive *bool `json:"recursive,omitempty"`
	Traverse  *bool `json:"traverse,omitempty"`
}

// NewOptions2 instantiates a new Options2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptions2() *Options2 {
	this := Options2{}
	var stripacl bool
	this.Stripacl = &stripacl
	var recursive bool
	this.Recursive = &recursive
	var traverse bool
	this.Traverse = &traverse
	return &this
}

// NewOptions2WithDefaults instantiates a new Options2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptions2WithDefaults() *Options2 {
	this := Options2{}
	var stripacl bool
	this.Stripacl = &stripacl
	var recursive bool
	this.Recursive = &recursive
	var traverse bool
	this.Traverse = &traverse
	return &this
}

// GetStripacl returns the Stripacl field value if set, zero value otherwise.
func (o *Options2) GetStripacl() bool {
	if o == nil || IsNil(o.Stripacl) {
		var ret bool
		return ret
	}
	return *o.Stripacl
}

// GetStripaclOk returns a tuple with the Stripacl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options2) GetStripaclOk() (*bool, bool) {
	if o == nil || IsNil(o.Stripacl) {
		return nil, false
	}
	return o.Stripacl, true
}

// HasStripacl returns a boolean if a field has been set.
func (o *Options2) HasStripacl() bool {
	if o != nil && !IsNil(o.Stripacl) {
		return true
	}

	return false
}

// SetStripacl gets a reference to the given bool and assigns it to the Stripacl field.
func (o *Options2) SetStripacl(v bool) {
	o.Stripacl = &v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *Options2) GetRecursive() bool {
	if o == nil || IsNil(o.Recursive) {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options2) GetRecursiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Recursive) {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *Options2) HasRecursive() bool {
	if o != nil && !IsNil(o.Recursive) {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *Options2) SetRecursive(v bool) {
	o.Recursive = &v
}

// GetTraverse returns the Traverse field value if set, zero value otherwise.
func (o *Options2) GetTraverse() bool {
	if o == nil || IsNil(o.Traverse) {
		var ret bool
		return ret
	}
	return *o.Traverse
}

// GetTraverseOk returns a tuple with the Traverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options2) GetTraverseOk() (*bool, bool) {
	if o == nil || IsNil(o.Traverse) {
		return nil, false
	}
	return o.Traverse, true
}

// HasTraverse returns a boolean if a field has been set.
func (o *Options2) HasTraverse() bool {
	if o != nil && !IsNil(o.Traverse) {
		return true
	}

	return false
}

// SetTraverse gets a reference to the given bool and assigns it to the Traverse field.
func (o *Options2) SetTraverse(v bool) {
	o.Traverse = &v
}

func (o Options2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Options2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stripacl) {
		toSerialize["stripacl"] = o.Stripacl
	}
	if !IsNil(o.Recursive) {
		toSerialize["recursive"] = o.Recursive
	}
	if !IsNil(o.Traverse) {
		toSerialize["traverse"] = o.Traverse
	}
	return toSerialize, nil
}

type NullableOptions2 struct {
	value *Options2
	isSet bool
}

func (v NullableOptions2) Get() *Options2 {
	return v.value
}

func (v *NullableOptions2) Set(val *Options2) {
	v.value = val
	v.isSet = true
}

func (v NullableOptions2) IsSet() bool {
	return v.isSet
}

func (v *NullableOptions2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptions2(val *Options2) *NullableOptions2 {
	return &NullableOptions2{value: val, isSet: true}
}

func (v NullableOptions2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptions2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
