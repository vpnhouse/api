// Package federation provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package federation

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

const (
	AuthorizerKeyScopes = "AuthorizerKey.Scopes"
	DiscoveryKeyScopes  = "DiscoveryKey.Scopes"
	ManagementKeyScopes = "ManagementKey.Scopes"
)

// Defines values for NodeState.
const (
	NodeStateDisabled NodeState = "disabled"

	NodeStateEnabled NodeState = "enabled"
)

// Node defines model for Node.
type Node struct {
	Addresses []string               `json:"addresses"`
	AliasOf   *string                `json:"alias_of,omitempty"`
	Created   time.Time              `json:"created"`
	Healthy   *bool                  `json:"healthy,omitempty"`
	Labels    map[string]interface{} `json:"labels"`
	Priority  int                    `json:"priority"`
	State     *NodeState             `json:"state,omitempty"`
	Updated   time.Time              `json:"updated"`
}

// NodeAction defines model for NodeAction.
type NodeAction struct {
	AddRestriction    *NodeActionAddRestriction `json:"add_restriction,omitempty"`
	DeleteRestriction *NodeActionDelRestriction `json:"delete_restriction,omitempty"`
}

// NodeActionAddRestriction defines model for NodeActionAddRestriction.
type NodeActionAddRestriction struct {
	InstallationId *string   `json:"installation_id,omitempty"`
	RestrictedTo   time.Time `json:"restricted_to"`
	SessionId      *string   `json:"session_id,omitempty"`
	UserId         *string   `json:"user_id,omitempty"`
}

// NodeActionDelRestriction defines model for NodeActionDelRestriction.
type NodeActionDelRestriction struct {
	InstallationId *string `json:"installation_id,omitempty"`
	SessionId      *string `json:"session_id,omitempty"`
	UserId         *string `json:"user_id,omitempty"`
}

// NodeActionRequest defines model for NodeActionRequest.
type NodeActionRequest struct {
	Action NodeAction `json:"action"`
	Node   *string    `json:"node,omitempty"`
}

// NodeRecord defines model for NodeRecord.
type NodeRecord struct {
	Id   string `json:"id"`
	Node Node   `json:"node"`
}

// NodeState defines model for NodeState.
type NodeState string

// PublicKey defines model for PublicKey.
type PublicKey struct {
	Expires *time.Time `json:"expires,omitempty"`
	Key     string     `json:"key"`
}

// PublicKeyRecord defines model for PublicKeyRecord.
type PublicKeyRecord struct {
	Id  string    `json:"id"`
	Key PublicKey `json:"key"`
}

// UpdateKeyJSONBody defines parameters for UpdateKey.
type UpdateKeyJSONBody PublicKey

// SetNodeLabelJSONBody defines parameters for SetNodeLabel.
type SetNodeLabelJSONBody string

// ListNodesParams defines parameters for ListNodes.
type ListNodesParams struct {
	Healthy *bool      `json:"healthy,omitempty"`
	State   *NodeState `json:"state,omitempty"`
}

// NodesActionJSONBody defines parameters for NodesAction.
type NodesActionJSONBody NodeActionRequest

// UpdateNodeJSONBody defines parameters for UpdateNode.
type UpdateNodeJSONBody Node

// RegisterNodeJSONBody defines parameters for RegisterNode.
type RegisterNodeJSONBody Node

// UnsetNodeAddressJSONBody defines parameters for UnsetNodeAddress.
type UnsetNodeAddressJSONBody string

// SetNodeAddressJSONBody defines parameters for SetNodeAddress.
type SetNodeAddressJSONBody string

// UpdateKeyJSONRequestBody defines body for UpdateKey for application/json ContentType.
type UpdateKeyJSONRequestBody UpdateKeyJSONBody

// SetNodeLabelJSONRequestBody defines body for SetNodeLabel for application/json ContentType.
type SetNodeLabelJSONRequestBody SetNodeLabelJSONBody

