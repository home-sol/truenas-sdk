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

// checks if the RestrictSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestrictSchedule{}

// RestrictSchedule * `restrict_schedule` restricts when replication task with bound periodic snapshot tasks runs. For example,   you can have periodic snapshot tasks that run every 15 minutes, but only run replication task every hour. * Enabling `only_matching_schedule` will only replicate snapshots that match `schedule` or   `restrict_schedule`
type RestrictSchedule struct {
	Minute *string `json:"minute,omitempty"`
	Hour   *string `json:"hour,omitempty"`
	Dom    *string `json:"dom,omitempty"`
	Month  *string `json:"month,omitempty"`
	Dow    *string `json:"dow,omitempty"`
	Begin  *string `json:"begin,omitempty"`
	End    *string `json:"end,omitempty"`
}

// NewRestrictSchedule instantiates a new RestrictSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictSchedule() *RestrictSchedule {
	this := RestrictSchedule{}
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

// NewRestrictScheduleWithDefaults instantiates a new RestrictSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictScheduleWithDefaults() *RestrictSchedule {
	this := RestrictSchedule{}
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
func (o *RestrictSchedule) GetMinute() string {
	if o == nil || IsNil(o.Minute) {
		var ret string
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictSchedule) GetMinuteOk() (*string, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *RestrictSchedule) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given string and assigns it to the Minute field.
func (o *RestrictSchedule) SetMinute(v string) {
	o.Minute = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *RestrictSchedule) GetHour() string {
	if o == nil || IsNil(o.Hour) {
		var ret string
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictSchedule) GetHourOk() (*string, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *RestrictSchedule) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given string and assigns it to the Hour field.
func (o *RestrictSchedule) SetHour(v string) {
	o.Hour = &v
}

// GetDom returns the Dom field value if set, zero value otherwise.
func (o *RestrictSchedule) GetDom() string {
	if o == nil || IsNil(o.Dom) {
		var ret string
		return ret
	}
	return *o.Dom
}

// GetDomOk returns a tuple with the Dom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictSchedule) GetDomOk() (*string, bool) {
	if o == nil || IsNil(o.Dom) {
		return nil, false
	}
	return o.Dom, true
}

// HasDom returns a boolean if a field has been set.
func (o *RestrictSchedule) HasDom() bool {
	if o != nil && !IsNil(o.Dom) {
		return true
	}

	return false
}

// SetDom gets a reference to the given string and assigns it to the Dom field.
func (o *RestrictSchedule) SetDom(v string) {
	o.Dom = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *RestrictSchedule) GetMonth() string {
	if o == nil || IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictSchedule) GetMonthOk() (*string, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *RestrictSchedule) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *RestrictSchedule) SetMonth(v string) {
	o.Month = &v
}

// GetDow returns the Dow field value if set, zero value otherwise.
func (o *RestrictSchedule) GetDow() string {
	if o == nil || IsNil(o.Dow) {
		var ret string
		return ret
	}
	return *o.Dow
}

// GetDowOk returns a tuple with the Dow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictSchedule) GetDowOk() (*string, bool) {
	if o == nil || IsNil(o.Dow) {
		return nil, false
	}
	return o.Dow, true
}

// HasDow returns a boolean if a field has been set.
func (o *RestrictSchedule) HasDow() bool {
	if o != nil && !IsNil(o.Dow) {
		return true
	}

	return false
}

// SetDow gets a reference to the given string and assigns it to the Dow field.
func (o *RestrictSchedule) SetDow(v string) {
	o.Dow = &v
}

// GetBegin returns the Begin field value if set, zero value otherwise.
func (o *RestrictSchedule) GetBegin() string {
	if o == nil || IsNil(o.Begin) {
		var ret string
		return ret
	}
	return *o.Begin
}

// GetBeginOk returns a tuple with the Begin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictSchedule) GetBeginOk() (*string, bool) {
	if o == nil || IsNil(o.Begin) {
		return nil, false
	}
	return o.Begin, true
}

// HasBegin returns a boolean if a field has been set.
func (o *RestrictSchedule) HasBegin() bool {
	if o != nil && !IsNil(o.Begin) {
		return true
	}

	return false
}

// SetBegin gets a reference to the given string and assigns it to the Begin field.
func (o *RestrictSchedule) SetBegin(v string) {
	o.Begin = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *RestrictSchedule) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictSchedule) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *RestrictSchedule) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *RestrictSchedule) SetEnd(v string) {
	o.End = &v
}

func (o RestrictSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestrictSchedule) ToMap() (map[string]interface{}, error) {
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

type NullableRestrictSchedule struct {
	value *RestrictSchedule
	isSet bool
}

func (v NullableRestrictSchedule) Get() *RestrictSchedule {
	return v.value
}

func (v *NullableRestrictSchedule) Set(val *RestrictSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictSchedule(val *RestrictSchedule) *NullableRestrictSchedule {
	return &NullableRestrictSchedule{value: val, isSet: true}
}

func (v NullableRestrictSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}