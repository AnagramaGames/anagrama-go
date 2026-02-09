// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagramasdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/AnagramaGames/anagrama-go"
	"github.com/AnagramaGames/anagrama-go/internal/testutil"
	"github.com/AnagramaGames/anagrama-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	response, err := client.Words.GetRandom(context.TODO(), anagramasdk.WordGetRandomParams{
		Count:     anagramasdk.Int(3),
		MaxLength: anagramasdk.Int(8),
		MinLength: anagramasdk.Int(5),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Count)
}
