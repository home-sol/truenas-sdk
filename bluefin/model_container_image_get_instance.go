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

// checks if the ContainerImageGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerImageGetInstance{}

// ContainerImageGetInstance struct for ContainerImageGetInstance
type ContainerImageGetInstance struct {
	Id                   *ContainerImageGetInstance0 `json:"id,omitempty"`
	QueryOptions         *ContainerImageGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContainerImageGetInstance ContainerImageGetInstance

// NewContainerImageGetInstance instantiates a new ContainerImageGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerImageGetInstance() *ContainerImageGetInstance {
	this := ContainerImageGetInstance{}
	var queryOptions ContainerImageGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewContainerImageGetInstanceWithDefaults instantiates a new ContainerImageGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerImageGetInstanceWithDefaults() *ContainerImageGetInstance {
	this := ContainerImageGetInstance{}
	var queryOptions ContainerImageGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerImageGetInstance) GetId() ContainerImageGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret ContainerImageGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGetInstance) GetIdOk() (*ContainerImageGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerImageGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given ContainerImageGetInstance0 and assigns it to the Id field.
func (o *ContainerImageGetInstance) SetId(v ContainerImageGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *ContainerImageGetInstance) GetQueryOptions() ContainerImageGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret ContainerImageGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGetInstance) GetQueryOptionsOk() (*ContainerImageGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *ContainerImageGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given ContainerImageGetInstance1 and assigns it to the QueryOptions field.
func (o *ContainerImageGetInstance) SetQueryOptions(v ContainerImageGetInstance1) {
	o.QueryOptions = &v
}

func (o ContainerImageGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerImageGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *ContainerImageGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varContainerImageGetInstance := _ContainerImageGetInstance{}

	if err = json.Unmarshal(bytes, &varContainerImageGetInstance); err == nil {
		*o = ContainerImageGetInstance(varContainerImageGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainerImageGetInstance struct {
	value *ContainerImageGetInstance
	isSet bool
}

func (v NullableContainerImageGetInstance) Get() *ContainerImageGetInstance {
	return v.value
}

func (v *NullableContainerImageGetInstance) Set(val *ContainerImageGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerImageGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerImageGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerImageGetInstance(val *ContainerImageGetInstance) *NullableContainerImageGetInstance {
	return &NullableContainerImageGetInstance{value: val, isSet: true}
}

func (v NullableContainerImageGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerImageGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
