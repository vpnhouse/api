// Package limit_service provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package limit_service

import (
    externalRef1 "github.com/vpnhouse/api/go/server/common"
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
	ServiceKeyScopes  = "ServiceKey.Scopes"
	ServiceNameScopes = "ServiceName.Scopes"
)

// DailyLimit defines model for DailyLimit.
type DailyLimit struct {
	AuthMethod string `json:"auth_method"`
	Limit      int    `json:"limit"`
}

// LimitConfig defines model for LimitConfig.
type LimitConfig struct {
	DailyLimits []DailyLimit `json:"daily_limits"`
	ProjectId   string       `json:"project_id"`
}

// UpdateLimitConfigParams defines model for UpdateLimitConfigParams.
type UpdateLimitConfigParams struct {
	Updates []LimitConfig `json:"updates"`
}

// UpdateLimitConfigJSONBody defines parameters for UpdateLimitConfig.
type UpdateLimitConfigJSONBody UpdateLimitConfigParams

// UpdateLimitConfigJSONRequestBody defines body for UpdateLimitConfig for application/json ContentType.
type UpdateLimitConfigJSONRequestBody UpdateLimitConfigJSONBody

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
	// ListLimitConfig request
	ListLimitConfig(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateLimitConfig request with any body
	UpdateLimitConfigWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateLimitConfig(ctx context.Context, body UpdateLimitConfigJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListLimitConfig(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListLimitConfigRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateLimitConfigWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateLimitConfigRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateLimitConfig(ctx context.Context, body UpdateLimitConfigJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateLimitConfigRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListLimitConfigRequest generates requests for ListLimitConfig
func NewListLimitConfigRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/limit-service/config")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateLimitConfigRequest calls the generic UpdateLimitConfig builder with application/json body
func NewUpdateLimitConfigRequest(server string, body UpdateLimitConfigJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateLimitConfigRequestWithBody(server, "application/json", bodyReader)
}

// NewUpdateLimitConfigRequestWithBody generates requests for UpdateLimitConfig with any type of body
func NewUpdateLimitConfigRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/limit-service/config")
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
	// ListLimitConfig request
	ListLimitConfigWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListLimitConfigResponse, error)

	// UpdateLimitConfig request with any body
	UpdateLimitConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLimitConfigResponse, error)

	UpdateLimitConfigWithResponse(ctx context.Context, body UpdateLimitConfigJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLimitConfigResponse, error)
}

type ListLimitConfigResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]LimitConfig
	JSON401      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r ListLimitConfigResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListLimitConfigResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateLimitConfigResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *externalRef1.Error
	JSON401      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r UpdateLimitConfigResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateLimitConfigResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListLimitConfigWithResponse request returning *ListLimitConfigResponse
func (c *ClientWithResponses) ListLimitConfigWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListLimitConfigResponse, error) {
	rsp, err := c.ListLimitConfig(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListLimitConfigResponse(rsp)
}

// UpdateLimitConfigWithBodyWithResponse request with arbitrary body returning *UpdateLimitConfigResponse
func (c *ClientWithResponses) UpdateLimitConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLimitConfigResponse, error) {
	rsp, err := c.UpdateLimitConfigWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateLimitConfigResponse(rsp)
}

func (c *ClientWithResponses) UpdateLimitConfigWithResponse(ctx context.Context, body UpdateLimitConfigJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLimitConfigResponse, error) {
	rsp, err := c.UpdateLimitConfig(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateLimitConfigResponse(rsp)
}

// ParseListLimitConfigResponse parses an HTTP response from a ListLimitConfigWithResponse call
func ParseListLimitConfigResponse(rsp *http.Response) (*ListLimitConfigResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListLimitConfigResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []LimitConfig
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseUpdateLimitConfigResponse parses an HTTP response from a UpdateLimitConfigWithResponse call
func ParseUpdateLimitConfigResponse(rsp *http.Response) (*UpdateLimitConfigResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateLimitConfigResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
