# VmGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**VmGetInstance0**](VmGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**VmGetInstance1**](VmGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewVmGetInstance

`func NewVmGetInstance() *VmGetInstance`

NewVmGetInstance instantiates a new VmGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmGetInstanceWithDefaults

`func NewVmGetInstanceWithDefaults() *VmGetInstance`

NewVmGetInstanceWithDefaults instantiates a new VmGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmGetInstance) GetId() VmGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmGetInstance) GetIdOk() (*VmGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmGetInstance) SetId(v VmGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *VmGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *VmGetInstance) GetQueryOptions() VmGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *VmGetInstance) GetQueryOptionsOk() (*VmGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *VmGetInstance) SetQueryOptions(v VmGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *VmGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


