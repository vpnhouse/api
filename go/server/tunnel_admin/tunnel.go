// Package tunnel_admin provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package tunnel_admin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
	externalRef0 "github.com/vpnhouse/api/go/server/common"
	externalRef1 "github.com/vpnhouse/api/go/server/tunnel"
)

const (
	Basic_authScopes = "Basic_auth.Scopes"
	Token_authScopes = "Token_auth.Scopes"
)

// Defines values for DomainConfigMode.
const (
	DomainConfigModeDirect DomainConfigMode = "direct"

	DomainConfigModeReverseProxy DomainConfigMode = "reverse-proxy"
)

// Defines values for DomainConfigSchema.
const (
	DomainConfigSchemaHttp DomainConfigSchema = "http"

	DomainConfigSchemaHttps DomainConfigSchema = "https"
)

// Defines values for PeerNetAccessPolicy.
const (
	PeerNetAccessPolicyN0 PeerNetAccessPolicy = 0

	PeerNetAccessPolicyN1 PeerNetAccessPolicy = 1

	PeerNetAccessPolicyN2 PeerNetAccessPolicy = 2
)

// Defines values for PeerActivationStatus.
const (
	PeerActivationStatusActivated PeerActivationStatus = "activated"

	PeerActivationStatusNotActivated PeerActivationStatus = "not_activated"
)

// AdminAuth defines model for AdminAuth.
type AdminAuth struct {
	// JWT for accessing other administrative endpoints.
	AccessToken string `json:"access_token"`
}

// Domain name, reverse proxy, and SSL configuration,
// used for the initial configuration.
type DomainConfig struct {
	// Domain name for the service, required.
	// In the "direct" mode we'll issue the SSL certificate for that domain name,
	// in "reverse-proxy" mode we need this name to build the extenral links.
	// If no configuraton is provided (InitialSetupRequest->domain is empty) - the external IP is used.
	DomainName string `json:"domain_name"`

	// We'll try to issue the SSL certificate if set.
	// For the "direct" mode only.
	IssueSsl bool `json:"issue_ssl"`

	// Shows how the http traffic delivered to the service.
	// The "direct" meand that we serve 80/443 by ourselves, so
	// we can also manage the SSL certificates.
	// The "reverse-proxy" means that we're behind the reverse proxy,
	// like nginx. We wont manage and serve the SSL traffic in that case.
	Mode DomainConfigMode `json:"mode"`

	// How the reverse-proxy serving our traffic for the external clients.
	// So the schema + domain_name produces a valid link to the service.
	// For the "reverse-proxy" mode only.
	Schema DomainConfigSchema `json:"schema"`
}

// Shows how the http traffic delivered to the service.
// The "direct" meand that we serve 80/443 by ourselves, so
// we can also manage the SSL certificates.
// The "reverse-proxy" means that we're behind the reverse proxy,
// like nginx. We wont manage and serve the SSL traffic in that case.
type DomainConfigMode string

// How the reverse-proxy serving our traffic for the external clients.
// So the schema + domain_name produces a valid link to the service.
// For the "reverse-proxy" mode only.
type DomainConfigSchema string

// InitialSetupRequest defines model for InitialSetupRequest.
type InitialSetupRequest struct {
	AdminPassword string `json:"admin_password"`

	// Domain name, reverse proxy, and SSL configuration,
	// used for the initial configuration.
	Domain *DomainConfig `json:"domain,omitempty"`

	// allow to send anonymous hearbeats and install notifications
	SendStats    *bool  `json:"send_stats,omitempty"`
	ServerIpMask string `json:"server_ip_mask"`
}

// IpPoolAddress defines model for IpPoolAddress.
type IpPoolAddress struct {
	IpAddress string `json:"ip_address"`
}

// Peer representation.
type Peer struct {
	// JWT information data.
	Claims *string `json:"claims"`

	// The date when the peer was created.
	Created *time.Time `json:"created,omitempty"`

	// Peer expiration time.
	Expires     *time.Time                          `json:"expires"`
	Identifiers *externalRef0.ConnectionIdentifiers `json:"identifiers,omitempty"`

	// Wireguard-specific tunnel information.
	InfoWireguard *externalRef1.PeerWireguard `json:"info_wireguard,omitempty"`

	// Tunneling IPv4 address of a peer.
	Ipv4 *string `json:"ipv4"`

	// Label of the peer.
	Label *string `json:"label"`

	// Network policy: isolate the peer (internet access only)
	// or allow to talk to its network neighbours (like in LANs)
	NetAccessPolicy *PeerNetAccessPolicy `json:"net_access_policy"`

	// How much of bandwidth the client is allowed to consume.
	// Takes no effect if the traffic control subsystem is
	// disabled on a node.
	// [!] Bits per Second, must follow SI.
	RateLimit *int `json:"rate_limit"`

	// The date when the peer was updated last time.
	Updated *time.Time `json:"updated,omitempty"`
}

