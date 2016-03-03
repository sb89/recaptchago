package recaptcha

import (
  "testing"
  "reflect"
  "net/http/httptest"
  "net/http"
  "encoding/json"
  )

func TestSecretGetsSet(t *testing.T) {
  secretKey := "My Secret Key"

  r := New(secretKey)

  if r.secret != secretKey {
    t.Error("Expected %s, got %s", secretKey, r.secret)
  }
}

func TestDefaultTimeoutSetTo30(t *testing.T) {
  r := New("")

  if r.timeout != 30 {
    t.Error("Expected 30, got %d", r.timeout)
  }
}

func TestTimeoutCanBeSetManually(t *testing.T) {
  timeout := 80

  r := New("", Timeout(timeout))

  if r.timeout != timeout {
    t.Error("Expected %d, got %d", timeout, r.timeout)
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
      Success:true,
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
      Success: false,
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
