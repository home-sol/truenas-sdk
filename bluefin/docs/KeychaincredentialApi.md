# \KeychaincredentialApi

All URIs are relative to *https://192.168.40.2/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeychaincredentialGenerateSshKeyPairGet**](KeychaincredentialApi.md#KeychaincredentialGenerateSshKeyPairGet) | **Get** /keychaincredential/generate_ssh_key_pair | 
[**KeychaincredentialGet**](KeychaincredentialApi.md#KeychaincredentialGet) | **Get** /keychaincredential | 
[**KeychaincredentialGetInstancePost**](KeychaincredentialApi.md#KeychaincredentialGetInstancePost) | **Post** /keychaincredential/get_instance | 
[**KeychaincredentialIdIdDelete**](KeychaincredentialApi.md#KeychaincredentialIdIdDelete) | **Delete** /keychaincredential/id/{id} | 
[**KeychaincredentialIdIdGet**](KeychaincredentialApi.md#KeychaincredentialIdIdGet) | **Get** /keychaincredential/id/{id} | 
[**KeychaincredentialIdIdPut**](KeychaincredentialApi.md#KeychaincredentialIdIdPut) | **Put** /keychaincredential/id/{id} | 
[**KeychaincredentialPost**](KeychaincredentialApi.md#KeychaincredentialPost) | **Post** /keychaincredential | 
[**KeychaincredentialRemoteSshHostKeyScanPost**](KeychaincredentialApi.md#KeychaincredentialRemoteSshHostKeyScanPost) | **Post** /keychaincredential/remote_ssh_host_key_scan | 
[**KeychaincredentialRemoteSshSemiautomaticSetupPost**](KeychaincredentialApi.md#KeychaincredentialRemoteSshSemiautomaticSetupPost) | **Post** /keychaincredential/remote_ssh_semiautomatic_setup | 
[**KeychaincredentialSetupSshConnectionPost**](KeychaincredentialApi.md#KeychaincredentialSetupSshConnectionPost) | **Post** /keychaincredential/setup_ssh_connection | 
[**KeychaincredentialUsedByPost**](KeychaincredentialApi.md#KeychaincredentialUsedByPost) | **Post** /keychaincredential/used_by | 



## KeychaincredentialGenerateSshKeyPairGet

> KeychaincredentialGenerateSshKeyPairGet(ctx).Execute()





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
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialGenerateSshKeyPairGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialGenerateSshKeyPairGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialGenerateSshKeyPairGetRequest struct via the builder pattern


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


## KeychaincredentialGet

> KeychaincredentialGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialGetRequest struct via the builder pattern


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


## KeychaincredentialGetInstancePost

> KeychaincredentialGetInstancePost(ctx).KeychaincredentialGetInstance(keychaincredentialGetInstance).Execute()





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
    keychaincredentialGetInstance := *openapiclient.NewKeychaincredentialGetInstance() // KeychaincredentialGetInstance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialGetInstancePost(context.Background()).KeychaincredentialGetInstance(keychaincredentialGetInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialGetInstancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialGetInstancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialGetInstance** | [**KeychaincredentialGetInstance**](KeychaincredentialGetInstance.md) |  | 

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


## KeychaincredentialIdIdDelete

> KeychaincredentialIdIdDelete(ctx, id).Execute()





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
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKeychaincredentialIdIdDeleteRequest struct via the builder pattern


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


## KeychaincredentialIdIdGet

> KeychaincredentialIdIdGet(ctx, id).Execute()





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
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKeychaincredentialIdIdGetRequest struct via the builder pattern


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


## KeychaincredentialIdIdPut

> KeychaincredentialIdIdPut(ctx, id).Execute()





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
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialIdIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKeychaincredentialIdIdPutRequest struct via the builder pattern


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


## KeychaincredentialPost

> KeychaincredentialPost(ctx).KeychaincredentialCreate0(keychaincredentialCreate0).Execute()





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
    keychaincredentialCreate0 := *openapiclient.NewKeychaincredentialCreate0() // KeychaincredentialCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialPost(context.Background()).KeychaincredentialCreate0(keychaincredentialCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialCreate0** | [**KeychaincredentialCreate0**](KeychaincredentialCreate0.md) |  | 

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


## KeychaincredentialRemoteSshHostKeyScanPost

> KeychaincredentialRemoteSshHostKeyScanPost(ctx).KeychaincredentialRemoteSshHostKeyScan0(keychaincredentialRemoteSshHostKeyScan0).Execute()





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
    keychaincredentialRemoteSshHostKeyScan0 := *openapiclient.NewKeychaincredentialRemoteSshHostKeyScan0() // KeychaincredentialRemoteSshHostKeyScan0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialRemoteSshHostKeyScanPost(context.Background()).KeychaincredentialRemoteSshHostKeyScan0(keychaincredentialRemoteSshHostKeyScan0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialRemoteSshHostKeyScanPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialRemoteSshHostKeyScanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialRemoteSshHostKeyScan0** | [**KeychaincredentialRemoteSshHostKeyScan0**](KeychaincredentialRemoteSshHostKeyScan0.md) |  | 

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


## KeychaincredentialRemoteSshSemiautomaticSetupPost

> KeychaincredentialRemoteSshSemiautomaticSetupPost(ctx).KeychaincredentialRemoteSshSemiautomaticSetup0(keychaincredentialRemoteSshSemiautomaticSetup0).Execute()





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
    keychaincredentialRemoteSshSemiautomaticSetup0 := *openapiclient.NewKeychaincredentialRemoteSshSemiautomaticSetup0() // KeychaincredentialRemoteSshSemiautomaticSetup0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialRemoteSshSemiautomaticSetupPost(context.Background()).KeychaincredentialRemoteSshSemiautomaticSetup0(keychaincredentialRemoteSshSemiautomaticSetup0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialRemoteSshSemiautomaticSetupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialRemoteSshSemiautomaticSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialRemoteSshSemiautomaticSetup0** | [**KeychaincredentialRemoteSshSemiautomaticSetup0**](KeychaincredentialRemoteSshSemiautomaticSetup0.md) |  | 

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


## KeychaincredentialSetupSshConnectionPost

> KeychaincredentialSetupSshConnectionPost(ctx).KeychaincredentialSetupSshConnection0(keychaincredentialSetupSshConnection0).Execute()





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
    keychaincredentialSetupSshConnection0 := *openapiclient.NewKeychaincredentialSetupSshConnection0() // KeychaincredentialSetupSshConnection0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialSetupSshConnectionPost(context.Background()).KeychaincredentialSetupSshConnection0(keychaincredentialSetupSshConnection0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialSetupSshConnectionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialSetupSshConnectionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialSetupSshConnection0** | [**KeychaincredentialSetupSshConnection0**](KeychaincredentialSetupSshConnection0.md) |  | 

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


## KeychaincredentialUsedByPost

> KeychaincredentialUsedByPost(ctx).Body(body).Execute()





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
    body := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeychaincredentialApi.KeychaincredentialUsedByPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialUsedByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialUsedByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **int32** |  | 

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

