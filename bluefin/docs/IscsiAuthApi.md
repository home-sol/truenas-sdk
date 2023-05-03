# \IscsiAuthApi

All URIs are relative to *https://192.168.40.2/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IscsiAuthGet**](IscsiAuthApi.md#IscsiAuthGet) | **Get** /iscsi/auth | 
[**IscsiAuthGetInstancePost**](IscsiAuthApi.md#IscsiAuthGetInstancePost) | **Post** /iscsi/auth/get_instance | 
[**IscsiAuthIdIdDelete**](IscsiAuthApi.md#IscsiAuthIdIdDelete) | **Delete** /iscsi/auth/id/{id} | 
[**IscsiAuthIdIdGet**](IscsiAuthApi.md#IscsiAuthIdIdGet) | **Get** /iscsi/auth/id/{id} | 
[**IscsiAuthIdIdPut**](IscsiAuthApi.md#IscsiAuthIdIdPut) | **Put** /iscsi/auth/id/{id} | 
[**IscsiAuthPost**](IscsiAuthApi.md#IscsiAuthPost) | **Post** /iscsi/auth | 



## IscsiAuthGet

> IscsiAuthGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    r, err := apiClient.IscsiAuthApi.IscsiAuthGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiAuthApi.IscsiAuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIscsiAuthGetRequest struct via the builder pattern


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


## IscsiAuthGetInstancePost

> IscsiAuthGetInstancePost(ctx).IscsiAuthGetInstance(iscsiAuthGetInstance).Execute()





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
    iscsiAuthGetInstance := *openapiclient.NewIscsiAuthGetInstance() // IscsiAuthGetInstance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IscsiAuthApi.IscsiAuthGetInstancePost(context.Background()).IscsiAuthGetInstance(iscsiAuthGetInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiAuthApi.IscsiAuthGetInstancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIscsiAuthGetInstancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iscsiAuthGetInstance** | [**IscsiAuthGetInstance**](IscsiAuthGetInstance.md) |  | 

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


## IscsiAuthIdIdDelete

> IscsiAuthIdIdDelete(ctx, id).Execute()





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
    r, err := apiClient.IscsiAuthApi.IscsiAuthIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiAuthApi.IscsiAuthIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIscsiAuthIdIdDeleteRequest struct via the builder pattern


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


## IscsiAuthIdIdGet

> IscsiAuthIdIdGet(ctx, id).Execute()





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
    r, err := apiClient.IscsiAuthApi.IscsiAuthIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiAuthApi.IscsiAuthIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIscsiAuthIdIdGetRequest struct via the builder pattern


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


## IscsiAuthIdIdPut

> IscsiAuthIdIdPut(ctx, id).Execute()





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
    r, err := apiClient.IscsiAuthApi.IscsiAuthIdIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiAuthApi.IscsiAuthIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIscsiAuthIdIdPutRequest struct via the builder pattern


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


## IscsiAuthPost

> IscsiAuthPost(ctx).IscsiAuthCreate0(iscsiAuthCreate0).Execute()





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
    iscsiAuthCreate0 := *openapiclient.NewIscsiAuthCreate0() // IscsiAuthCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IscsiAuthApi.IscsiAuthPost(context.Background()).IscsiAuthCreate0(iscsiAuthCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiAuthApi.IscsiAuthPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIscsiAuthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iscsiAuthCreate0** | [**IscsiAuthCreate0**](IscsiAuthCreate0.md) |  | 

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

