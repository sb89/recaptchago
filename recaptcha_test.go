package recaptcha

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSecretGetsSet(t *testing.T) {
	secretKey := "My Secret Key"

	r := New(secretKey)

	if r.secret != secretKey {
		t.Errorf("Expected %s, got %s", secretKey, r.secret)
	}
}

func TestDefaultHTTPClientIsUsedIfNotSet(t *testing.T) {
	r := New("")

	if r.httpClient != http.DefaultClient {
		t.Error("Expected httpClient to be default but was not")
	}
}

func TestGetErrorsReturnsNilWhenNoErrors(t *testing.T) {
	r := New("")

	if r.GetErrors() != nil {
		t.Error("Expected nil, got ", r.GetErrors())
	}
}

func TestVerifySetsSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		json.NewEncoder(w).Encode(verifyResponse{
			Success: true,
		})
	}))
	defer server.Close()
	postUrl = server.URL

	r := New("")
	if success, _ := r.Verify("", ""); success != true {
		t.Error("Expected true, got false")
	}
}

func TestVerifySetsErrorCodes(t *testing.T) {
	errorCodes := []string{"Error1", "Error2"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		json.NewEncoder(w).Encode(verifyResponse{
			Success:    false,
			ErrorCodes: errorCodes,
		})
	}))
	defer server.Close()
	postUrl = server.URL

	r := New("")
	r.Verify("", "")

	if !reflect.DeepEqual(r.GetErrors(), errorCodes) {
		t.Error("Expected ", errorCodes, " got ", r.GetErrors())
	}
}
