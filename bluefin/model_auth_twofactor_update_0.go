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

// checks if the AuthTwofactorUpdate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTwofactorUpdate0{}

// AuthTwofactorUpdate0 struct for AuthTwofactorUpdate0
type AuthTwofactorUpdate0 struct {
	Enabled *bool `json:"enabled,omitempty"`
	// `otp_digits` represents number of allowed digits in the OTP.
	OtpDigits *int32 `json:"otp_digits,omitempty"`
	// `window` extends the validity to `window` many counter ticks before and after the current one.
	Window   *int32    `json:"window,omitempty"`
	Interval *int32    `json:"interval,omitempty"`
	Services *Services `json:"services,omitempty"`
}

// NewAuthTwofactorUpdate0 instantiates a new AuthTwofactorUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTwofactorUpdate0() *AuthTwofactorUpdate0 {
	this := AuthTwofactorUpdate0{}
	var services Services
	this.Services = &services
	return &this
}

// NewAuthTwofactorUpdate0WithDefaults instantiates a new AuthTwofactorUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTwofactorUpdate0WithDefaults() *AuthTwofactorUpdate0 {
	this := AuthTwofactorUpdate0{}
	var services Services
	this.Services = &services
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AuthTwofactorUpdate0) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOtpDigits returns the OtpDigits field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetOtpDigits() int32 {
	if o == nil || IsNil(o.OtpDigits) {
		var ret int32
		return ret
	}
	return *o.OtpDigits
}

// GetOtpDigitsOk returns a tuple with the OtpDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetOtpDigitsOk() (*int32, bool) {
	if o == nil || IsNil(o.OtpDigits) {
		return nil, false
	}
	return o.OtpDigits, true
}

// HasOtpDigits returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasOtpDigits() bool {
	if o != nil && !IsNil(o.OtpDigits) {
		return true
	}

	return false
}

// SetOtpDigits gets a reference to the given int32 and assigns it to the OtpDigits field.
func (o *AuthTwofactorUpdate0) SetOtpDigits(v int32) {
	o.OtpDigits = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetWindow() int32 {
	if o == nil || IsNil(o.Window) {
		var ret int32
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.Window) {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasWindow() bool {
	if o != nil && !IsNil(o.Window) {
		return true
	}

	return false
}

// SetWindow gets a reference to the given int32 and assigns it to the Window field.
func (o *AuthTwofactorUpdate0) SetWindow(v int32) {
	o.Window = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *AuthTwofactorUpdate0) SetInterval(v int32) {
	o.Interval = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetServices() Services {
	if o == nil || IsNil(o.Services) {
		var ret Services
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetServicesOk() (*Services, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given Services and assigns it to the Services field.
func (o *AuthTwofactorUpdate0) SetServices(v Services) {
	o.Services = &v
}

func (o AuthTwofactorUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTwofactorUpdate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.OtpDigits) {
		toSerialize["otp_digits"] = o.OtpDigits
	}
	if !IsNil(o.Window) {
		toSerialize["window"] = o.Window
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	return toSerialize, nil
}

type NullableAuthTwofactorUpdate0 struct {
	value *AuthTwofactorUpdate0
	isSet bool
}

func (v NullableAuthTwofactorUpdate0) Get() *AuthTwofactorUpdate0 {
	return v.value
}

func (v *NullableAuthTwofactorUpdate0) Set(val *AuthTwofactorUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTwofactorUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTwofactorUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTwofactorUpdate0(val *AuthTwofactorUpdate0) *NullableAuthTwofactorUpdate0 {
	return &NullableAuthTwofactorUpdate0{value: val, isSet: true}
}

func (v NullableAuthTwofactorUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTwofactorUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}