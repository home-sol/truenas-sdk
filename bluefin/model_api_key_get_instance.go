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

// checks if the ApiKeyGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyGetInstance{}

// ApiKeyGetInstance struct for ApiKeyGetInstance
type ApiKeyGetInstance struct {
	Id                   *ApiKeyGetInstance0 `json:"id,omitempty"`
	QueryOptions         *ApiKeyGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiKeyGetInstance ApiKeyGetInstance

// NewApiKeyGetInstance instantiates a new ApiKeyGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyGetInstance() *ApiKeyGetInstance {
	this := ApiKeyGetInstance{}
	var queryOptions ApiKeyGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewApiKeyGetInstanceWithDefaults instantiates a new ApiKeyGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyGetInstanceWithDefaults() *ApiKeyGetInstance {
	this := ApiKeyGetInstance{}
	var queryOptions ApiKeyGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiKeyGetInstance) GetId() ApiKeyGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret ApiKeyGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyGetInstance) GetIdOk() (*ApiKeyGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiKeyGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given ApiKeyGetInstance0 and assigns it to the Id field.
func (o *ApiKeyGetInstance) SetId(v ApiKeyGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *ApiKeyGetInstance) GetQueryOptions() ApiKeyGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret ApiKeyGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyGetInstance) GetQueryOptionsOk() (*ApiKeyGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *ApiKeyGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given ApiKeyGetInstance1 and assigns it to the QueryOptions field.
func (o *ApiKeyGetInstance) SetQueryOptions(v ApiKeyGetInstance1) {
	o.QueryOptions = &v
}

func (o ApiKeyGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *ApiKeyGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varApiKeyGetInstance := _ApiKeyGetInstance{}

	if err = json.Unmarshal(bytes, &varApiKeyGetInstance); err == nil {
		*o = ApiKeyGetInstance(varApiKeyGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiKeyGetInstance struct {
	value *ApiKeyGetInstance
	isSet bool
}

func (v NullableApiKeyGetInstance) Get() *ApiKeyGetInstance {
	return v.value
}

func (v *NullableApiKeyGetInstance) Set(val *ApiKeyGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyGetInstance(val *ApiKeyGetInstance) *NullableApiKeyGetInstance {
	return &NullableApiKeyGetInstance{value: val, isSet: true}
}

func (v NullableApiKeyGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}