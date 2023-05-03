# PrivilegeCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | &#x60;name&#x60; is a name for privilege (must be unique). | [optional] 
**LocalGroups** | Pointer to **[]int32** | &#x60;local_groups&#x60; is a list of local user account group GIDs that gain this privilege. | [optional] [default to []]
**DsGroups** | Pointer to **[]int32** | &#x60;ds_groups&#x60; is list of Directory Service group GIDs that will gain this privilege. | [optional] [default to []]
**Allowlist** | Pointer to [**[]AllowlistItem**](AllowlistItem.md) | &#x60;allowlist&#x60; is a list of API endpoints allowed for this privilege. | [optional] [default to []]
**WebShell** | Pointer to **bool** |  | [optional] 

## Methods

### NewPrivilegeCreate0

`func NewPrivilegeCreate0() *PrivilegeCreate0`

NewPrivilegeCreate0 instantiates a new PrivilegeCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeCreate0WithDefaults

`func NewPrivilegeCreate0WithDefaults() *PrivilegeCreate0`

NewPrivilegeCreate0WithDefaults instantiates a new PrivilegeCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivilegeCreate0) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivilegeCreate0) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivilegeCreate0) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PrivilegeCreate0) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrivilegeCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivilegeCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivilegeCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivilegeCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocalGroups

`func (o *PrivilegeCreate0) GetLocalGroups() []int32`

GetLocalGroups returns the LocalGroups field if non-nil, zero value otherwise.

### GetLocalGroupsOk

`func (o *PrivilegeCreate0) GetLocalGroupsOk() (*[]int32, bool)`

GetLocalGroupsOk returns a tuple with the LocalGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroups

`func (o *PrivilegeCreate0) SetLocalGroups(v []int32)`

SetLocalGroups sets LocalGroups field to given value.

### HasLocalGroups

`func (o *PrivilegeCreate0) HasLocalGroups() bool`

HasLocalGroups returns a boolean if a field has been set.

### GetDsGroups

`func (o *PrivilegeCreate0) GetDsGroups() []int32`

GetDsGroups returns the DsGroups field if non-nil, zero value otherwise.

### GetDsGroupsOk

`func (o *PrivilegeCreate0) GetDsGroupsOk() (*[]int32, bool)`

GetDsGroupsOk returns a tuple with the DsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsGroups

`func (o *PrivilegeCreate0) SetDsGroups(v []int32)`

SetDsGroups sets DsGroups field to given value.

### HasDsGroups

`func (o *PrivilegeCreate0) HasDsGroups() bool`

HasDsGroups returns a boolean if a field has been set.

### GetAllowlist

`func (o *PrivilegeCreate0) GetAllowlist() []AllowlistItem`

GetAllowlist returns the Allowlist field if non-nil, zero value otherwise.

### GetAllowlistOk

`func (o *PrivilegeCreate0) GetAllowlistOk() (*[]AllowlistItem, bool)`

GetAllowlistOk returns a tuple with the Allowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlist

`func (o *PrivilegeCreate0) SetAllowlist(v []AllowlistItem)`

SetAllowlist sets Allowlist field to given value.

### HasAllowlist

`func (o *PrivilegeCreate0) HasAllowlist() bool`

HasAllowlist returns a boolean if a field has been set.

### GetWebShell

`func (o *PrivilegeCreate0) GetWebShell() bool`

GetWebShell returns the WebShell field if non-nil, zero value otherwise.

### GetWebShellOk

`func (o *PrivilegeCreate0) GetWebShellOk() (*bool, bool)`

GetWebShellOk returns a tuple with the WebShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebShell

`func (o *PrivilegeCreate0) SetWebShell(v bool)`

SetWebShell sets WebShell field to given value.

### HasWebShell

`func (o *PrivilegeCreate0) HasWebShell() bool`

HasWebShell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


