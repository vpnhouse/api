// Package client_service provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package client_service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

const (
	ServiceKeyScopes  = "ServiceKey.Scopes"
	ServiceNameScopes = "ServiceName.Scopes"
	BearerScopes      = "bearer.Scopes"
)

// Auth defines model for Auth.
type Auth struct {
	AuthMethodId *string                 `json:"auth_method_id,omitempty"`
	CreatedAt    *time.Time              `json:"created_at,omitempty"`
	ExtendedInfo *map[string]interface{} `json:"extended_info,omitempty"`
	Id           *string                 `json:"id,omitempty"`
	Identifier   *string                 `json:"identifier,omitempty"`
	UpdatedAt    *time.Time              `json:"updated_at,omitempty"`
	UserId       *string                 `json:"user_id,omitempty"`
}

// AuthMethod defines model for AuthMethod.
type AuthMethod struct {
	CreatedAt *time.Time              `json:"created_at,omitempty"`
	Id        *string                 `json:"id,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	ProjectId *string                 `json:"project_id,omitempty"`
	Settings  *map[string]interface{} `json:"settings,omitempty"`
	Type      *string                 `json:"type,omitempty"`
	UpdatedAt *time.Time              `json:"updated_at,omitempty"`
}

// Invite defines model for Invite.
type Invite struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	Email       *string                 `json:"email,omitempty"`
	ExpiresAt   *time.Time              `json:"expires_at,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	LocationId  *string                 `json:"location_id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	QueryId     *string                 `json:"query_id,omitempty"`
	QueryParams *map[string]interface{} `json:"query_params,omitempty"`
	Reminded    *bool                   `json:"reminded,omitempty"`
	Telegram    *string                 `json:"telegram,omitempty"`
	TokenId     *string                 `json:"token_id,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at,omitempty"`
	UserId      *string                 `json:"user_id,omitempty"`
}

// Mailing defines model for Mailing.
type Mailing struct {
	AcceptId   *string    `json:"accept_id,omitempty"`
	Accepted   *bool      `json:"accepted,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	Email      *string    `json:"email,omitempty"`
	Id         *string    `json:"id,omitempty"`
	MailingTag *string    `json:"mailing_tag,omitempty"`
	Status     *string    `json:"status,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

// Project defines model for Project.
type Project struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	Description *map[string]interface{} `json:"description,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at,omitempty"`
}

// Session defines model for Session.
type Session struct {
	ConnectedAt      *time.Time `json:"connected_at,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	Deleted          *bool      `json:"deleted,omitempty"`
	ExpiresAt        *time.Time `json:"expires_at,omitempty"`
	FirstConnectedAt *time.Time `json:"first_connected_at,omitempty"`
	Id               *string    `json:"id,omitempty"`
	Label            *string    `json:"label,omitempty"`
	Node             *string    `json:"node,omitempty"`
	PeerId           *int       `json:"peer_id,omitempty"`
	ToDelete         *bool      `json:"to_delete,omitempty"`
	TokenId          *string    `json:"token_id,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// User defines model for User.
type User struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	Description *map[string]interface{} `json:"description,omitempty"`
	Email       *string                 `json:"email,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	ProjectId   *string                 `json:"project_id,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at,omitempty"`
}

// GetClientFeaturesParams defines parameters for GetClientFeatures.
type GetClientFeaturesParams struct {
	ProjectId       string  `json:"project_id"`
	ClientType      *string `json:"client_type,omitempty"`
	ClientVersion   *string `json:"client_version,omitempty"`
	ClientDeviceId  *string `json:"client_device_id,omitempty"`
	ClientOsVersion *string `json:"client_os_version,omitempty"`
	ClientTimezone  *string `json:"client_timezone,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get client-features
	// (GET /api/client-service/client_features)
	GetClientFeatures(w http.ResponseWriter, r *http.Request, params GetClientFeaturesParams)
	// Get user
	// (GET /api/client-service/user)
	GetUser(w http.ResponseWriter, r *http.Request)
	// List user's auth
	// (GET /api/client-service/user/auth)
	ListUserAuth(w http.ResponseWriter, r *http.Request)
	// List user's auth methods
	// (GET /api/client-service/user/auth-method)
	ListUserAuthMethod(w http.ResponseWriter, r *http.Request)
	// List user's invite
	// (GET /api/client-service/user/invite)
	ListUserInvite(w http.ResponseWriter, r *http.Request)
	// List user's mailing
	// (GET /api/client-service/user/mailing)
	ListUserMailing(w http.ResponseWriter, r *http.Request)
	// Get user's project
	// (GET /api/client-service/user/project)
	GetUserProject(w http.ResponseWriter, r *http.Request)
	// List user's session
	// (GET /api/client-service/user/session)
	ListUserSession(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetClientFeatures operation middleware
func (siw *ServerInterfaceWrapper) GetClientFeatures(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetClientFeaturesParams

	// ------------- Required query parameter "project_id" -------------
	if paramValue := r.URL.Query().Get("project_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "project_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "project_id", r.URL.Query(), &params.ProjectId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "project_id", Err: err})
		return
	}

	// ------------- Optional query parameter "client_type" -------------
	if paramValue := r.URL.Query().Get("client_type"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "client_type", r.URL.Query(), &params.ClientType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "client_type", Err: err})
		return
	}

	// ------------- Optional query parameter "client_version" -------------
	if paramValue := r.URL.Query().Get("client_version"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "client_version", r.URL.Query(), &params.ClientVersion)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "client_version", Err: err})
		return
	}

	// ------------- Optional query parameter "client_device_id" -------------
	if paramValue := r.URL.Query().Get("client_device_id"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "client_device_id", r.URL.Query(), &params.ClientDeviceId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "client_device_id", Err: err})
		return
	}

	// ------------- Optional query parameter "client_os_version" -------------
	if paramValue := r.URL.Query().Get("client_os_version"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "client_os_version", r.URL.Query(), &params.ClientOsVersion)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "client_os_version", Err: err})
		return
	}

	// ------------- Optional query parameter "client_timezone" -------------
	if paramValue := r.URL.Query().Get("client_timezone"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "client_timezone", r.URL.Query(), &params.ClientTimezone)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "client_timezone", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetClientFeatures(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUser(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListUserAuth operation middleware
func (siw *ServerInterfaceWrapper) ListUserAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListUserAuth(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListUserAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) ListUserAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListUserAuthMethod(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListUserInvite operation middleware
func (siw *ServerInterfaceWrapper) ListUserInvite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListUserInvite(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListUserMailing operation middleware
func (siw *ServerInterfaceWrapper) ListUserMailing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListUserMailing(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUserProject operation middleware
func (siw *ServerInterfaceWrapper) GetUserProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserProject(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListUserSession operation middleware
func (siw *ServerInterfaceWrapper) ListUserSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListUserSession(w, r)
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
		r.Get(options.BaseURL+"/api/client-service/client_features", wrapper.GetClientFeatures)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client-service/user", wrapper.GetUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client-service/user/auth", wrapper.ListUserAuth)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client-service/user/auth-method", wrapper.ListUserAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client-service/user/invite", wrapper.ListUserInvite)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client-service/user/mailing", wrapper.ListUserMailing)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client-service/user/project", wrapper.GetUserProject)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client-service/user/session", wrapper.ListUserSession)
	})

	return r
}
