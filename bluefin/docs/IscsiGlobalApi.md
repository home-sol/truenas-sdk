# \IscsiGlobalApi

All URIs are relative to *https://192.168.40.2/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IscsiGlobalAluaEnabledGet**](IscsiGlobalApi.md#IscsiGlobalAluaEnabledGet) | **Get** /iscsi/global/alua_enabled | 
[**IscsiGlobalClientCountGet**](IscsiGlobalApi.md#IscsiGlobalClientCountGet) | **Get** /iscsi/global/client_count | 
[**IscsiGlobalGet**](IscsiGlobalApi.md#IscsiGlobalGet) | **Get** /iscsi/global | 
[**IscsiGlobalPut**](IscsiGlobalApi.md#IscsiGlobalPut) | **Put** /iscsi/global | 
[**IscsiGlobalSessionsGet**](IscsiGlobalApi.md#IscsiGlobalSessionsGet) | **Get** /iscsi/global/sessions | 



## IscsiGlobalAluaEnabledGet

> IscsiGlobalAluaEnabledGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IscsiGlobalApi.IscsiGlobalAluaEnabledGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiGlobalApi.IscsiGlobalAluaEnabledGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIscsiGlobalAluaEnabledGetRequest struct via the builder pattern


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


## IscsiGlobalClientCountGet

> IscsiGlobalClientCountGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IscsiGlobalApi.IscsiGlobalClientCountGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiGlobalApi.IscsiGlobalClientCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIscsiGlobalClientCountGetRequest struct via the builder pattern


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


## IscsiGlobalGet

> IscsiGlobalGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IscsiGlobalApi.IscsiGlobalGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiGlobalApi.IscsiGlobalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIscsiGlobalGetRequest struct via the builder pattern


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


## IscsiGlobalPut

> IscsiGlobalPut(ctx).IscsiGlobalUpdate0(iscsiGlobalUpdate0).Execute()





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
    iscsiGlobalUpdate0 := *openapiclient.NewIscsiGlobalUpdate0() // IscsiGlobalUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IscsiGlobalApi.IscsiGlobalPut(context.Background()).IscsiGlobalUpdate0(iscsiGlobalUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiGlobalApi.IscsiGlobalPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIscsiGlobalPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iscsiGlobalUpdate0** | [**IscsiGlobalUpdate0**](IscsiGlobalUpdate0.md) |  | 

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


## IscsiGlobalSessionsGet

> IscsiGlobalSessionsGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    r, err := apiClient.IscsiGlobalApi.IscsiGlobalSessionsGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiGlobalApi.IscsiGlobalSessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIscsiGlobalSessionsGetRequest struct via the builder pattern


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

