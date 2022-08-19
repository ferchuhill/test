package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	// DefaultURL is the default endpoint that client.NewClient use.
	DefaultURL = "http://localhost:8080"

	// DefaultTimeout is the default timeout to be used for the http client.
	DefaultTimeout = time.Second * 5

	// DefaultDialer is the default dialer that is being used when new Form 3 clients are created.
	DefaultDialer = &net.Dialer{}

	// DefaultTransport is the default transport configuration to be used when new Form 3 clients are created.
	DefaultTransport = &http.Transport{
		DialContext:     DefaultDialer.DialContext,
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
)

// Client is the class use to to connect and use to communicate with the API
type Client struct {

	//the specific Form3 host
	endpoint string

	// Http client to make the request and recive the response
	httpClient http.Client
}

// NewClient creates new API client instance.
func NewClient() *Client {
	//create the client with the default form3 host and configuration
	c := Client{
		endpoint:   DefaultURL,
		httpClient: http.Client{Timeout: DefaultTimeout, Transport: DefaultTransport},
	}
	return &c
}

// GetURL returns the current URL settled in the client
func (c *Client) GetURL() string {
	return c.endpoint
}

// GetDefaultUrl returns the default Url of the client
func (c *Client) GetDefaultUrl() string {
	return DefaultURL
}

// SetURL when require to change the URL
func (c *Client) SetURL(url string) *Client {
	c.endpoint = url
	return c
}

// HttpClient return the the HttpClient of the client
func (c *Client) HttpClient() http.Client {
	return c.httpClient
}

// SetHttpClient set the a new HttpClient
func (c *Client) SetHttpClient(clientHttp http.Client) {
	c.httpClient = clientHttp
}

// CallAPI to call a specific path in the form3 endpoint. This receives everything to make the request and process the response:
// method: get|post|put|delete
// path: the resource to call
// reqBody: the body to send in the request
// resType: the struct to parse in the response
// if everything is correct the resType will have the data and return nil, otherwise, this function returns an error
func (c *Client) CallAPI(method, path string, reqBody, resType interface{}) error {

	//create a new Request
	req, err := c.newRequest(method, path, reqBody)
	if err != nil {
		return err
	}
	//process the request
	response, err := c.do(req)
	if err != nil {
		return err
	}
	//parse the response recived
	return c.parseResponse(response, resType)
}

// newRequest create and returns a new HTTP request
func (c *Client) newRequest(method, path string, reqBody interface{}) (*http.Request, error) {
	var body []byte
	var err error

	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
	}

	target := fmt.Sprintf("%s%s", c.endpoint, path)
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	c.httpClient.Timeout = DefaultTimeout

	return req, nil
}

// do sends an HTTP request and returns an HTTP response
func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// parseResponse check the response, if the status is ok (20X), and then parse the resType with the data recive in the body
func (c *Client) parseResponse(response *http.Response, resType interface{}) error {
	// Read all the response body
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// < 200 && >= 300 : API error
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		apiError := &APIError{Code: response.StatusCode}

		err = json.Unmarshal(body, &apiError)
		if err != nil {
			apiError.Message = string(body)
		}
		return apiError
	}

	// Nothing to unmarshal
	if len(body) == 0 || resType == nil {
		return nil
	}

	err = json.Unmarshal(body, &resType)
	return err
}
