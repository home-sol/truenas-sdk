# CloudsyncCredentialsGetInstance1

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

### NewCloudsyncCredentialsGetInstance1

`func NewCloudsyncCredentialsGetInstance1() *CloudsyncCredentialsGetInstance1`

NewCloudsyncCredentialsGetInstance1 instantiates a new CloudsyncCredentialsGetInstance1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudsyncCredentialsGetInstance1WithDefaults

`func NewCloudsyncCredentialsGetInstance1WithDefaults() *CloudsyncCredentialsGetInstance1`

NewCloudsyncCredentialsGetInstance1WithDefaults instantiates a new CloudsyncCredentialsGetInstance1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *CloudsyncCredentialsGetInstance1) GetRelationships() bool`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CloudsyncCredentialsGetInstance1) GetRelationshipsOk() (*bool, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CloudsyncCredentialsGetInstance1) SetRelationships(v bool)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CloudsyncCredentialsGetInstance1) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetExtend

`func (o *CloudsyncCredentialsGetInstance1) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *CloudsyncCredentialsGetInstance1) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *CloudsyncCredentialsGetInstance1) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *CloudsyncCredentialsGetInstance1) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### SetExtendNil

`func (o *CloudsyncCredentialsGetInstance1) SetExtendNil(b bool)`

 SetExtendNil sets the value for Extend to be an explicit nil

### UnsetExtend
`func (o *CloudsyncCredentialsGetInstance1) UnsetExtend()`

UnsetExtend ensures that no value is present for Extend, not even an explicit nil
### GetExtendContext

`func (o *CloudsyncCredentialsGetInstance1) GetExtendContext() string`

GetExtendContext returns the ExtendContext field if non-nil, zero value otherwise.

### GetExtendContextOk

`func (o *CloudsyncCredentialsGetInstance1) GetExtendContextOk() (*string, bool)`

GetExtendContextOk returns a tuple with the ExtendContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendContext

`func (o *CloudsyncCredentialsGetInstance1) SetExtendContext(v string)`

SetExtendContext sets ExtendContext field to given value.

### HasExtendContext

`func (o *CloudsyncCredentialsGetInstance1) HasExtendContext() bool`

HasExtendContext returns a boolean if a field has been set.

### SetExtendContextNil

`func (o *CloudsyncCredentialsGetInstance1) SetExtendContextNil(b bool)`

 SetExtendContextNil sets the value for ExtendContext to be an explicit nil

### UnsetExtendContext
`func (o *CloudsyncCredentialsGetInstance1) UnsetExtendContext()`

UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
### GetPrefix

`func (o *CloudsyncCredentialsGetInstance1) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CloudsyncCredentialsGetInstance1) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CloudsyncCredentialsGetInstance1) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CloudsyncCredentialsGetInstance1) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *CloudsyncCredentialsGetInstance1) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *CloudsyncCredentialsGetInstance1) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetExtra

`func (o *CloudsyncCredentialsGetInstance1) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CloudsyncCredentialsGetInstance1) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CloudsyncCredentialsGetInstance1) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *CloudsyncCredentialsGetInstance1) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetOrderBy

`func (o *CloudsyncCredentialsGetInstance1) GetOrderBy() []interface{}`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *CloudsyncCredentialsGetInstance1) GetOrderByOk() (*[]interface{}, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *CloudsyncCredentialsGetInstance1) SetOrderBy(v []interface{})`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *CloudsyncCredentialsGetInstance1) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSelect

`func (o *CloudsyncCredentialsGetInstance1) GetSelect() []interface{}`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *CloudsyncCredentialsGetInstance1) GetSelectOk() (*[]interface{}, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *CloudsyncCredentialsGetInstance1) SetSelect(v []interface{})`

SetSelect sets Select field to given value.

### HasSelect

`func (o *CloudsyncCredentialsGetInstance1) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetCount

`func (o *CloudsyncCredentialsGetInstance1) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudsyncCredentialsGetInstance1) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudsyncCredentialsGetInstance1) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudsyncCredentialsGetInstance1) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGet

`func (o *CloudsyncCredentialsGetInstance1) GetGet() bool`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *CloudsyncCredentialsGetInstance1) GetGetOk() (*bool, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *CloudsyncCredentialsGetInstance1) SetGet(v bool)`

SetGet sets Get field to given value.

### HasGet

`func (o *CloudsyncCredentialsGetInstance1) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetOffset

`func (o *CloudsyncCredentialsGetInstance1) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CloudsyncCredentialsGetInstance1) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CloudsyncCredentialsGetInstance1) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CloudsyncCredentialsGetInstance1) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *CloudsyncCredentialsGetInstance1) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudsyncCredentialsGetInstance1) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudsyncCredentialsGetInstance1) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudsyncCredentialsGetInstance1) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetForceSqlFilters

`func (o *CloudsyncCredentialsGetInstance1) GetForceSqlFilters() bool`

GetForceSqlFilters returns the ForceSqlFilters field if non-nil, zero value otherwise.

### GetForceSqlFiltersOk

`func (o *CloudsyncCredentialsGetInstance1) GetForceSqlFiltersOk() (*bool, bool)`

GetForceSqlFiltersOk returns a tuple with the ForceSqlFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSqlFilters

`func (o *CloudsyncCredentialsGetInstance1) SetForceSqlFilters(v bool)`

SetForceSqlFilters sets ForceSqlFilters field to given value.

### HasForceSqlFilters

`func (o *CloudsyncCredentialsGetInstance1) HasForceSqlFilters() bool`

HasForceSqlFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

