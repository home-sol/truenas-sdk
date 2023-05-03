# CertificateGetInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**CertificateGetInstance0**](CertificateGetInstance0.md) |  | [optional] 
**QueryOptions** | Pointer to [**CertificateGetInstance1**](CertificateGetInstance1.md) |  | [optional] [default to {}]

## Methods

### NewCertificateGetInstance

`func NewCertificateGetInstance() *CertificateGetInstance`

NewCertificateGetInstance instantiates a new CertificateGetInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateGetInstanceWithDefaults

`func NewCertificateGetInstanceWithDefaults() *CertificateGetInstance`

NewCertificateGetInstanceWithDefaults instantiates a new CertificateGetInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateGetInstance) GetId() CertificateGetInstance0`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateGetInstance) GetIdOk() (*CertificateGetInstance0, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateGetInstance) SetId(v CertificateGetInstance0)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateGetInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueryOptions

`func (o *CertificateGetInstance) GetQueryOptions() CertificateGetInstance1`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *CertificateGetInstance) GetQueryOptionsOk() (*CertificateGetInstance1, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *CertificateGetInstance) SetQueryOptions(v CertificateGetInstance1)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *CertificateGetInstance) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

