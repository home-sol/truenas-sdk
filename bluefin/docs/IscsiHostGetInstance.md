# IscsiHostGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**IscsiHostGetInstance0**](IscsiHostGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**IscsiHostGetInstance1**](IscsiHostGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewIscsiHostGetInstance

`func NewIscsiHostGetInstance() *IscsiHostGetInstance`

NewIscsiHostGetInstance instantiates a new IscsiHostGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiHostGetInstanceWithDefaults

`func NewIscsiHostGetInstanceWithDefaults() *IscsiHostGetInstance`

NewIscsiHostGetInstanceWithDefaults instantiates a new IscsiHostGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IscsiHostGetInstance) GetId() IscsiHostGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IscsiHostGetInstance) GetIdOk() (*IscsiHostGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IscsiHostGetInstance) SetId(v IscsiHostGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *IscsiHostGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *IscsiHostGetInstance) GetQueryOptions() IscsiHostGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *IscsiHostGetInstance) GetQueryOptionsOk() (*IscsiHostGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *IscsiHostGetInstance) SetQueryOptions(v IscsiHostGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *IscsiHostGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


