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

// checks if the CertificateauthorityCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateauthorityCreate0{}

// CertificateauthorityCreate0 struct for CertificateauthorityCreate0
type CertificateauthorityCreate0 struct {
	Tos              *bool          `json:"tos,omitempty"`
	CsrId            *int32         `json:"csr_id,omitempty"`
	Signedby         *int32         `json:"signedby,omitempty"`
	KeyLength        *int32         `json:"key_length,omitempty"`
	RenewDays        *int32         `json:"renew_days,omitempty"`
	Type             *int32         `json:"type,omitempty"`
	Lifetime         *int32         `json:"lifetime,omitempty"`
	Serial           *int32         `json:"serial,omitempty"`
	AcmeDirectoryUri *string        `json:"acme_directory_uri,omitempty"`
	Certificate      *string        `json:"certificate,omitempty"`
	City             *string        `json:"city,omitempty"`
	Common           NullableString `json:"common,omitempty"`
	Country          *string        `json:"country,omitempty"`
	CSR              *string        `json:"CSR,omitempty"`
	// Created certificate authorities use RSA keys by default. If an Elliptic Curve Key is desired, then it can be specified with the `key_type` attribute. If the `ec_curve` attribute is not specified for the Elliptic Curve Key, default to using \"BrainpoolP384R1\" curve.
	EcCurve *string `json:"ec_curve,omitempty"`
	Email   *string `json:"email,omitempty"`
	// Created certificate authorities use RSA keys by default. If an Elliptic Curve Key is desired, then it can be specified with the `key_type` attribute. If the `ec_curve` attribute is not specified for the Elliptic Curve Key, default to using \"BrainpoolP384R1\" curve.
	KeyType            *string `json:"key_type,omitempty"`
	Name               *string `json:"name,omitempty"`
	Organization       *string `json:"organization,omitempty"`
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
	Passphrase         *string `json:"passphrase,omitempty"`
	Privatekey         *string `json:"privatekey,omitempty"`
	State              *string `json:"state,omitempty"`
	// Certificate Authorities are classified under following types with the necessary keywords to be passed for `create_type` attribute to create the respective type of certificate authority A type is selected by the Certificate Authority Service based on `create_type`. The rest of the values are validated accordingly and finally a certificate is made based on the selected type.
	CreateType        *string          `json:"create_type,omitempty"`
	DigestAlgorithm   *string          `json:"digest_algorithm,omitempty"`
	San               []string         `json:"san,omitempty"`
	CertExtensions    *CertExtensions1 `json:"cert_extensions,omitempty"`
	AddToTrustedStore *bool            `json:"add_to_trusted_store,omitempty"`
}

// NewCertificateauthorityCreate0 instantiates a new CertificateauthorityCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateauthorityCreate0() *CertificateauthorityCreate0 {
	this := CertificateauthorityCreate0{}
	var ecCurve string
	this.EcCurve = &ecCurve
	var keyType string
	this.KeyType = &keyType
	var certExtensions CertExtensions1
	this.CertExtensions = &certExtensions
	var addToTrustedStore bool
	this.AddToTrustedStore = &addToTrustedStore
	return &this
}

// NewCertificateauthorityCreate0WithDefaults instantiates a new CertificateauthorityCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateauthorityCreate0WithDefaults() *CertificateauthorityCreate0 {
	this := CertificateauthorityCreate0{}
	var ecCurve string
	this.EcCurve = &ecCurve
	var keyType string
	this.KeyType = &keyType
	var certExtensions CertExtensions1
	this.CertExtensions = &certExtensions
	var addToTrustedStore bool
	this.AddToTrustedStore = &addToTrustedStore
	return &this
}

// GetTos returns the Tos field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetTos() bool {
	if o == nil || IsNil(o.Tos) {
		var ret bool
		return ret
	}
	return *o.Tos
}

