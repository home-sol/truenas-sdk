# RsynctaskCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | See the comment in Rsyncmod about &#x60;path&#x60; length limits. | [optional] 
**User** | Pointer to **string** | In SSH mode, if &#x60;ssh_credentials&#x60; (a keychain credential of &#x60;SSH_CREDENTIALS&#x60; type) is specified then it is used to connect to the remote host. If it is not specified, then keys in &#x60;user&#x60;&#39;s .ssh directory are used. | [optional] 
**Mode** | Pointer to **string** | &#x60;mode&#x60; represents different operating mechanisms for Rsync i.e Rsync Module mode / Rsync SSH mode. &#x60;remotemodule&#x60; is the name of remote module, this attribute should be specified when &#x60;mode&#x60; is set to MODULE. | [optional] [default to "MODULE"]
**Remotehost** | Pointer to **NullableString** | &#x60;remotehost&#x60; is ip address or hostname of the remote system. If username differs on the remote host, \&quot;username@remote_host\&quot; format should be used. &#x60;remotehost&#x60; and &#x60;remoteport&#x60; are not used in this case. | [optional] 
**Remoteport** | Pointer to **NullableInt32** | &#x60;remotehost&#x60; and &#x60;remoteport&#x60; are not used in this case. | [optional] 
**Remotemodule** | Pointer to **NullableString** | &#x60;remotemodule&#x60; is the name of remote module, this attribute should be specified when &#x60;mode&#x60; is set to MODULE. | [optional] 
**SshCredentials** | Pointer to **NullableInt32** | In SSH mode, if &#x60;ssh_credentials&#x60; (a keychain credential of &#x60;SSH_CREDENTIALS&#x60; type) is specified then it is used to connect to the remote host. If it is not specified, then keys in &#x60;user&#x60;&#39;s .ssh directory are used. | [optional] 
**Remotepath** | Pointer to **string** | &#x60;remotepath&#x60; specifies the path on the remote system. | [optional] 
**ValidateRpath** | Pointer to **bool** | &#x60;validate_rpath&#x60; is a boolean which when sets validates the existence of the remote path. | [optional] [default to true]
**Direction** | Pointer to **string** | &#x60;direction&#x60; specifies if data should be PULLED or PUSHED from the remote system. | [optional] [default to "PUSH"]
**Desc** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] [default to {}]
**Recursive** | Pointer to **bool** |  | [optional] 
**Times** | Pointer to **bool** |  | [optional] 
**Compress** | Pointer to **bool** | &#x60;compress&#x60; when set reduces the size of the data which is to be transmitted. | [optional] 
**Archive** | Pointer to **bool** | &#x60;archive&#x60; when set makes rsync run recursively, preserving symlinks, permissions, modification times, group, and special files. | [optional] 
**Delete** | Pointer to **bool** | &#x60;delete&#x60; when set deletes files in the destination directory which do not exist in the source directory. | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**Preserveperm** | Pointer to **bool** | &#x60;preserveperm&#x60; when set preserves original file permissions. | [optional] 
**Preserveattr** | Pointer to **bool** |  | [optional] 
**Delayupdates** | Pointer to **bool** |  | [optional] 
**Extra** | Pointer to **[]string** |  | [optional] [default to []]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewRsynctaskCreate0

`func NewRsynctaskCreate0() *RsynctaskCreate0`

NewRsynctaskCreate0 instantiates a new RsynctaskCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRsynctaskCreate0WithDefaults

`func NewRsynctaskCreate0WithDefaults() *RsynctaskCreate0`

NewRsynctaskCreate0WithDefaults instantiates a new RsynctaskCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RsynctaskCreate0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RsynctaskCreate0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RsynctaskCreate0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RsynctaskCreate0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUser

`func (o *RsynctaskCreate0) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RsynctaskCreate0) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RsynctaskCreate0) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RsynctaskCreate0) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMode

`func (o *RsynctaskCreate0) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RsynctaskCreate0) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RsynctaskCreate0) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RsynctaskCreate0) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRemotehost

`func (o *RsynctaskCreate0) GetRemotehost() string`

