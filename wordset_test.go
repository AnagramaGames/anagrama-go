// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/anagrama-go"
	"github.com/stainless-sdks/anagrama-go/internal/testutil"
	"github.com/stainless-sdks/anagrama-go/option"
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
	client := anagramasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Wordsets.Get(
		context.TODO(),
		"ws__6fxAv1b",
		anagramasdk.WordsetGetParams{
			Category:  anagramasdk.String("category"),
			Limit:     anagramasdk.Int(1),
			MaxLength: anagramasdk.Int(1),
			MinLength: anagramasdk.Int(1),
			Random:    anagramasdk.WordsetGetParamsRandomTrue,
			Scramble:  anagramasdk.WordsetGetParamsScrambleTrue,
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
