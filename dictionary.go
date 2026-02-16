// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/AnagramaGames/anagrama-go/internal/apijson"
	"github.com/AnagramaGames/anagrama-go/internal/apiquery"
	"github.com/AnagramaGames/anagrama-go/internal/requestconfig"
	"github.com/AnagramaGames/anagrama-go/option"
	"github.com/AnagramaGames/anagrama-go/packages/param"
	"github.com/AnagramaGames/anagrama-go/packages/respjson"
)

// DictionaryService contains methods and other services that help with interacting
// with the anagrama API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDictionaryService] method instead.
type DictionaryService struct {
	Options []option.RequestOption
}

// NewDictionaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDictionaryService(opts ...option.RequestOption) (r DictionaryService) {
	r = DictionaryService{}
	r.Options = opts
	return
}

// Returns a list of all languages available in the Anagrama dictionary, along with
// the number of entries for each language. Requires a valid API key as a Bearer
// token.
func (r *DictionaryService) Languages(ctx context.Context, opts ...option.RequestOption) (res *DictionaryLanguagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/dictionary/languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves detailed dictionary entries for a word, including definitions,
// etymology, part of speech, synonyms, antonyms, and related words. When `exact`
// is set to `false`, performs a fuzzy search and returns matching candidates.
// Requires a valid API key as a Bearer token.
func (r *DictionaryService) Lookup(ctx context.Context, word string, query DictionaryLookupParams, opts ...option.RequestOption) (res *DictionaryLookupResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if word == "" {
		err = errors.New("missing required word parameter")
		return
	}
	path := fmt.Sprintf("v1/dictionary/%s", url.PathEscape(word))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A single definition of a word with an optional usage example.
type DictionaryDefinition struct {
	// The definition text.
	Gloss string `json:"gloss,required"`
	// An example sentence demonstrating usage.
	Example string `json:"example"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gloss       respjson.Field
		Example     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DictionaryDefinition) RawJSON() string { return r.JSON.raw }
func (r *DictionaryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dictionary entry for a specific part of speech, including definitions,
// etymology, and related words.
type DictionaryEntry struct {
	// Part of speech (e.g., "noun", "verb", "adjective").
	Pos string `json:"pos,required"`
	// Words with opposite meaning.
	Antonyms []string `json:"antonyms"`
	// Definitions for this part of speech.
	Definitions []DictionaryDefinition `json:"definitions"`
	// Words derived from this word.
	Derived []string `json:"derived"`
	// The etymology (origin) of the word.
	Etymology string `json:"etymology"`
	// Inflected forms of the word (e.g., plural, past tense).
	Forms []string `json:"forms"`
	// More general terms (e.g., "animal" is a hypernym of "dog").
	Hypernyms []string `json:"hypernyms"`
	// More specific terms (e.g., "poodle" is a hyponym of "dog").
	Hyponyms []string `json:"hyponyms"`
	// Semantically related words.
	Related []string `json:"related"`
	// Words with similar meaning.
	Synonyms []string `json:"synonyms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pos         respjson.Field
		Antonyms    respjson.Field
		Definitions respjson.Field
		Derived     respjson.Field
		Etymology   respjson.Field
		Forms       respjson.Field
		Hypernyms   respjson.Field
		Hyponyms    respjson.Field
		Related     respjson.Field
		Synonyms    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DictionaryEntry) RawJSON() string { return r.JSON.raw }
func (r *DictionaryEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for an exact dictionary word lookup.
type DictionaryLookupResponse struct {
	// The language code used for the lookup.
	Lang string `json:"lang,required"`
	// Dictionary entries for the word, one per part of speech.
	Results []DictionaryEntry `json:"results,required"`
	// The word that was looked up.
	Word string `json:"word,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lang        respjson.Field
		Results     respjson.Field
		Word        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DictionaryLookupResponse) RawJSON() string { return r.JSON.raw }
func (r *DictionaryLookupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a fuzzy dictionary search.
type DictionarySearchResponse struct {
	// Total number of results returned.
	Count int64 `json:"count,required"`
	// Always `false` for search responses.
	Exact bool `json:"exact,required"`
	// The language code used for the search.
	Lang string `json:"lang,required"`
	// The original search query.
	Query string `json:"query,required"`
	// Matching words from the fuzzy search.
	Results []DictionarySearchResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Exact       respjson.Field
		Lang        respjson.Field
		Query       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DictionarySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *DictionarySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single result from a fuzzy dictionary search.
type DictionarySearchResult struct {
	// The primary part of speech.
	Pos string `json:"pos,required"`
	// The matching word.
	Word string `json:"word,required"`
	// A brief definition, or null if unavailable.
	Gloss string `json:"gloss,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pos         respjson.Field
		Word        respjson.Field
		Gloss       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DictionarySearchResult) RawJSON() string { return r.JSON.raw }
func (r *DictionarySearchResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about a supported dictionary language.
type LanguageInfo struct {
	// ISO 639-1 language code (e.g., "en", "es", "fr").
	Code string `json:"code,required"`
	// Number of dictionary entries available for this language.
	Entries int64 `json:"entries,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LanguageInfo) RawJSON() string { return r.JSON.raw }
func (r *LanguageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DictionaryLanguagesResponse struct {
	// Total number of available languages.
	Count int64 `json:"count,required"`
	// Array of available languages with entry counts.
	Languages []LanguageInfo `json:"languages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Languages   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DictionaryLanguagesResponse) RawJSON() string { return r.JSON.raw }
func (r *DictionaryLanguagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DictionaryLookupResponseUnion contains all possible properties and values from
// [DictionaryLookupResponse], [DictionarySearchResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DictionaryLookupResponseUnion struct {
	Lang string `json:"lang"`
	// This field is a union of [[]DictionaryEntry], [[]DictionarySearchResult]
	Results DictionaryLookupResponseUnionResults `json:"results"`
	// This field is from variant [DictionaryLookupResponse].
	Word string `json:"word"`
	// This field is from variant [DictionarySearchResponse].
	Count int64 `json:"count"`
	// This field is from variant [DictionarySearchResponse].
	Exact bool `json:"exact"`
	// This field is from variant [DictionarySearchResponse].
	Query string `json:"query"`
	JSON  struct {
		Lang    respjson.Field
		Results respjson.Field
		Word    respjson.Field
		Count   respjson.Field
		Exact   respjson.Field
		Query   respjson.Field
		raw     string
	} `json:"-"`
}

func (u DictionaryLookupResponseUnion) AsDictionaryLookupResponse() (v DictionaryLookupResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DictionaryLookupResponseUnion) AsDictionarySearchResponse() (v DictionarySearchResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DictionaryLookupResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *DictionaryLookupResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DictionaryLookupResponseUnionResults is an implicit subunion of
// [DictionaryLookupResponseUnion]. DictionaryLookupResponseUnionResults provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [DictionaryLookupResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfDictionaryEntryArray OfDictionarySearchResultArray]
type DictionaryLookupResponseUnionResults struct {
	// This field will be present if the value is a [[]DictionaryEntry] instead of an
	// object.
	OfDictionaryEntryArray []DictionaryEntry `json:",inline"`
	// This field will be present if the value is a [[]DictionarySearchResult] instead
	// of an object.
	OfDictionarySearchResultArray []DictionarySearchResult `json:",inline"`
	JSON                          struct {
		OfDictionaryEntryArray        respjson.Field
		OfDictionarySearchResultArray respjson.Field
		raw                           string
	} `json:"-"`
}

func (r *DictionaryLookupResponseUnionResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DictionaryLookupParams struct {
	// Maximum number of definitions to return per entry. Clamped to [0, 50].
	Definitions param.Opt[int64] `query:"definitions,omitzero" json:"-"`
	// Comma-separated list of fields to include in the response (e.g.,
	// "definitions,etymology,synonyms"). When omitted, all fields are returned.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// ISO 639-1 language code. Defaults to "en" (English).
	Lang param.Opt[string] `query:"lang,omitzero" json:"-"`
	// Maximum number of results to return for fuzzy searches. Clamped to [1, 50].
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter results by part of speech (e.g., "noun", "verb", "adjective").
	Pos param.Opt[string] `query:"pos,omitzero" json:"-"`
	// When `"true"`, performs an exact lookup. When `"false"`, performs a fuzzy search
	// returning similar words.
	//
	// Any of "true", "false".
	Exact DictionaryLookupParamsExact `query:"exact,omitzero" json:"-"`
	// When `"true"`, includes usage examples in definitions.
	//
	// Any of "true", "false".
	Examples DictionaryLookupParamsExamples `query:"examples,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DictionaryLookupParams]'s query parameters as `url.Values`.
func (r DictionaryLookupParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When `"true"`, performs an exact lookup. When `"false"`, performs a fuzzy search
// returning similar words.
type DictionaryLookupParamsExact string

const (
	DictionaryLookupParamsExactTrue  DictionaryLookupParamsExact = "true"
	DictionaryLookupParamsExactFalse DictionaryLookupParamsExact = "false"
)

// When `"true"`, includes usage examples in definitions.
type DictionaryLookupParamsExamples string

const (
	DictionaryLookupParamsExamplesTrue  DictionaryLookupParamsExamples = "true"
	DictionaryLookupParamsExamplesFalse DictionaryLookupParamsExamples = "false"
)
