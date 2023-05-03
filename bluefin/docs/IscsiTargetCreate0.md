# IscsiTargetCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "ISCSI"]
**Groups** | Pointer to [**[]Group**](Group.md) | &#x60;groups&#x60; is a list of group dictionaries which provide information related to using a &#x60;portal&#x60;, &#x60;initiator&#x60;, &#x60;authmethod&#x60; and &#x60;auth&#x60; with this target. &#x60;auth&#x60; represents a valid iSCSI Authorized Access and defaults to null. | [optional] [default to []]
**AuthNetworks** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewIscsiTargetCreate0

`func NewIscsiTargetCreate0() *IscsiTargetCreate0`

NewIscsiTargetCreate0 instantiates a new IscsiTargetCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiTargetCreate0WithDefaults

`func NewIscsiTargetCreate0WithDefaults() *IscsiTargetCreate0`

NewIscsiTargetCreate0WithDefaults instantiates a new IscsiTargetCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IscsiTargetCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IscsiTargetCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IscsiTargetCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IscsiTargetCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlias

`func (o *IscsiTargetCreate0) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *IscsiTargetCreate0) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *IscsiTargetCreate0) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *IscsiTargetCreate0) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *IscsiTargetCreate0) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *IscsiTargetCreate0) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetMode

`func (o *IscsiTargetCreate0) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IscsiTargetCreate0) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IscsiTargetCreate0) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *IscsiTargetCreate0) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetGroups

`func (o *IscsiTargetCreate0) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IscsiTargetCreate0) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IscsiTargetCreate0) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IscsiTargetCreate0) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAuthNetworks

`func (o *IscsiTargetCreate0) GetAuthNetworks() []string`

GetAuthNetworks returns the AuthNetworks field if non-nil, zero value otherwise.

### GetAuthNetworksOk

`func (o *IscsiTargetCreate0) GetAuthNetworksOk() (*[]string, bool)`

GetAuthNetworksOk returns a tuple with the AuthNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNetworks

`func (o *IscsiTargetCreate0) SetAuthNetworks(v []string)`

SetAuthNetworks sets AuthNetworks field to given value.

### HasAuthNetworks

`func (o *IscsiTargetCreate0) HasAuthNetworks() bool`

HasAuthNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


