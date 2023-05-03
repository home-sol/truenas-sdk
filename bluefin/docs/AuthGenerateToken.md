# AuthGenerateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to **NullableInt32** | &#x60;ttl&#x60; stands for Time To Live, in seconds. The token will be invalidated if the connection has been inactive for a time greater than this. | [optional] [default to 600]
**Attrs** | Pointer to **map[string]interface{}** | &#x60;attrs&#x60; is a general purpose object/dictionary to hold information about the token. | [optional] [default to {}]
**MatchOrigin** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuthGenerateToken

`func NewAuthGenerateToken() *AuthGenerateToken`

NewAuthGenerateToken instantiates a new AuthGenerateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGenerateTokenWithDefaults

`func NewAuthGenerateTokenWithDefaults() *AuthGenerateToken`

NewAuthGenerateTokenWithDefaults instantiates a new AuthGenerateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *AuthGenerateToken) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AuthGenerateToken) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AuthGenerateToken) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *AuthGenerateToken) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *AuthGenerateToken) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *AuthGenerateToken) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetAttrs

`func (o *AuthGenerateToken) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AuthGenerateToken) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AuthGenerateToken) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AuthGenerateToken) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetMatchOrigin

`func (o *AuthGenerateToken) GetMatchOrigin() bool`

GetMatchOrigin returns the MatchOrigin field if non-nil, zero value otherwise.

### GetMatchOriginOk

`func (o *AuthGenerateToken) GetMatchOriginOk() (*bool, bool)`

GetMatchOriginOk returns a tuple with the MatchOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOrigin

`func (o *AuthGenerateToken) SetMatchOrigin(v bool)`

SetMatchOrigin sets MatchOrigin field to given value.

### HasMatchOrigin

`func (o *AuthGenerateToken) HasMatchOrigin() bool`

HasMatchOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


