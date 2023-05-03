# VmwareGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**VmwareGetInstance0**](VmwareGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**VmwareGetInstance1**](VmwareGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewVmwareGetInstance

`func NewVmwareGetInstance() *VmwareGetInstance`

NewVmwareGetInstance instantiates a new VmwareGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareGetInstanceWithDefaults

`func NewVmwareGetInstanceWithDefaults() *VmwareGetInstance`

NewVmwareGetInstanceWithDefaults instantiates a new VmwareGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmwareGetInstance) GetId() VmwareGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmwareGetInstance) GetIdOk() (*VmwareGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmwareGetInstance) SetId(v VmwareGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *VmwareGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *VmwareGetInstance) GetQueryOptions() VmwareGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *VmwareGetInstance) GetQueryOptionsOk() (*VmwareGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *VmwareGetInstance) SetQueryOptions(v VmwareGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *VmwareGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


