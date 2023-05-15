# Pool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Guid** | **string** |  | 
**Encrypt** | **int32** |  | 
**Encryptkey** | **string** |  | 
**EncryptkeyPath** | **NullableString** |  | 
**IsDecrypted** | **bool** |  | 
**Status** | **string** |  | 
**Path** | **string** |  | 
**Scan** | **map[string]interface{}** |  | 
**IsUpgraded** | Pointer to **bool** |  | [optional] 
**Healthy** | **bool** |  | 
**Warning** | **bool** |  | 
**StatusDetail** | **NullableString** |  | 
**Size** | **NullableInt64** |  | 
**Allocated** | **NullableInt64** |  | 
**Free** | **NullableInt64** |  | 
**Freeing** | **NullableInt64** |  | 
**Fragmentation** | **NullableString** |  | 
**Autotrim** | **map[string]interface{}** |  | 
**Topology** | [**PoolTopology**](PoolTopology.md) |  | 

## Methods

### NewPool

`func NewPool(id int32, name string, guid string, encrypt int32, encryptkey string, encryptkeyPath NullableString, isDecrypted bool, status string, path string, scan map[string]interface{}, healthy bool, warning bool, statusDetail NullableString, size NullableInt64, allocated NullableInt64, free NullableInt64, freeing NullableInt64, fragmentation NullableString, autotrim map[string]interface{}, topology PoolTopology, ) *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pool) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Pool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pool) SetName(v string)`

SetName sets Name field to given value.


### GetGuid

`func (o *Pool) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Pool) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Pool) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetEncrypt

`func (o *Pool) GetEncrypt() int32`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *Pool) GetEncryptOk() (*int32, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *Pool) SetEncrypt(v int32)`

SetEncrypt sets Encrypt field to given value.


### GetEncryptkey

`func (o *Pool) GetEncryptkey() string`

GetEncryptkey returns the Encryptkey field if non-nil, zero value otherwise.

### GetEncryptkeyOk

`func (o *Pool) GetEncryptkeyOk() (*string, bool)`

GetEncryptkeyOk returns a tuple with the Encryptkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptkey

`func (o *Pool) SetEncryptkey(v string)`

SetEncryptkey sets Encryptkey field to given value.


### GetEncryptkeyPath

`func (o *Pool) GetEncryptkeyPath() string`

GetEncryptkeyPath returns the EncryptkeyPath field if non-nil, zero value otherwise.

### GetEncryptkeyPathOk

`func (o *Pool) GetEncryptkeyPathOk() (*string, bool)`

GetEncryptkeyPathOk returns a tuple with the EncryptkeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptkeyPath

`func (o *Pool) SetEncryptkeyPath(v string)`

SetEncryptkeyPath sets EncryptkeyPath field to given value.


### SetEncryptkeyPathNil

`func (o *Pool) SetEncryptkeyPathNil(b bool)`

 SetEncryptkeyPathNil sets the value for EncryptkeyPath to be an explicit nil

### UnsetEncryptkeyPath
`func (o *Pool) UnsetEncryptkeyPath()`

UnsetEncryptkeyPath ensures that no value is present for EncryptkeyPath, not even an explicit nil
### GetIsDecrypted

`func (o *Pool) GetIsDecrypted() bool`

GetIsDecrypted returns the IsDecrypted field if non-nil, zero value otherwise.

### GetIsDecryptedOk

`func (o *Pool) GetIsDecryptedOk() (*bool, bool)`

GetIsDecryptedOk returns a tuple with the IsDecrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDecrypted

`func (o *Pool) SetIsDecrypted(v bool)`

SetIsDecrypted sets IsDecrypted field to given value.


### GetStatus

`func (o *Pool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pool) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPath

`func (o *Pool) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Pool) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Pool) SetPath(v string)`

SetPath sets Path field to given value.


### GetScan

`func (o *Pool) GetScan() map[string]interface{}`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *Pool) GetScanOk() (*map[string]interface{}, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *Pool) SetScan(v map[string]interface{})`

SetScan sets Scan field to given value.


### GetIsUpgraded

`func (o *Pool) GetIsUpgraded() bool`

GetIsUpgraded returns the IsUpgraded field if non-nil, zero value otherwise.

### GetIsUpgradedOk

