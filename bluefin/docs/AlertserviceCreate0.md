# AlertserviceCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Create an Alert Service of specified &#x60;type&#x60;. If &#x60;enabled&#x60;, it sends alerts to the configured &#x60;type&#x60; of Alert Service. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Level** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | If &#x60;enabled&#x60;, it sends alerts to the configured &#x60;type&#x60; of Alert Service. | [optional] [default to true]

## Methods

### NewAlertserviceCreate0

`func NewAlertserviceCreate0() *AlertserviceCreate0`

NewAlertserviceCreate0 instantiates a new AlertserviceCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertserviceCreate0WithDefaults

`func NewAlertserviceCreate0WithDefaults() *AlertserviceCreate0`

NewAlertserviceCreate0WithDefaults instantiates a new AlertserviceCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertserviceCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertserviceCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertserviceCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertserviceCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AlertserviceCreate0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertserviceCreate0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertserviceCreate0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertserviceCreate0) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *AlertserviceCreate0) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AlertserviceCreate0) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AlertserviceCreate0) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AlertserviceCreate0) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLevel

`func (o *AlertserviceCreate0) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AlertserviceCreate0) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AlertserviceCreate0) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AlertserviceCreate0) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertserviceCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertserviceCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertserviceCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertserviceCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


