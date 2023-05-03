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

// checks if the CloudsyncListDirectory0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudsyncListDirectory0{}

// CloudsyncListDirectory0 struct for CloudsyncListDirectory0
type CloudsyncListDirectory0 struct {
	Credentials        *int32  `json:"credentials,omitempty"`
	Encryption         *bool   `json:"encryption,omitempty"`
	FilenameEncryption *bool   `json:"filename_encryption,omitempty"`
	EncryptionPassword *string `json:"encryption_password,omitempty"`
	EncryptionSalt     *string `json:"encryption_salt,omitempty"`
	// If remote supports buckets, path is constructed by two keys \"bucket\"/\"folder\" in `attributes`. If remote does not support buckets, path is constructed using \"folder\" key only in `attributes`. \"folder\" is directory name and \"bucket\" is bucket name for remote.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Args       *string                `json:"args,omitempty"`
}

// NewCloudsyncListDirectory0 instantiates a new CloudsyncListDirectory0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudsyncListDirectory0() *CloudsyncListDirectory0 {
	this := CloudsyncListDirectory0{}
	var encryption bool
	this.Encryption = &encryption
	var filenameEncryption bool
	this.FilenameEncryption = &filenameEncryption
	var encryptionPassword string
	this.EncryptionPassword = &encryptionPassword
	var encryptionSalt string
	this.EncryptionSalt = &encryptionSalt
	var args string
	this.Args = &args
	return &this
}

// NewCloudsyncListDirectory0WithDefaults instantiates a new CloudsyncListDirectory0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudsyncListDirectory0WithDefaults() *CloudsyncListDirectory0 {
	this := CloudsyncListDirectory0{}
	var encryption bool
	this.Encryption = &encryption
	var filenameEncryption bool
	this.FilenameEncryption = &filenameEncryption
	var encryptionPassword string
	this.EncryptionPassword = &encryptionPassword
	var encryptionSalt string
	this.EncryptionSalt = &encryptionSalt
	var args string
	this.Args = &args
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CloudsyncListDirectory0) GetCredentials() int32 {
	if o == nil || IsNil(o.Credentials) {
		var ret int32
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncListDirectory0) GetCredentialsOk() (*int32, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CloudsyncListDirectory0) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given int32 and assigns it to the Credentials field.
func (o *CloudsyncListDirectory0) SetCredentials(v int32) {
	o.Credentials = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *CloudsyncListDirectory0) GetEncryption() bool {
	if o == nil || IsNil(o.Encryption) {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncListDirectory0) GetEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *CloudsyncListDirectory0) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *CloudsyncListDirectory0) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetFilenameEncryption returns the FilenameEncryption field value if set, zero value otherwise.
func (o *CloudsyncListDirectory0) GetFilenameEncryption() bool {
	if o == nil || IsNil(o.FilenameEncryption) {
		var ret bool
		return ret
	}
	return *o.FilenameEncryption
}

// GetFilenameEncryptionOk returns a tuple with the FilenameEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncListDirectory0) GetFilenameEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.FilenameEncryption) {
		return nil, false
	}
	return o.FilenameEncryption, true
}

// HasFilenameEncryption returns a boolean if a field has been set.
func (o *CloudsyncListDirectory0) HasFilenameEncryption() bool {
	if o != nil && !IsNil(o.FilenameEncryption) {
		return true
	}

	return false
}

// SetFilenameEncryption gets a reference to the given bool and assigns it to the FilenameEncryption field.
func (o *CloudsyncListDirectory0) SetFilenameEncryption(v bool) {
	o.FilenameEncryption = &v
}

// GetEncryptionPassword returns the EncryptionPassword field value if set, zero value otherwise.
func (o *CloudsyncListDirectory0) GetEncryptionPassword() string {
	if o == nil || IsNil(o.EncryptionPassword) {
		var ret string
		return ret
	}
	return *o.EncryptionPassword
}

// GetEncryptionPasswordOk returns a tuple with the EncryptionPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncListDirectory0) GetEncryptionPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionPassword) {
		return nil, false
	}
	return o.EncryptionPassword, true
}

// HasEncryptionPassword returns a boolean if a field has been set.
func (o *CloudsyncListDirectory0) HasEncryptionPassword() bool {
	if o != nil && !IsNil(o.EncryptionPassword) {
		return true
	}

	return false
}

// SetEncryptionPassword gets a reference to the given string and assigns it to the EncryptionPassword field.
func (o *CloudsyncListDirectory0) SetEncryptionPassword(v string) {
	o.EncryptionPassword = &v
}

// GetEncryptionSalt returns the EncryptionSalt field value if set, zero value otherwise.
func (o *CloudsyncListDirectory0) GetEncryptionSalt() string {
	if o == nil || IsNil(o.EncryptionSalt) {
		var ret string
		return ret
	}
	return *o.EncryptionSalt
}

// GetEncryptionSaltOk returns a tuple with the EncryptionSalt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncListDirectory0) GetEncryptionSaltOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSalt) {
		return nil, false
	}
	return o.EncryptionSalt, true
}

// HasEncryptionSalt returns a boolean if a field has been set.
func (o *CloudsyncListDirectory0) HasEncryptionSalt() bool {
	if o != nil && !IsNil(o.EncryptionSalt) {
		return true
	}

	return false
}

// SetEncryptionSalt gets a reference to the given string and assigns it to the EncryptionSalt field.
func (o *CloudsyncListDirectory0) SetEncryptionSalt(v string) {
	o.EncryptionSalt = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CloudsyncListDirectory0) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncListDirectory0) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CloudsyncListDirectory0) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *CloudsyncListDirectory0) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *CloudsyncListDirectory0) GetArgs() string {
	if o == nil || IsNil(o.Args) {
		var ret string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncListDirectory0) GetArgsOk() (*string, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *CloudsyncListDirectory0) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given string and assigns it to the Args field.
func (o *CloudsyncListDirectory0) SetArgs(v string) {
	o.Args = &v
}

func (o CloudsyncListDirectory0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudsyncListDirectory0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if !IsNil(o.FilenameEncryption) {
		toSerialize["filename_encryption"] = o.FilenameEncryption
	}
	if !IsNil(o.EncryptionPassword) {
		toSerialize["encryption_password"] = o.EncryptionPassword
	}
	if !IsNil(o.EncryptionSalt) {
		toSerialize["encryption_salt"] = o.EncryptionSalt
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	return toSerialize, nil
}

type NullableCloudsyncListDirectory0 struct {
	value *CloudsyncListDirectory0
	isSet bool
}

func (v NullableCloudsyncListDirectory0) Get() *CloudsyncListDirectory0 {
	return v.value
}

func (v *NullableCloudsyncListDirectory0) Set(val *CloudsyncListDirectory0) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudsyncListDirectory0) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudsyncListDirectory0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudsyncListDirectory0(val *CloudsyncListDirectory0) *NullableCloudsyncListDirectory0 {
	return &NullableCloudsyncListDirectory0{value: val, isSet: true}
}

func (v NullableCloudsyncListDirectory0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudsyncListDirectory0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}