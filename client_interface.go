package mailchimp

import (
	"net/url"
)

// ClientInterface defines exported methods
type ClientInterface interface {
	// Exported methods
	CheckSubscription(email string, listID string) (*MemberResponse, error)
	Subscribe(email string, listID string, mergeFields map[string]interface{}) (*MemberResponse, error)
	SetBaseURL(baseURL *url.URL)
	GetBaseURL() *url.URL
}
