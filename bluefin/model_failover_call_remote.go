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

// checks if the FailoverCallRemote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailoverCallRemote{}

// FailoverCallRemote struct for FailoverCallRemote
type FailoverCallRemote struct {
	// `method` name of the method to be called `args` list of arguments to be passed to `method`
	Method *string `json:"method,omitempty"`
	// `method` name of the method to be called `args` list of arguments to be passed to `method`
	Args                 []interface{}        `json:"args,omitempty"`
	Options              *FailoverCallRemote2 `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FailoverCallRemote FailoverCallRemote

// NewFailoverCallRemote instantiates a new FailoverCallRemote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverCallRemote() *FailoverCallRemote {
	this := FailoverCallRemote{}
	var options FailoverCallRemote2
	this.Options = &options
	return &this
}

// NewFailoverCallRemoteWithDefaults instantiates a new FailoverCallRemote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverCallRemoteWithDefaults() *FailoverCallRemote {
	this := FailoverCallRemote{}
	var options FailoverCallRemote2
	this.Options = &options
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *FailoverCallRemote) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverCallRemote) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *FailoverCallRemote) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *FailoverCallRemote) SetMethod(v string) {
	o.Method = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *FailoverCallRemote) GetArgs() []interface{} {
	if o == nil || IsNil(o.Args) {
		var ret []interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverCallRemote) GetArgsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *FailoverCallRemote) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []interface{} and assigns it to the Args field.
func (o *FailoverCallRemote) SetArgs(v []interface{}) {
	o.Args = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FailoverCallRemote) GetOptions() FailoverCallRemote2 {
	if o == nil || IsNil(o.Options) {
		var ret FailoverCallRemote2
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverCallRemote) GetOptionsOk() (*FailoverCallRemote2, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FailoverCallRemote) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given FailoverCallRemote2 and assigns it to the Options field.
func (o *FailoverCallRemote) SetOptions(v FailoverCallRemote2) {
	o.Options = &v
}

func (o FailoverCallRemote) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailoverCallRemote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FailoverCallRemote) UnmarshalJSON(bytes []byte) (err error) {
	varFailoverCallRemote := _FailoverCallRemote{}

	if err = json.Unmarshal(bytes, &varFailoverCallRemote); err == nil {
		*o = FailoverCallRemote(varFailoverCallRemote)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "args")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFailoverCallRemote struct {
	value *FailoverCallRemote
	isSet bool
}

func (v NullableFailoverCallRemote) Get() *FailoverCallRemote {
	return v.value
}

func (v *NullableFailoverCallRemote) Set(val *FailoverCallRemote) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverCallRemote) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverCallRemote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverCallRemote(val *FailoverCallRemote) *NullableFailoverCallRemote {
	return &NullableFailoverCallRemote{value: val, isSet: true}
}

func (v NullableFailoverCallRemote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverCallRemote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}