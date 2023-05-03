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

// checks if the ReplicationTargetUnmatchedSnapshots type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicationTargetUnmatchedSnapshots{}

// ReplicationTargetUnmatchedSnapshots struct for ReplicationTargetUnmatchedSnapshots
type ReplicationTargetUnmatchedSnapshots struct {
	Direction            *ReplicationTargetUnmatchedSnapshots0 `json:"direction,omitempty"`
	SourceDatasets       []string                              `json:"source_datasets,omitempty"`
	TargetDataset        *string                               `json:"target_dataset,omitempty"`
	Transport            *ReplicationTargetUnmatchedSnapshots3 `json:"transport,omitempty"`
	SshCredentials       NullableInt32                         `json:"ssh_credentials,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReplicationTargetUnmatchedSnapshots ReplicationTargetUnmatchedSnapshots

// NewReplicationTargetUnmatchedSnapshots instantiates a new ReplicationTargetUnmatchedSnapshots object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationTargetUnmatchedSnapshots() *ReplicationTargetUnmatchedSnapshots {
	this := ReplicationTargetUnmatchedSnapshots{}
	return &this
}

// NewReplicationTargetUnmatchedSnapshotsWithDefaults instantiates a new ReplicationTargetUnmatchedSnapshots object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationTargetUnmatchedSnapshotsWithDefaults() *ReplicationTargetUnmatchedSnapshots {
	this := ReplicationTargetUnmatchedSnapshots{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ReplicationTargetUnmatchedSnapshots) GetDirection() ReplicationTargetUnmatchedSnapshots0 {
	if o == nil || IsNil(o.Direction) {
		var ret ReplicationTargetUnmatchedSnapshots0
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationTargetUnmatchedSnapshots) GetDirectionOk() (*ReplicationTargetUnmatchedSnapshots0, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ReplicationTargetUnmatchedSnapshots) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given ReplicationTargetUnmatchedSnapshots0 and assigns it to the Direction field.
func (o *ReplicationTargetUnmatchedSnapshots) SetDirection(v ReplicationTargetUnmatchedSnapshots0) {
	o.Direction = &v
}

// GetSourceDatasets returns the SourceDatasets field value if set, zero value otherwise.
func (o *ReplicationTargetUnmatchedSnapshots) GetSourceDatasets() []string {
	if o == nil || IsNil(o.SourceDatasets) {
		var ret []string
		return ret
	}
	return o.SourceDatasets
}

// GetSourceDatasetsOk returns a tuple with the SourceDatasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationTargetUnmatchedSnapshots) GetSourceDatasetsOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceDatasets) {
		return nil, false
	}
	return o.SourceDatasets, true
}

// HasSourceDatasets returns a boolean if a field has been set.
func (o *ReplicationTargetUnmatchedSnapshots) HasSourceDatasets() bool {
	if o != nil && !IsNil(o.SourceDatasets) {
		return true
	}

	return false
}

// SetSourceDatasets gets a reference to the given []string and assigns it to the SourceDatasets field.
func (o *ReplicationTargetUnmatchedSnapshots) SetSourceDatasets(v []string) {
	o.SourceDatasets = v
}

// GetTargetDataset returns the TargetDataset field value if set, zero value otherwise.
func (o *ReplicationTargetUnmatchedSnapshots) GetTargetDataset() string {
	if o == nil || IsNil(o.TargetDataset) {
		var ret string
		return ret
	}
	return *o.TargetDataset
}

// GetTargetDatasetOk returns a tuple with the TargetDataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationTargetUnmatchedSnapshots) GetTargetDatasetOk() (*string, bool) {
	if o == nil || IsNil(o.TargetDataset) {
		return nil, false
	}
	return o.TargetDataset, true
}

// HasTargetDataset returns a boolean if a field has been set.
func (o *ReplicationTargetUnmatchedSnapshots) HasTargetDataset() bool {
	if o != nil && !IsNil(o.TargetDataset) {
		return true
	}

	return false
}

// SetTargetDataset gets a reference to the given string and assigns it to the TargetDataset field.
func (o *ReplicationTargetUnmatchedSnapshots) SetTargetDataset(v string) {
	o.TargetDataset = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *ReplicationTargetUnmatchedSnapshots) GetTransport() ReplicationTargetUnmatchedSnapshots3 {
	if o == nil || IsNil(o.Transport) {
		var ret ReplicationTargetUnmatchedSnapshots3
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationTargetUnmatchedSnapshots) GetTransportOk() (*ReplicationTargetUnmatchedSnapshots3, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *ReplicationTargetUnmatchedSnapshots) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given ReplicationTargetUnmatchedSnapshots3 and assigns it to the Transport field.
func (o *ReplicationTargetUnmatchedSnapshots) SetTransport(v ReplicationTargetUnmatchedSnapshots3) {
	o.Transport = &v
}

// GetSshCredentials returns the SshCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationTargetUnmatchedSnapshots) GetSshCredentials() int32 {
	if o == nil || IsNil(o.SshCredentials.Get()) {
		var ret int32
		return ret
	}
	return *o.SshCredentials.Get()
}

// GetSshCredentialsOk returns a tuple with the SshCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicationTargetUnmatchedSnapshots) GetSshCredentialsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshCredentials.Get(), o.SshCredentials.IsSet()
}

// HasSshCredentials returns a boolean if a field has been set.
func (o *ReplicationTargetUnmatchedSnapshots) HasSshCredentials() bool {
	if o != nil && o.SshCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshCredentials gets a reference to the given NullableInt32 and assigns it to the SshCredentials field.
func (o *ReplicationTargetUnmatchedSnapshots) SetSshCredentials(v int32) {
	o.SshCredentials.Set(&v)
}

// SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil
func (o *ReplicationTargetUnmatchedSnapshots) SetSshCredentialsNil() {
	o.SshCredentials.Set(nil)
}

// UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil
func (o *ReplicationTargetUnmatchedSnapshots) UnsetSshCredentials() {
	o.SshCredentials.Unset()
}

func (o ReplicationTargetUnmatchedSnapshots) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicationTargetUnmatchedSnapshots) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.SourceDatasets) {
		toSerialize["source_datasets"] = o.SourceDatasets
	}
	if !IsNil(o.TargetDataset) {
		toSerialize["target_dataset"] = o.TargetDataset
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	if o.SshCredentials.IsSet() {
		toSerialize["ssh_credentials"] = o.SshCredentials.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReplicationTargetUnmatchedSnapshots) UnmarshalJSON(bytes []byte) (err error) {
	varReplicationTargetUnmatchedSnapshots := _ReplicationTargetUnmatchedSnapshots{}

	if err = json.Unmarshal(bytes, &varReplicationTargetUnmatchedSnapshots); err == nil {
		*o = ReplicationTargetUnmatchedSnapshots(varReplicationTargetUnmatchedSnapshots)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "source_datasets")
		delete(additionalProperties, "target_dataset")
		delete(additionalProperties, "transport")
		delete(additionalProperties, "ssh_credentials")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReplicationTargetUnmatchedSnapshots struct {
	value *ReplicationTargetUnmatchedSnapshots
	isSet bool
}

func (v NullableReplicationTargetUnmatchedSnapshots) Get() *ReplicationTargetUnmatchedSnapshots {
	return v.value
}

func (v *NullableReplicationTargetUnmatchedSnapshots) Set(val *ReplicationTargetUnmatchedSnapshots) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationTargetUnmatchedSnapshots) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationTargetUnmatchedSnapshots) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationTargetUnmatchedSnapshots(val *ReplicationTargetUnmatchedSnapshots) *NullableReplicationTargetUnmatchedSnapshots {
	return &NullableReplicationTargetUnmatchedSnapshots{value: val, isSet: true}
}

func (v NullableReplicationTargetUnmatchedSnapshots) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationTargetUnmatchedSnapshots) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}