# FailoverCallRemote2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** |  | [optional] [default to 60]
**Job** | Pointer to **bool** |  | [optional] [default to false]
**JobReturn** | Pointer to **NullableBool** |  | [optional] 
**Callback** | Pointer to **interface{}** |  | [optional] [default to null]
**ConnectTimeout** | Pointer to **float32** |  | [optional] 

## Methods

### NewFailoverCallRemote2

`func NewFailoverCallRemote2() *FailoverCallRemote2`

NewFailoverCallRemote2 instantiates a new FailoverCallRemote2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverCallRemote2WithDefaults

`func NewFailoverCallRemote2WithDefaults() *FailoverCallRemote2`

NewFailoverCallRemote2WithDefaults instantiates a new FailoverCallRemote2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *FailoverCallRemote2) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FailoverCallRemote2) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FailoverCallRemote2) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FailoverCallRemote2) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetJob

`func (o *FailoverCallRemote2) GetJob() bool`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *FailoverCallRemote2) GetJobOk() (*bool, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *FailoverCallRemote2) SetJob(v bool)`

SetJob sets Job field to given value.

### HasJob

`func (o *FailoverCallRemote2) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobReturn

`func (o *FailoverCallRemote2) GetJobReturn() bool`

GetJobReturn returns the JobReturn field if non-nil, zero value otherwise.

### GetJobReturnOk

`func (o *FailoverCallRemote2) GetJobReturnOk() (*bool, bool)`

GetJobReturnOk returns a tuple with the JobReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobReturn

`func (o *FailoverCallRemote2) SetJobReturn(v bool)`

SetJobReturn sets JobReturn field to given value.

### HasJobReturn

`func (o *FailoverCallRemote2) HasJobReturn() bool`

HasJobReturn returns a boolean if a field has been set.

### SetJobReturnNil

`func (o *FailoverCallRemote2) SetJobReturnNil(b bool)`

 SetJobReturnNil sets the value for JobReturn to be an explicit nil

### UnsetJobReturn
`func (o *FailoverCallRemote2) UnsetJobReturn()`

UnsetJobReturn ensures that no value is present for JobReturn, not even an explicit nil
### GetCallback

`func (o *FailoverCallRemote2) GetCallback() interface{}`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *FailoverCallRemote2) GetCallbackOk() (*interface{}, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *FailoverCallRemote2) SetCallback(v interface{})`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *FailoverCallRemote2) HasCallback() bool`

HasCallback returns a boolean if a field has been set.

### SetCallbackNil

`func (o *FailoverCallRemote2) SetCallbackNil(b bool)`

 SetCallbackNil sets the value for Callback to be an explicit nil

### UnsetCallback
`func (o *FailoverCallRemote2) UnsetCallback()`

UnsetCallback ensures that no value is present for Callback, not even an explicit nil
### GetConnectTimeout

`func (o *FailoverCallRemote2) GetConnectTimeout() float32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *FailoverCallRemote2) GetConnectTimeoutOk() (*float32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *FailoverCallRemote2) SetConnectTimeout(v float32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *FailoverCallRemote2) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