// GetTosOk returns a tuple with the Tos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetTosOk() (*bool, bool) {
	if o == nil || IsNil(o.Tos) {
		return nil, false
	}
	return o.Tos, true
}

// HasTos returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasTos() bool {
	if o != nil && !IsNil(o.Tos) {
		return true
	}

	return false
}

// SetTos gets a reference to the given bool and assigns it to the Tos field.
func (o *CertificateauthorityCreate0) SetTos(v bool) {
	o.Tos = &v
}

// GetCsrId returns the CsrId field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetCsrId() int32 {
	if o == nil || IsNil(o.CsrId) {
		var ret int32
		return ret
	}
	return *o.CsrId
}

// GetCsrIdOk returns a tuple with the CsrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetCsrIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CsrId) {
		return nil, false
	}
	return o.CsrId, true
}

// HasCsrId returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCsrId() bool {
	if o != nil && !IsNil(o.CsrId) {
		return true
	}

	return false
}

// SetCsrId gets a reference to the given int32 and assigns it to the CsrId field.
func (o *CertificateauthorityCreate0) SetCsrId(v int32) {
	o.CsrId = &v
}

// GetSignedby returns the Signedby field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetSignedby() int32 {
	if o == nil || IsNil(o.Signedby) {
		var ret int32
		return ret
	}
	return *o.Signedby
}

// GetSignedbyOk returns a tuple with the Signedby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetSignedbyOk() (*int32, bool) {
	if o == nil || IsNil(o.Signedby) {
		return nil, false
	}
	return o.Signedby, true
}

// HasSignedby returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasSignedby() bool {
	if o != nil && !IsNil(o.Signedby) {
		return true
	}

	return false
}

// SetSignedby gets a reference to the given int32 and assigns it to the Signedby field.
func (o *CertificateauthorityCreate0) SetSignedby(v int32) {
	o.Signedby = &v
}

// GetKeyLength returns the KeyLength field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetKeyLength() int32 {
	if o == nil || IsNil(o.KeyLength) {
		var ret int32
		return ret
	}
	return *o.KeyLength
}

// GetKeyLengthOk returns a tuple with the KeyLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetKeyLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.KeyLength) {
		return nil, false
	}
	return o.KeyLength, true
}

// HasKeyLength returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasKeyLength() bool {
	if o != nil && !IsNil(o.KeyLength) {
		return true
	}

	return false
}

// SetKeyLength gets a reference to the given int32 and assigns it to the KeyLength field.
func (o *CertificateauthorityCreate0) SetKeyLength(v int32) {
	o.KeyLength = &v
}

// GetRenewDays returns the RenewDays field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetRenewDays() int32 {
	if o == nil || IsNil(o.RenewDays) {
		var ret int32
		return ret
	}
	return *o.RenewDays
}

// GetRenewDaysOk returns a tuple with the RenewDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetRenewDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.RenewDays) {
		return nil, false
	}
	return o.RenewDays, true
}

// HasRenewDays returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasRenewDays() bool {
	if o != nil && !IsNil(o.RenewDays) {
		return true
	}

	return false
}

