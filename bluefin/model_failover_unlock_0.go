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

// checks if the FailoverUnlock0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailoverUnlock0{}

// FailoverUnlock0 struct for FailoverUnlock0
type FailoverUnlock0 struct {
	Pools    []PoolKeys    `json:"pools,omitempty"`
	Datasets []DatasetKeys `json:"datasets,omitempty"`
}

// NewFailoverUnlock0 instantiates a new FailoverUnlock0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverUnlock0() *FailoverUnlock0 {
	this := FailoverUnlock0{}
	return &this
}

// NewFailoverUnlock0WithDefaults instantiates a new FailoverUnlock0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverUnlock0WithDefaults() *FailoverUnlock0 {
	this := FailoverUnlock0{}
	return &this
}

// GetPools returns the Pools field value if set, zero value otherwise.
func (o *FailoverUnlock0) GetPools() []PoolKeys {
	if o == nil || IsNil(o.Pools) {
		var ret []PoolKeys
		return ret
	}
	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverUnlock0) GetPoolsOk() ([]PoolKeys, bool) {
	if o == nil || IsNil(o.Pools) {
		return nil, false
	}
	return o.Pools, true
}

// HasPools returns a boolean if a field has been set.
func (o *FailoverUnlock0) HasPools() bool {
	if o != nil && !IsNil(o.Pools) {
		return true
	}

	return false
}

// SetPools gets a reference to the given []PoolKeys and assigns it to the Pools field.
func (o *FailoverUnlock0) SetPools(v []PoolKeys) {
	o.Pools = v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *FailoverUnlock0) GetDatasets() []DatasetKeys {
	if o == nil || IsNil(o.Datasets) {
		var ret []DatasetKeys
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverUnlock0) GetDatasetsOk() ([]DatasetKeys, bool) {
	if o == nil || IsNil(o.Datasets) {
		return nil, false
	}
	return o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *FailoverUnlock0) HasDatasets() bool {
	if o != nil && !IsNil(o.Datasets) {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given []DatasetKeys and assigns it to the Datasets field.
func (o *FailoverUnlock0) SetDatasets(v []DatasetKeys) {
	o.Datasets = v
}

func (o FailoverUnlock0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailoverUnlock0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pools) {
		toSerialize["pools"] = o.Pools
	}
	if !IsNil(o.Datasets) {
		toSerialize["datasets"] = o.Datasets
	}
	return toSerialize, nil
}

type NullableFailoverUnlock0 struct {
	value *FailoverUnlock0
	isSet bool
}

func (v NullableFailoverUnlock0) Get() *FailoverUnlock0 {
	return v.value
}

func (v *NullableFailoverUnlock0) Set(val *FailoverUnlock0) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverUnlock0) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverUnlock0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverUnlock0(val *FailoverUnlock0) *NullableFailoverUnlock0 {
	return &NullableFailoverUnlock0{value: val, isSet: true}
}

func (v NullableFailoverUnlock0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverUnlock0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}