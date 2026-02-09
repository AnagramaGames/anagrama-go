// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anagrama_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/anagrama-go"
	"github.com/stainless-sdks/anagrama-go/internal/testutil"
	"github.com/stainless-sdks/anagrama-go/option"
)

func TestCliAuthCompleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cli.Auth.Complete(context.TODO(), anagrama.CliAuthCompleteParams{
		BodyDeviceCode1: anagrama.String("device_code"),
		BodyDeviceCode2: anagrama.String("dvc_a1b2c3d4e5f6"),
		BodyUserCode1:   anagrama.String("user_code"),
		BodyUserCode2:   anagrama.String("ABCD-1234"),
	})
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCliAuthPollWithOptionalParams(t *testing.T) {
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
	_, err := client.Cli.Auth.Poll(context.TODO(), anagrama.CliAuthPollParams{
		BodyDeviceCode1: anagrama.String("dvc_a1b2c3d4e5f6"),
		BodyDeviceCode2: anagrama.String("deviceCode"),
		BodyUserCode1:   anagrama.String("user_code"),
		BodyUserCode2:   anagrama.String("userCode"),
	})
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCliAuthStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Cli.Auth.Start(context.TODO(), anagrama.CliAuthStartParams{
		ClientName: anagrama.String("anagrama-cli"),
		Label:      anagrama.String("My MacBook Pro"),
	})
	if err != nil {
		var apierr *anagrama.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
