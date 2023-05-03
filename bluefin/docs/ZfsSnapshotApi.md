# \ZfsSnapshotApi

All URIs are relative to *https://192.168.40.2/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZfsSnapshotClonePost**](ZfsSnapshotApi.md#ZfsSnapshotClonePost) | **Post** /zfs/snapshot/clone | 
[**ZfsSnapshotGet**](ZfsSnapshotApi.md#ZfsSnapshotGet) | **Get** /zfs/snapshot | 
[**ZfsSnapshotGetInstancePost**](ZfsSnapshotApi.md#ZfsSnapshotGetInstancePost) | **Post** /zfs/snapshot/get_instance | 
[**ZfsSnapshotHoldPost**](ZfsSnapshotApi.md#ZfsSnapshotHoldPost) | **Post** /zfs/snapshot/hold | 
[**ZfsSnapshotIdIdDelete**](ZfsSnapshotApi.md#ZfsSnapshotIdIdDelete) | **Delete** /zfs/snapshot/id/{id} | 
[**ZfsSnapshotIdIdGet**](ZfsSnapshotApi.md#ZfsSnapshotIdIdGet) | **Get** /zfs/snapshot/id/{id} | 
[**ZfsSnapshotIdIdPut**](ZfsSnapshotApi.md#ZfsSnapshotIdIdPut) | **Put** /zfs/snapshot/id/{id} | 
[**ZfsSnapshotPost**](ZfsSnapshotApi.md#ZfsSnapshotPost) | **Post** /zfs/snapshot | 
[**ZfsSnapshotReleasePost**](ZfsSnapshotApi.md#ZfsSnapshotReleasePost) | **Post** /zfs/snapshot/release | 
[**ZfsSnapshotRemovePost**](ZfsSnapshotApi.md#ZfsSnapshotRemovePost) | **Post** /zfs/snapshot/remove | 
[**ZfsSnapshotRollbackPost**](ZfsSnapshotApi.md#ZfsSnapshotRollbackPost) | **Post** /zfs/snapshot/rollback | 



## ZfsSnapshotClonePost

> ZfsSnapshotClonePost(ctx).ZfsSnapshotClone0(zfsSnapshotClone0).Execute()





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
    zfsSnapshotClone0 := *openapiclient.NewZfsSnapshotClone0() // ZfsSnapshotClone0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotClonePost(context.Background()).ZfsSnapshotClone0(zfsSnapshotClone0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotClonePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotClonePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotClone0** | [**ZfsSnapshotClone0**](ZfsSnapshotClone0.md) |  | 

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


## ZfsSnapshotGet

> ZfsSnapshotGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotGetRequest struct via the builder pattern


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


## ZfsSnapshotGetInstancePost

> ZfsSnapshotGetInstancePost(ctx).ZfsSnapshotGetInstance(zfsSnapshotGetInstance).Execute()





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
    zfsSnapshotGetInstance := *openapiclient.NewZfsSnapshotGetInstance() // ZfsSnapshotGetInstance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotGetInstancePost(context.Background()).ZfsSnapshotGetInstance(zfsSnapshotGetInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotGetInstancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotGetInstancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotGetInstance** | [**ZfsSnapshotGetInstance**](ZfsSnapshotGetInstance.md) |  | 

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


## ZfsSnapshotHoldPost

> ZfsSnapshotHoldPost(ctx).ZfsSnapshotHold(zfsSnapshotHold).Execute()





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
    zfsSnapshotHold := *openapiclient.NewZfsSnapshotHold() // ZfsSnapshotHold |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotHoldPost(context.Background()).ZfsSnapshotHold(zfsSnapshotHold).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotHoldPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotHoldPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotHold** | [**ZfsSnapshotHold**](ZfsSnapshotHold.md) |  | 

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


## ZfsSnapshotIdIdDelete

> ZfsSnapshotIdIdDelete(ctx, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotIdIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotIdIdDeleteRequest struct via the builder pattern


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


## ZfsSnapshotIdIdGet

> ZfsSnapshotIdIdGet(ctx, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotIdIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotIdIdGetRequest struct via the builder pattern


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


## ZfsSnapshotIdIdPut

> ZfsSnapshotIdIdPut(ctx, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotIdIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotIdIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotIdIdPutRequest struct via the builder pattern


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


## ZfsSnapshotPost

> ZfsSnapshotPost(ctx).ZfsSnapshotCreate0(zfsSnapshotCreate0).Execute()





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
    zfsSnapshotCreate0 := *openapiclient.NewZfsSnapshotCreate0() // ZfsSnapshotCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotPost(context.Background()).ZfsSnapshotCreate0(zfsSnapshotCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotCreate0** | [**ZfsSnapshotCreate0**](ZfsSnapshotCreate0.md) |  | 

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


## ZfsSnapshotReleasePost

> ZfsSnapshotReleasePost(ctx).ZfsSnapshotRelease(zfsSnapshotRelease).Execute()





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
    zfsSnapshotRelease := *openapiclient.NewZfsSnapshotRelease() // ZfsSnapshotRelease |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotReleasePost(context.Background()).ZfsSnapshotRelease(zfsSnapshotRelease).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotReleasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotReleasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotRelease** | [**ZfsSnapshotRelease**](ZfsSnapshotRelease.md) |  | 

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


## ZfsSnapshotRemovePost

> ZfsSnapshotRemovePost(ctx).ZfsSnapshotRemove0(zfsSnapshotRemove0).Execute()





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
    zfsSnapshotRemove0 := *openapiclient.NewZfsSnapshotRemove0() // ZfsSnapshotRemove0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotRemovePost(context.Background()).ZfsSnapshotRemove0(zfsSnapshotRemove0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotRemovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotRemove0** | [**ZfsSnapshotRemove0**](ZfsSnapshotRemove0.md) |  | 

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


## ZfsSnapshotRollbackPost

> ZfsSnapshotRollbackPost(ctx).ZfsSnapshotRollback(zfsSnapshotRollback).Execute()





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
    zfsSnapshotRollback := *openapiclient.NewZfsSnapshotRollback() // ZfsSnapshotRollback |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZfsSnapshotApi.ZfsSnapshotRollbackPost(context.Background()).ZfsSnapshotRollback(zfsSnapshotRollback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotRollbackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotRollbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotRollback** | [**ZfsSnapshotRollback**](ZfsSnapshotRollback.md) |  | 

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

