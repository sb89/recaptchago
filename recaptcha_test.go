package recaptcha

import "testing"

func TestSecretGetsSet(t *testing.T) {
  secretKey := "My Secret Key"

  r := New(secretKey)

  if r.secret != secretKey {
    t.Error("Expected %s, got %s", secretKey, r.secret)
  }
}
