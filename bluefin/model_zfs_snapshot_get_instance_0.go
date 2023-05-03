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

// ZfsSnapshotGetInstance0 Returns instance matching `id`. If `id` is not found, Validation error is raised.
type ZfsSnapshotGetInstance0 struct {
	int32  *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZfsSnapshotGetInstance0) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ZfsSnapshotGetInstance0)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZfsSnapshotGetInstance0) MarshalJSON() ([]byte, error) {
	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableZfsSnapshotGetInstance0 struct {
	value *ZfsSnapshotGetInstance0
	isSet bool
}

func (v NullableZfsSnapshotGetInstance0) Get() *ZfsSnapshotGetInstance0 {
	return v.value
}

func (v *NullableZfsSnapshotGetInstance0) Set(val *ZfsSnapshotGetInstance0) {
	v.value = val
	v.isSet = true
}

func (v NullableZfsSnapshotGetInstance0) IsSet() bool {
	return v.isSet
}

func (v *NullableZfsSnapshotGetInstance0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZfsSnapshotGetInstance0(val *ZfsSnapshotGetInstance0) *NullableZfsSnapshotGetInstance0 {
	return &NullableZfsSnapshotGetInstance0{value: val, isSet: true}
}

func (v NullableZfsSnapshotGetInstance0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZfsSnapshotGetInstance0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
