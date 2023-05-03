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

// checks if the KeychaincredentialCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeychaincredentialCreate0{}

// KeychaincredentialCreate0 struct for KeychaincredentialCreate0
type KeychaincredentialCreate0 struct {
	// Every Keychain Credential has a `name` which is used to distinguish it from others.
	Name *string `json:"name,omitempty"`
	// The following `type`s are supported:  * `SSH_KEY_PAIR`    Which `attributes` are:    * `private_key`    * `public_key` (which can be omitted and thus automatically derived from private key)    At least one attribute is required.
	Type *string `json:"type,omitempty"`
	// The following `type`s are supported:  * `SSH_KEY_PAIR`    Which `attributes` are:    * `private_key`    * `public_key` (which can be omitted and thus automatically derived from private key)    At least one attribute is required.  * `SSH_CREDENTIALS`    Which `attributes` are:    * `host`    * `port` (default 22)    * `username` (default root)    * `private_key` (Keychain Credential ID)    * `remote_host_key` (you can use `keychaincredential.remote_ssh_host_key_scan` do discover it)    * `cipher`: one of `STANDARD`, `FAST`, or `DISABLED` (last requires special support from both SSH server and      client)    * `connect_timeout` (default 10)
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// NewKeychaincredentialCreate0 instantiates a new KeychaincredentialCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeychaincredentialCreate0() *KeychaincredentialCreate0 {
	this := KeychaincredentialCreate0{}
	return &this
}

// NewKeychaincredentialCreate0WithDefaults instantiates a new KeychaincredentialCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeychaincredentialCreate0WithDefaults() *KeychaincredentialCreate0 {
	this := KeychaincredentialCreate0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeychaincredentialCreate0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeychaincredentialCreate0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeychaincredentialCreate0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeychaincredentialCreate0) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KeychaincredentialCreate0) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeychaincredentialCreate0) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KeychaincredentialCreate0) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KeychaincredentialCreate0) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *KeychaincredentialCreate0) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeychaincredentialCreate0) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *KeychaincredentialCreate0) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *KeychaincredentialCreate0) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o KeychaincredentialCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeychaincredentialCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableKeychaincredentialCreate0 struct {
	value *KeychaincredentialCreate0
	isSet bool
}

func (v NullableKeychaincredentialCreate0) Get() *KeychaincredentialCreate0 {
	return v.value
}

func (v *NullableKeychaincredentialCreate0) Set(val *KeychaincredentialCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableKeychaincredentialCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableKeychaincredentialCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeychaincredentialCreate0(val *KeychaincredentialCreate0) *NullableKeychaincredentialCreate0 {
	return &NullableKeychaincredentialCreate0{value: val, isSet: true}
}

func (v NullableKeychaincredentialCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeychaincredentialCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
