package mailchimp

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// CheckSubscription ...
func (c *Client) CheckSubscription(email string, listID string) (*MemberResponse, error) {
	// Hash email
	h := md5.New()
	io.WriteString(h, email)
	emailMD5 := fmt.Sprintf("%s", h.Sum(nil))
	// Make request
	resp, err := c.do(
		"GET",
		fmt.Sprintf("/lists/%s/members/%s", listID, emailMD5),
		nil,
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
