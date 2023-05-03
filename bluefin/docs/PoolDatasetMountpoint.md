# PoolDatasetMountpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** |  | [optional] 
**Raise** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPoolDatasetMountpoint

`func NewPoolDatasetMountpoint() *PoolDatasetMountpoint`

NewPoolDatasetMountpoint instantiates a new PoolDatasetMountpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetMountpointWithDefaults

`func NewPoolDatasetMountpointWithDefaults() *PoolDatasetMountpoint`

NewPoolDatasetMountpointWithDefaults instantiates a new PoolDatasetMountpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *PoolDatasetMountpoint) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *PoolDatasetMountpoint) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *PoolDatasetMountpoint) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *PoolDatasetMountpoint) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetRaise

`func (o *PoolDatasetMountpoint) GetRaise() bool`

GetRaise returns the Raise field if non-nil, zero value otherwise.

### GetRaiseOk

`func (o *PoolDatasetMountpoint) GetRaiseOk() (*bool, bool)`

GetRaiseOk returns a tuple with the Raise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaise

`func (o *PoolDatasetMountpoint) SetRaise(v bool)`

SetRaise sets Raise field to given value.

### HasRaise

`func (o *PoolDatasetMountpoint) HasRaise() bool`

HasRaise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


