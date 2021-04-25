package httpclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

const BaseURL = "https://dev.assetspec.com/api/v1"

type Client struct {
	HTTPClient http.Client
	Token      string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
		HTTPClient: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("User-Agent", "terraform-provider-assetspec-dev")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	parsedReq, _ := httputil.DumpRequest(req, true)
	log.Printf("[DEBUG] %v\n", string(parsedReq))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
