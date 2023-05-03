# \AcmeDnsAuthenticatorApi

All URIs are relative to *https://192.168.40.2/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcmeDnsAuthenticatorAuthenticatorSchemasGet**](AcmeDnsAuthenticatorApi.md#AcmeDnsAuthenticatorAuthenticatorSchemasGet) | **Get** /acme/dns/authenticator/authenticator_schemas | 
[**AcmeDnsAuthenticatorGet**](AcmeDnsAuthenticatorApi.md#AcmeDnsAuthenticatorGet) | **Get** /acme/dns/authenticator | 
[**AcmeDnsAuthenticatorGetInstancePost**](AcmeDnsAuthenticatorApi.md#AcmeDnsAuthenticatorGetInstancePost) | **Post** /acme/dns/authenticator/get_instance | 
[**AcmeDnsAuthenticatorIdIdDelete**](AcmeDnsAuthenticatorApi.md#AcmeDnsAuthenticatorIdIdDelete) | **Delete** /acme/dns/authenticator/id/{id} | 
[**AcmeDnsAuthenticatorIdIdGet**](AcmeDnsAuthenticatorApi.md#AcmeDnsAuthenticatorIdIdGet) | **Get** /acme/dns/authenticator/id/{id} | 
[**AcmeDnsAuthenticatorIdIdPut**](AcmeDnsAuthenticatorApi.md#AcmeDnsAuthenticatorIdIdPut) | **Put** /acme/dns/authenticator/id/{id} | 
[**AcmeDnsAuthenticatorPost**](AcmeDnsAuthenticatorApi.md#AcmeDnsAuthenticatorPost) | **Post** /acme/dns/authenticator | 



## AcmeDnsAuthenticatorAuthenticatorSchemasGet

> AcmeDnsAuthenticatorAuthenticatorSchemasGet(ctx).Execute()





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
    r, err := apiClient.AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorAuthenticatorSchemasGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorAuthenticatorSchemasGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeDnsAuthenticatorAuthenticatorSchemasGetRequest struct via the builder pattern


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


## AcmeDnsAuthenticatorGet

> AcmeDnsAuthenticatorGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    r, err := apiClient.AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeDnsAuthenticatorGetRequest struct via the builder pattern


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


## AcmeDnsAuthenticatorGetInstancePost

> AcmeDnsAuthenticatorGetInstancePost(ctx).AcmeDnsAuthenticatorGetInstance(acmeDnsAuthenticatorGetInstance).Execute()





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
    acmeDnsAuthenticatorGetInstance := *openapiclient.NewAcmeDnsAuthenticatorGetInstance() // AcmeDnsAuthenticatorGetInstance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorGetInstancePost(context.Background()).AcmeDnsAuthenticatorGetInstance(acmeDnsAuthenticatorGetInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorGetInstancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeDnsAuthenticatorGetInstancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acmeDnsAuthenticatorGetInstance** | [**AcmeDnsAuthenticatorGetInstance**](AcmeDnsAuthenticatorGetInstance.md) |  | 

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


## AcmeDnsAuthenticatorIdIdDelete

> AcmeDnsAuthenticatorIdIdDelete(ctx, id).Execute()





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
    r, err := apiClient.AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAcmeDnsAuthenticatorIdIdDeleteRequest struct via the builder pattern


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


## AcmeDnsAuthenticatorIdIdGet

> AcmeDnsAuthenticatorIdIdGet(ctx, id).Execute()





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
    r, err := apiClient.AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAcmeDnsAuthenticatorIdIdGetRequest struct via the builder pattern


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


## AcmeDnsAuthenticatorIdIdPut

> AcmeDnsAuthenticatorIdIdPut(ctx, id).Execute()





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
    r, err := apiClient.AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorIdIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAcmeDnsAuthenticatorIdIdPutRequest struct via the builder pattern


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


## AcmeDnsAuthenticatorPost

> AcmeDnsAuthenticatorPost(ctx).AcmeDnsAuthenticatorCreate0(acmeDnsAuthenticatorCreate0).Execute()





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
    acmeDnsAuthenticatorCreate0 := *openapiclient.NewAcmeDnsAuthenticatorCreate0() // AcmeDnsAuthenticatorCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorPost(context.Background()).AcmeDnsAuthenticatorCreate0(acmeDnsAuthenticatorCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcmeDnsAuthenticatorApi.AcmeDnsAuthenticatorPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeDnsAuthenticatorPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acmeDnsAuthenticatorCreate0** | [**AcmeDnsAuthenticatorCreate0**](AcmeDnsAuthenticatorCreate0.md) |  | 

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

