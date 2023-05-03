# MailUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fromemail** | Pointer to **string** | &#x60;fromemail&#x60; is used as a sending address which the mail server will use for sending emails. | [optional] 
**Fromname** | Pointer to **string** |  | [optional] 
**Outgoingserver** | Pointer to **string** | &#x60;outgoingserver&#x60; is the hostname or IP address of SMTP server used for sending an email. | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Security** | Pointer to **string** | &#x60;security&#x60; is type of encryption desired. | [optional] 
**Smtp** | Pointer to **bool** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 
**Pass** | Pointer to **NullableString** |  | [optional] 
**Oauth** | Pointer to [**Oauth**](Oauth.md) |  | [optional] [default to {}]

## Methods

### NewMailUpdate0

`func NewMailUpdate0() *MailUpdate0`

NewMailUpdate0 instantiates a new MailUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUpdate0WithDefaults

`func NewMailUpdate0WithDefaults() *MailUpdate0`

NewMailUpdate0WithDefaults instantiates a new MailUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromemail

`func (o *MailUpdate0) GetFromemail() string`

GetFromemail returns the Fromemail field if non-nil, zero value otherwise.

### GetFromemailOk

`func (o *MailUpdate0) GetFromemailOk() (*string, bool)`

GetFromemailOk returns a tuple with the Fromemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromemail

`func (o *MailUpdate0) SetFromemail(v string)`

SetFromemail sets Fromemail field to given value.

### HasFromemail

`func (o *MailUpdate0) HasFromemail() bool`

HasFromemail returns a boolean if a field has been set.

### GetFromname

`func (o *MailUpdate0) GetFromname() string`

GetFromname returns the Fromname field if non-nil, zero value otherwise.

### GetFromnameOk

`func (o *MailUpdate0) GetFromnameOk() (*string, bool)`

GetFromnameOk returns a tuple with the Fromname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromname

`func (o *MailUpdate0) SetFromname(v string)`

SetFromname sets Fromname field to given value.

### HasFromname

`func (o *MailUpdate0) HasFromname() bool`

HasFromname returns a boolean if a field has been set.

### GetOutgoingserver

`func (o *MailUpdate0) GetOutgoingserver() string`

GetOutgoingserver returns the Outgoingserver field if non-nil, zero value otherwise.

### GetOutgoingserverOk

`func (o *MailUpdate0) GetOutgoingserverOk() (*string, bool)`

GetOutgoingserverOk returns a tuple with the Outgoingserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingserver

`func (o *MailUpdate0) SetOutgoingserver(v string)`

SetOutgoingserver sets Outgoingserver field to given value.

### HasOutgoingserver

`func (o *MailUpdate0) HasOutgoingserver() bool`

HasOutgoingserver returns a boolean if a field has been set.

### GetPort

`func (o *MailUpdate0) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MailUpdate0) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MailUpdate0) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MailUpdate0) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *MailUpdate0) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *MailUpdate0) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *MailUpdate0) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *MailUpdate0) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSmtp

`func (o *MailUpdate0) GetSmtp() bool`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *MailUpdate0) GetSmtpOk() (*bool, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *MailUpdate0) SetSmtp(v bool)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *MailUpdate0) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetUser

`func (o *MailUpdate0) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MailUpdate0) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MailUpdate0) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MailUpdate0) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MailUpdate0) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MailUpdate0) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPass

`func (o *MailUpdate0) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *MailUpdate0) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *MailUpdate0) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *MailUpdate0) HasPass() bool`

HasPass returns a boolean if a field has been set.

### SetPassNil

`func (o *MailUpdate0) SetPassNil(b bool)`

 SetPassNil sets the value for Pass to be an explicit nil

### UnsetPass
`func (o *MailUpdate0) UnsetPass()`

UnsetPass ensures that no value is present for Pass, not even an explicit nil
### GetOauth

`func (o *MailUpdate0) GetOauth() Oauth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *MailUpdate0) GetOauthOk() (*Oauth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *MailUpdate0) SetOauth(v Oauth)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *MailUpdate0) HasOauth() bool`

HasOauth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


