// Package authorizer provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package authorizer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
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
	InstallationId string `json:"installation_id"`
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

// SendConfirmationLinkRequest defines model for SendConfirmationLinkRequest.
type SendConfirmationLinkRequest struct {
	Email string `json:"email"`
}

// TokenRequest defines model for TokenRequest.
type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// ConfirmParams defines parameters for Confirm.
type ConfirmParams struct {
	ConfirmationId string `json:"confirmation_id"`
}

// SendConfirmationLinkJSONBody defines parameters for SendConfirmationLink.
type SendConfirmationLinkJSONBody SendConfirmationLinkRequest

// AuthenticateJSONBody defines parameters for Authenticate.
type AuthenticateJSONBody AuthRequest

// RegisterJSONBody defines parameters for Register.
type RegisterJSONBody AuthRequest

// TokenJSONBody defines parameters for Token.
type TokenJSONBody TokenRequest

// ServiceAuthenticateJSONBody defines parameters for ServiceAuthenticate.
type ServiceAuthenticateJSONBody AuthServiceRequest

// SendConfirmationLinkJSONRequestBody defines body for SendConfirmationLink for application/json ContentType.
type SendConfirmationLinkJSONRequestBody SendConfirmationLinkJSONBody

// AuthenticateJSONRequestBody defines body for Authenticate for application/json ContentType.
type AuthenticateJSONRequestBody AuthenticateJSONBody

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody RegisterJSONBody

// TokenJSONRequestBody defines body for Token for application/json ContentType.
type TokenJSONRequestBody TokenJSONBody

// ServiceAuthenticateJSONRequestBody defines body for ServiceAuthenticate for application/json ContentType.
type ServiceAuthenticateJSONRequestBody ServiceAuthenticateJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Confirm email
	// (GET /api/client/confirm)
	Confirm(w http.ResponseWriter, r *http.Request, params ConfirmParams)
	// Send confirmation link
	// (POST /api/client/send-confirmation-link)
	SendConfirmationLink(w http.ResponseWriter, r *http.Request)
	// Authenticate user
	// (POST /api/client/signin)
	Authenticate(w http.ResponseWriter, r *http.Request)
	// Register user
	// (POST /api/client/signup)
	Register(w http.ResponseWriter, r *http.Request)
	// Refresh access token
	// (POST /api/client/token)
	Token(w http.ResponseWriter, r *http.Request)
	// Authenticate service
	// (POST /api/service/signin)
	ServiceAuthenticate(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// Confirm operation middleware
func (siw *ServerInterfaceWrapper) Confirm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ConfirmParams

	// ------------- Required query parameter "confirmation_id" -------------
	if paramValue := r.URL.Query().Get("confirmation_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "confirmation_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "confirmation_id", r.URL.Query(), &params.ConfirmationId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "confirmation_id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Confirm(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SendConfirmationLink operation middleware
func (siw *ServerInterfaceWrapper) SendConfirmationLink(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SendConfirmationLink(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Authenticate operation middleware
func (siw *ServerInterfaceWrapper) Authenticate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	ctx = context.WithValue(ctx, BasicScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Authenticate(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	ctx = context.WithValue(ctx, BasicScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Register(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Token operation middleware
func (siw *ServerInterfaceWrapper) Token(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	ctx = context.WithValue(ctx, BasicScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Token(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ServiceAuthenticate operation middleware
func (siw *ServerInterfaceWrapper) ServiceAuthenticate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ServiceAuthenticate(w, r)
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
		r.Get(options.BaseURL+"/api/client/confirm", wrapper.Confirm)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/send-confirmation-link", wrapper.SendConfirmationLink)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/signin", wrapper.Authenticate)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/signup", wrapper.Register)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/token", wrapper.Token)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/service/signin", wrapper.ServiceAuthenticate)
	})

	return r
}
