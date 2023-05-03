# DiskTemperatures1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to **NullableInt32** |  | [optional] [default to 290]
**OnlyCached** | Pointer to **bool** |  | [optional] [default to false]
**Powermode** | Pointer to **string** |  | [optional] [default to "NEVER"]

## Methods

### NewDiskTemperatures1

`func NewDiskTemperatures1() *DiskTemperatures1`

NewDiskTemperatures1 instantiates a new DiskTemperatures1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskTemperatures1WithDefaults

`func NewDiskTemperatures1WithDefaults() *DiskTemperatures1`

NewDiskTemperatures1WithDefaults instantiates a new DiskTemperatures1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *DiskTemperatures1) GetCache() int32`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *DiskTemperatures1) GetCacheOk() (*int32, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *DiskTemperatures1) SetCache(v int32)`

SetCache sets Cache field to given value.

### HasCache

`func (o *DiskTemperatures1) HasCache() bool`

HasCache returns a boolean if a field has been set.

### SetCacheNil

`func (o *DiskTemperatures1) SetCacheNil(b bool)`

 SetCacheNil sets the value for Cache to be an explicit nil

### UnsetCache
`func (o *DiskTemperatures1) UnsetCache()`

UnsetCache ensures that no value is present for Cache, not even an explicit nil
### GetOnlyCached

`func (o *DiskTemperatures1) GetOnlyCached() bool`

GetOnlyCached returns the OnlyCached field if non-nil, zero value otherwise.

### GetOnlyCachedOk

`func (o *DiskTemperatures1) GetOnlyCachedOk() (*bool, bool)`

GetOnlyCachedOk returns a tuple with the OnlyCached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyCached

`func (o *DiskTemperatures1) SetOnlyCached(v bool)`

SetOnlyCached sets OnlyCached field to given value.

### HasOnlyCached

`func (o *DiskTemperatures1) HasOnlyCached() bool`

HasOnlyCached returns a boolean if a field has been set.

### GetPowermode

`func (o *DiskTemperatures1) GetPowermode() string`

GetPowermode returns the Powermode field if non-nil, zero value otherwise.

### GetPowermodeOk

`func (o *DiskTemperatures1) GetPowermodeOk() (*string, bool)`

GetPowermodeOk returns a tuple with the Powermode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowermode

`func (o *DiskTemperatures1) SetPowermode(v string)`

SetPowermode sets Powermode field to given value.

### HasPowermode

`func (o *DiskTemperatures1) HasPowermode() bool`

HasPowermode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


