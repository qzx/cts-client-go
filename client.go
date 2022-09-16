package cts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:8558"
const APIVersion string = "v1"

// Client -
type Client struct {
	HTTPClient *http.Client
	HostURL    string
	APIVersion string
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"token"`
}

// NewClient -
func NewClient(host, version *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default CTS URL and API Version
		HostURL:    HostURL,
		APIVersion: APIVersion,
	}

	if host != nil {
		c.HostURL = *host
	}

	if version != nil {
		c.APIVersion = *version
	}

	return &c, nil
}

func (c *Client) doRequestWithCode(req *http.Request, acceptedCodes []int) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ok = false
	for _, code := range acceptedCodes {
		if res.StatusCode == code {
			ok = true
		}
	}

	if !statusInList(res.StatusCode, acceptedCodes) {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func (c *Client) doRequestOK(req *http.Request) ([]byte, error) {
	return c.doRequestWithCode(req, []int{http.StatusOK})
}

func (c *Client) doRequestAccepted(req *http.Request) ([]byte, error) {
	return c.doRequestWithCode(req, []int{http.StatusAccepted})
}

func (c *Client) doRequestCreated(req *http.Request) ([]byte, error) {
	return c.doRequestWithCode(req, []int{http.StatusCreated})
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	return c.doRequestWithCode(req, []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
	})
}

func statusInList(c int, accepted []int) bool {
	for _, a := range accpted {
		if c == a {
			return true
		}
	}
	return false
}
