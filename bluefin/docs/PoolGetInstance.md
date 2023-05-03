# PoolGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**PoolGetInstance0**](PoolGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**PoolGetInstance1**](PoolGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewPoolGetInstance

`func NewPoolGetInstance() *PoolGetInstance`

NewPoolGetInstance instantiates a new PoolGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolGetInstanceWithDefaults

`func NewPoolGetInstanceWithDefaults() *PoolGetInstance`

NewPoolGetInstanceWithDefaults instantiates a new PoolGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolGetInstance) GetId() PoolGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolGetInstance) GetIdOk() (*PoolGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolGetInstance) SetId(v PoolGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *PoolGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *PoolGetInstance) GetQueryOptions() PoolGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *PoolGetInstance) GetQueryOptionsOk() (*PoolGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *PoolGetInstance) SetQueryOptions(v PoolGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *PoolGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


