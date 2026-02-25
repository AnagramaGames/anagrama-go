// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/AnagramaGames/anagrama-go/internal/apijson"
	"github.com/AnagramaGames/anagrama-go/internal/apiquery"
	"github.com/AnagramaGames/anagrama-go/internal/requestconfig"
	"github.com/AnagramaGames/anagrama-go/option"
	"github.com/AnagramaGames/anagrama-go/packages/param"
	"github.com/AnagramaGames/anagrama-go/packages/respjson"
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
// calendar date. The word is selected from puzzle-quality words of 5-6 characters
// in length using a seeded random algorithm. An optional timezone parameter allows
// the date to be resolved in the caller's local timezone instead of UTC. Requires
// a valid API key as a Bearer token. Subject to daily rate limiting (1,000
// requests/day).
func (r *WordService) Daily(ctx context.Context, query WordDailyParams, opts ...option.RequestOption) (res *WordDailyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/words/daily"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a deterministic daily word that is the same for all users on a given
// calendar date. The word is selected from puzzle-quality words of 5-6 characters
// in length using a seeded random algorithm. An optional timezone parameter allows
// the date to be resolved in the caller's local timezone instead of UTC. Requires
// a valid API key as a Bearer token. Subject to daily rate limiting (1,000
// requests/day).
func (r *WordService) GetDaily(ctx context.Context, query WordGetDailyParams, opts ...option.RequestOption) (res *WordGetDailyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/words/daily"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns one or more random words from the Anagrama word pool. Words can be
// filtered by length. Requires a valid API key as a Bearer token. Subject to daily
// rate limiting (1,000 requests/day).
func (r *WordService) GetRandom(ctx context.Context, query WordGetRandomParams, opts ...option.RequestOption) (res *WordGetRandomResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/words/random"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns one or more random words from the Anagrama word pool. Words can be
// filtered by length. Requires a valid API key as a Bearer token. Subject to daily
// rate limiting (1,000 requests/day).
func (r *WordService) Random(ctx context.Context, query WordRandomParams, opts ...option.RequestOption) (res *WordRandomResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/words/random"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Checks whether a given word exists in the Anagrama dictionary for the specified
// language. Returns validation status and optional part-of-speech information.
// Requires a valid API key as a Bearer token.
func (r *WordService) Validate(ctx context.Context, query WordValidateParams, opts ...option.RequestOption) (res *WordValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/words/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WordDailyResponse struct {
	// The date (YYYY-MM-DD) for which this word was selected, resolved in the
	// requested timezone.
	Date time.Time `json:"date" api:"required" format:"date"`
	// The IANA timezone used to resolve the date.
	Timezone string `json:"timezone" api:"required"`
	// The daily word. Always 5-6 characters long.
	Word string `json:"word" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Timezone    respjson.Field
		Word        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WordDailyResponse) RawJSON() string { return r.JSON.raw }
func (r *WordDailyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WordGetDailyResponse struct {
	// The date (YYYY-MM-DD) for which this word was selected, resolved in the
	// requested timezone.
	Date time.Time `json:"date" api:"required" format:"date"`
	// The IANA timezone used to resolve the date.
	Timezone string `json:"timezone" api:"required"`
	// The daily word. Always 5-6 characters long.
	Word string `json:"word" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Timezone    respjson.Field
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
	Count int64 `json:"count" api:"required"`
	// Array of random words matching the length criteria.
	Words []string `json:"words" api:"required"`
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

type WordRandomResponse struct {
	// The number of words returned. May be 0 if no words match the length criteria.
	Count int64 `json:"count" api:"required"`
	// Array of random words matching the length criteria.
	Words []string `json:"words" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Words       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WordRandomResponse) RawJSON() string { return r.JSON.raw }
func (r *WordRandomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WordValidateResponse struct {
	// The language code used for validation.
	Lang string `json:"lang" api:"required"`
	// Whether the word exists in the dictionary for the specified language.
	Valid bool `json:"valid" api:"required"`
	// The word that was validated (lowercased).
	Word string `json:"word" api:"required"`
	// The primary part of speech for the word, if found (e.g., "noun", "verb",
	// "adjective").
	Pos string `json:"pos"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lang        respjson.Field
		Valid       respjson.Field
		Word        respjson.Field
		Pos         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WordValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *WordValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WordDailyParams struct {
	// IANA timezone identifier used to resolve the current calendar date (e.g.,
	// "America/New_York", "Europe/London"). Defaults to "UTC".
	Tz param.Opt[string] `query:"tz,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WordDailyParams]'s query parameters as `url.Values`.
func (r WordDailyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WordGetDailyParams struct {
	// IANA timezone identifier used to resolve the current calendar date (e.g.,
	// "America/New_York", "Europe/London"). Defaults to "UTC".
	Tz param.Opt[string] `query:"tz,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WordGetDailyParams]'s query parameters as `url.Values`.
func (r WordGetDailyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

type WordRandomParams struct {
	// Number of random words to return. Clamped to the range [1, 10]. Defaults to 1.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Maximum word length (inclusive). Clamped to the range [3, 15]. Defaults to 15.
	// If minLength > maxLength, they are swapped automatically.
	MaxLength param.Opt[int64] `query:"maxLength,omitzero" json:"-"`
	// Minimum word length (inclusive). Clamped to the range [3, 15]. Defaults to 3.
	MinLength param.Opt[int64] `query:"minLength,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WordRandomParams]'s query parameters as `url.Values`.
func (r WordRandomParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WordValidateParams struct {
	// The word to validate. Case-insensitive. Must be 1-100 characters.
	Word string `query:"word" api:"required" json:"-"`
	// ISO 639-1 language code to validate against. Defaults to "en" (English).
	Lang param.Opt[string] `query:"lang,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WordValidateParams]'s query parameters as `url.Values`.
func (r WordValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
