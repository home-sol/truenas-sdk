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

// checks if the ZfsSnapshotCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZfsSnapshotCreate0{}

// ZfsSnapshotCreate0 struct for ZfsSnapshotCreate0
type ZfsSnapshotCreate0 struct {
	Dataset      *string                `json:"dataset,omitempty"`
	Name         *string                `json:"name,omitempty"`
	NamingSchema *string                `json:"naming_schema,omitempty"`
	Recursive    *bool                  `json:"recursive,omitempty"`
	Exclude      []string               `json:"exclude,omitempty"`
	SuspendVms   *bool                  `json:"suspend_vms,omitempty"`
	VmwareSync   *bool                  `json:"vmware_sync,omitempty"`
	Properties   map[string]interface{} `json:"properties,omitempty"`
}

// NewZfsSnapshotCreate0 instantiates a new ZfsSnapshotCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZfsSnapshotCreate0() *ZfsSnapshotCreate0 {
	this := ZfsSnapshotCreate0{}
	var recursive bool
	this.Recursive = &recursive
	var suspendVms bool
	this.SuspendVms = &suspendVms
	var vmwareSync bool
	this.VmwareSync = &vmwareSync
	return &this
}

// NewZfsSnapshotCreate0WithDefaults instantiates a new ZfsSnapshotCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZfsSnapshotCreate0WithDefaults() *ZfsSnapshotCreate0 {
	this := ZfsSnapshotCreate0{}
	var recursive bool
	this.Recursive = &recursive
	var suspendVms bool
	this.SuspendVms = &suspendVms
	var vmwareSync bool
	this.VmwareSync = &vmwareSync
	return &this
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetDataset() string {
	if o == nil || IsNil(o.Dataset) {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetDatasetOk() (*string, bool) {
	if o == nil || IsNil(o.Dataset) {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasDataset() bool {
	if o != nil && !IsNil(o.Dataset) {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *ZfsSnapshotCreate0) SetDataset(v string) {
	o.Dataset = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZfsSnapshotCreate0) SetName(v string) {
	o.Name = &v
}

// GetNamingSchema returns the NamingSchema field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetNamingSchema() string {
	if o == nil || IsNil(o.NamingSchema) {
		var ret string
		return ret
	}
	return *o.NamingSchema
}

// GetNamingSchemaOk returns a tuple with the NamingSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetNamingSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.NamingSchema) {
		return nil, false
	}
	return o.NamingSchema, true
}

// HasNamingSchema returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasNamingSchema() bool {
	if o != nil && !IsNil(o.NamingSchema) {
		return true
	}

	return false
}

// SetNamingSchema gets a reference to the given string and assigns it to the NamingSchema field.
func (o *ZfsSnapshotCreate0) SetNamingSchema(v string) {
	o.NamingSchema = &v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetRecursive() bool {
	if o == nil || IsNil(o.Recursive) {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetRecursiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Recursive) {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasRecursive() bool {
	if o != nil && !IsNil(o.Recursive) {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *ZfsSnapshotCreate0) SetRecursive(v bool) {
	o.Recursive = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *ZfsSnapshotCreate0) SetExclude(v []string) {
	o.Exclude = v
}

// GetSuspendVms returns the SuspendVms field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetSuspendVms() bool {
	if o == nil || IsNil(o.SuspendVms) {
		var ret bool
		return ret
	}
	return *o.SuspendVms
}

// GetSuspendVmsOk returns a tuple with the SuspendVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetSuspendVmsOk() (*bool, bool) {
	if o == nil || IsNil(o.SuspendVms) {
		return nil, false
	}
	return o.SuspendVms, true
}

// HasSuspendVms returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasSuspendVms() bool {
	if o != nil && !IsNil(o.SuspendVms) {
		return true
	}

	return false
}

// SetSuspendVms gets a reference to the given bool and assigns it to the SuspendVms field.
func (o *ZfsSnapshotCreate0) SetSuspendVms(v bool) {
	o.SuspendVms = &v
}

// GetVmwareSync returns the VmwareSync field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetVmwareSync() bool {
	if o == nil || IsNil(o.VmwareSync) {
		var ret bool
		return ret
	}
	return *o.VmwareSync
}

// GetVmwareSyncOk returns a tuple with the VmwareSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetVmwareSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.VmwareSync) {
		return nil, false
	}
	return o.VmwareSync, true
}

// HasVmwareSync returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasVmwareSync() bool {
	if o != nil && !IsNil(o.VmwareSync) {
		return true
	}

	return false
}

// SetVmwareSync gets a reference to the given bool and assigns it to the VmwareSync field.
func (o *ZfsSnapshotCreate0) SetVmwareSync(v bool) {
	o.VmwareSync = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ZfsSnapshotCreate0) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotCreate0) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ZfsSnapshotCreate0) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *ZfsSnapshotCreate0) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o ZfsSnapshotCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZfsSnapshotCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dataset) {
		toSerialize["dataset"] = o.Dataset
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NamingSchema) {
		toSerialize["naming_schema"] = o.NamingSchema
	}
	if !IsNil(o.Recursive) {
		toSerialize["recursive"] = o.Recursive
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.SuspendVms) {
		toSerialize["suspend_vms"] = o.SuspendVms
	}
	if !IsNil(o.VmwareSync) {
		toSerialize["vmware_sync"] = o.VmwareSync
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableZfsSnapshotCreate0 struct {
	value *ZfsSnapshotCreate0
	isSet bool
}

func (v NullableZfsSnapshotCreate0) Get() *ZfsSnapshotCreate0 {
	return v.value
}

func (v *NullableZfsSnapshotCreate0) Set(val *ZfsSnapshotCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableZfsSnapshotCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableZfsSnapshotCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZfsSnapshotCreate0(val *ZfsSnapshotCreate0) *NullableZfsSnapshotCreate0 {
	return &NullableZfsSnapshotCreate0{value: val, isSet: true}
}

func (v NullableZfsSnapshotCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZfsSnapshotCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
