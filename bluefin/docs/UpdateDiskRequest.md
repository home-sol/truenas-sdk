# UpdateDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** |  | 
**Lunid** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**Critical** | **NullableInt32** |  | 
**Difference** | **NullableInt32** |  | 
**Informational** | **NullableInt32** |  | 
**Hddstandby** | [**HDDStandby**](HDDStandby.md) |  | 
**Advpowermgmt** | [**AdvPowermgmt**](AdvPowermgmt.md) |  | 
**Togglesmart** | **bool** |  | 
**SupportsSmart** | Pointer to **bool** |  | [optional] 
**Smartoptions** | Pointer to **string** |  | [optional] 
**Passwd** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDiskRequest

`func NewUpdateDiskRequest(number int32, lunid NullableString, description NullableString, critical NullableInt32, difference NullableInt32, informational NullableInt32, hddstandby HDDStandby, advpowermgmt AdvPowermgmt, togglesmart bool, ) *UpdateDiskRequest`

NewUpdateDiskRequest instantiates a new UpdateDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDiskRequestWithDefaults

`func NewUpdateDiskRequestWithDefaults() *UpdateDiskRequest`

NewUpdateDiskRequestWithDefaults instantiates a new UpdateDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *UpdateDiskRequest) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UpdateDiskRequest) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UpdateDiskRequest) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetLunid

`func (o *UpdateDiskRequest) GetLunid() string`

GetLunid returns the Lunid field if non-nil, zero value otherwise.

### GetLunidOk

`func (o *UpdateDiskRequest) GetLunidOk() (*string, bool)`

GetLunidOk returns a tuple with the Lunid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunid

`func (o *UpdateDiskRequest) SetLunid(v string)`

SetLunid sets Lunid field to given value.


### SetLunidNil

`func (o *UpdateDiskRequest) SetLunidNil(b bool)`

 SetLunidNil sets the value for Lunid to be an explicit nil

### UnsetLunid
`func (o *UpdateDiskRequest) UnsetLunid()`

UnsetLunid ensures that no value is present for Lunid, not even an explicit nil
### GetDescription

`func (o *UpdateDiskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDiskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDiskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *UpdateDiskRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateDiskRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCritical

`func (o *UpdateDiskRequest) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *UpdateDiskRequest) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *UpdateDiskRequest) SetCritical(v int32)`

SetCritical sets Critical field to given value.


### SetCriticalNil

`func (o *UpdateDiskRequest) SetCriticalNil(b bool)`

 SetCriticalNil sets the value for Critical to be an explicit nil

### UnsetCritical
`func (o *UpdateDiskRequest) UnsetCritical()`

UnsetCritical ensures that no value is present for Critical, not even an explicit nil
### GetDifference

`func (o *UpdateDiskRequest) GetDifference() int32`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *UpdateDiskRequest) GetDifferenceOk() (*int32, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *UpdateDiskRequest) SetDifference(v int32)`

SetDifference sets Difference field to given value.


### SetDifferenceNil

`func (o *UpdateDiskRequest) SetDifferenceNil(b bool)`

 SetDifferenceNil sets the value for Difference to be an explicit nil

### UnsetDifference
`func (o *UpdateDiskRequest) UnsetDifference()`

UnsetDifference ensures that no value is present for Difference, not even an explicit nil
### GetInformational

`func (o *UpdateDiskRequest) GetInformational() int32`

GetInformational returns the Informational field if non-nil, zero value otherwise.

### GetInformationalOk

`func (o *UpdateDiskRequest) GetInformationalOk() (*int32, bool)`

GetInformationalOk returns a tuple with the Informational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformational

`func (o *UpdateDiskRequest) SetInformational(v int32)`

SetInformational sets Informational field to given value.


### SetInformationalNil

`func (o *UpdateDiskRequest) SetInformationalNil(b bool)`

 SetInformationalNil sets the value for Informational to be an explicit nil

### UnsetInformational
`func (o *UpdateDiskRequest) UnsetInformational()`

UnsetInformational ensures that no value is present for Informational, not even an explicit nil
### GetHddstandby

`func (o *UpdateDiskRequest) GetHddstandby() HDDStandby`

GetHddstandby returns the Hddstandby field if non-nil, zero value otherwise.

### GetHddstandbyOk

`func (o *UpdateDiskRequest) GetHddstandbyOk() (*HDDStandby, bool)`

GetHddstandbyOk returns a tuple with the Hddstandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddstandby

`func (o *UpdateDiskRequest) SetHddstandby(v HDDStandby)`

SetHddstandby sets Hddstandby field to given value.


### GetAdvpowermgmt

`func (o *UpdateDiskRequest) GetAdvpowermgmt() AdvPowermgmt`

GetAdvpowermgmt returns the Advpowermgmt field if non-nil, zero value otherwise.

### GetAdvpowermgmtOk

`func (o *UpdateDiskRequest) GetAdvpowermgmtOk() (*AdvPowermgmt, bool)`

GetAdvpowermgmtOk returns a tuple with the Advpowermgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvpowermgmt

`func (o *UpdateDiskRequest) SetAdvpowermgmt(v AdvPowermgmt)`

SetAdvpowermgmt sets Advpowermgmt field to given value.


### GetTogglesmart

`func (o *UpdateDiskRequest) GetTogglesmart() bool`

GetTogglesmart returns the Togglesmart field if non-nil, zero value otherwise.

### GetTogglesmartOk

`func (o *UpdateDiskRequest) GetTogglesmartOk() (*bool, bool)`

GetTogglesmartOk returns a tuple with the Togglesmart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglesmart

`func (o *UpdateDiskRequest) SetTogglesmart(v bool)`

SetTogglesmart sets Togglesmart field to given value.


### GetSupportsSmart

`func (o *UpdateDiskRequest) GetSupportsSmart() bool`

GetSupportsSmart returns the SupportsSmart field if non-nil, zero value otherwise.

### GetSupportsSmartOk

`func (o *UpdateDiskRequest) GetSupportsSmartOk() (*bool, bool)`

GetSupportsSmartOk returns a tuple with the SupportsSmart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSmart

`func (o *UpdateDiskRequest) SetSupportsSmart(v bool)`

SetSupportsSmart sets SupportsSmart field to given value.

### HasSupportsSmart

`func (o *UpdateDiskRequest) HasSupportsSmart() bool`

HasSupportsSmart returns a boolean if a field has been set.

### GetSmartoptions

`func (o *UpdateDiskRequest) GetSmartoptions() string`

GetSmartoptions returns the Smartoptions field if non-nil, zero value otherwise.

### GetSmartoptionsOk

`func (o *UpdateDiskRequest) GetSmartoptionsOk() (*string, bool)`

GetSmartoptionsOk returns a tuple with the Smartoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartoptions

`func (o *UpdateDiskRequest) SetSmartoptions(v string)`

SetSmartoptions sets Smartoptions field to given value.

### HasSmartoptions

`func (o *UpdateDiskRequest) HasSmartoptions() bool`

HasSmartoptions returns a boolean if a field has been set.

### GetPasswd

`func (o *UpdateDiskRequest) GetPasswd() string`

GetPasswd returns the Passwd field if non-nil, zero value otherwise.

### GetPasswdOk

`func (o *UpdateDiskRequest) GetPasswdOk() (*string, bool)`

GetPasswdOk returns a tuple with the Passwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswd

`func (o *UpdateDiskRequest) SetPasswd(v string)`

SetPasswd sets Passwd field to given value.

### HasPasswd

`func (o *UpdateDiskRequest) HasPasswd() bool`

HasPasswd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


