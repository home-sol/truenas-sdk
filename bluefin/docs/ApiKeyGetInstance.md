# ApiKeyGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**ApiKeyGetInstance0**](ApiKeyGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**ApiKeyGetInstance1**](ApiKeyGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewApiKeyGetInstance

`func NewApiKeyGetInstance() *ApiKeyGetInstance`

NewApiKeyGetInstance instantiates a new ApiKeyGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyGetInstanceWithDefaults

`func NewApiKeyGetInstanceWithDefaults() *ApiKeyGetInstance`

NewApiKeyGetInstanceWithDefaults instantiates a new ApiKeyGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiKeyGetInstance) GetId() ApiKeyGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyGetInstance) GetIdOk() (*ApiKeyGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyGetInstance) SetId(v ApiKeyGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKeyGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *ApiKeyGetInstance) GetQueryOptions() ApiKeyGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *ApiKeyGetInstance) GetQueryOptionsOk() (*ApiKeyGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *ApiKeyGetInstance) SetQueryOptions(v ApiKeyGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *ApiKeyGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