// SetRenewDays gets a reference to the given int32 and assigns it to the RenewDays field.
func (o *CertificateauthorityCreate0) SetRenewDays(v int32) {
	o.RenewDays = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *CertificateauthorityCreate0) SetType(v int32) {
	o.Type = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetLifetime() int32 {
	if o == nil || IsNil(o.Lifetime) {
		var ret int32
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasLifetime() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given int32 and assigns it to the Lifetime field.
func (o *CertificateauthorityCreate0) SetLifetime(v int32) {
	o.Lifetime = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetSerial() int32 {
	if o == nil || IsNil(o.Serial) {
		var ret int32
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetSerialOk() (*int32, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given int32 and assigns it to the Serial field.
func (o *CertificateauthorityCreate0) SetSerial(v int32) {
	o.Serial = &v
}

// GetAcmeDirectoryUri returns the AcmeDirectoryUri field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetAcmeDirectoryUri() string {
	if o == nil || IsNil(o.AcmeDirectoryUri) {
		var ret string
		return ret
	}
	return *o.AcmeDirectoryUri
}

// GetAcmeDirectoryUriOk returns a tuple with the AcmeDirectoryUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetAcmeDirectoryUriOk() (*string, bool) {
	if o == nil || IsNil(o.AcmeDirectoryUri) {
		return nil, false
	}
	return o.AcmeDirectoryUri, true
}

// HasAcmeDirectoryUri returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasAcmeDirectoryUri() bool {
	if o != nil && !IsNil(o.AcmeDirectoryUri) {
		return true
	}

	return false
}

// SetAcmeDirectoryUri gets a reference to the given string and assigns it to the AcmeDirectoryUri field.
func (o *CertificateauthorityCreate0) SetAcmeDirectoryUri(v string) {
	o.AcmeDirectoryUri = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CertificateauthorityCreate0) SetCertificate(v string) {
	o.Certificate = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CertificateauthorityCreate0) SetCity(v string) {
	o.City = &v
}

// GetCommon returns the Common field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateauthorityCreate0) GetCommon() string {
	if o == nil || IsNil(o.Common.Get()) {
		var ret string
		return ret
	}
	return *o.Common.Get()
}

// GetCommonOk returns a tuple with the Common field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateauthorityCreate0) GetCommonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Common.Get(), o.Common.IsSet()
}

// HasCommon returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCommon() bool {
	if o != nil && o.Common.IsSet() {
		return true
	}

	return false
}

// SetCommon gets a reference to the given NullableString and assigns it to the Common field.
func (o *CertificateauthorityCreate0) SetCommon(v string) {
	o.Common.Set(&v)
}

// SetCommonNil sets the value for Common to be an explicit nil
func (o *CertificateauthorityCreate0) SetCommonNil() {
	o.Common.Set(nil)
}

// UnsetCommon ensures that no value is present for Common, not even an explicit nil
func (o *CertificateauthorityCreate0) UnsetCommon() {
	o.Common.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CertificateauthorityCreate0) SetCountry(v string) {
	o.Country = &v
}

// GetCSR returns the CSR field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetCSR() string {
	if o == nil || IsNil(o.CSR) {
		var ret string
		return ret
	}
	return *o.CSR
}

// GetCSROk returns a tuple with the CSR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetCSROk() (*string, bool) {
	if o == nil || IsNil(o.CSR) {
		return nil, false
	}
	return o.CSR, true
}

// HasCSR returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCSR() bool {
	if o != nil && !IsNil(o.CSR) {
		return true
	}

	return false
}

// SetCSR gets a reference to the given string and assigns it to the CSR field.
func (o *CertificateauthorityCreate0) SetCSR(v string) {
	o.CSR = &v
}

// GetEcCurve returns the EcCurve field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetEcCurve() string {
	if o == nil || IsNil(o.EcCurve) {
		var ret string
		return ret
	}
	return *o.EcCurve
}

// GetEcCurveOk returns a tuple with the EcCurve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetEcCurveOk() (*string, bool) {
	if o == nil || IsNil(o.EcCurve) {
		return nil, false
	}
	return o.EcCurve, true
}

// HasEcCurve returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasEcCurve() bool {
	if o != nil && !IsNil(o.EcCurve) {
		return true
	}

	return false
}

// SetEcCurve gets a reference to the given string and assigns it to the EcCurve field.
func (o *CertificateauthorityCreate0) SetEcCurve(v string) {
	o.EcCurve = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CertificateauthorityCreate0) SetEmail(v string) {
	o.Email = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetKeyType() string {
	if o == nil || IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *CertificateauthorityCreate0) SetKeyType(v string) {
	o.KeyType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CertificateauthorityCreate0) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *CertificateauthorityCreate0) SetOrganization(v string) {
	o.Organization = &v
}

