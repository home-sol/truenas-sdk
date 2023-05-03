# IpmiGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**IpmiGetInstance0**](IpmiGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**IpmiGetInstance1**](IpmiGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewIpmiGetInstance

`func NewIpmiGetInstance() *IpmiGetInstance`

NewIpmiGetInstance instantiates a new IpmiGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiGetInstanceWithDefaults

`func NewIpmiGetInstanceWithDefaults() *IpmiGetInstance`

NewIpmiGetInstanceWithDefaults instantiates a new IpmiGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpmiGetInstance) GetId() IpmiGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpmiGetInstance) GetIdOk() (*IpmiGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpmiGetInstance) SetId(v IpmiGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *IpmiGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *IpmiGetInstance) GetQueryOptions() IpmiGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *IpmiGetInstance) GetQueryOptionsOk() (*IpmiGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *IpmiGetInstance) SetQueryOptions(v IpmiGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *IpmiGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


