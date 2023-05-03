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

// checks if the DyndnsUpdate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DyndnsUpdate0{}

// DyndnsUpdate0 struct for DyndnsUpdate0
type DyndnsUpdate0 struct {
	Provider         *string  `json:"provider,omitempty"`
	CheckipSsl       *bool    `json:"checkip_ssl,omitempty"`
	CheckipServer    *string  `json:"checkip_server,omitempty"`
	CheckipPath      *string  `json:"checkip_path,omitempty"`
	Ssl              *bool    `json:"ssl,omitempty"`
	CustomDdnsServer *string  `json:"custom_ddns_server,omitempty"`
	CustomDdnsPath   *string  `json:"custom_ddns_path,omitempty"`
	Domain           []string `json:"domain,omitempty"`
	Username         *string  `json:"username,omitempty"`
	Password         *string  `json:"password,omitempty"`
	// `period` indicates how often the IP is checked in seconds.
	Period *int32 `json:"period,omitempty"`
}

// NewDyndnsUpdate0 instantiates a new DyndnsUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDyndnsUpdate0() *DyndnsUpdate0 {
	this := DyndnsUpdate0{}
	return &this
}

// NewDyndnsUpdate0WithDefaults instantiates a new DyndnsUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDyndnsUpdate0WithDefaults() *DyndnsUpdate0 {
	this := DyndnsUpdate0{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *DyndnsUpdate0) SetProvider(v string) {
	o.Provider = &v
}

// GetCheckipSsl returns the CheckipSsl field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetCheckipSsl() bool {
	if o == nil || IsNil(o.CheckipSsl) {
		var ret bool
		return ret
	}
	return *o.CheckipSsl
}

// GetCheckipSslOk returns a tuple with the CheckipSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetCheckipSslOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckipSsl) {
		return nil, false
	}
	return o.CheckipSsl, true
}

// HasCheckipSsl returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasCheckipSsl() bool {
	if o != nil && !IsNil(o.CheckipSsl) {
		return true
	}

	return false
}

// SetCheckipSsl gets a reference to the given bool and assigns it to the CheckipSsl field.
func (o *DyndnsUpdate0) SetCheckipSsl(v bool) {
	o.CheckipSsl = &v
}

// GetCheckipServer returns the CheckipServer field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetCheckipServer() string {
	if o == nil || IsNil(o.CheckipServer) {
		var ret string
		return ret
	}
	return *o.CheckipServer
}

// GetCheckipServerOk returns a tuple with the CheckipServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetCheckipServerOk() (*string, bool) {
	if o == nil || IsNil(o.CheckipServer) {
		return nil, false
	}
	return o.CheckipServer, true
}

// HasCheckipServer returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasCheckipServer() bool {
	if o != nil && !IsNil(o.CheckipServer) {
		return true
	}

	return false
}

// SetCheckipServer gets a reference to the given string and assigns it to the CheckipServer field.
func (o *DyndnsUpdate0) SetCheckipServer(v string) {
	o.CheckipServer = &v
}

// GetCheckipPath returns the CheckipPath field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetCheckipPath() string {
	if o == nil || IsNil(o.CheckipPath) {
		var ret string
		return ret
	}
	return *o.CheckipPath
}

// GetCheckipPathOk returns a tuple with the CheckipPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetCheckipPathOk() (*string, bool) {
	if o == nil || IsNil(o.CheckipPath) {
		return nil, false
	}
	return o.CheckipPath, true
}

// HasCheckipPath returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasCheckipPath() bool {
	if o != nil && !IsNil(o.CheckipPath) {
		return true
	}

	return false
}

