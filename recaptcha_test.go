package recaptcha

import (
  "testing"
  "reflect"
  )

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

func TestGetErrorsReturnsNilWhenNoErrors(t *testing.T) {
  r := New("My Secret Key")

  if r.GetErrors() != nil {
    t.Error("Expected nil, got ", r.GetErrors())
  }
}

func TestGetErrorsReturnsErrors(t *testing.T) {
  errors := []string{"Test1", "Test2"}

  r := New("My Secret Key")
  r.errors = errors

  if !reflect.DeepEqual(errors, r.GetErrors()) {
    t.Error("Expected ", errors, " got ", r.GetErrors())
  }
}

func TestExample(t *testing.T) {
  r := New("My Secret Key")

  t.Log(r.Verify("test", "test"))
}
