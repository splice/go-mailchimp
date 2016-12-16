package mailchimp_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mailchimp "github.com/nhocki/go-mailchimp"
	"github.com/nhocki/go-mailchimp/status"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSubscriptionError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(400)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, invalidMergeFieldsErrorResponse)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := mailchimp.NewClient("the_api_key-us13", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	memberResponse, err := client.Subscribe("list_id", "john@reese.com", map[string]interface{}{})
	assert.Nil(t, memberResponse)
	assert.Equal(t, "Error 400 Invalid Resource (Your merge fields were invalid.)", err.Error())

	errResponse, ok := err.(*mailchimp.ErrorResponse)
	assert.True(t, ok)
	assert.Equal(t, "Invalid Resource", errResponse.Title)
	assert.Equal(t, 400, errResponse.Status)
	assert.Equal(t, "Your merge fields were invalid.", errResponse.Detail)
}

func TestUpdateSubscriptionMalformedError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(500)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := mailchimp.NewClient("the_api_key-us13", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	memberResponse, err := client.UpdateSubscription("list_id", "john@reese.com", "subscribed", map[string]interface{}{})
	assert.Nil(t, memberResponse)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

func TestUpdateSubscription(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, successResponse)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := mailchimp.NewClient("the_api_key-us13", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	memberResponse, err := client.Subscribe("list_id", "john@reese.com", map[string]interface{}{})
	assert.NoError(t, err)

	assert.Equal(t, "11bf13d1eb58116eba1de370b2bd796b", memberResponse.ID)
	assert.Equal(t, "john@reese.com", memberResponse.EmailAddress)
	assert.Equal(t, status.Subscribed, memberResponse.Status)
}
