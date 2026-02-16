// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/AnagramaGames/anagrama-go"
	"github.com/AnagramaGames/anagrama-go/internal/testutil"
	"github.com/AnagramaGames/anagrama-go/option"
)

func TestDictionaryLanguages(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anagramasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Dictionary.Languages(context.TODO())
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDictionaryLookupWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anagramasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Dictionary.Lookup(
		context.TODO(),
		"x",
		anagramasdk.DictionaryLookupParams{
			Definitions: anagramasdk.Int(0),
			Exact:       anagramasdk.DictionaryLookupParamsExactTrue,
			Examples:    anagramasdk.DictionaryLookupParamsExamplesTrue,
			Fields:      anagramasdk.String("fields"),
			Lang:        anagramasdk.String("lang"),
			Limit:       anagramasdk.Int(1),
			Pos:         anagramasdk.String("pos"),
		},
	)
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
