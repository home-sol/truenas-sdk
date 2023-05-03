# ZfsSnapshotGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**ZfsSnapshotGetInstance0**](ZfsSnapshotGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**ZfsSnapshotGetInstance1**](ZfsSnapshotGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewZfsSnapshotGetInstance

`func NewZfsSnapshotGetInstance() *ZfsSnapshotGetInstance`

NewZfsSnapshotGetInstance instantiates a new ZfsSnapshotGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZfsSnapshotGetInstanceWithDefaults

`func NewZfsSnapshotGetInstanceWithDefaults() *ZfsSnapshotGetInstance`

NewZfsSnapshotGetInstanceWithDefaults instantiates a new ZfsSnapshotGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZfsSnapshotGetInstance) GetId() ZfsSnapshotGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZfsSnapshotGetInstance) GetIdOk() (*ZfsSnapshotGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZfsSnapshotGetInstance) SetId(v ZfsSnapshotGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *ZfsSnapshotGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *ZfsSnapshotGetInstance) GetQueryOptions() ZfsSnapshotGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *ZfsSnapshotGetInstance) GetQueryOptionsOk() (*ZfsSnapshotGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *ZfsSnapshotGetInstance) SetQueryOptions(v ZfsSnapshotGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *ZfsSnapshotGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


