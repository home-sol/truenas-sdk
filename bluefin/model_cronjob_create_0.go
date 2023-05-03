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

// checks if the CronjobCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CronjobCreate0{}

// CronjobCreate0 struct for CronjobCreate0
type CronjobCreate0 struct {
	Enabled *bool `json:"enabled,omitempty"`
	// `stderr` and `stdout` are boolean values which if `true`, represent that we would like to suppress standard error / standard output respectively.
	Stderr *bool `json:"stderr,omitempty"`
	// `stderr` and `stdout` are boolean values which if `true`, represent that we would like to suppress standard error / standard output respectively.
	Stdout      *bool     `json:"stdout,omitempty"`
	Schedule    *Schedule `json:"schedule,omitempty"`
	Command     *string   `json:"command,omitempty"`
	Description *string   `json:"description,omitempty"`
	User        *string   `json:"user,omitempty"`
}

// NewCronjobCreate0 instantiates a new CronjobCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronjobCreate0() *CronjobCreate0 {
	this := CronjobCreate0{}
	var stderr bool
	this.Stderr = &stderr
	var stdout bool
	this.Stdout = &stdout
	var schedule Schedule
	this.Schedule = &schedule
	return &this
}

// NewCronjobCreate0WithDefaults instantiates a new CronjobCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronjobCreate0WithDefaults() *CronjobCreate0 {
	this := CronjobCreate0{}
	var stderr bool
	this.Stderr = &stderr
	var stdout bool
	this.Stdout = &stdout
	var schedule Schedule
	this.Schedule = &schedule
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CronjobCreate0) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronjobCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CronjobCreate0) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CronjobCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStderr returns the Stderr field value if set, zero value otherwise.
func (o *CronjobCreate0) GetStderr() bool {
	if o == nil || IsNil(o.Stderr) {
		var ret bool
		return ret
	}
	return *o.Stderr
}

// GetStderrOk returns a tuple with the Stderr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronjobCreate0) GetStderrOk() (*bool, bool) {
	if o == nil || IsNil(o.Stderr) {
		return nil, false
	}
	return o.Stderr, true
}

// HasStderr returns a boolean if a field has been set.
func (o *CronjobCreate0) HasStderr() bool {
	if o != nil && !IsNil(o.Stderr) {
		return true
	}

	return false
}

// SetStderr gets a reference to the given bool and assigns it to the Stderr field.
func (o *CronjobCreate0) SetStderr(v bool) {
	o.Stderr = &v
}

// GetStdout returns the Stdout field value if set, zero value otherwise.
func (o *CronjobCreate0) GetStdout() bool {
	if o == nil || IsNil(o.Stdout) {
		var ret bool
		return ret
	}
	return *o.Stdout
}

// GetStdoutOk returns a tuple with the Stdout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronjobCreate0) GetStdoutOk() (*bool, bool) {
	if o == nil || IsNil(o.Stdout) {
		return nil, false
	}
	return o.Stdout, true
}

// HasStdout returns a boolean if a field has been set.
func (o *CronjobCreate0) HasStdout() bool {
	if o != nil && !IsNil(o.Stdout) {
		return true
	}

	return false
}

// SetStdout gets a reference to the given bool and assigns it to the Stdout field.
func (o *CronjobCreate0) SetStdout(v bool) {
	o.Stdout = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CronjobCreate0) GetSchedule() Schedule {
	if o == nil || IsNil(o.Schedule) {
		var ret Schedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronjobCreate0) GetScheduleOk() (*Schedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CronjobCreate0) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given Schedule and assigns it to the Schedule field.
func (o *CronjobCreate0) SetSchedule(v Schedule) {
	o.Schedule = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *CronjobCreate0) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronjobCreate0) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *CronjobCreate0) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *CronjobCreate0) SetCommand(v string) {
	o.Command = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CronjobCreate0) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronjobCreate0) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CronjobCreate0) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CronjobCreate0) SetDescription(v string) {
	o.Description = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CronjobCreate0) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronjobCreate0) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CronjobCreate0) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CronjobCreate0) SetUser(v string) {
	o.User = &v
}

func (o CronjobCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronjobCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Stderr) {
		toSerialize["stderr"] = o.Stderr
	}
	if !IsNil(o.Stdout) {
		toSerialize["stdout"] = o.Stdout
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableCronjobCreate0 struct {
	value *CronjobCreate0
	isSet bool
}

func (v NullableCronjobCreate0) Get() *CronjobCreate0 {
	return v.value
}

func (v *NullableCronjobCreate0) Set(val *CronjobCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableCronjobCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableCronjobCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronjobCreate0(val *CronjobCreate0) *NullableCronjobCreate0 {
	return &NullableCronjobCreate0{value: val, isSet: true}
}

func (v NullableCronjobCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronjobCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}