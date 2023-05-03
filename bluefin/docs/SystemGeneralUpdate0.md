# SystemGeneralUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UiHttpsport** | Pointer to **int32** |  | [optional] 
**UiHttpsredirect** | Pointer to **bool** | &#x60;ui_httpsredirect&#x60; when set, makes sure that all HTTP requests are converted to HTTPS requests to better enhance security. | [optional] 
**UiHttpsprotocols** | Pointer to **[]string** |  | [optional] [default to []]
**UiPort** | Pointer to **int32** |  | [optional] 
**UiAddress** | Pointer to **[]string** | &#x60;ui_address&#x60; and &#x60;ui_v6address&#x60; are a list of valid ipv4/ipv6 addresses respectively which the system will listen on. | [optional] [default to []]
**UiV6address** | Pointer to **[]string** | &#x60;ui_address&#x60; and &#x60;ui_v6address&#x60; are a list of valid ipv4/ipv6 addresses respectively which the system will listen on. | [optional] [default to []]
**UiAllowlist** | Pointer to **[]string** | &#x60;ui_allowlist&#x60; is a list of IP addresses and networks that are allow to use API and UI. If this list is empty, then all IP addresses are allowed to use API and UI. | [optional] [default to []]
**UiConsolemsg** | Pointer to **bool** |  | [optional] 
**UiXFrameOptions** | Pointer to **string** |  | [optional] 
**Kbdmap** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**CrashReporting** | Pointer to **NullableBool** |  | [optional] 
**UsageCollection** | Pointer to **NullableBool** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**DsAuth** | Pointer to **bool** | &#x60;ds_auth&#x60; controls whether configured Directory Service users that are granted with Privileges are allowed to log in to the Web UI or use TrueNAS API. | [optional] 
**UiCertificate** | Pointer to **NullableInt32** | &#x60;ui_certificate&#x60; is used to enable HTTPS access to the system. If &#x60;ui_certificate&#x60; is not configured on boot, it is automatically created by the system. | [optional] 
**RollbackTimeout** | Pointer to **NullableInt32** |  | [optional] 
**UiRestartDelay** | Pointer to **NullableInt32** | HTTP connections will be aborted) or specify &#x60;ui_restart_delay&#x60; (in seconds) to automatically apply them after some small amount of time necessary you might need to receive the response for your settings update request. | [optional] 

## Methods

### NewSystemGeneralUpdate0

`func NewSystemGeneralUpdate0() *SystemGeneralUpdate0`

NewSystemGeneralUpdate0 instantiates a new SystemGeneralUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemGeneralUpdate0WithDefaults

`func NewSystemGeneralUpdate0WithDefaults() *SystemGeneralUpdate0`

NewSystemGeneralUpdate0WithDefaults instantiates a new SystemGeneralUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUiHttpsport

`func (o *SystemGeneralUpdate0) GetUiHttpsport() int32`

GetUiHttpsport returns the UiHttpsport field if non-nil, zero value otherwise.

### GetUiHttpsportOk

`func (o *SystemGeneralUpdate0) GetUiHttpsportOk() (*int32, bool)`

GetUiHttpsportOk returns a tuple with the UiHttpsport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHttpsport

`func (o *SystemGeneralUpdate0) SetUiHttpsport(v int32)`

SetUiHttpsport sets UiHttpsport field to given value.

### HasUiHttpsport

`func (o *SystemGeneralUpdate0) HasUiHttpsport() bool`

HasUiHttpsport returns a boolean if a field has been set.

### GetUiHttpsredirect

`func (o *SystemGeneralUpdate0) GetUiHttpsredirect() bool`

GetUiHttpsredirect returns the UiHttpsredirect field if non-nil, zero value otherwise.

### GetUiHttpsredirectOk

`func (o *SystemGeneralUpdate0) GetUiHttpsredirectOk() (*bool, bool)`

GetUiHttpsredirectOk returns a tuple with the UiHttpsredirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHttpsredirect

`func (o *SystemGeneralUpdate0) SetUiHttpsredirect(v bool)`

SetUiHttpsredirect sets UiHttpsredirect field to given value.

### HasUiHttpsredirect

`func (o *SystemGeneralUpdate0) HasUiHttpsredirect() bool`

