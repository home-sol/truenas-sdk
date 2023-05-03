# InterfaceCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Type** | Pointer to **string** | For BRIDGE &#x60;type&#x60; the following attribute is required: bridge_members. For LINK_AGGREGATION &#x60;type&#x60; the following attributes are required: lag_ports, lag_protocol. For VLAN &#x60;type&#x60; the following attributes are required: vlan_parent_interface, vlan_tag and vlan_pcp. | [optional] 
**Ipv4Dhcp** | Pointer to **bool** |  | [optional] [default to false]
**Ipv6Auto** | Pointer to **bool** |  | [optional] [default to false]
**Aliases** | Pointer to [**[]InterfaceAlias**](InterfaceAlias.md) |  | [optional] [default to []]
**FailoverCritical** | Pointer to **bool** |  | [optional] [default to false]
**FailoverGroup** | Pointer to **NullableInt32** |  | [optional] 
**FailoverVhid** | Pointer to **NullableInt32** |  | [optional] 
**FailoverAliases** | Pointer to [**[]InterfaceFailoverAlias**](InterfaceFailoverAlias.md) |  | [optional] [default to []]
**FailoverVirtualAliases** | Pointer to [**[]InterfaceVirtualAlias**](InterfaceVirtualAlias.md) |  | [optional] [default to []]
**BridgeMembers** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Stp** | Pointer to **bool** |  | [optional] [default to true]
**LagProtocol** | Pointer to **string** |  | [optional] 
**XmitHashPolicy** | Pointer to **NullableString** |  | [optional] 
**LacpduRate** | Pointer to **NullableString** |  | [optional] 
**LagPorts** | Pointer to **[]string** |  | [optional] [default to []]
**VlanParentInterface** | Pointer to **string** |  | [optional] 
**VlanTag** | Pointer to **int32** |  | [optional] 
**VlanPcp** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewInterfaceCreate0

`func NewInterfaceCreate0() *InterfaceCreate0`

NewInterfaceCreate0 instantiates a new InterfaceCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceCreate0WithDefaults

`func NewInterfaceCreate0WithDefaults() *InterfaceCreate0`

NewInterfaceCreate0WithDefaults instantiates a new InterfaceCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InterfaceCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InterfaceCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceCreate0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceCreate0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceCreate0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceCreate0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *InterfaceCreate0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceCreate0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceCreate0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InterfaceCreate0) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpv4Dhcp

`func (o *InterfaceCreate0) GetIpv4Dhcp() bool`

GetIpv4Dhcp returns the Ipv4Dhcp field if non-nil, zero value otherwise.

### GetIpv4DhcpOk

`func (o *InterfaceCreate0) GetIpv4DhcpOk() (*bool, bool)`

GetIpv4DhcpOk returns a tuple with the Ipv4Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Dhcp

`func (o *InterfaceCreate0) SetIpv4Dhcp(v bool)`

SetIpv4Dhcp sets Ipv4Dhcp field to given value.

### HasIpv4Dhcp

`func (o *InterfaceCreate0) HasIpv4Dhcp() bool`

HasIpv4Dhcp returns a boolean if a field has been set.

### GetIpv6Auto

`func (o *InterfaceCreate0) GetIpv6Auto() bool`

GetIpv6Auto returns the Ipv6Auto field if non-nil, zero value otherwise.

### GetIpv6AutoOk

`func (o *InterfaceCreate0) GetIpv6AutoOk() (*bool, bool)`

GetIpv6AutoOk returns a tuple with the Ipv6Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Auto

`func (o *InterfaceCreate0) SetIpv6Auto(v bool)`

SetIpv6Auto sets Ipv6Auto field to given value.

### HasIpv6Auto

`func (o *InterfaceCreate0) HasIpv6Auto() bool`

HasIpv6Auto returns a boolean if a field has been set.

### GetAliases

