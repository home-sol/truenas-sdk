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

// checks if the Dosmode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dosmode{}

// Dosmode struct for Dosmode
type Dosmode struct {
	Readonly *bool `json:"readonly,omitempty"`
	Hidden   *bool `json:"hidden,omitempty"`
	System   *bool `json:"system,omitempty"`
	Archive  *bool `json:"archive,omitempty"`
	Reparse  *bool `json:"reparse,omitempty"`
	Offline  *bool `json:"offline,omitempty"`
	Sparse   *bool `json:"sparse,omitempty"`
}

// NewDosmode instantiates a new Dosmode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDosmode() *Dosmode {
	this := Dosmode{}
	return &this
}

// NewDosmodeWithDefaults instantiates a new Dosmode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDosmodeWithDefaults() *Dosmode {
	this := Dosmode{}
	return &this
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *Dosmode) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dosmode) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *Dosmode) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *Dosmode) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *Dosmode) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dosmode) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *Dosmode) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *Dosmode) SetHidden(v bool) {
	o.Hidden = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *Dosmode) GetSystem() bool {
	if o == nil || IsNil(o.System) {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dosmode) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *Dosmode) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *Dosmode) SetSystem(v bool) {
	o.System = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *Dosmode) GetArchive() bool {
	if o == nil || IsNil(o.Archive) {
		var ret bool
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dosmode) GetArchiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *Dosmode) HasArchive() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given bool and assigns it to the Archive field.
func (o *Dosmode) SetArchive(v bool) {
	o.Archive = &v
}

// GetReparse returns the Reparse field value if set, zero value otherwise.
func (o *Dosmode) GetReparse() bool {
	if o == nil || IsNil(o.Reparse) {
		var ret bool
		return ret
	}
	return *o.Reparse
}

// GetReparseOk returns a tuple with the Reparse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dosmode) GetReparseOk() (*bool, bool) {
	if o == nil || IsNil(o.Reparse) {
		return nil, false
	}
	return o.Reparse, true
}

// HasReparse returns a boolean if a field has been set.
func (o *Dosmode) HasReparse() bool {
	if o != nil && !IsNil(o.Reparse) {
		return true
	}

	return false
}

// SetReparse gets a reference to the given bool and assigns it to the Reparse field.
func (o *Dosmode) SetReparse(v bool) {
	o.Reparse = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *Dosmode) GetOffline() bool {
	if o == nil || IsNil(o.Offline) {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dosmode) GetOfflineOk() (*bool, bool) {
	if o == nil || IsNil(o.Offline) {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *Dosmode) HasOffline() bool {
	if o != nil && !IsNil(o.Offline) {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *Dosmode) SetOffline(v bool) {
	o.Offline = &v
}

// GetSparse returns the Sparse field value if set, zero value otherwise.
func (o *Dosmode) GetSparse() bool {
	if o == nil || IsNil(o.Sparse) {
		var ret bool
		return ret
	}
	return *o.Sparse
}

// GetSparseOk returns a tuple with the Sparse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dosmode) GetSparseOk() (*bool, bool) {
	if o == nil || IsNil(o.Sparse) {
		return nil, false
	}
	return o.Sparse, true
}

// HasSparse returns a boolean if a field has been set.
func (o *Dosmode) HasSparse() bool {
	if o != nil && !IsNil(o.Sparse) {
		return true
	}

	return false
}

// SetSparse gets a reference to the given bool and assigns it to the Sparse field.
func (o *Dosmode) SetSparse(v bool) {
	o.Sparse = &v
}

func (o Dosmode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dosmode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	if !IsNil(o.Reparse) {
		toSerialize["reparse"] = o.Reparse
	}
	if !IsNil(o.Offline) {
		toSerialize["offline"] = o.Offline
	}
	if !IsNil(o.Sparse) {
		toSerialize["sparse"] = o.Sparse
	}
	return toSerialize, nil
}

type NullableDosmode struct {
	value *Dosmode
	isSet bool
}

func (v NullableDosmode) Get() *Dosmode {
	return v.value
}

func (v *NullableDosmode) Set(val *Dosmode) {
	v.value = val
	v.isSet = true
}

func (v NullableDosmode) IsSet() bool {
	return v.isSet
}

func (v *NullableDosmode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDosmode(val *Dosmode) *NullableDosmode {
	return &NullableDosmode{value: val, isSet: true}
}

func (v NullableDosmode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDosmode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}