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

// checks if the Dataset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dataset{}

// Dataset struct for Dataset
type Dataset struct {
	Force      *bool   `json:"force,omitempty"`
	Name       *string `json:"name,omitempty"`
	Key        *string `json:"key,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewDataset instantiates a new Dataset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataset() *Dataset {
	this := Dataset{}
	var force bool
	this.Force = &force
	return &this
}

// NewDatasetWithDefaults instantiates a new Dataset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetWithDefaults() *Dataset {
	this := Dataset{}
	var force bool
	this.Force = &force
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *Dataset) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *Dataset) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *Dataset) SetForce(v bool) {
	o.Force = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dataset) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dataset) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dataset) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Dataset) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Dataset) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Dataset) SetKey(v string) {
	o.Key = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *Dataset) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *Dataset) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *Dataset) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o Dataset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dataset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return toSerialize, nil
}

type NullableDataset struct {
	value *Dataset
	isSet bool
}

func (v NullableDataset) Get() *Dataset {
	return v.value
}

func (v *NullableDataset) Set(val *Dataset) {
	v.value = val
	v.isSet = true
}

func (v NullableDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataset(val *Dataset) *NullableDataset {
	return &NullableDataset{value: val, isSet: true}
}

func (v NullableDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}