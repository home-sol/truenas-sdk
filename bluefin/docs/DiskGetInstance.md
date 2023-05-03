# DiskGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**DiskGetInstance0**](DiskGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**DiskGetInstance1**](DiskGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewDiskGetInstance

`func NewDiskGetInstance() *DiskGetInstance`

NewDiskGetInstance instantiates a new DiskGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskGetInstanceWithDefaults

`func NewDiskGetInstanceWithDefaults() *DiskGetInstance`

NewDiskGetInstanceWithDefaults instantiates a new DiskGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiskGetInstance) GetId() DiskGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskGetInstance) GetIdOk() (*DiskGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskGetInstance) SetId(v DiskGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *DiskGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *DiskGetInstance) GetQueryOptions() DiskGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *DiskGetInstance) GetQueryOptionsOk() (*DiskGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *DiskGetInstance) SetQueryOptions(v DiskGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *DiskGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


