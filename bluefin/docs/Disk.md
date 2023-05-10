# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Subsystem** | **string** |  | [readonly] 
**Number** | **int32** |  | 
**Serial** | **string** |  | [readonly] 
**Lunid** | **NullableString** |  | 
**Size** | **int64** |  | [readonly] 
**Description** | **string** |  | 
**Transfermode** | **string** |  | [readonly] 
**Hddstandby** | [**HDDStandby**](HDDStandby.md) |  | 
**Advpowermgmt** | [**AdvPowermgmt**](AdvPowermgmt.md) |  | 
**Togglesmart** | **bool** |  | 
**Smartoptions** | **string** |  | 
**Expiretime** | **NullableString** |  | [readonly] 
**Critical** | **NullableInt32** |  | 
**Difference** | **NullableInt32** |  | 
**Informational** | **NullableInt32** |  | 
**Model** | **NullableString** |  | [readonly] 
**Rotationrate** | **NullableInt32** |  | [readonly] 
**Type** | **NullableString** |  | [readonly] 
**ZfsGuid** | **NullableString** |  | [readonly] 
**Bus** | **string** |  | 
**Devname** | **string** |  | [readonly] 
**Enclosure** | [**DiskEnclosure**](DiskEnclosure.md) |  | 
**Pool** | **NullableString** |  | 
**Passwd** | Pointer to **string** |  | [optional] 
**KmipUid** | Pointer to **NullableString** |  | [optional] [readonly] 
**SupportsSmart** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDisk

`func NewDisk(identifier string, name string, subsystem string, number int32, serial string, lunid NullableString, size int64, description string, transfermode string, hddstandby HDDStandby, advpowermgmt AdvPowermgmt, togglesmart bool, smartoptions string, expiretime NullableString, critical NullableInt32, difference NullableInt32, informational NullableInt32, model NullableString, rotationrate NullableInt32, type_ NullableString, zfsGuid NullableString, bus string, devname string, enclosure DiskEnclosure, pool NullableString, ) *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *Disk) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Disk) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Disk) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *Disk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Disk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Disk) SetName(v string)`

SetName sets Name field to given value.


### GetSubsystem

`func (o *Disk) GetSubsystem() string`

GetSubsystem returns the Subsystem field if non-nil, zero value otherwise.

### GetSubsystemOk

`func (o *Disk) GetSubsystemOk() (*string, bool)`

GetSubsystemOk returns a tuple with the Subsystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystem

`func (o *Disk) SetSubsystem(v string)`

SetSubsystem sets Subsystem field to given value.


### GetNumber

`func (o *Disk) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Disk) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Disk) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetSerial

`func (o *Disk) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Disk) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Disk) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetLunid

`func (o *Disk) GetLunid() string`

GetLunid returns the Lunid field if non-nil, zero value otherwise.

### GetLunidOk

`func (o *Disk) GetLunidOk() (*string, bool)`

GetLunidOk returns a tuple with the Lunid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunid

`func (o *Disk) SetLunid(v string)`

SetLunid sets Lunid field to given value.


### SetLunidNil

`func (o *Disk) SetLunidNil(b bool)`

 SetLunidNil sets the value for Lunid to be an explicit nil

### UnsetLunid
`func (o *Disk) UnsetLunid()`

UnsetLunid ensures that no value is present for Lunid, not even an explicit nil
### GetSize

`func (o *Disk) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Disk) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Disk) SetSize(v int64)`

SetSize sets Size field to given value.


### GetDescription

