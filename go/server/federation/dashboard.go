// Package federation provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package federation

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/chi/v5"
)

const (
	ManagementKeyScopes = "ManagementKey.Scopes"
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

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Summary sessions statistics per countries on daily basis
	// (GET /api/dashboard/daily_sessions_summary)
	DailySessionsSummary(w http.ResponseWriter, r *http.Request, params DailySessionsSummaryParams)
	// User sessions
	// (GET /api/dashboard/daily_user_sessions)
	DailyUserSessions(w http.ResponseWriter, r *http.Request, params DailyUserSessionsParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// DailySessionsSummary operation middleware
func (siw *ServerInterfaceWrapper) DailySessionsSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params DailySessionsSummaryParams

	// ------------- Required query parameter "start" -------------
	if paramValue := r.URL.Query().Get("start"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "start"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "start", r.URL.Query(), &params.Start)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "start", Err: err})
		return
	}

	// ------------- Required query parameter "end" -------------
	if paramValue := r.URL.Query().Get("end"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "end"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "end", r.URL.Query(), &params.End)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "end", Err: err})
		return
	}

	// ------------- Required query parameter "country" -------------
	if paramValue := r.URL.Query().Get("country"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "country"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "country", r.URL.Query(), &params.Country)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "country", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "cursor" -------------
	if paramValue := r.URL.Query().Get("cursor"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "cursor", r.URL.Query(), &params.Cursor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cursor", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DailySessionsSummary(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DailyUserSessions operation middleware
func (siw *ServerInterfaceWrapper) DailyUserSessions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params DailyUserSessionsParams

	// ------------- Required query parameter "start" -------------
	if paramValue := r.URL.Query().Get("start"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "start"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "start", r.URL.Query(), &params.Start)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "start", Err: err})
		return
	}

	// ------------- Required query parameter "end" -------------
	if paramValue := r.URL.Query().Get("end"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "end"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "end", r.URL.Query(), &params.End)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "end", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "cursor" -------------
	if paramValue := r.URL.Query().Get("cursor"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "cursor", r.URL.Query(), &params.Cursor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cursor", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DailyUserSessions(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/dashboard/daily_sessions_summary", wrapper.DailySessionsSummary)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/dashboard/daily_user_sessions", wrapper.DailyUserSessions)
	})

	return r
}
