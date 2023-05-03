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

// checks if the IdmapLdapOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdmapLdapOptions{}

// IdmapLdapOptions struct for IdmapLdapOptions
type IdmapLdapOptions struct {
	LdapBaseDn           *string `json:"ldap_base_dn,omitempty"`
	LdapUserDn           *string `json:"ldap_user_dn,omitempty"`
	LdapUserDnPassword   *string `json:"ldap_user_dn_password,omitempty"`
	LdapUrl              *string `json:"ldap_url,omitempty"`
	Readonly             *bool   `json:"readonly,omitempty"`
	Ssl                  *string `json:"ssl,omitempty"`
	ValidateCertificates *bool   `json:"validate_certificates,omitempty"`
}

// NewIdmapLdapOptions instantiates a new IdmapLdapOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdmapLdapOptions() *IdmapLdapOptions {
	this := IdmapLdapOptions{}
	var readonly bool
	this.Readonly = &readonly
	var ssl string
	this.Ssl = &ssl
	var validateCertificates bool
	this.ValidateCertificates = &validateCertificates
	return &this
}

// NewIdmapLdapOptionsWithDefaults instantiates a new IdmapLdapOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdmapLdapOptionsWithDefaults() *IdmapLdapOptions {
	this := IdmapLdapOptions{}
	var readonly bool
	this.Readonly = &readonly
	var ssl string
	this.Ssl = &ssl
	var validateCertificates bool
	this.ValidateCertificates = &validateCertificates
	return &this
}

// GetLdapBaseDn returns the LdapBaseDn field value if set, zero value otherwise.
func (o *IdmapLdapOptions) GetLdapBaseDn() string {
	if o == nil || IsNil(o.LdapBaseDn) {
		var ret string
		return ret
	}
	return *o.LdapBaseDn
}

// GetLdapBaseDnOk returns a tuple with the LdapBaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapLdapOptions) GetLdapBaseDnOk() (*string, bool) {
	if o == nil || IsNil(o.LdapBaseDn) {
		return nil, false
	}
	return o.LdapBaseDn, true
}

// HasLdapBaseDn returns a boolean if a field has been set.
func (o *IdmapLdapOptions) HasLdapBaseDn() bool {
	if o != nil && !IsNil(o.LdapBaseDn) {
		return true
	}

	return false
}

// SetLdapBaseDn gets a reference to the given string and assigns it to the LdapBaseDn field.
func (o *IdmapLdapOptions) SetLdapBaseDn(v string) {
	o.LdapBaseDn = &v
}

// GetLdapUserDn returns the LdapUserDn field value if set, zero value otherwise.
func (o *IdmapLdapOptions) GetLdapUserDn() string {
	if o == nil || IsNil(o.LdapUserDn) {
		var ret string
		return ret
	}
	return *o.LdapUserDn
}

// GetLdapUserDnOk returns a tuple with the LdapUserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapLdapOptions) GetLdapUserDnOk() (*string, bool) {
	if o == nil || IsNil(o.LdapUserDn) {
		return nil, false
	}
	return o.LdapUserDn, true
}

// HasLdapUserDn returns a boolean if a field has been set.
func (o *IdmapLdapOptions) HasLdapUserDn() bool {
	if o != nil && !IsNil(o.LdapUserDn) {
		return true
	}

	return false
}

// SetLdapUserDn gets a reference to the given string and assigns it to the LdapUserDn field.
func (o *IdmapLdapOptions) SetLdapUserDn(v string) {
	o.LdapUserDn = &v
}

// GetLdapUserDnPassword returns the LdapUserDnPassword field value if set, zero value otherwise.
func (o *IdmapLdapOptions) GetLdapUserDnPassword() string {
	if o == nil || IsNil(o.LdapUserDnPassword) {
		var ret string
		return ret
	}
	return *o.LdapUserDnPassword
}

// GetLdapUserDnPasswordOk returns a tuple with the LdapUserDnPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapLdapOptions) GetLdapUserDnPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.LdapUserDnPassword) {
		return nil, false
	}
	return o.LdapUserDnPassword, true
}