`func (o *InterfaceCreate0) GetAliases() []InterfaceAlias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InterfaceCreate0) GetAliasesOk() (*[]InterfaceAlias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InterfaceCreate0) SetAliases(v []InterfaceAlias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InterfaceCreate0) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFailoverCritical

`func (o *InterfaceCreate0) GetFailoverCritical() bool`

GetFailoverCritical returns the FailoverCritical field if non-nil, zero value otherwise.

### GetFailoverCriticalOk

`func (o *InterfaceCreate0) GetFailoverCriticalOk() (*bool, bool)`

GetFailoverCriticalOk returns a tuple with the FailoverCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverCritical

`func (o *InterfaceCreate0) SetFailoverCritical(v bool)`

SetFailoverCritical sets FailoverCritical field to given value.

### HasFailoverCritical

`func (o *InterfaceCreate0) HasFailoverCritical() bool`

HasFailoverCritical returns a boolean if a field has been set.

### GetFailoverGroup

`func (o *InterfaceCreate0) GetFailoverGroup() int32`

GetFailoverGroup returns the FailoverGroup field if non-nil, zero value otherwise.

### GetFailoverGroupOk

`func (o *InterfaceCreate0) GetFailoverGroupOk() (*int32, bool)`

GetFailoverGroupOk returns a tuple with the FailoverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverGroup

`func (o *InterfaceCreate0) SetFailoverGroup(v int32)`

SetFailoverGroup sets FailoverGroup field to given value.

### HasFailoverGroup

`func (o *InterfaceCreate0) HasFailoverGroup() bool`

HasFailoverGroup returns a boolean if a field has been set.

### SetFailoverGroupNil

`func (o *InterfaceCreate0) SetFailoverGroupNil(b bool)`

 SetFailoverGroupNil sets the value for FailoverGroup to be an explicit nil

### UnsetFailoverGroup
`func (o *InterfaceCreate0) UnsetFailoverGroup()`

UnsetFailoverGroup ensures that no value is present for FailoverGroup, not even an explicit nil
### GetFailoverVhid

`func (o *InterfaceCreate0) GetFailoverVhid() int32`

GetFailoverVhid returns the FailoverVhid field if non-nil, zero value otherwise.

### GetFailoverVhidOk

`func (o *InterfaceCreate0) GetFailoverVhidOk() (*int32, bool)`

GetFailoverVhidOk returns a tuple with the FailoverVhid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverVhid

`func (o *InterfaceCreate0) SetFailoverVhid(v int32)`

SetFailoverVhid sets FailoverVhid field to given value.

### HasFailoverVhid

`func (o *InterfaceCreate0) HasFailoverVhid() bool`

HasFailoverVhid returns a boolean if a field has been set.

### SetFailoverVhidNil

`func (o *InterfaceCreate0) SetFailoverVhidNil(b bool)`

 SetFailoverVhidNil sets the value for FailoverVhid to be an explicit nil

### UnsetFailoverVhid
`func (o *InterfaceCreate0) UnsetFailoverVhid()`

UnsetFailoverVhid ensures that no value is present for FailoverVhid, not even an explicit nil
### GetFailoverAliases

`func (o *InterfaceCreate0) GetFailoverAliases() []InterfaceFailoverAlias`

GetFailoverAliases returns the FailoverAliases field if non-nil, zero value otherwise.

### GetFailoverAliasesOk

`func (o *InterfaceCreate0) GetFailoverAliasesOk() (*[]InterfaceFailoverAlias, bool)`

GetFailoverAliasesOk returns a tuple with the FailoverAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAliases

`func (o *InterfaceCreate0) SetFailoverAliases(v []InterfaceFailoverAlias)`

SetFailoverAliases sets FailoverAliases field to given value.

### HasFailoverAliases

`func (o *InterfaceCreate0) HasFailoverAliases() bool`

HasFailoverAliases returns a boolean if a field has been set.

### GetFailoverVirtualAliases

`func (o *InterfaceCreate0) GetFailoverVirtualAliases() []InterfaceVirtualAlias`

GetFailoverVirtualAliases returns the FailoverVirtualAliases field if non-nil, zero value otherwise.

### GetFailoverVirtualAliasesOk

`func (o *InterfaceCreate0) GetFailoverVirtualAliasesOk() (*[]InterfaceVirtualAlias, bool)`

GetFailoverVirtualAliasesOk returns a tuple with the FailoverVirtualAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverVirtualAliases

`func (o *InterfaceCreate0) SetFailoverVirtualAliases(v []InterfaceVirtualAlias)`

SetFailoverVirtualAliases sets FailoverVirtualAliases field to given value.

### HasFailoverVirtualAliases

`func (o *InterfaceCreate0) HasFailoverVirtualAliases() bool`

HasFailoverVirtualAliases returns a boolean if a field has been set.

### GetBridgeMembers

`func (o *InterfaceCreate0) GetBridgeMembers() []interface{}`

GetBridgeMembers returns the BridgeMembers field if non-nil, zero value otherwise.

### GetBridgeMembersOk

`func (o *InterfaceCreate0) GetBridgeMembersOk() (*[]interface{}, bool)`

GetBridgeMembersOk returns a tuple with the BridgeMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeMembers

`func (o *InterfaceCreate0) SetBridgeMembers(v []interface{})`

SetBridgeMembers sets BridgeMembers field to given value.

### HasBridgeMembers

`func (o *InterfaceCreate0) HasBridgeMembers() bool`

HasBridgeMembers returns a boolean if a field has been set.

### GetStp

`func (o *InterfaceCreate0) GetStp() bool`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *InterfaceCreate0) GetStpOk() (*bool, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *InterfaceCreate0) SetStp(v bool)`

