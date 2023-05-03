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

// FtpApiService FtpApi service
type FtpApiService service

type ApiFtpGetRequest struct {
	ctx        context.Context
	ApiService *FtpApiService
}

func (r ApiFtpGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.FtpGetExecute(r)
}

/*
FtpGet Method for FtpGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFtpGetRequest
*/
func (a *FtpApiService) FtpGet(ctx context.Context) ApiFtpGetRequest {
	return ApiFtpGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *FtpApiService) FtpGetExecute(r ApiFtpGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FtpApiService.FtpGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ftp"

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

type ApiFtpPutRequest struct {
	ctx        context.Context
	ApiService *FtpApiService
	ftpUpdate0 *FtpUpdate0
}

func (r ApiFtpPutRequest) FtpUpdate0(ftpUpdate0 FtpUpdate0) ApiFtpPutRequest {
	r.ftpUpdate0 = &ftpUpdate0
	return r
}

func (r ApiFtpPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.FtpPutExecute(r)
}

/*
FtpPut Method for FtpPut

Update ftp service configuration.

`clients` is an integer value which sets the maximum number of simultaneous clients allowed. It defaults to 32.

`ipconnections` is an integer value which shows the maximum number of connections per IP address. It defaults
to 0 which equals to unlimited.

`timeout` is the maximum number of seconds that proftpd will allow clients to stay connected without receiving
any data on either the control or data connection.

`timeout_notransfer` is the maximum number of seconds a client is allowed to spend connected, after
authentication, without issuing a command which results in creating an active or passive data connection
(i.e. sending/receiving a file, or receiving a directory listing).

`rootlogin` is a boolean value which when configured to true enables login as root. This is generally
discouraged because of the security risks.

`onlyanonymous` allows anonymous FTP logins with access to the directory specified by `anonpath`.

`banner` is a message displayed to local login users after they successfully authenticate. It is not displayed
to anonymous login users.

`filemask` sets the default permissions for newly created files which by default are 077.

`dirmask` sets the default permissions for newly created directories which by default are 077.

`resume` if set allows FTP clients to resume interrupted transfers.

`fxp` if set to true indicates that File eXchange Protocol is enabled. Generally it is discouraged as it
makes the server vulnerable to FTP bounce attacks.

`defaultroot` when set ensures that for local users, home directory access is only granted if the user
is a member of group wheel.

`ident` is a boolean value which when set to true indicates that IDENT authentication is required. If identd
is not running on the client, this can result in timeouts.

`masqaddress` is the public IP address or hostname which is set if FTP clients cannot connect through a
NAT device.

`localuserbw` is a positive integer value which indicates maximum upload bandwidth in KB/s for local user.
Default of zero indicates unlimited upload bandwidth ( from the FTP server configuration ).

`localuserdlbw` is a positive integer value which indicates maximum download bandwidth in KB/s for local user.
Default of zero indicates unlimited download bandwidth ( from the FTP server configuration ).

`anonuserbw` is a positive integer value which indicates maximum upload bandwidth in KB/s for anonymous user.
Default of zero indicates unlimited upload bandwidth ( from the FTP server configuration ).

`anonuserdlbw` is a positive integer value which indicates maximum download bandwidth in KB/s for anonymous
user. Default of zero indicates unlimited download bandwidth ( from the FTP server configuration ).

`tls` is a boolean value which when set indicates that encrypted connections are enabled. This requires a
certificate to be configured first with the certificate service and the id of certificate is passed on in
`ssltls_certificate`.

`tls_policy` defines whether the control channel, data channel, both channels, or neither channel of an FTP
session must occur over SSL/TLS.

`tls_opt_enable_diags` is a boolean value when set, logs verbosely. This is helpful when troubleshooting a
connection.

`options` is a string used to add proftpd(8) parameters not covered by ftp service.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFtpPutRequest
*/
func (a *FtpApiService) FtpPut(ctx context.Context) ApiFtpPutRequest {
	return ApiFtpPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *FtpApiService) FtpPutExecute(r ApiFtpPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FtpApiService.FtpPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ftp"

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
	localVarPostBody = r.ftpUpdate0
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
