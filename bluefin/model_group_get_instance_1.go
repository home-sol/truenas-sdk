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

// checks if the GroupGetInstance1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupGetInstance1{}

// GroupGetInstance1 struct for GroupGetInstance1
type GroupGetInstance1 struct {
	Relationships   *bool                  `json:"relationships,omitempty"`
	Extend          NullableString         `json:"extend,omitempty"`
	ExtendContext   NullableString         `json:"extend_context,omitempty"`
	Prefix          NullableString         `json:"prefix,omitempty"`
	Extra           map[string]interface{} `json:"extra,omitempty"`
	OrderBy         []interface{}          `json:"order_by,omitempty"`
	Select          []interface{}          `json:"select,omitempty"`
	Count           *bool                  `json:"count,omitempty"`
	Get             *bool                  `json:"get,omitempty"`
	Offset          *int32                 `json:"offset,omitempty"`
	Limit           *int32                 `json:"limit,omitempty"`
	ForceSqlFilters *bool                  `json:"force_sql_filters,omitempty"`
}

// NewGroupGetInstance1 instantiates a new GroupGetInstance1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupGetInstance1() *GroupGetInstance1 {
	this := GroupGetInstance1{}
	var relationships bool
	this.Relationships = &relationships
	var count bool
	this.Count = &count
	var get bool
	this.Get = &get
	var offset int32
	this.Offset = &offset
	var limit int32
	this.Limit = &limit
	var forceSqlFilters bool
	this.ForceSqlFilters = &forceSqlFilters
	return &this
}

// NewGroupGetInstance1WithDefaults instantiates a new GroupGetInstance1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupGetInstance1WithDefaults() *GroupGetInstance1 {
	this := GroupGetInstance1{}
	var relationships bool
	this.Relationships = &relationships
	var count bool
	this.Count = &count
	var get bool
	this.Get = &get
	var offset int32
	this.Offset = &offset
	var limit int32
	this.Limit = &limit
	var forceSqlFilters bool
	this.ForceSqlFilters = &forceSqlFilters
	return &this
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetRelationships() bool {
	if o == nil || IsNil(o.Relationships) {
		var ret bool
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetRelationshipsOk() (*bool, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given bool and assigns it to the Relationships field.
func (o *GroupGetInstance1) SetRelationships(v bool) {
	o.Relationships = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGetInstance1) GetExtend() string {
	if o == nil || IsNil(o.Extend.Get()) {
		var ret string
		return ret
	}
	return *o.Extend.Get()
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGetInstance1) GetExtendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extend.Get(), o.Extend.IsSet()
}

// HasExtend returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasExtend() bool {
	if o != nil && o.Extend.IsSet() {
		return true
	}

	return false
}

// SetExtend gets a reference to the given NullableString and assigns it to the Extend field.
func (o *GroupGetInstance1) SetExtend(v string) {
	o.Extend.Set(&v)
}

// SetExtendNil sets the value for Extend to be an explicit nil
func (o *GroupGetInstance1) SetExtendNil() {
	o.Extend.Set(nil)
}

// UnsetExtend ensures that no value is present for Extend, not even an explicit nil
func (o *GroupGetInstance1) UnsetExtend() {
	o.Extend.Unset()
}

// GetExtendContext returns the ExtendContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGetInstance1) GetExtendContext() string {
	if o == nil || IsNil(o.ExtendContext.Get()) {
		var ret string
		return ret
	}
	return *o.ExtendContext.Get()
}

// GetExtendContextOk returns a tuple with the ExtendContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGetInstance1) GetExtendContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtendContext.Get(), o.ExtendContext.IsSet()
}

// HasExtendContext returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasExtendContext() bool {
	if o != nil && o.ExtendContext.IsSet() {
		return true
	}

	return false
}

// SetExtendContext gets a reference to the given NullableString and assigns it to the ExtendContext field.
func (o *GroupGetInstance1) SetExtendContext(v string) {
	o.ExtendContext.Set(&v)
}

// SetExtendContextNil sets the value for ExtendContext to be an explicit nil
func (o *GroupGetInstance1) SetExtendContextNil() {
	o.ExtendContext.Set(nil)
}

// UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
func (o *GroupGetInstance1) UnsetExtendContext() {
	o.ExtendContext.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGetInstance1) GetPrefix() string {
	if o == nil || IsNil(o.Prefix.Get()) {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGetInstance1) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *GroupGetInstance1) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *GroupGetInstance1) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *GroupGetInstance1) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetExtra() map[string]interface{} {
	if o == nil || IsNil(o.Extra) {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetExtraOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Extra) {
		return map[string]interface{}{}, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *GroupGetInstance1) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetOrderBy() []interface{} {
	if o == nil || IsNil(o.OrderBy) {
		var ret []interface{}
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetOrderByOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []interface{} and assigns it to the OrderBy field.
func (o *GroupGetInstance1) SetOrderBy(v []interface{}) {
	o.OrderBy = v
}

// GetSelect returns the Select field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetSelect() []interface{} {
	if o == nil || IsNil(o.Select) {
		var ret []interface{}
		return ret
	}
	return o.Select
}

// GetSelectOk returns a tuple with the Select field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetSelectOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Select) {
		return nil, false
	}
	return o.Select, true
}

// HasSelect returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasSelect() bool {
	if o != nil && !IsNil(o.Select) {
		return true
	}

	return false
}

// SetSelect gets a reference to the given []interface{} and assigns it to the Select field.
func (o *GroupGetInstance1) SetSelect(v []interface{}) {
	o.Select = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetCount() bool {
	if o == nil || IsNil(o.Count) {
		var ret bool
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetCountOk() (*bool, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given bool and assigns it to the Count field.
func (o *GroupGetInstance1) SetCount(v bool) {
	o.Count = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetGet() bool {
	if o == nil || IsNil(o.Get) {
		var ret bool
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetGetOk() (*bool, bool) {
	if o == nil || IsNil(o.Get) {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasGet() bool {
	if o != nil && !IsNil(o.Get) {
		return true
	}

	return false
}

// SetGet gets a reference to the given bool and assigns it to the Get field.
func (o *GroupGetInstance1) SetGet(v bool) {
	o.Get = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GroupGetInstance1) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GroupGetInstance1) SetLimit(v int32) {
	o.Limit = &v
}

// GetForceSqlFilters returns the ForceSqlFilters field value if set, zero value otherwise.
func (o *GroupGetInstance1) GetForceSqlFilters() bool {
	if o == nil || IsNil(o.ForceSqlFilters) {
		var ret bool
		return ret
	}
	return *o.ForceSqlFilters
}

// GetForceSqlFiltersOk returns a tuple with the ForceSqlFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGetInstance1) GetForceSqlFiltersOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceSqlFilters) {
		return nil, false
	}
	return o.ForceSqlFilters, true
}

// HasForceSqlFilters returns a boolean if a field has been set.
func (o *GroupGetInstance1) HasForceSqlFilters() bool {
	if o != nil && !IsNil(o.ForceSqlFilters) {
		return true
	}

	return false
}

// SetForceSqlFilters gets a reference to the given bool and assigns it to the ForceSqlFilters field.
func (o *GroupGetInstance1) SetForceSqlFilters(v bool) {
	o.ForceSqlFilters = &v
}

func (o GroupGetInstance1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupGetInstance1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Extend.IsSet() {
		toSerialize["extend"] = o.Extend.Get()
	}
	if o.ExtendContext.IsSet() {
		toSerialize["extend_context"] = o.ExtendContext.Get()
	}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !IsNil(o.OrderBy) {
		toSerialize["order_by"] = o.OrderBy
	}
	if !IsNil(o.Select) {
		toSerialize["select"] = o.Select
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Get) {
		toSerialize["get"] = o.Get
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.ForceSqlFilters) {
		toSerialize["force_sql_filters"] = o.ForceSqlFilters
	}
	return toSerialize, nil
}

type NullableGroupGetInstance1 struct {
	value *GroupGetInstance1
	isSet bool
}

func (v NullableGroupGetInstance1) Get() *GroupGetInstance1 {
	return v.value
}

func (v *NullableGroupGetInstance1) Set(val *GroupGetInstance1) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupGetInstance1) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupGetInstance1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupGetInstance1(val *GroupGetInstance1) *NullableGroupGetInstance1 {
	return &NullableGroupGetInstance1{value: val, isSet: true}
}

func (v NullableGroupGetInstance1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupGetInstance1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
