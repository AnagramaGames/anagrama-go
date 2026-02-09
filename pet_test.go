// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagrama_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/anagrama-go"
	"github.com/stainless-sdks/anagrama-go/internal/testutil"
	"github.com/stainless-sdks/anagrama-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.New(context.TODO(), anagrama.PetNewParams{
		Pet: anagrama.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        anagrama.Int(10),
			Category: anagrama.PetCategoryParam{
				ID:   anagrama.Int(1),
				Name: anagrama.String("Dogs"),
			},
			Status: anagrama.PetStatusAvailable,
			Tags: []anagrama.PetTagParam{{
				ID:   anagrama.Int(0),
				Name: anagrama.String("name"),
			}},
		},
	})
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetGet(t *testing.T) {
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
	_, err := client.Pet.Get(context.TODO(), 0)
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.Update(context.TODO(), anagrama.PetUpdateParams{
		Pet: anagrama.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        anagrama.Int(10),
			Category: anagrama.PetCategoryParam{
				ID:   anagrama.Int(1),
				Name: anagrama.String("Dogs"),
			},
			Status: anagrama.PetStatusAvailable,
			Tags: []anagrama.PetTagParam{{
				ID:   anagrama.Int(0),
				Name: anagrama.String("name"),
			}},
		},
	})
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetDelete(t *testing.T) {
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
	err := client.Pet.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.FindByStatus(context.TODO(), anagrama.PetFindByStatusParams{
		Status: anagrama.PetFindByStatusParamsStatusAvailable,
	})
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.FindByTags(context.TODO(), anagrama.PetFindByTagsParams{
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithFormDataWithOptionalParams(t *testing.T) {
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
	err := client.Pet.UpdateWithFormData(
		context.TODO(),
		0,
		anagrama.PetUpdateWithFormDataParams{
			Name:   anagrama.String("name"),
			Status: anagrama.String("status"),
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

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.UploadImage(
		context.TODO(),
		0,
		io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		anagrama.PetUploadImageParams{
			AdditionalMetadata: anagrama.String("additionalMetadata"),
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
