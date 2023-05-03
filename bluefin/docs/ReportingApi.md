# \ReportingApi

All URIs are relative to *https://192.168.40.2/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportingClearGet**](ReportingApi.md#ReportingClearGet) | **Get** /reporting/clear | 
[**ReportingGet**](ReportingApi.md#ReportingGet) | **Get** /reporting | 
[**ReportingGetDataPost**](ReportingApi.md#ReportingGetDataPost) | **Post** /reporting/get_data | 
[**ReportingGraphsGet**](ReportingApi.md#ReportingGraphsGet) | **Get** /reporting/graphs | 
[**ReportingPut**](ReportingApi.md#ReportingPut) | **Put** /reporting | 



## ReportingClearGet

> ReportingClearGet(ctx).Execute()





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
    r, err := apiClient.ReportingApi.ReportingClearGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.ReportingClearGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReportingClearGetRequest struct via the builder pattern


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


## ReportingGet

> ReportingGet(ctx).Execute()





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
    r, err := apiClient.ReportingApi.ReportingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.ReportingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReportingGetRequest struct via the builder pattern


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


## ReportingGetDataPost

> ReportingGetDataPost(ctx).ReportingGetData(reportingGetData).Execute()





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
    reportingGetData := *openapiclient.NewReportingGetData() // ReportingGetData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReportingApi.ReportingGetDataPost(context.Background()).ReportingGetData(reportingGetData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.ReportingGetDataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportingGetDataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportingGetData** | [**ReportingGetData**](ReportingGetData.md) |  | 

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


## ReportingGraphsGet

> ReportingGraphsGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    r, err := apiClient.ReportingApi.ReportingGraphsGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.ReportingGraphsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportingGraphsGetRequest struct via the builder pattern


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


## ReportingPut

> ReportingPut(ctx).ReportingUpdate0(reportingUpdate0).Execute()





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
    reportingUpdate0 := *openapiclient.NewReportingUpdate0() // ReportingUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReportingApi.ReportingPut(context.Background()).ReportingUpdate0(reportingUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.ReportingPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportingUpdate0** | [**ReportingUpdate0**](ReportingUpdate0.md) |  | 

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