HasUiHttpsredirect returns a boolean if a field has been set.

### GetUiHttpsprotocols

`func (o *SystemGeneralUpdate0) GetUiHttpsprotocols() []string`

GetUiHttpsprotocols returns the UiHttpsprotocols field if non-nil, zero value otherwise.

### GetUiHttpsprotocolsOk

`func (o *SystemGeneralUpdate0) GetUiHttpsprotocolsOk() (*[]string, bool)`

GetUiHttpsprotocolsOk returns a tuple with the UiHttpsprotocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHttpsprotocols

`func (o *SystemGeneralUpdate0) SetUiHttpsprotocols(v []string)`

SetUiHttpsprotocols sets UiHttpsprotocols field to given value.

### HasUiHttpsprotocols

`func (o *SystemGeneralUpdate0) HasUiHttpsprotocols() bool`

HasUiHttpsprotocols returns a boolean if a field has been set.

### GetUiPort

`func (o *SystemGeneralUpdate0) GetUiPort() int32`

GetUiPort returns the UiPort field if non-nil, zero value otherwise.

### GetUiPortOk

`func (o *SystemGeneralUpdate0) GetUiPortOk() (*int32, bool)`

GetUiPortOk returns a tuple with the UiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiPort

`func (o *SystemGeneralUpdate0) SetUiPort(v int32)`

SetUiPort sets UiPort field to given value.

### HasUiPort

`func (o *SystemGeneralUpdate0) HasUiPort() bool`

HasUiPort returns a boolean if a field has been set.

### GetUiAddress

`func (o *SystemGeneralUpdate0) GetUiAddress() []string`

GetUiAddress returns the UiAddress field if non-nil, zero value otherwise.

### GetUiAddressOk

`func (o *SystemGeneralUpdate0) GetUiAddressOk() (*[]string, bool)`

GetUiAddressOk returns a tuple with the UiAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiAddress

`func (o *SystemGeneralUpdate0) SetUiAddress(v []string)`

SetUiAddress sets UiAddress field to given value.

### HasUiAddress

`func (o *SystemGeneralUpdate0) HasUiAddress() bool`

HasUiAddress returns a boolean if a field has been set.

### GetUiV6address

`func (o *SystemGeneralUpdate0) GetUiV6address() []string`

GetUiV6address returns the UiV6address field if non-nil, zero value otherwise.

### GetUiV6addressOk

`func (o *SystemGeneralUpdate0) GetUiV6addressOk() (*[]string, bool)`

GetUiV6addressOk returns a tuple with the UiV6address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiV6address

`func (o *SystemGeneralUpdate0) SetUiV6address(v []string)`

SetUiV6address sets UiV6address field to given value.

### HasUiV6address

`func (o *SystemGeneralUpdate0) HasUiV6address() bool`

HasUiV6address returns a boolean if a field has been set.

### GetUiAllowlist

`func (o *SystemGeneralUpdate0) GetUiAllowlist() []string`

GetUiAllowlist returns the UiAllowlist field if non-nil, zero value otherwise.

### GetUiAllowlistOk

`func (o *SystemGeneralUpdate0) GetUiAllowlistOk() (*[]string, bool)`

GetUiAllowlistOk returns a tuple with the UiAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiAllowlist

`func (o *SystemGeneralUpdate0) SetUiAllowlist(v []string)`

SetUiAllowlist sets UiAllowlist field to given value.

### HasUiAllowlist

`func (o *SystemGeneralUpdate0) HasUiAllowlist() bool`

HasUiAllowlist returns a boolean if a field has been set.

### GetUiConsolemsg

`func (o *SystemGeneralUpdate0) GetUiConsolemsg() bool`

GetUiConsolemsg returns the UiConsolemsg field if non-nil, zero value otherwise.

### GetUiConsolemsgOk

`func (o *SystemGeneralUpdate0) GetUiConsolemsgOk() (*bool, bool)`

GetUiConsolemsgOk returns a tuple with the UiConsolemsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiConsolemsg

`func (o *SystemGeneralUpdate0) SetUiConsolemsg(v bool)`

SetUiConsolemsg sets UiConsolemsg field to given value.

