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

// checks if the PoolDatasetLock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolDatasetLock{}

// PoolDatasetLock struct for PoolDatasetLock
type PoolDatasetLock struct {
	// Locks `id` dataset. It will unmount the dataset and its children before locking.
	Id                   *string           `json:"id,omitempty"`
	LockOptions          *PoolDatasetLock1 `json:"lock_options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolDatasetLock PoolDatasetLock

// NewPoolDatasetLock instantiates a new PoolDatasetLock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetLock() *PoolDatasetLock {
	this := PoolDatasetLock{}
	var lockOptions PoolDatasetLock1
	this.LockOptions = &lockOptions
	return &this
}

// NewPoolDatasetLockWithDefaults instantiates a new PoolDatasetLock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetLockWithDefaults() *PoolDatasetLock {
	this := PoolDatasetLock{}
	var lockOptions PoolDatasetLock1
	this.LockOptions = &lockOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolDatasetLock) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetLock) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolDatasetLock) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PoolDatasetLock) SetId(v string) {
	o.Id = &v
}

// GetLockOptions returns the LockOptions field value if set, zero value otherwise.
func (o *PoolDatasetLock) GetLockOptions() PoolDatasetLock1 {
	if o == nil || IsNil(o.LockOptions) {
		var ret PoolDatasetLock1
		return ret
	}
	return *o.LockOptions
}

// GetLockOptionsOk returns a tuple with the LockOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetLock) GetLockOptionsOk() (*PoolDatasetLock1, bool) {
	if o == nil || IsNil(o.LockOptions) {
		return nil, false
	}
	return o.LockOptions, true
}

// HasLockOptions returns a boolean if a field has been set.
func (o *PoolDatasetLock) HasLockOptions() bool {
	if o != nil && !IsNil(o.LockOptions) {
		return true
	}

	return false
}

// SetLockOptions gets a reference to the given PoolDatasetLock1 and assigns it to the LockOptions field.
func (o *PoolDatasetLock) SetLockOptions(v PoolDatasetLock1) {
	o.LockOptions = &v
}

func (o PoolDatasetLock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolDatasetLock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LockOptions) {
		toSerialize["lock_options"] = o.LockOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PoolDatasetLock) UnmarshalJSON(bytes []byte) (err error) {
	varPoolDatasetLock := _PoolDatasetLock{}

	if err = json.Unmarshal(bytes, &varPoolDatasetLock); err == nil {
		*o = PoolDatasetLock(varPoolDatasetLock)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "lock_options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolDatasetLock struct {
	value *PoolDatasetLock
	isSet bool
}

func (v NullablePoolDatasetLock) Get() *PoolDatasetLock {
	return v.value
}

func (v *NullablePoolDatasetLock) Set(val *PoolDatasetLock) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetLock) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetLock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetLock(val *PoolDatasetLock) *NullablePoolDatasetLock {
	return &NullablePoolDatasetLock{value: val, isSet: true}
}

func (v NullablePoolDatasetLock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetLock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}