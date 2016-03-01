package recaptcha

import "testing"

func TestSecretGetsSet(t *testing.T) {
  secretKey := "My Secret Key"

  r := New(secretKey)

  if r.secret != secretKey {
    t.Error("Expected %s, got %s", secretKey, r.secret)
  }
}

func TestDefaultTimeoutSetTo30(t *testing.T) {
  r := New("My Secret Key")

  if r.timeout != 30 {
    t.Error("Expected 30, got %d", r.timeout)
  }
}

func TestTimeoutCanBeSetManually(t *testing.T) {
  timeout := 80

  r := New("My Secret Key", Timeout(timeout))

  if r.timeout != timeout {
    t.Error("Expected %d, got %d", timeout, r.timeout)
  }
}

func TestExample(t *testing.T) {
  r := New("My Secret Key")

  t.Log(r.Verify("test", "test"))
}
