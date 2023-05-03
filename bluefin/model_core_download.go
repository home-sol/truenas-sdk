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

// checks if the CoreDownload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreDownload{}

// CoreDownload struct for CoreDownload
type CoreDownload struct {
	Method   *string       `json:"method,omitempty"`
	Args     []interface{} `json:"args,omitempty"`
	Filename *string       `json:"filename,omitempty"`
	// Non-`buffered` downloads will allow job to write to pipe as soon as download URL is requested, job will stay blocked meanwhile. `buffered` downloads must wait for job to complete before requesting download URL, job's pipe output will be buffered to ramfs.
	Buffered             *bool `json:"buffered,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CoreDownload CoreDownload

// NewCoreDownload instantiates a new CoreDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDownload() *CoreDownload {
	this := CoreDownload{}
	var buffered bool
	this.Buffered = &buffered
	return &this
}

// NewCoreDownloadWithDefaults instantiates a new CoreDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDownloadWithDefaults() *CoreDownload {
	this := CoreDownload{}
	var buffered bool
	this.Buffered = &buffered
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CoreDownload) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDownload) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CoreDownload) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CoreDownload) SetMethod(v string) {
	o.Method = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *CoreDownload) GetArgs() []interface{} {
	if o == nil || IsNil(o.Args) {
		var ret []interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDownload) GetArgsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *CoreDownload) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []interface{} and assigns it to the Args field.
func (o *CoreDownload) SetArgs(v []interface{}) {
	o.Args = v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CoreDownload) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDownload) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CoreDownload) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CoreDownload) SetFilename(v string) {
	o.Filename = &v
}

// GetBuffered returns the Buffered field value if set, zero value otherwise.
func (o *CoreDownload) GetBuffered() bool {
	if o == nil || IsNil(o.Buffered) {
		var ret bool
		return ret
	}
	return *o.Buffered
}

// GetBufferedOk returns a tuple with the Buffered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDownload) GetBufferedOk() (*bool, bool) {
	if o == nil || IsNil(o.Buffered) {
		return nil, false
	}
	return o.Buffered, true
}

// HasBuffered returns a boolean if a field has been set.
func (o *CoreDownload) HasBuffered() bool {
	if o != nil && !IsNil(o.Buffered) {
		return true
	}

	return false
}

// SetBuffered gets a reference to the given bool and assigns it to the Buffered field.
func (o *CoreDownload) SetBuffered(v bool) {
	o.Buffered = &v
}

func (o CoreDownload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreDownload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Buffered) {
		toSerialize["buffered"] = o.Buffered
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreDownload) UnmarshalJSON(bytes []byte) (err error) {
	varCoreDownload := _CoreDownload{}

	if err = json.Unmarshal(bytes, &varCoreDownload); err == nil {
		*o = CoreDownload(varCoreDownload)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "args")
		delete(additionalProperties, "filename")
		delete(additionalProperties, "buffered")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreDownload struct {
	value *CoreDownload
	isSet bool
}

func (v NullableCoreDownload) Get() *CoreDownload {
	return v.value
}

func (v *NullableCoreDownload) Set(val *CoreDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDownload(val *CoreDownload) *NullableCoreDownload {
	return &NullableCoreDownload{value: val, isSet: true}
}

func (v NullableCoreDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}