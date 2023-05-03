/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RsyncmodApiService RsyncmodApi service
type RsyncmodApiService service

type ApiRsyncmodGetRequest struct {
	ctx        context.Context
	ApiService *RsyncmodApiService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiRsyncmodGetRequest) Limit(limit int32) ApiRsyncmodGetRequest {
	r.limit = &limit
	return r
}

func (r ApiRsyncmodGetRequest) Offset(offset int32) ApiRsyncmodGetRequest {
	r.offset = &offset
	return r
}

func (r ApiRsyncmodGetRequest) Count(count bool) ApiRsyncmodGetRequest {
	r.count = &count
	return r
}

func (r ApiRsyncmodGetRequest) Sort(sort string) ApiRsyncmodGetRequest {
	r.sort = &sort
	return r
}

func (r ApiRsyncmodGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.RsyncmodGetExecute(r)
}

/*
RsyncmodGet Method for RsyncmodGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRsyncmodGetRequest
*/
func (a *RsyncmodApiService) RsyncmodGet(ctx context.Context) ApiRsyncmodGetRequest {
	return ApiRsyncmodGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *RsyncmodApiService) RsyncmodGetExecute(r ApiRsyncmodGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RsyncmodApiService.RsyncmodGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rsyncmod"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRsyncmodGetInstancePostRequest struct {
	ctx                 context.Context
	ApiService          *RsyncmodApiService
	rsyncmodGetInstance *RsyncmodGetInstance
}

func (r ApiRsyncmodGetInstancePostRequest) RsyncmodGetInstance(rsyncmodGetInstance RsyncmodGetInstance) ApiRsyncmodGetInstancePostRequest {
	r.rsyncmodGetInstance = &rsyncmodGetInstance
	return r
}

func (r ApiRsyncmodGetInstancePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.RsyncmodGetInstancePostExecute(r)
}

/*
RsyncmodGetInstancePost Method for RsyncmodGetInstancePost

Returns instance matching `id`. If `id` is not found, Validation error is raised.

Please see `query` method documentation for `options`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRsyncmodGetInstancePostRequest
*/
func (a *RsyncmodApiService) RsyncmodGetInstancePost(ctx context.Context) ApiRsyncmodGetInstancePostRequest {
	return ApiRsyncmodGetInstancePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *RsyncmodApiService) RsyncmodGetInstancePostExecute(r ApiRsyncmodGetInstancePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RsyncmodApiService.RsyncmodGetInstancePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rsyncmod/get_instance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.rsyncmodGetInstance
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRsyncmodIdIdDeleteRequest struct {
	ctx        context.Context
	ApiService *RsyncmodApiService
	id         int32
}

func (r ApiRsyncmodIdIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.RsyncmodIdIdDeleteExecute(r)
}

/*
RsyncmodIdIdDelete Method for RsyncmodIdIdDelete

Delete Rsyncmod module of `id`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiRsyncmodIdIdDeleteRequest
*/
func (a *RsyncmodApiService) RsyncmodIdIdDelete(ctx context.Context, id int32) ApiRsyncmodIdIdDeleteRequest {
	return ApiRsyncmodIdIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *RsyncmodApiService) RsyncmodIdIdDeleteExecute(r ApiRsyncmodIdIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RsyncmodApiService.RsyncmodIdIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rsyncmod/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRsyncmodIdIdGetRequest struct {
	ctx        context.Context
	ApiService *RsyncmodApiService
	id         int32
}

func (r ApiRsyncmodIdIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.RsyncmodIdIdGetExecute(r)
}

/*
RsyncmodIdIdGet Method for RsyncmodIdIdGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiRsyncmodIdIdGetRequest
*/
func (a *RsyncmodApiService) RsyncmodIdIdGet(ctx context.Context, id int32) ApiRsyncmodIdIdGetRequest {
	return ApiRsyncmodIdIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *RsyncmodApiService) RsyncmodIdIdGetExecute(r ApiRsyncmodIdIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RsyncmodApiService.RsyncmodIdIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rsyncmod/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRsyncmodIdIdPutRequest struct {
	ctx        context.Context
	ApiService *RsyncmodApiService
	id         int32
}

func (r ApiRsyncmodIdIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.RsyncmodIdIdPutExecute(r)
}

/*
RsyncmodIdIdPut Method for RsyncmodIdIdPut

Update Rsyncmod module of `id`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiRsyncmodIdIdPutRequest
*/
func (a *RsyncmodApiService) RsyncmodIdIdPut(ctx context.Context, id int32) ApiRsyncmodIdIdPutRequest {
	return ApiRsyncmodIdIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *RsyncmodApiService) RsyncmodIdIdPutExecute(r ApiRsyncmodIdIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RsyncmodApiService.RsyncmodIdIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rsyncmod/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRsyncmodPostRequest struct {
	ctx             context.Context
	ApiService      *RsyncmodApiService
	rsyncmodCreate0 *RsyncmodCreate0
}

func (r ApiRsyncmodPostRequest) RsyncmodCreate0(rsyncmodCreate0 RsyncmodCreate0) ApiRsyncmodPostRequest {
	r.rsyncmodCreate0 = &rsyncmodCreate0
	return r
}

func (r ApiRsyncmodPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.RsyncmodPostExecute(r)
}

/*
RsyncmodPost Method for RsyncmodPost

Create a Rsyncmod module.

`path` represents the path to a dataset. Path length is limited to 1023 characters maximum as per the limit
enforced by FreeBSD. It is possible that we reach this max length recursively while transferring data. In that
case, the user must ensure the maximum path will not be too long or modify the recursed path to shorter
than the limit.

`maxconn` is an integer value representing the maximum number of simultaneous connections. Zero represents
unlimited.

`hostsallow` is a list of patterns to match hostname/ip address of a connecting client. If list is empty,
all hosts are allowed.

`hostsdeny` is a list of patterns to match hostname/ip address of a connecting client. If the pattern is
matched, access is denied to the client. If no client should be denied, this should be left empty.

`auxiliary` attribute can be used to pass on any additional parameters from rsyncd.conf(5).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRsyncmodPostRequest
*/
func (a *RsyncmodApiService) RsyncmodPost(ctx context.Context) ApiRsyncmodPostRequest {
	return ApiRsyncmodPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *RsyncmodApiService) RsyncmodPostExecute(r ApiRsyncmodPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RsyncmodApiService.RsyncmodPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rsyncmod"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.rsyncmodCreate0
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
