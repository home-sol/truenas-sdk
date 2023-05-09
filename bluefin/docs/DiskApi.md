# \DiskApi

All URIs are relative to *https://truenas/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisk**](DiskApi.md#GetDisk) | **Get** /disk/id/{id} | 
[**ListDisks**](DiskApi.md#ListDisks) | **Get** /disk | 
[**UpdateDisk**](DiskApi.md#UpdateDisk) | **Put** /disk/id/{id} | 



## GetDisk

> Disk GetDisk(ctx, id).Execute()



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
    id := "id_example" // string | Disk ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiskApi.GetDisk(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiskApi.GetDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDisk`: Disk
    fmt.Fprintf(os.Stdout, "Response from `DiskApi.GetDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Disk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Disk**](Disk.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDisks

> []Disk ListDisks(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).ExtraIncludeExpired(extraIncludeExpired).ExtraPasswords(extraPasswords).ExtraSupportsSmart(extraSupportsSmart).ExtraPools(extraPools).Execute()



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
    extraIncludeExpired := true // bool | will also include expired disks (optional) (default to false)
    extraPasswords := true // bool | will not hide KMIP password for the disks (optional) (default to false)
    extraSupportsSmart := true // bool | will query if disks support S.M.A.R.T. Only supported if resulting disks count is not larger than one; otherwise, raises an error. (optional) (default to false)
    extraPools := true // bool | will join pool name for each disk (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiskApi.ListDisks(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).ExtraIncludeExpired(extraIncludeExpired).ExtraPasswords(extraPasswords).ExtraSupportsSmart(extraSupportsSmart).ExtraPools(extraPools).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiskApi.ListDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDisks`: []Disk
    fmt.Fprintf(os.Stdout, "Response from `DiskApi.ListDisks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 
 **extraIncludeExpired** | **bool** | will also include expired disks | [default to false]
 **extraPasswords** | **bool** | will not hide KMIP password for the disks | [default to false]
 **extraSupportsSmart** | **bool** | will query if disks support S.M.A.R.T. Only supported if resulting disks count is not larger than one; otherwise, raises an error. | [default to false]
 **extraPools** | **bool** | will join pool name for each disk | [default to false]

### Return type

[**[]Disk**](Disk.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDisk

> Disk UpdateDisk(ctx, id).UpdateDiskRequest(updateDiskRequest).Execute()



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
    id := "id_example" // string | Disk ID
    updateDiskRequest := *openapiclient.NewUpdateDiskRequest(int32(123), "Lunid_example", "Description_example", NullableInt32(123), NullableInt32(123), NullableInt32(123), openapiclient.HDDStandby("ALWAYS ON"), openapiclient.AdvPowermgmt("DISABLED"), false) // UpdateDiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiskApi.UpdateDisk(context.Background(), id).UpdateDiskRequest(updateDiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiskApi.UpdateDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDisk`: Disk
    fmt.Fprintf(os.Stdout, "Response from `DiskApi.UpdateDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Disk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDiskRequest** | [**UpdateDiskRequest**](UpdateDiskRequest.md) |  | 

### Return type

[**Disk**](Disk.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

