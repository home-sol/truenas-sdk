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

// checks if the LdapUpdate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapUpdate0{}

// LdapUpdate0 struct for LdapUpdate0
type LdapUpdate0 struct {
	// `hostname` list of ip addresses or hostnames of LDAP servers with which to communicate in order of preference. Failover only occurs if the current LDAP server is unresponsive.
	Hostname []interface{} `json:"hostname,omitempty"`
	// `basedn` specifies the default base DN to use when performing ldap operations. The base must be specified as a Distinguished Name in LDAP format.
	Basedn *string `json:"basedn,omitempty"`
	// `binddn` specifies the default bind DN to use when performing ldap operations. The bind DN must be specified as a Distinguished Name in LDAP format. `kerberos_principal` kerberos principal to use for SASL GSSAPI authentication to the remote server. If `kerberos_realm` is specified without a keytab, then the `binddn` and `bindpw` are used to perform to obtain the ticket necessary for GSSAPI authentication.
	Binddn *string `json:"binddn,omitempty"`
	// `kerberos_principal` kerberos principal to use for SASL GSSAPI authentication to the remote server. If `kerberos_realm` is specified without a keytab, then the `binddn` and `bindpw` are used to perform to obtain the ticket necessary for GSSAPI authentication.
	Bindpw *string `json:"bindpw,omitempty"`
	// `anonbind` use anonymous authentication.
	Anonbind *bool `json:"anonbind,omitempty"`
	// `ssl` establish SSL/TLS-protected connections to the LDAP server(s). GSSAPI signing is disabled on SSL/TLS-protected connections if kerberos authentication is used.
	Ssl *string `json:"ssl,omitempty"`
	// `certificate` LDAPs client certificate to be used for certificate- based authentication.
	Certificate NullableInt32 `json:"certificate,omitempty"`
	// `validate_certificates` specifies whether to perform checks on server certificates in a TLS session. If enabled, TLS_REQCERT demand is set. The server certificate is requested. If no certificate is provided or if a bad certificate is provided, the session is immediately terminated. If disabled, TLS_REQCERT allow is set. The server certificate is requested, but all errors are ignored.
	ValidateCertificates *bool `json:"validate_certificates,omitempty"`
	DisableFreenasCache  *bool `json:"disable_freenas_cache,omitempty"`
	// `timeout` specifies  a  timeout  (in  seconds) after which calls to synchronous LDAP APIs will abort if no response is received.
	Timeout *int32 `json:"timeout,omitempty"`
	// `dns_timeout` specifies the timeout (in seconds) after which the poll(2)/select(2) following a connect(2) returns in case of no activity for openldap. For nslcd this specifies the time limit (in seconds) to use when connecting to the directory server. This directly impacts the length of time that the LDAP service tries before failing over to a secondary LDAP URI.
	DnsTimeout *int32 `json:"dns_timeout,omitempty"`
	// `kerberos_realm` in which the server is located. This parameter is only required for SASL GSSAPI authentication to the remote LDAP server. `kerberos_principal` kerberos principal to use for SASL GSSAPI authentication to the remote server. If `kerberos_realm` is specified without a keytab, then the `binddn` and `bindpw` are used to perform to obtain the ticket necessary for GSSAPI authentication.
	KerberosRealm NullableInt32 `json:"kerberos_realm,omitempty"`
	// `kerberos_principal` kerberos principal to use for SASL GSSAPI authentication to the remote server. If `kerberos_realm` is specified without a keytab, then the `binddn` and `bindpw` are used to perform to obtain the ticket necessary for GSSAPI authentication.
	KerberosPrincipal   *string `json:"kerberos_principal,omitempty"`
	HasSambaSchema      *bool   `json:"has_samba_schema,omitempty"`
	AuxiliaryParameters *string `json:"auxiliary_parameters,omitempty"`
	Schema              *string `json:"schema,omitempty"`
	Enable              *bool   `json:"enable,omitempty"`
}

// NewLdapUpdate0 instantiates a new LdapUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapUpdate0() *LdapUpdate0 {
	this := LdapUpdate0{}
	var anonbind bool
	this.Anonbind = &anonbind
	var ssl string
	this.Ssl = &ssl
	var validateCertificates bool
	this.ValidateCertificates = &validateCertificates
	var timeout int32
	this.Timeout = &timeout
	var dnsTimeout int32
	this.DnsTimeout = &dnsTimeout
	var hasSambaSchema bool
	this.HasSambaSchema = &hasSambaSchema
	var auxiliaryParameters string
	this.AuxiliaryParameters = &auxiliaryParameters
	var schema string
	this.Schema = &schema
	return &this
}

