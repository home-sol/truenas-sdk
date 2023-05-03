# DiskResize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | Pointer to [**[]DiskResize0Inner**](DiskResize0Inner.md) |  | [optional] [default to []]
**Sync** | Pointer to **bool** | &#x60;sync&#x60;: boolean, when true (default) will synchronize the new size of the disk(s)     with the database cache. | [optional] [default to true]
**RaiseError** | Pointer to **bool** | &#x60;raise_error&#x60;: boolean     when true, will raise a &#x60;CallError&#x60; if any failures occur     when false, will will log the errors if any failures occur | [optional] [default to false]

## Methods

### NewDiskResize

`func NewDiskResize() *DiskResize`

NewDiskResize instantiates a new DiskResize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskResizeWithDefaults

`func NewDiskResizeWithDefaults() *DiskResize`

NewDiskResizeWithDefaults instantiates a new DiskResize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisks

`func (o *DiskResize) GetDisks() []DiskResize0Inner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *DiskResize) GetDisksOk() (*[]DiskResize0Inner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *DiskResize) SetDisks(v []DiskResize0Inner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *DiskResize) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetSync

`func (o *DiskResize) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *DiskResize) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *DiskResize) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *DiskResize) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetRaiseError

`func (o *DiskResize) GetRaiseError() bool`

GetRaiseError returns the RaiseError field if non-nil, zero value otherwise.

### GetRaiseErrorOk

`func (o *DiskResize) GetRaiseErrorOk() (*bool, bool)`

GetRaiseErrorOk returns a tuple with the RaiseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseError

`func (o *DiskResize) SetRaiseError(v bool)`

SetRaiseError sets RaiseError field to given value.

### HasRaiseError

`func (o *DiskResize) HasRaiseError() bool`

HasRaiseError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


