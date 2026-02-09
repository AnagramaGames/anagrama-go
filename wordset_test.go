// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagrama_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/AnagramaGames/anagrama-go"
	"github.com/AnagramaGames/anagrama-go/internal/testutil"
	"github.com/AnagramaGames/anagrama-go/option"
)

func TestWordsetGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anagrama.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Wordsets.Get(
		context.TODO(),
		"ws__6fxAv1b",
		anagrama.WordsetGetParams{
			Category:  anagrama.String("category"),
			Limit:     anagrama.Int(1),
			MaxLength: anagrama.Int(1),
			MinLength: anagrama.Int(1),
			Random:    anagrama.WordsetGetParamsRandomTrue,
			Scramble:  anagrama.WordsetGetParamsScrambleTrue,
		},
	)
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