`func (o *Disk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Disk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Disk) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTransfermode

`func (o *Disk) GetTransfermode() string`

GetTransfermode returns the Transfermode field if non-nil, zero value otherwise.

### GetTransfermodeOk

`func (o *Disk) GetTransfermodeOk() (*string, bool)`

GetTransfermodeOk returns a tuple with the Transfermode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfermode

`func (o *Disk) SetTransfermode(v string)`

SetTransfermode sets Transfermode field to given value.


### GetHddstandby

`func (o *Disk) GetHddstandby() HDDStandby`

GetHddstandby returns the Hddstandby field if non-nil, zero value otherwise.

### GetHddstandbyOk

`func (o *Disk) GetHddstandbyOk() (*HDDStandby, bool)`

GetHddstandbyOk returns a tuple with the Hddstandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddstandby

`func (o *Disk) SetHddstandby(v HDDStandby)`

SetHddstandby sets Hddstandby field to given value.


### GetAdvpowermgmt

`func (o *Disk) GetAdvpowermgmt() AdvPowermgmt`

GetAdvpowermgmt returns the Advpowermgmt field if non-nil, zero value otherwise.

### GetAdvpowermgmtOk

`func (o *Disk) GetAdvpowermgmtOk() (*AdvPowermgmt, bool)`

GetAdvpowermgmtOk returns a tuple with the Advpowermgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvpowermgmt

`func (o *Disk) SetAdvpowermgmt(v AdvPowermgmt)`

SetAdvpowermgmt sets Advpowermgmt field to given value.


### GetTogglesmart

`func (o *Disk) GetTogglesmart() bool`

GetTogglesmart returns the Togglesmart field if non-nil, zero value otherwise.

### GetTogglesmartOk

`func (o *Disk) GetTogglesmartOk() (*bool, bool)`

GetTogglesmartOk returns a tuple with the Togglesmart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglesmart

`func (o *Disk) SetTogglesmart(v bool)`

SetTogglesmart sets Togglesmart field to given value.


### GetSmartoptions

`func (o *Disk) GetSmartoptions() string`

GetSmartoptions returns the Smartoptions field if non-nil, zero value otherwise.

### GetSmartoptionsOk

`func (o *Disk) GetSmartoptionsOk() (*string, bool)`

GetSmartoptionsOk returns a tuple with the Smartoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartoptions

`func (o *Disk) SetSmartoptions(v string)`

SetSmartoptions sets Smartoptions field to given value.


### GetExpiretime

`func (o *Disk) GetExpiretime() string`

GetExpiretime returns the Expiretime field if non-nil, zero value otherwise.

### GetExpiretimeOk

`func (o *Disk) GetExpiretimeOk() (*string, bool)`

GetExpiretimeOk returns a tuple with the Expiretime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiretime

`func (o *Disk) SetExpiretime(v string)`

SetExpiretime sets Expiretime field to given value.


### SetExpiretimeNil

`func (o *Disk) SetExpiretimeNil(b bool)`

 SetExpiretimeNil sets the value for Expiretime to be an explicit nil

### UnsetExpiretime
`func (o *Disk) UnsetExpiretime()`

UnsetExpiretime ensures that no value is present for Expiretime, not even an explicit nil
### GetCritical

`func (o *Disk) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *Disk) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *Disk) SetCritical(v int32)`

SetCritical sets Critical field to given value.


### SetCriticalNil

`func (o *Disk) SetCriticalNil(b bool)`

 SetCriticalNil sets the value for Critical to be an explicit nil

### UnsetCritical
`func (o *Disk) UnsetCritical()`

UnsetCritical ensures that no value is present for Critical, not even an explicit nil
### GetDifference

`func (o *Disk) GetDifference() int32`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *Disk) GetDifferenceOk() (*int32, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *Disk) SetDifference(v int32)`

SetDifference sets Difference field to given value.


### SetDifferenceNil

`func (o *Disk) SetDifferenceNil(b bool)`

 SetDifferenceNil sets the value for Difference to be an explicit nil

### UnsetDifference
`func (o *Disk) UnsetDifference()`

UnsetDifference ensures that no value is present for Difference, not even an explicit nil
### GetInformational

`func (o *Disk) GetInformational() int32`

GetInformational returns the Informational field if non-nil, zero value otherwise.

### GetInformationalOk

`func (o *Disk) GetInformationalOk() (*int32, bool)`

GetInformationalOk returns a tuple with the Informational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformational

`func (o *Disk) SetInformational(v int32)`

SetInformational sets Informational field to given value.


### SetInformationalNil

`func (o *Disk) SetInformationalNil(b bool)`

 SetInformationalNil sets the value for Informational to be an explicit nil

### UnsetInformational
`func (o *Disk) UnsetInformational()`

UnsetInformational ensures that no value is present for Informational, not even an explicit nil
### GetModel

`func (o *Disk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Disk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Disk) SetModel(v string)`

SetModel sets Model field to given value.


### SetModelNil

`func (o *Disk) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Disk) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetRotationrate

`func (o *Disk) GetRotationrate() int32`

GetRotationrate returns the Rotationrate field if non-nil, zero value otherwise.

### GetRotationrateOk

`func (o *Disk) GetRotationrateOk() (*int32, bool)`

GetRotationrateOk returns a tuple with the Rotationrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationrate

`func (o *Disk) SetRotationrate(v int32)`

SetRotationrate sets Rotationrate field to given value.


### SetRotationrateNil

`func (o *Disk) SetRotationrateNil(b bool)`

 SetRotationrateNil sets the value for Rotationrate to be an explicit nil

### UnsetRotationrate
`func (o *Disk) UnsetRotationrate()`

UnsetRotationrate ensures that no value is present for Rotationrate, not even an explicit nil
### GetType

`func (o *Disk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Disk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Disk) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Disk) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Disk) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetZfsGuid

