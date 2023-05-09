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

// checks if the UpdateDiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDiskRequest{}

// UpdateDiskRequest struct for UpdateDiskRequest
type UpdateDiskRequest struct {
	Number               int32          `json:"number"`
	Lunid                NullableString `json:"lunid"`
	Description          NullableString `json:"description"`
	Critical             NullableInt32  `json:"critical"`
	Difference           NullableInt32  `json:"difference"`
	Informational        NullableInt32  `json:"informational"`
	Hddstandby           HDDStandby     `json:"hddstandby"`
	Advpowermgmt         AdvPowermgmt   `json:"advpowermgmt"`
	Togglesmart          bool           `json:"togglesmart"`
	SupportsSmart        *bool          `json:"supports_smart,omitempty"`
	Smartoptions         *string        `json:"smartoptions,omitempty"`
	Passwd               *string        `json:"passwd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDiskRequest UpdateDiskRequest

// NewUpdateDiskRequest instantiates a new UpdateDiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDiskRequest(number int32, lunid NullableString, description NullableString, critical NullableInt32, difference NullableInt32, informational NullableInt32, hddstandby HDDStandby, advpowermgmt AdvPowermgmt, togglesmart bool) *UpdateDiskRequest {
	this := UpdateDiskRequest{}
	this.Number = number
	this.Lunid = lunid
	this.Description = description
	this.Critical = critical
	this.Difference = difference
	this.Informational = informational
	this.Hddstandby = hddstandby
	this.Advpowermgmt = advpowermgmt
	this.Togglesmart = togglesmart
	return &this
}

// NewUpdateDiskRequestWithDefaults instantiates a new UpdateDiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDiskRequestWithDefaults() *UpdateDiskRequest {
	this := UpdateDiskRequest{}
	return &this
}

// GetNumber returns the Number field value
func (o *UpdateDiskRequest) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *UpdateDiskRequest) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *UpdateDiskRequest) SetNumber(v int32) {
	o.Number = v
}

// GetLunid returns the Lunid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateDiskRequest) GetLunid() string {
	if o == nil || o.Lunid.Get() == nil {
		var ret string
		return ret
	}

	return *o.Lunid.Get()
}

// GetLunidOk returns a tuple with the Lunid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiskRequest) GetLunidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lunid.Get(), o.Lunid.IsSet()
}

// SetLunid sets field value
func (o *UpdateDiskRequest) SetLunid(v string) {
	o.Lunid.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateDiskRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiskRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *UpdateDiskRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCritical returns the Critical field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UpdateDiskRequest) GetCritical() int32 {
	if o == nil || o.Critical.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Critical.Get()
}

// GetCriticalOk returns a tuple with the Critical field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiskRequest) GetCriticalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Critical.Get(), o.Critical.IsSet()
}

// SetCritical sets field value
func (o *UpdateDiskRequest) SetCritical(v int32) {
	o.Critical.Set(&v)
}

// GetDifference returns the Difference field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UpdateDiskRequest) GetDifference() int32 {
	if o == nil || o.Difference.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Difference.Get()
}

// GetDifferenceOk returns a tuple with the Difference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiskRequest) GetDifferenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Difference.Get(), o.Difference.IsSet()
}

// SetDifference sets field value
func (o *UpdateDiskRequest) SetDifference(v int32) {
	o.Difference.Set(&v)
}

// GetInformational returns the Informational field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UpdateDiskRequest) GetInformational() int32 {
	if o == nil || o.Informational.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Informational.Get()
}

// GetInformationalOk returns a tuple with the Informational field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiskRequest) GetInformationalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Informational.Get(), o.Informational.IsSet()
}

// SetInformational sets field value
func (o *UpdateDiskRequest) SetInformational(v int32) {
	o.Informational.Set(&v)
}

// GetHddstandby returns the Hddstandby field value
func (o *UpdateDiskRequest) GetHddstandby() HDDStandby {
	if o == nil {
		var ret HDDStandby
		return ret
	}

	return o.Hddstandby
}

// GetHddstandbyOk returns a tuple with the Hddstandby field value
// and a boolean to check if the value has been set.
func (o *UpdateDiskRequest) GetHddstandbyOk() (*HDDStandby, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hddstandby, true
}

// SetHddstandby sets field value
func (o *UpdateDiskRequest) SetHddstandby(v HDDStandby) {
	o.Hddstandby = v
}

// GetAdvpowermgmt returns the Advpowermgmt field value
func (o *UpdateDiskRequest) GetAdvpowermgmt() AdvPowermgmt {
	if o == nil {
		var ret AdvPowermgmt
		return ret
	}

	return o.Advpowermgmt
}

// GetAdvpowermgmtOk returns a tuple with the Advpowermgmt field value
// and a boolean to check if the value has been set.
func (o *UpdateDiskRequest) GetAdvpowermgmtOk() (*AdvPowermgmt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Advpowermgmt, true
}

// SetAdvpowermgmt sets field value
func (o *UpdateDiskRequest) SetAdvpowermgmt(v AdvPowermgmt) {
	o.Advpowermgmt = v
}

// GetTogglesmart returns the Togglesmart field value
func (o *UpdateDiskRequest) GetTogglesmart() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Togglesmart
}

// GetTogglesmartOk returns a tuple with the Togglesmart field value
// and a boolean to check if the value has been set.
func (o *UpdateDiskRequest) GetTogglesmartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Togglesmart, true
}

// SetTogglesmart sets field value
func (o *UpdateDiskRequest) SetTogglesmart(v bool) {
	o.Togglesmart = v
}

// GetSupportsSmart returns the SupportsSmart field value if set, zero value otherwise.
func (o *UpdateDiskRequest) GetSupportsSmart() bool {
	if o == nil || IsNil(o.SupportsSmart) {
		var ret bool
		return ret
	}
	return *o.SupportsSmart
}

// GetSupportsSmartOk returns a tuple with the SupportsSmart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiskRequest) GetSupportsSmartOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsSmart) {
		return nil, false
	}
	return o.SupportsSmart, true
}

// HasSupportsSmart returns a boolean if a field has been set.
func (o *UpdateDiskRequest) HasSupportsSmart() bool {
	if o != nil && !IsNil(o.SupportsSmart) {
		return true
	}

	return false
}

// SetSupportsSmart gets a reference to the given bool and assigns it to the SupportsSmart field.
func (o *UpdateDiskRequest) SetSupportsSmart(v bool) {
	o.SupportsSmart = &v
}

// GetSmartoptions returns the Smartoptions field value if set, zero value otherwise.
func (o *UpdateDiskRequest) GetSmartoptions() string {
	if o == nil || IsNil(o.Smartoptions) {
		var ret string
		return ret
	}
	return *o.Smartoptions
}

// GetSmartoptionsOk returns a tuple with the Smartoptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiskRequest) GetSmartoptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Smartoptions) {
		return nil, false
	}
	return o.Smartoptions, true
}

// HasSmartoptions returns a boolean if a field has been set.
func (o *UpdateDiskRequest) HasSmartoptions() bool {
	if o != nil && !IsNil(o.Smartoptions) {
		return true
	}

	return false
}

// SetSmartoptions gets a reference to the given string and assigns it to the Smartoptions field.
func (o *UpdateDiskRequest) SetSmartoptions(v string) {
	o.Smartoptions = &v
}

// GetPasswd returns the Passwd field value if set, zero value otherwise.
func (o *UpdateDiskRequest) GetPasswd() string {
	if o == nil || IsNil(o.Passwd) {
		var ret string
		return ret
	}
	return *o.Passwd
}

// GetPasswdOk returns a tuple with the Passwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiskRequest) GetPasswdOk() (*string, bool) {
	if o == nil || IsNil(o.Passwd) {
		return nil, false
	}
	return o.Passwd, true
}

// HasPasswd returns a boolean if a field has been set.
func (o *UpdateDiskRequest) HasPasswd() bool {
	if o != nil && !IsNil(o.Passwd) {
		return true
	}

	return false
}

// SetPasswd gets a reference to the given string and assigns it to the Passwd field.
func (o *UpdateDiskRequest) SetPasswd(v string) {
	o.Passwd = &v
}

func (o UpdateDiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	toSerialize["lunid"] = o.Lunid.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["critical"] = o.Critical.Get()
	toSerialize["difference"] = o.Difference.Get()
	toSerialize["informational"] = o.Informational.Get()
	toSerialize["hddstandby"] = o.Hddstandby
	toSerialize["advpowermgmt"] = o.Advpowermgmt
	toSerialize["togglesmart"] = o.Togglesmart
	if !IsNil(o.SupportsSmart) {
		toSerialize["supports_smart"] = o.SupportsSmart
	}
	if !IsNil(o.Smartoptions) {
		toSerialize["smartoptions"] = o.Smartoptions
	}
	if !IsNil(o.Passwd) {
		toSerialize["passwd"] = o.Passwd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateDiskRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateDiskRequest := _UpdateDiskRequest{}

	if err = json.Unmarshal(bytes, &varUpdateDiskRequest); err == nil {
		*o = UpdateDiskRequest(varUpdateDiskRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "number")
		delete(additionalProperties, "lunid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "critical")
		delete(additionalProperties, "difference")
		delete(additionalProperties, "informational")
		delete(additionalProperties, "hddstandby")
		delete(additionalProperties, "advpowermgmt")
		delete(additionalProperties, "togglesmart")
		delete(additionalProperties, "supports_smart")
		delete(additionalProperties, "smartoptions")
		delete(additionalProperties, "passwd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDiskRequest struct {
	value *UpdateDiskRequest
	isSet bool
}

func (v NullableUpdateDiskRequest) Get() *UpdateDiskRequest {
	return v.value
}

func (v *NullableUpdateDiskRequest) Set(val *UpdateDiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDiskRequest(val *UpdateDiskRequest) *NullableUpdateDiskRequest {
	return &NullableUpdateDiskRequest{value: val, isSet: true}
}

func (v NullableUpdateDiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}