GetRemotehost returns the Remotehost field if non-nil, zero value otherwise.

### GetRemotehostOk

`func (o *RsynctaskCreate0) GetRemotehostOk() (*string, bool)`

GetRemotehostOk returns a tuple with the Remotehost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotehost

`func (o *RsynctaskCreate0) SetRemotehost(v string)`

SetRemotehost sets Remotehost field to given value.

### HasRemotehost

`func (o *RsynctaskCreate0) HasRemotehost() bool`

HasRemotehost returns a boolean if a field has been set.

### SetRemotehostNil

`func (o *RsynctaskCreate0) SetRemotehostNil(b bool)`

 SetRemotehostNil sets the value for Remotehost to be an explicit nil

### UnsetRemotehost
`func (o *RsynctaskCreate0) UnsetRemotehost()`

UnsetRemotehost ensures that no value is present for Remotehost, not even an explicit nil
### GetRemoteport

`func (o *RsynctaskCreate0) GetRemoteport() int32`

GetRemoteport returns the Remoteport field if non-nil, zero value otherwise.

### GetRemoteportOk

`func (o *RsynctaskCreate0) GetRemoteportOk() (*int32, bool)`

GetRemoteportOk returns a tuple with the Remoteport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteport

`func (o *RsynctaskCreate0) SetRemoteport(v int32)`

SetRemoteport sets Remoteport field to given value.

### HasRemoteport

`func (o *RsynctaskCreate0) HasRemoteport() bool`

HasRemoteport returns a boolean if a field has been set.

### SetRemoteportNil

`func (o *RsynctaskCreate0) SetRemoteportNil(b bool)`

 SetRemoteportNil sets the value for Remoteport to be an explicit nil

### UnsetRemoteport
`func (o *RsynctaskCreate0) UnsetRemoteport()`

UnsetRemoteport ensures that no value is present for Remoteport, not even an explicit nil
### GetRemotemodule

`func (o *RsynctaskCreate0) GetRemotemodule() string`

GetRemotemodule returns the Remotemodule field if non-nil, zero value otherwise.

### GetRemotemoduleOk

`func (o *RsynctaskCreate0) GetRemotemoduleOk() (*string, bool)`

GetRemotemoduleOk returns a tuple with the Remotemodule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotemodule

`func (o *RsynctaskCreate0) SetRemotemodule(v string)`

SetRemotemodule sets Remotemodule field to given value.

### HasRemotemodule

`func (o *RsynctaskCreate0) HasRemotemodule() bool`

HasRemotemodule returns a boolean if a field has been set.

### SetRemotemoduleNil

`func (o *RsynctaskCreate0) SetRemotemoduleNil(b bool)`

 SetRemotemoduleNil sets the value for Remotemodule to be an explicit nil

### UnsetRemotemodule
`func (o *RsynctaskCreate0) UnsetRemotemodule()`

UnsetRemotemodule ensures that no value is present for Remotemodule, not even an explicit nil
### GetSshCredentials

`func (o *RsynctaskCreate0) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *RsynctaskCreate0) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *RsynctaskCreate0) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *RsynctaskCreate0) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *RsynctaskCreate0) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *RsynctaskCreate0) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil
### GetRemotepath

`func (o *RsynctaskCreate0) GetRemotepath() string`

GetRemotepath returns the Remotepath field if non-nil, zero value otherwise.

### GetRemotepathOk

`func (o *RsynctaskCreate0) GetRemotepathOk() (*string, bool)`

GetRemotepathOk returns a tuple with the Remotepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotepath

`func (o *RsynctaskCreate0) SetRemotepath(v string)`

SetRemotepath sets Remotepath field to given value.

### HasRemotepath

`func (o *RsynctaskCreate0) HasRemotepath() bool`

HasRemotepath returns a boolean if a field has been set.

### GetValidateRpath

`func (o *RsynctaskCreate0) GetValidateRpath() bool`

GetValidateRpath returns the ValidateRpath field if non-nil, zero value otherwise.

### GetValidateRpathOk

`func (o *RsynctaskCreate0) GetValidateRpathOk() (*bool, bool)`

GetValidateRpathOk returns a tuple with the ValidateRpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRpath

`func (o *RsynctaskCreate0) SetValidateRpath(v bool)`

SetValidateRpath sets ValidateRpath field to given value.

### HasValidateRpath

`func (o *RsynctaskCreate0) HasValidateRpath() bool`

HasValidateRpath returns a boolean if a field has been set.

### GetDirection

`func (o *RsynctaskCreate0) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RsynctaskCreate0) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RsynctaskCreate0) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RsynctaskCreate0) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDesc

