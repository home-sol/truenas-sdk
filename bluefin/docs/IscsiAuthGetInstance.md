# IscsiAuthGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**IscsiAuthGetInstance0**](IscsiAuthGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**IscsiAuthGetInstance1**](IscsiAuthGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewIscsiAuthGetInstance

`func NewIscsiAuthGetInstance() *IscsiAuthGetInstance`

NewIscsiAuthGetInstance instantiates a new IscsiAuthGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiAuthGetInstanceWithDefaults

`func NewIscsiAuthGetInstanceWithDefaults() *IscsiAuthGetInstance`

NewIscsiAuthGetInstanceWithDefaults instantiates a new IscsiAuthGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IscsiAuthGetInstance) GetId() IscsiAuthGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IscsiAuthGetInstance) GetIdOk() (*IscsiAuthGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IscsiAuthGetInstance) SetId(v IscsiAuthGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *IscsiAuthGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *IscsiAuthGetInstance) GetQueryOptions() IscsiAuthGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *IscsiAuthGetInstance) GetQueryOptionsOk() (*IscsiAuthGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *IscsiAuthGetInstance) SetQueryOptions(v IscsiAuthGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *IscsiAuthGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


