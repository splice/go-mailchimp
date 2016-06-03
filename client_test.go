package mailchimp

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadAPIKey(t *testing.T) {
	client, err := NewClient("bogus", nil)
	assert.Nil(t, client)
	assert.Error(t, err)
}

func TestURL(t *testing.T) {
	client, err := NewClient("the_api_key-us13", nil)
	assert.NoError(t, err)

	expected, _ := url.Parse("https://us13.api.mailchimp.com/3.0")
	assert.Equal(t, expected, client.GetBaseURL())
}