// Network policy: isolate the peer (internet access only)
// or allow to talk to its network neighbours (like in LANs)
type PeerNetAccessPolicy int

// Returns the status of the shared peer.
// "not_activated" - no configuration has been given, we can
//   activate it immeadietly.
// "activated" - the peer has already been activated,
//   we must ask a user about a re-activation (previously
//   issued credentials will be invalidated).
type PeerActivation struct {
	Status PeerActivationStatus `json:"status"`
}

// PeerActivationStatus defines model for PeerActivation.Status.
type PeerActivationStatus string

// PeerRecord defines model for PeerRecord.
type PeerRecord struct {
	Id int64 `json:"id"`

	// Peer representation.
	Peer Peer `json:"peer"`
}

// Holds current staus flags of the service
type ServiceStatus struct {
	PeersActive    *int `json:"peers_active,omitempty"`
	PeersConnected *int `json:"peers_connected,omitempty"`
	PeersTotal     *int `json:"peers_total,omitempty"`

	// Indicate, whether service requires restart to apply latest settings.
	RestartRequired bool   `json:"restart_required"`
	TrafficRx       *int64 `json:"traffic_rx,omitempty"`
	TrafficTx       *int64 `json:"traffic_tx,omitempty"`
}

// Server-side configuration.
type Settings struct {
	// Admin password (write-only, never returned).
	AdminPassword     *string `json:"admin_password,omitempty"`
	ConnectionTimeout *int    `json:"connection_timeout,omitempty"`

	// DNS servers to announce to a peer
	Dns *[]string `json:"dns,omitempty"`

	// Domain name, reverse proxy, and SSL configuration,
	// used for the initial configuration.
	Domain *DomainConfig `json:"domain,omitempty"`

	// HTTP listening IP:Port pair.
	HttpListenAddr *string `json:"http_listen_addr,omitempty"`
	PingInterval   *int    `json:"ping_interval,omitempty"`

	// allow to send anonymous hearbeats
	SendStats *bool `json:"send_stats,omitempty"`

	// Keepalive interval for wireguard peers.
	WireguardKeepalive *int `json:"wireguard_keepalive,omitempty"`

	// Wireguard listen port inside the container.
	// In 99% cases it matches the `wireguard_server_port` value.
	WireguardListenPort *int `json:"wireguard_listen_port,omitempty"`

	// Wireguard public key (read only).
	WireguardPublicKey *string `json:"wireguard_public_key,omitempty"`

	// Public ipv4 address of a wireguard server.
	WireguardServerIpv4 *string `json:"wireguard_server_ipv4,omitempty"`

	// Public UDP port of a wireguard server.
	// This value is announced to peers, in 99% cases it is the same as the `wireguard_listen_port`.
	// May differs from the `wireguard_listen_port`'s value if NATed (especially with docker).
	WireguardServerPort *int `json:"wireguard_server_port,omitempty"`

	// Subnet for wireguard peers.
	WireguardSubnet *string `json:"wireguard_subnet,omitempty"`
}

// Peer-independent wireguard configuration from a server
type WireguardOptions struct {
	// List of subnets, allowed to be sent to tunnel.
	AllowedIps []string `json:"allowed_ips"`

	// List of DNS servers.
	Dns []string `json:"dns"`

	// Keepalive interval to be set on client side.
	Keepalive int `json:"keepalive"`

	// Public IPv4 of a wireguard server.
	ServerIpv4 string `json:"server_ipv4"`

	// Public wireguard port.
	ServerPort int `json:"server_port"`

	// Server public key.
	ServerPublicKey string `json:"server_public_key"`

	// Network subnet/mask for wireguard clients, e.g 10.235.0.0/24
	Subnet string `json:"subnet"`
}

// AdminAuthResponse defines model for AdminAuthResponse.
type AdminAuthResponse AdminAuth

// IpPoolSuggestResult defines model for IpPoolSuggestResult.
type IpPoolSuggestResult IpPoolAddress

// PeerActivationResponse defines model for PeerActivationResponse.
type PeerActivationResponse struct {
	Peer PeerRecord `json:"peer"`

	// Peer-independent wireguard configuration from a server
	WireguardOptions WireguardOptions `json:"wireguard_options"`
}

// Peer representation.
type PeerInfo Peer

// PeerLink defines model for PeerLink.
type PeerLink struct {
	Link string `json:"link"`
}

