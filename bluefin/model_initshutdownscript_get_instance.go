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

// checks if the InitshutdownscriptGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitshutdownscriptGetInstance{}

// InitshutdownscriptGetInstance struct for InitshutdownscriptGetInstance
type InitshutdownscriptGetInstance struct {
	Id                   *InitshutdownscriptGetInstance0 `json:"id,omitempty"`
	QueryOptions         *InitshutdownscriptGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InitshutdownscriptGetInstance InitshutdownscriptGetInstance

// NewInitshutdownscriptGetInstance instantiates a new InitshutdownscriptGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitshutdownscriptGetInstance() *InitshutdownscriptGetInstance {
	this := InitshutdownscriptGetInstance{}
	var queryOptions InitshutdownscriptGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewInitshutdownscriptGetInstanceWithDefaults instantiates a new InitshutdownscriptGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitshutdownscriptGetInstanceWithDefaults() *InitshutdownscriptGetInstance {
	this := InitshutdownscriptGetInstance{}
	var queryOptions InitshutdownscriptGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InitshutdownscriptGetInstance) GetId() InitshutdownscriptGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret InitshutdownscriptGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitshutdownscriptGetInstance) GetIdOk() (*InitshutdownscriptGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InitshutdownscriptGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given InitshutdownscriptGetInstance0 and assigns it to the Id field.
func (o *InitshutdownscriptGetInstance) SetId(v InitshutdownscriptGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *InitshutdownscriptGetInstance) GetQueryOptions() InitshutdownscriptGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret InitshutdownscriptGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitshutdownscriptGetInstance) GetQueryOptionsOk() (*InitshutdownscriptGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *InitshutdownscriptGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given InitshutdownscriptGetInstance1 and assigns it to the QueryOptions field.
func (o *InitshutdownscriptGetInstance) SetQueryOptions(v InitshutdownscriptGetInstance1) {
	o.QueryOptions = &v
}

func (o InitshutdownscriptGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitshutdownscriptGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *InitshutdownscriptGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varInitshutdownscriptGetInstance := _InitshutdownscriptGetInstance{}

	if err = json.Unmarshal(bytes, &varInitshutdownscriptGetInstance); err == nil {
		*o = InitshutdownscriptGetInstance(varInitshutdownscriptGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInitshutdownscriptGetInstance struct {
	value *InitshutdownscriptGetInstance
	isSet bool
}

func (v NullableInitshutdownscriptGetInstance) Get() *InitshutdownscriptGetInstance {
	return v.value
}

func (v *NullableInitshutdownscriptGetInstance) Set(val *InitshutdownscriptGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableInitshutdownscriptGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInitshutdownscriptGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitshutdownscriptGetInstance(val *InitshutdownscriptGetInstance) *NullableInitshutdownscriptGetInstance {
	return &NullableInitshutdownscriptGetInstance{value: val, isSet: true}
}

func (v NullableInitshutdownscriptGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitshutdownscriptGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
