# \KerberosRealmApi

All URIs are relative to *https://192.168.40.2/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KerberosRealmGet**](KerberosRealmApi.md#KerberosRealmGet) | **Get** /kerberos/realm | 
[**KerberosRealmGetInstancePost**](KerberosRealmApi.md#KerberosRealmGetInstancePost) | **Post** /kerberos/realm/get_instance | 
[**KerberosRealmIdIdDelete**](KerberosRealmApi.md#KerberosRealmIdIdDelete) | **Delete** /kerberos/realm/id/{id} | 
[**KerberosRealmIdIdGet**](KerberosRealmApi.md#KerberosRealmIdIdGet) | **Get** /kerberos/realm/id/{id} | 
[**KerberosRealmIdIdPut**](KerberosRealmApi.md#KerberosRealmIdIdPut) | **Put** /kerberos/realm/id/{id} | 
[**KerberosRealmPost**](KerberosRealmApi.md#KerberosRealmPost) | **Post** /kerberos/realm | 



## KerberosRealmGet

> KerberosRealmGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/home-sol/truenas-sdk/bluefin"
)

func main() {
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KerberosRealmApi.KerberosRealmGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmApi.KerberosRealmGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKerberosRealmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosRealmGetInstancePost

> KerberosRealmGetInstancePost(ctx).KerberosRealmGetInstance(kerberosRealmGetInstance).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/home-sol/truenas-sdk/bluefin"
)

func main() {
    kerberosRealmGetInstance := *openapiclient.NewKerberosRealmGetInstance() // KerberosRealmGetInstance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KerberosRealmApi.KerberosRealmGetInstancePost(context.Background()).KerberosRealmGetInstance(kerberosRealmGetInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmApi.KerberosRealmGetInstancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKerberosRealmGetInstancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosRealmGetInstance** | [**KerberosRealmGetInstance**](KerberosRealmGetInstance.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosRealmIdIdDelete

> KerberosRealmIdIdDelete(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/home-sol/truenas-sdk/bluefin"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KerberosRealmApi.KerberosRealmIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmApi.KerberosRealmIdIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKerberosRealmIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosRealmIdIdGet

> KerberosRealmIdIdGet(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/home-sol/truenas-sdk/bluefin"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KerberosRealmApi.KerberosRealmIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmApi.KerberosRealmIdIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKerberosRealmIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosRealmIdIdPut

> KerberosRealmIdIdPut(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/home-sol/truenas-sdk/bluefin"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KerberosRealmApi.KerberosRealmIdIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmApi.KerberosRealmIdIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKerberosRealmIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosRealmPost

> KerberosRealmPost(ctx).KerberosRealmCreate0(kerberosRealmCreate0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/home-sol/truenas-sdk/bluefin"
)

func main() {
    kerberosRealmCreate0 := *openapiclient.NewKerberosRealmCreate0() // KerberosRealmCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KerberosRealmApi.KerberosRealmPost(context.Background()).KerberosRealmCreate0(kerberosRealmCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmApi.KerberosRealmPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKerberosRealmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosRealmCreate0** | [**KerberosRealmCreate0**](KerberosRealmCreate0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

