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

// checks if the KerberosRealmGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KerberosRealmGetInstance{}

// KerberosRealmGetInstance struct for KerberosRealmGetInstance
type KerberosRealmGetInstance struct {
	Id                   *KerberosRealmGetInstance0 `json:"id,omitempty"`
	QueryOptions         *KerberosRealmGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KerberosRealmGetInstance KerberosRealmGetInstance

// NewKerberosRealmGetInstance instantiates a new KerberosRealmGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosRealmGetInstance() *KerberosRealmGetInstance {
	this := KerberosRealmGetInstance{}
	var queryOptions KerberosRealmGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewKerberosRealmGetInstanceWithDefaults instantiates a new KerberosRealmGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosRealmGetInstanceWithDefaults() *KerberosRealmGetInstance {
	this := KerberosRealmGetInstance{}
	var queryOptions KerberosRealmGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KerberosRealmGetInstance) GetId() KerberosRealmGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret KerberosRealmGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealmGetInstance) GetIdOk() (*KerberosRealmGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KerberosRealmGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given KerberosRealmGetInstance0 and assigns it to the Id field.
func (o *KerberosRealmGetInstance) SetId(v KerberosRealmGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *KerberosRealmGetInstance) GetQueryOptions() KerberosRealmGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret KerberosRealmGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealmGetInstance) GetQueryOptionsOk() (*KerberosRealmGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *KerberosRealmGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given KerberosRealmGetInstance1 and assigns it to the QueryOptions field.
func (o *KerberosRealmGetInstance) SetQueryOptions(v KerberosRealmGetInstance1) {
	o.QueryOptions = &v
}

func (o KerberosRealmGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KerberosRealmGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *KerberosRealmGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varKerberosRealmGetInstance := _KerberosRealmGetInstance{}

	if err = json.Unmarshal(bytes, &varKerberosRealmGetInstance); err == nil {
		*o = KerberosRealmGetInstance(varKerberosRealmGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKerberosRealmGetInstance struct {
	value *KerberosRealmGetInstance
	isSet bool
}

func (v NullableKerberosRealmGetInstance) Get() *KerberosRealmGetInstance {
	return v.value
}

func (v *NullableKerberosRealmGetInstance) Set(val *KerberosRealmGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosRealmGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosRealmGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosRealmGetInstance(val *KerberosRealmGetInstance) *NullableKerberosRealmGetInstance {
	return &NullableKerberosRealmGetInstance{value: val, isSet: true}
}

func (v NullableKerberosRealmGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosRealmGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
