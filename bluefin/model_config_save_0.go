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

// checks if the ConfigSave0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSave0{}

// ConfigSave0 struct for ConfigSave0
type ConfigSave0 struct {
	// `secretseed` bool: When true, include password secret seed.
	Secretseed *bool `json:"secretseed,omitempty"`
	// `pool_keys` bool: IGNORED and DEPRECATED as it does not apply on SCALE systems.
	PoolKeys *bool `json:"pool_keys,omitempty"`
	// `root_authorized_keys` bool: When true, include \"/root/.ssh/authorized_keys\" file for the root user.
	RootAuthorizedKeys *bool `json:"root_authorized_keys,omitempty"`
	// `gluster_config` bool: When true, include the directory that stores the gluster configuration files.
	GlusterConfig *bool `json:"gluster_config,omitempty"`
}

// NewConfigSave0 instantiates a new ConfigSave0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSave0() *ConfigSave0 {
	this := ConfigSave0{}
	var secretseed bool
	this.Secretseed = &secretseed
	var poolKeys bool
	this.PoolKeys = &poolKeys
	var rootAuthorizedKeys bool
	this.RootAuthorizedKeys = &rootAuthorizedKeys
	var glusterConfig bool
	this.GlusterConfig = &glusterConfig
	return &this
}

// NewConfigSave0WithDefaults instantiates a new ConfigSave0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSave0WithDefaults() *ConfigSave0 {
	this := ConfigSave0{}
	var secretseed bool
	this.Secretseed = &secretseed
	var poolKeys bool
	this.PoolKeys = &poolKeys
	var rootAuthorizedKeys bool
	this.RootAuthorizedKeys = &rootAuthorizedKeys
	var glusterConfig bool
	this.GlusterConfig = &glusterConfig
	return &this
}

// GetSecretseed returns the Secretseed field value if set, zero value otherwise.
func (o *ConfigSave0) GetSecretseed() bool {
	if o == nil || IsNil(o.Secretseed) {
		var ret bool
		return ret
	}
	return *o.Secretseed
}

// GetSecretseedOk returns a tuple with the Secretseed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSave0) GetSecretseedOk() (*bool, bool) {
	if o == nil || IsNil(o.Secretseed) {
		return nil, false
	}
	return o.Secretseed, true
}

// HasSecretseed returns a boolean if a field has been set.
func (o *ConfigSave0) HasSecretseed() bool {
	if o != nil && !IsNil(o.Secretseed) {
		return true
	}

	return false
}

// SetSecretseed gets a reference to the given bool and assigns it to the Secretseed field.
func (o *ConfigSave0) SetSecretseed(v bool) {
	o.Secretseed = &v
}

// GetPoolKeys returns the PoolKeys field value if set, zero value otherwise.
func (o *ConfigSave0) GetPoolKeys() bool {
	if o == nil || IsNil(o.PoolKeys) {
		var ret bool
		return ret
	}
	return *o.PoolKeys
}

// GetPoolKeysOk returns a tuple with the PoolKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSave0) GetPoolKeysOk() (*bool, bool) {
	if o == nil || IsNil(o.PoolKeys) {
		return nil, false
	}
	return o.PoolKeys, true
}

// HasPoolKeys returns a boolean if a field has been set.
func (o *ConfigSave0) HasPoolKeys() bool {
	if o != nil && !IsNil(o.PoolKeys) {
		return true
	}

	return false
}

// SetPoolKeys gets a reference to the given bool and assigns it to the PoolKeys field.
func (o *ConfigSave0) SetPoolKeys(v bool) {
	o.PoolKeys = &v
}

// GetRootAuthorizedKeys returns the RootAuthorizedKeys field value if set, zero value otherwise.
func (o *ConfigSave0) GetRootAuthorizedKeys() bool {
	if o == nil || IsNil(o.RootAuthorizedKeys) {
		var ret bool
		return ret
	}
	return *o.RootAuthorizedKeys
}

// GetRootAuthorizedKeysOk returns a tuple with the RootAuthorizedKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSave0) GetRootAuthorizedKeysOk() (*bool, bool) {
	if o == nil || IsNil(o.RootAuthorizedKeys) {
		return nil, false
	}
	return o.RootAuthorizedKeys, true
}

// HasRootAuthorizedKeys returns a boolean if a field has been set.
func (o *ConfigSave0) HasRootAuthorizedKeys() bool {
	if o != nil && !IsNil(o.RootAuthorizedKeys) {
		return true
	}

	return false
}

// SetRootAuthorizedKeys gets a reference to the given bool and assigns it to the RootAuthorizedKeys field.
func (o *ConfigSave0) SetRootAuthorizedKeys(v bool) {
	o.RootAuthorizedKeys = &v
}

// GetGlusterConfig returns the GlusterConfig field value if set, zero value otherwise.
func (o *ConfigSave0) GetGlusterConfig() bool {
	if o == nil || IsNil(o.GlusterConfig) {
		var ret bool
		return ret
	}
	return *o.GlusterConfig
}

// GetGlusterConfigOk returns a tuple with the GlusterConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSave0) GetGlusterConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.GlusterConfig) {
		return nil, false
	}
	return o.GlusterConfig, true
}

// HasGlusterConfig returns a boolean if a field has been set.
func (o *ConfigSave0) HasGlusterConfig() bool {
	if o != nil && !IsNil(o.GlusterConfig) {
		return true
	}

	return false
}

// SetGlusterConfig gets a reference to the given bool and assigns it to the GlusterConfig field.
func (o *ConfigSave0) SetGlusterConfig(v bool) {
	o.GlusterConfig = &v
}

func (o ConfigSave0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSave0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Secretseed) {
		toSerialize["secretseed"] = o.Secretseed
	}
	if !IsNil(o.PoolKeys) {
		toSerialize["pool_keys"] = o.PoolKeys
	}
	if !IsNil(o.RootAuthorizedKeys) {
		toSerialize["root_authorized_keys"] = o.RootAuthorizedKeys
	}
	if !IsNil(o.GlusterConfig) {
		toSerialize["gluster_config"] = o.GlusterConfig
	}
	return toSerialize, nil
}

type NullableConfigSave0 struct {
	value *ConfigSave0
	isSet bool
}

func (v NullableConfigSave0) Get() *ConfigSave0 {
	return v.value
}

func (v *NullableConfigSave0) Set(val *ConfigSave0) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSave0) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSave0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSave0(val *ConfigSave0) *NullableConfigSave0 {
	return &NullableConfigSave0{value: val, isSet: true}
}

func (v NullableConfigSave0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSave0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
