package ncaa

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	DefaultBaseURL = &url.URL{
		Host:   "ncaa-api.henrygd.me",
		Scheme: "https",
		Path:   "",
	}
)

type Client struct {
	BaseUrl    *url.URL
	HTTPClient *http.Client
}

var nccaClientInstance *Client

func NewDefaultClient() *Client {
	if nccaClientInstance == nil {
		nccaClientInstance = &Client{
			BaseUrl:    DefaultBaseURL,
			HTTPClient: http.DefaultClient,
		}
	}

	return nccaClientInstance
}

func (c *Client) Do(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s status code: %d", req.URL.String(), res.StatusCode)
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) Get(endpoint string, query map[string]string) ([]byte, error) {
	// Build base http request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.BaseUrl.String(), endpoint), nil)
	if err != nil {
		return nil, err
	}

	// Add query params for request
	q := req.URL.Query()

	for key, val := range query {
		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()

	// Do the request
	b, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return b, nil
}
