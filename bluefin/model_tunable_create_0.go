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

// checks if the TunableCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunableCreate0{}

// TunableCreate0 struct for TunableCreate0
type TunableCreate0 struct {
	// If `type` is `SYSCTL` then `var` is a sysctl name (e.g. `kernel.watchdog`) and `value` is its corresponding value (e.g. `0`).
	Type *string `json:"type,omitempty"`
	// If `type` is `SYSCTL` then `var` is a sysctl name (e.g. `kernel.watchdog`) and `value` is its corresponding value (e.g. `0`).
	Var *string `json:"var,omitempty"`
	// If `type` is `SYSCTL` then `var` is a sysctl name (e.g. `kernel.watchdog`) and `value` is its corresponding value (e.g. `0`).
	Value   *string `json:"value,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

// NewTunableCreate0 instantiates a new TunableCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunableCreate0() *TunableCreate0 {
	this := TunableCreate0{}
	var type_ string
	this.Type = &type_
	var comment string
	this.Comment = &comment
	var enabled bool
	this.Enabled = &enabled
	return &this
}

// NewTunableCreate0WithDefaults instantiates a new TunableCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunableCreate0WithDefaults() *TunableCreate0 {
	this := TunableCreate0{}
	var type_ string
	this.Type = &type_
	var comment string
	this.Comment = &comment
	var enabled bool
	this.Enabled = &enabled
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TunableCreate0) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TunableCreate0) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TunableCreate0) SetType(v string) {
	o.Type = &v
}

// GetVar returns the Var field value if set, zero value otherwise.
func (o *TunableCreate0) GetVar() string {
	if o == nil || IsNil(o.Var) {
		var ret string
		return ret
	}
	return *o.Var
}

// GetVarOk returns a tuple with the Var field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetVarOk() (*string, bool) {
	if o == nil || IsNil(o.Var) {
		return nil, false
	}
	return o.Var, true
}

// HasVar returns a boolean if a field has been set.
func (o *TunableCreate0) HasVar() bool {
	if o != nil && !IsNil(o.Var) {
		return true
	}

	return false
}

// SetVar gets a reference to the given string and assigns it to the Var field.
func (o *TunableCreate0) SetVar(v string) {
	o.Var = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TunableCreate0) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TunableCreate0) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TunableCreate0) SetValue(v string) {
	o.Value = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TunableCreate0) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TunableCreate0) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TunableCreate0) SetComment(v string) {
	o.Comment = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TunableCreate0) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TunableCreate0) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TunableCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o TunableCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunableCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Var) {
		toSerialize["var"] = o.Var
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableTunableCreate0 struct {
	value *TunableCreate0
	isSet bool
}

func (v NullableTunableCreate0) Get() *TunableCreate0 {
	return v.value
}

func (v *NullableTunableCreate0) Set(val *TunableCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableTunableCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableTunableCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunableCreate0(val *TunableCreate0) *NullableTunableCreate0 {
	return &NullableTunableCreate0{value: val, isSet: true}
}

func (v NullableTunableCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunableCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
