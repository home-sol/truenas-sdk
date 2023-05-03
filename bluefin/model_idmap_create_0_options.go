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

// IdmapCreate0Options `options` are additional parameters that are backend-dependent:
type IdmapCreate0Options struct {
	IdmapAdOptions      *IdmapAdOptions
	IdmapAutoridOptions *IdmapAutoridOptions
	IdmapLdapOptions    *IdmapLdapOptions
	IdmapNssOptions     *IdmapNssOptions
	IdmapRfc2307Options *IdmapRfc2307Options
	IdmapRidOptions     *IdmapRidOptions
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IdmapCreate0Options) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into IdmapAdOptions
	err = json.Unmarshal(data, &dst.IdmapAdOptions)
	if err == nil {
		jsonIdmapAdOptions, _ := json.Marshal(dst.IdmapAdOptions)
		if string(jsonIdmapAdOptions) == "{}" { // empty struct
			dst.IdmapAdOptions = nil
		} else {
			return nil // data stored in dst.IdmapAdOptions, return on the first match
		}
	} else {
		dst.IdmapAdOptions = nil
	}

	// try to unmarshal JSON data into IdmapAutoridOptions
	err = json.Unmarshal(data, &dst.IdmapAutoridOptions)
	if err == nil {
		jsonIdmapAutoridOptions, _ := json.Marshal(dst.IdmapAutoridOptions)
		if string(jsonIdmapAutoridOptions) == "{}" { // empty struct
			dst.IdmapAutoridOptions = nil
		} else {
			return nil // data stored in dst.IdmapAutoridOptions, return on the first match
		}
	} else {
		dst.IdmapAutoridOptions = nil
	}

	// try to unmarshal JSON data into IdmapLdapOptions
	err = json.Unmarshal(data, &dst.IdmapLdapOptions)
	if err == nil {
		jsonIdmapLdapOptions, _ := json.Marshal(dst.IdmapLdapOptions)
		if string(jsonIdmapLdapOptions) == "{}" { // empty struct
			dst.IdmapLdapOptions = nil
		} else {
			return nil // data stored in dst.IdmapLdapOptions, return on the first match
		}
	} else {
		dst.IdmapLdapOptions = nil
	}

	// try to unmarshal JSON data into IdmapNssOptions
	err = json.Unmarshal(data, &dst.IdmapNssOptions)
	if err == nil {
		jsonIdmapNssOptions, _ := json.Marshal(dst.IdmapNssOptions)
		if string(jsonIdmapNssOptions) == "{}" { // empty struct
			dst.IdmapNssOptions = nil
		} else {
			return nil // data stored in dst.IdmapNssOptions, return on the first match
		}
	} else {
		dst.IdmapNssOptions = nil
	}

	// try to unmarshal JSON data into IdmapRfc2307Options
	err = json.Unmarshal(data, &dst.IdmapRfc2307Options)
	if err == nil {
		jsonIdmapRfc2307Options, _ := json.Marshal(dst.IdmapRfc2307Options)
		if string(jsonIdmapRfc2307Options) == "{}" { // empty struct
			dst.IdmapRfc2307Options = nil
		} else {
			return nil // data stored in dst.IdmapRfc2307Options, return on the first match
		}
	} else {
		dst.IdmapRfc2307Options = nil
	}

	// try to unmarshal JSON data into IdmapRidOptions
	err = json.Unmarshal(data, &dst.IdmapRidOptions)
	if err == nil {
		jsonIdmapRidOptions, _ := json.Marshal(dst.IdmapRidOptions)
		if string(jsonIdmapRidOptions) == "{}" { // empty struct
			dst.IdmapRidOptions = nil
		} else {
			return nil // data stored in dst.IdmapRidOptions, return on the first match
		}
	} else {
		dst.IdmapRidOptions = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IdmapCreate0Options)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IdmapCreate0Options) MarshalJSON() ([]byte, error) {
	if src.IdmapAdOptions != nil {
		return json.Marshal(&src.IdmapAdOptions)
	}

	if src.IdmapAutoridOptions != nil {
		return json.Marshal(&src.IdmapAutoridOptions)
	}

	if src.IdmapLdapOptions != nil {
		return json.Marshal(&src.IdmapLdapOptions)
	}

	if src.IdmapNssOptions != nil {
		return json.Marshal(&src.IdmapNssOptions)
	}

	if src.IdmapRfc2307Options != nil {
		return json.Marshal(&src.IdmapRfc2307Options)
	}

	if src.IdmapRidOptions != nil {
		return json.Marshal(&src.IdmapRidOptions)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIdmapCreate0Options struct {
	value *IdmapCreate0Options
	isSet bool
}

func (v NullableIdmapCreate0Options) Get() *IdmapCreate0Options {
	return v.value
}

func (v *NullableIdmapCreate0Options) Set(val *IdmapCreate0Options) {
	v.value = val
	v.isSet = true
}

func (v NullableIdmapCreate0Options) IsSet() bool {
	return v.isSet
}

func (v *NullableIdmapCreate0Options) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdmapCreate0Options(val *IdmapCreate0Options) *NullableIdmapCreate0Options {
	return &NullableIdmapCreate0Options{value: val, isSet: true}
}

func (v NullableIdmapCreate0Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdmapCreate0Options) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
