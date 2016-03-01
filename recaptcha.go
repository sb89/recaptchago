package recaptcha

import (
  "net/http"
  "net/url"
  "io/ioutil"
  "time"
  "log"
)

const PostUrl string = "https://www.google.com/recaptcha/api/siteverify"

type Recaptcha struct {
  secret string
  timeout int
}

func Timeout(timeout int) func(*Recaptcha) {
  return func(r *Recaptcha) {
    r.timeout = timeout
  }
}

func New(secret string, options ...func(*Recaptcha)) *Recaptcha {
  r := &Recaptcha{secret: secret, timeout: 30}

  for _, option := range options {
    option(r)
  }

  return r
}

func (recaptcha *Recaptcha) Verify(ipAddress string, response string) (bool, error) {
  client := http.Client{Timeout: 30 * time.Second}

  resp, err := client.PostForm(PostUrl, url.Values{"secret": {recaptcha.secret}, "response": {response}, "remoteip": {ipAddress}})
  if err != nil {
    return false, err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return false, err
  }

  log.Println(string(body))

  return true, nil
}
