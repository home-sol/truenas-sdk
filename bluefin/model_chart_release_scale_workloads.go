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

// checks if the ChartReleaseScaleWorkloads type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChartReleaseScaleWorkloads{}

// ChartReleaseScaleWorkloads struct for ChartReleaseScaleWorkloads
type ChartReleaseScaleWorkloads struct {
	ReleaseName          *string         `json:"release_name,omitempty"`
	Workloads            []ScaleWorkload `json:"workloads,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChartReleaseScaleWorkloads ChartReleaseScaleWorkloads

// NewChartReleaseScaleWorkloads instantiates a new ChartReleaseScaleWorkloads object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartReleaseScaleWorkloads() *ChartReleaseScaleWorkloads {
	this := ChartReleaseScaleWorkloads{}
	return &this
}

// NewChartReleaseScaleWorkloadsWithDefaults instantiates a new ChartReleaseScaleWorkloads object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartReleaseScaleWorkloadsWithDefaults() *ChartReleaseScaleWorkloads {
	this := ChartReleaseScaleWorkloads{}
	return &this
}

// GetReleaseName returns the ReleaseName field value if set, zero value otherwise.
func (o *ChartReleaseScaleWorkloads) GetReleaseName() string {
	if o == nil || IsNil(o.ReleaseName) {
		var ret string
		return ret
	}
	return *o.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartReleaseScaleWorkloads) GetReleaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseName) {
		return nil, false
	}
	return o.ReleaseName, true
}

// HasReleaseName returns a boolean if a field has been set.
func (o *ChartReleaseScaleWorkloads) HasReleaseName() bool {
	if o != nil && !IsNil(o.ReleaseName) {
		return true
	}

	return false
}

// SetReleaseName gets a reference to the given string and assigns it to the ReleaseName field.
func (o *ChartReleaseScaleWorkloads) SetReleaseName(v string) {
	o.ReleaseName = &v
}

// GetWorkloads returns the Workloads field value if set, zero value otherwise.
func (o *ChartReleaseScaleWorkloads) GetWorkloads() []ScaleWorkload {
	if o == nil || IsNil(o.Workloads) {
		var ret []ScaleWorkload
		return ret
	}
	return o.Workloads
}

// GetWorkloadsOk returns a tuple with the Workloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartReleaseScaleWorkloads) GetWorkloadsOk() ([]ScaleWorkload, bool) {
	if o == nil || IsNil(o.Workloads) {
		return nil, false
	}
	return o.Workloads, true
}

// HasWorkloads returns a boolean if a field has been set.
func (o *ChartReleaseScaleWorkloads) HasWorkloads() bool {
	if o != nil && !IsNil(o.Workloads) {
		return true
	}

	return false
}

// SetWorkloads gets a reference to the given []ScaleWorkload and assigns it to the Workloads field.
func (o *ChartReleaseScaleWorkloads) SetWorkloads(v []ScaleWorkload) {
	o.Workloads = v
}

func (o ChartReleaseScaleWorkloads) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChartReleaseScaleWorkloads) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReleaseName) {
		toSerialize["release_name"] = o.ReleaseName
	}
	if !IsNil(o.Workloads) {
		toSerialize["workloads"] = o.Workloads
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChartReleaseScaleWorkloads) UnmarshalJSON(bytes []byte) (err error) {
	varChartReleaseScaleWorkloads := _ChartReleaseScaleWorkloads{}

	if err = json.Unmarshal(bytes, &varChartReleaseScaleWorkloads); err == nil {
		*o = ChartReleaseScaleWorkloads(varChartReleaseScaleWorkloads)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "release_name")
		delete(additionalProperties, "workloads")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChartReleaseScaleWorkloads struct {
	value *ChartReleaseScaleWorkloads
	isSet bool
}

func (v NullableChartReleaseScaleWorkloads) Get() *ChartReleaseScaleWorkloads {
	return v.value
}

func (v *NullableChartReleaseScaleWorkloads) Set(val *ChartReleaseScaleWorkloads) {
	v.value = val
	v.isSet = true
}

func (v NullableChartReleaseScaleWorkloads) IsSet() bool {
	return v.isSet
}

func (v *NullableChartReleaseScaleWorkloads) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartReleaseScaleWorkloads(val *ChartReleaseScaleWorkloads) *NullableChartReleaseScaleWorkloads {
	return &NullableChartReleaseScaleWorkloads{value: val, isSet: true}
}

func (v NullableChartReleaseScaleWorkloads) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartReleaseScaleWorkloads) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
