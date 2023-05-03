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

// checks if the PoolCreate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolCreate0{}

// PoolCreate0 struct for PoolCreate0
type PoolCreate0 struct {
	// `encryption` when enabled will create an ZFS encrypted root dataset for `name` pool. `encryption_options` specifies configuration for encryption of root dataset for `name` pool. `encryption_options.passphrase` must be specified if encryption for root dataset is desired with a passphrase as a key. Otherwise a hex encoded key can be specified by providing `encryption_options.key`. `encryption_options.generate_key` when enabled automatically generates the key to be used for dataset encryption.
	Name *string `json:"name,omitempty"`
	// `encryption` when enabled will create an ZFS encrypted root dataset for `name` pool.
	Encryption *bool `json:"encryption,omitempty"`
	// `deduplication` when set to ON or VERIFY makes sure that no block of data is duplicated in the pool. When VERIFY is specified, if two blocks have similar signatures, byte to byte comparison is performed to ensure that the blocks are identical. This should be used in special circumstances as it carries a significant overhead.
	Deduplication         NullableString     `json:"deduplication,omitempty"`
	Checksum              NullableString     `json:"checksum,omitempty"`
	EncryptionOptions     *EncryptionOptions `json:"encryption_options,omitempty"`
	Topology              *Topology          `json:"topology,omitempty"`
	AllowDuplicateSerials *bool              `json:"allow_duplicate_serials,omitempty"`
}

// NewPoolCreate0 instantiates a new PoolCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolCreate0() *PoolCreate0 {
	this := PoolCreate0{}
	var encryption bool
	this.Encryption = &encryption
	var encryptionOptions EncryptionOptions
	this.EncryptionOptions = &encryptionOptions
	var topology Topology
	this.Topology = &topology
	var allowDuplicateSerials bool
	this.AllowDuplicateSerials = &allowDuplicateSerials
	return &this
}

// NewPoolCreate0WithDefaults instantiates a new PoolCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolCreate0WithDefaults() *PoolCreate0 {
	this := PoolCreate0{}
	var encryption bool
	this.Encryption = &encryption
	var encryptionOptions EncryptionOptions
	this.EncryptionOptions = &encryptionOptions
	var topology Topology
	this.Topology = &topology
	var allowDuplicateSerials bool
	this.AllowDuplicateSerials = &allowDuplicateSerials
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolCreate0) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolCreate0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolCreate0) SetName(v string) {
	o.Name = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *PoolCreate0) GetEncryption() bool {
	if o == nil || IsNil(o.Encryption) {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0) GetEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *PoolCreate0) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *PoolCreate0) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetDeduplication returns the Deduplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolCreate0) GetDeduplication() string {
	if o == nil || IsNil(o.Deduplication.Get()) {
		var ret string
		return ret
	}
	return *o.Deduplication.Get()
}

// GetDeduplicationOk returns a tuple with the Deduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolCreate0) GetDeduplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deduplication.Get(), o.Deduplication.IsSet()
}

// HasDeduplication returns a boolean if a field has been set.
func (o *PoolCreate0) HasDeduplication() bool {
	if o != nil && o.Deduplication.IsSet() {
		return true
	}

	return false
}

// SetDeduplication gets a reference to the given NullableString and assigns it to the Deduplication field.
func (o *PoolCreate0) SetDeduplication(v string) {
	o.Deduplication.Set(&v)
}

// SetDeduplicationNil sets the value for Deduplication to be an explicit nil
func (o *PoolCreate0) SetDeduplicationNil() {
	o.Deduplication.Set(nil)
}

// UnsetDeduplication ensures that no value is present for Deduplication, not even an explicit nil
func (o *PoolCreate0) UnsetDeduplication() {
	o.Deduplication.Unset()
}

// GetChecksum returns the Checksum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolCreate0) GetChecksum() string {
	if o == nil || IsNil(o.Checksum.Get()) {
		var ret string
		return ret
	}
	return *o.Checksum.Get()
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolCreate0) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Checksum.Get(), o.Checksum.IsSet()
}

