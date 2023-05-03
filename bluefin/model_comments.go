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

// Comments struct for Comments
type Comments struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Comments) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(Comments)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Comments) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableComments struct {
	value *Comments
	isSet bool
}

func (v NullableComments) Get() *Comments {
	return v.value
}

func (v *NullableComments) Set(val *Comments) {
	v.value = val
	v.isSet = true
}

func (v NullableComments) IsSet() bool {
	return v.isSet
}

func (v *NullableComments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComments(val *Comments) *NullableComments {
	return &NullableComments{value: val, isSet: true}
}

func (v NullableComments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}