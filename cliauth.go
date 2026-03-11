// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/AnagramaGames/anagrama-go/internal/apijson"
	"github.com/AnagramaGames/anagrama-go/internal/requestconfig"
	"github.com/AnagramaGames/anagrama-go/option"
	"github.com/AnagramaGames/anagrama-go/packages/param"
	"github.com/AnagramaGames/anagrama-go/packages/respjson"
	"github.com/AnagramaGames/anagrama-go/shared/constant"
)

// CLI device-flow authentication. Used by the Anagrama CLI and SDKs to obtain an
// API token by having the user authorize via the web browser.
//
// CliAuthService contains methods and other services that help with interacting
// with the anagrama API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCliAuthService] method instead.
type CliAuthService struct {
	Options []option.RequestOption
}

// NewCliAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCliAuthService(opts ...option.RequestOption) (r CliAuthService) {
	r = CliAuthService{}
	r.Options = opts
	return
}

// Called by the web application after the user approves the CLI authentication
// request. This endpoint requires a valid Anagrama web session (browser
// cookie-based auth) and is not intended for direct SDK use.
//
// Approves the pending session identified by `deviceCode` or `userCode`, generates
// an API token, and stores it. The CLI can then retrieve the token by polling
// `/cli/auth/poll`.
func (r *CliAuthService) Complete(ctx context.Context, body CliAuthCompleteParams, opts ...option.RequestOption) (res *CliAuthCompleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cli/auth/complete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Polls the status of a device-flow authentication session. The CLI should call
// this endpoint at the interval specified in the `/cli/auth/start` response until
// it receives a status of `"approved"` with a token.
//
// The token is returned exactly once -- if the session has already been consumed
// by a previous successful poll, a `409` status is returned.
//
// Rate limited to 30 requests per minute per IP address.
func (r *CliAuthService) Poll(ctx context.Context, body CliAuthPollParams, opts ...option.RequestOption) (res *CliAuthPollResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "cli/auth/poll"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Initiates a device-flow authentication session. Returns a device code, a
// human-readable user code, and a verification URL. The CLI should display the
// verification URL and user code to the user, then poll `/cli/auth/poll` until the
// session is approved.
//
// Rate limited to 10 requests per minute per IP address.
func (r *CliAuthService) Start(ctx context.Context, body CliAuthStartParams, opts ...option.RequestOption) (res *CliAuthStartResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "cli/auth/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CliAuthCompleteResponse struct {
	// Always `true` on success.
	Ok bool `json:"ok" api:"required"`
	// The resulting session status.
	Status constant.Approved `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliAuthCompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CliAuthCompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliAuthPollResponse struct {
	// The API token (prefixed with `cli_`). Store this securely and use it as a Bearer
	// token for authenticated API calls.
	Token string `json:"token" api:"required"`
	// The session status. Always `"approved"` for a 200 response.
	Status constant.Approved `json:"status" api:"required"`
	// Basic profile information for the authenticated user.
	User CliAuthPollResponseUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Status      respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliAuthPollResponse) RawJSON() string { return r.JSON.raw }
func (r *CliAuthPollResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic profile information for the authenticated user.
type CliAuthPollResponseUser struct {
	// The user's display name. Falls back to username, first name, last name, or
	// "Player".
	DisplayName string `json:"displayName" api:"required"`
	// The user's unique identifier.
	UserID string `json:"userId" api:"required"`
	// The user's username, or null if not set.
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliAuthPollResponseUser) RawJSON() string { return r.JSON.raw }
func (r *CliAuthPollResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliAuthStartResponse struct {
	// Unique device code for this authentication session. Used when polling for
	// approval.
	DeviceCode string `json:"device_code" api:"required"`
	// Number of seconds until this authentication session expires. Default is 900 (15
	// minutes).
	ExpiresIn int64 `json:"expires_in" api:"required"`
	// Minimum number of seconds the CLI should wait between poll requests. Default
	// is 3.
	Interval int64 `json:"interval" api:"required"`
	// A short, human-readable code displayed to the user for verification.
	UserCode string `json:"user_code" api:"required"`
	// The full URL the user should visit to authorize the CLI. Includes the
	// device_code and user_code as query parameters.
	VerificationURL string `json:"verification_url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceCode      respjson.Field
		ExpiresIn       respjson.Field
		Interval        respjson.Field
		UserCode        respjson.Field
		VerificationURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliAuthStartResponse) RawJSON() string { return r.JSON.raw }
func (r *CliAuthStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliAuthCompleteParams struct {
	// The device code from the authentication session.
	DeviceCode param.Opt[string] `json:"device_code,omitzero"`
	// The user code from the authentication session.
	UserCode param.Opt[string] `json:"user_code,omitzero"`
	paramObj
}

func (r CliAuthCompleteParams) MarshalJSON() (data []byte, err error) {
	type shadow CliAuthCompleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CliAuthCompleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliAuthPollParams struct {
	// The device code returned from `/cli/auth/start`.
	DeviceCode param.Opt[string] `json:"device_code,omitzero"`
	// The user code returned from `/cli/auth/start`.
	UserCode param.Opt[string] `json:"user_code,omitzero"`
	paramObj
}

func (r CliAuthPollParams) MarshalJSON() (data []byte, err error) {
	type shadow CliAuthPollParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CliAuthPollParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliAuthStartParams struct {
	// The name of the CLI client application (e.g., "anagrama-cli").
	ClientName param.Opt[string] `json:"clientName,omitzero"`
	// A human-readable label for this token (e.g., "My Laptop"). HTML special
	// characters are stripped. Truncated to 100 characters.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r CliAuthStartParams) MarshalJSON() (data []byte, err error) {
	type shadow CliAuthStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CliAuthStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
