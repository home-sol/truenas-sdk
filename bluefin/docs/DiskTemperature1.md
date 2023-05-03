# DiskTemperature1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **NullableInt32** |  | [optional] 
**Powermode** | Pointer to **string** |  | [optional] [default to "NEVER"]

## Methods

### NewDiskTemperature1

`func NewDiskTemperature1() *DiskTemperature1`

NewDiskTemperature1 instantiates a new DiskTemperature1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskTemperature1WithDefaults

`func NewDiskTemperature1WithDefaults() *DiskTemperature1`

NewDiskTemperature1WithDefaults instantiates a new DiskTemperature1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *DiskTemperature1) GetCache() int32`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *DiskTemperature1) GetCacheOk() (*int32, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *DiskTemperature1) SetCache(v int32)`

SetCache sets Cache field to given value.

### HasCache

`func (o *DiskTemperature1) HasCache() bool`

HasCache returns a boolean if a field has been set.

### SetCacheNil

`func (o *DiskTemperature1) SetCacheNil(b bool)`

 SetCacheNil sets the value for Cache to be an explicit nil

### UnsetCache
`func (o *DiskTemperature1) UnsetCache()`

UnsetCache ensures that no value is present for Cache, not even an explicit nil
### GetPowermode

`func (o *DiskTemperature1) GetPowermode() string`

GetPowermode returns the Powermode field if non-nil, zero value otherwise.

### GetPowermodeOk

`func (o *DiskTemperature1) GetPowermodeOk() (*string, bool)`

GetPowermodeOk returns a tuple with the Powermode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowermode

`func (o *DiskTemperature1) SetPowermode(v string)`

SetPowermode sets Powermode field to given value.

### HasPowermode

`func (o *DiskTemperature1) HasPowermode() bool`

HasPowermode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

