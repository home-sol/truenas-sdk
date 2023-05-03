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

// checks if the PoolSnapshottaskGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolSnapshottaskGetInstance{}

// PoolSnapshottaskGetInstance struct for PoolSnapshottaskGetInstance
type PoolSnapshottaskGetInstance struct {
	Id                   *PoolSnapshottaskGetInstance0 `json:"id,omitempty"`
	QueryOptions         *PoolSnapshottaskGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolSnapshottaskGetInstance PoolSnapshottaskGetInstance

// NewPoolSnapshottaskGetInstance instantiates a new PoolSnapshottaskGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolSnapshottaskGetInstance() *PoolSnapshottaskGetInstance {
	this := PoolSnapshottaskGetInstance{}
	var queryOptions PoolSnapshottaskGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewPoolSnapshottaskGetInstanceWithDefaults instantiates a new PoolSnapshottaskGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolSnapshottaskGetInstanceWithDefaults() *PoolSnapshottaskGetInstance {
	this := PoolSnapshottaskGetInstance{}
	var queryOptions PoolSnapshottaskGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolSnapshottaskGetInstance) GetId() PoolSnapshottaskGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret PoolSnapshottaskGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSnapshottaskGetInstance) GetIdOk() (*PoolSnapshottaskGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolSnapshottaskGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given PoolSnapshottaskGetInstance0 and assigns it to the Id field.
func (o *PoolSnapshottaskGetInstance) SetId(v PoolSnapshottaskGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *PoolSnapshottaskGetInstance) GetQueryOptions() PoolSnapshottaskGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret PoolSnapshottaskGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSnapshottaskGetInstance) GetQueryOptionsOk() (*PoolSnapshottaskGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *PoolSnapshottaskGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given PoolSnapshottaskGetInstance1 and assigns it to the QueryOptions field.
func (o *PoolSnapshottaskGetInstance) SetQueryOptions(v PoolSnapshottaskGetInstance1) {
	o.QueryOptions = &v
}

func (o PoolSnapshottaskGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolSnapshottaskGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *PoolSnapshottaskGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varPoolSnapshottaskGetInstance := _PoolSnapshottaskGetInstance{}

	if err = json.Unmarshal(bytes, &varPoolSnapshottaskGetInstance); err == nil {
		*o = PoolSnapshottaskGetInstance(varPoolSnapshottaskGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolSnapshottaskGetInstance struct {
	value *PoolSnapshottaskGetInstance
	isSet bool
}

func (v NullablePoolSnapshottaskGetInstance) Get() *PoolSnapshottaskGetInstance {
	return v.value
}

func (v *NullablePoolSnapshottaskGetInstance) Set(val *PoolSnapshottaskGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolSnapshottaskGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolSnapshottaskGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolSnapshottaskGetInstance(val *PoolSnapshottaskGetInstance) *NullablePoolSnapshottaskGetInstance {
	return &NullablePoolSnapshottaskGetInstance{value: val, isSet: true}
}

func (v NullablePoolSnapshottaskGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolSnapshottaskGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}