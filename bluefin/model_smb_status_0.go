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

// SmbStatus0 `info_level` type of information requests. Defaults to ALL.
type SmbStatus0 string

// List of smb_status_0
const (
	AUTH_LOG      SmbStatus0 = "AUTH_LOG"
	ALL           SmbStatus0 = "ALL"
	SESSIONS      SmbStatus0 = "SESSIONS"
	SHARES        SmbStatus0 = "SHARES"
	LOCKS         SmbStatus0 = "LOCKS"
	BYTERANGE     SmbStatus0 = "BYTERANGE"
	NOTIFICATIONS SmbStatus0 = "NOTIFICATIONS"
)

// All allowed values of SmbStatus0 enum
var AllowedSmbStatus0EnumValues = []SmbStatus0{
	"AUTH_LOG",
	"ALL",
	"SESSIONS",
	"SHARES",
	"LOCKS",
	"BYTERANGE",
	"NOTIFICATIONS",
}

func (v *SmbStatus0) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmbStatus0(value)
	for _, existing := range AllowedSmbStatus0EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SmbStatus0", value)
}

// NewSmbStatus0FromValue returns a pointer to a valid SmbStatus0
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSmbStatus0FromValue(v string) (*SmbStatus0, error) {
	ev := SmbStatus0(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SmbStatus0: valid values are %v", v, AllowedSmbStatus0EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SmbStatus0) IsValid() bool {
	for _, existing := range AllowedSmbStatus0EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to smb_status_0 value
func (v SmbStatus0) Ptr() *SmbStatus0 {
	return &v
}

type NullableSmbStatus0 struct {
	value *SmbStatus0
	isSet bool
}

func (v NullableSmbStatus0) Get() *SmbStatus0 {
	return v.value
}

func (v *NullableSmbStatus0) Set(val *SmbStatus0) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbStatus0) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbStatus0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbStatus0(val *SmbStatus0) *NullableSmbStatus0 {
	return &NullableSmbStatus0{value: val, isSet: true}
}

func (v NullableSmbStatus0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbStatus0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}