`func (o *Pool) GetIsUpgradedOk() (*bool, bool)`

GetIsUpgradedOk returns a tuple with the IsUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgraded

`func (o *Pool) SetIsUpgraded(v bool)`

SetIsUpgraded sets IsUpgraded field to given value.

### HasIsUpgraded

`func (o *Pool) HasIsUpgraded() bool`

HasIsUpgraded returns a boolean if a field has been set.

### GetHealthy

`func (o *Pool) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *Pool) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *Pool) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.


### GetWarning

`func (o *Pool) GetWarning() bool`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *Pool) GetWarningOk() (*bool, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *Pool) SetWarning(v bool)`

SetWarning sets Warning field to given value.


### GetStatusDetail

`func (o *Pool) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *Pool) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *Pool) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.


### SetStatusDetailNil

`func (o *Pool) SetStatusDetailNil(b bool)`

 SetStatusDetailNil sets the value for StatusDetail to be an explicit nil

### UnsetStatusDetail
`func (o *Pool) UnsetStatusDetail()`

UnsetStatusDetail ensures that no value is present for StatusDetail, not even an explicit nil
### GetSize

`func (o *Pool) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Pool) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Pool) SetSize(v int64)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *Pool) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Pool) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetAllocated

`func (o *Pool) GetAllocated() int64`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *Pool) GetAllocatedOk() (*int64, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *Pool) SetAllocated(v int64)`

SetAllocated sets Allocated field to given value.


### SetAllocatedNil

`func (o *Pool) SetAllocatedNil(b bool)`

 SetAllocatedNil sets the value for Allocated to be an explicit nil

### UnsetAllocated
`func (o *Pool) UnsetAllocated()`

UnsetAllocated ensures that no value is present for Allocated, not even an explicit nil
### GetFree

`func (o *Pool) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *Pool) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *Pool) SetFree(v int64)`

SetFree sets Free field to given value.


### SetFreeNil

`func (o *Pool) SetFreeNil(b bool)`

 SetFreeNil sets the value for Free to be an explicit nil

### UnsetFree
`func (o *Pool) UnsetFree()`

UnsetFree ensures that no value is present for Free, not even an explicit nil
### GetFreeing

`func (o *Pool) GetFreeing() int64`

GetFreeing returns the Freeing field if non-nil, zero value otherwise.

### GetFreeingOk

`func (o *Pool) GetFreeingOk() (*int64, bool)`

GetFreeingOk returns a tuple with the Freeing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeing

`func (o *Pool) SetFreeing(v int64)`

SetFreeing sets Freeing field to given value.


### SetFreeingNil

`func (o *Pool) SetFreeingNil(b bool)`

 SetFreeingNil sets the value for Freeing to be an explicit nil

### UnsetFreeing
`func (o *Pool) UnsetFreeing()`

UnsetFreeing ensures that no value is present for Freeing, not even an explicit nil
### GetFragmentation

`func (o *Pool) GetFragmentation() string`

GetFragmentation returns the Fragmentation field if non-nil, zero value otherwise.

### GetFragmentationOk

`func (o *Pool) GetFragmentationOk() (*string, bool)`

GetFragmentationOk returns a tuple with the Fragmentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragmentation

`func (o *Pool) SetFragmentation(v string)`

SetFragmentation sets Fragmentation field to given value.


### SetFragmentationNil

`func (o *Pool) SetFragmentationNil(b bool)`

 SetFragmentationNil sets the value for Fragmentation to be an explicit nil

### UnsetFragmentation
`func (o *Pool) UnsetFragmentation()`

UnsetFragmentation ensures that no value is present for Fragmentation, not even an explicit nil
### GetAutotrim

`func (o *Pool) GetAutotrim() map[string]interface{}`

GetAutotrim returns the Autotrim field if non-nil, zero value otherwise.

### GetAutotrimOk

`func (o *Pool) GetAutotrimOk() (*map[string]interface{}, bool)`

GetAutotrimOk returns a tuple with the Autotrim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotrim

`func (o *Pool) SetAutotrim(v map[string]interface{})`

SetAutotrim sets Autotrim field to given value.


### GetTopology

`func (o *Pool) GetTopology() PoolTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *Pool) GetTopologyOk() (*PoolTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *Pool) SetTopology(v PoolTopology)`

SetTopology sets Topology field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


