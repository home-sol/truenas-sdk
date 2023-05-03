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

// checks if the CloudsyncCreateBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudsyncCreateBucket{}

// CloudsyncCreateBucket struct for CloudsyncCreateBucket
type CloudsyncCreateBucket struct {
	CredentialsId        *int32  `json:"credentials_id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudsyncCreateBucket CloudsyncCreateBucket

// NewCloudsyncCreateBucket instantiates a new CloudsyncCreateBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudsyncCreateBucket() *CloudsyncCreateBucket {
	this := CloudsyncCreateBucket{}
	return &this
}

// NewCloudsyncCreateBucketWithDefaults instantiates a new CloudsyncCreateBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudsyncCreateBucketWithDefaults() *CloudsyncCreateBucket {
	this := CloudsyncCreateBucket{}
	return &this
}

// GetCredentialsId returns the CredentialsId field value if set, zero value otherwise.
func (o *CloudsyncCreateBucket) GetCredentialsId() int32 {
	if o == nil || IsNil(o.CredentialsId) {
		var ret int32
		return ret
	}
	return *o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncCreateBucket) GetCredentialsIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CredentialsId) {
		return nil, false
	}
	return o.CredentialsId, true
}

// HasCredentialsId returns a boolean if a field has been set.
func (o *CloudsyncCreateBucket) HasCredentialsId() bool {
	if o != nil && !IsNil(o.CredentialsId) {
		return true
	}

	return false
}

// SetCredentialsId gets a reference to the given int32 and assigns it to the CredentialsId field.
func (o *CloudsyncCreateBucket) SetCredentialsId(v int32) {
	o.CredentialsId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudsyncCreateBucket) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncCreateBucket) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudsyncCreateBucket) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudsyncCreateBucket) SetName(v string) {
	o.Name = &v
}

func (o CloudsyncCreateBucket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudsyncCreateBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialsId) {
		toSerialize["credentials_id"] = o.CredentialsId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudsyncCreateBucket) UnmarshalJSON(bytes []byte) (err error) {
	varCloudsyncCreateBucket := _CloudsyncCreateBucket{}

	if err = json.Unmarshal(bytes, &varCloudsyncCreateBucket); err == nil {
		*o = CloudsyncCreateBucket(varCloudsyncCreateBucket)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials_id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudsyncCreateBucket struct {
	value *CloudsyncCreateBucket
	isSet bool
}

func (v NullableCloudsyncCreateBucket) Get() *CloudsyncCreateBucket {
	return v.value
}

func (v *NullableCloudsyncCreateBucket) Set(val *CloudsyncCreateBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudsyncCreateBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudsyncCreateBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudsyncCreateBucket(val *CloudsyncCreateBucket) *NullableCloudsyncCreateBucket {
	return &NullableCloudsyncCreateBucket{value: val, isSet: true}
}

func (v NullableCloudsyncCreateBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudsyncCreateBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}