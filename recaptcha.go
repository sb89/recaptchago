package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var postUrl = "https://www.google.com/recaptcha/api/siteverify"

// Recaptcha struct stores the recaptcha configuration as well as the errors received
// from the Verify function.
type Recaptcha struct {
	httpClient *http.Client
	secret     string
	errors     []string
}

type verifyResponse struct {
	Success    bool
	ErrorCodes []string `json:"error-codes"`
}

func HTTPClient(c *http.Client) func(*Recaptcha) {
	return func(r *Recaptcha) {
		r.httpClient = c
	}
}

// New returns a new Recaptcha struct using specified secret and any additional options.
// Default timeout is 30 seconds.
func New(secret string, options ...func(*Recaptcha)) *Recaptcha {
	r := &Recaptcha{secret: secret}

	for _, option := range options {
		option(r)
	}

	if r.httpClient == nil {
		r.httpClient = http.DefaultClient
	}

	return r
}

// Verify the recaptcha response, will return true or false. Any errors received will be
// stored in recaptcha struct.
func (recaptcha *Recaptcha) Verify(ipAddress string, response string) (bool, error) {
	recaptcha.errors = nil
	client := http.Client{Timeout: 30 * time.Second}

	resp, err := client.PostForm(postUrl, url.Values{"secret": {recaptcha.secret}, "response": {response}, "remoteip": {ipAddress}})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	vr := new(verifyResponse)
	if err := json.Unmarshal(body, &vr); err != nil {
		return false, err
	}

	recaptcha.errors = vr.ErrorCodes

	return vr.Success, nil
}

// GetErrors returns the error that occurred during last recaptcha attempt.
func (recaptcha *Recaptcha) GetErrors() []string {
	return recaptcha.errors
}
