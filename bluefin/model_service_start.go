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

// checks if the ServiceStart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceStart{}

// ServiceStart struct for ServiceStart
type ServiceStart struct {
	// Start the service specified by `service`.
	Service              *string        `json:"service,omitempty"`
	ServiceControl       *ServiceStart1 `json:"service-control,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceStart ServiceStart

// NewServiceStart instantiates a new ServiceStart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStart() *ServiceStart {
	this := ServiceStart{}
	var serviceControl ServiceStart1
	this.ServiceControl = &serviceControl
	return &this
}

// NewServiceStartWithDefaults instantiates a new ServiceStart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStartWithDefaults() *ServiceStart {
	this := ServiceStart{}
	var serviceControl ServiceStart1
	this.ServiceControl = &serviceControl
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceStart) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStart) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceStart) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceStart) SetService(v string) {
	o.Service = &v
}

// GetServiceControl returns the ServiceControl field value if set, zero value otherwise.
func (o *ServiceStart) GetServiceControl() ServiceStart1 {
	if o == nil || IsNil(o.ServiceControl) {
		var ret ServiceStart1
		return ret
	}
	return *o.ServiceControl
}

// GetServiceControlOk returns a tuple with the ServiceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStart) GetServiceControlOk() (*ServiceStart1, bool) {
	if o == nil || IsNil(o.ServiceControl) {
		return nil, false
	}
	return o.ServiceControl, true
}

// HasServiceControl returns a boolean if a field has been set.
func (o *ServiceStart) HasServiceControl() bool {
	if o != nil && !IsNil(o.ServiceControl) {
		return true
	}

	return false
}

// SetServiceControl gets a reference to the given ServiceStart1 and assigns it to the ServiceControl field.
func (o *ServiceStart) SetServiceControl(v ServiceStart1) {
	o.ServiceControl = &v
}

func (o ServiceStart) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceStart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.ServiceControl) {
		toSerialize["service-control"] = o.ServiceControl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceStart) UnmarshalJSON(bytes []byte) (err error) {
	varServiceStart := _ServiceStart{}

	if err = json.Unmarshal(bytes, &varServiceStart); err == nil {
		*o = ServiceStart(varServiceStart)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		delete(additionalProperties, "service-control")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceStart struct {
	value *ServiceStart
	isSet bool
}

func (v NullableServiceStart) Get() *ServiceStart {
	return v.value
}

func (v *NullableServiceStart) Set(val *ServiceStart) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStart) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStart(val *ServiceStart) *NullableServiceStart {
	return &NullableServiceStart{value: val, isSet: true}
}

func (v NullableServiceStart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}