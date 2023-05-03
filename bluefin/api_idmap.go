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

// IdmapApiService IdmapApi service
type IdmapApiService service

type ApiIdmapBackendChoicesGetRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
}

func (r ApiIdmapBackendChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapBackendChoicesGetExecute(r)
}

/*
IdmapBackendChoicesGet Method for IdmapBackendChoicesGet

Returns array of valid idmap backend choices per directory service.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIdmapBackendChoicesGetRequest
*/
func (a *IdmapApiService) IdmapBackendChoicesGet(ctx context.Context) ApiIdmapBackendChoicesGetRequest {
	return ApiIdmapBackendChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapBackendChoicesGetExecute(r ApiIdmapBackendChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapBackendChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/backend_choices"

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

type ApiIdmapBackendOptionsGetRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
}

func (r ApiIdmapBackendOptionsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapBackendOptionsGetExecute(r)
}

/*
IdmapBackendOptionsGet Method for IdmapBackendOptionsGet

This returns full information about idmap backend options. Not all
`options` are valid for every backend.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIdmapBackendOptionsGetRequest
*/
func (a *IdmapApiService) IdmapBackendOptionsGet(ctx context.Context) ApiIdmapBackendOptionsGetRequest {
	return ApiIdmapBackendOptionsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapBackendOptionsGetExecute(r ApiIdmapBackendOptionsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapBackendOptionsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/backend_options"

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

type ApiIdmapClearIdmapCacheGetRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
}

func (r ApiIdmapClearIdmapCacheGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapClearIdmapCacheGetExecute(r)
}

/*
IdmapClearIdmapCacheGet Method for IdmapClearIdmapCacheGet

Stop samba, remove the winbindd_cache.tdb file, start samba, flush samba's cache.
This should be performed after finalizing idmap changes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIdmapClearIdmapCacheGetRequest
*/
func (a *IdmapApiService) IdmapClearIdmapCacheGet(ctx context.Context) ApiIdmapClearIdmapCacheGetRequest {
	return ApiIdmapClearIdmapCacheGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapClearIdmapCacheGetExecute(r ApiIdmapClearIdmapCacheGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapClearIdmapCacheGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/clear_idmap_cache"

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

type ApiIdmapGetRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiIdmapGetRequest) Limit(limit int32) ApiIdmapGetRequest {
	r.limit = &limit
	return r
}

func (r ApiIdmapGetRequest) Offset(offset int32) ApiIdmapGetRequest {
	r.offset = &offset
	return r
}

func (r ApiIdmapGetRequest) Count(count bool) ApiIdmapGetRequest {
	r.count = &count
	return r
}

func (r ApiIdmapGetRequest) Sort(sort string) ApiIdmapGetRequest {
	r.sort = &sort
	return r
}

func (r ApiIdmapGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapGetExecute(r)
}

/*
IdmapGet Method for IdmapGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIdmapGetRequest
*/
func (a *IdmapApiService) IdmapGet(ctx context.Context) ApiIdmapGetRequest {
	return ApiIdmapGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapGetExecute(r ApiIdmapGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap"

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

type ApiIdmapGetInstancePostRequest struct {
	ctx              context.Context
	ApiService       *IdmapApiService
	idmapGetInstance *IdmapGetInstance
}

func (r ApiIdmapGetInstancePostRequest) IdmapGetInstance(idmapGetInstance IdmapGetInstance) ApiIdmapGetInstancePostRequest {
	r.idmapGetInstance = &idmapGetInstance
	return r
}

func (r ApiIdmapGetInstancePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapGetInstancePostExecute(r)
}

/*
IdmapGetInstancePost Method for IdmapGetInstancePost

Returns instance matching `id`. If `id` is not found, Validation error is raised.

Please see `query` method documentation for `options`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIdmapGetInstancePostRequest
*/
func (a *IdmapApiService) IdmapGetInstancePost(ctx context.Context) ApiIdmapGetInstancePostRequest {
	return ApiIdmapGetInstancePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapGetInstancePostExecute(r ApiIdmapGetInstancePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapGetInstancePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/get_instance"

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
	localVarPostBody = r.idmapGetInstance
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

type ApiIdmapIdIdDeleteRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
	id         int32
}

func (r ApiIdmapIdIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapIdIdDeleteExecute(r)
}

/*
IdmapIdIdDelete Method for IdmapIdIdDelete

Delete a domain by id. Deletion of default system domains is not permitted.
In case of registry config for clustered server, this will remove all smb4.conf
entries for the domain associated with the id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiIdmapIdIdDeleteRequest
*/
func (a *IdmapApiService) IdmapIdIdDelete(ctx context.Context, id int32) ApiIdmapIdIdDeleteRequest {
	return ApiIdmapIdIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapIdIdDeleteExecute(r ApiIdmapIdIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapIdIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/id/{id}"
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

type ApiIdmapIdIdGetRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
	id         int32
}

func (r ApiIdmapIdIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapIdIdGetExecute(r)
}

/*
IdmapIdIdGet Method for IdmapIdIdGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiIdmapIdIdGetRequest
*/
func (a *IdmapApiService) IdmapIdIdGet(ctx context.Context, id int32) ApiIdmapIdIdGetRequest {
	return ApiIdmapIdIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapIdIdGetExecute(r ApiIdmapIdIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapIdIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/id/{id}"
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

type ApiIdmapIdIdPutRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
	id         int32
}

func (r ApiIdmapIdIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapIdIdPutExecute(r)
}

/*
IdmapIdIdPut Method for IdmapIdIdPut

Update a domain by id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiIdmapIdIdPutRequest
*/
func (a *IdmapApiService) IdmapIdIdPut(ctx context.Context, id int32) ApiIdmapIdIdPutRequest {
	return ApiIdmapIdIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapIdIdPutExecute(r ApiIdmapIdIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapIdIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/id/{id}"
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

type ApiIdmapOptionsChoicesPostRequest struct {
	ctx        context.Context
	ApiService *IdmapApiService
	body       *string
}

func (r ApiIdmapOptionsChoicesPostRequest) Body(body string) ApiIdmapOptionsChoicesPostRequest {
	r.body = &body
	return r
}

func (r ApiIdmapOptionsChoicesPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapOptionsChoicesPostExecute(r)
}

/*
IdmapOptionsChoicesPost Method for IdmapOptionsChoicesPost

Returns a list of supported keys for the specified idmap backend.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIdmapOptionsChoicesPostRequest
*/
func (a *IdmapApiService) IdmapOptionsChoicesPost(ctx context.Context) ApiIdmapOptionsChoicesPostRequest {
	return ApiIdmapOptionsChoicesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapOptionsChoicesPostExecute(r ApiIdmapOptionsChoicesPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapOptionsChoicesPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/options_choices"

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
	localVarPostBody = r.body
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

type ApiIdmapPostRequest struct {
	ctx          context.Context
	ApiService   *IdmapApiService
	idmapCreate0 *IdmapCreate0
}

func (r ApiIdmapPostRequest) IdmapCreate0(idmapCreate0 IdmapCreate0) ApiIdmapPostRequest {
	r.idmapCreate0 = &idmapCreate0
	return r
}

func (r ApiIdmapPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IdmapPostExecute(r)
}

/*
IdmapPost Method for IdmapPost

Create a new IDMAP domain. These domains must be unique. This table
will be automatically populated after joining an Active Directory domain
if "allow trusted domains" is set to True in the AD service configuration.
There are three default system domains: DS_TYPE_ACTIVEDIRECTORY, DS_TYPE_LDAP, DS_TYPE_DEFAULT_DOMAIN.
The system domains correspond with the idmap settings under Active Directory, LDAP, and SMB
respectively.

`name` the pre-windows 2000 domain name.

`DNS_domain_name` DNS name of the domain.

`idmap_backend` provides a plugin interface for Winbind to use varying
backends to store SID/uid/gid mapping tables. The correct setting
depends on the environment in which the NAS is deployed.

`range_low` and `range_high` specify the UID and GID range for which this backend is authoritative.

`certificate_id` references the certificate ID of the SSL certificate to use for certificate-based
authentication to a remote LDAP server. This parameter is not supported for all idmap backends as some
backends will generate SID to ID mappings algorithmically without causing network traffic.

`options` are additional parameters that are backend-dependent:

`AD` idmap backend options:
`unix_primary_group` If True, the primary group membership is fetched from the LDAP attributes (gidNumber).
If False, the primary group membership is calculated via the "primaryGroupID" LDAP attribute.

`unix_nss_info` if True winbind will retrieve the login shell and home directory from the LDAP attributes.
If False or if the AD LDAP entry lacks the SFU attributes the smb4.conf parameters `template shell` and `template homedir` are used.

`schema_mode` Defines the schema that idmap_ad should use when querying Active Directory regarding user and group information.
This can be either the RFC2307 schema support included in Windows 2003 R2 or the Service for Unix (SFU) schema.
For SFU 3.0 or 3.5 please choose "SFU", for SFU 2.0 please choose "SFU20". The behavior of primary group membership is
controlled by the unix_primary_group option.

`AUTORID` idmap backend options:
`readonly` sets the module to read-only mode. No new ranges will be allocated and new mappings
will not be created in the idmap pool.

`ignore_builtin` ignores mapping requests for the BUILTIN domain.

`LDAP` idmap backend options:
`ldap_base_dn` defines the directory base suffix to use for SID/uid/gid mapping entries.

`ldap_user_dn` defines the user DN to be used for authentication.

`ldap_url` specifies the LDAP server to use for SID/uid/gid map entries.

`ssl` specifies whether to encrypt the LDAP transport for the idmap backend.

`NSS` idmap backend options:
`linked_service` specifies the auxiliary directory service ID provider.

`RFC2307` idmap backend options:
`domain` specifies the domain for which the idmap backend is being created. Numeric id, short-form
domain name, or long-form DNS domain name of the domain may be specified. Entry must be entered as
it appears in `idmap.domain`.

`range_low` and `range_high` specify the UID and GID range for which this backend is authoritative.

`ldap_server` defines the type of LDAP server to use. This can either be an LDAP server provided
by the Active Directory Domain (ad) or a stand-alone LDAP server.

`bind_path_user` specfies the search base where user objects can be found in the LDAP server.

`bind_path_group` specifies the search base where group objects can be found in the LDAP server.

`user_cn` query cn attribute instead of uid attribute for the user name in LDAP.

`realm` append @realm to cn for groups (and users if user_cn is set) in LDAP queries.

`ldmap_domain` when using the LDAP server in the Active Directory server, this allows one to
specify the domain where to access the Active Directory server. This allows using trust relationships
while keeping all RFC 2307 records in one place. This parameter is optional, the default is to access
the AD server in the current domain to query LDAP records.

`ldap_url` when using a stand-alone LDAP server, this parameter specifies the LDAP URL for accessing the LDAP server.

`ldap_user_dn` defines the user DN to be used for authentication.

`ldap_user_dn_password` is the password to be used for LDAP authentication.

`realm` defines the realm to use in the user and group names. This is only required when using cn_realm together with

	a stand-alone ldap server.

`RID` backend options:
`sssd_compat` generate idmap low range based on same algorithm that SSSD uses by default.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIdmapPostRequest
*/
func (a *IdmapApiService) IdmapPost(ctx context.Context) ApiIdmapPostRequest {
	return ApiIdmapPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IdmapApiService) IdmapPostExecute(r ApiIdmapPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap"

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
	localVarPostBody = r.idmapCreate0
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
