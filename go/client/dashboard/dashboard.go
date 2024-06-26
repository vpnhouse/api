// Package dashboard provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package dashboard

import (
    externalRef1 "github.com/vpnhouse/api/go/server/common"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	DashboardKeyScopes = "DashboardKey.Scopes"
)

// DailySessionsSummary defines model for DailySessionsSummary.
type DailySessionsSummary struct {
	NextCursor      *string         `json:"next_cursor,omitempty"`
	SessionsSummary *[]DailySummary `json:"sessions_summary,omitempty"`
}

// DailySummary defines model for DailySummary.
type DailySummary struct {
	Date            *openapi_types.Date `json:"date,omitempty"`
	SessionsSummary *[]SessionsSummary  `json:"sessions_summary,omitempty"`
}

// DailyUserSession defines model for DailyUserSession.
type DailyUserSession struct {
	ConnectionTime *string             `json:"connection_time,omitempty"`
	Date           *openapi_types.Date `json:"date,omitempty"`
	EndTime        *time.Time          `json:"end_time,omitempty"`
	Location       *string             `json:"location,omitempty"`
	Received       *string             `json:"received,omitempty"`
	Sent           *string             `json:"sent,omitempty"`
	Sessions       *int                `json:"sessions,omitempty"`
	StartTime      *time.Time          `json:"start_time,omitempty"`
	User           *string             `json:"user,omitempty"`
	UserCountry    *string             `json:"user_country,omitempty"`
}

// DailyUserSessions defines model for DailyUserSessions.
type DailyUserSessions struct {
	NextCursor   *string             `json:"next_cursor,omitempty"`
	UserSessions *[]DailyUserSession `json:"user_sessions,omitempty"`
}

// SessionsSummary defines model for SessionsSummary.
type SessionsSummary struct {
	Connections *int    `json:"connections,omitempty"`
	Country     *string `json:"country,omitempty"`
	Users       *int    `json:"users,omitempty"`
}

// DailySessionsSummaryParams defines parameters for DailySessionsSummary.
type DailySessionsSummaryParams struct {
	Start   openapi_types.Date `json:"start"`
	End     openapi_types.Date `json:"end"`
	Country string             `json:"country"`
	Limit   *int               `json:"limit,omitempty"`
	Cursor  *string            `json:"cursor,omitempty"`
}

// DailyUserSessionsParams defines parameters for DailyUserSessions.
type DailyUserSessionsParams struct {
	Start  openapi_types.Date `json:"start"`
	End    openapi_types.Date `json:"end"`
	Limit  *int               `json:"limit,omitempty"`
	Cursor *string            `json:"cursor,omitempty"`
}

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
	// DailySessionsSummary request
	DailySessionsSummary(ctx context.Context, params *DailySessionsSummaryParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DailyUserSessions request
	DailyUserSessions(ctx context.Context, params *DailyUserSessionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) DailySessionsSummary(ctx context.Context, params *DailySessionsSummaryParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDailySessionsSummaryRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DailyUserSessions(ctx context.Context, params *DailyUserSessionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDailyUserSessionsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDailySessionsSummaryRequest generates requests for DailySessionsSummary
func NewDailySessionsSummaryRequest(server string, params *DailySessionsSummaryParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/dashboard/daily_sessions_summary")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start", runtime.ParamLocationQuery, params.Start); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end", runtime.ParamLocationQuery, params.End); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "country", runtime.ParamLocationQuery, params.Country); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Cursor != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "cursor", runtime.ParamLocationQuery, *params.Cursor); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDailyUserSessionsRequest generates requests for DailyUserSessions
func NewDailyUserSessionsRequest(server string, params *DailyUserSessionsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/dashboard/daily_user_sessions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start", runtime.ParamLocationQuery, params.Start); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end", runtime.ParamLocationQuery, params.End); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Cursor != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "cursor", runtime.ParamLocationQuery, *params.Cursor); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

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
	// DailySessionsSummary request
	DailySessionsSummaryWithResponse(ctx context.Context, params *DailySessionsSummaryParams, reqEditors ...RequestEditorFn) (*DailySessionsSummaryResponse, error)

	// DailyUserSessions request
	DailyUserSessionsWithResponse(ctx context.Context, params *DailyUserSessionsParams, reqEditors ...RequestEditorFn) (*DailyUserSessionsResponse, error)
}

type DailySessionsSummaryResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DailySessionsSummary
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r DailySessionsSummaryResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DailySessionsSummaryResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DailyUserSessionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DailyUserSessions
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r DailyUserSessionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DailyUserSessionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DailySessionsSummaryWithResponse request returning *DailySessionsSummaryResponse
func (c *ClientWithResponses) DailySessionsSummaryWithResponse(ctx context.Context, params *DailySessionsSummaryParams, reqEditors ...RequestEditorFn) (*DailySessionsSummaryResponse, error) {
	rsp, err := c.DailySessionsSummary(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDailySessionsSummaryResponse(rsp)
}

// DailyUserSessionsWithResponse request returning *DailyUserSessionsResponse
func (c *ClientWithResponses) DailyUserSessionsWithResponse(ctx context.Context, params *DailyUserSessionsParams, reqEditors ...RequestEditorFn) (*DailyUserSessionsResponse, error) {
	rsp, err := c.DailyUserSessions(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDailyUserSessionsResponse(rsp)
}

// ParseDailySessionsSummaryResponse parses an HTTP response from a DailySessionsSummaryWithResponse call
func ParseDailySessionsSummaryResponse(rsp *http.Response) (*DailySessionsSummaryResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DailySessionsSummaryResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DailySessionsSummary
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDailyUserSessionsResponse parses an HTTP response from a DailyUserSessionsWithResponse call
func ParseDailyUserSessionsResponse(rsp *http.Response) (*DailyUserSessionsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DailyUserSessionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DailyUserSessions
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
