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

// checks if the DatasetKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetKeys{}

// DatasetKeys struct for DatasetKeys
type DatasetKeys struct {
	Name       *string `json:"name,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewDatasetKeys instantiates a new DatasetKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetKeys() *DatasetKeys {
	this := DatasetKeys{}
	return &this
}

// NewDatasetKeysWithDefaults instantiates a new DatasetKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetKeysWithDefaults() *DatasetKeys {
	this := DatasetKeys{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatasetKeys) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetKeys) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatasetKeys) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatasetKeys) SetName(v string) {
	o.Name = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *DatasetKeys) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetKeys) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *DatasetKeys) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *DatasetKeys) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o DatasetKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return toSerialize, nil
}

type NullableDatasetKeys struct {
	value *DatasetKeys
	isSet bool
}

func (v NullableDatasetKeys) Get() *DatasetKeys {
	return v.value
}

func (v *NullableDatasetKeys) Set(val *DatasetKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetKeys(val *DatasetKeys) *NullableDatasetKeys {
	return &NullableDatasetKeys{value: val, isSet: true}
}

func (v NullableDatasetKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}