// GetOrganizationalUnit returns the OrganizationalUnit field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetOrganizationalUnit() string {
	if o == nil || IsNil(o.OrganizationalUnit) {
		var ret string
		return ret
	}
	return *o.OrganizationalUnit
}

// GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetOrganizationalUnitOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationalUnit) {
		return nil, false
	}
	return o.OrganizationalUnit, true
}

// HasOrganizationalUnit returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasOrganizationalUnit() bool {
	if o != nil && !IsNil(o.OrganizationalUnit) {
		return true
	}

	return false
}

// SetOrganizationalUnit gets a reference to the given string and assigns it to the OrganizationalUnit field.
func (o *CertificateauthorityCreate0) SetOrganizationalUnit(v string) {
	o.OrganizationalUnit = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *CertificateauthorityCreate0) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetPrivatekey returns the Privatekey field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetPrivatekey() string {
	if o == nil || IsNil(o.Privatekey) {
		var ret string
		return ret
	}
	return *o.Privatekey
}

// GetPrivatekeyOk returns a tuple with the Privatekey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetPrivatekeyOk() (*string, bool) {
	if o == nil || IsNil(o.Privatekey) {
		return nil, false
	}
	return o.Privatekey, true
}

// HasPrivatekey returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasPrivatekey() bool {
	if o != nil && !IsNil(o.Privatekey) {
		return true
	}

	return false
}

// SetPrivatekey gets a reference to the given string and assigns it to the Privatekey field.
func (o *CertificateauthorityCreate0) SetPrivatekey(v string) {
	o.Privatekey = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CertificateauthorityCreate0) SetState(v string) {
	o.State = &v
}

// GetCreateType returns the CreateType field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetCreateType() string {
	if o == nil || IsNil(o.CreateType) {
		var ret string
		return ret
	}
	return *o.CreateType
}

// GetCreateTypeOk returns a tuple with the CreateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetCreateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateType) {
		return nil, false
	}
	return o.CreateType, true
}

// HasCreateType returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCreateType() bool {
	if o != nil && !IsNil(o.CreateType) {
		return true
	}

	return false
}

// SetCreateType gets a reference to the given string and assigns it to the CreateType field.
func (o *CertificateauthorityCreate0) SetCreateType(v string) {
	o.CreateType = &v
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetDigestAlgorithm() string {
	if o == nil || IsNil(o.DigestAlgorithm) {
		var ret string
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetDigestAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.DigestAlgorithm) {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasDigestAlgorithm() bool {
	if o != nil && !IsNil(o.DigestAlgorithm) {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given string and assigns it to the DigestAlgorithm field.
func (o *CertificateauthorityCreate0) SetDigestAlgorithm(v string) {
	o.DigestAlgorithm = &v
}

// GetSan returns the San field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetSan() []string {
	if o == nil || IsNil(o.San) {
		var ret []string
		return ret
	}
	return o.San
}

// GetSanOk returns a tuple with the San field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetSanOk() ([]string, bool) {
	if o == nil || IsNil(o.San) {
		return nil, false
	}
	return o.San, true
}

// HasSan returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasSan() bool {
	if o != nil && !IsNil(o.San) {
		return true
	}

	return false
}

// SetSan gets a reference to the given []string and assigns it to the San field.
func (o *CertificateauthorityCreate0) SetSan(v []string) {
	o.San = v
}

// GetCertExtensions returns the CertExtensions field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetCertExtensions() CertExtensions1 {
	if o == nil || IsNil(o.CertExtensions) {
		var ret CertExtensions1
		return ret
	}
	return *o.CertExtensions
}

// GetCertExtensionsOk returns a tuple with the CertExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetCertExtensionsOk() (*CertExtensions1, bool) {
	if o == nil || IsNil(o.CertExtensions) {
		return nil, false
	}
	return o.CertExtensions, true
}

