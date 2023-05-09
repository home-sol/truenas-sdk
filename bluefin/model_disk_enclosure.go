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

// checks if the DiskEnclosure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskEnclosure{}

// DiskEnclosure struct for DiskEnclosure
type DiskEnclosure struct {
	Number *int32 `json:"number,omitempty"`
	Slot   *int32 `json:"slot,omitempty"`
}

// NewDiskEnclosure instantiates a new DiskEnclosure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskEnclosure() *DiskEnclosure {
	this := DiskEnclosure{}
	return &this
}

// NewDiskEnclosureWithDefaults instantiates a new DiskEnclosure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskEnclosureWithDefaults() *DiskEnclosure {
	this := DiskEnclosure{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *DiskEnclosure) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskEnclosure) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *DiskEnclosure) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *DiskEnclosure) SetNumber(v int32) {
	o.Number = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *DiskEnclosure) GetSlot() int32 {
	if o == nil || IsNil(o.Slot) {
		var ret int32
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskEnclosure) GetSlotOk() (*int32, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *DiskEnclosure) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given int32 and assigns it to the Slot field.
func (o *DiskEnclosure) SetSlot(v int32) {
	o.Slot = &v
}

func (o DiskEnclosure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskEnclosure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	return toSerialize, nil
}

type NullableDiskEnclosure struct {
	value *DiskEnclosure
	isSet bool
}

func (v NullableDiskEnclosure) Get() *DiskEnclosure {
	return v.value
}

func (v *NullableDiskEnclosure) Set(val *DiskEnclosure) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskEnclosure) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskEnclosure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskEnclosure(val *DiskEnclosure) *NullableDiskEnclosure {
	return &NullableDiskEnclosure{value: val, isSet: true}
}

func (v NullableDiskEnclosure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskEnclosure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}