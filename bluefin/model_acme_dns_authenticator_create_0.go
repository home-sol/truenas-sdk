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

// checks if the AcmeDnsAuthenticatorCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcmeDnsAuthenticatorCreate0{}

// AcmeDnsAuthenticatorCreate0 struct for AcmeDnsAuthenticatorCreate0
type AcmeDnsAuthenticatorCreate0 struct {
	Authenticator *string `json:"authenticator,omitempty"`
	// Specific attributes of each `authenticator`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// User defined name of authenticator
	Name *string `json:"name,omitempty"`
}

// NewAcmeDnsAuthenticatorCreate0 instantiates a new AcmeDnsAuthenticatorCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeDnsAuthenticatorCreate0() *AcmeDnsAuthenticatorCreate0 {
	this := AcmeDnsAuthenticatorCreate0{}
	return &this
}

// NewAcmeDnsAuthenticatorCreate0WithDefaults instantiates a new AcmeDnsAuthenticatorCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeDnsAuthenticatorCreate0WithDefaults() *AcmeDnsAuthenticatorCreate0 {
	this := AcmeDnsAuthenticatorCreate0{}
	return &this
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorCreate0) GetAuthenticator() string {
	if o == nil || IsNil(o.Authenticator) {
		var ret string
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorCreate0) GetAuthenticatorOk() (*string, bool) {
	if o == nil || IsNil(o.Authenticator) {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorCreate0) HasAuthenticator() bool {
	if o != nil && !IsNil(o.Authenticator) {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given string and assigns it to the Authenticator field.
func (o *AcmeDnsAuthenticatorCreate0) SetAuthenticator(v string) {
	o.Authenticator = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorCreate0) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorCreate0) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorCreate0) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *AcmeDnsAuthenticatorCreate0) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorCreate0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorCreate0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorCreate0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AcmeDnsAuthenticatorCreate0) SetName(v string) {
	o.Name = &v
}

func (o AcmeDnsAuthenticatorCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcmeDnsAuthenticatorCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authenticator) {
		toSerialize["authenticator"] = o.Authenticator
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAcmeDnsAuthenticatorCreate0 struct {
	value *AcmeDnsAuthenticatorCreate0
	isSet bool
}

func (v NullableAcmeDnsAuthenticatorCreate0) Get() *AcmeDnsAuthenticatorCreate0 {
	return v.value
}

func (v *NullableAcmeDnsAuthenticatorCreate0) Set(val *AcmeDnsAuthenticatorCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeDnsAuthenticatorCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeDnsAuthenticatorCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeDnsAuthenticatorCreate0(val *AcmeDnsAuthenticatorCreate0) *NullableAcmeDnsAuthenticatorCreate0 {
	return &NullableAcmeDnsAuthenticatorCreate0{value: val, isSet: true}
}

func (v NullableAcmeDnsAuthenticatorCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeDnsAuthenticatorCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}