// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagrama_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/anagrama-go"
	"github.com/stainless-sdks/anagrama-go/internal/testutil"
	"github.com/stainless-sdks/anagrama-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	pet, err := client.Pet.Update(context.TODO(), anagrama.PetUpdateParams{
		Pet: anagrama.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", pet.ID)
}
