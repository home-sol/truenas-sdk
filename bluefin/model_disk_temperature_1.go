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

// checks if the DiskTemperature1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskTemperature1{}

// DiskTemperature1 struct for DiskTemperature1
type DiskTemperature1 struct {
	Cache     NullableInt32 `json:"cache,omitempty"`
	Powermode *string       `json:"powermode,omitempty"`
}

// NewDiskTemperature1 instantiates a new DiskTemperature1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskTemperature1() *DiskTemperature1 {
	this := DiskTemperature1{}
	var powermode string
	this.Powermode = &powermode
	return &this
}

// NewDiskTemperature1WithDefaults instantiates a new DiskTemperature1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskTemperature1WithDefaults() *DiskTemperature1 {
	this := DiskTemperature1{}
	var powermode string
	this.Powermode = &powermode
	return &this
}

// GetCache returns the Cache field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiskTemperature1) GetCache() int32 {
	if o == nil || IsNil(o.Cache.Get()) {
		var ret int32
		return ret
	}
	return *o.Cache.Get()
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiskTemperature1) GetCacheOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cache.Get(), o.Cache.IsSet()
}

// HasCache returns a boolean if a field has been set.
func (o *DiskTemperature1) HasCache() bool {
	if o != nil && o.Cache.IsSet() {
		return true
	}

	return false
}

// SetCache gets a reference to the given NullableInt32 and assigns it to the Cache field.
func (o *DiskTemperature1) SetCache(v int32) {
	o.Cache.Set(&v)
}

// SetCacheNil sets the value for Cache to be an explicit nil
func (o *DiskTemperature1) SetCacheNil() {
	o.Cache.Set(nil)
}

// UnsetCache ensures that no value is present for Cache, not even an explicit nil
func (o *DiskTemperature1) UnsetCache() {
	o.Cache.Unset()
}

// GetPowermode returns the Powermode field value if set, zero value otherwise.
func (o *DiskTemperature1) GetPowermode() string {
	if o == nil || IsNil(o.Powermode) {
		var ret string
		return ret
	}
	return *o.Powermode
}

// GetPowermodeOk returns a tuple with the Powermode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskTemperature1) GetPowermodeOk() (*string, bool) {
	if o == nil || IsNil(o.Powermode) {
		return nil, false
	}
	return o.Powermode, true
}

// HasPowermode returns a boolean if a field has been set.
func (o *DiskTemperature1) HasPowermode() bool {
	if o != nil && !IsNil(o.Powermode) {
		return true
	}

	return false
}

// SetPowermode gets a reference to the given string and assigns it to the Powermode field.
func (o *DiskTemperature1) SetPowermode(v string) {
	o.Powermode = &v
}

func (o DiskTemperature1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskTemperature1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cache.IsSet() {
		toSerialize["cache"] = o.Cache.Get()
	}
	if !IsNil(o.Powermode) {
		toSerialize["powermode"] = o.Powermode
	}
	return toSerialize, nil
}

type NullableDiskTemperature1 struct {
	value *DiskTemperature1
	isSet bool
}

func (v NullableDiskTemperature1) Get() *DiskTemperature1 {
	return v.value
}

func (v *NullableDiskTemperature1) Set(val *DiskTemperature1) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskTemperature1) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskTemperature1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskTemperature1(val *DiskTemperature1) *NullableDiskTemperature1 {
	return &NullableDiskTemperature1{value: val, isSet: true}
}

func (v NullableDiskTemperature1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskTemperature1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
