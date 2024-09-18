package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	header := http.Header{}
	header.Set("Authorization", "ApiKey sampleKey123")

	apiKey, err := GetAPIKey(header)

	if apiKey != "sampleKey123" {
		t.Errorf("expected 'ApiKey sampleKey123', got '%s'", apiKey)
	}

	if err != nil {
		t.Errorf("error getting ApiKey %v", err)
	}

}