SetStp sets Stp field to given value.

### HasStp

`func (o *InterfaceCreate0) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetLagProtocol

`func (o *InterfaceCreate0) GetLagProtocol() string`

GetLagProtocol returns the LagProtocol field if non-nil, zero value otherwise.

### GetLagProtocolOk

`func (o *InterfaceCreate0) GetLagProtocolOk() (*string, bool)`

GetLagProtocolOk returns a tuple with the LagProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagProtocol

`func (o *InterfaceCreate0) SetLagProtocol(v string)`

SetLagProtocol sets LagProtocol field to given value.

### HasLagProtocol

`func (o *InterfaceCreate0) HasLagProtocol() bool`

HasLagProtocol returns a boolean if a field has been set.

### GetXmitHashPolicy

`func (o *InterfaceCreate0) GetXmitHashPolicy() string`

GetXmitHashPolicy returns the XmitHashPolicy field if non-nil, zero value otherwise.

### GetXmitHashPolicyOk

`func (o *InterfaceCreate0) GetXmitHashPolicyOk() (*string, bool)`

GetXmitHashPolicyOk returns a tuple with the XmitHashPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmitHashPolicy

`func (o *InterfaceCreate0) SetXmitHashPolicy(v string)`

SetXmitHashPolicy sets XmitHashPolicy field to given value.

### HasXmitHashPolicy

`func (o *InterfaceCreate0) HasXmitHashPolicy() bool`

HasXmitHashPolicy returns a boolean if a field has been set.

### SetXmitHashPolicyNil

`func (o *InterfaceCreate0) SetXmitHashPolicyNil(b bool)`

 SetXmitHashPolicyNil sets the value for XmitHashPolicy to be an explicit nil

### UnsetXmitHashPolicy
`func (o *InterfaceCreate0) UnsetXmitHashPolicy()`

UnsetXmitHashPolicy ensures that no value is present for XmitHashPolicy, not even an explicit nil
### GetLacpduRate

`func (o *InterfaceCreate0) GetLacpduRate() string`

GetLacpduRate returns the LacpduRate field if non-nil, zero value otherwise.

### GetLacpduRateOk

`func (o *InterfaceCreate0) GetLacpduRateOk() (*string, bool)`

GetLacpduRateOk returns a tuple with the LacpduRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacpduRate

`func (o *InterfaceCreate0) SetLacpduRate(v string)`

SetLacpduRate sets LacpduRate field to given value.

### HasLacpduRate

`func (o *InterfaceCreate0) HasLacpduRate() bool`

HasLacpduRate returns a boolean if a field has been set.

### SetLacpduRateNil

`func (o *InterfaceCreate0) SetLacpduRateNil(b bool)`

 SetLacpduRateNil sets the value for LacpduRate to be an explicit nil

### UnsetLacpduRate
`func (o *InterfaceCreate0) UnsetLacpduRate()`

UnsetLacpduRate ensures that no value is present for LacpduRate, not even an explicit nil
### GetLagPorts

`func (o *InterfaceCreate0) GetLagPorts() []string`

GetLagPorts returns the LagPorts field if non-nil, zero value otherwise.

### GetLagPortsOk

`func (o *InterfaceCreate0) GetLagPortsOk() (*[]string, bool)`

GetLagPortsOk returns a tuple with the LagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPorts

`func (o *InterfaceCreate0) SetLagPorts(v []string)`

SetLagPorts sets LagPorts field to given value.

### HasLagPorts

`func (o *InterfaceCreate0) HasLagPorts() bool`

HasLagPorts returns a boolean if a field has been set.

### GetVlanParentInterface

`func (o *InterfaceCreate0) GetVlanParentInterface() string`

GetVlanParentInterface returns the VlanParentInterface field if non-nil, zero value otherwise.

### GetVlanParentInterfaceOk

`func (o *InterfaceCreate0) GetVlanParentInterfaceOk() (*string, bool)`

GetVlanParentInterfaceOk returns a tuple with the VlanParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParentInterface

`func (o *InterfaceCreate0) SetVlanParentInterface(v string)`

SetVlanParentInterface sets VlanParentInterface field to given value.

### HasVlanParentInterface

`func (o *InterfaceCreate0) HasVlanParentInterface() bool`

HasVlanParentInterface returns a boolean if a field has been set.

### GetVlanTag

`func (o *InterfaceCreate0) GetVlanTag() int32`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *InterfaceCreate0) GetVlanTagOk() (*int32, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *InterfaceCreate0) SetVlanTag(v int32)`

