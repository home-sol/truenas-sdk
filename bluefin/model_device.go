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

// checks if the Device type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Device{}

// Device struct for Device
type Device struct {
	Name                 *string                  `json:"name,omitempty"`
	Path                 *string                  `json:"path,omitempty"`
	Guid                 *string                  `json:"guid,omitempty"`
	Status               *string                  `json:"status,omitempty"`
	Stats                map[string]interface{}   `json:"stats,omitempty"`
	Device               *string                  `json:"device,omitempty"`
	Disk                 *string                  `json:"disk,omitempty"`
	Children             []map[string]interface{} `json:"children,omitempty"`
	UnavailDisk          map[string]interface{}   `json:"unavail_disk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Device Device

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice() *Device {
	this := Device{}
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Device) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Device) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Device) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Device) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Device) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Device) SetPath(v string) {
	o.Path = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *Device) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *Device) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *Device) SetGuid(v string) {
	o.Guid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Device) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Device) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Device) SetStatus(v string) {
	o.Status = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *Device) GetStats() map[string]interface{} {
	if o == nil || IsNil(o.Stats) {
		var ret map[string]interface{}
		return ret
	}
	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetStatsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Stats) {
		return map[string]interface{}{}, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *Device) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given map[string]interface{} and assigns it to the Stats field.
func (o *Device) SetStats(v map[string]interface{}) {
	o.Stats = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *Device) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *Device) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *Device) SetDevice(v string) {
	o.Device = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *Device) GetDisk() string {
	if o == nil || IsNil(o.Disk) {
		var ret string
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDiskOk() (*string, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *Device) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given string and assigns it to the Disk field.
func (o *Device) SetDisk(v string) {
	o.Disk = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Device) GetChildren() []map[string]interface{} {
	if o == nil || IsNil(o.Children) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetChildrenOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Device) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []map[string]interface{} and assigns it to the Children field.
func (o *Device) SetChildren(v []map[string]interface{}) {
	o.Children = v
}

// GetUnavailDisk returns the UnavailDisk field value if set, zero value otherwise.
func (o *Device) GetUnavailDisk() map[string]interface{} {
	if o == nil || IsNil(o.UnavailDisk) {
		var ret map[string]interface{}
		return ret
	}
	return o.UnavailDisk
}

// GetUnavailDiskOk returns a tuple with the UnavailDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetUnavailDiskOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UnavailDisk) {
		return map[string]interface{}{}, false
	}
	return o.UnavailDisk, true
}

// HasUnavailDisk returns a boolean if a field has been set.
func (o *Device) HasUnavailDisk() bool {
	if o != nil && !IsNil(o.UnavailDisk) {
		return true
	}

	return false
}

// SetUnavailDisk gets a reference to the given map[string]interface{} and assigns it to the UnavailDisk field.
func (o *Device) SetUnavailDisk(v map[string]interface{}) {
	o.UnavailDisk = v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Device) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.UnavailDisk) {
		toSerialize["unavail_disk"] = o.UnavailDisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Device) UnmarshalJSON(bytes []byte) (err error) {
	varDevice := _Device{}

	if err = json.Unmarshal(bytes, &varDevice); err == nil {
		*o = Device(varDevice)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "path")
		delete(additionalProperties, "guid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "stats")
		delete(additionalProperties, "device")
		delete(additionalProperties, "disk")
		delete(additionalProperties, "children")
		delete(additionalProperties, "unavail_disk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}