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
)

// OpenvpnServerApiService OpenvpnServerApi service
type OpenvpnServerApiService service

type ApiOpenvpnServerAuthenticationAlgorithmChoicesGetRequest struct {
	ctx        context.Context
	ApiService *OpenvpnServerApiService
}

func (r ApiOpenvpnServerAuthenticationAlgorithmChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnServerAuthenticationAlgorithmChoicesGetExecute(r)
}

/*
OpenvpnServerAuthenticationAlgorithmChoicesGet Method for OpenvpnServerAuthenticationAlgorithmChoicesGet

Returns a dictionary of valid authentication algorithms which can be used with OpenVPN server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenvpnServerAuthenticationAlgorithmChoicesGetRequest
*/
func (a *OpenvpnServerApiService) OpenvpnServerAuthenticationAlgorithmChoicesGet(ctx context.Context) ApiOpenvpnServerAuthenticationAlgorithmChoicesGetRequest {
	return ApiOpenvpnServerAuthenticationAlgorithmChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnServerApiService) OpenvpnServerAuthenticationAlgorithmChoicesGetExecute(r ApiOpenvpnServerAuthenticationAlgorithmChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnServerApiService.OpenvpnServerAuthenticationAlgorithmChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/server/authentication_algorithm_choices"

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

type ApiOpenvpnServerCipherChoicesGetRequest struct {
	ctx        context.Context
	ApiService *OpenvpnServerApiService
}

func (r ApiOpenvpnServerCipherChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnServerCipherChoicesGetExecute(r)
}

/*
OpenvpnServerCipherChoicesGet Method for OpenvpnServerCipherChoicesGet

Returns a dictionary of valid ciphers which can be used with OpenVPN server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenvpnServerCipherChoicesGetRequest
*/
func (a *OpenvpnServerApiService) OpenvpnServerCipherChoicesGet(ctx context.Context) ApiOpenvpnServerCipherChoicesGetRequest {
	return ApiOpenvpnServerCipherChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnServerApiService) OpenvpnServerCipherChoicesGetExecute(r ApiOpenvpnServerCipherChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnServerApiService.OpenvpnServerCipherChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/server/cipher_choices"

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

type ApiOpenvpnServerClientConfigurationGenerationPostRequest struct {
	ctx                                        context.Context
	ApiService                                 *OpenvpnServerApiService
	openvpnServerClientConfigurationGeneration *OpenvpnServerClientConfigurationGeneration
}

func (r ApiOpenvpnServerClientConfigurationGenerationPostRequest) OpenvpnServerClientConfigurationGeneration(openvpnServerClientConfigurationGeneration OpenvpnServerClientConfigurationGeneration) ApiOpenvpnServerClientConfigurationGenerationPostRequest {
	r.openvpnServerClientConfigurationGeneration = &openvpnServerClientConfigurationGeneration
	return r
}

func (r ApiOpenvpnServerClientConfigurationGenerationPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnServerClientConfigurationGenerationPostExecute(r)
}

/*
OpenvpnServerClientConfigurationGenerationPost Method for OpenvpnServerClientConfigurationGenerationPost

Returns a configuration for OpenVPN client which can be used with any client to connect to FN/TN OpenVPN
server.

`client_certificate_id` should be a valid certificate issued for use with OpenVPN client service.

`server_address` if specified auto-fills the remote directive in the OpenVPN configuration enabling the end
user to use the file without making any edits to connect to OpenVPN server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenvpnServerClientConfigurationGenerationPostRequest
*/
func (a *OpenvpnServerApiService) OpenvpnServerClientConfigurationGenerationPost(ctx context.Context) ApiOpenvpnServerClientConfigurationGenerationPostRequest {
	return ApiOpenvpnServerClientConfigurationGenerationPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnServerApiService) OpenvpnServerClientConfigurationGenerationPostExecute(r ApiOpenvpnServerClientConfigurationGenerationPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnServerApiService.OpenvpnServerClientConfigurationGenerationPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/server/client_configuration_generation"

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
	localVarPostBody = r.openvpnServerClientConfigurationGeneration
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

type ApiOpenvpnServerGetRequest struct {
	ctx        context.Context
	ApiService *OpenvpnServerApiService
}

func (r ApiOpenvpnServerGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnServerGetExecute(r)
}

/*
OpenvpnServerGet Method for OpenvpnServerGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenvpnServerGetRequest
*/
func (a *OpenvpnServerApiService) OpenvpnServerGet(ctx context.Context) ApiOpenvpnServerGetRequest {
	return ApiOpenvpnServerGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnServerApiService) OpenvpnServerGetExecute(r ApiOpenvpnServerGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnServerApiService.OpenvpnServerGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/server"

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

type ApiOpenvpnServerPutRequest struct {
	ctx                  context.Context
	ApiService           *OpenvpnServerApiService
	openvpnServerUpdate0 *OpenvpnServerUpdate0
}

func (r ApiOpenvpnServerPutRequest) OpenvpnServerUpdate0(openvpnServerUpdate0 OpenvpnServerUpdate0) ApiOpenvpnServerPutRequest {
	r.openvpnServerUpdate0 = &openvpnServerUpdate0
	return r
}

func (r ApiOpenvpnServerPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnServerPutExecute(r)
}

/*
OpenvpnServerPut Method for OpenvpnServerPut

Update OpenVPN Server configuration.

When `tls_crypt_auth_enabled` is enabled and `tls_crypt_auth` not provided, a static key is automatically
generated to be used with OpenVPN server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenvpnServerPutRequest
*/
func (a *OpenvpnServerApiService) OpenvpnServerPut(ctx context.Context) ApiOpenvpnServerPutRequest {
	return ApiOpenvpnServerPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnServerApiService) OpenvpnServerPutExecute(r ApiOpenvpnServerPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnServerApiService.OpenvpnServerPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/server"

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
	localVarPostBody = r.openvpnServerUpdate0
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

type ApiOpenvpnServerRenewStaticKeyGetRequest struct {
	ctx        context.Context
	ApiService *OpenvpnServerApiService
}

func (r ApiOpenvpnServerRenewStaticKeyGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnServerRenewStaticKeyGetExecute(r)
}

/*
OpenvpnServerRenewStaticKeyGet Method for OpenvpnServerRenewStaticKeyGet

Reset OpenVPN server's TLS static key which will be used to encrypt/authenticate control channel packets.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenvpnServerRenewStaticKeyGetRequest
*/
func (a *OpenvpnServerApiService) OpenvpnServerRenewStaticKeyGet(ctx context.Context) ApiOpenvpnServerRenewStaticKeyGetRequest {
	return ApiOpenvpnServerRenewStaticKeyGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnServerApiService) OpenvpnServerRenewStaticKeyGetExecute(r ApiOpenvpnServerRenewStaticKeyGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnServerApiService.OpenvpnServerRenewStaticKeyGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/server/renew_static_key"

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
