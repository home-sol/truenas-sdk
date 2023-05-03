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

// IscsiInitiatorGetInstance0 Returns instance matching `id`. If `id` is not found, Validation error is raised.
type IscsiInitiatorGetInstance0 struct {
	int32  *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IscsiInitiatorGetInstance0) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(IscsiInitiatorGetInstance0)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IscsiInitiatorGetInstance0) MarshalJSON() ([]byte, error) {
	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIscsiInitiatorGetInstance0 struct {
	value *IscsiInitiatorGetInstance0
	isSet bool
}

func (v NullableIscsiInitiatorGetInstance0) Get() *IscsiInitiatorGetInstance0 {
	return v.value
}

func (v *NullableIscsiInitiatorGetInstance0) Set(val *IscsiInitiatorGetInstance0) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiInitiatorGetInstance0) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiInitiatorGetInstance0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiInitiatorGetInstance0(val *IscsiInitiatorGetInstance0) *NullableIscsiInitiatorGetInstance0 {
	return &NullableIscsiInitiatorGetInstance0{value: val, isSet: true}
}

func (v NullableIscsiInitiatorGetInstance0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiInitiatorGetInstance0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
