# VDev

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Path** | **NullableString** |  | 
**Status** | **string** |  | 
**Children** | Pointer to [**[]Device**](Device.md) |  | [optional] 

## Methods

### NewVDev

`func NewVDev(guid string, name string, type_ string, path NullableString, status string, ) *VDev`

NewVDev instantiates a new VDev object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDevWithDefaults

`func NewVDevWithDefaults() *VDev`

NewVDevWithDefaults instantiates a new VDev object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *VDev) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *VDev) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *VDev) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetName

`func (o *VDev) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VDev) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VDev) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VDev) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VDev) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VDev) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *VDev) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VDev) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VDev) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *VDev) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *VDev) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetStatus

`func (o *VDev) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VDev) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VDev) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetChildren

`func (o *VDev) GetChildren() []Device`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *VDev) GetChildrenOk() (*[]Device, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *VDev) SetChildren(v []Device)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *VDev) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


