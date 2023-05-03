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

// checks if the CertExtensions1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertExtensions1{}

// CertExtensions1 `cert_extensions` can be specified to set X509v3 extensions.
type CertExtensions1 struct {
	BasicConstraints       *BasicConstraints1      `json:"BasicConstraints,omitempty"`
	AuthorityKeyIdentifier *AuthorityKeyIdentifier `json:"AuthorityKeyIdentifier,omitempty"`
	ExtendedKeyUsage       *ExtendedKeyUsage1      `json:"ExtendedKeyUsage,omitempty"`
	KeyUsage               *KeyUsage1              `json:"KeyUsage,omitempty"`
}

// NewCertExtensions1 instantiates a new CertExtensions1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertExtensions1() *CertExtensions1 {
	this := CertExtensions1{}
	var basicConstraints BasicConstraints1
	this.BasicConstraints = &basicConstraints
	var authorityKeyIdentifier AuthorityKeyIdentifier
	this.AuthorityKeyIdentifier = &authorityKeyIdentifier
	var extendedKeyUsage ExtendedKeyUsage1
	this.ExtendedKeyUsage = &extendedKeyUsage
	var keyUsage KeyUsage1
	this.KeyUsage = &keyUsage
	return &this
}

// NewCertExtensions1WithDefaults instantiates a new CertExtensions1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertExtensions1WithDefaults() *CertExtensions1 {
	this := CertExtensions1{}
	var basicConstraints BasicConstraints1
	this.BasicConstraints = &basicConstraints
	var authorityKeyIdentifier AuthorityKeyIdentifier
	this.AuthorityKeyIdentifier = &authorityKeyIdentifier
	var extendedKeyUsage ExtendedKeyUsage1
	this.ExtendedKeyUsage = &extendedKeyUsage
	var keyUsage KeyUsage1
	this.KeyUsage = &keyUsage
	return &this
}

// GetBasicConstraints returns the BasicConstraints field value if set, zero value otherwise.
func (o *CertExtensions1) GetBasicConstraints() BasicConstraints1 {
	if o == nil || IsNil(o.BasicConstraints) {
		var ret BasicConstraints1
		return ret
	}
	return *o.BasicConstraints
}

// GetBasicConstraintsOk returns a tuple with the BasicConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions1) GetBasicConstraintsOk() (*BasicConstraints1, bool) {
	if o == nil || IsNil(o.BasicConstraints) {
		return nil, false
	}
	return o.BasicConstraints, true
}

// HasBasicConstraints returns a boolean if a field has been set.
func (o *CertExtensions1) HasBasicConstraints() bool {
	if o != nil && !IsNil(o.BasicConstraints) {
		return true
	}

	return false
}

// SetBasicConstraints gets a reference to the given BasicConstraints1 and assigns it to the BasicConstraints field.
func (o *CertExtensions1) SetBasicConstraints(v BasicConstraints1) {
	o.BasicConstraints = &v
}

// GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field value if set, zero value otherwise.
func (o *CertExtensions1) GetAuthorityKeyIdentifier() AuthorityKeyIdentifier {
	if o == nil || IsNil(o.AuthorityKeyIdentifier) {
		var ret AuthorityKeyIdentifier
		return ret
	}
	return *o.AuthorityKeyIdentifier
}

// GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions1) GetAuthorityKeyIdentifierOk() (*AuthorityKeyIdentifier, bool) {
	if o == nil || IsNil(o.AuthorityKeyIdentifier) {
		return nil, false
	}
	return o.AuthorityKeyIdentifier, true
}

// HasAuthorityKeyIdentifier returns a boolean if a field has been set.
func (o *CertExtensions1) HasAuthorityKeyIdentifier() bool {
	if o != nil && !IsNil(o.AuthorityKeyIdentifier) {
		return true
	}

	return false
}

// SetAuthorityKeyIdentifier gets a reference to the given AuthorityKeyIdentifier and assigns it to the AuthorityKeyIdentifier field.
func (o *CertExtensions1) SetAuthorityKeyIdentifier(v AuthorityKeyIdentifier) {
	o.AuthorityKeyIdentifier = &v
}

// GetExtendedKeyUsage returns the ExtendedKeyUsage field value if set, zero value otherwise.
func (o *CertExtensions1) GetExtendedKeyUsage() ExtendedKeyUsage1 {
	if o == nil || IsNil(o.ExtendedKeyUsage) {
		var ret ExtendedKeyUsage1
		return ret
	}
	return *o.ExtendedKeyUsage
}

// GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions1) GetExtendedKeyUsageOk() (*ExtendedKeyUsage1, bool) {
	if o == nil || IsNil(o.ExtendedKeyUsage) {
		return nil, false
	}
	return o.ExtendedKeyUsage, true
}

// HasExtendedKeyUsage returns a boolean if a field has been set.
func (o *CertExtensions1) HasExtendedKeyUsage() bool {
	if o != nil && !IsNil(o.ExtendedKeyUsage) {
		return true
	}

	return false
}

// SetExtendedKeyUsage gets a reference to the given ExtendedKeyUsage1 and assigns it to the ExtendedKeyUsage field.
func (o *CertExtensions1) SetExtendedKeyUsage(v ExtendedKeyUsage1) {
	o.ExtendedKeyUsage = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *CertExtensions1) GetKeyUsage() KeyUsage1 {
	if o == nil || IsNil(o.KeyUsage) {
		var ret KeyUsage1
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions1) GetKeyUsageOk() (*KeyUsage1, bool) {
	if o == nil || IsNil(o.KeyUsage) {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *CertExtensions1) HasKeyUsage() bool {
	if o != nil && !IsNil(o.KeyUsage) {
		return true
	}

	return false
}

// SetKeyUsage gets a reference to the given KeyUsage1 and assigns it to the KeyUsage field.
func (o *CertExtensions1) SetKeyUsage(v KeyUsage1) {
	o.KeyUsage = &v
}

func (o CertExtensions1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertExtensions1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasicConstraints) {
		toSerialize["BasicConstraints"] = o.BasicConstraints
	}
	if !IsNil(o.AuthorityKeyIdentifier) {
		toSerialize["AuthorityKeyIdentifier"] = o.AuthorityKeyIdentifier
	}
	if !IsNil(o.ExtendedKeyUsage) {
		toSerialize["ExtendedKeyUsage"] = o.ExtendedKeyUsage
	}
	if !IsNil(o.KeyUsage) {
		toSerialize["KeyUsage"] = o.KeyUsage
	}
	return toSerialize, nil
}

type NullableCertExtensions1 struct {
	value *CertExtensions1
	isSet bool
}

func (v NullableCertExtensions1) Get() *CertExtensions1 {
	return v.value
}

func (v *NullableCertExtensions1) Set(val *CertExtensions1) {
	v.value = val
	v.isSet = true
}

func (v NullableCertExtensions1) IsSet() bool {
	return v.isSet
}

func (v *NullableCertExtensions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertExtensions1(val *CertExtensions1) *NullableCertExtensions1 {
	return &NullableCertExtensions1{value: val, isSet: true}
}

func (v NullableCertExtensions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertExtensions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}