### HasUiConsolemsg

`func (o *SystemGeneralUpdate0) HasUiConsolemsg() bool`

HasUiConsolemsg returns a boolean if a field has been set.

### GetUiXFrameOptions

`func (o *SystemGeneralUpdate0) GetUiXFrameOptions() string`

GetUiXFrameOptions returns the UiXFrameOptions field if non-nil, zero value otherwise.

### GetUiXFrameOptionsOk

`func (o *SystemGeneralUpdate0) GetUiXFrameOptionsOk() (*string, bool)`

GetUiXFrameOptionsOk returns a tuple with the UiXFrameOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiXFrameOptions

`func (o *SystemGeneralUpdate0) SetUiXFrameOptions(v string)`

SetUiXFrameOptions sets UiXFrameOptions field to given value.

### HasUiXFrameOptions

`func (o *SystemGeneralUpdate0) HasUiXFrameOptions() bool`

HasUiXFrameOptions returns a boolean if a field has been set.

### GetKbdmap

`func (o *SystemGeneralUpdate0) GetKbdmap() string`

GetKbdmap returns the Kbdmap field if non-nil, zero value otherwise.

### GetKbdmapOk

`func (o *SystemGeneralUpdate0) GetKbdmapOk() (*string, bool)`

GetKbdmapOk returns a tuple with the Kbdmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbdmap

`func (o *SystemGeneralUpdate0) SetKbdmap(v string)`

SetKbdmap sets Kbdmap field to given value.

### HasKbdmap

`func (o *SystemGeneralUpdate0) HasKbdmap() bool`

HasKbdmap returns a boolean if a field has been set.

### GetLanguage

`func (o *SystemGeneralUpdate0) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SystemGeneralUpdate0) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SystemGeneralUpdate0) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SystemGeneralUpdate0) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTimezone

`func (o *SystemGeneralUpdate0) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SystemGeneralUpdate0) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SystemGeneralUpdate0) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SystemGeneralUpdate0) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCrashReporting

`func (o *SystemGeneralUpdate0) GetCrashReporting() bool`

GetCrashReporting returns the CrashReporting field if non-nil, zero value otherwise.

### GetCrashReportingOk

`func (o *SystemGeneralUpdate0) GetCrashReportingOk() (*bool, bool)`

GetCrashReportingOk returns a tuple with the CrashReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashReporting

`func (o *SystemGeneralUpdate0) SetCrashReporting(v bool)`

SetCrashReporting sets CrashReporting field to given value.

### HasCrashReporting

`func (o *SystemGeneralUpdate0) HasCrashReporting() bool`

HasCrashReporting returns a boolean if a field has been set.

### SetCrashReportingNil

`func (o *SystemGeneralUpdate0) SetCrashReportingNil(b bool)`

 SetCrashReportingNil sets the value for CrashReporting to be an explicit nil

### UnsetCrashReporting
`func (o *SystemGeneralUpdate0) UnsetCrashReporting()`

UnsetCrashReporting ensures that no value is present for CrashReporting, not even an explicit nil
### GetUsageCollection

`func (o *SystemGeneralUpdate0) GetUsageCollection() bool`

GetUsageCollection returns the UsageCollection field if non-nil, zero value otherwise.

### GetUsageCollectionOk

`func (o *SystemGeneralUpdate0) GetUsageCollectionOk() (*bool, bool)`

GetUsageCollectionOk returns a tuple with the UsageCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCollection

`func (o *SystemGeneralUpdate0) SetUsageCollection(v bool)`

SetUsageCollection sets UsageCollection field to given value.

### HasUsageCollection

`func (o *SystemGeneralUpdate0) HasUsageCollection() bool`

HasUsageCollection returns a boolean if a field has been set.

### SetUsageCollectionNil

`func (o *SystemGeneralUpdate0) SetUsageCollectionNil(b bool)`

 SetUsageCollectionNil sets the value for UsageCollection to be an explicit nil

### UnsetUsageCollection
`func (o *SystemGeneralUpdate0) UnsetUsageCollection()`

UnsetUsageCollection ensures that no value is present for UsageCollection, not even an explicit nil
### GetBirthday