`func (o *Disk) GetZfsGuid() string`

GetZfsGuid returns the ZfsGuid field if non-nil, zero value otherwise.

### GetZfsGuidOk

`func (o *Disk) GetZfsGuidOk() (*string, bool)`

GetZfsGuidOk returns a tuple with the ZfsGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZfsGuid

`func (o *Disk) SetZfsGuid(v string)`

SetZfsGuid sets ZfsGuid field to given value.


### SetZfsGuidNil

`func (o *Disk) SetZfsGuidNil(b bool)`

 SetZfsGuidNil sets the value for ZfsGuid to be an explicit nil

### UnsetZfsGuid
`func (o *Disk) UnsetZfsGuid()`

UnsetZfsGuid ensures that no value is present for ZfsGuid, not even an explicit nil
### GetBus

`func (o *Disk) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *Disk) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *Disk) SetBus(v string)`

SetBus sets Bus field to given value.


### GetDevname

`func (o *Disk) GetDevname() string`

GetDevname returns the Devname field if non-nil, zero value otherwise.

### GetDevnameOk

`func (o *Disk) GetDevnameOk() (*string, bool)`

GetDevnameOk returns a tuple with the Devname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevname

`func (o *Disk) SetDevname(v string)`

SetDevname sets Devname field to given value.


### GetEnclosure

`func (o *Disk) GetEnclosure() DiskEnclosure`

GetEnclosure returns the Enclosure field if non-nil, zero value otherwise.

### GetEnclosureOk

`func (o *Disk) GetEnclosureOk() (*DiskEnclosure, bool)`

GetEnclosureOk returns a tuple with the Enclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosure

`func (o *Disk) SetEnclosure(v DiskEnclosure)`

SetEnclosure sets Enclosure field to given value.


### GetPool

`func (o *Disk) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Disk) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Disk) SetPool(v string)`

SetPool sets Pool field to given value.


### SetPoolNil

`func (o *Disk) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *Disk) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetPasswd

`func (o *Disk) GetPasswd() string`

GetPasswd returns the Passwd field if non-nil, zero value otherwise.

### GetPasswdOk

`func (o *Disk) GetPasswdOk() (*string, bool)`

GetPasswdOk returns a tuple with the Passwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswd

`func (o *Disk) SetPasswd(v string)`

SetPasswd sets Passwd field to given value.

### HasPasswd

`func (o *Disk) HasPasswd() bool`

HasPasswd returns a boolean if a field has been set.

### GetKmipUid

`func (o *Disk) GetKmipUid() string`

GetKmipUid returns the KmipUid field if non-nil, zero value otherwise.

### GetKmipUidOk

`func (o *Disk) GetKmipUidOk() (*string, bool)`

GetKmipUidOk returns a tuple with the KmipUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipUid

`func (o *Disk) SetKmipUid(v string)`

SetKmipUid sets KmipUid field to given value.

### HasKmipUid

`func (o *Disk) HasKmipUid() bool`

HasKmipUid returns a boolean if a field has been set.

### SetKmipUidNil

`func (o *Disk) SetKmipUidNil(b bool)`

 SetKmipUidNil sets the value for KmipUid to be an explicit nil

### UnsetKmipUid
`func (o *Disk) UnsetKmipUid()`

UnsetKmipUid ensures that no value is present for KmipUid, not even an explicit nil
### GetSupportsSmart

`func (o *Disk) GetSupportsSmart() bool`

GetSupportsSmart returns the SupportsSmart field if non-nil, zero value otherwise.

### GetSupportsSmartOk

`func (o *Disk) GetSupportsSmartOk() (*bool, bool)`

GetSupportsSmartOk returns a tuple with the SupportsSmart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSmart

`func (o *Disk) SetSupportsSmart(v bool)`

SetSupportsSmart sets SupportsSmart field to given value.

### HasSupportsSmart

`func (o *Disk) HasSupportsSmart() bool`

HasSupportsSmart returns a boolean if a field has been set.

### SetSupportsSmartNil

`func (o *Disk) SetSupportsSmartNil(b bool)`

 SetSupportsSmartNil sets the value for SupportsSmart to be an explicit nil

### UnsetSupportsSmart
`func (o *Disk) UnsetSupportsSmart()`

UnsetSupportsSmart ensures that no value is present for SupportsSmart, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


