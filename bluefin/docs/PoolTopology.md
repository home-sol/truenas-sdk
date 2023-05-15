# PoolTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VDev**](VDev.md) |  | 
**Log** | [**[]VDev**](VDev.md) |  | 
**Cache** | [**[]VDev**](VDev.md) |  | 
**Spare** | [**[]VDev**](VDev.md) |  | 
**Special** | [**[]VDev**](VDev.md) |  | 
**Dedup** | [**[]VDev**](VDev.md) |  | 

## Methods

### NewPoolTopology

`func NewPoolTopology(data []VDev, log []VDev, cache []VDev, spare []VDev, special []VDev, dedup []VDev, ) *PoolTopology`

NewPoolTopology instantiates a new PoolTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolTopologyWithDefaults

`func NewPoolTopologyWithDefaults() *PoolTopology`

NewPoolTopologyWithDefaults instantiates a new PoolTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PoolTopology) GetData() []VDev`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PoolTopology) GetDataOk() (*[]VDev, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PoolTopology) SetData(v []VDev)`

SetData sets Data field to given value.


### GetLog

`func (o *PoolTopology) GetLog() []VDev`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *PoolTopology) GetLogOk() (*[]VDev, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *PoolTopology) SetLog(v []VDev)`

SetLog sets Log field to given value.


### GetCache

`func (o *PoolTopology) GetCache() []VDev`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *PoolTopology) GetCacheOk() (*[]VDev, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *PoolTopology) SetCache(v []VDev)`

SetCache sets Cache field to given value.


### GetSpare

`func (o *PoolTopology) GetSpare() []VDev`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *PoolTopology) GetSpareOk() (*[]VDev, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *PoolTopology) SetSpare(v []VDev)`

SetSpare sets Spare field to given value.


### GetSpecial

`func (o *PoolTopology) GetSpecial() []VDev`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *PoolTopology) GetSpecialOk() (*[]VDev, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *PoolTopology) SetSpecial(v []VDev)`

SetSpecial sets Special field to given value.


### GetDedup

`func (o *PoolTopology) GetDedup() []VDev`

GetDedup returns the Dedup field if non-nil, zero value otherwise.

### GetDedupOk

`func (o *PoolTopology) GetDedupOk() (*[]VDev, bool)`

GetDedupOk returns a tuple with the Dedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedup

`func (o *PoolTopology) SetDedup(v []VDev)`

SetDedup sets Dedup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


