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

// checks if the ServiceAnnouncement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAnnouncement{}

// ServiceAnnouncement struct for ServiceAnnouncement
type ServiceAnnouncement struct {
	Netbios *bool `json:"netbios,omitempty"`
	Mdns    *bool `json:"mdns,omitempty"`
	Wsd     *bool `json:"wsd,omitempty"`
}

// NewServiceAnnouncement instantiates a new ServiceAnnouncement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAnnouncement() *ServiceAnnouncement {
	this := ServiceAnnouncement{}
	return &this
}

// NewServiceAnnouncementWithDefaults instantiates a new ServiceAnnouncement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAnnouncementWithDefaults() *ServiceAnnouncement {
	this := ServiceAnnouncement{}
	return &this
}

// GetNetbios returns the Netbios field value if set, zero value otherwise.
func (o *ServiceAnnouncement) GetNetbios() bool {
	if o == nil || IsNil(o.Netbios) {
		var ret bool
		return ret
	}
	return *o.Netbios
}

// GetNetbiosOk returns a tuple with the Netbios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAnnouncement) GetNetbiosOk() (*bool, bool) {
	if o == nil || IsNil(o.Netbios) {
		return nil, false
	}
	return o.Netbios, true
}

// HasNetbios returns a boolean if a field has been set.
func (o *ServiceAnnouncement) HasNetbios() bool {
	if o != nil && !IsNil(o.Netbios) {
		return true
	}

	return false
}

// SetNetbios gets a reference to the given bool and assigns it to the Netbios field.
func (o *ServiceAnnouncement) SetNetbios(v bool) {
	o.Netbios = &v
}

// GetMdns returns the Mdns field value if set, zero value otherwise.
func (o *ServiceAnnouncement) GetMdns() bool {
	if o == nil || IsNil(o.Mdns) {
		var ret bool
		return ret
	}
	return *o.Mdns
}

// GetMdnsOk returns a tuple with the Mdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAnnouncement) GetMdnsOk() (*bool, bool) {
	if o == nil || IsNil(o.Mdns) {
		return nil, false
	}
	return o.Mdns, true
}

// HasMdns returns a boolean if a field has been set.
func (o *ServiceAnnouncement) HasMdns() bool {
	if o != nil && !IsNil(o.Mdns) {
		return true
	}

	return false
}

// SetMdns gets a reference to the given bool and assigns it to the Mdns field.
func (o *ServiceAnnouncement) SetMdns(v bool) {
	o.Mdns = &v
}

// GetWsd returns the Wsd field value if set, zero value otherwise.
func (o *ServiceAnnouncement) GetWsd() bool {
	if o == nil || IsNil(o.Wsd) {
		var ret bool
		return ret
	}
	return *o.Wsd
}

// GetWsdOk returns a tuple with the Wsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAnnouncement) GetWsdOk() (*bool, bool) {
	if o == nil || IsNil(o.Wsd) {
		return nil, false
	}
	return o.Wsd, true
}

// HasWsd returns a boolean if a field has been set.
func (o *ServiceAnnouncement) HasWsd() bool {
	if o != nil && !IsNil(o.Wsd) {
		return true
	}

	return false
}

// SetWsd gets a reference to the given bool and assigns it to the Wsd field.
func (o *ServiceAnnouncement) SetWsd(v bool) {
	o.Wsd = &v
}

func (o ServiceAnnouncement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAnnouncement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Netbios) {
		toSerialize["netbios"] = o.Netbios
	}
	if !IsNil(o.Mdns) {
		toSerialize["mdns"] = o.Mdns
	}
	if !IsNil(o.Wsd) {
		toSerialize["wsd"] = o.Wsd
	}
	return toSerialize, nil
}

type NullableServiceAnnouncement struct {
	value *ServiceAnnouncement
	isSet bool
}

func (v NullableServiceAnnouncement) Get() *ServiceAnnouncement {
	return v.value
}

func (v *NullableServiceAnnouncement) Set(val *ServiceAnnouncement) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAnnouncement) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAnnouncement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAnnouncement(val *ServiceAnnouncement) *NullableServiceAnnouncement {
	return &NullableServiceAnnouncement{value: val, isSet: true}
}

func (v NullableServiceAnnouncement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAnnouncement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}