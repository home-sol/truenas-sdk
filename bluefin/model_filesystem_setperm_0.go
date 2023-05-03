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

// checks if the FilesystemSetperm0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesystemSetperm0{}

// FilesystemSetperm0 struct for FilesystemSetperm0
type FilesystemSetperm0 struct {
	// Set unix permissions on given `path`. `stripacl` setperm will fail if an extended ACL is present on `path`, unless `stripacl` is set to True.
	Path *string `json:"path,omitempty"`
	// If `mode` is specified then the mode will be applied to the path and files and subdirectories depending on which `options` are selected. Mode should be formatted as string representation of octal permissions bits.
	Mode NullableString `json:"mode,omitempty"`
	// `uid` the desired UID of the file user. If set to None (the default), then user is not changed.
	Uid NullableInt32 `json:"uid,omitempty"`
	// `gid` the desired GID of the file group. If set to None (the default), then group is not changed.
	Gid     NullableInt32 `json:"gid,omitempty"`
	Options *Options2     `json:"options,omitempty"`
}

// NewFilesystemSetperm0 instantiates a new FilesystemSetperm0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemSetperm0() *FilesystemSetperm0 {
	this := FilesystemSetperm0{}
	var options Options2
	this.Options = &options
	return &this
}

// NewFilesystemSetperm0WithDefaults instantiates a new FilesystemSetperm0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemSetperm0WithDefaults() *FilesystemSetperm0 {
	this := FilesystemSetperm0{}
	var options Options2
	this.Options = &options
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FilesystemSetperm0) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemSetperm0) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FilesystemSetperm0) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FilesystemSetperm0) SetPath(v string) {
	o.Path = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesystemSetperm0) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesystemSetperm0) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *FilesystemSetperm0) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *FilesystemSetperm0) SetMode(v string) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil
func (o *FilesystemSetperm0) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *FilesystemSetperm0) UnsetMode() {
	o.Mode.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesystemSetperm0) GetUid() int32 {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret int32
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesystemSetperm0) GetUidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *FilesystemSetperm0) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableInt32 and assigns it to the Uid field.
func (o *FilesystemSetperm0) SetUid(v int32) {
	o.Uid.Set(&v)
}

// SetUidNil sets the value for Uid to be an explicit nil
func (o *FilesystemSetperm0) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *FilesystemSetperm0) UnsetUid() {
	o.Uid.Unset()
}

// GetGid returns the Gid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesystemSetperm0) GetGid() int32 {
	if o == nil || IsNil(o.Gid.Get()) {
		var ret int32
		return ret
	}
	return *o.Gid.Get()
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesystemSetperm0) GetGidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gid.Get(), o.Gid.IsSet()
}

// HasGid returns a boolean if a field has been set.
func (o *FilesystemSetperm0) HasGid() bool {
	if o != nil && o.Gid.IsSet() {
		return true
	}

	return false
}

// SetGid gets a reference to the given NullableInt32 and assigns it to the Gid field.
func (o *FilesystemSetperm0) SetGid(v int32) {
	o.Gid.Set(&v)
}

// SetGidNil sets the value for Gid to be an explicit nil
func (o *FilesystemSetperm0) SetGidNil() {
	o.Gid.Set(nil)
}

// UnsetGid ensures that no value is present for Gid, not even an explicit nil
func (o *FilesystemSetperm0) UnsetGid() {
	o.Gid.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FilesystemSetperm0) GetOptions() Options2 {
	if o == nil || IsNil(o.Options) {
		var ret Options2
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemSetperm0) GetOptionsOk() (*Options2, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FilesystemSetperm0) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Options2 and assigns it to the Options field.
func (o *FilesystemSetperm0) SetOptions(v Options2) {
	o.Options = &v
}

func (o FilesystemSetperm0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesystemSetperm0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["uid"] = o.Uid.Get()
	}
	if o.Gid.IsSet() {
		toSerialize["gid"] = o.Gid.Get()
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableFilesystemSetperm0 struct {
	value *FilesystemSetperm0
	isSet bool
}

func (v NullableFilesystemSetperm0) Get() *FilesystemSetperm0 {
	return v.value
}

func (v *NullableFilesystemSetperm0) Set(val *FilesystemSetperm0) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemSetperm0) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemSetperm0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemSetperm0(val *FilesystemSetperm0) *NullableFilesystemSetperm0 {
	return &NullableFilesystemSetperm0{value: val, isSet: true}
}

func (v NullableFilesystemSetperm0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemSetperm0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}