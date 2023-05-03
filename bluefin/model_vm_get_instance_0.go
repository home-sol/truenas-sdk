/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"encoding/json"
	"fmt"
)

// VmGetInstance0 Returns instance matching `id`. If `id` is not found, Validation error is raised.
type VmGetInstance0 struct {
	int32  *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *VmGetInstance0) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32)
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(VmGetInstance0)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *VmGetInstance0) MarshalJSON() ([]byte, error) {
	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableVmGetInstance0 struct {
	value *VmGetInstance0
	isSet bool
}

func (v NullableVmGetInstance0) Get() *VmGetInstance0 {
	return v.value
}

func (v *NullableVmGetInstance0) Set(val *VmGetInstance0) {
	v.value = val
	v.isSet = true
}

func (v NullableVmGetInstance0) IsSet() bool {
	return v.isSet
}

func (v *NullableVmGetInstance0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmGetInstance0(val *VmGetInstance0) *NullableVmGetInstance0 {
	return &NullableVmGetInstance0{value: val, isSet: true}
}

func (v NullableVmGetInstance0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmGetInstance0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
