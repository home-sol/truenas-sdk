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

// IscsiTargetApiService IscsiTargetApi service
type IscsiTargetApiService service

type ApiIscsiTargetGetRequest struct {
	ctx        context.Context
	ApiService *IscsiTargetApiService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiIscsiTargetGetRequest) Limit(limit int32) ApiIscsiTargetGetRequest {
	r.limit = &limit
	return r
}

func (r ApiIscsiTargetGetRequest) Offset(offset int32) ApiIscsiTargetGetRequest {
	r.offset = &offset
	return r
}

func (r ApiIscsiTargetGetRequest) Count(count bool) ApiIscsiTargetGetRequest {
	r.count = &count
	return r
}

func (r ApiIscsiTargetGetRequest) Sort(sort string) ApiIscsiTargetGetRequest {
	r.sort = &sort
	return r
}

func (r ApiIscsiTargetGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.IscsiTargetGetExecute(r)
}

/*
IscsiTargetGet Method for IscsiTargetGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIscsiTargetGetRequest
*/
func (a *IscsiTargetApiService) IscsiTargetGet(ctx context.Context) ApiIscsiTargetGetRequest {
	return ApiIscsiTargetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IscsiTargetApiService) IscsiTargetGetExecute(r ApiIscsiTargetGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetApiService.IscsiTargetGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/target"

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

type ApiIscsiTargetGetInstancePostRequest struct {
	ctx                    context.Context
	ApiService             *IscsiTargetApiService
	iscsiTargetGetInstance *IscsiTargetGetInstance
}

func (r ApiIscsiTargetGetInstancePostRequest) IscsiTargetGetInstance(iscsiTargetGetInstance IscsiTargetGetInstance) ApiIscsiTargetGetInstancePostRequest {
	r.iscsiTargetGetInstance = &iscsiTargetGetInstance
	return r
}

func (r ApiIscsiTargetGetInstancePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IscsiTargetGetInstancePostExecute(r)
}

/*
IscsiTargetGetInstancePost Method for IscsiTargetGetInstancePost

Returns instance matching `id`. If `id` is not found, Validation error is raised.

Please see `query` method documentation for `options`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIscsiTargetGetInstancePostRequest
*/
func (a *IscsiTargetApiService) IscsiTargetGetInstancePost(ctx context.Context) ApiIscsiTargetGetInstancePostRequest {
	return ApiIscsiTargetGetInstancePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IscsiTargetApiService) IscsiTargetGetInstancePostExecute(r ApiIscsiTargetGetInstancePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetApiService.IscsiTargetGetInstancePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/target/get_instance"

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
	localVarPostBody = r.iscsiTargetGetInstance
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

type ApiIscsiTargetIdIdDeleteRequest struct {
	ctx        context.Context
	ApiService *IscsiTargetApiService
	id         int32
}

func (r ApiIscsiTargetIdIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.IscsiTargetIdIdDeleteExecute(r)
}

/*
IscsiTargetIdIdDelete Method for IscsiTargetIdIdDelete

Delete iSCSI Target of `id`.

Deleting an iSCSI Target makes sure we delete all Associated Targets which use `id` iSCSI Target.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiIscsiTargetIdIdDeleteRequest
*/
func (a *IscsiTargetApiService) IscsiTargetIdIdDelete(ctx context.Context, id int32) ApiIscsiTargetIdIdDeleteRequest {
	return ApiIscsiTargetIdIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IscsiTargetApiService) IscsiTargetIdIdDeleteExecute(r ApiIscsiTargetIdIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetApiService.IscsiTargetIdIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/target/id/{id}"
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

type ApiIscsiTargetIdIdGetRequest struct {
	ctx        context.Context
	ApiService *IscsiTargetApiService
	id         int32
}

func (r ApiIscsiTargetIdIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.IscsiTargetIdIdGetExecute(r)
}

/*
IscsiTargetIdIdGet Method for IscsiTargetIdIdGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiIscsiTargetIdIdGetRequest
*/
func (a *IscsiTargetApiService) IscsiTargetIdIdGet(ctx context.Context, id int32) ApiIscsiTargetIdIdGetRequest {
	return ApiIscsiTargetIdIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IscsiTargetApiService) IscsiTargetIdIdGetExecute(r ApiIscsiTargetIdIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetApiService.IscsiTargetIdIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/target/id/{id}"
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

type ApiIscsiTargetIdIdPutRequest struct {
	ctx        context.Context
	ApiService *IscsiTargetApiService
	id         int32
}

func (r ApiIscsiTargetIdIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.IscsiTargetIdIdPutExecute(r)
}

/*
IscsiTargetIdIdPut Method for IscsiTargetIdIdPut

Update iSCSI Target of `id`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiIscsiTargetIdIdPutRequest
*/
func (a *IscsiTargetApiService) IscsiTargetIdIdPut(ctx context.Context, id int32) ApiIscsiTargetIdIdPutRequest {
	return ApiIscsiTargetIdIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IscsiTargetApiService) IscsiTargetIdIdPutExecute(r ApiIscsiTargetIdIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetApiService.IscsiTargetIdIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/target/id/{id}"
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

type ApiIscsiTargetPostRequest struct {
	ctx                context.Context
	ApiService         *IscsiTargetApiService
	iscsiTargetCreate0 *IscsiTargetCreate0
}

func (r ApiIscsiTargetPostRequest) IscsiTargetCreate0(iscsiTargetCreate0 IscsiTargetCreate0) ApiIscsiTargetPostRequest {
	r.iscsiTargetCreate0 = &iscsiTargetCreate0
	return r
}

func (r ApiIscsiTargetPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IscsiTargetPostExecute(r)
}

/*
IscsiTargetPost Method for IscsiTargetPost

Create an iSCSI Target.

`groups` is a list of group dictionaries which provide information related to using a `portal`, `initiator`,
`authmethod` and `auth` with this target. `auth` represents a valid iSCSI Authorized Access and defaults to
null.

`auth_networks` is a list of IP/CIDR addresses which are allowed to use this initiator. If all networks are
to be allowed, this field should be left empty.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIscsiTargetPostRequest
*/
func (a *IscsiTargetApiService) IscsiTargetPost(ctx context.Context) ApiIscsiTargetPostRequest {
	return ApiIscsiTargetPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IscsiTargetApiService) IscsiTargetPostExecute(r ApiIscsiTargetPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetApiService.IscsiTargetPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/target"

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
	localVarPostBody = r.iscsiTargetCreate0
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
