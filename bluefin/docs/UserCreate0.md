# UserCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | If &#x60;uid&#x60; is not provided it is automatically filled with the next one available. | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **int32** | &#x60;group&#x60; is required if &#x60;group_create&#x60; is false. | [optional] 
**GroupCreate** | Pointer to **bool** | &#x60;group&#x60; is required if &#x60;group_create&#x60; is false. | [optional] [default to false]
**Home** | Pointer to **string** |  | [optional] [default to "/nonexistent"]
**HomeMode** | Pointer to **string** |  | [optional] [default to "700"]
**Shell** | Pointer to **string** | Available choices for &#x60;shell&#x60; can be retrieved with &#x60;user.shell_choices&#x60;. | [optional] [default to "/usr/bin/zsh"]
**FullName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **string** | &#x60;password&#x60; is required if &#x60;password_disabled&#x60; is false. | [optional] 
**PasswordDisabled** | Pointer to **bool** | &#x60;password&#x60; is required if &#x60;password_disabled&#x60; is false. | [optional] [default to false]
**Locked** | Pointer to **bool** |  | [optional] [default to false]
**Smb** | Pointer to **bool** |  | [optional] [default to true]
**SudoCommands** | Pointer to **[]string** |  | [optional] [default to []]
**SudoCommandsNopasswd** | Pointer to **[]string** |  | [optional] [default to []]
**Sshpubkey** | Pointer to **NullableString** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] [default to []]
**Attributes** | Pointer to **map[string]interface{}** | &#x60;attributes&#x60; is a general-purpose object for storing arbitrary user information. | [optional] [default to {}]

## Methods

### NewUserCreate0

`func NewUserCreate0() *UserCreate0`

NewUserCreate0 instantiates a new UserCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreate0WithDefaults

`func NewUserCreate0WithDefaults() *UserCreate0`

NewUserCreate0WithDefaults instantiates a new UserCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *UserCreate0) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserCreate0) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserCreate0) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserCreate0) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *UserCreate0) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCreate0) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCreate0) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserCreate0) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetGroup

`func (o *UserCreate0) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UserCreate0) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UserCreate0) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UserCreate0) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupCreate

`func (o *UserCreate0) GetGroupCreate() bool`

GetGroupCreate returns the GroupCreate field if non-nil, zero value otherwise.

### GetGroupCreateOk

`func (o *UserCreate0) GetGroupCreateOk() (*bool, bool)`

GetGroupCreateOk returns a tuple with the GroupCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCreate

`func (o *UserCreate0) SetGroupCreate(v bool)`

SetGroupCreate sets GroupCreate field to given value.

### HasGroupCreate

`func (o *UserCreate0) HasGroupCreate() bool`

HasGroupCreate returns a boolean if a field has been set.

### GetHome

`func (o *UserCreate0) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *UserCreate0) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *UserCreate0) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *UserCreate0) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetHomeMode

`func (o *UserCreate0) GetHomeMode() string`

GetHomeMode returns the HomeMode field if non-nil, zero value otherwise.

### GetHomeModeOk

`func (o *UserCreate0) GetHomeModeOk() (*string, bool)`

GetHomeModeOk returns a tuple with the HomeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMode

`func (o *UserCreate0) SetHomeMode(v string)`

SetHomeMode sets HomeMode field to given value.

### HasHomeMode

`func (o *UserCreate0) HasHomeMode() bool`

HasHomeMode returns a boolean if a field has been set.

### GetShell

`func (o *UserCreate0) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *UserCreate0) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *UserCreate0) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *UserCreate0) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetFullName

`func (o *UserCreate0) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserCreate0) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserCreate0) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserCreate0) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetEmail

`func (o *UserCreate0) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreate0) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreate0) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreate0) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserCreate0) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserCreate0) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *UserCreate0) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreate0) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreate0) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreate0) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordDisabled

`func (o *UserCreate0) GetPasswordDisabled() bool`

GetPasswordDisabled returns the PasswordDisabled field if non-nil, zero value otherwise.

### GetPasswordDisabledOk

`func (o *UserCreate0) GetPasswordDisabledOk() (*bool, bool)`

GetPasswordDisabledOk returns a tuple with the PasswordDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordDisabled

`func (o *UserCreate0) SetPasswordDisabled(v bool)`

SetPasswordDisabled sets PasswordDisabled field to given value.

### HasPasswordDisabled

`func (o *UserCreate0) HasPasswordDisabled() bool`

HasPasswordDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *UserCreate0) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UserCreate0) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UserCreate0) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UserCreate0) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetSmb

`func (o *UserCreate0) GetSmb() bool`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *UserCreate0) GetSmbOk() (*bool, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *UserCreate0) SetSmb(v bool)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *UserCreate0) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetSudoCommands

`func (o *UserCreate0) GetSudoCommands() []string`

GetSudoCommands returns the SudoCommands field if non-nil, zero value otherwise.

### GetSudoCommandsOk

`func (o *UserCreate0) GetSudoCommandsOk() (*[]string, bool)`

GetSudoCommandsOk returns a tuple with the SudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommands

`func (o *UserCreate0) SetSudoCommands(v []string)`

SetSudoCommands sets SudoCommands field to given value.

### HasSudoCommands

`func (o *UserCreate0) HasSudoCommands() bool`

HasSudoCommands returns a boolean if a field has been set.

### GetSudoCommandsNopasswd

`func (o *UserCreate0) GetSudoCommandsNopasswd() []string`

GetSudoCommandsNopasswd returns the SudoCommandsNopasswd field if non-nil, zero value otherwise.

### GetSudoCommandsNopasswdOk

`func (o *UserCreate0) GetSudoCommandsNopasswdOk() (*[]string, bool)`

GetSudoCommandsNopasswdOk returns a tuple with the SudoCommandsNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommandsNopasswd

`func (o *UserCreate0) SetSudoCommandsNopasswd(v []string)`

SetSudoCommandsNopasswd sets SudoCommandsNopasswd field to given value.

### HasSudoCommandsNopasswd

`func (o *UserCreate0) HasSudoCommandsNopasswd() bool`

HasSudoCommandsNopasswd returns a boolean if a field has been set.

### GetSshpubkey

`func (o *UserCreate0) GetSshpubkey() string`

GetSshpubkey returns the Sshpubkey field if non-nil, zero value otherwise.

### GetSshpubkeyOk

`func (o *UserCreate0) GetSshpubkeyOk() (*string, bool)`

GetSshpubkeyOk returns a tuple with the Sshpubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshpubkey

`func (o *UserCreate0) SetSshpubkey(v string)`

SetSshpubkey sets Sshpubkey field to given value.

### HasSshpubkey

`func (o *UserCreate0) HasSshpubkey() bool`

HasSshpubkey returns a boolean if a field has been set.

### SetSshpubkeyNil

`func (o *UserCreate0) SetSshpubkeyNil(b bool)`

 SetSshpubkeyNil sets the value for Sshpubkey to be an explicit nil

### UnsetSshpubkey
`func (o *UserCreate0) UnsetSshpubkey()`

UnsetSshpubkey ensures that no value is present for Sshpubkey, not even an explicit nil
### GetGroups

`func (o *UserCreate0) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserCreate0) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserCreate0) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserCreate0) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAttributes

`func (o *UserCreate0) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserCreate0) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserCreate0) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserCreate0) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


