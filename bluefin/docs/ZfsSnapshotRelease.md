# ZfsSnapshotRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Release held snapshot &#x60;id&#x60;. | [optional] 
**Options** | Pointer to [**ZfsSnapshotRelease1**](ZfsSnapshotRelease1.md) |  | [optional] [default to {}]

## Methods

### NewZfsSnapshotRelease

`func NewZfsSnapshotRelease() *ZfsSnapshotRelease`

NewZfsSnapshotRelease instantiates a new ZfsSnapshotRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZfsSnapshotReleaseWithDefaults

`func NewZfsSnapshotReleaseWithDefaults() *ZfsSnapshotRelease`

NewZfsSnapshotReleaseWithDefaults instantiates a new ZfsSnapshotRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZfsSnapshotRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZfsSnapshotRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZfsSnapshotRelease) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZfsSnapshotRelease) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptions

`func (o *ZfsSnapshotRelease) GetOptions() ZfsSnapshotRelease1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ZfsSnapshotRelease) GetOptionsOk() (*ZfsSnapshotRelease1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ZfsSnapshotRelease) SetOptions(v ZfsSnapshotRelease1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ZfsSnapshotRelease) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