SetVlanTag sets VlanTag field to given value.

### HasVlanTag

`func (o *InterfaceCreate0) HasVlanTag() bool`

HasVlanTag returns a boolean if a field has been set.

### GetVlanPcp

`func (o *InterfaceCreate0) GetVlanPcp() int32`

GetVlanPcp returns the VlanPcp field if non-nil, zero value otherwise.

### GetVlanPcpOk

`func (o *InterfaceCreate0) GetVlanPcpOk() (*int32, bool)`

GetVlanPcpOk returns a tuple with the VlanPcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPcp

`func (o *InterfaceCreate0) SetVlanPcp(v int32)`

SetVlanPcp sets VlanPcp field to given value.

### HasVlanPcp

`func (o *InterfaceCreate0) HasVlanPcp() bool`

HasVlanPcp returns a boolean if a field has been set.

### SetVlanPcpNil

`func (o *InterfaceCreate0) SetVlanPcpNil(b bool)`

 SetVlanPcpNil sets the value for VlanPcp to be an explicit nil

### UnsetVlanPcp
`func (o *InterfaceCreate0) UnsetVlanPcp()`

UnsetVlanPcp ensures that no value is present for VlanPcp, not even an explicit nil
### GetMtu

`func (o *InterfaceCreate0) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InterfaceCreate0) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InterfaceCreate0) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InterfaceCreate0) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InterfaceCreate0) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InterfaceCreate0) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


