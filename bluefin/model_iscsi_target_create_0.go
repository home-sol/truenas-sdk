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

// checks if the IscsiTargetCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IscsiTargetCreate0{}

// IscsiTargetCreate0 struct for IscsiTargetCreate0
type IscsiTargetCreate0 struct {
	Name  *string        `json:"name,omitempty"`
	Alias NullableString `json:"alias,omitempty"`
	Mode  *string        `json:"mode,omitempty"`
	// `groups` is a list of group dictionaries which provide information related to using a `portal`, `initiator`, `authmethod` and `auth` with this target. `auth` represents a valid iSCSI Authorized Access and defaults to null.
	Groups       []Group  `json:"groups,omitempty"`
	AuthNetworks []string `json:"auth_networks,omitempty"`
}

// NewIscsiTargetCreate0 instantiates a new IscsiTargetCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiTargetCreate0() *IscsiTargetCreate0 {
	this := IscsiTargetCreate0{}
	var mode string
	this.Mode = &mode
	return &this
}

// NewIscsiTargetCreate0WithDefaults instantiates a new IscsiTargetCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiTargetCreate0WithDefaults() *IscsiTargetCreate0 {
	this := IscsiTargetCreate0{}
	var mode string
	this.Mode = &mode
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IscsiTargetCreate0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiTargetCreate0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IscsiTargetCreate0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IscsiTargetCreate0) SetName(v string) {
	o.Name = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiTargetCreate0) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiTargetCreate0) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *IscsiTargetCreate0) HasAlias() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *IscsiTargetCreate0) SetAlias(v string) {
	o.Alias.Set(&v)
}

// SetAliasNil sets the value for Alias to be an explicit nil
func (o *IscsiTargetCreate0) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *IscsiTargetCreate0) UnsetAlias() {
	o.Alias.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *IscsiTargetCreate0) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiTargetCreate0) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *IscsiTargetCreate0) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *IscsiTargetCreate0) SetMode(v string) {
	o.Mode = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *IscsiTargetCreate0) GetGroups() []Group {
	if o == nil || IsNil(o.Groups) {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiTargetCreate0) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *IscsiTargetCreate0) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *IscsiTargetCreate0) SetGroups(v []Group) {
	o.Groups = v
}

// GetAuthNetworks returns the AuthNetworks field value if set, zero value otherwise.
func (o *IscsiTargetCreate0) GetAuthNetworks() []string {
	if o == nil || IsNil(o.AuthNetworks) {
		var ret []string
		return ret
	}
	return o.AuthNetworks
}

// GetAuthNetworksOk returns a tuple with the AuthNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiTargetCreate0) GetAuthNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthNetworks) {
		return nil, false
	}
	return o.AuthNetworks, true
}

// HasAuthNetworks returns a boolean if a field has been set.
func (o *IscsiTargetCreate0) HasAuthNetworks() bool {
	if o != nil && !IsNil(o.AuthNetworks) {
		return true
	}

	return false
}

// SetAuthNetworks gets a reference to the given []string and assigns it to the AuthNetworks field.
func (o *IscsiTargetCreate0) SetAuthNetworks(v []string) {
	o.AuthNetworks = v
}

func (o IscsiTargetCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IscsiTargetCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.AuthNetworks) {
		toSerialize["auth_networks"] = o.AuthNetworks
	}
	return toSerialize, nil
}

type NullableIscsiTargetCreate0 struct {
	value *IscsiTargetCreate0
	isSet bool
}

func (v NullableIscsiTargetCreate0) Get() *IscsiTargetCreate0 {
	return v.value
}

func (v *NullableIscsiTargetCreate0) Set(val *IscsiTargetCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiTargetCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiTargetCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiTargetCreate0(val *IscsiTargetCreate0) *NullableIscsiTargetCreate0 {
	return &NullableIscsiTargetCreate0{value: val, isSet: true}
}

func (v NullableIscsiTargetCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiTargetCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