// Peer-independent wireguard configuration from a server
type ServerWireguardOptions WireguardOptions

// Holds current staus flags of the service
type ServiceStatusResponse ServiceStatus

// Server-side configuration.
type SettingsInfo Settings

// AdminInitialSetupJSONBody defines parameters for AdminInitialSetup.
type AdminInitialSetupJSONBody InitialSetupRequest

// AdminIppoolIsUsedJSONBody defines parameters for AdminIppoolIsUsed.
type AdminIppoolIsUsedJSONBody IpPoolAddress

// AdminCreatePeerJSONBody defines parameters for AdminCreatePeer.
type AdminCreatePeerJSONBody Peer

// AdminCreateSharedPeerJSONBody defines parameters for AdminCreateSharedPeer.
type AdminCreateSharedPeerJSONBody Peer

// AdminUpdatePeerJSONBody defines parameters for AdminUpdatePeer.
type AdminUpdatePeerJSONBody Peer

// AdminUpdateSettingsJSONBody defines parameters for AdminUpdateSettings.
type AdminUpdateSettingsJSONBody Settings

// PublicPeerActivateJSONBody defines parameters for PublicPeerActivate.
type PublicPeerActivateJSONBody externalRef1.PeerWireguard

// AdminInitialSetupJSONRequestBody defines body for AdminInitialSetup for application/json ContentType.
type AdminInitialSetupJSONRequestBody AdminInitialSetupJSONBody

// AdminIppoolIsUsedJSONRequestBody defines body for AdminIppoolIsUsed for application/json ContentType.
type AdminIppoolIsUsedJSONRequestBody AdminIppoolIsUsedJSONBody

// AdminCreatePeerJSONRequestBody defines body for AdminCreatePeer for application/json ContentType.
type AdminCreatePeerJSONRequestBody AdminCreatePeerJSONBody

// AdminCreateSharedPeerJSONRequestBody defines body for AdminCreateSharedPeer for application/json ContentType.
type AdminCreateSharedPeerJSONRequestBody AdminCreateSharedPeerJSONBody

// AdminUpdatePeerJSONRequestBody defines body for AdminUpdatePeer for application/json ContentType.
type AdminUpdatePeerJSONRequestBody AdminUpdatePeerJSONBody

// AdminUpdateSettingsJSONRequestBody defines body for AdminUpdateSettings for application/json ContentType.
type AdminUpdateSettingsJSONRequestBody AdminUpdateSettingsJSONBody

