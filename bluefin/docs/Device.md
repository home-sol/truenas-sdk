# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **string** |  | [optional] 
**Children** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnavailDisk** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *Device) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Device) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Device) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Device) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetGuid

`func (o *Device) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Device) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Device) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Device) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetStatus

`func (o *Device) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStats

`func (o *Device) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Device) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Device) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *Device) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDevice

`func (o *Device) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Device) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Device) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Device) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDisk

`func (o *Device) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Device) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Device) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Device) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetChildren

`func (o *Device) GetChildren() []map[string]interface{}`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Device) GetChildrenOk() (*[]map[string]interface{}, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Device) SetChildren(v []map[string]interface{})`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Device) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetUnavailDisk

`func (o *Device) GetUnavailDisk() map[string]interface{}`

GetUnavailDisk returns the UnavailDisk field if non-nil, zero value otherwise.

### GetUnavailDiskOk

`func (o *Device) GetUnavailDiskOk() (*map[string]interface{}, bool)`

GetUnavailDiskOk returns a tuple with the UnavailDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailDisk

`func (o *Device) SetUnavailDisk(v map[string]interface{})`

SetUnavailDisk sets UnavailDisk field to given value.

### HasUnavailDisk

`func (o *Device) HasUnavailDisk() bool`

HasUnavailDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


