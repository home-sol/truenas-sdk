# IscsiHostGetInstance1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | Pointer to **bool** |  | [optional] [default to true]
**Extend** | Pointer to **NullableString** |  | [optional] 
**ExtendContext** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Extra** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**OrderBy** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Select** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Count** | Pointer to **bool** |  | [optional] [default to false]
**Get** | Pointer to **bool** |  | [optional] [default to false]
**Offset** | Pointer to **int32** |  | [optional] [default to 0]
**Limit** | Pointer to **int32** |  | [optional] [default to 0]
**ForceSqlFilters** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewIscsiHostGetInstance1

`func NewIscsiHostGetInstance1() *IscsiHostGetInstance1`

NewIscsiHostGetInstance1 instantiates a new IscsiHostGetInstance1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiHostGetInstance1WithDefaults

`func NewIscsiHostGetInstance1WithDefaults() *IscsiHostGetInstance1`

NewIscsiHostGetInstance1WithDefaults instantiates a new IscsiHostGetInstance1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *IscsiHostGetInstance1) GetRelationships() bool`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IscsiHostGetInstance1) GetRelationshipsOk() (*bool, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IscsiHostGetInstance1) SetRelationships(v bool)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IscsiHostGetInstance1) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetExtend

`func (o *IscsiHostGetInstance1) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *IscsiHostGetInstance1) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *IscsiHostGetInstance1) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *IscsiHostGetInstance1) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### SetExtendNil

`func (o *IscsiHostGetInstance1) SetExtendNil(b bool)`

 SetExtendNil sets the value for Extend to be an explicit nil

### UnsetExtend
`func (o *IscsiHostGetInstance1) UnsetExtend()`

UnsetExtend ensures that no value is present for Extend, not even an explicit nil
### GetExtendContext

`func (o *IscsiHostGetInstance1) GetExtendContext() string`

GetExtendContext returns the ExtendContext field if non-nil, zero value otherwise.

### GetExtendContextOk

`func (o *IscsiHostGetInstance1) GetExtendContextOk() (*string, bool)`

GetExtendContextOk returns a tuple with the ExtendContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendContext

`func (o *IscsiHostGetInstance1) SetExtendContext(v string)`

SetExtendContext sets ExtendContext field to given value.

### HasExtendContext

`func (o *IscsiHostGetInstance1) HasExtendContext() bool`

HasExtendContext returns a boolean if a field has been set.

### SetExtendContextNil

`func (o *IscsiHostGetInstance1) SetExtendContextNil(b bool)`

 SetExtendContextNil sets the value for ExtendContext to be an explicit nil

### UnsetExtendContext
`func (o *IscsiHostGetInstance1) UnsetExtendContext()`

UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
### GetPrefix

`func (o *IscsiHostGetInstance1) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IscsiHostGetInstance1) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IscsiHostGetInstance1) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IscsiHostGetInstance1) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *IscsiHostGetInstance1) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *IscsiHostGetInstance1) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetExtra

`func (o *IscsiHostGetInstance1) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *IscsiHostGetInstance1) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *IscsiHostGetInstance1) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *IscsiHostGetInstance1) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetOrderBy

`func (o *IscsiHostGetInstance1) GetOrderBy() []interface{}`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *IscsiHostGetInstance1) GetOrderByOk() (*[]interface{}, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *IscsiHostGetInstance1) SetOrderBy(v []interface{})`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *IscsiHostGetInstance1) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSelect

`func (o *IscsiHostGetInstance1) GetSelect() []interface{}`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *IscsiHostGetInstance1) GetSelectOk() (*[]interface{}, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *IscsiHostGetInstance1) SetSelect(v []interface{})`

SetSelect sets Select field to given value.

### HasSelect

`func (o *IscsiHostGetInstance1) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetCount

`func (o *IscsiHostGetInstance1) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IscsiHostGetInstance1) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IscsiHostGetInstance1) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *IscsiHostGetInstance1) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGet

`func (o *IscsiHostGetInstance1) GetGet() bool`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *IscsiHostGetInstance1) GetGetOk() (*bool, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *IscsiHostGetInstance1) SetGet(v bool)`

SetGet sets Get field to given value.

### HasGet

`func (o *IscsiHostGetInstance1) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetOffset

`func (o *IscsiHostGetInstance1) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IscsiHostGetInstance1) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IscsiHostGetInstance1) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IscsiHostGetInstance1) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *IscsiHostGetInstance1) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IscsiHostGetInstance1) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IscsiHostGetInstance1) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IscsiHostGetInstance1) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetForceSqlFilters

`func (o *IscsiHostGetInstance1) GetForceSqlFilters() bool`

GetForceSqlFilters returns the ForceSqlFilters field if non-nil, zero value otherwise.

### GetForceSqlFiltersOk

`func (o *IscsiHostGetInstance1) GetForceSqlFiltersOk() (*bool, bool)`

GetForceSqlFiltersOk returns a tuple with the ForceSqlFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSqlFilters

`func (o *IscsiHostGetInstance1) SetForceSqlFilters(v bool)`

SetForceSqlFilters sets ForceSqlFilters field to given value.

### HasForceSqlFilters

`func (o *IscsiHostGetInstance1) HasForceSqlFilters() bool`

HasForceSqlFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