`func (o *RsynctaskCreate0) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *RsynctaskCreate0) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *RsynctaskCreate0) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *RsynctaskCreate0) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetSchedule

`func (o *RsynctaskCreate0) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RsynctaskCreate0) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RsynctaskCreate0) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *RsynctaskCreate0) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRecursive

`func (o *RsynctaskCreate0) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *RsynctaskCreate0) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *RsynctaskCreate0) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *RsynctaskCreate0) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetTimes

`func (o *RsynctaskCreate0) GetTimes() bool`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *RsynctaskCreate0) GetTimesOk() (*bool, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *RsynctaskCreate0) SetTimes(v bool)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *RsynctaskCreate0) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetCompress

`func (o *RsynctaskCreate0) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *RsynctaskCreate0) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *RsynctaskCreate0) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *RsynctaskCreate0) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetArchive

`func (o *RsynctaskCreate0) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *RsynctaskCreate0) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *RsynctaskCreate0) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *RsynctaskCreate0) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetDelete

`func (o *RsynctaskCreate0) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *RsynctaskCreate0) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *RsynctaskCreate0) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *RsynctaskCreate0) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetQuiet

`func (o *RsynctaskCreate0) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *RsynctaskCreate0) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *RsynctaskCreate0) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *RsynctaskCreate0) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetPreserveperm

`func (o *RsynctaskCreate0) GetPreserveperm() bool`

GetPreserveperm returns the Preserveperm field if non-nil, zero value otherwise.

### GetPreservepermOk

`func (o *RsynctaskCreate0) GetPreservepermOk() (*bool, bool)`

GetPreservepermOk returns a tuple with the Preserveperm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveperm

`func (o *RsynctaskCreate0) SetPreserveperm(v bool)`

SetPreserveperm sets Preserveperm field to given value.

### HasPreserveperm

`func (o *RsynctaskCreate0) HasPreserveperm() bool`

HasPreserveperm returns a boolean if a field has been set.

### GetPreserveattr

`func (o *RsynctaskCreate0) GetPreserveattr() bool`

GetPreserveattr returns the Preserveattr field if non-nil, zero value otherwise.

### GetPreserveattrOk

`func (o *RsynctaskCreate0) GetPreserveattrOk() (*bool, bool)`

GetPreserveattrOk returns a tuple with the Preserveattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveattr

`func (o *RsynctaskCreate0) SetPreserveattr(v bool)`

SetPreserveattr sets Preserveattr field to given value.

### HasPreserveattr

`func (o *RsynctaskCreate0) HasPreserveattr() bool`

HasPreserveattr returns a boolean if a field has been set.

### GetDelayupdates

`func (o *RsynctaskCreate0) GetDelayupdates() bool`

GetDelayupdates returns the Delayupdates field if non-nil, zero value otherwise.

### GetDelayupdatesOk

`func (o *RsynctaskCreate0) GetDelayupdatesOk() (*bool, bool)`

GetDelayupdatesOk returns a tuple with the Delayupdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayupdates

`func (o *RsynctaskCreate0) SetDelayupdates(v bool)`

SetDelayupdates sets Delayupdates field to given value.

### HasDelayupdates

`func (o *RsynctaskCreate0) HasDelayupdates() bool`

HasDelayupdates returns a boolean if a field has been set.

### GetExtra

`func (o *RsynctaskCreate0) GetExtra() []string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *RsynctaskCreate0) GetExtraOk() (*[]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *RsynctaskCreate0) SetExtra(v []string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *RsynctaskCreate0) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetEnabled

`func (o *RsynctaskCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RsynctaskCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RsynctaskCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RsynctaskCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


