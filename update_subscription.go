package mailchimp

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/splicers/go-mailchimp/status"
)

// UpdateSubscription is deprecated and you should use RawUpdateSubscription.
func (c *Client) UpdateSubscription(listID string, email string, status string, mergeFields map[string]interface{}) (*MemberResponse, error) {
	params := map[string]interface{}{
		"email_address": email,
		"status":        status,
		"merge_fields":  mergeFields,
	}

	return c.RawUpdateSubscription(listID, email, params)
}

// RawUpdateSubscription will update the subscription identified by `email` in
// the list with ID `listID`. You can send any parameters from the ones that
// are documented in Mailchimps docs.
//
// This is a better version of the deprecated UpdateSubscription.
//
// http://developer.mailchimp.com/documentation/mailchimp/reference/lists/members/#edit-put_lists_list_id_members_subscriber_hash
func (c *Client) RawUpdateSubscription(listID, email string, params map[string]interface{}) (*MemberResponse, error) {
	// Mailchimp uses the MD5 of the email as the subscription's key.
	emailMD5 := fmt.Sprintf("%x", md5.Sum([]byte(email)))

	// Default parameters for subscriptions
	reqParams := map[string]interface{}{
		"email_address": email,
		"status":        status.Subscribed,
	}

	// Override default parameters with whatever was sent in `params` argument.
	for k, v := range params {
		reqParams[k] = v
	}

	resp, err := c.do(
		"PUT",
		fmt.Sprintf("/lists/%s/members/%s", listID, emailMD5),
		&reqParams,
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

	// Allow any success status (2xx)
	if resp.StatusCode/100 == 2 {
		// Unmarshal response into MemberResponse struct
		memberResponse := new(MemberResponse)
		if err := json.Unmarshal(data, memberResponse); err != nil {
			return nil, err
		}
		return memberResponse, nil
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}
