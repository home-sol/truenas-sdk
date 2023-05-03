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

// checks if the ZfsSnapshotRollback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZfsSnapshotRollback{}

// ZfsSnapshotRollback struct for ZfsSnapshotRollback
type ZfsSnapshotRollback struct {
	// Rollback to a given snapshot `id`.
	Id                   *string               `json:"id,omitempty"`
	Options              *ZfsSnapshotRollback1 `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ZfsSnapshotRollback ZfsSnapshotRollback

// NewZfsSnapshotRollback instantiates a new ZfsSnapshotRollback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZfsSnapshotRollback() *ZfsSnapshotRollback {
	this := ZfsSnapshotRollback{}
	var options ZfsSnapshotRollback1
	this.Options = &options
	return &this
}

// NewZfsSnapshotRollbackWithDefaults instantiates a new ZfsSnapshotRollback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZfsSnapshotRollbackWithDefaults() *ZfsSnapshotRollback {
	this := ZfsSnapshotRollback{}
	var options ZfsSnapshotRollback1
	this.Options = &options
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZfsSnapshotRollback) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotRollback) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZfsSnapshotRollback) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ZfsSnapshotRollback) SetId(v string) {
	o.Id = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ZfsSnapshotRollback) GetOptions() ZfsSnapshotRollback1 {
	if o == nil || IsNil(o.Options) {
		var ret ZfsSnapshotRollback1
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotRollback) GetOptionsOk() (*ZfsSnapshotRollback1, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ZfsSnapshotRollback) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ZfsSnapshotRollback1 and assigns it to the Options field.
func (o *ZfsSnapshotRollback) SetOptions(v ZfsSnapshotRollback1) {
	o.Options = &v
}

func (o ZfsSnapshotRollback) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZfsSnapshotRollback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ZfsSnapshotRollback) UnmarshalJSON(bytes []byte) (err error) {
	varZfsSnapshotRollback := _ZfsSnapshotRollback{}

	if err = json.Unmarshal(bytes, &varZfsSnapshotRollback); err == nil {
		*o = ZfsSnapshotRollback(varZfsSnapshotRollback)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableZfsSnapshotRollback struct {
	value *ZfsSnapshotRollback
	isSet bool
}

func (v NullableZfsSnapshotRollback) Get() *ZfsSnapshotRollback {
	return v.value
}

func (v *NullableZfsSnapshotRollback) Set(val *ZfsSnapshotRollback) {
	v.value = val
	v.isSet = true
}

func (v NullableZfsSnapshotRollback) IsSet() bool {
	return v.isSet
}

func (v *NullableZfsSnapshotRollback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZfsSnapshotRollback(val *ZfsSnapshotRollback) *NullableZfsSnapshotRollback {
	return &NullableZfsSnapshotRollback{value: val, isSet: true}
}

func (v NullableZfsSnapshotRollback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZfsSnapshotRollback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}