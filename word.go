// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/anagrama-go/internal/apijson"
	"github.com/stainless-sdks/anagrama-go/internal/apiquery"
	"github.com/stainless-sdks/anagrama-go/internal/requestconfig"
	"github.com/stainless-sdks/anagrama-go/option"
	"github.com/stainless-sdks/anagrama-go/packages/param"
	"github.com/stainless-sdks/anagrama-go/packages/respjson"
)

// WordService contains methods and other services that help with interacting with
// the anagrama API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWordService] method instead.
type WordService struct {
	Options []option.RequestOption
}

// NewWordService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWordService(opts ...option.RequestOption) (r WordService) {
	r = WordService{}
	r.Options = opts
	return
}

// Returns a deterministic daily word that is the same for all users on a given
// calendar date (UTC). The word is selected from puzzle-quality words of 5-6
// characters in length using a seeded random algorithm. Requires a valid API key
// as a Bearer token. Subject to daily rate limiting (1,000 requests/day).
func (r *WordService) GetDaily(ctx context.Context, opts ...option.RequestOption) (res *WordGetDailyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/words/daily"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns one or more random words from the Anagrama word pool. Words can be
// filtered by length. Requires a valid API key as a Bearer token. Subject to daily
// rate limiting (1,000 requests/day).
func (r *WordService) GetRandom(ctx context.Context, query WordGetRandomParams, opts ...option.RequestOption) (res *WordGetRandomResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/words/random"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WordGetDailyResponse struct {
	// The UTC date (YYYY-MM-DD) for which this word was selected.
	Date time.Time `json:"date,required" format:"date"`
	// The daily word. Always 5-6 characters long.
	Word string `json:"word,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Word        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WordGetDailyResponse) RawJSON() string { return r.JSON.raw }
func (r *WordGetDailyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WordGetRandomResponse struct {
	// The number of words returned. May be 0 if no words match the length criteria.
	Count int64 `json:"count,required"`
	// Array of random words matching the length criteria.
	Words []string `json:"words,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Words       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WordGetRandomResponse) RawJSON() string { return r.JSON.raw }
func (r *WordGetRandomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WordGetRandomParams struct {
	// Number of random words to return. Clamped to the range [1, 10]. Defaults to 1.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Maximum word length (inclusive). Clamped to the range [3, 15]. Defaults to 15.
	// If minLength > maxLength, they are swapped automatically.
	MaxLength param.Opt[int64] `query:"maxLength,omitzero" json:"-"`
	// Minimum word length (inclusive). Clamped to the range [3, 15]. Defaults to 3.
	MinLength param.Opt[int64] `query:"minLength,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WordGetRandomParams]'s query parameters as `url.Values`.
func (r WordGetRandomParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
