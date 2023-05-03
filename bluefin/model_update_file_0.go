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

// checks if the UpdateFile0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFile0{}

// UpdateFile0 struct for UpdateFile0
type UpdateFile0 struct {
	Destination NullableString `json:"destination,omitempty"`
}

// NewUpdateFile0 instantiates a new UpdateFile0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFile0() *UpdateFile0 {
	this := UpdateFile0{}
	return &this
}

// NewUpdateFile0WithDefaults instantiates a new UpdateFile0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFile0WithDefaults() *UpdateFile0 {
	this := UpdateFile0{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFile0) GetDestination() string {
	if o == nil || IsNil(o.Destination.Get()) {
		var ret string
		return ret
	}
	return *o.Destination.Get()
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFile0) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destination.Get(), o.Destination.IsSet()
}

// HasDestination returns a boolean if a field has been set.
func (o *UpdateFile0) HasDestination() bool {
	if o != nil && o.Destination.IsSet() {
		return true
	}

	return false
}

// SetDestination gets a reference to the given NullableString and assigns it to the Destination field.
func (o *UpdateFile0) SetDestination(v string) {
	o.Destination.Set(&v)
}

// SetDestinationNil sets the value for Destination to be an explicit nil
func (o *UpdateFile0) SetDestinationNil() {
	o.Destination.Set(nil)
}

// UnsetDestination ensures that no value is present for Destination, not even an explicit nil
func (o *UpdateFile0) UnsetDestination() {
	o.Destination.Unset()
}

func (o UpdateFile0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFile0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination.IsSet() {
		toSerialize["destination"] = o.Destination.Get()
	}
	return toSerialize, nil
}

type NullableUpdateFile0 struct {
	value *UpdateFile0
	isSet bool
}

func (v NullableUpdateFile0) Get() *UpdateFile0 {
	return v.value
}

func (v *NullableUpdateFile0) Set(val *UpdateFile0) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFile0) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFile0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFile0(val *UpdateFile0) *NullableUpdateFile0 {
	return &NullableUpdateFile0{value: val, isSet: true}
}

func (v NullableUpdateFile0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFile0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}