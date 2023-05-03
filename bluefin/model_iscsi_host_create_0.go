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

// checks if the IscsiHostCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IscsiHostCreate0{}

// IscsiHostCreate0 struct for IscsiHostCreate0
type IscsiHostCreate0 struct {
	// `ip` indicates an IP address of the host.
	Ip *string `json:"ip,omitempty"`
	// `description` is a human-readable name for the host.
	Description        *string  `json:"description,omitempty"`
	Iqns               []string `json:"iqns,omitempty"`
	AddedAutomatically *bool    `json:"added_automatically,omitempty"`
}

// NewIscsiHostCreate0 instantiates a new IscsiHostCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiHostCreate0() *IscsiHostCreate0 {
	this := IscsiHostCreate0{}
	var description string
	this.Description = &description
	var addedAutomatically bool
	this.AddedAutomatically = &addedAutomatically
	return &this
}

// NewIscsiHostCreate0WithDefaults instantiates a new IscsiHostCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiHostCreate0WithDefaults() *IscsiHostCreate0 {
	this := IscsiHostCreate0{}
	var description string
	this.Description = &description
	var addedAutomatically bool
	this.AddedAutomatically = &addedAutomatically
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *IscsiHostCreate0) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiHostCreate0) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *IscsiHostCreate0) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *IscsiHostCreate0) SetIp(v string) {
	o.Ip = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IscsiHostCreate0) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiHostCreate0) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IscsiHostCreate0) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IscsiHostCreate0) SetDescription(v string) {
	o.Description = &v
}

// GetIqns returns the Iqns field value if set, zero value otherwise.
func (o *IscsiHostCreate0) GetIqns() []string {
	if o == nil || IsNil(o.Iqns) {
		var ret []string
		return ret
	}
	return o.Iqns
}

// GetIqnsOk returns a tuple with the Iqns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiHostCreate0) GetIqnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Iqns) {
		return nil, false
	}
	return o.Iqns, true
}

// HasIqns returns a boolean if a field has been set.
func (o *IscsiHostCreate0) HasIqns() bool {
	if o != nil && !IsNil(o.Iqns) {
		return true
	}

	return false
}

// SetIqns gets a reference to the given []string and assigns it to the Iqns field.
func (o *IscsiHostCreate0) SetIqns(v []string) {
	o.Iqns = v
}

// GetAddedAutomatically returns the AddedAutomatically field value if set, zero value otherwise.
func (o *IscsiHostCreate0) GetAddedAutomatically() bool {
	if o == nil || IsNil(o.AddedAutomatically) {
		var ret bool
		return ret
	}
	return *o.AddedAutomatically
}

// GetAddedAutomaticallyOk returns a tuple with the AddedAutomatically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiHostCreate0) GetAddedAutomaticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.AddedAutomatically) {
		return nil, false
	}
	return o.AddedAutomatically, true
}

// HasAddedAutomatically returns a boolean if a field has been set.
func (o *IscsiHostCreate0) HasAddedAutomatically() bool {
	if o != nil && !IsNil(o.AddedAutomatically) {
		return true
	}

	return false
}

// SetAddedAutomatically gets a reference to the given bool and assigns it to the AddedAutomatically field.
func (o *IscsiHostCreate0) SetAddedAutomatically(v bool) {
	o.AddedAutomatically = &v
}

func (o IscsiHostCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IscsiHostCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Iqns) {
		toSerialize["iqns"] = o.Iqns
	}
	if !IsNil(o.AddedAutomatically) {
		toSerialize["added_automatically"] = o.AddedAutomatically
	}
	return toSerialize, nil
}

type NullableIscsiHostCreate0 struct {
	value *IscsiHostCreate0
	isSet bool
}

func (v NullableIscsiHostCreate0) Get() *IscsiHostCreate0 {
	return v.value
}

func (v *NullableIscsiHostCreate0) Set(val *IscsiHostCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiHostCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiHostCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiHostCreate0(val *IscsiHostCreate0) *NullableIscsiHostCreate0 {
	return &NullableIscsiHostCreate0{value: val, isSet: true}
}

func (v NullableIscsiHostCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiHostCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
