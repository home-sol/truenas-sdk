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

// checks if the SmartTestGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartTestGetInstance{}

// SmartTestGetInstance struct for SmartTestGetInstance
type SmartTestGetInstance struct {
	Id                   *SmartTestGetInstance0 `json:"id,omitempty"`
	QueryOptions         *SmartTestGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SmartTestGetInstance SmartTestGetInstance

// NewSmartTestGetInstance instantiates a new SmartTestGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartTestGetInstance() *SmartTestGetInstance {
	this := SmartTestGetInstance{}
	var queryOptions SmartTestGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewSmartTestGetInstanceWithDefaults instantiates a new SmartTestGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartTestGetInstanceWithDefaults() *SmartTestGetInstance {
	this := SmartTestGetInstance{}
	var queryOptions SmartTestGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SmartTestGetInstance) GetId() SmartTestGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret SmartTestGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartTestGetInstance) GetIdOk() (*SmartTestGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SmartTestGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given SmartTestGetInstance0 and assigns it to the Id field.
func (o *SmartTestGetInstance) SetId(v SmartTestGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *SmartTestGetInstance) GetQueryOptions() SmartTestGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret SmartTestGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartTestGetInstance) GetQueryOptionsOk() (*SmartTestGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *SmartTestGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given SmartTestGetInstance1 and assigns it to the QueryOptions field.
func (o *SmartTestGetInstance) SetQueryOptions(v SmartTestGetInstance1) {
	o.QueryOptions = &v
}

func (o SmartTestGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartTestGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *SmartTestGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varSmartTestGetInstance := _SmartTestGetInstance{}

	if err = json.Unmarshal(bytes, &varSmartTestGetInstance); err == nil {
		*o = SmartTestGetInstance(varSmartTestGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSmartTestGetInstance struct {
	value *SmartTestGetInstance
	isSet bool
}

func (v NullableSmartTestGetInstance) Get() *SmartTestGetInstance {
	return v.value
}

func (v *NullableSmartTestGetInstance) Set(val *SmartTestGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartTestGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartTestGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartTestGetInstance(val *SmartTestGetInstance) *NullableSmartTestGetInstance {
	return &NullableSmartTestGetInstance{value: val, isSet: true}
}

func (v NullableSmartTestGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartTestGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}