package mininghamster

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BaseURL    = "https://www.mininghamster.com/api/v2/"
	DemoKey    = "288b2113-28ac-4b14-801f-f4d9cf9d87ad"
	TimeFormat = "2006-01-02 15:04:05"
)

type Client struct {
	URL string
	Key string
}

func New(apiKey string) *Client {
	return &Client{
		URL: BaseURL,
		Key: apiKey,
	}
}

func (client *Client) Get() ([]byte, error) {
	var err error

	// parse the mininghamster URL
	var endpoint *url.URL
	if endpoint, err = url.Parse(client.URL); err != nil {
		return nil, err
	}

	// set the endpoint for this request
	endpoint.Path += client.Key

	// create the request
	var req *http.Request
	if req, err = http.NewRequest("GET", endpoint.String(), nil); err != nil {
		return nil, err
	}

	// add signature to http header
	mac := hmac.New(sha512.New, []byte(client.Key))
	mac.Write([]byte(endpoint.String()))
	req.Header.Set("apisign", hex.EncodeToString(mac.Sum(nil)))

	// submit the http request
	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the body of the http message into a byte array
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return body, nil
}
