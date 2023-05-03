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

// checks if the IscsiExtentCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IscsiExtentCreate0{}

// IscsiExtentCreate0 struct for IscsiExtentCreate0
type IscsiExtentCreate0 struct {
	Name *string `json:"name,omitempty"`
	// When `type` is set to FILE, attribute `filesize` is used and it represents number of bytes. `filesize` if not zero should be a multiple of `blocksize`. `path` is a required attribute with `type` set as FILE. With `type` being set to DISK, a valid ZFS volume is required.
	Type   *string        `json:"type,omitempty"`
	Disk   NullableString `json:"disk,omitempty"`
	Serial NullableString `json:"serial,omitempty"`
	// When `type` is set to FILE, attribute `filesize` is used and it represents number of bytes. `filesize` if not zero should be a multiple of `blocksize`. `path` is a required attribute with `type` set as FILE.
	Path NullableString `json:"path,omitempty"`
	// When `type` is set to FILE, attribute `filesize` is used and it represents number of bytes. `filesize` if not zero should be a multiple of `blocksize`. `path` is a required attribute with `type` set as FILE.
	Filesize *int32 `json:"filesize,omitempty"`
	// When `type` is set to FILE, attribute `filesize` is used and it represents number of bytes. `filesize` if not zero should be a multiple of `blocksize`. `path` is a required attribute with `type` set as FILE.
	Blocksize      *int32        `json:"blocksize,omitempty"`
	Pblocksize     *bool         `json:"pblocksize,omitempty"`
	AvailThreshold NullableInt32 `json:"avail_threshold,omitempty"`
	Comment        *string       `json:"comment,omitempty"`
	// `insecure_tpc` when enabled allows an initiator to bypass normal access control and access any scannable target. This allows xcopy operations otherwise blocked by access control.
	InsecureTpc *bool `json:"insecure_tpc,omitempty"`
	// `xen` is a boolean value which is set to true if Xen is being used as the iSCSI initiator.
	Xen     *bool   `json:"xen,omitempty"`
	Rpm     *string `json:"rpm,omitempty"`
	Ro      *bool   `json:"ro,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

// NewIscsiExtentCreate0 instantiates a new IscsiExtentCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiExtentCreate0() *IscsiExtentCreate0 {
	this := IscsiExtentCreate0{}
	var type_ string
	this.Type = &type_
	var filesize int32
	this.Filesize = &filesize
	var blocksize int32
	this.Blocksize = &blocksize
	var insecureTpc bool
	this.InsecureTpc = &insecureTpc
	var rpm string
	this.Rpm = &rpm
	var ro bool
	this.Ro = &ro
	var enabled bool
	this.Enabled = &enabled
	return &this
}

// NewIscsiExtentCreate0WithDefaults instantiates a new IscsiExtentCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiExtentCreate0WithDefaults() *IscsiExtentCreate0 {
	this := IscsiExtentCreate0{}
	var type_ string
	this.Type = &type_
	var filesize int32
	this.Filesize = &filesize
	var blocksize int32
	this.Blocksize = &blocksize
	var insecureTpc bool
	this.InsecureTpc = &insecureTpc
	var rpm string
	this.Rpm = &rpm
	var ro bool
	this.Ro = &ro
	var enabled bool
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IscsiExtentCreate0) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IscsiExtentCreate0) SetType(v string) {
	o.Type = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiExtentCreate0) GetDisk() string {
	if o == nil || IsNil(o.Disk.Get()) {
		var ret string
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiExtentCreate0) GetDiskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableString and assigns it to the Disk field.
func (o *IscsiExtentCreate0) SetDisk(v string) {
	o.Disk.Set(&v)
}

// SetDiskNil sets the value for Disk to be an explicit nil
func (o *IscsiExtentCreate0) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *IscsiExtentCreate0) UnsetDisk() {
	o.Disk.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiExtentCreate0) GetSerial() string {
	if o == nil || IsNil(o.Serial.Get()) {
		var ret string
		return ret
	}
	return *o.Serial.Get()
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiExtentCreate0) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serial.Get(), o.Serial.IsSet()
}

// HasSerial returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasSerial() bool {
	if o != nil && o.Serial.IsSet() {
		return true
	}

	return false
}

// SetSerial gets a reference to the given NullableString and assigns it to the Serial field.
func (o *IscsiExtentCreate0) SetSerial(v string) {
	o.Serial.Set(&v)
}

// SetSerialNil sets the value for Serial to be an explicit nil
func (o *IscsiExtentCreate0) SetSerialNil() {
	o.Serial.Set(nil)
}

// UnsetSerial ensures that no value is present for Serial, not even an explicit nil
func (o *IscsiExtentCreate0) UnsetSerial() {
	o.Serial.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiExtentCreate0) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiExtentCreate0) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *IscsiExtentCreate0) SetPath(v string) {
	o.Path.Set(&v)
}

// SetPathNil sets the value for Path to be an explicit nil
func (o *IscsiExtentCreate0) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *IscsiExtentCreate0) UnsetPath() {
	o.Path.Unset()
}

// GetFilesize returns the Filesize field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetFilesize() int32 {
	if o == nil || IsNil(o.Filesize) {
		var ret int32
		return ret
	}
	return *o.Filesize
}

// GetFilesizeOk returns a tuple with the Filesize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetFilesizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Filesize) {
		return nil, false
	}
	return o.Filesize, true
}

// HasFilesize returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasFilesize() bool {
	if o != nil && !IsNil(o.Filesize) {
		return true
	}

	return false
}

// SetFilesize gets a reference to the given int32 and assigns it to the Filesize field.
func (o *IscsiExtentCreate0) SetFilesize(v int32) {
	o.Filesize = &v
}

// GetBlocksize returns the Blocksize field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetBlocksize() int32 {
	if o == nil || IsNil(o.Blocksize) {
		var ret int32
		return ret
	}
	return *o.Blocksize
}

// GetBlocksizeOk returns a tuple with the Blocksize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetBlocksizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Blocksize) {
		return nil, false
	}
	return o.Blocksize, true
}

// HasBlocksize returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasBlocksize() bool {
	if o != nil && !IsNil(o.Blocksize) {
		return true
	}

	return false
}

// SetBlocksize gets a reference to the given int32 and assigns it to the Blocksize field.
func (o *IscsiExtentCreate0) SetBlocksize(v int32) {
	o.Blocksize = &v
}

// GetPblocksize returns the Pblocksize field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetPblocksize() bool {
	if o == nil || IsNil(o.Pblocksize) {
		var ret bool
		return ret
	}
	return *o.Pblocksize
}

// GetPblocksizeOk returns a tuple with the Pblocksize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetPblocksizeOk() (*bool, bool) {
	if o == nil || IsNil(o.Pblocksize) {
		return nil, false
	}
	return o.Pblocksize, true
}

// HasPblocksize returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasPblocksize() bool {
	if o != nil && !IsNil(o.Pblocksize) {
		return true
	}

	return false
}

// SetPblocksize gets a reference to the given bool and assigns it to the Pblocksize field.
func (o *IscsiExtentCreate0) SetPblocksize(v bool) {
	o.Pblocksize = &v
}

// GetAvailThreshold returns the AvailThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiExtentCreate0) GetAvailThreshold() int32 {
	if o == nil || IsNil(o.AvailThreshold.Get()) {
		var ret int32
		return ret
	}
	return *o.AvailThreshold.Get()
}

// GetAvailThresholdOk returns a tuple with the AvailThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiExtentCreate0) GetAvailThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailThreshold.Get(), o.AvailThreshold.IsSet()
}

// HasAvailThreshold returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasAvailThreshold() bool {
	if o != nil && o.AvailThreshold.IsSet() {
		return true
	}

	return false
}

// SetAvailThreshold gets a reference to the given NullableInt32 and assigns it to the AvailThreshold field.
func (o *IscsiExtentCreate0) SetAvailThreshold(v int32) {
	o.AvailThreshold.Set(&v)
}

// SetAvailThresholdNil sets the value for AvailThreshold to be an explicit nil
func (o *IscsiExtentCreate0) SetAvailThresholdNil() {
	o.AvailThreshold.Set(nil)
}

// UnsetAvailThreshold ensures that no value is present for AvailThreshold, not even an explicit nil
func (o *IscsiExtentCreate0) UnsetAvailThreshold() {
	o.AvailThreshold.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IscsiExtentCreate0) SetComment(v string) {
	o.Comment = &v
}

// GetInsecureTpc returns the InsecureTpc field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetInsecureTpc() bool {
	if o == nil || IsNil(o.InsecureTpc) {
		var ret bool
		return ret
	}
	return *o.InsecureTpc
}

// GetInsecureTpcOk returns a tuple with the InsecureTpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetInsecureTpcOk() (*bool, bool) {
	if o == nil || IsNil(o.InsecureTpc) {
		return nil, false
	}
	return o.InsecureTpc, true
}

// HasInsecureTpc returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasInsecureTpc() bool {
	if o != nil && !IsNil(o.InsecureTpc) {
		return true
	}

	return false
}

// SetInsecureTpc gets a reference to the given bool and assigns it to the InsecureTpc field.
func (o *IscsiExtentCreate0) SetInsecureTpc(v bool) {
	o.InsecureTpc = &v
}

// GetXen returns the Xen field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetXen() bool {
	if o == nil || IsNil(o.Xen) {
		var ret bool
		return ret
	}
	return *o.Xen
}

// GetXenOk returns a tuple with the Xen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetXenOk() (*bool, bool) {
	if o == nil || IsNil(o.Xen) {
		return nil, false
	}
	return o.Xen, true
}

// HasXen returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasXen() bool {
	if o != nil && !IsNil(o.Xen) {
		return true
	}

	return false
}

// SetXen gets a reference to the given bool and assigns it to the Xen field.
func (o *IscsiExtentCreate0) SetXen(v bool) {
	o.Xen = &v
}

// GetRpm returns the Rpm field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetRpm() string {
	if o == nil || IsNil(o.Rpm) {
		var ret string
		return ret
	}
	return *o.Rpm
}

// GetRpmOk returns a tuple with the Rpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetRpmOk() (*string, bool) {
	if o == nil || IsNil(o.Rpm) {
		return nil, false
	}
	return o.Rpm, true
}

// HasRpm returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasRpm() bool {
	if o != nil && !IsNil(o.Rpm) {
		return true
	}

	return false
}

// SetRpm gets a reference to the given string and assigns it to the Rpm field.
func (o *IscsiExtentCreate0) SetRpm(v string) {
	o.Rpm = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetRo() bool {
	if o == nil || IsNil(o.Ro) {
		var ret bool
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetRoOk() (*bool, bool) {
	if o == nil || IsNil(o.Ro) {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasRo() bool {
	if o != nil && !IsNil(o.Ro) {
		return true
	}

	return false
}

// SetRo gets a reference to the given bool and assigns it to the Ro field.
func (o *IscsiExtentCreate0) SetRo(v bool) {
	o.Ro = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IscsiExtentCreate0) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiExtentCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IscsiExtentCreate0) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IscsiExtentCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o IscsiExtentCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IscsiExtentCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
	}
	if o.Serial.IsSet() {
		toSerialize["serial"] = o.Serial.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !IsNil(o.Filesize) {
		toSerialize["filesize"] = o.Filesize
	}
	if !IsNil(o.Blocksize) {
		toSerialize["blocksize"] = o.Blocksize
	}
	if !IsNil(o.Pblocksize) {
		toSerialize["pblocksize"] = o.Pblocksize
	}
	if o.AvailThreshold.IsSet() {
		toSerialize["avail_threshold"] = o.AvailThreshold.Get()
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.InsecureTpc) {
		toSerialize["insecure_tpc"] = o.InsecureTpc
	}
	if !IsNil(o.Xen) {
		toSerialize["xen"] = o.Xen
	}
	if !IsNil(o.Rpm) {
		toSerialize["rpm"] = o.Rpm
	}
	if !IsNil(o.Ro) {
		toSerialize["ro"] = o.Ro
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableIscsiExtentCreate0 struct {
	value *IscsiExtentCreate0
	isSet bool
}

func (v NullableIscsiExtentCreate0) Get() *IscsiExtentCreate0 {
	return v.value
}

func (v *NullableIscsiExtentCreate0) Set(val *IscsiExtentCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiExtentCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiExtentCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiExtentCreate0(val *IscsiExtentCreate0) *NullableIscsiExtentCreate0 {
	return &NullableIscsiExtentCreate0{value: val, isSet: true}
}

func (v NullableIscsiExtentCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiExtentCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
