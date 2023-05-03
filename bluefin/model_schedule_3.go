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

// checks if the Schedule3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule3{}

// Schedule3 struct for Schedule3
type Schedule3 struct {
	Minute *string `json:"minute,omitempty"`
	Hour   *string `json:"hour,omitempty"`
	Dom    *string `json:"dom,omitempty"`
	Month  *string `json:"month,omitempty"`
	Dow    *string `json:"dow,omitempty"`
	Begin  *string `json:"begin,omitempty"`
	End    *string `json:"end,omitempty"`
}

// NewSchedule3 instantiates a new Schedule3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule3() *Schedule3 {
	this := Schedule3{}
	var minute string
	this.Minute = &minute
	var hour string
	this.Hour = &hour
	var dom string
	this.Dom = &dom
	var month string
	this.Month = &month
	var dow string
	this.Dow = &dow
	var begin string
	this.Begin = &begin
	var end string
	this.End = &end
	return &this
}

// NewSchedule3WithDefaults instantiates a new Schedule3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule3WithDefaults() *Schedule3 {
	this := Schedule3{}
	var minute string
	this.Minute = &minute
	var hour string
	this.Hour = &hour
	var dom string
	this.Dom = &dom
	var month string
	this.Month = &month
	var dow string
	this.Dow = &dow
	var begin string
	this.Begin = &begin
	var end string
	this.End = &end
	return &this
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *Schedule3) GetMinute() string {
	if o == nil || IsNil(o.Minute) {
		var ret string
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule3) GetMinuteOk() (*string, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *Schedule3) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given string and assigns it to the Minute field.
func (o *Schedule3) SetMinute(v string) {
	o.Minute = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *Schedule3) GetHour() string {
	if o == nil || IsNil(o.Hour) {
		var ret string
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule3) GetHourOk() (*string, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *Schedule3) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given string and assigns it to the Hour field.
func (o *Schedule3) SetHour(v string) {
	o.Hour = &v
}

// GetDom returns the Dom field value if set, zero value otherwise.
func (o *Schedule3) GetDom() string {
	if o == nil || IsNil(o.Dom) {
		var ret string
		return ret
	}
	return *o.Dom
}

// GetDomOk returns a tuple with the Dom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule3) GetDomOk() (*string, bool) {
	if o == nil || IsNil(o.Dom) {
		return nil, false
	}
	return o.Dom, true
}

// HasDom returns a boolean if a field has been set.
func (o *Schedule3) HasDom() bool {
	if o != nil && !IsNil(o.Dom) {
		return true
	}

	return false
}

// SetDom gets a reference to the given string and assigns it to the Dom field.
func (o *Schedule3) SetDom(v string) {
	o.Dom = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *Schedule3) GetMonth() string {
	if o == nil || IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule3) GetMonthOk() (*string, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *Schedule3) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *Schedule3) SetMonth(v string) {
	o.Month = &v
}

// GetDow returns the Dow field value if set, zero value otherwise.
func (o *Schedule3) GetDow() string {
	if o == nil || IsNil(o.Dow) {
		var ret string
		return ret
	}
	return *o.Dow
}

// GetDowOk returns a tuple with the Dow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule3) GetDowOk() (*string, bool) {
	if o == nil || IsNil(o.Dow) {
		return nil, false
	}
	return o.Dow, true
}

// HasDow returns a boolean if a field has been set.
func (o *Schedule3) HasDow() bool {
	if o != nil && !IsNil(o.Dow) {
		return true
	}

	return false
}

// SetDow gets a reference to the given string and assigns it to the Dow field.
func (o *Schedule3) SetDow(v string) {
	o.Dow = &v
}

// GetBegin returns the Begin field value if set, zero value otherwise.
func (o *Schedule3) GetBegin() string {
	if o == nil || IsNil(o.Begin) {
		var ret string
		return ret
	}
	return *o.Begin
}

// GetBeginOk returns a tuple with the Begin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule3) GetBeginOk() (*string, bool) {
	if o == nil || IsNil(o.Begin) {
		return nil, false
	}
	return o.Begin, true
}

// HasBegin returns a boolean if a field has been set.
func (o *Schedule3) HasBegin() bool {
	if o != nil && !IsNil(o.Begin) {
		return true
	}

	return false
}

// SetBegin gets a reference to the given string and assigns it to the Begin field.
func (o *Schedule3) SetBegin(v string) {
	o.Begin = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Schedule3) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule3) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Schedule3) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *Schedule3) SetEnd(v string) {
	o.End = &v
}

func (o Schedule3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.Dom) {
		toSerialize["dom"] = o.Dom
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !IsNil(o.Dow) {
		toSerialize["dow"] = o.Dow
	}
	if !IsNil(o.Begin) {
		toSerialize["begin"] = o.Begin
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableSchedule3 struct {
	value *Schedule3
	isSet bool
}

func (v NullableSchedule3) Get() *Schedule3 {
	return v.value
}

func (v *NullableSchedule3) Set(val *Schedule3) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule3) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule3(val *Schedule3) *NullableSchedule3 {
	return &NullableSchedule3{value: val, isSet: true}
}

func (v NullableSchedule3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
