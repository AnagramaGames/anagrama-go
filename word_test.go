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

func TestWordDailyWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Words.Daily(context.TODO(), anagramasdk.WordDailyParams{
		Tz: anagramasdk.String("tz"),
	})
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWordGetDailyWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Words.GetDaily(context.TODO(), anagramasdk.WordGetDailyParams{
		Tz: anagramasdk.String("tz"),
	})
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWordGetRandomWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Words.GetRandom(context.TODO(), anagramasdk.WordGetRandomParams{
		Count:     anagramasdk.Int(1),
		MaxLength: anagramasdk.Int(3),
		MinLength: anagramasdk.Int(3),
	})
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWordRandomWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Words.Random(context.TODO(), anagramasdk.WordRandomParams{
		Count:     anagramasdk.Int(1),
		MaxLength: anagramasdk.Int(3),
		MinLength: anagramasdk.Int(3),
	})
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWordValidateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Words.Validate(context.TODO(), anagramasdk.WordValidateParams{
		Word: "x",
		Lang: anagramasdk.String("lang"),
	})
	if err != nil {
		var apierr *anagramasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
