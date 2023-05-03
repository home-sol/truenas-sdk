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

// checks if the GlusterVolumeGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlusterVolumeGetInstance{}

// GlusterVolumeGetInstance struct for GlusterVolumeGetInstance
type GlusterVolumeGetInstance struct {
	Id                   *GlusterVolumeGetInstance0 `json:"id,omitempty"`
	QueryOptions         *GlusterVolumeGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GlusterVolumeGetInstance GlusterVolumeGetInstance

// NewGlusterVolumeGetInstance instantiates a new GlusterVolumeGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterVolumeGetInstance() *GlusterVolumeGetInstance {
	this := GlusterVolumeGetInstance{}
	var queryOptions GlusterVolumeGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewGlusterVolumeGetInstanceWithDefaults instantiates a new GlusterVolumeGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterVolumeGetInstanceWithDefaults() *GlusterVolumeGetInstance {
	this := GlusterVolumeGetInstance{}
	var queryOptions GlusterVolumeGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GlusterVolumeGetInstance) GetId() GlusterVolumeGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret GlusterVolumeGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterVolumeGetInstance) GetIdOk() (*GlusterVolumeGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GlusterVolumeGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given GlusterVolumeGetInstance0 and assigns it to the Id field.
func (o *GlusterVolumeGetInstance) SetId(v GlusterVolumeGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *GlusterVolumeGetInstance) GetQueryOptions() GlusterVolumeGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret GlusterVolumeGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterVolumeGetInstance) GetQueryOptionsOk() (*GlusterVolumeGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *GlusterVolumeGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given GlusterVolumeGetInstance1 and assigns it to the QueryOptions field.
func (o *GlusterVolumeGetInstance) SetQueryOptions(v GlusterVolumeGetInstance1) {
	o.QueryOptions = &v
}

func (o GlusterVolumeGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlusterVolumeGetInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.QueryOptions) {
		toSerialize["query-options"] = o.QueryOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GlusterVolumeGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varGlusterVolumeGetInstance := _GlusterVolumeGetInstance{}

	if err = json.Unmarshal(bytes, &varGlusterVolumeGetInstance); err == nil {
		*o = GlusterVolumeGetInstance(varGlusterVolumeGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGlusterVolumeGetInstance struct {
	value *GlusterVolumeGetInstance
	isSet bool
}

func (v NullableGlusterVolumeGetInstance) Get() *GlusterVolumeGetInstance {
	return v.value
}

func (v *NullableGlusterVolumeGetInstance) Set(val *GlusterVolumeGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterVolumeGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterVolumeGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterVolumeGetInstance(val *GlusterVolumeGetInstance) *NullableGlusterVolumeGetInstance {
	return &NullableGlusterVolumeGetInstance{value: val, isSet: true}
}

func (v NullableGlusterVolumeGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterVolumeGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
