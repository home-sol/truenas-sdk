# SupportNewTicket0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**AttachDebug** | Pointer to **bool** |  | [optional] [default to false]
**Token** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Criticality** | Pointer to **string** | For SCALE &#x60;criticality&#x60;, &#x60;environment&#x60;, &#x60;phone&#x60;, &#x60;name&#x60; and &#x60;email&#x60; attributes are not required. | [optional] 
**Environment** | Pointer to **string** | For SCALE &#x60;criticality&#x60;, &#x60;environment&#x60;, &#x60;phone&#x60;, &#x60;name&#x60; and &#x60;email&#x60; attributes are not required. | [optional] 
**Phone** | Pointer to **string** | For SCALE &#x60;criticality&#x60;, &#x60;environment&#x60;, &#x60;phone&#x60;, &#x60;name&#x60; and &#x60;email&#x60; attributes are not required. | [optional] 
**Name** | Pointer to **string** | For SCALE &#x60;criticality&#x60;, &#x60;environment&#x60;, &#x60;phone&#x60;, &#x60;name&#x60; and &#x60;email&#x60; attributes are not required. | [optional] 
**Email** | Pointer to **string** | For SCALE &#x60;criticality&#x60;, &#x60;environment&#x60;, &#x60;phone&#x60;, &#x60;name&#x60; and &#x60;email&#x60; attributes are not required. | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewSupportNewTicket0

`func NewSupportNewTicket0() *SupportNewTicket0`

NewSupportNewTicket0 instantiates a new SupportNewTicket0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportNewTicket0WithDefaults

`func NewSupportNewTicket0WithDefaults() *SupportNewTicket0`

NewSupportNewTicket0WithDefaults instantiates a new SupportNewTicket0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SupportNewTicket0) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportNewTicket0) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportNewTicket0) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportNewTicket0) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *SupportNewTicket0) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SupportNewTicket0) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SupportNewTicket0) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SupportNewTicket0) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *SupportNewTicket0) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SupportNewTicket0) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SupportNewTicket0) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SupportNewTicket0) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAttachDebug

`func (o *SupportNewTicket0) GetAttachDebug() bool`

GetAttachDebug returns the AttachDebug field if non-nil, zero value otherwise.

### GetAttachDebugOk

`func (o *SupportNewTicket0) GetAttachDebugOk() (*bool, bool)`

GetAttachDebugOk returns a tuple with the AttachDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachDebug

`func (o *SupportNewTicket0) SetAttachDebug(v bool)`

SetAttachDebug sets AttachDebug field to given value.

### HasAttachDebug

`func (o *SupportNewTicket0) HasAttachDebug() bool`

HasAttachDebug returns a boolean if a field has been set.

### GetToken

`func (o *SupportNewTicket0) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SupportNewTicket0) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SupportNewTicket0) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SupportNewTicket0) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *SupportNewTicket0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupportNewTicket0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupportNewTicket0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SupportNewTicket0) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCriticality

`func (o *SupportNewTicket0) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *SupportNewTicket0) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *SupportNewTicket0) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *SupportNewTicket0) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetEnvironment

`func (o *SupportNewTicket0) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SupportNewTicket0) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SupportNewTicket0) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SupportNewTicket0) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetPhone

`func (o *SupportNewTicket0) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SupportNewTicket0) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SupportNewTicket0) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SupportNewTicket0) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetName

`func (o *SupportNewTicket0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportNewTicket0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportNewTicket0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupportNewTicket0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *SupportNewTicket0) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupportNewTicket0) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupportNewTicket0) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SupportNewTicket0) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCc

`func (o *SupportNewTicket0) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *SupportNewTicket0) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *SupportNewTicket0) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *SupportNewTicket0) HasCc() bool`

HasCc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


