package mailchimp

import (
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
