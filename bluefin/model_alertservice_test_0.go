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

// checks if the AlertserviceTest0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertserviceTest0{}

// AlertserviceTest0 struct for AlertserviceTest0
type AlertserviceTest0 struct {
	Name *string `json:"name,omitempty"`
	// Send a test alert using `type` of Alert Service.
	Type       *string                `json:"type,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Level      *string                `json:"level,omitempty"`
	Enabled    *bool                  `json:"enabled,omitempty"`
}

// NewAlertserviceTest0 instantiates a new AlertserviceTest0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertserviceTest0() *AlertserviceTest0 {
	this := AlertserviceTest0{}
	var enabled bool
	this.Enabled = &enabled
	return &this
}

// NewAlertserviceTest0WithDefaults instantiates a new AlertserviceTest0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertserviceTest0WithDefaults() *AlertserviceTest0 {
	this := AlertserviceTest0{}
	var enabled bool
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertserviceTest0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertserviceTest0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertserviceTest0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertserviceTest0) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlertserviceTest0) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertserviceTest0) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlertserviceTest0) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlertserviceTest0) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AlertserviceTest0) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertserviceTest0) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AlertserviceTest0) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *AlertserviceTest0) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *AlertserviceTest0) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertserviceTest0) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *AlertserviceTest0) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *AlertserviceTest0) SetLevel(v string) {
	o.Level = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertserviceTest0) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertserviceTest0) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertserviceTest0) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertserviceTest0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AlertserviceTest0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertserviceTest0) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableAlertserviceTest0 struct {
	value *AlertserviceTest0
	isSet bool
}

func (v NullableAlertserviceTest0) Get() *AlertserviceTest0 {
	return v.value
}

func (v *NullableAlertserviceTest0) Set(val *AlertserviceTest0) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertserviceTest0) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertserviceTest0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertserviceTest0(val *AlertserviceTest0) *NullableAlertserviceTest0 {
	return &NullableAlertserviceTest0{value: val, isSet: true}
}

func (v NullableAlertserviceTest0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertserviceTest0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}