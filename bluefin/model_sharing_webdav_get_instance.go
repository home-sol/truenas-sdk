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

// checks if the SharingWebdavGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharingWebdavGetInstance{}

// SharingWebdavGetInstance struct for SharingWebdavGetInstance
type SharingWebdavGetInstance struct {
	Id                   *SharingWebdavGetInstance0 `json:"id,omitempty"`
	QueryOptions         *SharingWebdavGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SharingWebdavGetInstance SharingWebdavGetInstance

// NewSharingWebdavGetInstance instantiates a new SharingWebdavGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharingWebdavGetInstance() *SharingWebdavGetInstance {
	this := SharingWebdavGetInstance{}
	var queryOptions SharingWebdavGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewSharingWebdavGetInstanceWithDefaults instantiates a new SharingWebdavGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharingWebdavGetInstanceWithDefaults() *SharingWebdavGetInstance {
	this := SharingWebdavGetInstance{}
	var queryOptions SharingWebdavGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SharingWebdavGetInstance) GetId() SharingWebdavGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret SharingWebdavGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingWebdavGetInstance) GetIdOk() (*SharingWebdavGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SharingWebdavGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given SharingWebdavGetInstance0 and assigns it to the Id field.
func (o *SharingWebdavGetInstance) SetId(v SharingWebdavGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *SharingWebdavGetInstance) GetQueryOptions() SharingWebdavGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret SharingWebdavGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingWebdavGetInstance) GetQueryOptionsOk() (*SharingWebdavGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *SharingWebdavGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given SharingWebdavGetInstance1 and assigns it to the QueryOptions field.
func (o *SharingWebdavGetInstance) SetQueryOptions(v SharingWebdavGetInstance1) {
	o.QueryOptions = &v
}

func (o SharingWebdavGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharingWebdavGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *SharingWebdavGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varSharingWebdavGetInstance := _SharingWebdavGetInstance{}

	if err = json.Unmarshal(bytes, &varSharingWebdavGetInstance); err == nil {
		*o = SharingWebdavGetInstance(varSharingWebdavGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSharingWebdavGetInstance struct {
	value *SharingWebdavGetInstance
	isSet bool
}

func (v NullableSharingWebdavGetInstance) Get() *SharingWebdavGetInstance {
	return v.value
}

func (v *NullableSharingWebdavGetInstance) Set(val *SharingWebdavGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingWebdavGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingWebdavGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingWebdavGetInstance(val *SharingWebdavGetInstance) *NullableSharingWebdavGetInstance {
	return &NullableSharingWebdavGetInstance{value: val, isSet: true}
}

func (v NullableSharingWebdavGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingWebdavGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