`func (o *SystemGeneralUpdate0) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *SystemGeneralUpdate0) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *SystemGeneralUpdate0) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *SystemGeneralUpdate0) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetDsAuth

`func (o *SystemGeneralUpdate0) GetDsAuth() bool`

GetDsAuth returns the DsAuth field if non-nil, zero value otherwise.

### GetDsAuthOk

`func (o *SystemGeneralUpdate0) GetDsAuthOk() (*bool, bool)`

GetDsAuthOk returns a tuple with the DsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsAuth

`func (o *SystemGeneralUpdate0) SetDsAuth(v bool)`

SetDsAuth sets DsAuth field to given value.

### HasDsAuth

`func (o *SystemGeneralUpdate0) HasDsAuth() bool`

HasDsAuth returns a boolean if a field has been set.

### GetUiCertificate

`func (o *SystemGeneralUpdate0) GetUiCertificate() int32`

GetUiCertificate returns the UiCertificate field if non-nil, zero value otherwise.

### GetUiCertificateOk

`func (o *SystemGeneralUpdate0) GetUiCertificateOk() (*int32, bool)`

GetUiCertificateOk returns a tuple with the UiCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCertificate

`func (o *SystemGeneralUpdate0) SetUiCertificate(v int32)`

SetUiCertificate sets UiCertificate field to given value.

### HasUiCertificate

`func (o *SystemGeneralUpdate0) HasUiCertificate() bool`

HasUiCertificate returns a boolean if a field has been set.

### SetUiCertificateNil

`func (o *SystemGeneralUpdate0) SetUiCertificateNil(b bool)`

 SetUiCertificateNil sets the value for UiCertificate to be an explicit nil

### UnsetUiCertificate
`func (o *SystemGeneralUpdate0) UnsetUiCertificate()`

UnsetUiCertificate ensures that no value is present for UiCertificate, not even an explicit nil
### GetRollbackTimeout

`func (o *SystemGeneralUpdate0) GetRollbackTimeout() int32`

GetRollbackTimeout returns the RollbackTimeout field if non-nil, zero value otherwise.

### GetRollbackTimeoutOk

`func (o *SystemGeneralUpdate0) GetRollbackTimeoutOk() (*int32, bool)`

GetRollbackTimeoutOk returns a tuple with the RollbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackTimeout

`func (o *SystemGeneralUpdate0) SetRollbackTimeout(v int32)`

SetRollbackTimeout sets RollbackTimeout field to given value.

### HasRollbackTimeout

`func (o *SystemGeneralUpdate0) HasRollbackTimeout() bool`

HasRollbackTimeout returns a boolean if a field has been set.

### SetRollbackTimeoutNil

`func (o *SystemGeneralUpdate0) SetRollbackTimeoutNil(b bool)`

 SetRollbackTimeoutNil sets the value for RollbackTimeout to be an explicit nil

### UnsetRollbackTimeout
`func (o *SystemGeneralUpdate0) UnsetRollbackTimeout()`

UnsetRollbackTimeout ensures that no value is present for RollbackTimeout, not even an explicit nil
### GetUiRestartDelay

`func (o *SystemGeneralUpdate0) GetUiRestartDelay() int32`

GetUiRestartDelay returns the UiRestartDelay field if non-nil, zero value otherwise.

### GetUiRestartDelayOk

`func (o *SystemGeneralUpdate0) GetUiRestartDelayOk() (*int32, bool)`

GetUiRestartDelayOk returns a tuple with the UiRestartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiRestartDelay

`func (o *SystemGeneralUpdate0) SetUiRestartDelay(v int32)`

SetUiRestartDelay sets UiRestartDelay field to given value.

### HasUiRestartDelay

`func (o *SystemGeneralUpdate0) HasUiRestartDelay() bool`

HasUiRestartDelay returns a boolean if a field has been set.

### SetUiRestartDelayNil

`func (o *SystemGeneralUpdate0) SetUiRestartDelayNil(b bool)`

 SetUiRestartDelayNil sets the value for UiRestartDelay to be an explicit nil

### UnsetUiRestartDelay
`func (o *SystemGeneralUpdate0) UnsetUiRestartDelay()`

UnsetUiRestartDelay ensures that no value is present for UiRestartDelay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


