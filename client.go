package mailchimp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/AreaHQ/mailchimp/status"
)

// Client manages communication with the Mailchimp API.
type Client struct {
	client  *http.Client
	baseURL *url.URL
	dc      string
	apiKey  string
}

// ErrorResponse ...
type ErrorResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

// Error ...
func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d %s (%s)", e.Status, e.Title, e.Detail)
}

// NewClient returns a new Mailchimp API client.  If a nil httpClient is
// provided, http.DefaultClient will be used. The apiKey must be in the format xyz-us11.
func NewClient(apiKey string, httpClient *http.Client) (ClientInterface, error) {
	if len(strings.Split(apiKey, "-")) != 2 {
		return nil, errors.New("Mailchimp API Key must be formatted like: xyz-zys")
	}
	dc := strings.Split(apiKey, "-")[1] // data center
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, err := url.Parse(fmt.Sprintf("https://%s.api.mailchimp.com/3.0", dc))
	if err != nil {
		return nil, err
	}
	return &Client{
		client:  httpClient,
		baseURL: baseURL,
		apiKey:  apiKey,
		dc:      dc,
	}, nil
}

// GetBaseURL ...
func (c *Client) GetBaseURL() *url.URL {
	return c.baseURL
}

// SetBaseURL ...
func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.baseURL = baseURL
}

// Subscribe ...
func (c *Client) Subscribe(email string, listID string) (*MemberResponse, error) {
	// Make request
	resp, err := c.do(
		"POST",
		fmt.Sprintf("/lists/%s/members/", listID),
		&map[string]string{
			"email_address": email,
			"status":        status.Subscribed,
		},
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// log.Print(string(data))

	// If the request failed
	if resp.StatusCode > 299 {
		errorResponse, err := extractError(data)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}

	// Unmarshal response into MemberResponse struct
	memberResponse := new(MemberResponse)
	if err := json.Unmarshal(data, memberResponse); err != nil {
		return nil, err
	}
	return memberResponse, nil
}

func (c *Client) do(method string, path string, body interface{}) (*http.Response, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	apiURL := fmt.Sprintf("%s%s", c.GetBaseURL(), path)

	req, err := http.NewRequest(method, apiURL, buf)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("", c.apiKey)

	return c.client.Do(req)
}

func extractError(data []byte) (*ErrorResponse, error) {
	errorResponse := new(ErrorResponse)
	if err := json.Unmarshal(data, errorResponse); err != nil {
		return nil, err
	}
	return errorResponse, nil
}
