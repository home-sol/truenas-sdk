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

// checks if the FilesystemPut1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesystemPut1{}

// FilesystemPut1 struct for FilesystemPut1
type FilesystemPut1 struct {
	Append *bool  `json:"append,omitempty"`
	Mode   *int32 `json:"mode,omitempty"`
}

// NewFilesystemPut1 instantiates a new FilesystemPut1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemPut1() *FilesystemPut1 {
	this := FilesystemPut1{}
	var append bool
	this.Append = &append
	return &this
}

// NewFilesystemPut1WithDefaults instantiates a new FilesystemPut1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemPut1WithDefaults() *FilesystemPut1 {
	this := FilesystemPut1{}
	var append bool
	this.Append = &append
	return &this
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *FilesystemPut1) GetAppend() bool {
	if o == nil || IsNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemPut1) GetAppendOk() (*bool, bool) {
	if o == nil || IsNil(o.Append) {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *FilesystemPut1) HasAppend() bool {
	if o != nil && !IsNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *FilesystemPut1) SetAppend(v bool) {
	o.Append = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *FilesystemPut1) GetMode() int32 {
	if o == nil || IsNil(o.Mode) {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemPut1) GetModeOk() (*int32, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *FilesystemPut1) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *FilesystemPut1) SetMode(v int32) {
	o.Mode = &v
}

func (o FilesystemPut1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesystemPut1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Append) {
		toSerialize["append"] = o.Append
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableFilesystemPut1 struct {
	value *FilesystemPut1
	isSet bool
}

func (v NullableFilesystemPut1) Get() *FilesystemPut1 {
	return v.value
}

func (v *NullableFilesystemPut1) Set(val *FilesystemPut1) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemPut1) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemPut1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemPut1(val *FilesystemPut1) *NullableFilesystemPut1 {
	return &NullableFilesystemPut1{value: val, isSet: true}
}

func (v NullableFilesystemPut1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemPut1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}