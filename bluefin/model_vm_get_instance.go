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

// checks if the VmGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmGetInstance{}

// VmGetInstance struct for VmGetInstance
type VmGetInstance struct {
	Id                   *VmGetInstance0 `json:"id,omitempty"`
	QueryOptions         *VmGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VmGetInstance VmGetInstance

// NewVmGetInstance instantiates a new VmGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmGetInstance() *VmGetInstance {
	this := VmGetInstance{}
	var queryOptions VmGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewVmGetInstanceWithDefaults instantiates a new VmGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmGetInstanceWithDefaults() *VmGetInstance {
	this := VmGetInstance{}
	var queryOptions VmGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VmGetInstance) GetId() VmGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret VmGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGetInstance) GetIdOk() (*VmGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VmGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given VmGetInstance0 and assigns it to the Id field.
func (o *VmGetInstance) SetId(v VmGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *VmGetInstance) GetQueryOptions() VmGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret VmGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGetInstance) GetQueryOptionsOk() (*VmGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *VmGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given VmGetInstance1 and assigns it to the QueryOptions field.
func (o *VmGetInstance) SetQueryOptions(v VmGetInstance1) {
	o.QueryOptions = &v
}

func (o VmGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *VmGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varVmGetInstance := _VmGetInstance{}

	if err = json.Unmarshal(bytes, &varVmGetInstance); err == nil {
		*o = VmGetInstance(varVmGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVmGetInstance struct {
	value *VmGetInstance
	isSet bool
}

func (v NullableVmGetInstance) Get() *VmGetInstance {
	return v.value
}

func (v *NullableVmGetInstance) Set(val *VmGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableVmGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableVmGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmGetInstance(val *VmGetInstance) *NullableVmGetInstance {
	return &NullableVmGetInstance{value: val, isSet: true}
}

func (v NullableVmGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