// NewLdapUpdate0WithDefaults instantiates a new LdapUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapUpdate0WithDefaults() *LdapUpdate0 {
	this := LdapUpdate0{}
	var anonbind bool
	this.Anonbind = &anonbind
	var ssl string
	this.Ssl = &ssl
	var validateCertificates bool
	this.ValidateCertificates = &validateCertificates
	var timeout int32
	this.Timeout = &timeout
	var dnsTimeout int32
	this.DnsTimeout = &dnsTimeout
	var hasSambaSchema bool
	this.HasSambaSchema = &hasSambaSchema
	var auxiliaryParameters string
	this.AuxiliaryParameters = &auxiliaryParameters
	var schema string
	this.Schema = &schema
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *LdapUpdate0) GetHostname() []interface{} {
	if o == nil || IsNil(o.Hostname) {
		var ret []interface{}
		return ret
	}
	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetHostnameOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *LdapUpdate0) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given []interface{} and assigns it to the Hostname field.
func (o *LdapUpdate0) SetHostname(v []interface{}) {
	o.Hostname = v
}

// GetBasedn returns the Basedn field value if set, zero value otherwise.
func (o *LdapUpdate0) GetBasedn() string {
	if o == nil || IsNil(o.Basedn) {
		var ret string
		return ret
	}
	return *o.Basedn
}

// GetBasednOk returns a tuple with the Basedn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetBasednOk() (*string, bool) {
	if o == nil || IsNil(o.Basedn) {
		return nil, false
	}
	return o.Basedn, true
}

// HasBasedn returns a boolean if a field has been set.
func (o *LdapUpdate0) HasBasedn() bool {
	if o != nil && !IsNil(o.Basedn) {
		return true
	}

	return false
}

// SetBasedn gets a reference to the given string and assigns it to the Basedn field.
func (o *LdapUpdate0) SetBasedn(v string) {
	o.Basedn = &v
}

// GetBinddn returns the Binddn field value if set, zero value otherwise.
func (o *LdapUpdate0) GetBinddn() string {
	if o == nil || IsNil(o.Binddn) {
		var ret string
		return ret
	}
	return *o.Binddn
}

// GetBinddnOk returns a tuple with the Binddn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetBinddnOk() (*string, bool) {
	if o == nil || IsNil(o.Binddn) {
		return nil, false
	}
	return o.Binddn, true
}

// HasBinddn returns a boolean if a field has been set.
func (o *LdapUpdate0) HasBinddn() bool {
	if o != nil && !IsNil(o.Binddn) {
		return true
	}

	return false
}

// SetBinddn gets a reference to the given string and assigns it to the Binddn field.
func (o *LdapUpdate0) SetBinddn(v string) {
	o.Binddn = &v
}

// GetBindpw returns the Bindpw field value if set, zero value otherwise.
func (o *LdapUpdate0) GetBindpw() string {
	if o == nil || IsNil(o.Bindpw) {
		var ret string
		return ret
	}
	return *o.Bindpw
}

// GetBindpwOk returns a tuple with the Bindpw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetBindpwOk() (*string, bool) {
	if o == nil || IsNil(o.Bindpw) {
		return nil, false
	}
	return o.Bindpw, true
}

// HasBindpw returns a boolean if a field has been set.
func (o *LdapUpdate0) HasBindpw() bool {
	if o != nil && !IsNil(o.Bindpw) {
		return true
	}

	return false
}

// SetBindpw gets a reference to the given string and assigns it to the Bindpw field.
func (o *LdapUpdate0) SetBindpw(v string) {
	o.Bindpw = &v
}

// GetAnonbind returns the Anonbind field value if set, zero value otherwise.
func (o *LdapUpdate0) GetAnonbind() bool {
	if o == nil || IsNil(o.Anonbind) {
		var ret bool
		return ret
	}
	return *o.Anonbind
}

// GetAnonbindOk returns a tuple with the Anonbind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetAnonbindOk() (*bool, bool) {
	if o == nil || IsNil(o.Anonbind) {
		return nil, false
	}
	return o.Anonbind, true
}

// HasAnonbind returns a boolean if a field has been set.
func (o *LdapUpdate0) HasAnonbind() bool {
	if o != nil && !IsNil(o.Anonbind) {
		return true
	}

	return false
}

// SetAnonbind gets a reference to the given bool and assigns it to the Anonbind field.
func (o *LdapUpdate0) SetAnonbind(v bool) {
	o.Anonbind = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *LdapUpdate0) GetSsl() string {
	if o == nil || IsNil(o.Ssl) {
		var ret string
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetSslOk() (*string, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *LdapUpdate0) HasSsl() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given string and assigns it to the Ssl field.
func (o *LdapUpdate0) SetSsl(v string) {
	o.Ssl = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LdapUpdate0) GetCertificate() int32 {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret int32
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LdapUpdate0) GetCertificateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *LdapUpdate0) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableInt32 and assigns it to the Certificate field.
func (o *LdapUpdate0) SetCertificate(v int32) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *LdapUpdate0) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *LdapUpdate0) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetValidateCertificates returns the ValidateCertificates field value if set, zero value otherwise.
func (o *LdapUpdate0) GetValidateCertificates() bool {
	if o == nil || IsNil(o.ValidateCertificates) {
		var ret bool
		return ret
	}
	return *o.ValidateCertificates
}

