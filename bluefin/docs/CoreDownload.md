# CoreDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Filename** | Pointer to **string** |  | [optional] 
**Buffered** | Pointer to **bool** | Non-&#x60;buffered&#x60; downloads will allow job to write to pipe as soon as download URL is requested, job will stay blocked meanwhile. &#x60;buffered&#x60; downloads must wait for job to complete before requesting download URL, job&#39;s pipe output will be buffered to ramfs. | [optional] [default to false]

## Methods

### NewCoreDownload

`func NewCoreDownload() *CoreDownload`

NewCoreDownload instantiates a new CoreDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDownloadWithDefaults

`func NewCoreDownloadWithDefaults() *CoreDownload`

NewCoreDownloadWithDefaults instantiates a new CoreDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *CoreDownload) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CoreDownload) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CoreDownload) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CoreDownload) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetArgs

`func (o *CoreDownload) GetArgs() []interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CoreDownload) GetArgsOk() (*[]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CoreDownload) SetArgs(v []interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CoreDownload) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetFilename

`func (o *CoreDownload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CoreDownload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CoreDownload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CoreDownload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetBuffered

`func (o *CoreDownload) GetBuffered() bool`

GetBuffered returns the Buffered field if non-nil, zero value otherwise.

### GetBufferedOk

`func (o *CoreDownload) GetBufferedOk() (*bool, bool)`

GetBufferedOk returns a tuple with the Buffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuffered

`func (o *CoreDownload) SetBuffered(v bool)`

SetBuffered sets Buffered field to given value.

### HasBuffered

`func (o *CoreDownload) HasBuffered() bool`

HasBuffered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


