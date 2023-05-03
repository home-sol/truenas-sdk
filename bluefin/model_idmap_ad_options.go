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

// checks if the IdmapAdOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdmapAdOptions{}

// IdmapAdOptions struct for IdmapAdOptions
type IdmapAdOptions struct {
	SchemaMode       *string `json:"schema_mode,omitempty"`
	UnixPrimaryGroup *bool   `json:"unix_primary_group,omitempty"`
	UnixNssInfo      *bool   `json:"unix_nss_info,omitempty"`
}

// NewIdmapAdOptions instantiates a new IdmapAdOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdmapAdOptions() *IdmapAdOptions {
	this := IdmapAdOptions{}
	var schemaMode string
	this.SchemaMode = &schemaMode
	var unixPrimaryGroup bool
	this.UnixPrimaryGroup = &unixPrimaryGroup
	var unixNssInfo bool
	this.UnixNssInfo = &unixNssInfo
	return &this
}

// NewIdmapAdOptionsWithDefaults instantiates a new IdmapAdOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdmapAdOptionsWithDefaults() *IdmapAdOptions {
	this := IdmapAdOptions{}
	var schemaMode string
	this.SchemaMode = &schemaMode
	var unixPrimaryGroup bool
	this.UnixPrimaryGroup = &unixPrimaryGroup
	var unixNssInfo bool
	this.UnixNssInfo = &unixNssInfo
	return &this
}

// GetSchemaMode returns the SchemaMode field value if set, zero value otherwise.
func (o *IdmapAdOptions) GetSchemaMode() string {
	if o == nil || IsNil(o.SchemaMode) {
		var ret string
		return ret
	}
	return *o.SchemaMode
}

// GetSchemaModeOk returns a tuple with the SchemaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapAdOptions) GetSchemaModeOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaMode) {
		return nil, false
	}
	return o.SchemaMode, true
}

// HasSchemaMode returns a boolean if a field has been set.
func (o *IdmapAdOptions) HasSchemaMode() bool {
	if o != nil && !IsNil(o.SchemaMode) {
		return true
	}

	return false
}

// SetSchemaMode gets a reference to the given string and assigns it to the SchemaMode field.
func (o *IdmapAdOptions) SetSchemaMode(v string) {
	o.SchemaMode = &v
}

// GetUnixPrimaryGroup returns the UnixPrimaryGroup field value if set, zero value otherwise.
func (o *IdmapAdOptions) GetUnixPrimaryGroup() bool {
	if o == nil || IsNil(o.UnixPrimaryGroup) {
		var ret bool
		return ret
	}
	return *o.UnixPrimaryGroup
}

// GetUnixPrimaryGroupOk returns a tuple with the UnixPrimaryGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapAdOptions) GetUnixPrimaryGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.UnixPrimaryGroup) {
		return nil, false
	}
	return o.UnixPrimaryGroup, true
}

// HasUnixPrimaryGroup returns a boolean if a field has been set.
func (o *IdmapAdOptions) HasUnixPrimaryGroup() bool {
	if o != nil && !IsNil(o.UnixPrimaryGroup) {
		return true
	}

	return false
}

// SetUnixPrimaryGroup gets a reference to the given bool and assigns it to the UnixPrimaryGroup field.
func (o *IdmapAdOptions) SetUnixPrimaryGroup(v bool) {
	o.UnixPrimaryGroup = &v
}

// GetUnixNssInfo returns the UnixNssInfo field value if set, zero value otherwise.
func (o *IdmapAdOptions) GetUnixNssInfo() bool {
	if o == nil || IsNil(o.UnixNssInfo) {
		var ret bool
		return ret
	}
	return *o.UnixNssInfo
}

// GetUnixNssInfoOk returns a tuple with the UnixNssInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapAdOptions) GetUnixNssInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.UnixNssInfo) {
		return nil, false
	}
	return o.UnixNssInfo, true
}

// HasUnixNssInfo returns a boolean if a field has been set.
func (o *IdmapAdOptions) HasUnixNssInfo() bool {
	if o != nil && !IsNil(o.UnixNssInfo) {
		return true
	}

	return false
}

// SetUnixNssInfo gets a reference to the given bool and assigns it to the UnixNssInfo field.
func (o *IdmapAdOptions) SetUnixNssInfo(v bool) {
	o.UnixNssInfo = &v
}

func (o IdmapAdOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdmapAdOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchemaMode) {
		toSerialize["schema_mode"] = o.SchemaMode
	}
	if !IsNil(o.UnixPrimaryGroup) {
		toSerialize["unix_primary_group"] = o.UnixPrimaryGroup
	}
	if !IsNil(o.UnixNssInfo) {
		toSerialize["unix_nss_info"] = o.UnixNssInfo
	}
	return toSerialize, nil
}

type NullableIdmapAdOptions struct {
	value *IdmapAdOptions
	isSet bool
}

func (v NullableIdmapAdOptions) Get() *IdmapAdOptions {
	return v.value
}

func (v *NullableIdmapAdOptions) Set(val *IdmapAdOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableIdmapAdOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableIdmapAdOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdmapAdOptions(val *IdmapAdOptions) *NullableIdmapAdOptions {
	return &NullableIdmapAdOptions{value: val, isSet: true}
}

func (v NullableIdmapAdOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdmapAdOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}