// PublicPeerActivateJSONRequestBody defines body for PublicPeerActivate for application/json ContentType.
type PublicPeerActivateJSONRequestBody PublicPeerActivateJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Authorize as the node admin
	// (GET /api/tunnel/admin/auth)
	AdminDoAuth(w http.ResponseWriter, r *http.Request)
	// Get peer-independent wireguard configuration from a server
	// (GET /api/tunnel/admin/connection-info/wireguard)
	AdminConnectionInfoWireguard(w http.ResponseWriter, r *http.Request)
	// Set initial parameters
	// (POST /api/tunnel/admin/initial-setup)
	AdminInitialSetup(w http.ResponseWriter, r *http.Request)
	// Suggest an available IP address by the server pool
	// (GET /api/tunnel/admin/ip-pool/suggest)
	AdminIppoolSuggest(w http.ResponseWriter, r *http.Request)
	// Check that the IP address is used by the server pool
	// (POST /api/tunnel/admin/ip-pool/suggest)
	AdminIppoolIsUsed(w http.ResponseWriter, r *http.Request)
	// Get list of peers
	// (GET /api/tunnel/admin/peers)
	AdminListPeers(w http.ResponseWriter, r *http.Request)
	// Create peer
	// (POST /api/tunnel/admin/peers)
	AdminCreatePeer(w http.ResponseWriter, r *http.Request)
	// create a peer for sharing via the unique link
	// (POST /api/tunnel/admin/peers/shared)
	AdminCreateSharedPeer(w http.ResponseWriter, r *http.Request)
	// Delete peer
	// (DELETE /api/tunnel/admin/peers/{id})
	AdminDeletePeer(w http.ResponseWriter, r *http.Request, id int64)
	// Get peer info
	// (GET /api/tunnel/admin/peers/{id})
	AdminGetPeer(w http.ResponseWriter, r *http.Request, id int64)
	// Update peer
	// (PUT /api/tunnel/admin/peers/{id})
	AdminUpdatePeer(w http.ResponseWriter, r *http.Request, id int64)
	// Reloads service with new configuration
	// (GET /api/tunnel/admin/reload)
	AdminReloadService(w http.ResponseWriter, r *http.Request)
	// Get current server settings
	// (GET /api/tunnel/admin/settings)
	AdminGetSettings(w http.ResponseWriter, r *http.Request)
	// Update server settings
	// (PATCH /api/tunnel/admin/settings)
	AdminUpdateSettings(w http.ResponseWriter, r *http.Request)
	// Get current service status
	// (GET /api/tunnel/admin/status)
	AdminGetStatus(w http.ResponseWriter, r *http.Request)
	// Chech the shared peer status before the activation request
	// (GET /api/tunnel/public/activate-peer/{key})
	PublicPeerStatus(w http.ResponseWriter, r *http.Request, key string)
	// Activate the shared peer via the unique URL
	// (POST /api/tunnel/public/activate-peer/{key})
	PublicPeerActivate(w http.ResponseWriter, r *http.Request, key string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// AdminDoAuth operation middleware
func (siw *ServerInterfaceWrapper) AdminDoAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Basic_authScopes, []string{""})

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminDoAuth(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminConnectionInfoWireguard operation middleware
func (siw *ServerInterfaceWrapper) AdminConnectionInfoWireguard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminConnectionInfoWireguard(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminInitialSetup operation middleware
func (siw *ServerInterfaceWrapper) AdminInitialSetup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminInitialSetup(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminIppoolSuggest operation middleware
func (siw *ServerInterfaceWrapper) AdminIppoolSuggest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminIppoolSuggest(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminIppoolIsUsed operation middleware
func (siw *ServerInterfaceWrapper) AdminIppoolIsUsed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminIppoolIsUsed(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminListPeers operation middleware
func (siw *ServerInterfaceWrapper) AdminListPeers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminListPeers(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminCreatePeer operation middleware
func (siw *ServerInterfaceWrapper) AdminCreatePeer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminCreatePeer(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminCreateSharedPeer operation middleware
func (siw *ServerInterfaceWrapper) AdminCreateSharedPeer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminCreateSharedPeer(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminDeletePeer operation middleware
func (siw *ServerInterfaceWrapper) AdminDeletePeer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminDeletePeer(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminGetPeer operation middleware
func (siw *ServerInterfaceWrapper) AdminGetPeer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminGetPeer(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminUpdatePeer operation middleware
func (siw *ServerInterfaceWrapper) AdminUpdatePeer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminUpdatePeer(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminReloadService operation middleware
func (siw *ServerInterfaceWrapper) AdminReloadService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminReloadService(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminGetSettings operation middleware
func (siw *ServerInterfaceWrapper) AdminGetSettings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminGetSettings(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminUpdateSettings operation middleware
func (siw *ServerInterfaceWrapper) AdminUpdateSettings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminUpdateSettings(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AdminGetStatus operation middleware
func (siw *ServerInterfaceWrapper) AdminGetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AdminGetStatus(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PublicPeerStatus operation middleware
func (siw *ServerInterfaceWrapper) PublicPeerStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameter("simple", false, "key", chi.URLParam(r, "key"), &key)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PublicPeerStatus(w, r, key)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PublicPeerActivate operation middleware
func (siw *ServerInterfaceWrapper) PublicPeerActivate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameter("simple", false, "key", chi.URLParam(r, "key"), &key)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Token_authScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PublicPeerActivate(w, r, key)
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
		r.Get(options.BaseURL+"/api/tunnel/admin/auth", wrapper.AdminDoAuth)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/admin/connection-info/wireguard", wrapper.AdminConnectionInfoWireguard)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/tunnel/admin/initial-setup", wrapper.AdminInitialSetup)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/admin/ip-pool/suggest", wrapper.AdminIppoolSuggest)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/tunnel/admin/ip-pool/suggest", wrapper.AdminIppoolIsUsed)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/admin/peers", wrapper.AdminListPeers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/tunnel/admin/peers", wrapper.AdminCreatePeer)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/tunnel/admin/peers/shared", wrapper.AdminCreateSharedPeer)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/tunnel/admin/peers/{id}", wrapper.AdminDeletePeer)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/admin/peers/{id}", wrapper.AdminGetPeer)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/tunnel/admin/peers/{id}", wrapper.AdminUpdatePeer)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/admin/reload", wrapper.AdminReloadService)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/admin/settings", wrapper.AdminGetSettings)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/api/tunnel/admin/settings", wrapper.AdminUpdateSettings)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/admin/status", wrapper.AdminGetStatus)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/tunnel/public/activate-peer/{key}", wrapper.PublicPeerStatus)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/tunnel/public/activate-peer/{key}", wrapper.PublicPeerActivate)
	})

	return r
}
