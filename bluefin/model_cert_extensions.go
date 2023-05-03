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

// checks if the CertExtensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertExtensions{}

// CertExtensions `cert_extensions` can be specified to set X509v3 extensions.
type CertExtensions struct {
	BasicConstraints       *BasicConstraints       `json:"BasicConstraints,omitempty"`
	AuthorityKeyIdentifier *AuthorityKeyIdentifier `json:"AuthorityKeyIdentifier,omitempty"`
	ExtendedKeyUsage       *ExtendedKeyUsage       `json:"ExtendedKeyUsage,omitempty"`
	KeyUsage               *KeyUsage               `json:"KeyUsage,omitempty"`
}

// NewCertExtensions instantiates a new CertExtensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertExtensions() *CertExtensions {
	this := CertExtensions{}
	var basicConstraints BasicConstraints
	this.BasicConstraints = &basicConstraints
	var authorityKeyIdentifier AuthorityKeyIdentifier
	this.AuthorityKeyIdentifier = &authorityKeyIdentifier
	var extendedKeyUsage ExtendedKeyUsage
	this.ExtendedKeyUsage = &extendedKeyUsage
	var keyUsage KeyUsage
	this.KeyUsage = &keyUsage
	return &this
}

// NewCertExtensionsWithDefaults instantiates a new CertExtensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertExtensionsWithDefaults() *CertExtensions {
	this := CertExtensions{}
	var basicConstraints BasicConstraints
	this.BasicConstraints = &basicConstraints
	var authorityKeyIdentifier AuthorityKeyIdentifier
	this.AuthorityKeyIdentifier = &authorityKeyIdentifier
	var extendedKeyUsage ExtendedKeyUsage
	this.ExtendedKeyUsage = &extendedKeyUsage
	var keyUsage KeyUsage
	this.KeyUsage = &keyUsage
	return &this
}

// GetBasicConstraints returns the BasicConstraints field value if set, zero value otherwise.
func (o *CertExtensions) GetBasicConstraints() BasicConstraints {
	if o == nil || IsNil(o.BasicConstraints) {
		var ret BasicConstraints
		return ret
	}
	return *o.BasicConstraints
}

// GetBasicConstraintsOk returns a tuple with the BasicConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions) GetBasicConstraintsOk() (*BasicConstraints, bool) {
	if o == nil || IsNil(o.BasicConstraints) {
		return nil, false
	}
	return o.BasicConstraints, true
}

// HasBasicConstraints returns a boolean if a field has been set.
func (o *CertExtensions) HasBasicConstraints() bool {
	if o != nil && !IsNil(o.BasicConstraints) {
		return true
	}

	return false
}

// SetBasicConstraints gets a reference to the given BasicConstraints and assigns it to the BasicConstraints field.
func (o *CertExtensions) SetBasicConstraints(v BasicConstraints) {
	o.BasicConstraints = &v
}

// GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field value if set, zero value otherwise.
func (o *CertExtensions) GetAuthorityKeyIdentifier() AuthorityKeyIdentifier {
	if o == nil || IsNil(o.AuthorityKeyIdentifier) {
		var ret AuthorityKeyIdentifier
		return ret
	}
	return *o.AuthorityKeyIdentifier
}

// GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions) GetAuthorityKeyIdentifierOk() (*AuthorityKeyIdentifier, bool) {
	if o == nil || IsNil(o.AuthorityKeyIdentifier) {
		return nil, false
	}
	return o.AuthorityKeyIdentifier, true
}

// HasAuthorityKeyIdentifier returns a boolean if a field has been set.
func (o *CertExtensions) HasAuthorityKeyIdentifier() bool {
	if o != nil && !IsNil(o.AuthorityKeyIdentifier) {
		return true
	}

	return false
}

// SetAuthorityKeyIdentifier gets a reference to the given AuthorityKeyIdentifier and assigns it to the AuthorityKeyIdentifier field.
func (o *CertExtensions) SetAuthorityKeyIdentifier(v AuthorityKeyIdentifier) {
	o.AuthorityKeyIdentifier = &v
}

// GetExtendedKeyUsage returns the ExtendedKeyUsage field value if set, zero value otherwise.
func (o *CertExtensions) GetExtendedKeyUsage() ExtendedKeyUsage {
	if o == nil || IsNil(o.ExtendedKeyUsage) {
		var ret ExtendedKeyUsage
		return ret
	}
	return *o.ExtendedKeyUsage
}

// GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions) GetExtendedKeyUsageOk() (*ExtendedKeyUsage, bool) {
	if o == nil || IsNil(o.ExtendedKeyUsage) {
		return nil, false
	}
	return o.ExtendedKeyUsage, true
}

// HasExtendedKeyUsage returns a boolean if a field has been set.
func (o *CertExtensions) HasExtendedKeyUsage() bool {
	if o != nil && !IsNil(o.ExtendedKeyUsage) {
		return true
	}

	return false
}

// SetExtendedKeyUsage gets a reference to the given ExtendedKeyUsage and assigns it to the ExtendedKeyUsage field.
func (o *CertExtensions) SetExtendedKeyUsage(v ExtendedKeyUsage) {
	o.ExtendedKeyUsage = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *CertExtensions) GetKeyUsage() KeyUsage {
	if o == nil || IsNil(o.KeyUsage) {
		var ret KeyUsage
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertExtensions) GetKeyUsageOk() (*KeyUsage, bool) {
	if o == nil || IsNil(o.KeyUsage) {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *CertExtensions) HasKeyUsage() bool {
	if o != nil && !IsNil(o.KeyUsage) {
		return true
	}

	return false
}

// SetKeyUsage gets a reference to the given KeyUsage and assigns it to the KeyUsage field.
func (o *CertExtensions) SetKeyUsage(v KeyUsage) {
	o.KeyUsage = &v
}

func (o CertExtensions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertExtensions) ToMap() (map[string]interface{}, error) {
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

type NullableCertExtensions struct {
	value *CertExtensions
	isSet bool
}

func (v NullableCertExtensions) Get() *CertExtensions {
	return v.value
}

func (v *NullableCertExtensions) Set(val *CertExtensions) {
	v.value = val
	v.isSet = true
}

func (v NullableCertExtensions) IsSet() bool {
	return v.isSet
}

func (v *NullableCertExtensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertExtensions(val *CertExtensions) *NullableCertExtensions {
	return &NullableCertExtensions{value: val, isSet: true}
}

func (v NullableCertExtensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertExtensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
