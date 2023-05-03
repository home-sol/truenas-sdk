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

// checks if the DiskRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskRun{}

// DiskRun struct for DiskRun
type DiskRun struct {
	Identifier *string `json:"identifier,omitempty"`
	Mode       *string `json:"mode,omitempty"`
	Type       *string `json:"type,omitempty"`
}

// NewDiskRun instantiates a new DiskRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskRun() *DiskRun {
	this := DiskRun{}
	var mode string
	this.Mode = &mode
	return &this
}

// NewDiskRunWithDefaults instantiates a new DiskRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskRunWithDefaults() *DiskRun {
	this := DiskRun{}
	var mode string
	this.Mode = &mode
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *DiskRun) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRun) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *DiskRun) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *DiskRun) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DiskRun) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRun) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DiskRun) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DiskRun) SetMode(v string) {
	o.Mode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DiskRun) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRun) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DiskRun) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DiskRun) SetType(v string) {
	o.Type = &v
}

func (o DiskRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDiskRun struct {
	value *DiskRun
	isSet bool
}

func (v NullableDiskRun) Get() *DiskRun {
	return v.value
}

func (v *NullableDiskRun) Set(val *DiskRun) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskRun) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskRun(val *DiskRun) *NullableDiskRun {
	return &NullableDiskRun{value: val, isSet: true}
}

func (v NullableDiskRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}