# IscsiInitiatorGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**IscsiInitiatorGetInstance0**](IscsiInitiatorGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**IscsiInitiatorGetInstance1**](IscsiInitiatorGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewIscsiInitiatorGetInstance

`func NewIscsiInitiatorGetInstance() *IscsiInitiatorGetInstance`

NewIscsiInitiatorGetInstance instantiates a new IscsiInitiatorGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiInitiatorGetInstanceWithDefaults

`func NewIscsiInitiatorGetInstanceWithDefaults() *IscsiInitiatorGetInstance`

NewIscsiInitiatorGetInstanceWithDefaults instantiates a new IscsiInitiatorGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IscsiInitiatorGetInstance) GetId() IscsiInitiatorGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IscsiInitiatorGetInstance) GetIdOk() (*IscsiInitiatorGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IscsiInitiatorGetInstance) SetId(v IscsiInitiatorGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *IscsiInitiatorGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *IscsiInitiatorGetInstance) GetQueryOptions() IscsiInitiatorGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *IscsiInitiatorGetInstance) GetQueryOptionsOk() (*IscsiInitiatorGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *IscsiInitiatorGetInstance) SetQueryOptions(v IscsiInitiatorGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *IscsiInitiatorGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


