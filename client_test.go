package mailchimp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadAPIKey(t *testing.T) {
	client, err := NewClient("asz", nil)
	assert.Nil(t, client)
	assert.Error(t, err)
}

func TestURL(t *testing.T) {
	client, err := NewClient("a-lit11", nil)
	assert.NoError(t, err)

	expected, _ := url.Parse("https://lit11.api.mailchimp.com/3.0")
	assert.Equal(t, expected, client.GetBaseURL())
}

func TestSubscribeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(500)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := NewClient("a-lit11", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	_, err = client.Subscribe("john@doe.com", "abc_test")
	assert.Equal(t, "Error 0  ()", err.Error())
}

func TestSubscribe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"email": "bob@example.com",
			"status": "sent",
			"reject_reason": "hard-bounce",
			"_id": "1"
		}`)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := NewClient("a-lit11", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	_, err = client.Subscribe("john@doe.com", "abc_test")
	assert.NoError(t, err)
}
