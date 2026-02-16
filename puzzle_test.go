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

func TestPuzzleGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Puzzles.Generate(context.TODO(), anagramasdk.PuzzleGenerateParams{
		Difficulty: anagramasdk.PuzzleGenerateParamsDifficultyEasy,
		HardMode:   anagramasdk.PuzzleGenerateParamsHardModeTrue,
		Seed:       anagramasdk.String("seed"),
	})
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
