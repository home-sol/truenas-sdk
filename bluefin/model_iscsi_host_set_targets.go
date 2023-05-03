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

// checks if the IscsiHostSetTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IscsiHostSetTargets{}

// IscsiHostSetTargets struct for IscsiHostSetTargets
type IscsiHostSetTargets struct {
	Id                   *int32  `json:"id,omitempty"`
	Ids                  []int32 `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IscsiHostSetTargets IscsiHostSetTargets

// NewIscsiHostSetTargets instantiates a new IscsiHostSetTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiHostSetTargets() *IscsiHostSetTargets {
	this := IscsiHostSetTargets{}
	return &this
}

// NewIscsiHostSetTargetsWithDefaults instantiates a new IscsiHostSetTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiHostSetTargetsWithDefaults() *IscsiHostSetTargets {
	this := IscsiHostSetTargets{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IscsiHostSetTargets) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiHostSetTargets) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IscsiHostSetTargets) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IscsiHostSetTargets) SetId(v int32) {
	o.Id = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *IscsiHostSetTargets) GetIds() []int32 {
	if o == nil || IsNil(o.Ids) {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiHostSetTargets) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *IscsiHostSetTargets) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *IscsiHostSetTargets) SetIds(v []int32) {
	o.Ids = v
}

func (o IscsiHostSetTargets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IscsiHostSetTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IscsiHostSetTargets) UnmarshalJSON(bytes []byte) (err error) {
	varIscsiHostSetTargets := _IscsiHostSetTargets{}

	if err = json.Unmarshal(bytes, &varIscsiHostSetTargets); err == nil {
		*o = IscsiHostSetTargets(varIscsiHostSetTargets)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIscsiHostSetTargets struct {
	value *IscsiHostSetTargets
	isSet bool
}

func (v NullableIscsiHostSetTargets) Get() *IscsiHostSetTargets {
	return v.value
}

func (v *NullableIscsiHostSetTargets) Set(val *IscsiHostSetTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiHostSetTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiHostSetTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiHostSetTargets(val *IscsiHostSetTargets) *NullableIscsiHostSetTargets {
	return &NullableIscsiHostSetTargets{value: val, isSet: true}
}

func (v NullableIscsiHostSetTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiHostSetTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
