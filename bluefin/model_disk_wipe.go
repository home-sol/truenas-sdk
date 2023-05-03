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

// checks if the DiskWipe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskWipe{}

// DiskWipe struct for DiskWipe
type DiskWipe struct {
	Dev                  *string    `json:"dev,omitempty"`
	Mode                 *DiskWipe1 `json:"mode,omitempty"`
	Synccache            *bool      `json:"synccache,omitempty"`
	SwapRemovalOptions   *DiskWipe3 `json:"swap_removal_options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiskWipe DiskWipe

// NewDiskWipe instantiates a new DiskWipe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskWipe() *DiskWipe {
	this := DiskWipe{}
	var synccache bool
	this.Synccache = &synccache
	var swapRemovalOptions DiskWipe3
	this.SwapRemovalOptions = &swapRemovalOptions
	return &this
}

// NewDiskWipeWithDefaults instantiates a new DiskWipe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskWipeWithDefaults() *DiskWipe {
	this := DiskWipe{}
	var synccache bool
	this.Synccache = &synccache
	var swapRemovalOptions DiskWipe3
	this.SwapRemovalOptions = &swapRemovalOptions
	return &this
}

// GetDev returns the Dev field value if set, zero value otherwise.
func (o *DiskWipe) GetDev() string {
	if o == nil || IsNil(o.Dev) {
		var ret string
		return ret
	}
	return *o.Dev
}

// GetDevOk returns a tuple with the Dev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskWipe) GetDevOk() (*string, bool) {
	if o == nil || IsNil(o.Dev) {
		return nil, false
	}
	return o.Dev, true
}

// HasDev returns a boolean if a field has been set.
func (o *DiskWipe) HasDev() bool {
	if o != nil && !IsNil(o.Dev) {
		return true
	}

	return false
}

// SetDev gets a reference to the given string and assigns it to the Dev field.
func (o *DiskWipe) SetDev(v string) {
	o.Dev = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DiskWipe) GetMode() DiskWipe1 {
	if o == nil || IsNil(o.Mode) {
		var ret DiskWipe1
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskWipe) GetModeOk() (*DiskWipe1, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DiskWipe) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given DiskWipe1 and assigns it to the Mode field.
func (o *DiskWipe) SetMode(v DiskWipe1) {
	o.Mode = &v
}

// GetSynccache returns the Synccache field value if set, zero value otherwise.
func (o *DiskWipe) GetSynccache() bool {
	if o == nil || IsNil(o.Synccache) {
		var ret bool
		return ret
	}
	return *o.Synccache
}

// GetSynccacheOk returns a tuple with the Synccache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskWipe) GetSynccacheOk() (*bool, bool) {
	if o == nil || IsNil(o.Synccache) {
		return nil, false
	}
	return o.Synccache, true
}

// HasSynccache returns a boolean if a field has been set.
func (o *DiskWipe) HasSynccache() bool {
	if o != nil && !IsNil(o.Synccache) {
		return true
	}

	return false
}

// SetSynccache gets a reference to the given bool and assigns it to the Synccache field.
func (o *DiskWipe) SetSynccache(v bool) {
	o.Synccache = &v
}

// GetSwapRemovalOptions returns the SwapRemovalOptions field value if set, zero value otherwise.
func (o *DiskWipe) GetSwapRemovalOptions() DiskWipe3 {
	if o == nil || IsNil(o.SwapRemovalOptions) {
		var ret DiskWipe3
		return ret
	}
	return *o.SwapRemovalOptions
}

// GetSwapRemovalOptionsOk returns a tuple with the SwapRemovalOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskWipe) GetSwapRemovalOptionsOk() (*DiskWipe3, bool) {
	if o == nil || IsNil(o.SwapRemovalOptions) {
		return nil, false
	}
	return o.SwapRemovalOptions, true
}

// HasSwapRemovalOptions returns a boolean if a field has been set.
func (o *DiskWipe) HasSwapRemovalOptions() bool {
	if o != nil && !IsNil(o.SwapRemovalOptions) {
		return true
	}

	return false
}

// SetSwapRemovalOptions gets a reference to the given DiskWipe3 and assigns it to the SwapRemovalOptions field.
func (o *DiskWipe) SetSwapRemovalOptions(v DiskWipe3) {
	o.SwapRemovalOptions = &v
}

func (o DiskWipe) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskWipe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dev) {
		toSerialize["dev"] = o.Dev
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Synccache) {
		toSerialize["synccache"] = o.Synccache
	}
	if !IsNil(o.SwapRemovalOptions) {
		toSerialize["swap_removal_options"] = o.SwapRemovalOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiskWipe) UnmarshalJSON(bytes []byte) (err error) {
	varDiskWipe := _DiskWipe{}

	if err = json.Unmarshal(bytes, &varDiskWipe); err == nil {
		*o = DiskWipe(varDiskWipe)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dev")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "synccache")
		delete(additionalProperties, "swap_removal_options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiskWipe struct {
	value *DiskWipe
	isSet bool
}

func (v NullableDiskWipe) Get() *DiskWipe {
	return v.value
}

func (v *NullableDiskWipe) Set(val *DiskWipe) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskWipe) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskWipe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskWipe(val *DiskWipe) *NullableDiskWipe {
	return &NullableDiskWipe{value: val, isSet: true}
}

func (v NullableDiskWipe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskWipe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}