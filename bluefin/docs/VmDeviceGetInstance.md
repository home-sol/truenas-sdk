# VmDeviceGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**VmDeviceGetInstance0**](VmDeviceGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**VmDeviceGetInstance1**](VmDeviceGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewVmDeviceGetInstance

`func NewVmDeviceGetInstance() *VmDeviceGetInstance`

NewVmDeviceGetInstance instantiates a new VmDeviceGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDeviceGetInstanceWithDefaults

`func NewVmDeviceGetInstanceWithDefaults() *VmDeviceGetInstance`

NewVmDeviceGetInstanceWithDefaults instantiates a new VmDeviceGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmDeviceGetInstance) GetId() VmDeviceGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmDeviceGetInstance) GetIdOk() (*VmDeviceGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmDeviceGetInstance) SetId(v VmDeviceGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *VmDeviceGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *VmDeviceGetInstance) GetQueryOptions() VmDeviceGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *VmDeviceGetInstance) GetQueryOptionsOk() (*VmDeviceGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *VmDeviceGetInstance) SetQueryOptions(v VmDeviceGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *VmDeviceGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


