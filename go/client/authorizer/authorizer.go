// Package authorizer provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package authorizer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ServiceKeyScopes = "ServiceKey.Scopes"
	BasicScopes      = "basic.Scopes"
	BearerScopes     = "bearer.Scopes"
)

// AuthRequest defines model for AuthRequest.
type AuthRequest struct {
	AuthInfo       string `json:"auth_info"`
	AuthType       string `json:"auth_type"`
	ClientPlatform string `json:"client_platform"`
	ClientVersion  string `json:"client_version"`
	DeviceId       string `json:"device_id"`
	Project        string `json:"project"`
}

// AuthResponse defines model for AuthResponse.
type AuthResponse struct {
	AccessToken        string    `json:"access_token"`
	DiscoveryAddresses *[]string `json:"discovery_addresses,omitempty"`
	RefreshToken       *string   `json:"refresh_token,omitempty"`
}

// AuthServiceRequest defines model for AuthServiceRequest.
type AuthServiceRequest struct {
	Project   string `json:"project"`
	ServiceId string `json:"service_id"`
}

// AuthenticateJSONBody defines parameters for Authenticate.
type AuthenticateJSONBody AuthRequest

// ServiceAuthenticateJSONBody defines parameters for ServiceAuthenticate.
type ServiceAuthenticateJSONBody AuthServiceRequest

// AuthenticateJSONRequestBody defines body for Authenticate for application/json ContentType.
type AuthenticateJSONRequestBody AuthenticateJSONBody

// ServiceAuthenticateJSONRequestBody defines body for ServiceAuthenticate for application/json ContentType.
type ServiceAuthenticateJSONRequestBody ServiceAuthenticateJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// Authenticate request with any body
	AuthenticateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Authenticate(ctx context.Context, body AuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ServiceAuthenticate request with any body
	ServiceAuthenticateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ServiceAuthenticate(ctx context.Context, body ServiceAuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AuthenticateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAuthenticateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Authenticate(ctx context.Context, body AuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAuthenticateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ServiceAuthenticateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewServiceAuthenticateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ServiceAuthenticate(ctx context.Context, body ServiceAuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewServiceAuthenticateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAuthenticateRequest calls the generic Authenticate builder with application/json body
func NewAuthenticateRequest(server string, body AuthenticateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAuthenticateRequestWithBody(server, "application/json", bodyReader)
}

// NewAuthenticateRequestWithBody generates requests for Authenticate with any type of body
func NewAuthenticateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client/signin")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewServiceAuthenticateRequest calls the generic ServiceAuthenticate builder with application/json body
func NewServiceAuthenticateRequest(server string, body ServiceAuthenticateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewServiceAuthenticateRequestWithBody(server, "application/json", bodyReader)
}

// NewServiceAuthenticateRequestWithBody generates requests for ServiceAuthenticate with any type of body
func NewServiceAuthenticateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/service/signin")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// Authenticate request with any body
	AuthenticateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error)

	AuthenticateWithResponse(ctx context.Context, body AuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error)

	// ServiceAuthenticate request with any body
	ServiceAuthenticateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ServiceAuthenticateResponse, error)

	ServiceAuthenticateWithResponse(ctx context.Context, body ServiceAuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*ServiceAuthenticateResponse, error)
}

type AuthenticateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AuthResponse
	JSON400      *Error
	JSON401      *Error
	JSON403      *Error
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r AuthenticateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AuthenticateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ServiceAuthenticateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AuthResponse
	JSON400      *Error
	JSON401      *Error
	JSON403      *Error
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r ServiceAuthenticateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ServiceAuthenticateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AuthenticateWithBodyWithResponse request with arbitrary body returning *AuthenticateResponse
func (c *ClientWithResponses) AuthenticateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error) {
	rsp, err := c.AuthenticateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAuthenticateResponse(rsp)
}

func (c *ClientWithResponses) AuthenticateWithResponse(ctx context.Context, body AuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error) {
	rsp, err := c.Authenticate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAuthenticateResponse(rsp)
}

// ServiceAuthenticateWithBodyWithResponse request with arbitrary body returning *ServiceAuthenticateResponse
func (c *ClientWithResponses) ServiceAuthenticateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ServiceAuthenticateResponse, error) {
	rsp, err := c.ServiceAuthenticateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseServiceAuthenticateResponse(rsp)
}

func (c *ClientWithResponses) ServiceAuthenticateWithResponse(ctx context.Context, body ServiceAuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*ServiceAuthenticateResponse, error) {
	rsp, err := c.ServiceAuthenticate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseServiceAuthenticateResponse(rsp)
}

// ParseAuthenticateResponse parses an HTTP response from a AuthenticateWithResponse call
func ParseAuthenticateResponse(rsp *http.Response) (*AuthenticateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AuthenticateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AuthResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseServiceAuthenticateResponse parses an HTTP response from a ServiceAuthenticateWithResponse call
func ParseServiceAuthenticateResponse(rsp *http.Response) (*ServiceAuthenticateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ServiceAuthenticateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AuthResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}