// HasCertExtensions returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasCertExtensions() bool {
	if o != nil && !IsNil(o.CertExtensions) {
		return true
	}

	return false
}

// SetCertExtensions gets a reference to the given CertExtensions1 and assigns it to the CertExtensions field.
func (o *CertificateauthorityCreate0) SetCertExtensions(v CertExtensions1) {
	o.CertExtensions = &v
}

// GetAddToTrustedStore returns the AddToTrustedStore field value if set, zero value otherwise.
func (o *CertificateauthorityCreate0) GetAddToTrustedStore() bool {
	if o == nil || IsNil(o.AddToTrustedStore) {
		var ret bool
		return ret
	}
	return *o.AddToTrustedStore
}

// GetAddToTrustedStoreOk returns a tuple with the AddToTrustedStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateauthorityCreate0) GetAddToTrustedStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.AddToTrustedStore) {
		return nil, false
	}
	return o.AddToTrustedStore, true
}

// HasAddToTrustedStore returns a boolean if a field has been set.
func (o *CertificateauthorityCreate0) HasAddToTrustedStore() bool {
	if o != nil && !IsNil(o.AddToTrustedStore) {
		return true
	}

	return false
}

// SetAddToTrustedStore gets a reference to the given bool and assigns it to the AddToTrustedStore field.
func (o *CertificateauthorityCreate0) SetAddToTrustedStore(v bool) {
	o.AddToTrustedStore = &v
}

func (o CertificateauthorityCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateauthorityCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tos) {
		toSerialize["tos"] = o.Tos
	}
	if !IsNil(o.CsrId) {
		toSerialize["csr_id"] = o.CsrId
	}
	if !IsNil(o.Signedby) {
		toSerialize["signedby"] = o.Signedby
	}
	if !IsNil(o.KeyLength) {
		toSerialize["key_length"] = o.KeyLength
	}
	if !IsNil(o.RenewDays) {
		toSerialize["renew_days"] = o.RenewDays
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Lifetime) {
		toSerialize["lifetime"] = o.Lifetime
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.AcmeDirectoryUri) {
		toSerialize["acme_directory_uri"] = o.AcmeDirectoryUri
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if o.Common.IsSet() {
		toSerialize["common"] = o.Common.Get()
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CSR) {
		toSerialize["CSR"] = o.CSR
	}
	if !IsNil(o.EcCurve) {
		toSerialize["ec_curve"] = o.EcCurve
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.KeyType) {
		toSerialize["key_type"] = o.KeyType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.OrganizationalUnit) {
		toSerialize["organizational_unit"] = o.OrganizationalUnit
	}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if !IsNil(o.Privatekey) {
		toSerialize["privatekey"] = o.Privatekey
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CreateType) {
		toSerialize["create_type"] = o.CreateType
	}
	if !IsNil(o.DigestAlgorithm) {
		toSerialize["digest_algorithm"] = o.DigestAlgorithm
	}
	if !IsNil(o.San) {
		toSerialize["san"] = o.San
	}
	if !IsNil(o.CertExtensions) {
		toSerialize["cert_extensions"] = o.CertExtensions
	}
	if !IsNil(o.AddToTrustedStore) {
		toSerialize["add_to_trusted_store"] = o.AddToTrustedStore
	}
	return toSerialize, nil
}

type NullableCertificateauthorityCreate0 struct {
	value *CertificateauthorityCreate0
	isSet bool
}

func (v NullableCertificateauthorityCreate0) Get() *CertificateauthorityCreate0 {
	return v.value
}

func (v *NullableCertificateauthorityCreate0) Set(val *CertificateauthorityCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateauthorityCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateauthorityCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateauthorityCreate0(val *CertificateauthorityCreate0) *NullableCertificateauthorityCreate0 {
	return &NullableCertificateauthorityCreate0{value: val, isSet: true}
}

func (v NullableCertificateauthorityCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateauthorityCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