// SetCheckipPath gets a reference to the given string and assigns it to the CheckipPath field.
func (o *DyndnsUpdate0) SetCheckipPath(v string) {
	o.CheckipPath = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetSsl() bool {
	if o == nil || IsNil(o.Ssl) {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetSslOk() (*bool, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasSsl() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *DyndnsUpdate0) SetSsl(v bool) {
	o.Ssl = &v
}

// GetCustomDdnsServer returns the CustomDdnsServer field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetCustomDdnsServer() string {
	if o == nil || IsNil(o.CustomDdnsServer) {
		var ret string
		return ret
	}
	return *o.CustomDdnsServer
}

// GetCustomDdnsServerOk returns a tuple with the CustomDdnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetCustomDdnsServerOk() (*string, bool) {
	if o == nil || IsNil(o.CustomDdnsServer) {
		return nil, false
	}
	return o.CustomDdnsServer, true
}

// HasCustomDdnsServer returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasCustomDdnsServer() bool {
	if o != nil && !IsNil(o.CustomDdnsServer) {
		return true
	}

	return false
}

// SetCustomDdnsServer gets a reference to the given string and assigns it to the CustomDdnsServer field.
func (o *DyndnsUpdate0) SetCustomDdnsServer(v string) {
	o.CustomDdnsServer = &v
}

// GetCustomDdnsPath returns the CustomDdnsPath field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetCustomDdnsPath() string {
	if o == nil || IsNil(o.CustomDdnsPath) {
		var ret string
		return ret
	}
	return *o.CustomDdnsPath
}

// GetCustomDdnsPathOk returns a tuple with the CustomDdnsPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetCustomDdnsPathOk() (*string, bool) {
	if o == nil || IsNil(o.CustomDdnsPath) {
		return nil, false
	}
	return o.CustomDdnsPath, true
}

// HasCustomDdnsPath returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasCustomDdnsPath() bool {
	if o != nil && !IsNil(o.CustomDdnsPath) {
		return true
	}

	return false
}

// SetCustomDdnsPath gets a reference to the given string and assigns it to the CustomDdnsPath field.
func (o *DyndnsUpdate0) SetCustomDdnsPath(v string) {
	o.CustomDdnsPath = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetDomain() []string {
	if o == nil || IsNil(o.Domain) {
		var ret []string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetDomainOk() ([]string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given []string and assigns it to the Domain field.
func (o *DyndnsUpdate0) SetDomain(v []string) {
	o.Domain = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DyndnsUpdate0) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DyndnsUpdate0) SetPassword(v string) {
	o.Password = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DyndnsUpdate0) GetPeriod() int32 {
	if o == nil || IsNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DyndnsUpdate0) GetPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DyndnsUpdate0) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *DyndnsUpdate0) SetPeriod(v int32) {
	o.Period = &v
}

func (o DyndnsUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DyndnsUpdate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.CheckipSsl) {
		toSerialize["checkip_ssl"] = o.CheckipSsl
	}
	if !IsNil(o.CheckipServer) {
		toSerialize["checkip_server"] = o.CheckipServer
	}
	if !IsNil(o.CheckipPath) {
		toSerialize["checkip_path"] = o.CheckipPath
	}
	if !IsNil(o.Ssl) {
		toSerialize["ssl"] = o.Ssl
	}
	if !IsNil(o.CustomDdnsServer) {
		toSerialize["custom_ddns_server"] = o.CustomDdnsServer
	}
	if !IsNil(o.CustomDdnsPath) {
		toSerialize["custom_ddns_path"] = o.CustomDdnsPath
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	return toSerialize, nil
}

type NullableDyndnsUpdate0 struct {
	value *DyndnsUpdate0
	isSet bool
}

func (v NullableDyndnsUpdate0) Get() *DyndnsUpdate0 {
	return v.value
}

func (v *NullableDyndnsUpdate0) Set(val *DyndnsUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableDyndnsUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableDyndnsUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDyndnsUpdate0(val *DyndnsUpdate0) *NullableDyndnsUpdate0 {
	return &NullableDyndnsUpdate0{value: val, isSet: true}
}

func (v NullableDyndnsUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDyndnsUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