// HasLdapUserDnPassword returns a boolean if a field has been set.
func (o *IdmapLdapOptions) HasLdapUserDnPassword() bool {
	if o != nil && !IsNil(o.LdapUserDnPassword) {
		return true
	}

	return false
}

// SetLdapUserDnPassword gets a reference to the given string and assigns it to the LdapUserDnPassword field.
func (o *IdmapLdapOptions) SetLdapUserDnPassword(v string) {
	o.LdapUserDnPassword = &v
}

// GetLdapUrl returns the LdapUrl field value if set, zero value otherwise.
func (o *IdmapLdapOptions) GetLdapUrl() string {
	if o == nil || IsNil(o.LdapUrl) {
		var ret string
		return ret
	}
	return *o.LdapUrl
}

// GetLdapUrlOk returns a tuple with the LdapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapLdapOptions) GetLdapUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LdapUrl) {
		return nil, false
	}
	return o.LdapUrl, true
}

// HasLdapUrl returns a boolean if a field has been set.
func (o *IdmapLdapOptions) HasLdapUrl() bool {
	if o != nil && !IsNil(o.LdapUrl) {
		return true
	}

	return false
}

// SetLdapUrl gets a reference to the given string and assigns it to the LdapUrl field.
func (o *IdmapLdapOptions) SetLdapUrl(v string) {
	o.LdapUrl = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *IdmapLdapOptions) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapLdapOptions) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *IdmapLdapOptions) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *IdmapLdapOptions) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *IdmapLdapOptions) GetSsl() string {
	if o == nil || IsNil(o.Ssl) {
		var ret string
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapLdapOptions) GetSslOk() (*string, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *IdmapLdapOptions) HasSsl() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given string and assigns it to the Ssl field.
func (o *IdmapLdapOptions) SetSsl(v string) {
	o.Ssl = &v
}

// GetValidateCertificates returns the ValidateCertificates field value if set, zero value otherwise.
func (o *IdmapLdapOptions) GetValidateCertificates() bool {
	if o == nil || IsNil(o.ValidateCertificates) {
		var ret bool
		return ret
	}
	return *o.ValidateCertificates
}

// GetValidateCertificatesOk returns a tuple with the ValidateCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdmapLdapOptions) GetValidateCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateCertificates) {
		return nil, false
	}
	return o.ValidateCertificates, true
}

// HasValidateCertificates returns a boolean if a field has been set.
func (o *IdmapLdapOptions) HasValidateCertificates() bool {
	if o != nil && !IsNil(o.ValidateCertificates) {
		return true
	}

	return false
}

// SetValidateCertificates gets a reference to the given bool and assigns it to the ValidateCertificates field.
func (o *IdmapLdapOptions) SetValidateCertificates(v bool) {
	o.ValidateCertificates = &v
}

func (o IdmapLdapOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdmapLdapOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LdapBaseDn) {
		toSerialize["ldap_base_dn"] = o.LdapBaseDn
	}
	if !IsNil(o.LdapUserDn) {
		toSerialize["ldap_user_dn"] = o.LdapUserDn
	}
	if !IsNil(o.LdapUserDnPassword) {
		toSerialize["ldap_user_dn_password"] = o.LdapUserDnPassword
	}
	if !IsNil(o.LdapUrl) {
		toSerialize["ldap_url"] = o.LdapUrl
	}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !IsNil(o.Ssl) {
		toSerialize["ssl"] = o.Ssl
	}
	if !IsNil(o.ValidateCertificates) {
		toSerialize["validate_certificates"] = o.ValidateCertificates
	}
	return toSerialize, nil
}

type NullableIdmapLdapOptions struct {
	value *IdmapLdapOptions
	isSet bool
}

func (v NullableIdmapLdapOptions) Get() *IdmapLdapOptions {
	return v.value
}

func (v *NullableIdmapLdapOptions) Set(val *IdmapLdapOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableIdmapLdapOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableIdmapLdapOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdmapLdapOptions(val *IdmapLdapOptions) *NullableIdmapLdapOptions {
	return &NullableIdmapLdapOptions{value: val, isSet: true}
}

func (v NullableIdmapLdapOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdmapLdapOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}