// HasChecksum returns a boolean if a field has been set.
func (o *PoolCreate0) HasChecksum() bool {
	if o != nil && o.Checksum.IsSet() {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given NullableString and assigns it to the Checksum field.
func (o *PoolCreate0) SetChecksum(v string) {
	o.Checksum.Set(&v)
}

// SetChecksumNil sets the value for Checksum to be an explicit nil
func (o *PoolCreate0) SetChecksumNil() {
	o.Checksum.Set(nil)
}

// UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
func (o *PoolCreate0) UnsetChecksum() {
	o.Checksum.Unset()
}

// GetEncryptionOptions returns the EncryptionOptions field value if set, zero value otherwise.
func (o *PoolCreate0) GetEncryptionOptions() EncryptionOptions {
	if o == nil || IsNil(o.EncryptionOptions) {
		var ret EncryptionOptions
		return ret
	}
	return *o.EncryptionOptions
}

// GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0) GetEncryptionOptionsOk() (*EncryptionOptions, bool) {
	if o == nil || IsNil(o.EncryptionOptions) {
		return nil, false
	}
	return o.EncryptionOptions, true
}

// HasEncryptionOptions returns a boolean if a field has been set.
func (o *PoolCreate0) HasEncryptionOptions() bool {
	if o != nil && !IsNil(o.EncryptionOptions) {
		return true
	}

	return false
}

// SetEncryptionOptions gets a reference to the given EncryptionOptions and assigns it to the EncryptionOptions field.
func (o *PoolCreate0) SetEncryptionOptions(v EncryptionOptions) {
	o.EncryptionOptions = &v
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *PoolCreate0) GetTopology() Topology {
	if o == nil || IsNil(o.Topology) {
		var ret Topology
		return ret
	}
	return *o.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0) GetTopologyOk() (*Topology, bool) {
	if o == nil || IsNil(o.Topology) {
		return nil, false
	}
	return o.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *PoolCreate0) HasTopology() bool {
	if o != nil && !IsNil(o.Topology) {
		return true
	}

	return false
}

// SetTopology gets a reference to the given Topology and assigns it to the Topology field.
func (o *PoolCreate0) SetTopology(v Topology) {
	o.Topology = &v
}

// GetAllowDuplicateSerials returns the AllowDuplicateSerials field value if set, zero value otherwise.
func (o *PoolCreate0) GetAllowDuplicateSerials() bool {
	if o == nil || IsNil(o.AllowDuplicateSerials) {
		var ret bool
		return ret
	}
	return *o.AllowDuplicateSerials
}

// GetAllowDuplicateSerialsOk returns a tuple with the AllowDuplicateSerials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0) GetAllowDuplicateSerialsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowDuplicateSerials) {
		return nil, false
	}
	return o.AllowDuplicateSerials, true
}

// HasAllowDuplicateSerials returns a boolean if a field has been set.
func (o *PoolCreate0) HasAllowDuplicateSerials() bool {
	if o != nil && !IsNil(o.AllowDuplicateSerials) {
		return true
	}

	return false
}

// SetAllowDuplicateSerials gets a reference to the given bool and assigns it to the AllowDuplicateSerials field.
func (o *PoolCreate0) SetAllowDuplicateSerials(v bool) {
	o.AllowDuplicateSerials = &v
}

func (o PoolCreate0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolCreate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if o.Deduplication.IsSet() {
		toSerialize["deduplication"] = o.Deduplication.Get()
	}
	if o.Checksum.IsSet() {
		toSerialize["checksum"] = o.Checksum.Get()
	}
	if !IsNil(o.EncryptionOptions) {
		toSerialize["encryption_options"] = o.EncryptionOptions
	}
	if !IsNil(o.Topology) {
		toSerialize["topology"] = o.Topology
	}
	if !IsNil(o.AllowDuplicateSerials) {
		toSerialize["allow_duplicate_serials"] = o.AllowDuplicateSerials
	}
	return toSerialize, nil
}

type NullablePoolCreate0 struct {
	value *PoolCreate0
	isSet bool
}

func (v NullablePoolCreate0) Get() *PoolCreate0 {
	return v.value
}

func (v *NullablePoolCreate0) Set(val *PoolCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolCreate0(val *PoolCreate0) *NullablePoolCreate0 {
	return &NullablePoolCreate0{value: val, isSet: true}
}

func (v NullablePoolCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}