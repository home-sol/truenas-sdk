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

// checks if the AlertserviceGetInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertserviceGetInstance{}

// AlertserviceGetInstance struct for AlertserviceGetInstance
type AlertserviceGetInstance struct {
	Id                   *AlertserviceGetInstance0 `json:"id,omitempty"`
	QueryOptions         *AlertserviceGetInstance1 `json:"query-options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlertserviceGetInstance AlertserviceGetInstance

// NewAlertserviceGetInstance instantiates a new AlertserviceGetInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertserviceGetInstance() *AlertserviceGetInstance {
	this := AlertserviceGetInstance{}
	var queryOptions AlertserviceGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// NewAlertserviceGetInstanceWithDefaults instantiates a new AlertserviceGetInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertserviceGetInstanceWithDefaults() *AlertserviceGetInstance {
	this := AlertserviceGetInstance{}
	var queryOptions AlertserviceGetInstance1
	this.QueryOptions = &queryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertserviceGetInstance) GetId() AlertserviceGetInstance0 {
	if o == nil || IsNil(o.Id) {
		var ret AlertserviceGetInstance0
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertserviceGetInstance) GetIdOk() (*AlertserviceGetInstance0, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertserviceGetInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given AlertserviceGetInstance0 and assigns it to the Id field.
func (o *AlertserviceGetInstance) SetId(v AlertserviceGetInstance0) {
	o.Id = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *AlertserviceGetInstance) GetQueryOptions() AlertserviceGetInstance1 {
	if o == nil || IsNil(o.QueryOptions) {
		var ret AlertserviceGetInstance1
		return ret
	}
	return *o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertserviceGetInstance) GetQueryOptionsOk() (*AlertserviceGetInstance1, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *AlertserviceGetInstance) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given AlertserviceGetInstance1 and assigns it to the QueryOptions field.
func (o *AlertserviceGetInstance) SetQueryOptions(v AlertserviceGetInstance1) {
	o.QueryOptions = &v
}

func (o AlertserviceGetInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertserviceGetInstance) ToMap() (map[string]interface{}, error) {
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

func (o *AlertserviceGetInstance) UnmarshalJSON(bytes []byte) (err error) {
	varAlertserviceGetInstance := _AlertserviceGetInstance{}

	if err = json.Unmarshal(bytes, &varAlertserviceGetInstance); err == nil {
		*o = AlertserviceGetInstance(varAlertserviceGetInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "query-options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlertserviceGetInstance struct {
	value *AlertserviceGetInstance
	isSet bool
}

func (v NullableAlertserviceGetInstance) Get() *AlertserviceGetInstance {
	return v.value
}

func (v *NullableAlertserviceGetInstance) Set(val *AlertserviceGetInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertserviceGetInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertserviceGetInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertserviceGetInstance(val *AlertserviceGetInstance) *NullableAlertserviceGetInstance {
	return &NullableAlertserviceGetInstance{value: val, isSet: true}
}

func (v NullableAlertserviceGetInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertserviceGetInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
