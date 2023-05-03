# UserGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**UserGetInstance0**](UserGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**UserGetInstance1**](UserGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewUserGetInstance

`func NewUserGetInstance() *UserGetInstance`

NewUserGetInstance instantiates a new UserGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGetInstanceWithDefaults

`func NewUserGetInstanceWithDefaults() *UserGetInstance`

NewUserGetInstanceWithDefaults instantiates a new UserGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGetInstance) GetId() UserGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGetInstance) GetIdOk() (*UserGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGetInstance) SetId(v UserGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *UserGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *UserGetInstance) GetQueryOptions() UserGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *UserGetInstance) GetQueryOptionsOk() (*UserGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *UserGetInstance) SetQueryOptions(v UserGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *UserGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


