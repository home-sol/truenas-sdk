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

// checks if the DiskTemperature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskTemperature{}

// DiskTemperature struct for DiskTemperature
type DiskTemperature struct {
	Name                 *string           `json:"name,omitempty"`
	Options              *DiskTemperature1 `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiskTemperature DiskTemperature

// NewDiskTemperature instantiates a new DiskTemperature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskTemperature() *DiskTemperature {
	this := DiskTemperature{}
	var options DiskTemperature1
	this.Options = &options
	return &this
}

// NewDiskTemperatureWithDefaults instantiates a new DiskTemperature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskTemperatureWithDefaults() *DiskTemperature {
	this := DiskTemperature{}
	var options DiskTemperature1
	this.Options = &options
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DiskTemperature) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskTemperature) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DiskTemperature) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DiskTemperature) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DiskTemperature) GetOptions() DiskTemperature1 {
	if o == nil || IsNil(o.Options) {
		var ret DiskTemperature1
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskTemperature) GetOptionsOk() (*DiskTemperature1, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DiskTemperature) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given DiskTemperature1 and assigns it to the Options field.
func (o *DiskTemperature) SetOptions(v DiskTemperature1) {
	o.Options = &v
}

func (o DiskTemperature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskTemperature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiskTemperature) UnmarshalJSON(bytes []byte) (err error) {
	varDiskTemperature := _DiskTemperature{}

	if err = json.Unmarshal(bytes, &varDiskTemperature); err == nil {
		*o = DiskTemperature(varDiskTemperature)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiskTemperature struct {
	value *DiskTemperature
	isSet bool
}

func (v NullableDiskTemperature) Get() *DiskTemperature {
	return v.value
}

func (v *NullableDiskTemperature) Set(val *DiskTemperature) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskTemperature) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskTemperature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskTemperature(val *DiskTemperature) *NullableDiskTemperature {
	return &NullableDiskTemperature{value: val, isSet: true}
}

func (v NullableDiskTemperature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskTemperature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
