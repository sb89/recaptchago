package recaptcha

import (
  "net/http"
  "net/url"
  "time"
)

const PostUrl string = "https://www.google.com/recaptcha/api/siteverify"

type Recaptcha struct {
  secret string
}

func New(secret string) *Recaptcha {
  return &Recaptcha{secret: secret}
}

func (recaptcha *Recaptcha) Verify(ipAddress string, response string) (bool, error) {
  client := http.Client{Timeout: 30 * time.Second}

  resp, err := client.PostForm(PostUrl, url.Values{"secret": {recaptcha.secret}, "response": {response}, "remoteip": {ipAddress}})
  if err != nil {
    return false, err
  }
  defer resp.Body.Close()

}
