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

// checks if the SmbGetRemoteAcl0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmbGetRemoteAcl0{}

// SmbGetRemoteAcl0 struct for SmbGetRemoteAcl0
type SmbGetRemoteAcl0 struct {
	// `server` IP Address or hostname of the remote server
	Server *string `json:"server,omitempty"`
	// `share` Share name
	Share *string `json:"share,omitempty"`
	// `path` path on the remote SMB server. Use \"\" to separate path components
	Path *string `json:"path,omitempty"`
	// `username` username to use for authentication
	Username *string `json:"username,omitempty"`
	// `password` password to use for authentication
	Password *string   `json:"password,omitempty"`
	Options  *Options3 `json:"options,omitempty"`
}

// NewSmbGetRemoteAcl0 instantiates a new SmbGetRemoteAcl0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbGetRemoteAcl0() *SmbGetRemoteAcl0 {
	this := SmbGetRemoteAcl0{}
	var path string
	this.Path = &path
	var options Options3
	this.Options = &options
	return &this
}

// NewSmbGetRemoteAcl0WithDefaults instantiates a new SmbGetRemoteAcl0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbGetRemoteAcl0WithDefaults() *SmbGetRemoteAcl0 {
	this := SmbGetRemoteAcl0{}
	var path string
	this.Path = &path
	var options Options3
	this.Options = &options
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *SmbGetRemoteAcl0) SetServer(v string) {
	o.Server = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0) GetShare() string {
	if o == nil || IsNil(o.Share) {
		var ret string
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0) GetShareOk() (*string, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0) HasShare() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given string and assigns it to the Share field.
func (o *SmbGetRemoteAcl0) SetShare(v string) {
	o.Share = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SmbGetRemoteAcl0) SetPath(v string) {
	o.Path = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SmbGetRemoteAcl0) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SmbGetRemoteAcl0) SetPassword(v string) {
	o.Password = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0) GetOptions() Options3 {
	if o == nil || IsNil(o.Options) {
		var ret Options3
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0) GetOptionsOk() (*Options3, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Options3 and assigns it to the Options field.
func (o *SmbGetRemoteAcl0) SetOptions(v Options3) {
	o.Options = &v
}

func (o SmbGetRemoteAcl0) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmbGetRemoteAcl0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Share) {
		toSerialize["share"] = o.Share
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableSmbGetRemoteAcl0 struct {
	value *SmbGetRemoteAcl0
	isSet bool
}

func (v NullableSmbGetRemoteAcl0) Get() *SmbGetRemoteAcl0 {
	return v.value
}

func (v *NullableSmbGetRemoteAcl0) Set(val *SmbGetRemoteAcl0) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbGetRemoteAcl0) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbGetRemoteAcl0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbGetRemoteAcl0(val *SmbGetRemoteAcl0) *NullableSmbGetRemoteAcl0 {
	return &NullableSmbGetRemoteAcl0{value: val, isSet: true}
}

func (v NullableSmbGetRemoteAcl0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbGetRemoteAcl0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
