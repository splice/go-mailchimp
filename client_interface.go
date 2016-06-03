package mailchimp

import (
	"net/url"
)

// ClientInterface defines exported methods
type ClientInterface interface {
	// Exported methods
	Subscribe(email string, listID string) (*MemberResponse, error)
	SetBaseURL(baseURL *url.URL)
	GetBaseURL() *url.URL
}
