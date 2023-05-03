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

// checks if the StatsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatsData{}

// StatsData struct for StatsData
type StatsData struct {
	Source  *string `json:"source,omitempty"`
	Type    *string `json:"type,omitempty"`
	Dataset *string `json:"dataset,omitempty"`
	Cf      *string `json:"cf,omitempty"`
}

// NewStatsData instantiates a new StatsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsData() *StatsData {
	this := StatsData{}
	var cf string
	this.Cf = &cf
	return &this
}

// NewStatsDataWithDefaults instantiates a new StatsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsDataWithDefaults() *StatsData {
	this := StatsData{}
	var cf string
	this.Cf = &cf
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StatsData) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsData) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StatsData) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StatsData) SetSource(v string) {
	o.Source = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StatsData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StatsData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StatsData) SetType(v string) {
	o.Type = &v
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *StatsData) GetDataset() string {
	if o == nil || IsNil(o.Dataset) {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsData) GetDatasetOk() (*string, bool) {
	if o == nil || IsNil(o.Dataset) {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *StatsData) HasDataset() bool {
	if o != nil && !IsNil(o.Dataset) {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *StatsData) SetDataset(v string) {
	o.Dataset = &v
}

// GetCf returns the Cf field value if set, zero value otherwise.
func (o *StatsData) GetCf() string {
	if o == nil || IsNil(o.Cf) {
		var ret string
		return ret
	}
	return *o.Cf
}

// GetCfOk returns a tuple with the Cf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsData) GetCfOk() (*string, bool) {
	if o == nil || IsNil(o.Cf) {
		return nil, false
	}
	return o.Cf, true
}

// HasCf returns a boolean if a field has been set.
func (o *StatsData) HasCf() bool {
	if o != nil && !IsNil(o.Cf) {
		return true
	}

	return false
}

// SetCf gets a reference to the given string and assigns it to the Cf field.
func (o *StatsData) SetCf(v string) {
	o.Cf = &v
}

func (o StatsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Dataset) {
		toSerialize["dataset"] = o.Dataset
	}
	if !IsNil(o.Cf) {
		toSerialize["cf"] = o.Cf
	}
	return toSerialize, nil
}

type NullableStatsData struct {
	value *StatsData
	isSet bool
}

func (v NullableStatsData) Get() *StatsData {
	return v.value
}

func (v *NullableStatsData) Set(val *StatsData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsData(val *StatsData) *NullableStatsData {
	return &NullableStatsData{value: val, isSet: true}
}

func (v NullableStatsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