// GetValidateCertificatesOk returns a tuple with the ValidateCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetValidateCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateCertificates) {
		return nil, false
	}
	return o.ValidateCertificates, true
}

// HasValidateCertificates returns a boolean if a field has been set.
func (o *LdapUpdate0) HasValidateCertificates() bool {
	if o != nil && !IsNil(o.ValidateCertificates) {
		return true
	}

	return false
}

// SetValidateCertificates gets a reference to the given bool and assigns it to the ValidateCertificates field.
func (o *LdapUpdate0) SetValidateCertificates(v bool) {
	o.ValidateCertificates = &v
}

// GetDisableFreenasCache returns the DisableFreenasCache field value if set, zero value otherwise.
func (o *LdapUpdate0) GetDisableFreenasCache() bool {
	if o == nil || IsNil(o.DisableFreenasCache) {
		var ret bool
		return ret
	}
	return *o.DisableFreenasCache
}

// GetDisableFreenasCacheOk returns a tuple with the DisableFreenasCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetDisableFreenasCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableFreenasCache) {
		return nil, false
	}
	return o.DisableFreenasCache, true
}

// HasDisableFreenasCache returns a boolean if a field has been set.
func (o *LdapUpdate0) HasDisableFreenasCache() bool {
	if o != nil && !IsNil(o.DisableFreenasCache) {
		return true
	}

	return false
}

// SetDisableFreenasCache gets a reference to the given bool and assigns it to the DisableFreenasCache field.
func (o *LdapUpdate0) SetDisableFreenasCache(v bool) {
	o.DisableFreenasCache = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *LdapUpdate0) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *LdapUpdate0) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *LdapUpdate0) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetDnsTimeout returns the DnsTimeout field value if set, zero value otherwise.
func (o *LdapUpdate0) GetDnsTimeout() int32 {
	if o == nil || IsNil(o.DnsTimeout) {
		var ret int32
		return ret
	}
	return *o.DnsTimeout
}

// GetDnsTimeoutOk returns a tuple with the DnsTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetDnsTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.DnsTimeout) {
		return nil, false
	}
	return o.DnsTimeout, true
}

// HasDnsTimeout returns a boolean if a field has been set.
func (o *LdapUpdate0) HasDnsTimeout() bool {
	if o != nil && !IsNil(o.DnsTimeout) {
		return true
	}

	return false
}

// SetDnsTimeout gets a reference to the given int32 and assigns it to the DnsTimeout field.
func (o *LdapUpdate0) SetDnsTimeout(v int32) {
	o.DnsTimeout = &v
}

// GetKerberosRealm returns the KerberosRealm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LdapUpdate0) GetKerberosRealm() int32 {
	if o == nil || IsNil(o.KerberosRealm.Get()) {
		var ret int32
		return ret
	}
	return *o.KerberosRealm.Get()
}

// GetKerberosRealmOk returns a tuple with the KerberosRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LdapUpdate0) GetKerberosRealmOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.KerberosRealm.Get(), o.KerberosRealm.IsSet()
}

// HasKerberosRealm returns a boolean if a field has been set.
func (o *LdapUpdate0) HasKerberosRealm() bool {
	if o != nil && o.KerberosRealm.IsSet() {
		return true
	}

	return false
}

// SetKerberosRealm gets a reference to the given NullableInt32 and assigns it to the KerberosRealm field.
func (o *LdapUpdate0) SetKerberosRealm(v int32) {
	o.KerberosRealm.Set(&v)
}

// SetKerberosRealmNil sets the value for KerberosRealm to be an explicit nil
func (o *LdapUpdate0) SetKerberosRealmNil() {
	o.KerberosRealm.Set(nil)
}

// UnsetKerberosRealm ensures that no value is present for KerberosRealm, not even an explicit nil
func (o *LdapUpdate0) UnsetKerberosRealm() {
	o.KerberosRealm.Unset()
}

// GetKerberosPrincipal returns the KerberosPrincipal field value if set, zero value otherwise.
func (o *LdapUpdate0) GetKerberosPrincipal() string {
	if o == nil || IsNil(o.KerberosPrincipal) {
		var ret string
		return ret
	}
	return *o.KerberosPrincipal
}

// GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetKerberosPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.KerberosPrincipal) {
		return nil, false
	}
	return o.KerberosPrincipal, true
}

// HasKerberosPrincipal returns a boolean if a field has been set.
func (o *LdapUpdate0) HasKerberosPrincipal() bool {
	if o != nil && !IsNil(o.KerberosPrincipal) {
		return true
	}

	return false
}

// SetKerberosPrincipal gets a reference to the given string and assigns it to the KerberosPrincipal field.
func (o *LdapUpdate0) SetKerberosPrincipal(v string) {
	o.KerberosPrincipal = &v
}

// GetHasSambaSchema returns the HasSambaSchema field value if set, zero value otherwise.
func (o *LdapUpdate0) GetHasSambaSchema() bool {
	if o == nil || IsNil(o.HasSambaSchema) {
		var ret bool
		return ret
	}
	return *o.HasSambaSchema
}

// GetHasSambaSchemaOk returns a tuple with the HasSambaSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetHasSambaSchemaOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSambaSchema) {
		return nil, false
	}
	return o.HasSambaSchema, true
}

// HasHasSambaSchema returns a boolean if a field has been set.
func (o *LdapUpdate0) HasHasSambaSchema() bool {
	if o != nil && !IsNil(o.HasSambaSchema) {
		return true
	}

	return false
}

// SetHasSambaSchema gets a reference to the given bool and assigns it to the HasSambaSchema field.
func (o *LdapUpdate0) SetHasSambaSchema(v bool) {
	o.HasSambaSchema = &v
}

// GetAuxiliaryParameters returns the AuxiliaryParameters field value if set, zero value otherwise.
func (o *LdapUpdate0) GetAuxiliaryParameters() string {
	if o == nil || IsNil(o.AuxiliaryParameters) {
		var ret string
		return ret
	}
	return *o.AuxiliaryParameters
}

// GetAuxiliaryParametersOk returns a tuple with the AuxiliaryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetAuxiliaryParametersOk() (*string, bool) {
	if o == nil || IsNil(o.AuxiliaryParameters) {
		return nil, false
	}
	return o.AuxiliaryParameters, true
}

// HasAuxiliaryParameters returns a boolean if a field has been set.
func (o *LdapUpdate0) HasAuxiliaryParameters() bool {
	if o != nil && !IsNil(o.AuxiliaryParameters) {
		return true
	}

	return false
}

// SetAuxiliaryParameters gets a reference to the given string and assigns it to the AuxiliaryParameters field.
func (o *LdapUpdate0) SetAuxiliaryParameters(v string) {
	o.AuxiliaryParameters = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *LdapUpdate0) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *LdapUpdate0) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *LdapUpdate0) SetSchema(v string) {
	o.Schema = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *LdapUpdate0) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUpdate0) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *LdapUpdate0) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *LdapUpdate0) SetEnable(v bool) {
	o.Enable = &v
}

func (o LdapUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapUpdate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Basedn) {
		toSerialize["basedn"] = o.Basedn
	}
	if !IsNil(o.Binddn) {
		toSerialize["binddn"] = o.Binddn
	}
	if !IsNil(o.Bindpw) {
		toSerialize["bindpw"] = o.Bindpw
	}
	if !IsNil(o.Anonbind) {
		toSerialize["anonbind"] = o.Anonbind
	}
	if !IsNil(o.Ssl) {
		toSerialize["ssl"] = o.Ssl
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if !IsNil(o.ValidateCertificates) {
		toSerialize["validate_certificates"] = o.ValidateCertificates
	}
	if !IsNil(o.DisableFreenasCache) {
		toSerialize["disable_freenas_cache"] = o.DisableFreenasCache
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.DnsTimeout) {
		toSerialize["dns_timeout"] = o.DnsTimeout
	}
	if o.KerberosRealm.IsSet() {
		toSerialize["kerberos_realm"] = o.KerberosRealm.Get()
	}
	if !IsNil(o.KerberosPrincipal) {
		toSerialize["kerberos_principal"] = o.KerberosPrincipal
	}
	if !IsNil(o.HasSambaSchema) {
		toSerialize["has_samba_schema"] = o.HasSambaSchema
	}
	if !IsNil(o.AuxiliaryParameters) {
		toSerialize["auxiliary_parameters"] = o.AuxiliaryParameters
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	return toSerialize, nil
}

type NullableLdapUpdate0 struct {
	value *LdapUpdate0
	isSet bool
}

func (v NullableLdapUpdate0) Get() *LdapUpdate0 {
	return v.value
}

func (v *NullableLdapUpdate0) Set(val *LdapUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapUpdate0(val *LdapUpdate0) *NullableLdapUpdate0 {
	return &NullableLdapUpdate0{value: val, isSet: true}
}

func (v NullableLdapUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
