// Package tunnel provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package tunnel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	externalRef0 "github.com/vpnhouse/api/go/server/common"
)

const (
	Token_authScopes = "Token_auth.Scopes"
)

// Information to configure client on client side.
type ClientConfiguration struct {
	// Connection information for wireguard peers.
	InfoWireguard *ConnectInfoWireguard `json:"info_wireguard,omitempty"`
}

// Connection information for wireguard peers.
type ConnectInfoWireguard struct {
	// Array of subnet, allowed to be sent to tunnel.
	AllowedIps []string `json:"allowed_ips"`

	// List of DNS servers.
	Dns []string `json:"dns"`

	// Keepalive interval to be set on client side.
	Keepalive    int `json:"keepalive"`
	PingInterval int `json:"ping_interval"`

	// Public IPv4 of a wireguard server.
	ServerIpv4 string `json:"server_ipv4"`

	// Public wireguard port.
	ServerPort int `json:"server_port"`

	// Server public key.
	ServerPublicKey string `json:"server_public_key"`

	// Client's tunnel IPv4 address.
	TunnelIpv4 string `json:"tunnel_ipv4"`
}

// Wireguard-specific tunnel information.
type PeerWireguard struct {
	// Wireguard public key.
	PublicKey *string `json:"public_key,omitempty"`
}

// ClientConnectJSONBody defines parameters for ClientConnect.
type ClientConnectJSONBody struct {
	ExpireSeconds *int                               `json:"expire_seconds,omitempty"`
	Identifiers   externalRef0.ConnectionIdentifiers `json:"identifiers"`

	// Wireguard-specific tunnel information.
	InfoWireguard *PeerWireguard `json:"info_wireguard,omitempty"`
	Location      *string        `json:"location,omitempty"`
}

// ClientDisconnectJSONBody defines parameters for ClientDisconnect.
type ClientDisconnectJSONBody externalRef0.ConnectionIdentifiers

// ClientPingJSONBody defines parameters for ClientPing.
type ClientPingJSONBody externalRef0.ConnectionIdentifiers

// ClientConnectJSONRequestBody defines body for ClientConnect for application/json ContentType.
type ClientConnectJSONRequestBody ClientConnectJSONBody

// ClientDisconnectJSONRequestBody defines body for ClientDisconnect for application/json ContentType.
type ClientDisconnectJSONRequestBody ClientDisconnectJSONBody

// ClientPingJSONRequestBody defines body for ClientPing for application/json ContentType.
type ClientPingJSONRequestBody ClientPingJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Client connect
	// (POST /api/client/connect)
	ClientConnect(w http.ResponseWriter, r *http.Request)

	// (GET /api/client/connect_unsafe)
	ClientConnectUnsafe(w http.ResponseWriter, r *http.Request)
	// Client disconnect
	// (POST /api/client/disconnect)
	ClientDisconnect(w http.ResponseWriter, r *http.Request)
	// Client ping
	// (POST /api/client/ping)
	ClientPing(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ClientConnect operation middleware
func (siw *ServerInterfaceWrapper) ClientConnect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ClientConnect(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ClientConnectUnsafe operation middleware
func (siw *ServerInterfaceWrapper) ClientConnectUnsafe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ClientConnectUnsafe(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ClientDisconnect operation middleware
func (siw *ServerInterfaceWrapper) ClientDisconnect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ClientDisconnect(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ClientPing operation middleware
func (siw *ServerInterfaceWrapper) ClientPing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ClientPing(w, r)
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
		r.Post(options.BaseURL+"/api/client/connect", wrapper.ClientConnect)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client/connect_unsafe", wrapper.ClientConnectUnsafe)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/disconnect", wrapper.ClientDisconnect)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/ping", wrapper.ClientPing)
	})

	return r
}
