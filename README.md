# recaptchago
Recaptcha package for Go

Install
----------------

```
go get github.com/sb89/recaptchago

```

Example
----------------

```
r := recpatcha.New("secret key", recaptcha.Timeout(40))  // Omit Timeout() if not required

success, err := r.Verify("ip address", r.PostFormValue("g-recaptcha-response"))

if !success {
  errors := r.GetErrors() // []string
  ....
}
```

Documentation
----------------
http://godoc.org/github.com/sb89/recaptchago
