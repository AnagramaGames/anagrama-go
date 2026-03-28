// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk

import (
	"context"
	"errors"
	"fmt"
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

// Access user-created word collections (wordsets). Public wordsets can be read
// without authentication. Private wordsets require the creator's API key.
//
// WordsetService contains methods and other services that help with interacting
// with the anagrama API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWordsetService] method instead.
type WordsetService struct {
	Options []option.RequestOption
}

// NewWordsetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWordsetService(opts ...option.RequestOption) (r WordsetService) {
	r = WordsetService{}
	r.Options = opts
	return
}

// Retrieves a user-created wordset by its ID. Public wordsets (`isPublic: true`)
// can be accessed without authentication. Private wordsets require the creator's
// API key passed as a Bearer token.
//
// Words in the response can be filtered, shuffled, limited, and optionally
// scrambled using query parameters. Each request increments the wordset's hit
// counter for analytics.
func (r *WordsetService) Get(ctx context.Context, wordsetID string, query WordsetGetParams, opts ...option.RequestOption) (res *WordsetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if wordsetID == "" {
		err = errors.New("missing required wordsetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wordsets/%s", url.PathEscape(wordsetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A wordset with its words, optionally filtered and/or scrambled.
type WordsetResponse struct {
	// ISO 8601 timestamp of when the wordset was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// A description of the wordset.
	Description string `json:"description" api:"required"`
	// Whether the wordset is publicly accessible.
	IsPublic bool `json:"isPublic" api:"required"`
	// The name of the wordset.
	Name string `json:"name" api:"required"`
	// The total number of words in the wordset (before filtering).
	TotalWordCount int64 `json:"totalWordCount" api:"required"`
	// The number of words in the `words` array (after filtering).
	WordCount int64 `json:"wordCount" api:"required"`
	// The list of words in the wordset, after applying any query filters (minLength,
	// maxLength, category, random, limit, scramble).
	Words []WordsetWord `json:"words" api:"required"`
	// The unique wordset identifier.
	WordsetID string `json:"wordsetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		Description    respjson.Field
		IsPublic       respjson.Field
		Name           respjson.Field
		TotalWordCount respjson.Field
		WordCount      respjson.Field
		Words          respjson.Field
		WordsetID      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WordsetResponse) RawJSON() string { return r.JSON.raw }
func (r *WordsetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A word entry in a wordset.
type WordsetWord struct {
	// The word string. If `scramble=true` was requested, the letters are shuffled.
	Word string `json:"word" api:"required"`
	// An optional category label for the word (e.g., "animals", "food").
	Category string `json:"category"`
	// An optional hint or clue for the word.
	Hint string `json:"hint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Word        respjson.Field
		Category    respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WordsetWord) RawJSON() string { return r.JSON.raw }
func (r *WordsetWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WordsetGetParams struct {
	// Filter words by category (case-insensitive match). Only words with a matching
	// `category` field are returned.
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Maximum number of words to return. Applied after filtering and optional
	// shuffling.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter words to only include those with at most this many characters.
	MaxLength param.Opt[int64] `query:"maxLength,omitzero" json:"-"`
	// Filter words to only include those with at least this many characters.
	MinLength param.Opt[int64] `query:"minLength,omitzero" json:"-"`
	// When set to `"true"`, shuffles the word list into a random order before applying
	// the limit.
	//
	// Any of "true", "false".
	Random WordsetGetParamsRandom `query:"random,omitzero" json:"-"`
	// When set to `"true"`, scrambles the letters of each word in the response
	// (Fisher-Yates shuffle on characters). Useful for building anagram puzzles.
	//
	// Any of "true", "false".
	Scramble WordsetGetParamsScramble `query:"scramble,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WordsetGetParams]'s query parameters as `url.Values`.
func (r WordsetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When set to `"true"`, shuffles the word list into a random order before applying
// the limit.
type WordsetGetParamsRandom string

const (
	WordsetGetParamsRandomTrue  WordsetGetParamsRandom = "true"
	WordsetGetParamsRandomFalse WordsetGetParamsRandom = "false"
)

// When set to `"true"`, scrambles the letters of each word in the response
// (Fisher-Yates shuffle on characters). Useful for building anagram puzzles.
type WordsetGetParamsScramble string

const (
	WordsetGetParamsScrambleTrue  WordsetGetParamsScramble = "true"
	WordsetGetParamsScrambleFalse WordsetGetParamsScramble = "false"
)
