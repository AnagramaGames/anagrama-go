// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagrama

import (
	"github.com/AnagramaGames/anagrama-go/option"
)

// CliService contains methods and other services that help with interacting with
// the anagrama API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCliService] method instead.
type CliService struct {
	Options []option.RequestOption
	Auth    CliAuthService
}

// NewCliService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCliService(opts ...option.RequestOption) (r CliService) {
	r = CliService{}
	r.Options = opts
	r.Auth = NewCliAuthService(opts...)
	return
}
