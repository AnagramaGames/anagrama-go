// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk

import (
	"context"
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

// Generate anagram puzzles with configurable difficulty. Returns scrambled words,
// hints, alternative solutions, and metadata for building interactive puzzle
// experiences.
//
// PuzzleService contains methods and other services that help with interacting
// with the anagrama API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPuzzleService] method instead.
type PuzzleService struct {
	Options []option.RequestOption
}

// NewPuzzleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPuzzleService(opts ...option.RequestOption) (r PuzzleService) {
	r = PuzzleService{}
	r.Options = opts
	return
}

// Generates a new anagram puzzle with a target word, scrambled letters, hints,
// alternative solutions, and metadata. Difficulty controls word length and hint
// availability. An optional seed allows reproducible puzzle generation. Requires a
// valid API key as a Bearer token.
func (r *PuzzleService) Generate(ctx context.Context, query PuzzleGenerateParams, opts ...option.RequestOption) (res *PuzzleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/puzzles/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A positional hint revealing one letter of the target word.
type PuzzleHint struct {
	// The letter revealed by this hint.
	Letter string `json:"letter" api:"required"`
	// The order in which this hint should be revealed (0-based).
	Order int64 `json:"order" api:"required"`
	// The zero-based index of the letter in the target word.
	Position int64 `json:"position" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Letter      respjson.Field
		Order       respjson.Field
		Position    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PuzzleHint) RawJSON() string { return r.JSON.raw }
func (r *PuzzleHint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the generated puzzle for analytics and difficulty calibration.
type PuzzleMetadata struct {
	// Number of alternative valid solutions.
	AltCount int64 `json:"altCount" api:"required"`
	// Whether the target word contains repeated letters.
	HasRepeats bool `json:"hasRepeats" api:"required"`
	// Number of unique letters in the letter pool.
	PoolSize int64 `json:"poolSize" api:"required"`
	// Number of vowels in the target word.
	VowelCount int64 `json:"vowelCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AltCount    respjson.Field
		HasRepeats  respjson.Field
		PoolSize    respjson.Field
		VowelCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PuzzleMetadata) RawJSON() string { return r.JSON.raw }
func (r *PuzzleMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A generated anagram puzzle with target word, scrambled letters, hints, and
// metadata.
type PuzzleResponse struct {
	// Alternative valid anagram solutions.
	Alts []string `json:"alts" api:"required"`
	// Definitions for the target word and alternatives. Keys are words, values are
	// definition strings or null if no definition is available.
	Definitions map[string]string `json:"definitions" api:"required"`
	// The difficulty level of the puzzle.
	//
	// Any of "easy", "medium", "hard".
	Difficulty PuzzleResponseDifficulty `json:"difficulty" api:"required"`
	// The feedback mode for guesses (e.g., "exact", "positional").
	FeedbackMode string `json:"feedbackMode" api:"required"`
	// Pre-generated hints revealing individual letters.
	Hints []PuzzleHint `json:"hints" api:"required"`
	// Maximum number of attempts allowed.
	MaxAttempts int64 `json:"maxAttempts" api:"required"`
	// Maximum number of hints available.
	MaxHints int64 `json:"maxHints" api:"required"`
	// Metadata about the generated puzzle for analytics and difficulty calibration.
	Metadata PuzzleMetadata `json:"metadata" api:"required"`
	// The sorted pool of available letters.
	Pool string `json:"pool" api:"required"`
	// The scrambled version of the target word.
	Scramble string `json:"scramble" api:"required"`
	// The seed used for puzzle generation. Can be passed back to reproduce this exact
	// puzzle.
	Seed string `json:"seed" api:"required"`
	// The target word the player must unscramble.
	Target string `json:"target" api:"required"`
	// The length of the target word.
	WordLength int64 `json:"wordLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alts         respjson.Field
		Definitions  respjson.Field
		Difficulty   respjson.Field
		FeedbackMode respjson.Field
		Hints        respjson.Field
		MaxAttempts  respjson.Field
		MaxHints     respjson.Field
		Metadata     respjson.Field
		Pool         respjson.Field
		Scramble     respjson.Field
		Seed         respjson.Field
		Target       respjson.Field
		WordLength   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PuzzleResponse) RawJSON() string { return r.JSON.raw }
func (r *PuzzleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The difficulty level of the puzzle.
type PuzzleResponseDifficulty string

const (
	PuzzleResponseDifficultyEasy   PuzzleResponseDifficulty = "easy"
	PuzzleResponseDifficultyMedium PuzzleResponseDifficulty = "medium"
	PuzzleResponseDifficultyHard   PuzzleResponseDifficulty = "hard"
)

type PuzzleGenerateParams struct {
	// Optional seed string for deterministic puzzle generation. The same seed and
	// difficulty always produce the same puzzle.
	Seed param.Opt[string] `query:"seed,omitzero" json:"-"`
	// Puzzle difficulty level. Controls word length, number of hints, and attempts
	// allowed.
	//
	// Any of "easy", "medium", "hard".
	Difficulty PuzzleGenerateParamsDifficulty `query:"difficulty,omitzero" json:"-"`
	// When `"true"`, reduces hints and attempts for an extra challenge.
	//
	// Any of "true", "false".
	HardMode PuzzleGenerateParamsHardMode `query:"hardMode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PuzzleGenerateParams]'s query parameters as `url.Values`.
func (r PuzzleGenerateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Puzzle difficulty level. Controls word length, number of hints, and attempts
// allowed.
type PuzzleGenerateParamsDifficulty string

const (
	PuzzleGenerateParamsDifficultyEasy   PuzzleGenerateParamsDifficulty = "easy"
	PuzzleGenerateParamsDifficultyMedium PuzzleGenerateParamsDifficulty = "medium"
	PuzzleGenerateParamsDifficultyHard   PuzzleGenerateParamsDifficulty = "hard"
)

// When `"true"`, reduces hints and attempts for an extra challenge.
type PuzzleGenerateParamsHardMode string

const (
	PuzzleGenerateParamsHardModeTrue  PuzzleGenerateParamsHardMode = "true"
	PuzzleGenerateParamsHardModeFalse PuzzleGenerateParamsHardMode = "false"
)
