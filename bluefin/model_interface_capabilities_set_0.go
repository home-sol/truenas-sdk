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

// checks if the InterfaceCapabilitiesSet0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceCapabilitiesSet0{}

// InterfaceCapabilitiesSet0 struct for InterfaceCapabilitiesSet0
type InterfaceCapabilitiesSet0 struct {
	// `name` String representing name of the interface `capabilities` List representing capabilities to be acted upon
	Name        *string       `json:"name,omitempty"`
	Capabilties []interface{} `json:"capabilties,omitempty"`
	Action      *string       `json:"action,omitempty"`
}

// NewInterfaceCapabilitiesSet0 instantiates a new InterfaceCapabilitiesSet0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceCapabilitiesSet0() *InterfaceCapabilitiesSet0 {
	this := InterfaceCapabilitiesSet0{}
	return &this
}

// NewInterfaceCapabilitiesSet0WithDefaults instantiates a new InterfaceCapabilitiesSet0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceCapabilitiesSet0WithDefaults() *InterfaceCapabilitiesSet0 {
	this := InterfaceCapabilitiesSet0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InterfaceCapabilitiesSet0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceCapabilitiesSet0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InterfaceCapabilitiesSet0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InterfaceCapabilitiesSet0) SetName(v string) {
	o.Name = &v
}

// GetCapabilties returns the Capabilties field value if set, zero value otherwise.
func (o *InterfaceCapabilitiesSet0) GetCapabilties() []interface{} {
	if o == nil || IsNil(o.Capabilties) {
		var ret []interface{}
		return ret
	}
	return o.Capabilties
}

// GetCapabiltiesOk returns a tuple with the Capabilties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceCapabilitiesSet0) GetCapabiltiesOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Capabilties) {
		return nil, false
	}
	return o.Capabilties, true
}

// HasCapabilties returns a boolean if a field has been set.
func (o *InterfaceCapabilitiesSet0) HasCapabilties() bool {
	if o != nil && !IsNil(o.Capabilties) {
		return true
	}

	return false
}

// SetCapabilties gets a reference to the given []interface{} and assigns it to the Capabilties field.
func (o *InterfaceCapabilitiesSet0) SetCapabilties(v []interface{}) {
	o.Capabilties = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *InterfaceCapabilitiesSet0) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceCapabilitiesSet0) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *InterfaceCapabilitiesSet0) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *InterfaceCapabilitiesSet0) SetAction(v string) {
	o.Action = &v
}

func (o InterfaceCapabilitiesSet0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceCapabilitiesSet0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Capabilties) {
		toSerialize["capabilties"] = o.Capabilties
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableInterfaceCapabilitiesSet0 struct {
	value *InterfaceCapabilitiesSet0
	isSet bool
}

func (v NullableInterfaceCapabilitiesSet0) Get() *InterfaceCapabilitiesSet0 {
	return v.value
}

func (v *NullableInterfaceCapabilitiesSet0) Set(val *InterfaceCapabilitiesSet0) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceCapabilitiesSet0) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceCapabilitiesSet0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceCapabilitiesSet0(val *InterfaceCapabilitiesSet0) *NullableInterfaceCapabilitiesSet0 {
	return &NullableInterfaceCapabilitiesSet0{value: val, isSet: true}
}

func (v NullableInterfaceCapabilitiesSet0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceCapabilitiesSet0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}