// NodesActionJSONRequestBody defines body for NodesAction for application/json ContentType.
type NodesActionJSONRequestBody NodesActionJSONBody

// UpdateNodeJSONRequestBody defines body for UpdateNode for application/json ContentType.
type UpdateNodeJSONRequestBody UpdateNodeJSONBody

// RegisterNodeJSONRequestBody defines body for RegisterNode for application/json ContentType.
type RegisterNodeJSONRequestBody RegisterNodeJSONBody

// UnsetNodeAddressJSONRequestBody defines body for UnsetNodeAddress for application/json ContentType.
type UnsetNodeAddressJSONRequestBody UnsetNodeAddressJSONBody

// SetNodeAddressJSONRequestBody defines body for SetNodeAddress for application/json ContentType.
type SetNodeAddressJSONRequestBody SetNodeAddressJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Import inventory
	// (POST /api/federation/import/inventory)
	ImportInventory(w http.ResponseWriter, r *http.Request)
	// List public keys
	// (GET /api/federation/keys)
	ListKeys(w http.ResponseWriter, r *http.Request)

	// (DELETE /api/federation/keys/{id})
	DeleteKey(w http.ResponseWriter, r *http.Request, id string)
	// Manage public keys
	// (GET /api/federation/keys/{id})
	GetKeyById(w http.ResponseWriter, r *http.Request, id string)

	// (POST /api/federation/keys/{id})
	UpdateKey(w http.ResponseWriter, r *http.Request, id string)
	// Delete node label
	// (DELETE /api/federation/node/{id}/labels/{label})
	DeleteNodeLabel(w http.ResponseWriter, r *http.Request, id string, label string)
	// Get node label value
	// (GET /api/federation/node/{id}/labels/{label})
	GetNodeLabel(w http.ResponseWriter, r *http.Request, id string, label string)
	// Set node label value
	// (POST /api/federation/node/{id}/labels/{label})
	SetNodeLabel(w http.ResponseWriter, r *http.Request, id string, label string)
	// List nodes
	// (GET /api/federation/nodes)
	ListNodes(w http.ResponseWriter, r *http.Request, params ListNodesParams)
	// Send action to nodes
	// (POST /api/federation/nodes/action)
	NodesAction(w http.ResponseWriter, r *http.Request)
	// Delete node
	// (DELETE /api/federation/nodes/{id})
	DeleteNode(w http.ResponseWriter, r *http.Request, id string)
	// Get node info
	// (GET /api/federation/nodes/{id})
	GetNodeInfo(w http.ResponseWriter, r *http.Request, id string)
	// Update node
	// (PATCH /api/federation/nodes/{id})
	UpdateNode(w http.ResponseWriter, r *http.Request, id string)
	// Register node
	// (POST /api/federation/nodes/{id})
	RegisterNode(w http.ResponseWriter, r *http.Request, id string)
	// Unset node address
	// (DELETE /api/federation/nodes/{id}/addresses)
	UnsetNodeAddress(w http.ResponseWriter, r *http.Request, id string)
	// List node addresses
	// (GET /api/federation/nodes/{id}/addresses)
	ListNodeAddresses(w http.ResponseWriter, r *http.Request, id string)
	// Set node address
	// (POST /api/federation/nodes/{id}/addresses)
	SetNodeAddress(w http.ResponseWriter, r *http.Request, id string)
	// List node labels
	// (GET /api/federation/nodes/{id}/labels)
	ListNodeLabels(w http.ResponseWriter, r *http.Request, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ImportInventory operation middleware
func (siw *ServerInterfaceWrapper) ImportInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ImportInventory(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListKeys operation middleware
func (siw *ServerInterfaceWrapper) ListKeys(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListKeys(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteKey operation middleware
func (siw *ServerInterfaceWrapper) DeleteKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteKey(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetKeyById operation middleware
func (siw *ServerInterfaceWrapper) GetKeyById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetKeyById(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateKey operation middleware
func (siw *ServerInterfaceWrapper) UpdateKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	ctx = context.WithValue(ctx, AuthorizerKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateKey(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteNodeLabel operation middleware
func (siw *ServerInterfaceWrapper) DeleteNodeLabel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// ------------- Path parameter "label" -------------
	var label string

	err = runtime.BindStyledParameter("simple", false, "label", chi.URLParam(r, "label"), &label)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "label", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteNodeLabel(w, r, id, label)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetNodeLabel operation middleware
func (siw *ServerInterfaceWrapper) GetNodeLabel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// ------------- Path parameter "label" -------------
	var label string

	err = runtime.BindStyledParameter("simple", false, "label", chi.URLParam(r, "label"), &label)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "label", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetNodeLabel(w, r, id, label)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SetNodeLabel operation middleware
func (siw *ServerInterfaceWrapper) SetNodeLabel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// ------------- Path parameter "label" -------------
	var label string

	err = runtime.BindStyledParameter("simple", false, "label", chi.URLParam(r, "label"), &label)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "label", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetNodeLabel(w, r, id, label)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListNodes operation middleware
func (siw *ServerInterfaceWrapper) ListNodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	ctx = context.WithValue(ctx, DiscoveryKeyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListNodesParams

	// ------------- Optional query parameter "healthy" -------------
	if paramValue := r.URL.Query().Get("healthy"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "healthy", r.URL.Query(), &params.Healthy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "healthy", Err: err})
		return
	}

	// ------------- Optional query parameter "state" -------------
	if paramValue := r.URL.Query().Get("state"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "state", r.URL.Query(), &params.State)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "state", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListNodes(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// NodesAction operation middleware
func (siw *ServerInterfaceWrapper) NodesAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.NodesAction(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteNode operation middleware
func (siw *ServerInterfaceWrapper) DeleteNode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteNode(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetNodeInfo operation middleware
func (siw *ServerInterfaceWrapper) GetNodeInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetNodeInfo(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateNode operation middleware
func (siw *ServerInterfaceWrapper) UpdateNode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateNode(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RegisterNode operation middleware
func (siw *ServerInterfaceWrapper) RegisterNode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RegisterNode(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UnsetNodeAddress operation middleware
func (siw *ServerInterfaceWrapper) UnsetNodeAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UnsetNodeAddress(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListNodeAddresses operation middleware
func (siw *ServerInterfaceWrapper) ListNodeAddresses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListNodeAddresses(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SetNodeAddress operation middleware
func (siw *ServerInterfaceWrapper) SetNodeAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetNodeAddress(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListNodeLabels operation middleware
func (siw *ServerInterfaceWrapper) ListNodeLabels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ManagementKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListNodeLabels(w, r, id)
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
		r.Post(options.BaseURL+"/api/federation/import/inventory", wrapper.ImportInventory)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/federation/keys", wrapper.ListKeys)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/federation/keys/{id}", wrapper.DeleteKey)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/federation/keys/{id}", wrapper.GetKeyById)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/federation/keys/{id}", wrapper.UpdateKey)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/federation/node/{id}/labels/{label}", wrapper.DeleteNodeLabel)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/federation/node/{id}/labels/{label}", wrapper.GetNodeLabel)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/federation/node/{id}/labels/{label}", wrapper.SetNodeLabel)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/federation/nodes", wrapper.ListNodes)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/federation/nodes/action", wrapper.NodesAction)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/federation/nodes/{id}", wrapper.DeleteNode)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/federation/nodes/{id}", wrapper.GetNodeInfo)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/api/federation/nodes/{id}", wrapper.UpdateNode)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/federation/nodes/{id}", wrapper.RegisterNode)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/federation/nodes/{id}/addresses", wrapper.UnsetNodeAddress)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/federation/nodes/{id}/addresses", wrapper.ListNodeAddresses)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/federation/nodes/{id}/addresses", wrapper.SetNodeAddress)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/federation/nodes/{id}/labels", wrapper.ListNodeLabels)
	})

	return r
}
