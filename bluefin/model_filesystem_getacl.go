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

// checks if the FilesystemGetacl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesystemGetacl{}

// FilesystemGetacl struct for FilesystemGetacl
type FilesystemGetacl struct {
	Path *string `json:"path,omitempty"`
	// `simplified` - effect of this depends on ACL type on underlying filesystem. In the case of NFSv4 ACLs simplified permissions and flags are returned for ACL entries where applicable. NFSv4 errata below. In the case of POSIX1E ACls, this setting has no impact on returned ACL. `simplified` returns a shortened form of the ACL permset and flags where applicable. If permissions have been simplified, then the `perms` object will contain only a single `BASIC` key with a string describing the underlying permissions set.
	Simplified *bool `json:"simplified,omitempty"`
	// `resolve_ids` - adds additional `who` key to each ACL entry, that converts the numeric id to a user name or group name. In the case of owner@ and group@ (NFSv4) or USER_OBJ and GROUP_OBJ (POSIX1E), st_uid or st_gid will be converted from stat() return for file. In the case of MASK (POSIX1E), OTHER (POSIX1E), everyone@ (NFSv4), key `who` will be included, but set to null. In case of failure to resolve the id to a name, `who` will be set to null. This option should only be used if resolving ids to names is required.
	ResolveIds           *bool `json:"resolve_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FilesystemGetacl FilesystemGetacl

// NewFilesystemGetacl instantiates a new FilesystemGetacl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemGetacl() *FilesystemGetacl {
	this := FilesystemGetacl{}
	var simplified bool
	this.Simplified = &simplified
	var resolveIds bool
	this.ResolveIds = &resolveIds
	return &this
}

// NewFilesystemGetaclWithDefaults instantiates a new FilesystemGetacl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemGetaclWithDefaults() *FilesystemGetacl {
	this := FilesystemGetacl{}
	var simplified bool
	this.Simplified = &simplified
	var resolveIds bool
	this.ResolveIds = &resolveIds
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FilesystemGetacl) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemGetacl) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FilesystemGetacl) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FilesystemGetacl) SetPath(v string) {
	o.Path = &v
}

// GetSimplified returns the Simplified field value if set, zero value otherwise.
func (o *FilesystemGetacl) GetSimplified() bool {
	if o == nil || IsNil(o.Simplified) {
		var ret bool
		return ret
	}
	return *o.Simplified
}

// GetSimplifiedOk returns a tuple with the Simplified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemGetacl) GetSimplifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Simplified) {
		return nil, false
	}
	return o.Simplified, true
}

// HasSimplified returns a boolean if a field has been set.
func (o *FilesystemGetacl) HasSimplified() bool {
	if o != nil && !IsNil(o.Simplified) {
		return true
	}

	return false
}

// SetSimplified gets a reference to the given bool and assigns it to the Simplified field.
func (o *FilesystemGetacl) SetSimplified(v bool) {
	o.Simplified = &v
}

// GetResolveIds returns the ResolveIds field value if set, zero value otherwise.
func (o *FilesystemGetacl) GetResolveIds() bool {
	if o == nil || IsNil(o.ResolveIds) {
		var ret bool
		return ret
	}
	return *o.ResolveIds
}

// GetResolveIdsOk returns a tuple with the ResolveIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemGetacl) GetResolveIdsOk() (*bool, bool) {
	if o == nil || IsNil(o.ResolveIds) {
		return nil, false
	}
	return o.ResolveIds, true
}

// HasResolveIds returns a boolean if a field has been set.
func (o *FilesystemGetacl) HasResolveIds() bool {
	if o != nil && !IsNil(o.ResolveIds) {
		return true
	}

	return false
}

// SetResolveIds gets a reference to the given bool and assigns it to the ResolveIds field.
func (o *FilesystemGetacl) SetResolveIds(v bool) {
	o.ResolveIds = &v
}

func (o FilesystemGetacl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesystemGetacl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Simplified) {
		toSerialize["simplified"] = o.Simplified
	}
	if !IsNil(o.ResolveIds) {
		toSerialize["resolve_ids"] = o.ResolveIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesystemGetacl) UnmarshalJSON(bytes []byte) (err error) {
	varFilesystemGetacl := _FilesystemGetacl{}

	if err = json.Unmarshal(bytes, &varFilesystemGetacl); err == nil {
		*o = FilesystemGetacl(varFilesystemGetacl)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "simplified")
		delete(additionalProperties, "resolve_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesystemGetacl struct {
	value *FilesystemGetacl
	isSet bool
}

func (v NullableFilesystemGetacl) Get() *FilesystemGetacl {
	return v.value
}

func (v *NullableFilesystemGetacl) Set(val *FilesystemGetacl) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemGetacl) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemGetacl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemGetacl(val *FilesystemGetacl) *NullableFilesystemGetacl {
	return &NullableFilesystemGetacl{value: val, isSet: true}
}

func (v NullableFilesystemGetacl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemGetacl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}